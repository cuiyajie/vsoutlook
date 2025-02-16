<script lang="ts" setup>
import { draftVol, draftTitle } from '../utils';

const props = defineProps<{
  modelValue: LayoutDataItem,
  bound: LayoutDimension,
  base: LayoutDimension
}>()

const emit = defineEmits<{
  (e: 'update:model-value', value: LayoutDataItem): void
}>()

const winw = computed({
  get: () => Math.round(props.modelValue.win.w * props.base.w),
  set: (v: number) => emit('update:model-value', { ...props.modelValue, win: { ...props.modelValue.win, w: v / props.base.w } })
})

const winh = computed({
  get: () => Math.round(props.modelValue.win.h * props.base.h),
  set: (v: number) => emit('update:model-value', { ...props.modelValue, win: { ...props.modelValue.win, h: v / props.base.h } })
})

const winx = computed({
  get: () => Math.round(props.modelValue.win.x * props.base.w),
  set: (v: number) => emit('update:model-value', { ...props.modelValue, win: { ...props.modelValue.win, x: v / props.base.w } })
})

const winy = computed({
  get: () => Math.round(props.modelValue.win.y * props.base.h),
  set: (v: number) => emit('update:model-value', { ...props.modelValue, win: { ...props.modelValue.win, y: v / props.base.h } })
})

const reservedTitle = ref<LayoutDataItem['title']>(null)
const title = computed({
  get: () => Boolean(props.modelValue.title),
  set: (v: boolean) => {
    if (v) {
      reservedTitle.value = reservedTitle.value || draftTitle(props.modelValue.win.w * props.bound.w, props.modelValue.win.h * props.bound.h, props.bound)
    } else {
      reservedTitle.value = props.modelValue.title
    }
    emit('update:model-value', { ...props.modelValue, title: v ? reservedTitle.value : null})
  }
})

const reservedVol = ref<LayoutDataItem['vol']>(null)
const vol = computed({
  get: () => Boolean(props.modelValue.vol),
  set: (v: boolean) => {
    if (v) {
      reservedVol.value = reservedVol.value || draftVol(props.modelValue.win.h * props.bound.h, props.bound)
    } else {
      reservedVol.value = props.modelValue.vol
    }
    emit('update:model-value', { ...props.modelValue, vol: v ? reservedVol.value : null })
  }
})
</script>
<template>
  <section class="layout-row">
    <div>尺寸</div>
    <div class="layout-row-inner">
      <div class="layout-cell half">
        <label for="winw">w</label>
        <input id="winw" v-model="winw" type="number" min="0" :max="base.w">
      </div>
      <div class="layout-cell half">
        <label for="winh">h</label>
        <input id="winh" v-model="winh" type="number" min="0" :max="base.h">
      </div>
    </div>
  </section>
  <section class="layout-row">
    <div>位置</div>
    <div class="layout-row-inner">
      <div class="layout-cell half">
        <label for="winx">x</label>
        <input id="winx" v-model="winx" type="number" min="0" :max="base.w">
      </div>
      <div class="layout-cell half">
        <label for="winy">y</label>
        <input id="winy" v-model="winy" type="number" min="0" :max="base.h">
      </div>
    </div>
  </section>
  <section class="layout-row">
    <div class="layout-row-inner">
      <div class="layout-cell">
        <div>显示名称</div>
        <VCheckbox v-model="title" color="primary" />
      </div>
    </div>
  </section>
  <section class="layout-row">
    <div class="layout-row-inner">
      <div class="layout-cell">
        <div>显示音柱</div>
        <VCheckbox v-model="vol" color="primary" />
      </div>
    </div>
  </section>
</template>
<style lang="scss"></style>
