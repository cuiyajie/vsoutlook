<script lang="ts" setup>
const props = defineProps<{
  modelValue: LayoutDataItem,
  bound: LayoutDimension
}>()

const emit = defineEmits<{
  (e: 'update:model-value', value: LayoutDataItem): void,
  (e: 'select'): void
}>()

const mv = computed({
  get: () => props.modelValue,
  set: (v: LayoutDataItem) => emit('update:model-value', v)
})

const ratiow = computed(() => props.bound.w / mv.value.win.w)
const ratioh = computed(() => props.bound.h / mv.value.win.h)

</script>
<template>
  <div
    v-if="mv.oscillogram"
    class="layout-oscillogram"
    data-role="oscillogram"
    :style="{
      left: `${mv.oscillogram.x * ratiow}px`,
      top: `${mv.oscillogram.y * ratioh}px`,
      width: `${mv.oscillogram.w * ratiow}px`,
      height: `${mv.oscillogram.h * ratioh}px`,
      opacity: mv.oscillogram.alpha,
    }"
    role="button"
    tabindex="0"
    @click.prevent="emit('select')"
    @keyup.enter.prevent="emit('select')"
  />
</template>
<style lang="scss">
.layout-oscillogram {
  background: url('@src/assets/oscillogram.svg') no-repeat center center/cover;
}
</style>
