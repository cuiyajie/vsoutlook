<script lang="ts" setup>
import { defaultTitleColor } from "@src/components/pages/app/MtvLayouts/utils";
import { TallyType } from "@src/utils/enums";

const props = defineProps<{
  modelValue: LayoutDataItem,
  bound: LayoutDimension
  className?: any
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

const color = computed(() => mv.value.title?.color || defaultTitleColor.Text)
const bgColor = computed(() => mv.value.title?.bgColor || defaultTitleColor.Bg)
const tallyBgColor = computed(() => mv.value.title?.tallyBgColor || defaultTitleColor.Tally)
const tallyBorderColor = computed(() => mv.value.title?.tallyBorderColor || defaultTitleColor.TallyBorder)
const tb = computed(() => (mv.value.title?.tallyBorderWidth || 8 / 1920) * ratiow.value)
const tw = computed(() => (mv.value.title?.tallyWidth || 288 / 1920) * ratiow.value)
const bl = computed(() => mv.value.win.border.left * ratiow.value)
const br = computed(() => mv.value.win.border.right * ratiow.value)
const bt = computed(() => mv.value.win.border.top * ratiow.value)
const bb = computed(() => mv.value.win.border.bottom * ratiow.value)

</script>
<template>
  <div
    v-if="mv.title"
    data-role="title"
    :class="['layout-title', className]"
    :style="{
      left: `${mv.title.x * ratiow}px`,
      top: `${mv.title.y * ratioh}px`,
      width: `${mv.title.w * ratiow}px`,
      height: `${mv.title.h * ratioh}px`,
      '--tally-light-width': `${tw}px`,
      '--tally-light-bg': tallyBgColor,
      '--tally-text-fz': `${mv.title.fontSize * ratiow}px`,
      '--tally-text-fa': `'${mv.title.fontFamily}'`,
      '--tally-text-bg': bgColor,
      '--tally-text-color': color,
      '--tally-rotate': mv.title.isVertial ? '90deg' : '0deg',
    }"
    role="button"
    tabindex="0"
    @click.prevent="emit('select')"
    @keyup.enter.prevent="emit('select')"
  >
    <div v-if="mv.title.tallyType === TallyType.Light" class="tally-light" />
    <div class="tally-text">窗口名称</div>
    <div v-if="mv.title.tallyType === TallyType.Light" class="tally-light" />
  </div>
  <div
    v-if="mv.title?.tallyType === TallyType.Border"
    :class="['layout-tally-border', className]"
    :style="{
      '--tally-border-color': tallyBorderColor,
      '--tally-border': `${tb}px`,
      '--layout-border-l': `${bl}px`,
      '--layout-border-r': `${br}px`,
      '--layout-border-t': `${bt}px`,
      '--layout-border-b': `${bb}px`
    }"
  >
    <div class="ta-border ta-border-t" />
    <div class="ta-border ta-border-r" />
    <div class="ta-border ta-border-b" />
    <div class="ta-border ta-border-l" />
  </div>
</template>
<style lang="scss">
.layout-title {
  gap: 2%;
  --tally-light-width: 10px;
  --tally-light-bg: rgba(255, 0, 0, 1.0),
  --tally-text-fz: 12px;
  --tally-text-fa: 'Arial';
  --tally-text-bg: rgba(0, 0, 0, 0.5);
  --tally-text-color: rgba(255, 255, 255, 1.0);
  transform: rotate(var(--tally-rotate));
  transform-origin: top left;

  .tally-light {
    width: var(--tally-light-width);
    background-color: var(--tally-light-bg);
    height: 100%;
    flex: 0 0 auto;
  }

  .tally-text {
    flex: 1;
    min-width: 0;
    white-space: nowrap;
    text-wrap: nowrap;
    overflow: hidden;
    display: flex;
    align-items: center;
    justify-content: center;
    font-size: var(--tally-text-fz);
    font-family: var(--tally-text-fa);
    background-color: var(--tally-text-bg);
    color: var(--tally-text-color);
  }
}
.layout-tally-border {
  position: absolute;
  inset: 0;
  pointer-events: none;
  z-index: 0;
  --tally-border-color: rgba(240, 240, 240, 1.0);
  --tally-border: 1px;
  --layout-border-l: 1px;
  --layout-border-r: 1px;
  --layout-border-t: 1px;
  --layout-border-b: 1px;

  .ta-border {
    position: absolute;
    background-color: var(--tally-border-color);

    &.ta-border-t {
      top: var(--layout-border-t);
      left: var(--layout-border-l);
      right: var(--layout-border-r);
      height: var(--tally-border);
    }

    &.ta-border-r {
      top: var(--layout-border-t);
      bottom: var(--layout-border-b);
      right: var(--layout-border-r);
      width: var(--tally-border);
    }

    &.ta-border-b {
      bottom: var(--layout-border-b);
      left: var(--layout-border-l);
      right: var(--layout-border-r);
      height: var(--tally-border);
    }

    &.ta-border-l {
      top: var(--layout-border-t);
      bottom: var(--layout-border-b);
      left: var(--layout-border-l);
      width: var(--tally-border);
    }
  }
}
</style>
