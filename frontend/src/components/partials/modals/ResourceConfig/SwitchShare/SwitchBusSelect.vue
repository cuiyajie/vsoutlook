<script lang="ts" setup>
import { type SwitchBusMeParams, type SwitchBusKeyParams } from "./Consts";

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
const options = computed(() => {
  if (props.type === 'key_bus') {
    return props.busKeys.map(item => ({
      bus_name: item.bus_name,
      bus_id: item.bus_id
    }))
  } else if (props.type === 'me_bus') {
    return props.busMes.map(item => ({
      bus_name: item.bus_name,
      bus_id: item.bus_id
    }))
  }
  return []
})

watch(() => props.type, () => {
  busId.value = ''
})

</script>
<template>
  <Multiselect
    v-model="busId"
    class="tippy-select"
    :placeholder="`选择${title}`"
    value-prop="bus_id"
    label="bus_name"
    :can-deselect="false"
    :max-height="145"
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
