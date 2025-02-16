<script lang="ts" setup>
import { useColorPickerRef } from "./useColorPickerRef";

const props = defineProps<{
  modelValue: LayoutMeta,
  base: LayoutDimension
}>()

const emit = defineEmits<{
  (e: 'update:model-value', value: LayoutMeta): void,
}>()

const fz = computed({
  get: () => Math.round(props.modelValue.fontSize * props.base.w),
  set: (v: number) => emit('update:model-value', { ...props.modelValue, fontSize: v / props.base.w })
})

const color = computed({
  get: () => props.modelValue.color,
  set: (v: string) => emit('update:model-value', { ...props.modelValue, color: v })
})

const bgColor = computed({
  get: () => props.modelValue.bgColor,
  set: (v: string) => emit('update:model-value', { ...props.modelValue, bgColor: v })
})

const mx = computed({
  get: () => Math.round(props.modelValue.x * props.base.w),
  set: (v: number) => emit('update:model-value', { ...props.modelValue, x: v / props.base.w })
})

const my = computed({
  get: () => Math.round(props.modelValue.y * props.base.h),
  set: (v: number) => emit('update:model-value', { ...props.modelValue, y: v / props.base.h })
})

const mw = computed({
  get: () => Math.round(props.modelValue.w * props.base.w),
  set: (v: number) => emit('update:model-value', { ...props.modelValue, w: v / props.base.w })
})

const mh = computed({
  get: () => Math.round(props.modelValue.h * props.base.h),
  set: (v: number) => emit('update:model-value', { ...props.modelValue, h: v / props.base.h })
})

const pickerRef = useColorPickerRef()

</script>
<template>
  <div class="layout-form">
    <section class="layout-row">
      <div>背景位置</div>
      <div class="layout-row-inner">
        <div class="layout-cell half">
          <label for="mx">x</label>
          <input id="mx" v-model="mx" type="number" min="0" :max="base.w">
        </div>
        <div class="layout-cell half">
          <label for="my">y</label>
          <input id="my" v-model="my" type="number" min="0" :max="base.h">
        </div>
      </div>
    </section>
    <section class="layout-row">
      <div>背景宽高</div>
      <div class="layout-row-inner">
        <div class="layout-cell half">
          <label for="mw">w</label>
          <input id="mw" v-model="mw" type="number" min="0" :max="base.w">
        </div>
        <div class="layout-cell half">
          <label for="mh">h</label>
          <input id="mh" v-model="mh" type="number" min="0" :max="base.h">
        </div>
      </div>
    </section>
    <section class="layout-row">
      <div class="layout-row-inner">
        <div class="layout-cell">
          <div>背景颜色</div>
          <color-picker v-model:pureColor="bgColor" format="rgb" shape="square" theme="black" picker-type="chrome" :picker-container="pickerRef" />
        </div>
      </div>
    </section>
    <section class="layout-row">
      <div class="layout-row-inner">
        <div class="layout-cell">
          <div>文字颜色</div>
          <color-picker v-model:pureColor="color" format="rgb" shape="square" theme="black" picker-type="chrome" :picker-container="pickerRef" />
        </div>
      </div>
    </section>
    <section class="layout-row">
      <div class="layout-row-inner">
        <div class="layout-cell">
          <div>字号</div>
          <input v-model="fz" type="number" min="0" :max="Math.floor(props.base.w / 5)">
        </div>
      </div>
    </section>
  </div>
</template>
<style lang="scss">
</style>
