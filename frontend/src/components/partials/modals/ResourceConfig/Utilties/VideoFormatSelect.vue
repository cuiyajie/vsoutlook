<script lang="ts" setup>
import { useUserSession } from "@src/stores/userSession";
import { VideoFormatProtocols } from '@src/components/partials/modals/PresetConfig/Consts';

const mv = defineModel<string>({
  default: "",
  local: true,
});

const props = defineProps<{
  usedSignalType: number,
}>()

const emit = defineEmits<{
  (e: 'video-selected', name: string): void,
  (e: 'video-unselected', name: string): void
}>()

const usStore = useUserSession()
const videoFormats = computed(() => usStore.settings.video_formats || [])

const filteredProtocols = computed(() => {
  switch(props.usedSignalType) {
    case 0:
      return ['st2110-20', 'st2110-22']
    case 1:
      return VideoFormatProtocols
        .filter(vfp => !['file', 'st2110-20', 'st2110-22'].includes(vfp.value))
        .map(vfp => vfp.value)
    case 2:
      return VideoFormatProtocols
        .filter(vfp => vfp.value !== 'file')
        .map(vfp => vfp.value)
    default:
      return [];
  }
});

const filterVideoFormats = computed(() => {
  return videoFormats.value.filter(vf => filteredProtocols.value.includes(vf.protocol))
})

watch(() => mv.value, (nv, ov) => {
  if (nv === ov) return
  if (nv) {
    emit('video-selected', nv)
  }
  if (ov) {
    emit('video-unselected', ov)
  }
}, { immediate: true })
</script>
<template>
  <Multiselect
    v-model="mv"
    class="tippy-select"
    placeholder="选择视频格式名称"
    value-prop="name"
    label="name"
    :can-deselect="false"
    :style="{'--ms-max-height': '245px'}"
    no-options-text="暂时没有配置视频格式名称"
    :options="filterVideoFormats"
  >
    <template #option="{ option }">
      <span class="select-option-text">
        {{ option.name }}
      </span>
      <VideoFormatTooltip :value="option">
        <i class="iconify" data-icon="feather:alert-circle" aria-hidden="true"></i>
      </VideoFormatTooltip>
    </template>
  </Multiselect>
</template>
