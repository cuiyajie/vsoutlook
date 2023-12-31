import type { WatchSource, WatchOptions } from "vue";
import { defs, formats, v_compression_format, v_compression_ratio, v_protocols, v_width, val_udx } from './Consts';
import omit from 'lodash-es/omit';

function setFormat(dst: any, tokens: string[]) {
  Object.assign(dst, {
    'v_width': +tokens[0],
    'v_height': +tokens[1],
    'v_interlaced': Boolean(+tokens[2]),
    'v_fps': +tokens[3],
    'v_gamma': tokens[4],
    'v_gamut': tokens[5],
  });
}

export function watchFormat(format: WatchSource<string>, mv: any) {
  watch(
    format,
    (nv) => {
      const tokens = nv.split("_");
      if (mv) {
        setFormat(mv.videoformat, tokens);
      } else {
        mv.videoformat = omit(mv.videoformat, 'v_width', 'v_height', 'v_interlaced', 'v_fps', 'v_gamma', 'v_gamut');
      }
    },
    { immediate: true }
  );
}

export function getFormat(mv: any) {
  const vf = mv.videoformat;
  return `${vf.v_width}_${vf.v_height}_${Number(vf.v_interlaced)}_${vf.v_fps}_${vf.v_gamma}_${vf.v_gamut}`;
}

export function watchFormat2(format: WatchSource<string>, mvs: any[]) {
  watch(
    format,
    (nv) => {
      const tokens = nv.split("_");
      mvs.forEach(m => {
        if (m) {
          setFormat(m.videoformat, tokens);
        } else {
          m.videoformat = omit(m.videoformat, 'v_width', 'v_height', 'v_interlaced', 'v_fps', 'v_gamma', 'v_gamut');
        }
      })
    },
    { immediate: true }
  );
}

export function watchProtocol(mv: any) {
  watch(
    () => mv.v_protocol,
    (nv) => {
      if (nv === v_protocols[1]) {
        mv.videoformat.v_compression_format = v_compression_format;
      } else {
        mv.videoformat.v_compression_format = null;
      }
    },
    { immediate: true }
  );
}

export function watchCpressFormat(mv: any) {
  watch(
    [() => mv.videoformat.v_compression_format, () => mv.videoformat.v_width],
    ([df, w]) => {
      if (df === v_compression_format) {
        mv.videoformat.v_compression_ratio = String(w) === v_width[0] ? v_compression_ratio[0] : v_compression_ratio[1];
      } else {
        mv.videoformat.v_compression_ratio = null;
      }
    },
    { immediate: true }
  );
}

export function useFormat(mv: any, def = defs.format) {
  const v_format = ref(def);
  watchFormat(v_format, mv);
  return v_format;
}

export function useProtocolDC(mv: any) {
  watchProtocol(mv);
  watchCpressFormat(mv);
}

export function watchInput(key: string, input: any, mvs: any[], options: WatchOptions  = {}) {
  watch(
    () => input[key],
    (nv) => {
      mvs.forEach(m => {
        Object.assign(m[key], nv)
      });
    },
    { immediate: true, ...options }
  );
}

export function watchModeVFormat(
  watcher: WatchSource<'upscale' | 'downscale'>,
  shouldFormats: Ref<any[]>
) {
  watch(
    watcher,
    (nv) => {
      if (nv === val_udx[0]) {
        shouldFormats.value = formats.slice(0, 2)
      } else {
        shouldFormats.value = formats.slice(-1);
      }
    },
    { immediate: true }
  );
}

export function wrap(src: any, prefix: string, useBackup?: boolean, isBackup?: boolean, m_local_ip?: string, b_local_ip?: string) {
  const dst: any = {}
  for (const key in src) {
    let wrapKey = `${prefix}${key}`;
    if ([
      'ipstream_master', 'ipstream_backup', 'videoformat', 'audioformat', 'pip_params',
      'input_key_params', 'input_fill_params', 'input_video_params',
      'video_bus_master', 'video_bus_backup', 'keyfill_bus_master', 'keyfill_bus_backup',
    ].includes(key)) {
      wrapKey = key
    }
    if (src[key] instanceof Array) {
      const isBackup = key.endsWith('backup')
      if (isBackup && !useBackup) {
        continue
      } else {
        dst[wrapKey] = src[key].map((v: any) => wrap(v, prefix, useBackup, isBackup, m_local_ip, b_local_ip));
      }
    } else if (typeof src[key] === 'object') {
      const isBackup = key.endsWith('backup')
      if (isBackup && !useBackup) {
        continue
      } else {
        dst[wrapKey] = wrap(src[key], prefix, useBackup, isBackup, m_local_ip, b_local_ip);
      }
      if (Object.keys(dst[wrapKey]).length === 0) {
        dst[wrapKey] = null;
      }
    } else {
      if (prefix === 'out_' && key.endsWith('src_address') && isBackup !== undefined) {
        dst[wrapKey] = `${isBackup ? b_local_ip : m_local_ip}:${src[key]}`;
      } else {
        dst[wrapKey] = src[key];
      }
    }
  }
  return dst
}

export function unwrap(src: any, prefix: string) {
  const dst: any = {}
  for (const key in src) {
    const unwrapKey = key.replace(prefix, '');
    if (src[key] instanceof Array) {
      dst[unwrapKey] = src[key].map((v: any) => unwrap(v, prefix));
    } else if (typeof src[key] === 'object') {
      if (src[key] === null) {
        dst[unwrapKey] = null
      } else {
        dst[unwrapKey] = unwrap(src[key], prefix);
      }
    } else {
      if (prefix === 'out_' && key.endsWith('src_address')) {
        dst[unwrapKey] = src[key].split(':')?.[1] || '';
      } else {
        dst[unwrapKey] = src[key];
      }
    }
  }
  return dst
}
