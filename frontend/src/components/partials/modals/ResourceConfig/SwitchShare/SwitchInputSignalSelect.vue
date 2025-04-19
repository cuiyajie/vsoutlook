<script lang="ts" setup>
import { type SwitchInputKeyParams, type SwitchBusLevelParams, type SwitchInputVideoParams, def_switch_input_signal_select, InjectSwitchInputSignalSelect, InjectSwitchKeySignalSelect, def_switch_key_signal_select } from "./Consts";

const signalId = defineModel('id', {
  default: "",
  local: true,
});

const signalName = defineModel('name', {
  default: "",
  local: true,
});

const props =withDefaults(defineProps<{
  title: string,
  type: string,
  inputKeys: SwitchInputKeyParams[],
  inputVideos: SwitchInputVideoParams[],
  busLevels: SwitchBusLevelParams[],
  levelIndex: number
  uniqueSignal?: boolean
  uniqueKeySignal?: boolean
  emitSignal?: boolean
  emitKeySignal?: boolean
}>(), {
  uniqueSignal: false,
  uniqueKeySignal: false,
  emitSignal: false,
  emitKeySignal: false
});
const { selectedSignal, signalSelected, signalUnSelected } = inject<InjectSwitchInputSignalSelect>('switch_input_signal_select', def_switch_input_signal_select())
const { selectedKeySignal, keySignalSelected, keySignalUnSelected } = inject<InjectSwitchKeySignalSelect>('switch_key_signal_select', def_switch_key_signal_select())

const options = computed(() => {
  if (props.type === 'key') {
    return props.inputKeys.map(item => ({
      signal_name: item.signal_name,
      signal_id: item.signal_id,
    }))
  } else if (props.type === 'video') {
    const isDisabled = (sid: string) => {
      let disabled = false
      if (props.uniqueSignal) {
        disabled ||= selectedSignal.value.has(sid)
      }
      if (props.uniqueKeySignal) {
        disabled ||= selectedKeySignal.value.has(sid)
      }
      return disabled && signalId.value !== sid
    }
    return props.inputVideos.map(item => ({
      signal_name: item.signal_name,
      signal_id: item.signal_id,
      disabled: isDisabled(item.signal_id)
    }))
  } else if (props.type === 'bus_output') {
    return props.busLevels.reduce((acc, item) => {
      item.pgm_bus.out_signal?.forEach(signal => {
        if (item.level !== props.levelIndex) {
          acc.push({
            signal_name: signal.signal_name,
            signal_id: signal.signal_id,
          })
        }
      })
      item.pvw_bus.out_signal?.forEach(signal => {
        if (item.level !== props.levelIndex) {
          acc.push({
            signal_name: signal.signal_name,
            signal_id: signal.signal_id,
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

watch(signalId, (nv, ov) => {
  if (nv === ov) return
  if (nv && nv.startsWith('video')) {
    if (props.emitSignal) {
      signalSelected(nv)
    }
    if (props.emitKeySignal) {
      keySignalSelected(nv)
    }
  }
  if (ov && ov.startsWith('video')) {
    if (props.emitSignal) {
      signalUnSelected(ov)
    }
    if (props.emitKeySignal) {
      keySignalUnSelected(ov)
    }
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
