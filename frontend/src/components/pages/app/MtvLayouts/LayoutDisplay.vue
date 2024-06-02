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
      <div v-if="data.title" data-role="title" class="layout-title" :style="{left: `${data.title.x * bound.w}px`, top: `${data.title.y * bound.h}px`, width: `${data.title.w * bound.w}px`, height: `${data.title.h * bound.h}px`, fontSize: `${data.title.fontSize * bound.w}px`, fontFamily:`'${data.title.fontFamily}'`}">
        窗口名称
      </div>
      <div v-if="data.vol" data-role="vol" class="layout-vol" :style="{left: `${data.vol.x * bound.w}px`, top: `${data.vol.y * bound.h}px`, height: `${data.vol.h * bound.h}px`, gap: `${data.vol.g * bound.w}px`}">
        <div v-for="(v, vi) in new Array(data.vol.cnt).fill(0)" :key="vi" class="vol" :style="{width: `${data.vol.one_w * bound.w}px`}" />
      </div>
    </template>
  </div>
</template>
<style lang="scss">
</style>
