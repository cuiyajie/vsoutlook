<script lang="ts" setup>

const props = defineProps<{
  modelValue: LayoutSingleVol,
  base: LayoutDimension
}>()

const emit = defineEmits<{
  (e: 'update:model-value', value: LayoutSingleVol): void
}>()

const start = computed({
  get: () => props.modelValue.start,
  set: (v: number) => emit('update:model-value', { ...props.modelValue, start: v })
})

const end = computed({
  get: () => props.modelValue.end,
  set: (v: number) => emit('update:model-value', { ...props.modelValue, end: Math.max(start.value, v) })
})

const volx = computed({
  get: () => Math.round(props.modelValue.x * props.base.w),
  set: (v: number) => {
    emit('update:model-value', { ...props.modelValue, x: v / props.base.w })
  }
})

const voly = computed({
  get: () => Math.round(props.modelValue.y * props.base.h),
  set: (v: number) => {
    emit('update:model-value', { ...props.modelValue, y: v / props.base.h })
  }
})

const onew = computed({
  get: () => Math.round(props.modelValue.one_w * props.base.w),
  set: (v: number) => {
    emit('update:model-value', { ...props.modelValue, one_w: v / props.base.w })
  }
})

const gap = computed({
  get: () => Math.round(props.modelValue.g * props.base.w),
  set: (v: number) => {
    emit('update:model-value', { ...props.modelValue, g: v / props.base.w })
  }
})

const h = computed({
  get: () => Math.round(props.modelValue.h * props.base.h),
  set: (v: number) => {
    emit('update:model-value', { ...props.modelValue, h: v / props.base.h })
  }
})

</script>
<template>
  <div class="layout-section">
    <slot />
    <section class="layout-row">
      <div>声道序号</div>
      <div class="layout-row-inner">
        <div class="layout-cell half">
          <label for="start">开始</label>
          <input id="start" v-model="start" type="number" min="0" :max="end">
        </div>
        <div class="layout-cell half">
          <label for="end">结束</label>
          <input id="end" v-model="end" type="number" :min="start">
        </div>
      </div>
    </section>
    <section class="layout-row">
      <div>位置</div>
      <div class="layout-row-inner">
        <div class="layout-cell half">
          <label for="volx">X</label>
          <input id="volx" v-model="volx" type="number" min="0" :max="base.w">
        </div>
        <div class="layout-cell half">
          <label for="voly">Y</label>
          <input id="voly" v-model="voly" type="number" min="0" :max="base.h">
        </div>
      </div>
    </section>
    <section class="layout-row">
      <div class="layout-row-inner">
        <div class="layout-cell">
          <div>音柱宽度</div>
          <input id="onew" v-model="onew" type="number" min="0" :max="base.w">
        </div>
      </div>
    </section>
    <section class="layout-row">
      <div class="layout-row-inner">
        <div class="layout-cell">
          <div>音柱间隔</div>
          <input id="gap" v-model="gap" type="number" min="0" :max="base.w">
        </div>
      </div>
    </section>
    <section class="layout-row">
      <div class="layout-row-inner">
        <div class="layout-cell">
          <div>音柱高度</div>
          <input id="h" v-model="h" type="number" min="0" :max="base.h">
        </div>
      </div>
    </section>
  </div>
</template>
<style lang="scss">
.layout-panel {
}
</style>
