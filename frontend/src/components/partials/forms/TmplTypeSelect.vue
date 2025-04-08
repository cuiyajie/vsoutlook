<script setup lang="ts">
import { useTemplateType } from '@src/stores/templateType'

const tmplTypeStore = useTemplateType()

const modalValue = defineModel<any>('modelValue', {
  default: '',
  local: true,
})

defineProps<{
  validate?: boolean
  formId?: string
  label?: string
}>()
</script>
<template>
  <VField
    :id="formId"
    v-slot="{ field }"
    :label="label"
    class="is-image-select is-tt-select"
  >
    <VControl>
      <Multiselect
        v-model="modalValue"
        placeholder="选择应用类型"
        value-prop="id"
        label="name"
        :style="{'--ms-max-height': '245px'}"
        :options="tmplTypeStore.tmplTypes"
        @change="(val: any) => field?.setValue(val)"
      >
        <template #singlelabel="{ value }">
          <div class="multiselect-single-label">
            <img class="select-label-icon" :src="value.icon" alt="" />
            <span class="select-label-text">
              {{ value.name }}
            </span>
          </div>
        </template>
        <template #option="{ option }">
          <img class="select-option-icon" :src="option.icon" alt="" />
          <span class="select-label-text">
            {{ option.name }}
          </span>
        </template>
      </Multiselect>
      <p v-if="field?.errorMessage" class="help is-danger">
        {{ field.errorMessage }}
      </p>
    </VControl>
  </VField>
</template>
<style lang="scss">
.is-image-select.is-tt-select {
  .multiselect {
    .select-option-icon {
      margin-inline-end: 10px;
      height: 18px;
      min-width: 18px;
    }

    .select-label-icon {
      margin-inline-start: 2px;
      margin-inline-end: 6px;
      height: 18px;
      min-width: 18px;
    }
  }
}
</style>
