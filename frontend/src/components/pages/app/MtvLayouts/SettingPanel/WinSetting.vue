<script lang="ts" setup>
import { type Ref } from 'vue';
import { type CellComponentProp, CellComponents, getDefaultCheckStore } from '../utils';

const props = defineProps<{
  modelValue: LayoutDataItem
  bound: LayoutDimension
  base: LayoutDimension
}>()

const emit = defineEmits<{
  (e: 'update:model-value', value: LayoutDataItem): void
}>()

const addedComponents = computed(() => CellComponents.filter(v => Boolean(props.modelValue[v.key])))
const componentCheckStore = inject<Ref<Record<string, Record<CellComponentProp | 'border', boolean>>>>('componentCheckStore', ref({}))
const componentChecked = computed({
  get: () => componentCheckStore.value[props.modelValue.id] || getDefaultCheckStore(),
  set: (v: Record<CellComponentProp | 'border', boolean>) => {
    componentCheckStore.value[props.modelValue.id] = v
  }
})

const allChecked = computed({
  get: () => {
    const result = addedComponents.value.every(v => componentChecked.value[v.key])
    if (props.modelValue.win.showBorder) {
      return result && componentChecked.value.border
    }
    return result
  },
  set: (v: boolean) => {
    for (const comp of addedComponents.value) {
      componentChecked.value = {
        ...componentChecked.value,
        [comp.key]: v
      }
    }
    if (props.modelValue.win.showBorder) {
      componentChecked.value = {
        ...componentChecked.value,
        border: v
      }
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
        <div>显示全部组件</div>
        <VCheckbox v-model="allChecked" color="primary" circle />
      </div>
    </div>
    <div v-if="props.modelValue.win.showBorder" class="layout-row-inner">
      <div class="layout-cell">
        <div>边框</div>
        <VCheckbox
          :model-value="componentChecked.border"
          color="primary"
          circle
          @update:model-value="v => componentChecked = { ...componentChecked, border: v }"
        />
      </div>
    </div>
    <div v-for="comp in addedComponents" :key="comp.key" class="layout-row-inner">
      <div class="layout-cell">
        <div>{{ comp.name }}</div>
        <VCheckbox
          :model-value="componentChecked[comp.key]"
          color="primary"
          circle
          @update:model-value="v => componentChecked = { ...componentChecked, [comp.key]: v }"
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
