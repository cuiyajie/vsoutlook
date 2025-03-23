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
const audioFormats = computed(() => usStore.settings.audio_formats || [])

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
    :max-height="145"
    no-options-text="暂时没有配置音频格式名称"
    :options="audioFormats"
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
