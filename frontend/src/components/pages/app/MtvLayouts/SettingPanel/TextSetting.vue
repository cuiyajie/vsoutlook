<script lang="ts" setup>
import { clamp } from "@vueuse/core";
import { defaultTextBgColor, defaultTextColor } from "../utils";
import roundToNearest4 from '@src/utils/round-to-nearest-four';
import roundToNearestEven from "@src/utils/round-to-nearest-even";

const props = defineProps<{
  modelValue: LayoutDataItem,
  bound: LayoutDimension,
  base: LayoutDimension
}>()

const emit = defineEmits<{
  (e: 'update:model-value', value: LayoutDataItem): void
}>()

const txtw = computed({
  get: () => props.modelValue?.text ? roundToNearest4(props.modelValue.text.w * props.base.w) : 0,
  set: (v: number) => props.modelValue?.text && emit('update:model-value', { ...props.modelValue, text: { ...props.modelValue.text, w: v / props.base.w } })
})

const txth = computed({
  get: () => props.modelValue?.text ? roundToNearestEven(props.modelValue.text.h * props.base.h) : 0,
  set: (v: number) => props.modelValue?.text && emit('update:model-value', { ...props.modelValue, text: { ...props.modelValue.text, h: v / props.base.h } })
})

const txtx = computed({
  get: () => props.modelValue?.text ? Math.round(props.modelValue.text.x * props.base.w) : 0,
  set: (v: number) => props.modelValue?.text && emit('update:model-value', { ...props.modelValue, text: { ...props.modelValue.text, x: v / props.base.w } })
})

const txty = computed({
  get: () => props.modelValue?.text ? Math.round(props.modelValue.text.y * props.base.h) : 0,
  set: (v: number) => props.modelValue?.text && emit('update:model-value', { ...props.modelValue, text: { ...props.modelValue.text, y: v / props.base.h } })
})

const gap = computed({
  get: () => props.modelValue?.text ? roundToNearestEven(props.modelValue.text.gap * props.base.w) : 0,
  set: (v: number) => props.modelValue?.text && emit('update:model-value', { ...props.modelValue, text: { ...props.modelValue.text, gap: v / props.base.w } })
})

const fz1 = computed({
  get: () => props.modelValue?.text?.rect1 ? Math.round(props.modelValue.text.rect1.fontSize * props.base.w) : 0,
  set: (v: number) => {
    if (props.modelValue?.text?.rect1) {
      const realV = clamp(v, 0, Math.floor(props.base.w / 5))
      emit('update:model-value', { ...props.modelValue, text: {
        ...props.modelValue.text,
        rect1: {
          ...props.modelValue.text.rect1,
          fontSize: realV / props.base.w
        }
      } })
    }
  }
})

const fa1 = computed({
  get: () => props.modelValue?.text?.rect1?.fontFamily || '',
  set: (v: string) => {
    if (props.modelValue?.text?.rect1) {
      emit('update:model-value', { ...props.modelValue, text: {
        ...props.modelValue.text,
        rect1: {
          ...props.modelValue.text.rect1,
          fontFamily: v
        }
      } })
    }
  }
})

const color1 = computed({
  get: () => props.modelValue?.text?.rect1?.color || defaultTextColor,
  set: (v: string) => {
    if (props.modelValue?.text?.rect1) {
      emit('update:model-value', { ...props.modelValue, text: {
        ...props.modelValue.text,
        rect1: {
          ...props.modelValue.text.rect1,
          color: v
        }
      } })
    }
  }
})

const bgColor1 = computed({
  get: () => props.modelValue?.text?.rect1?.bgColor || defaultTextBgColor,
  set: (v: string) => {
    if (props.modelValue?.text?.rect1) {
      emit('update:model-value', { ...props.modelValue, text: {
        ...props.modelValue.text,
        rect1: {
          ...props.modelValue.text.rect1,
          bgColor: v
        }
      } })
    }
  }
})

const fz2 = computed({
  get: () => props.modelValue?.text?.rect2 ? Math.round(props.modelValue.text.rect2.fontSize * props.base.w) : 0,
  set: (v: number) => {
    if (props.modelValue?.text?.rect2) {
      const realV = clamp(v, 0, Math.floor(props.base.w / 5))
      emit('update:model-value', { ...props.modelValue, text: {
        ...props.modelValue.text,
        rect2: {
          ...props.modelValue.text.rect2,
          fontSize: realV / props.base.w
        }
      } })
    }
  }
})

