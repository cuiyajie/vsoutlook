<script lang="ts" setup>
import { defaultAlarmColor } from "@src/components/pages/app/MtvLayouts/utils";
import { AlarmLang } from "@src/utils/enums";

const props = defineProps<{
  modelValue: LayoutDataItem,
  bound: LayoutDimension,
  activeRect?: TypedLayoutRect,
}>()

const emit = defineEmits<{
  (e: 'update:model-value', value: LayoutDataItem): void,
  (e: 'update:active-rect', value: TypedLayoutRect): void,
  (e: 'select', index: number): void,
  (e: 'reset-active-rect'): void
}>()

const mv = computed({
  get: () => props.modelValue,
  set: (v: LayoutDataItem) => emit('update:model-value', v)
})

const activeRect = computed({
  get: () => props.activeRect || { type: 'win', index: 0, x: 0, y: 0, w: 0, h: 0 },
  set: (v: TypedLayoutRect) => emit('update:active-rect', v)
})

const ratiow = computed(() => props.bound.w / mv.value.win.w)
const ratioh = computed(() => props.bound.h / mv.value.win.h)

const ax = computed(() => mv.value.alarm!.x * ratiow.value)
const ay = computed(() => mv.value.alarm!.y * ratioh.value)
const aw = computed(() => mv.value.alarm!.w * ratiow.value)
const ah = computed(() => mv.value.alarm!.h * ratioh.value)
const fz = computed(() => mv.value.alarm!.fontSize * ratiow.value)
const color = computed(() => mv.value.alarm!.color || defaultAlarmColor.Text)
const bgColor = computed(() => mv.value.alarm!.bgColor || defaultAlarmColor.Bg)
const lang = computed(() => mv.value.alarm!.lang)

const bx = computed(() => mv.value.alarm!.border.x * ratiow.value)
const by = computed(() => mv.value.alarm!.border.y * ratioh.value)
const bw = computed(() => mv.value.alarm!.border.w * ratiow.value)
const bh = computed(() => mv.value.alarm!.border.h * ratioh.value)
const borderw = computed(() => mv.value.alarm!.border.width * ratiow.value)
const borderColor = computed(() => mv.value.alarm!.border.color || defaultAlarmColor.Border)

const isResizing = ref(false)
function onResizing(x: number, y: number, pw: number, ph: number) {
  isResizing.value = true
  handleActiveRect(x, y, pw, ph)
}

function onResizeStop(x: number, y: number, pw: number, ph: number) {
  handleActiveRect(x, y, pw, ph)
  requestAnimationFrame(() => {
    isResizing.value = false
  })
}

function handleActiveRect(x: number, y: number, pw: number, ph: number) {
  if (!activeRect.value) return
  const ac = {
    ...activeRect.value,
    x: x + bx.value,
    y: y + by.value,
    w: pw || activeRect.value.w,
    h: ph || activeRect.value.h
  }
  activeRect.value = ac
  mv.value = {
    ...mv.value,
    alarm: {
      ...mv.value.alarm!,
      x: ac.x / ratiow.value,
      y: ac.y / ratioh.value,
      w: ac.w / ratiow.value,
      h: ac.h / ratioh.value
    }
  }
}

function onAlarmClick(event: MouseEvent | KeyboardEvent) {
  const target = event.target as HTMLElement
  if (target.closest('.vdr-nest')) return
  emit('select', 0)
}

const dragHandle = ref<HTMLElement | null>(null)
onClickOutside(dragHandle, (event) => {
  if (!activeRect.value || isResizing.value) return
  const target = event.target as HTMLElement
  if (target.closest('.vdr, [data-role=LayoutSetting], .vc-colorpicker')) return
  emit('reset-active-rect')
})

</script>
<template>
  <div
    v-if="mv.alarm"
    data-role="alarm"
    class="layout-alarm-container"
    :style="{
      left: `${bx}px`,
      top: `${by}px`,
      width: `${bw}px`,
      height: `${bh}px`,
      zIndex: 1,
      '--alarm-border-color': borderColor,
      '--alarm-border-width': `${borderw}px`,
      '--alarm-fade-duration': `${mv.alarm.border.interval}ms`,
    }"
    role="button"
    tabindex="0"
    @click.prevent="onAlarmClick($event)"
    @keyup.enter.prevent="onAlarmClick($event)"
  >
    <div class="layout-alarm-border">
      <div class="a-border a-border-t" />
      <div class="a-border a-border-r" />
      <div class="a-border a-border-b" />
      <div class="a-border a-border-l" />
    </div>
    <div
      data-role="alarm"
      class="layout-alarm"
      :style="{
        left: `${ax - bx}px`,
        top: `${ay - by}px`,
        width: `${aw}px`,
        height: `${ah}px`,
        fontSize: `${fz}px`,
        color,
        backgroundColor: bgColor,
      }"
      role="button"
      tabindex="0"
      @click.prevent.stop="emit('select', 1)"
      @keyup.enter.prevent.stop="emit('select', 1)"
    >
      <div v-if="lang === AlarmLang.English">Video Lost<br>Audio Lost</div>
      <div v-else-if="lang === AlarmLang.Chinese">视频丢失<br>音频丢失</div>
    </div>
    <vue-draggable-resizable
      v-if="activeRect.type === 'alarm' && activeRect.index === 1"
      class-name="vdr in-modal vdr-nest"
      :parent="true"
      :active="true"
      :prevent-deactivation="true"
      :drag-handle="`.drag-handle`"
      :max-width="bw"
      :max-height="bh"
      :min-width="10"
      :min-height="10"
      :w="activeRect.w"
      :h="activeRect.h"
      :x="activeRect.x - bx"
      :y="activeRect.y - by"
      @resizing="onResizing"
      @resize-stop="onResizeStop"
      @dragging="handleActiveRect"
      @drag-stop="handleActiveRect"
    >
      <div ref="dragHandle" class="drag-handle fixed-width" />
    </vue-draggable-resizable>
  </div>
</template>
<style lang="scss">
@keyframes alarmFade {
  0%, 100% {
    opacity: 0;
  }

  50% {
    opacity: 1;
  }
}

.layout-alarm {
  > div {
    color: inherit;
  }
}
.layout-alarm-container {
  position: absolute;
  --alarm-border-color: rgba(255, 0, 0, 1.0);
  --alarm-border-width: 1px;
  --alarm-fade-duration: 0.5s;
}

.layout-alarm-border {
  position: absolute;
  inset: 0;
  pointer-events: none;
  animation: var(--alarm-fade-duration) linear 0s infinite alarmFade;

  .a-border {
    position: absolute;
    background-color: var(--alarm-border-color);

    &.a-border-t {
      top: 0;
      left: 0;
      right: 0;
      height: var(--alarm-border-width);
    }

    &.a-border-r {
      top: 0;
      right: 0;
      bottom: 0;
      width: var(--alarm-border-width);
    }

    &.a-border-b {
      bottom: 0;
      left: 0;
      right: 0;
      height: var(--alarm-border-width);
    }

    &.a-border-l {
      top: 0;
      left: 0;
      bottom: 0;
      width: var(--alarm-border-width);
    }
  }
}
</style>
