<script lang="ts" setup>
import { useColorPickerRef } from "./useColorPickerRef";


const props = defineProps<{
  modelValue: LayoutWin,
  base: LayoutDimension
}>()

const emit = defineEmits<{
  (e: 'update:model-value', value: LayoutWin): void,
}>()

const winw = computed(() => Math.round(props.modelValue.w * props.base.w))
const winh = computed(() => Math.round(props.modelValue.h * props.base.h))

const showBorder = computed({
  get: () => props.modelValue.showBorder,
  set: (v: boolean) => emit('update:model-value', { ...props.modelValue, showBorder: v })
})

// hide the coordinate of border
// const winx = computed(() => Math.round(props.modelValue.x * props.base.w))
// const winy = computed(() => Math.round(props.modelValue.y * props.base.h))

const bl = computed({
  get: () => Math.round(props.modelValue.border.left * props.base.w),
  set: (v: number) => emit('update:model-value', { ...props.modelValue, border: { ...props.modelValue.border, left: v / props.base.w } })
})

const br = computed({
  get: () => Math.round(props.modelValue.border.right * props.base.w),
  set: (v: number) => emit('update:model-value', { ...props.modelValue, border: { ...props.modelValue.border, right: v / props.base.w } })
})

const bt = computed({
  get: () => Math.round(props.modelValue.border.top * props.base.w),
  set: (v: number) => emit('update:model-value', { ...props.modelValue, border: { ...props.modelValue.border, top: v / props.base.w } })
})

const bb = computed({
  get: () => Math.round(props.modelValue.border.bottom * props.base.w),
  set: (v: number) => emit('update:model-value', { ...props.modelValue, border: { ...props.modelValue.border, bottom: v / props.base.w } })
})

const bcolor = computed({
  get: () => props.modelValue.border.color,
  set: (v: string) => emit('update:model-value', { ...props.modelValue, border: { ...props.modelValue.border, color: v } })
})

const pickerRef = useColorPickerRef()

</script>
<template>
  <div class="layout-form">
    <section class="layout-row">
      <div>窗口尺寸</div>
      <div class="layout-row-inner">
        <div class="layout-cell half">
          <label for="winw">w</label>
          <input id="winw" v-model="winw" type="number" readonly>
        </div>
        <div class="layout-cell half">
          <label for="winh">h</label>
          <input id="winh" v-model="winh" type="number" readonly>
        </div>
      </div>
    </section>
    <section class="layout-row">
      <div class="layout-row-inner">
        <div class="layout-cell">
          <div>边框显示</div>
          <VCheckbox v-model="showBorder" color="primary" circle />
        </div>
      </div>
    </section>
    <!-- <Transition name="fade-slow">
      <section v-if="showBorder" class="layout-row">
        <div>位置</div>
        <div class="layout-row-inner">
          <div class="layout-cell half">
            <label for="winx">x</label>
            <input id="winx" v-model="winx" type="number" readonly>
          </div>
          <div class="layout-cell half">
            <label for="winy">y</label>
            <input id="winy" v-model="winy" type="number" readonly>
          </div>
        </div>
      </section>
    </Transition> -->
    <Transition name="fade-slow">
      <section v-if="showBorder" class="layout-row">
        <div>边框宽度</div>
        <div class="layout-row-inner">
          <div class="layout-cell half">
            <label for="bl">左</label>
            <input id="bl" v-model="bl" type="number" min="0" :max="base.w" step="1">
          </div>
          <div class="layout-cell half">
            <label for="br">右</label>
            <input id="br" v-model="br" type="number" min="0" :max="base.w" step="1">
          </div>
        </div>
        <div class="layout-row-inner">
          <div class="layout-cell half">
            <label for="bt">上</label>
            <input id="bt" v-model="bt" type="number" min="0" :max="base.h" step="1">
          </div>
          <div class="layout-cell half">
            <label for="bb">下</label>
            <input id="bb" v-model="bb" type="number" min="0" :max="base.h" step="1">
          </div>
        </div>
      </section>
    </Transition>
    <Transition name="fade-slow">
      <section v-if="showBorder" class="layout-row">
        <div class="layout-row-inner">
          <div class="layout-cell">
            <div>边框颜色</div>
            <color-picker v-model:pureColor="bcolor" format="rgb" shape="square" theme="black" picker-type="chrome" disable-alpha :picker-container="pickerRef" />
          </div>
        </div>
      </section>
    </Transition>
  </div>
</template>
<style lang="scss">
</style>
