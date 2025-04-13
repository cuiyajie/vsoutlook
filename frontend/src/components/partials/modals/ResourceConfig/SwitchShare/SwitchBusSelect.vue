<script lang="ts" setup>
import { type SwitchBusMeParams, type SwitchBusKeyParams, type InjectSwitchBusSelect, def_switch_bus_select } from "./Consts";

const busId = defineModel<string>({
  default: "",
  local: true,
});

const props = defineProps<{
  title: string,
  type: string,
  busKeys: SwitchBusKeyParams[],
  busMes: SwitchBusMeParams[],
}>();

const { selectedBus, busSelected, busUnSelected } = inject<InjectSwitchBusSelect>('switch_bus_select', def_switch_bus_select())

const options = computed(() => {
  const isDisabled = (bid: string) => selectedBus.value.has(bid) && bid !== busId.value
  if (props.type === 'key_bus') {
    return props.busKeys.map(item => ({
      bus_name: item.bus_name,
      bus_id: item.bus_id,
      disabled: isDisabled(item.bus_id)
    }))
  } else if (props.type === 'me_bus') {
    return props.busMes.map(item => ({
      bus_name: item.bus_name,
      bus_id: item.bus_id,
      disabled: isDisabled(item.bus_id)
    }))
  }
  return []
})

watch(() => props.type, () => {
  busId.value = ''
})

watch(busId, (nv, ov) => {
  if (nv === ov) return
  if (nv) {
    busSelected(nv)
  }
  if (ov) {
    busUnSelected(ov)
  }
}, { immediate: true })

</script>
<template>
  <Multiselect
    v-model="busId"
    class="tippy-select"
    :placeholder="`选择${title}`"
    value-prop="bus_id"
    label="bus_name"
    :can-deselect="false"
    :style="{'--ms-max-height': '245px'}"
    :no-options-text="`暂时没有配置${title}`"
    :options="options"
  >
    <template #option="{ option }">
      <span class="select-option-text">
        {{ option.bus_name }}
      </span>
    </template>
  </Multiselect>
</template>
