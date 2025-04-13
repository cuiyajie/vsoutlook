<script lang="ts" setup>
import { def_switch_bus_select, type SwitchBusLevelParams, type SwitchInputVideoParams } from "./Consts";
import { type InjectSwitchBusSelect } from "../SwitchShare/Consts";

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
const { selectedBus, busSelected, busUnSelected } = inject<InjectSwitchBusSelect>('switch_bus_select', def_switch_bus_select())

const options = computed(() => {
  if (props.type === 'video') {
    return props.inputVideos.map(item => ({
      signal_name: item.signal_name,
      signal_id: item.signal_id
    }))
  } else if (props.type === 'bus_output') {
    const isDisabled = (bid: string) => selectedBus.value.has(bid) && bid !== signalId.value
    return props.busLevels.reduce((acc, item) => {
      item.pgm_bus.out_signal?.forEach(signal => {
        if (props.outType === 'pgm' && signal.signal_type === 'final') {
          acc.push({
            signal_name: signal.signal_name,
            signal_id: signal.signal_id,
            disabled: isDisabled(signal.signal_id)
          })
        } else if (props.outType === 'clean' && signal.signal_type === 'clean') {
          acc.push({
            signal_name: signal.signal_name,
            signal_id: signal.signal_id,
            disabled: isDisabled(signal.signal_id)
          })
        }
      })
      item.pvw_bus.out_signal?.forEach(signal => {
        if (props.outType === 'pvw' && signal.signal_type === 'final') {
          acc.push({
            signal_name: signal.signal_name,
            signal_id: signal.signal_id,
            disabled: isDisabled(signal.signal_id)
          })
        }
      })
      return acc
    }, [] as { signal_name: string, signal_id: string, disabled?: boolean }[])
  }
  return []
})

watch(() => props.type, () => {
  signalId.value = ''
  signalName.value = ''
})

watch(signalId, (nv, ov) => {
  if (nv === ov) return
  if (nv && nv.startsWith('level')) {
    busSelected(nv)
  }
  if (ov && ov.startsWith('level')) {
    busUnSelected(ov)
  }
}, { immediate: true })

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
    :style="{'--ms-max-height': '245px'}"
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
