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
const activeComponent = ref<TypedLayoutRect | null>(null)

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
  mv.value = {
    ...mv.value,
    [activeComponent.value.type]: {
      ...mv.value[activeComponent.value.type],
      x: ac.x / ratiow.value,
      y: ac.y / ratioh.value,
      w: ac.w / ratiow.value,
      h: ac.h / ratioh.value
    }
  }
}

function syncComponentData(type: LayoutProps) {
  let rect = mv.value[type]
  if (!rect) return
  let rectw: number
  if (type === 'vol') {
    rect = rect as LayoutVol
    rectw = rect.one_w * rect.cnt + rect.g * (rect.cnt - 1)
  } else {
    rect = rect as LayoutTitle
    rectw = rect.w
  }
  activeComponent.value = {
    type,
    x: rect.x * ratiow.value,
    y: rect.y * ratioh.value,
    w: rectw * ratiow.value,
    h: rect.h * ratioh.value
  }
}

watch(mv, () => {
  if (!activeComponent.value) return
  syncComponentData(activeComponent.value.type)
}, { immediate: true })

// const dragHandle = ref<HTMLElement | null>(null)
// onClickOutside(dragHandle, (event) => {
//   if (!activeComponent.value) return null
//   const target = event.target as HTMLElement
//   if (target.closest('.vdr, [data-role=LayoutSetting], .vc-colorpicker')) return
//   activeComponent.value = null
// })

function selectComponent(type: LayoutProps) {
  activeComponent.value = null
  nextTick(() => {
    syncComponentData(type)
  })
}

watch(activeComponent, (v) => {
  emit('active', v?.type || null)
}, { immediate: true })

defineExpose({
  clearComponent: () => activeComponent.value = null
})
</script>
<template>
  <div class="layout-display draggable">
    <div v-if="mv.title" data-role="title" class="layout-title" :style="{left: `${mv.title.x * ratiow}px`, top: `${mv.title.y * ratioh}px`, width: `${mv.title.w * ratiow}px`, height: `${mv.title.h * ratioh}px`, fontSize: `${mv.title.fontSize * ratiow}px`, fontFamily:`'${mv.title.fontFamily}'`}" role="button" tabindex="0" @click.prevent="selectComponent('title')" @keyup.enter.prevent="selectComponent('title')">
      窗口名称
    </div>
    <div v-if="mv.vol" class="layout-vol" :style="{left: `${mv.vol.x * ratiow}px`, top: `${mv.vol.y * ratioh}px`, height: `${mv.vol.h * ratioh}px`, gap: `${mv.vol.g * ratiow}px`}" role="button" tabindex="0" @click.prevent="selectComponent('vol')" @keyup.enter.prevent="selectComponent('vol')">
      <div v-for="(v, vi) in new Array(mv.vol.cnt).fill(0)" :key="vi" class="vol" :style="{width: `${mv.vol.one_w * ratiow}px`}" />
    </div>
    <div v-if="mv.timer" data-role="timer" class="layout-timer" :style="{left: `${mv.timer.x * ratiow}px`, top: `${mv.timer.y * ratioh}px`, fontSize: `${mv.timer.fontSize * ratiow}px`, color: mv.timer.color}" role="button" tabindex="0" @click.prevent="selectComponent('timer')" @keyup.enter.prevent="selectComponent('timer')">
      16:00:00<br>2024-04-12
    </div>
    <vue-draggable-resizable
      v-if="activeComponent"
      class-name="vdr in-modal"
      :handles="activeComponent.type === 'vol' ? ['tm', 'bm'] : ['tl', 'tm', 'tr', 'mr', 'br', 'bm', 'bl', 'ml']"
      :parent="true"
      :active="true"
      :prevent-deactivation="true"
      :drag-handle="`.drag-handle`"
      :max-width="bound.w"
      :max-height="bound.h"
      :min-width="10"
      :min-height="10"
      :w="activeComponent.w"
      :h="activeComponent.h"
      :x="activeComponent.x"
      :y="activeComponent.y"
      @resizing="handleComponent"
      @resize-stop="handleComponent"
      @dragging="handleComponent"
      @drag-stop="handleComponent"
    >
      <div ref="dragHandle" class="drag-handle" />
    </vue-draggable-resizable>
  </div>
</template>
<style lang="scss">
</style>
