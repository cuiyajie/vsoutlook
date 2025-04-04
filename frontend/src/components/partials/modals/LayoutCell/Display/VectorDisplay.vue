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
    v-if="mv.vector"
    class="layout-vector"
    data-role="vector"
    :style="{
      left: `${mv.vector.x * ratiow}px`,
      top: `${mv.vector.y * ratioh}px`,
      width: `${mv.vector.w * ratiow}px`,
      height: `${mv.vector.h * ratioh}px`,
      opacity: mv.vector.alpha,
    }"
    role="button"
    tabindex="0"
    @click.prevent="emit('select')"
    @keyup.enter.prevent="emit('select')"
  />
</template>
<style lang="scss">
.layout-vector {
  background: url('@src/assets/color-vector.png') no-repeat center center/cover;
}
</style>
