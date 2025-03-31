<script lang="ts" setup>
import { bus_sub_types, type SwitchBusLevelSubParams, type SwitchBusKeyParams, type SwitchBusMeParams, bus_sub_types_map } from './Consts'

const mv = defineModel<SwitchBusLevelSubParams>({
  default: {} as SwitchBusLevelSubParams,
  local: true,
})

defineProps<{
  index: number
  isLast: boolean
  type: 'pgm' | 'pvw'
  busKeys: SwitchBusKeyParams[]
  busMes: SwitchBusMeParams[]
}>()

const emit = defineEmits<{
  (e: 'remove'): void
}>()

const switchTypeCategory = inject('switch_type_category')
const busSubTypes = computed(() => {
  return bus_sub_types.filter(bst => bst.value === 'key_bus' || (bst.value === 'me_bus' && switchTypeCategory !== 'bcswt'))
})
</script>
<template>
  <div class="form-fieldset-nested-4 collapse-form">
    <div class="fieldset-heading collapse-header is-nested">
      <h5>包含的第 {{ index + 1 }} 个子总线设置&nbsp;&nbsp;索引: {{ index }}</h5>
      <div class="collapse-icons">
        <div
          class="collapse-icon is-close-hidden"
          role="button"
          tabindex="0"
          @click.prevent.stop="emit('remove')"
          @keydown.space.prevent.stop="emit('remove')"
        >
          <VIcon icon="feather:x" />
        </div>
      </div>
    </div>
    <div class="form-fieldset-nested-3">
      <div class="columns is-multiline">
        <div class="column is-4">
          <VField>
            <VLabel>总线类型</VLabel>
            <VControl>
              <VSelect v-model="mv.bus_type" class="is-rounded">
                <VOption
                  v-for="bpst in busSubTypes"
                  :key="bpst.value"
                  :value="bpst.value"
                >
                  {{ bpst.label }}
                </VOption>
              </VSelect>
            </VControl>
          </VField>
        </div>
        <div class="column is-4">
          <VField>
            <VLabel>总线名称</VLabel>
            <VControl>
              <SwitchBusSelect
                v-model="mv.bus_id"
                :type="mv.bus_type"
                :title="bus_sub_types_map.get(mv.bus_type)?.label || '总线名称'"
                :bus-keys="busKeys"
                :bus-mes="busMes"
              ></SwitchBusSelect>
            </VControl>
          </VField>
        </div>
        <div class="column is-4">
          <VField>
            <VLabel>总线ID</VLabel>
            <VControl>
              <VLabel class="is-static">{{ mv.bus_id || '未设置' }}</VLabel>
            </VControl>
          </VField>
        </div>
      </div>
    </div>
  </div>
</template>
