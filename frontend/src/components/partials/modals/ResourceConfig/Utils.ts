import type { WatchSource } from "vue";
import { defs, formats } from "./Consts"

export function watchFormat(format: WatchSource<string>, mv: Ref<any>) {
  watch(
    format,
    (nv) => {
      const tokens = nv.split("_");
      Object.assign(mv.value, {
        'V_Width': tokens[0],
        'V_Height': tokens[1],
        'V_Interlaced': tokens[2],
        'V_FPS': tokens[3],
        'V_Gamma': tokens[4],
        'V_ColorGamut': tokens[5],
      });
    },
    { immediate: true }
  );
}

export function watchFormat2(format: WatchSource<string>, mvs: Ref<any>[]) {
  watch(
    format,
    (nv) => {
      const tokens = nv.split("_");
      mvs.forEach(m => {
        Object.assign(m.value, {
          'V_Width': tokens[0],
          'V_Height': tokens[1],
          'V_Interlaced': tokens[2],
          'V_FPS': tokens[3],
          'V_Gamma': tokens[4],
          'V_ColorGamut': tokens[5],
        });
      })
    },
    { immediate: true }
  );
}

export function watchProtocol(mv: Ref<any>) {
  watch(
    () => mv.value.V_Protocol,
    (nv) => {
      if (nv === "ST2110-22") {
        mv.value.V_DecFormat = "JPEG-XS";
      } else {
        mv.value.V_DecFormat = "";
      }
    },
    { immediate: true }
  );
}

export function watchDecFormat(mv: Ref<any>) {
  watch(
    [() => mv.value.V_DecFormat, () => mv.value.V_Width],
    ([df, w]) => {
      if (df === "JPEG-XS") {
        mv.value.V_CompressionRatio = w === "1920" ? "5:1" : "8:1";
      } else {
        mv.value.V_CompressionRatio = "";
      }
    },
    { immediate: true }
  );
}

export function useFormat(mv: Ref<any>, def = defs.format) {
  const V_Format = ref(def);
  watchFormat(V_Format, mv);
  return V_Format;
}

export function useProtocolDC(mv: Ref<any>) {
  watchProtocol(mv);
  watchDecFormat(mv);
}

export function watchInput(watchKeys: string[], input: Ref<any>, mvs: Ref<any>[]) {
  watchKeys.forEach(key => {
    watch(
      () => input.value[key],
      (nv) => {
        mvs.forEach(m => {
          m.value[key] = nv
        });
      },
      { immediate: true }
    );
  });
}

export function watchModeVFormat(
  watcher: WatchSource<'UpScale' | 'DownScale'>,
  shouldFormats: Ref<any[]>
) {
  watch(
    watcher,
    (nv) => {
      if (nv === "UpScale") {
        shouldFormats.value = formats.slice(0, 2)
      } else {
        shouldFormats.value = formats.slice(-1);
      }
    },
    { immediate: true }
  );
}
