<script lang="ts" setup>
defineProps<{
  no: number,
  data: LayoutDataItem | null,
  bound: LayoutDimension
}>()
</script>
<template>
  <div>
    <template v-if="data">
      <slot />
      <div class="layout-no">序号{{ no }}</div>
      <title-display v-if="data.title" :model-value="data" :bound="bound" />
      <vol-display v-if="data.vol" :model-value="data" :bound="bound" />
      <vector-display v-if="data.vector" :model-value="data" :bound="bound" />
      <oscillogram-display v-if="data.oscillogram" :model-value="data" :bound="bound" />
      <meta-display v-if="data.meta" :model-value="data" :bound="bound" />
      <alarm-display v-if="data.alarm" :model-value="data" :bound="bound" />
    </template>
  </div>
</template>
<style lang="scss">
</style>
