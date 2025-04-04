<script lang="ts" setup>
import { type CellComponentProp, getDefaultCheckStore } from '../utils';

const props = defineProps<{
  no: number,
  data: LayoutDataItem | null,
  bound: LayoutDimension
}>()

const componentCheckStore = inject<Ref<Record<string, Record<CellComponentProp | 'border', boolean>>>>('componentCheckStore', ref({}))
const componentChecked = computed(() => componentCheckStore.value[props.data?.id || ''] || getDefaultCheckStore())
</script>
<template>
  <div>
    <template v-if="data">
      <slot />
      <div class="layout-no">序号{{ no }}</div>
      <border-display v-if="data.win" :class="{ 'is-hidden': !componentChecked.border }" :model-value="data" :bound="bound" />
      <title-display v-if="data.title" :class-name="{ 'is-hidden': !componentChecked.title }" :model-value="data" :bound="bound" />
      <vol-display v-if="data.vol" :class="{ 'is-hidden': !componentChecked.vol }" :model-value="data" :bound="bound" />
      <vector-display v-if="data.vector" :class="{ 'is-hidden': !componentChecked.vector }" :model-value="data" :bound="bound" />
      <oscillogram-display v-if="data.oscillogram" :class="{ 'is-hidden': !componentChecked.oscillogram }" :model-value="data" :bound="bound" />
      <meta-display v-if="data.meta" :class="{ 'is-hidden': !componentChecked.meta }" :model-value="data" :bound="bound" />
      <alarm-display v-if="data.alarm" :class="{ 'is-hidden': !componentChecked.alarm }" :model-value="data" :bound="bound" />
    </template>
  </div>
</template>
<style lang="scss">
</style>
