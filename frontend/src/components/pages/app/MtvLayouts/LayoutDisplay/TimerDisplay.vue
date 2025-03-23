<script lang="ts" setup>
import { CLockDisplayType, DateDisplayType } from '@src/utils/enums'
import { format } from 'date-fns'
import { getFontFamily } from '../utils';
const props = defineProps<{
  no: number
  data: LayoutDataItem | null
  bound: LayoutDimension
  type: 'date' | 'time'
}>()
const time = new Date()

const fontFamily = computed(() => {
  const t = props.data?.timer
  const defaultFontFamily = getFontFamily()
  if (!t) return defaultFontFamily
  return t.clockDisplayType === CLockDisplayType.Digital
    ? t.fontFamily || defaultFontFamily
    : 'Digital Dismay'
})

const fontFormat = computed(() => {
  const t = props.data?.timer
  if (!t) return 'yyyy-MM-dd'
  return t.clockDisplayType === CLockDisplayType.Digital
    ? t.dateDisplayType === DateDisplayType.Chinese
      ? 'yyyy年MM月dd日'
      : 'yyyy-MM-dd'
    : 'yyyy-MM-dd'
})

const color = computed(() => {
  const t = props.data?.timer
  if (!t) return ''
  return t.clockDisplayType === CLockDisplayType.Digital ? t.color : t.timeColor
})

const dateText = computed(() => {
  const t = props.data?.timer
  if (!t) return ''
  let str = ''
  if (t.showDate && props.type === 'date') {
    str = format(time, fontFormat.value)
  }
  if (props.type === 'time') {
    str = format(time, t.time24 === 24 ? 'HH:mm:ss' : 'hh:mm:ss')
  }
  return str
})
</script>
<template>
  <div
    v-if="data?.timer && dateText"
    data-role="timer"
    class="layout-timer"
    :style="{ fontSize: `${data.timer.fontSize * bound.w}px`, color, fontFamily }"
    style="max-width: 100%"
  >
    <slot />
    <span style="max-width: 100%; overflow: hidden" v-html="dateText" />
  </div>
</template>
<style lang="scss"></style>
