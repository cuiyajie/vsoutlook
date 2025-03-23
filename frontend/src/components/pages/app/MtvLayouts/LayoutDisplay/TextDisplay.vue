<script lang="ts" setup>
import { getFontFamily, defaultTextColor, defaultTextBgColor } from '../utils';
const props = defineProps<{
  no: number
  data: LayoutDataItem | null
  bound: LayoutDimension
}>()

const font1 = computed(() => props.data?.text?.rect1.fontFamily || getFontFamily())
const fontSize1 = computed(() => props.data?.text?.rect1.fontSize || 0)
const color1 = computed(() => props.data?.text?.rect1.color || defaultTextColor)
const bgColor1 = computed(() => props.data?.text?.rect1.bgColor || defaultTextBgColor)

const font2 = computed(() => props.data?.text?.rect2.fontFamily || getFontFamily())
const fontSize2 = computed(() => props.data?.text?.rect2.fontSize || 0)
const color2 = computed(() => props.data?.text?.rect2.color || defaultTextColor)
const bgColor2 = computed(() => props.data?.text?.rect2.bgColor || defaultTextBgColor)
</script>
<template>
  <div
    v-if="data?.text"
    data-role="text"
    class="layout-text"
    style="max-width: 100%"
    :style="{ gap: `${data.text.gap * bound.w}px` }"
  >
    <slot />
    <div
      class="text-rect"
      :style="{
        fontFamily: font1,
        color: color1,
        backgroundColor: bgColor1,
        fontSize: `${fontSize1 * bound.w}px`,
      }"
    >
      文字窗口 1
    </div>
    <div
      class="text-rect"
      :style="{
        fontFamily: font2,
        color: color2,
        backgroundColor: bgColor2,
        fontSize: `${fontSize2 * bound.w}px`,
      }"
    >
      文字窗口 2
    </div>
  </div>
</template>
<style lang="scss"></style>
