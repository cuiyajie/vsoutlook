<script lang="ts" setup>
const props = defineProps<{
  modelValue: LayoutDataItem,
  bound: LayoutDimension
}>()

const emit = defineEmits<{
  (e: 'update:model-value', value: LayoutDataItem): void,
  (e: 'active', type: LayoutProps | null): void
}>()

const mv = computed({
  get: () => props.modelValue,
  set: (v: LayoutDataItem) => emit('update:model-value', v)
})

const ratiow = computed(() => props.bound.w / mv.value.win.w)
const ratioh = computed(() => props.bound.h / mv.value.win.h)
const defWin = () => ({ type: 'win', index: 0, x: 0, y: 0, w: 1, h: 1 } as TypedLayoutRect)
const activeComponent = ref<TypedLayoutRect>(defWin())
const isResizing = ref(false)
const isAlarmBorder = computed(() => activeComponent.value?.type === 'alarm' && activeComponent.value.index === 0)
const isAlarmText = computed(() => activeComponent.value?.type === 'alarm' && activeComponent.value.index === 1)
const alarmTextWidth = computed(() => mv.value.alarm!.w * ratiow.value)
const alarmTextHeight = computed(() => mv.value.alarm!.h * ratioh.value)

function onResizing(x: number, y: number, pw: number, ph: number) {
  isResizing.value = true
  handleComponent(x, y, pw, ph)
}

function onResizeStop(x: number, y: number, pw: number, ph: number) {
  handleComponent(x, y, pw, ph)
  requestAnimationFrame(() => {
    isResizing.value = false
  })
}

