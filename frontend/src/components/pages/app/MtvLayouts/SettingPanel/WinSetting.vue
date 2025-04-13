<script lang="ts" setup>
import { type CellComponentProp, CellComponents } from '../utils';

const props = defineProps<{
  modelValue: LayoutDataItem
  bound: LayoutDimension
  base: LayoutDimension
}>()

const emit = defineEmits<{
  (e: 'update:model-value', value: LayoutDataItem): void,
  (e: 'swap', key?: CellComponentProp | 'border'): void
}>()

const allChecked = computed({
  get: () => {
    const result = CellComponents.every(v => Boolean(props.modelValue[v.key]))
    if (props.modelValue.win.showBorder) {
      return result && Boolean(props.modelValue.win.showBorder)
    }
    return result
  },
  set: (v: boolean) => {
    if (v) {
      emit('swap')
    } else {
      const newVal: Partial<LayoutDataItem> = {}
      for (const comp of CellComponents) {
        newVal[comp.key] = null
      }
      newVal.win = { ...props.modelValue.win, showBorder: false }
      emit('update:model-value', { ...props.modelValue, ...newVal })
    }
  }
})

const winw = computed({
  get: () => Math.round(props.modelValue.win.w * props.base.w),
  set: (v: number) =>
    emit('update:model-value', {
      ...props.modelValue,
      win: { ...props.modelValue.win, w: v / props.base.w },
    }),
})

const winh = computed({
  get: () => Math.round(props.modelValue.win.h * props.base.h),
  set: (v: number) =>
    emit('update:model-value', {
      ...props.modelValue,
      win: { ...props.modelValue.win, h: v / props.base.h },
    }),
})

const winx = computed({
  get: () => Math.round(props.modelValue.win.x * props.base.w),
  set: (v: number) =>
    emit('update:model-value', {
      ...props.modelValue,
      win: { ...props.modelValue.win, x: v / props.base.w },
    }),
})

const winy = computed({
  get: () => Math.round(props.modelValue.win.y * props.base.h),
  set: (v: number) =>
    emit('update:model-value', {
      ...props.modelValue,
      win: { ...props.modelValue.win, y: v / props.base.h },
    }),
})

function toggleComponent(key: string, checked: boolean) {
  if (checked) {
    emit('swap', key as (CellComponentProp | 'border'))
  } else {
    if (key === 'border') {
      emit('update:model-value', { ...props.modelValue, win: { ...props.modelValue.win, showBorder: false } })
    } else {
      emit('update:model-value', { ...props.modelValue, [key]: null })
    }
  }
}


</script>
<template>
  <section class="layout-row">
    <div>尺寸</div>
    <div class="layout-row-inner">
      <div class="layout-cell half">
        <label for="winw">w</label>
        <input id="winw" v-model="winw" type="number" min="0" :max="base.w" />
      </div>
      <div class="layout-cell half">
        <label for="winh">h</label>
        <input id="winh" v-model="winh" type="number" min="0" :max="base.h" />
      </div>
    </div>
  </section>
  <section class="layout-row">
    <div>位置</div>
    <div class="layout-row-inner">
      <div class="layout-cell half">
        <label for="winx">x</label>
        <input id="winx" v-model="winx" type="number" min="0" :max="base.w" />
      </div>
      <div class="layout-cell half">
        <label for="winy">y</label>
        <input id="winy" v-model="winy" type="number" min="0" :max="base.h" />
      </div>
    </div>
  </section>
  <section class="layout-row is-row-check">
    <div class="layout-row-inner is-head">
      <div class="layout-cell">
        <div>{{ allChecked ? '删除' : '使用' }}全部组件</div>
        <VCheckbox v-model="allChecked" color="primary" circle />
      </div>
    </div>
    <div class="layout-row-inner">
      <div class="layout-cell">
        <div>边框</div>
        <VCheckbox
          :model-value="modelValue.win.showBorder"
          color="primary"
          circle
          @update:model-value="v => toggleComponent('border', v)"
        />
      </div>
    </div>
    <div v-for="comp in CellComponents" :key="comp.key" class="layout-row-inner">
      <div class="layout-cell">
        <div>{{ comp.name }}</div>
        <VCheckbox
          :model-value="Boolean(modelValue[comp.key])"
          color="primary"
          circle
          @update:model-value="v => toggleComponent(comp.key, v)"
        />
      </div>
    </div>
  </section>
</template>
<style lang="scss">
.layout-row.is-row-check {
  padding-top: 12px;

  .layout-row-inner {

    &.is-head {
      font-weight: 500;
      margin-bottom: 4px;
    }

    & > .layout-cell {
      width: 100%;
      justify-content: space-between;
    }
  }
}
</style>
