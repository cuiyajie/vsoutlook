<script lang="ts" setup>
import { defaultWinColor } from "@src/components/pages/app/MtvLayouts/utils";

const props = defineProps<{
  modelValue: LayoutDataItem,
  bound: LayoutDimension
}>()

const mv = computed(() => props.modelValue.win)

const ratiow = computed(() => props.bound.w / mv.value.w)

const bcolor = computed(() => mv.value.border.color || defaultWinColor)
const bl = computed(() => mv.value.border.left * ratiow.value)
const br = computed(() => mv.value.border.right * ratiow.value)
const bt = computed(() => mv.value.border.top * ratiow.value)
const bb = computed(() => mv.value.border.bottom * ratiow.value)

</script>
<template>
  <div
    v-if="mv.showBorder"
    class="layout-border"
    :style="{
      '--layout-border-color': bcolor,
      '--layout-border-l': `${bl}px`,
      '--layout-border-r': `${br}px`,
      '--layout-border-t': `${bt}px`,
      '--layout-border-b': `${bb}px`
    }"
  >
    <div class="l-border l-border-t" />
    <div class="l-border l-border-r" />
    <div class="l-border l-border-b" />
    <div class="l-border l-border-l" />
  </div>
</template>
<style lang="scss">
.layout-border {
  position: absolute;
  inset: 0;
  pointer-events: none;
  z-index: 0;
  --layout-border-color: #f00;
  --layout-border-l: 1px;
  --layout-border-r: 1px;
  --layout-border-t: 1px;
  --layout-border-b: 1px;

  .l-border {
    position: absolute;
    background-color: var(--layout-border-color);

    &.l-border-t {
      top: 0;
      left: 0;
      right: 0;
      height: var(--layout-border-t);
    }

    &.l-border-r {
      top: 0;
      right: 0;
      bottom: 0;
      width: var(--layout-border-r);
    }

    &.l-border-b {
      bottom: 0;
      left: 0;
      right: 0;
      height: var(--layout-border-b);
    }

    &.l-border-l {
      top: 0;
      left: 0;
      bottom: 0;
      width: var(--layout-border-l);
    }
  }
}
</style>
