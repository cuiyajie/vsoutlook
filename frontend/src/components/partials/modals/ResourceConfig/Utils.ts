import type { WatchSource, WatchOptions } from "vue";
import { defs, formats, v_compression_format, v_compression_ratio, v_protocols, v_width, val_udx } from './Consts';

export function watchFormat(format: WatchSource<string>, mv: any) {
  watch(
    format,
    (nv) => {
      const tokens = nv.split("_");
      if (!mv) {
        debugger
      }
      Object.assign(mv.videoformat, {
        'v_width': tokens[0],
        'v_height': tokens[1],
        'v_interlaced': tokens[2],
        'v_fps': tokens[3],
        'v_gamma': tokens[4],
        'v_color_gamut': tokens[5],
      });
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
        Object.assign(m.videoformat, {
          'v_width': tokens[0],
          'v_height': tokens[1],
          'v_interlaced': tokens[2],
          'v_fps': tokens[3],
          'v_gamma': tokens[4],
          'v_color_gamut': tokens[5],
        });
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

export function wrap(src: any, prefix: string) {
  const dst: any = {}
  for (const key in src) {
    const wrapKey = `${prefix}${key}`;
    if (src[key] instanceof Array) {
      dst[wrapKey] = src[key].map((v: any) => wrap(v, prefix));
    } else if (typeof src[key] === 'object') {
      dst[wrapKey] = wrap(src[key], prefix);
    } else {
      dst[wrapKey] = src[key];
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
      dst[unwrapKey] = unwrap(src[key], prefix);
    } else {
      if (prefix === 'out_' && key.endsWith('dst_address')) {
        dst[unwrapKey] = src[key].split(':')?.[1] || '';
        continue;
      } else {
        dst[unwrapKey] = src[key];
      }
    }
  }
  return dst
}
