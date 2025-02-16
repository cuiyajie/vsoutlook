<script lang="ts" setup>

const props = defineProps<{
  modelValue: LayoutArea,
  base: LayoutDimension
}>()

const emit = defineEmits<{
  (e: 'update:model-value', value: LayoutArea): void,
}>()

const vx = computed({
  get: () => Math.round(props.modelValue.x * props.base.w),
  set: (v: number) => emit('update:model-value', { ...props.modelValue, x: v / props.base.w })
})

const vy = computed({
  get: () => Math.round(props.modelValue.y * props.base.h),
  set: (v: number) => emit('update:model-value', { ...props.modelValue, y: v / props.base.h })
})

const vw = computed({
  get: () => Math.round(props.modelValue.w * props.base.w),
  set: (v: number) => emit('update:model-value', {
    ...props.modelValue,
    w: v / props.base.w,
    h: v / props.base.h,
  })
})

const vh = computed({
  get: () => Math.round(props.modelValue.h * props.base.h),
  set: (v: number) => emit('update:model-value', {
    ...props.modelValue,
    h: v / props.base.h,
    w: v / props.base.w,
  })
})

const alpha = computed({
  get: () => Math.round(props.modelValue.alpha * 100),
  set: (v: number) => emit('update:model-value', { ...props.modelValue, alpha: v / 100 })
})

</script>
<template>
  <div class="layout-form">
    <section class="layout-row">
      <div>矢量图位置</div>
      <div class="layout-row-inner">
        <div class="layout-cell half">
          <label for="vx">x</label>
          <input id="vx" v-model="vx" type="number" min="0" :max="base.w">
        </div>
        <div class="layout-cell half">
          <label for="vy">y</label>
          <input id="vy" v-model="vy" type="number" min="0" :max="base.h">
        </div>
      </div>
    </section>
    <section class="layout-row">
      <div>矢量图宽高</div>
      <div class="layout-row-inner">
        <div class="layout-cell half">
          <label for="vw">w</label>
          <input id="vw" v-model="vw" type="number" min="0" :max="base.w">
        </div>
        <div class="layout-cell half">
          <label for="vh">h</label>
          <input id="vh" v-model="vh" type="number" min="0" :max="base.h">
        </div>
      </div>
    </section>
    <section class="layout-row">
      <div class="layout-row-inner">
        <div class="layout-cell">
          <div>矢量透明度</div>
          <input v-model="alpha" type="number" min="0" :max="100" step="1" style="width: 28px;">
          <div>%</div>
        </div>
      </div>
    </section>
  </div>
</template>
<style lang="scss">
</style>
