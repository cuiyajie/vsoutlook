<script lang="ts" setup>
import { type TallySignalParams } from './Consts'

const mv = defineModel<TallySignalParams>({
  default: {} as TallySignalParams,
  local: true,
})

defineProps<{
  index: number
  isLast: boolean
}>()

</script>
<template>
  <div class="form-fieldset-nested-5 collapse-form">
    <div class="fieldset-heading collapse-header is-nested">
      <h5>第 {{ index + 1 }} 路信源配置&nbsp;&nbsp;索引: {{ index }}</h5>
    </div>
    <div class="form-fieldset-nested-4">
      <div class="columns is-multiline">
        <div class="column is-6">
          <VField>
            <VLabel>切换台输入序号</VLabel>
            <VControl>
              <VInputNumber v-model="mv.level_input_index" centered :min="0" :step="1" />
            </VControl>
          </VField>
        </div>
        <div class="column is-6">
          <VField>
            <VLabel>tally索引</VLabel>
            <VControl>
              <VInputNumber v-model="mv.tally_index" centered :min="0" :step="1" />
            </VControl>
          </VField>
        </div>
      </div>
    </div>
  </div>
</template>
