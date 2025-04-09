<script lang="ts" setup>
import { reDraftVol } from '@src/components/pages/app/MtvLayouts/utils';


const props = defineProps<{
  modelValue: LayoutDataItem,
  type: LayoutProps | null,
  base: LayoutDimension
  bound: LayoutDimension
}>()

const emit = defineEmits<{
  (e: 'reset'): void,
  (e: 'delete'): void,
  (e: 'update:model-value', value: LayoutDataItem): void,
}>()

const ac = computed({
  get: () => props.type ? props.modelValue[props.type] : null,
  set: (v: LayoutDataItem[LayoutProps]) => props.type && emit('update:model-value', { ...props.modelValue, [props.type]: v })
})

const relayoutVol = (len: number) => {
  const { vol, win } = props.modelValue
  if (!vol || !win) return
  emit('update:model-value', {
    ...props.modelValue,
    vol: reDraftVol(vol, len, win.w * props.bound.w, win.h * props.bound.h, props.bound)
  })
}

</script>
<template>
  <div className="layout-panel is-sticky">
    <div class="layout-header mb-4">
      <h3 class="title is-5">属性设置</h3>
      <div v-if="ac">
        <a
          v-if="type !== 'win'"
          role="button"
          class="action-link is-warning mr-4"
          tabindex="0"
          @keydown.space.prevent="emit('delete')"
          @click.prevent="emit('delete')"
        >删除</a>
        <a
          role="button"
          class="action-link"
          tabindex="0"
          @keydown.space.prevent="emit('reset')"
          @click.prevent="emit('reset')"
        >重置</a>
      </div>
    </div>
    <template v-if="ac">
      <border-setting v-if="type === 'win'" v-model="ac" :base="base" />
      <title-setting v-if="type === 'title'" v-model="ac" :base="base" />
      <vol-setting v-if="type === 'vol'" v-model="ac" :base="base" :bound="bound" @relayout="relayoutVol($event)" />
      <vector-setting v-if="type === 'vector'" v-model="ac" :base="base" />
      <oscillogram-setting v-if="type === 'oscillogram'" v-model="ac" :base="base" />
      <meta-setting v-if="type === 'meta'" v-model="ac" :base="base" />
      <alarm-setting v-if="type === 'alarm'" v-model="ac" :base="base" />
    </template>
  </div>
</template>
<style lang="scss">
.layout-panel {
}
</style>
