<script lang="ts" setup>
import { defaultMetaColor } from "@src/components/pages/app/MtvLayouts/utils";

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

const color = computed(() => mv.value.meta?.color || defaultMetaColor.Text)
const bgColor = computed(() => mv.value.meta?.bgColor || defaultMetaColor.Bg)
const fz = computed(() => (mv.value.meta?.fontSize || (24 / 1920)) * ratiow.value)

</script>
<template>
  <div
    v-if="mv.meta"
    data-role="meta"
    class="layout-meta"
    :style="{
      left: `${mv.meta.x * ratiow}px`,
      top: `${mv.meta.y * ratioh}px`,
      width: `${mv.meta.w * ratiow}px`,
      height: `${mv.meta.h * ratioh}px`,
      fontSize: `${fz}px`,
      color,
      backgroundColor: bgColor,
    }"
    role="button"
    tabindex="0"
    @click.prevent="emit('select')"
    @keyup.enter.prevent="emit('select')"
  >
    视频信息
  </div>
</template>
<style lang="scss">
</style>
