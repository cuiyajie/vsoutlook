<script lang="ts" setup>
import { useUserSession } from "@src/stores/userSession";

const mv = defineModel<string>({
  default: "",
  local: true,
});

const emit = defineEmits<{
  (e: 'audio-selected', name: string): void,
  (e: 'audio-unselected', name: string): void
}>()

const usStore = useUserSession()
const videoFormats = computed(() => usStore.settings.video_formats || [])
const audioFormats = computed(() => usStore.settings.audio_formats || [])
const props = defineProps<{
  videoformat: string,
}>()

const filterAudioFormats = computed(() => {
  const vf = videoFormats.value.find(vf => vf.name === props.videoformat)
  const isSt = ['st2110-20', 'st2110-22'].includes(vf?.protocol || '')
  if (isSt) {
    return audioFormats.value.filter(af => af.compression_format === 'pcm')
  } else {
    return audioFormats.value
  }
})

watch<string>(() => mv.value, (nv, ov) => {
  if (nv === ov) return
  if (nv) {
    emit('audio-selected', nv)
  } else {
    emit('audio-unselected', ov)
  }
})
</script>
<template>
  <Multiselect
    v-model="mv"
    class="tippy-select"
    placeholder="选择音频格式名称"
    value-prop="name"
    label="name"
    :can-deselect="false"
    :style="{'--ms-max-height': '245px'}"
    no-options-text="没有可选择的音频格式"
    :options="filterAudioFormats"
  >
    <template #option="{ option }">
      <span class="select-option-text">
        {{ option.name }}
      </span>
      <AudioFormatTooltip :value="option">
        <i class="iconify" data-icon="feather:alert-circle" aria-hidden="true"></i>
      </AudioFormatTooltip>
    </template>
  </Multiselect>
</template>
