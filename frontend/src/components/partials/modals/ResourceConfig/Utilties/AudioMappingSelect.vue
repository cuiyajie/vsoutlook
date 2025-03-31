<script lang="ts" setup>
import { useUserSession } from "@src/stores/userSession";

const mv = defineModel<string>({
  default: "",
  local: true,
});

const emit = defineEmits<{
  (e: 'mapping-selected', name: string): void,
  (e: 'mapping-unselected', name: string): void
}>()

const usStore = useUserSession()
const audioMappings = computed(() => usStore.settings.audio_mappings || [])

watch<string>(() => mv.value, (nv, ov) => {
  if (nv === ov) return
  if (nv) {
    emit('mapping-selected', nv)
  } else {
    emit('mapping-unselected', ov)
  }
})
</script>
<template>
  <Multiselect
    v-model="mv"
    class="tippy-select"
    placeholder="选择音频映射模板"
    value-prop="name"
    label="name"
    :can-deselect="false"
    :max-height="145"
    no-options-text="暂时没有配置音频映射模板"
    :options="audioMappings"
  >
    <template #option="{ option }">
      <span class="select-option-text">
        {{ option.name }}
      </span>
      <AudioMappingTooltip :value="option">
        <i class="iconify" data-icon="feather:alert-circle" aria-hidden="true"></i>
      </AudioMappingTooltip>
    </template>
  </Multiselect>
</template>
