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
    </template>
  </div>
</template>
<style lang="scss">
</style>
