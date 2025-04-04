<script lang="ts" setup>
import { type SwitchBusLevelParams, type SwitchInputVideoParams } from "./Consts";

const signalId = defineModel('id', {
  default: "",
  local: true,
});

const signalName = defineModel('name', {
  default: "",
  local: true,
});

const props = defineProps<{
  title: string,
  type: string,
  outType: string,
  inputVideos: SwitchInputVideoParams[],
  busLevels: SwitchBusLevelParams[]
}>();

const options = computed(() => {
  if (props.type === 'video') {
    return props.inputVideos.map(item => ({
      signal_name: item.signal_name,
      signal_id: item.signal_id
    }))
  } else if (props.type === 'bus_output') {
    return props.busLevels.reduce((acc, item) => {
      item.pgm_bus.out_signal?.forEach(signal => {
        if (props.outType === 'pgm' && signal.signal_type === 'final') {
          acc.push({
            signal_name: signal.signal_name,
            signal_id: signal.signal_id
          })
        } else if (props.outType === 'clean' && signal.signal_type === 'clean') {
          acc.push({
            signal_name: signal.signal_name,
            signal_id: signal.signal_id
          })
        }
      })
      item.pvw_bus.out_signal?.forEach(signal => {
        if (props.outType === 'pvw' && signal.signal_type === 'final') {
          acc.push({
            signal_name: signal.signal_name,
            signal_id: signal.signal_id
          })
        }
      })
      return acc
    }, [] as { signal_name: string, signal_id: string }[])
  }
  return []
})

watch(() => props.type, () => {
  signalId.value = ''
  signalName.value = ''
})

function onSignalSelect(value: string) {
  signalId.value = value
  signalName.value = options.value.find(item => item.signal_id === value)?.signal_name || ''
}

</script>
<template>
  <Multiselect
    v-model="signalId"
    class="tippy-select"
    :placeholder="`选择${title}`"
    value-prop="signal_id"
    label="signal_name"
    :can-deselect="false"
    :max-height="145"
    :no-options-text="`暂时没有配置${title}`"
    :options="options"
    @select="onSignalSelect"
  >
    <template #option="{ option }">
      <span class="select-option-text">
        {{ option.signal_name }}
      </span>
    </template>
  </Multiselect>
</template>
