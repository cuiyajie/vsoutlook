<script lang="ts" setup>
import { format } from 'date-fns'
const props = defineProps<{
  no: number,
  data: LayoutDataItem | null,
  bound: LayoutDimension
}>()
const time = new Date()

const dateText = computed(() => {
  const t = props.data?.timer
  if (!t) return ""
  let str = ""
  if (t.showDate) {
    str = `${str}${format(time, t.dateDisplayType === 0 ? "yyyy年MM月dd日" : "yyyy-MM-dd")}${t.dateNewLine === 1 ? "<br>" : " "}`
  }
  str = `${str}${format(time, t.time24 === 24 ? "HH:mm:ss" : "hh:mm:ss")}`
  if (t.showFrame === 0) {
    str = `${str}:00`
  }
  return str
})
</script>
<template>
  <div v-if="data?.timer" data-role="timer" class="layout-timer" :style="{fontSize: `${data.timer.fontSize * bound.w}px`, color: data.timer.color, fontFamily: data.timer.fontFamily}" style="max-width: 100%;">
    <slot />
    <span style="max-width: 100%; overflow: hidden;" v-html="dateText" />
  </div>
</template>
<style lang="scss">
</style>
