<script lang="ts" setup>
import { type IApiParams } from './Consts_V1';

const mv = defineModel<IApiParams>({
  default: {} as IApiParams,
  local: true,
})

defineProps<{
  label: string
}>()
</script>
<template>
  <VField :label="label" class="field-check">
    <VControl raw subcontrol class="check-control">
      <VCheckbox v-model="mv.checked" color="primary" />
    </VControl>
    <VControl>
      <VInputNumber
        v-model="mv.service_port"
        centered
        :min="0"
        :step="1"
      />
    </VControl>
  </VField>
</template>