function handleComponent(x: number, y: number, pw: number, ph: number) {
  if (!activeComponent.value) return
  activeComponent.value = {
    ...activeComponent.value,
    x,
    y,
    w: pw || activeComponent.value.w,
    h: ph || activeComponent.value.h
  }
  const ac = activeComponent.value
  if (activeComponent.value.type === 'vol') {
    const idx = ac.index || 0
    mv.value = {
      ...mv.value,
      vol: {
        ...mv.value.vol!,
        vols: mv.value.vol!.vols.map((vol, i) => {
          if (i === idx) {
            return {
              ...vol,
              x: ac.x / ratiow.value,
              y: ac.y / ratioh.value,
              h: ac.h / ratioh.value,
            }
          }
          return vol
        })
      }
    }
  } else if (activeComponent.value.type === 'alarm') {
    if (ac.index === 0) {
      mv.value = {
        ...mv.value,
        alarm: {
          ...mv.value.alarm!,
          border: {
            ...mv.value.alarm!.border,
            x: ac.x / ratiow.value,
            y: ac.y / ratioh.value,
            w: ac.w / ratiow.value,
            h: ac.h / ratioh.value
          }
        }
      }
    } else {
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
  } else {
    mv.value = {
      ...mv.value,
      [ac.type]: {
        ...mv.value[ac.type],
        x: ac.x / ratiow.value,
        y: ac.y / ratioh.value,
        w: ac.w / ratiow.value,
        h: ac.h / ratioh.value
      }
    }
  }
}

function syncComponentData(type: LayoutProps, componentIndex?: number) {
  if (type === 'vol') {
    syncVolData(componentIndex || 0)
  } else if (type === 'alarm') {
    syncAlarmData(componentIndex || 0)
  } else {
    let rect = mv.value[type] as LayoutRect
    if (!rect) return
    activeComponent.value = {
      type,
      index: 0,
      x: rect.x * ratiow.value,
      y: rect.y * ratioh.value,
      w: rect.w * ratiow.value,
      h: rect.h * ratioh.value,
      rotated: type === 'title' && (rect as LayoutTitle).isVertial
    }
  }
}

function syncVolData(index: number) {
  const vol = mv.value.vol!.vols[index]
  const cnt = vol.end - vol.start + 1
  const rectw = vol.one_w * cnt + vol.g * (cnt - 1)
  activeComponent.value = {
    type: 'vol',
    index,
    x: vol.x * ratiow.value,
    y: vol.y * ratioh.value,
    w: rectw * ratiow.value,
    h: vol.h * ratioh.value
  }
}

function syncAlarmData(index: number) {
  const rect = (index === 0 ? mv.value.alarm!.border : mv.value.alarm) as LayoutRect
  activeComponent.value = {
    type: 'alarm',
    index,
    x: rect.x * ratiow.value,
    y: rect.y * ratioh.value,
    w: rect.w * ratiow.value,
    h: rect.h * ratioh.value
  }
}

watch(mv, () => {
  if (!activeComponent.value) return
  syncComponentData(activeComponent.value.type, activeComponent.value.index)
}, { immediate: true })

watch(() => mv.value.vol?.len, (nv) => {
  if (!nv) return
  const ac = activeComponent.value
  if (ac && ac.type === 'vol' && ac.index! >= nv) {
    selectComponent('win')
  }
}, { immediate: true })

const dragHandle = ref<HTMLElement | null>(null)
onClickOutside(dragHandle, (event) => {
  if (!activeComponent.value || isResizing.value) return
  const target = event.target as HTMLElement
  if (target.closest('.vdr, [data-role=LayoutSetting], .vc-colorpicker')) return
  selectComponent('win')
})

function selectComponent(type: LayoutProps, componentIndex = 0) {
  activeComponent.value = defWin()
  nextTick(() => {
    syncComponentData(type, componentIndex)
  })
}

const resizeHandles = computed(() => {
  const ac = activeComponent.value
  if (!ac || ac.type === 'win') return []
  if (ac.type === 'vol') return ['tm', 'bm']
  if (ac.type === 'vector') return ['tl', 'tr', 'br', 'bl']
  return ['tl', 'tm', 'tr', 'mr', 'br', 'bm', 'bl', 'ml']
})

watch(activeComponent, (v) => {
  emit('active', v?.type || null)
}, { immediate: true })

defineExpose({
  selectComponent,
  clearComponent: () => selectComponent('win')
})
</script>
<template>
  <div class="layout-display draggable">
    <border-display v-model="mv" :bound="bound" />
    <title-display v-if="mv.title" v-model="mv" :bound="bound" @select="selectComponent('title')" />
    <vol-display v-if="mv.vol" v-model="mv" :bound="bound" @select="selectComponent('vol', $event)" />
    <vector-display v-if="mv.vector" v-model="mv" :bound="bound" @select="selectComponent('vector')" />
    <oscillogram-display v-if="mv.oscillogram" v-model="mv" :bound="bound" @select="selectComponent('oscillogram')" />
    <meta-display v-if="mv.meta" v-model="mv" :bound="bound" @select="selectComponent('meta')" />
    <alarm-display
      v-if="mv.alarm"
      v-model="mv"
      v-model:active-rect="activeComponent"
      :bound="bound"
      @select="selectComponent('alarm', $event)"
      @reset-active-rect="selectComponent('win')"
    />
    <vue-draggable-resizable
      v-if="activeComponent && activeComponent.type !== 'win' && !isAlarmText"
      class-name="vdr in-modal"
      :style="{'--vdr-rotate': activeComponent.rotated ? '90deg' : '0deg'}"
      :handles="resizeHandles"
      :parent="true"
      :active="true"
      :prevent-deactivation="true"
      :drag-handle="`.drag-handle`"
      :draggable="!isAlarmBorder"
      :max-width="bound.w"
      :max-height="bound.h"
      :min-width="isAlarmBorder ? alarmTextWidth: 10"
      :min-height="isAlarmBorder ? alarmTextHeight : 10"
      :lock-aspect-ratio="activeComponent.type === 'vector'"
      :centered-scaling="isAlarmBorder"
      :w="activeComponent.w"
      :h="activeComponent.h"
      :x="activeComponent.x"
      :y="activeComponent.y"
      @resizing="onResizing"
      @resize-stop="onResizeStop"
      @dragging="handleComponent"
      @drag-stop="handleComponent"
    >
      <div ref="dragHandle" class="drag-handle fixed-width" />
    </vue-draggable-resizable>
  </div>
</template>
<style lang="scss">
</style>
