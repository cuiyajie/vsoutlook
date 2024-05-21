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
      <div v-if="data.title" data-role="title" class="layout-title" :style="{left: `${data.title.x * bound.w}px`, top: `${data.title.y * bound.h}px`, width: `${data.title.w * bound.w}px`, height: `${data.title.h * bound.h}px`, fontSize: `${data.title.fontSize * bound.w}px`}">
        窗口名称 {{ no }}
      </div>
      <div v-if="data.vol" data-role="vol" class="layout-vol" :style="{left: `${data.vol.x * bound.w}px`, top: `${data.vol.y * bound.h}px`, height: `${data.vol.h * bound.h}px`, gap: `${data.vol.g * bound.w}px`}">
        <div v-for="(v, vi) in new Array(data.vol.cnt).fill(0)" :key="vi" class="vol" :style="{width: `${data.vol.one_w * bound.w}px`}" />
      </div>
      <div v-if="data.timer" data-role="timer" class="layout-timer" :style="{left: `${data.timer.x * bound.w}px`, top: `${data.timer.y * bound.h}px`, width: `${data.timer.w * bound.w}px`, height: `${data.timer.h * bound.h}px`, fontSize: `${data.timer.fontSize * bound.w}px`, color: data.timer.color}">
        16:00:00<br>2024-04-12
      </div>
    </template>
  </div>
</template>
<style lang="scss">
</style>