const fa2 = computed({
  get: () => props.modelValue?.text?.rect2?.fontFamily || '',
  set: (v: string) => {
    if (props.modelValue?.text?.rect2) {
      emit('update:model-value', { ...props.modelValue, text: {
        ...props.modelValue.text,
        rect2: {
          ...props.modelValue.text.rect2,
          fontFamily: v
        }
      } })
    }
  }
})

const color2 = computed({
  get: () => props.modelValue?.text?.rect2?.color || defaultTextColor,
  set: (v: string) => {
    if (props.modelValue?.text?.rect2) {
      emit('update:model-value', { ...props.modelValue, text: {
        ...props.modelValue.text,
        rect2: {
          ...props.modelValue.text.rect2,
          color: v
        }
      } })
    }
  }
})

const bgColor2 = computed({
  get: () => props.modelValue?.text?.rect2?.bgColor || defaultTextBgColor,
  set: (v: string) => {
    if (props.modelValue?.text?.rect2) {
      emit('update:model-value', { ...props.modelValue, text: {
        ...props.modelValue.text,
        rect2: {
          ...props.modelValue.text.rect2,
          bgColor: v
        }
      } })
    }
  }
})


</script>
<template>
  <section class="layout-row">
    <div>尺寸</div>
    <div class="layout-row-inner">
      <div class="layout-cell half">
        <label for="txtw">w</label>
        <input id="txtw" v-model="txtw" type="number" min="0" :max="base.w" step="4">
      </div>
      <div class="layout-cell half">
        <label for="txth">h</label>
        <input id="txth" v-model="txth" type="number" min="0" :max="base.h" step="2">
      </div>
    </div>
  </section>
  <section class="layout-row">
    <div>位置</div>
    <div class="layout-row-inner">
      <div class="layout-cell half">
        <label for="txtx">x</label>
        <input id="txtx" v-model="txtx" type="number" min="0" :max="base.w">
      </div>
      <div class="layout-cell half">
        <label for="txty">y</label>
        <input id="txty" v-model="txty" type="number" min="0" :max="base.h">
      </div>
    </div>
  </section>
  <section class="layout-row">
    <div class="layout-row-inner">
      <div class="layout-cell">
        <div>间距</div>
        <input v-model="gap" type="number" min="0" :max="txtw / 2" step="2">
      </div>
    </div>
  </section>
  <div class="layout-section">
    <div class="section-header">窗口 1</div>
    <section class="layout-row">
      <div class="layout-row-inner">
        <div class="layout-cell">
          <div>字号</div>
          <input v-model="fz1" type="number" min="0" :max="Math.floor(props.base.w / 5)">
        </div>
      </div>
    </section>
    <section class="layout-row">
      <div class="layout-row-inner">
        <div class="layout-cell fullwidth">
          <div>字体</div>
          <input v-model="fa1" type="text">
        </div>
      </div>
    </section>
    <section class="layout-row">
      <div class="layout-row-inner">
        <div class="layout-cell">
          <div>字体颜色</div>
          <color-picker v-model:pureColor="color1" format="rgb" shape="square" theme="black" picker-type="chrome" />
        </div>
      </div>
    </section>
    <section class="layout-row">
      <div class="layout-row-inner">
        <div class="layout-cell">
          <div>背景颜色</div>
          <color-picker v-model:pureColor="bgColor1" format="rgb" shape="square" theme="black" picker-type="chrome" />
        </div>
      </div>
    </section>
  </div>
  <div class="layout-section">
    <div class="section-header">窗口 2</div>
    <section class="layout-row">
      <div class="layout-row-inner">
        <div class="layout-cell">
          <div>字号</div>
          <input v-model="fz2" type="number" min="0" :max="Math.floor(props.base.w / 5)">
        </div>
      </div>
    </section>
    <section class="layout-row">
      <div class="layout-row-inner">
        <div class="layout-cell fullwidth">
          <div>字体</div>
          <input v-model="fa2" type="text">
        </div>
      </div>
    </section>
    <section class="layout-row">
      <div class="layout-row-inner">
        <div class="layout-cell">
          <div>字体颜色</div>
          <color-picker v-model:pureColor="color2" format="rgb" shape="square" theme="black" picker-type="chrome" />
        </div>
      </div>
    </section>
    <section class="layout-row">
      <div class="layout-row-inner">
        <div class="layout-cell">
          <div>背景颜色</div>
          <color-picker v-model:pureColor="bgColor2" format="rgb" shape="square" theme="black" picker-type="chrome" />
        </div>
      </div>
    </section>
  </div>
</template>
<style lang="scss"></style>
