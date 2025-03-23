<script setup lang="ts">
import { useVFieldContext } from '@src/composable/useVFieldContext'

export interface VSelectProps {
  raw?: boolean
  multiple?: boolean
}

export interface VSelectEmits {
  (e: 'change', payload: Event): void
}

defineOptions({
  inheritAttrs: false,
})

const emit = defineEmits<VSelectEmits>()
const modelValue = defineModel<any>({
  default: '',
  local: true,
})
const props = defineProps<VSelectProps>()
const attrs = useAttrs()

const { field, id } = useVFieldContext({
  create: false,
  help: 'VSelect',
})

const internal = computed({
  get() {
    if (field?.value) {
      return field.value.value
    } else {
      return modelValue.value
    }
  },
  set(value: any) {
    if (field?.value) {
      field.value.setValue(value)
    }
    modelValue.value = value
  },
})

const classes = computed(() => {
  if (props.raw) return []

  return ['select', props.multiple && 'is-multiple']
})

function handleChange(payload: Event) {
  emit('change', payload)
  field.value?.handleChange(payload)
}
</script>

<template>
  <div :class="classes">
    <select
      :id="id"
      v-bind="attrs"
      v-model="internal"
      :name="id"
      :multiple="props.multiple"
      @change="handleChange"
      @blur="field?.handleBlur"
    >
      <slot v-bind="{ selected: internal, id }" />
    </select>
  </div>
</template>
