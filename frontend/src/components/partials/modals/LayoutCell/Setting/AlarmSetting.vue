<script lang="ts" setup>
import { AlarmLang } from "@src/utils/enums";
import { useColorPickerRef } from "./useColorPickerRef";

const props = defineProps<{
  modelValue: LayoutAlarm,
  base: LayoutDimension
}>()

const emit = defineEmits<{
  (e: 'update:model-value', value: LayoutAlarm): void,
}>()

const bx = computed({
  get: () => Math.round(props.modelValue.border.x * props.base.w),
  set: (v: number) => {
    const nx = v / props.base.w
    emit('update:model-value', {
      ...props.modelValue,
      border: {
        ...props.modelValue.border,
        x: nx,
        w: props.modelValue.border.w + (props.modelValue.border.x - nx) * 2
      }
    })
  }
})

const by = computed({
  get: () => Math.round(props.modelValue.border.y * props.base.h),
  set: (v: number) => {
    const ny = v / props.base.h
    emit('update:model-value', {
      ...props.modelValue,
      border: {
        ...props.modelValue.border,
        y: ny,
        h: props.modelValue.border.h + (props.modelValue.border.y - ny) * 2
      }
    })
  }
})

const bw = computed({
  get: () => Math.round(props.modelValue.border.w * props.base.w),
  set: (v: number) => {
    const nw = v / props.base.w
    emit('update:model-value', {
      ...props.modelValue,
      border: {
        ...props.modelValue.border,
        w: nw,
        x: props.modelValue.border.x + (props.modelValue.border.w - nw) / 2
      }
    })
  }
})

const bh = computed({
  get: () => Math.round(props.modelValue.border.h * props.base.h),
  set: (v: number) => {
    const nh = v / props.base.h
    emit('update:model-value', {
      ...props.modelValue,
      border: {
        ...props.modelValue.border,
        h: nh,
        y: props.modelValue.border.y + (props.modelValue.border.h - nh) / 2
      }
    })
  }
})

const borderW = computed({
  get: () => Math.round(props.modelValue.border.width * props.base.w),
  set: (v: number) => emit('update:model-value', {
    ...props.modelValue,
    border: { ...props.modelValue.border, width: v / props.base.w }
  })
})

const borderColor = computed({
  get: () => props.modelValue.border.color,
  set: (v: string) => emit('update:model-value', {
    ...props.modelValue,
    border: { ...props.modelValue.border, color: v }
  })
})

const interval = computed({
  get: () => Math.round(props.modelValue.border.interval),
  set: (v: number) => emit('update:model-value', {
    ...props.modelValue,
    border: { ...props.modelValue.border, interval: v }
  })
})

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

const ax = computed({
  get: () => Math.round(props.modelValue.x * props.base.w),
  set: (v: number) => emit('update:model-value', { ...props.modelValue, x: v / props.base.w })
})

const ay = computed({
  get: () => Math.round(props.modelValue.y * props.base.h),
  set: (v: number) => emit('update:model-value', { ...props.modelValue, y: v / props.base.h })
})

const aw = computed({
  get: () => Math.round(props.modelValue.w * props.base.w),
  set: (v: number) => emit('update:model-value', { ...props.modelValue, w: v / props.base.w })
})

const ah = computed({
  get: () => Math.round(props.modelValue.h * props.base.h),
  set: (v: number) => emit('update:model-value', { ...props.modelValue, h: v / props.base.h })
})

const lang = computed({
  get: () => props.modelValue.lang,
  set: (v: AlarmLang) => emit('update:model-value', { ...props.modelValue, lang: v })
})

const pickerRef = useColorPickerRef()

</script>
<template>
  <div class="layout-form">
    <div class="layout-section">
      <div class="section-header">报警边框</div>
      <section class="layout-row">
        <div>边框位置</div>
        <div class="layout-row-inner">
          <div class="layout-cell half">
            <label for="bx">x</label>
            <input id="bx" v-model="bx" type="number" min="0" :max="base.w" :step="1">
          </div>
          <div class="layout-cell half">
            <label for="by">y</label>
            <input id="by" v-model="by" type="number" min="0" :max="base.h" :step="1">
          </div>
        </div>
      </section>
      <section class="layout-row">
        <div>边框宽高</div>
        <div class="layout-row-inner">
          <div class="layout-cell half">
            <label for="bw">w</label>
            <input id="bw" v-model="bw" type="number" min="0" :max="base.w" :step="1">
          </div>
          <div class="layout-cell half">
            <label for="bh">h</label>
            <input id="bh" v-model="bh" type="number" min="0" :max="base.h" :step="1">
          </div>
        </div>
      </section>
      <section class="layout-row">
        <div class="layout-row-inner">
          <div class="layout-cell">
            <div>边框宽度</div>
            <input id="borderW" v-model="borderW" type="number" min="0" :max="base.h">
          </div>
        </div>
      </section>
      <section class="layout-row">
        <div class="layout-row-inner">
          <div class="layout-cell">
            <div>边框颜色</div>
            <color-picker v-model:pureColor="borderColor" format="rgb" shape="square" theme="black" picker-type="chrome" :picker-container="pickerRef" />
          </div>
        </div>
      </section>
      <section class="layout-row">
        <div class="layout-row-inner">
          <div class="layout-cell">
            <div>闪烁间隔</div>
            <input id="borderW" v-model="interval" type="number" min="0" :step="100">
          </div>
        </div>
      </section>
    </div>
    <div class="layout-section">
      <div class="section-header">报警文字</div>
      <section class="layout-row">
        <div>语言</div>
        <div class="layout-row-inner">
          <div class="layout-cell">
            <VRadio v-model="lang" :value="AlarmLang.Chinese" label="中文" name="lang" color="primary" />
          </div>
          <div class="layout-cell">
            <VRadio v-model="lang" :value="AlarmLang.English" label="英文" name="lang" color="primary" />
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
      <section class="layout-row">
        <div class="layout-row-inner">
          <div class="layout-cell">
            <div>字体颜色</div>
            <color-picker v-model:pureColor="color" format="rgb" shape="square" theme="black" picker-type="chrome" :picker-container="pickerRef" />
          </div>
        </div>
      </section>
      <section class="layout-row">
        <div>背景位置</div>
        <div class="layout-row-inner">
          <div class="layout-cell half">
            <label for="ax">x</label>
            <input id="ax" v-model="ax" type="number" :min="bx" :max="bx + bw - aw">
          </div>
          <div class="layout-cell half">
            <label for="ay">y</label>
            <input id="ay" v-model="ay" type="number" :min="by" :max="by + bh - ah">
          </div>
        </div>
      </section>
      <section class="layout-row">
        <div>背景宽高</div>
        <div class="layout-row-inner">
          <div class="layout-cell half">
            <label for="aw">w</label>
            <input id="aw" v-model="aw" type="number" min="0" :max="bw">
          </div>
          <div class="layout-cell half">
            <label for="ah">h</label>
            <input id="ah" v-model="ah" type="number" min="0" :max="bh">
          </div>
        </div>
      </section>
      <section class="layout-row">
        <div class="layout-row-inner">
          <div class="layout-cell">
            <div>背景框颜色</div>
            <color-picker v-model:pureColor="bgColor" format="rgb" shape="square" theme="black" picker-type="chrome" :picker-container="pickerRef" />
          </div>
        </div>
      </section>
    </div>
  </div>
</template>
<style lang="scss">
.layout-panel {
}
</style>
