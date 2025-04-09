<script setup lang="ts">
import { draftVol, draftTitle, draftVector, draftOscillogram, draftMeta, draftAlarm, CellComponents as lcControls, draftWinBorder, draftWin } from '@src/components/pages/app/MtvLayouts/utils';
import type { GlobalComponents } from "vue";
import { confirm } from "@src/utils/dialog";

const props = defineProps<{
  index: number,
  parent: LayoutDimension,
  base: LayoutDimension
}>()

const opened = ref(false);
const PADDING = 16
let callbacks: any = null

const data = ref<LayoutDataItem | null>(null)
const layoutContainer = ref<HTMLElement | null>(null)
const { width: containerW } = useElementSize(layoutContainer)
const bounding = computed(() => {
  if (!data.value) return { w: 0, h: 0 }
  const ratio = (data.value.win.h * props.base.h) / (data.value.win.w * props.base.w)
  let rw = containerW.value - PADDING * 2
  let rh = rw * ratio
  return { w: rw, h: rh }
})
watch(bounding, () => {
  nextTick(() => {
    if (layoutContainer.value) {
      const parent = layoutContainer.value.parentElement
      if (parent) {
        parent.style.setProperty('--available-height', `${parent.getBoundingClientRect().height}px`)
      }
    }
  })
}, { immediate: true })

const { pause, resume } = useIntervalFn(() => {
  callbacks?.update?.(data.value)
}, 1000, { immediate: false })

const originData = ref<LayoutDataItem | null>(null)
useListener(Signal.OpenLayoutCellSetting, (p: any) => {
  opened.value = true;
  data.value = p.item;
  originData.value = p.item ? JSON.parse(JSON.stringify(p.item)) : null;
  callbacks = p.callbacks;
  resume()
});

const activeType = ref<LayoutProps | null>(null)
function activeComponent(type: LayoutProps | null) {
  activeType.value = type
}

const displayRef = ref<InstanceType<GlobalComponents['LayoutCellDisplay']> | null>(null)
function resetCell() {
  if (!originData.value) return
  let updates: Partial<LayoutDataItem> = {}
  const win = data.value?.win || draftWin(0, 0, 0, 0, bounding.value)
  const winW = win.w * bounding.value.w
  const winH = win.h * bounding.value.h
  switch(activeType.value) {
    case 'win':
      updates = {
        win: {
          ...(data.value?.win || {}),
          showBorder: false,
          border: draftWinBorder(winW, bounding.value)
        } as LayoutDataItem['win']
      }
      break
    case 'title':
      updates = { title: draftTitle(winW, winH, bounding.value) }
      break
    case 'vol':
      updates = { vol: draftVol(winW, winH, bounding.value) }
      break
    case 'vector':
      updates = { vector: draftVector(winW, winH, bounding.value) }
      break
    case 'oscillogram':
      updates = { oscillogram: draftOscillogram(winW, winH, bounding.value) }
      break
    case 'alarm':
      updates = { alarm: draftAlarm(winW, winH, bounding.value) }
      break
    case 'meta':
      updates = { meta: draftMeta(winW, winH, bounding.value) }
      break
  }
  data.value = { ...data.value, ...updates } as LayoutDataItem
  setTimeout(() => {
    displayRef.value?.syncComponentData()
  }, 200)
}

function deleteCell() {
  if (!activeType.value || !data.value || activeType.value === 'win') return
  data.value[activeType.value] = null
  activeComponent(null)
  displayRef.value?.clearComponent()
}

function toggleComponent(type: Exclude<LayoutProps, 'win' | 'text' | 'timer'>) {
  const lcControl = lcControls.find(lc => lc.key === type)
  if (!lcControl) return
  if (data.value?.[type]) {
    confirm({
      title: "删除组件",
      content: `确定要删除 ${lcControl.name} 组件吗？`,
      onConfirm: (hide) => {
        if (!data.value) return
        data.value[type] = null
        if (activeType.value === type) {
          activeComponent(null)
          displayRef.value?.clearComponent()
        }
        hide()
      },
    });
  } else {
    confirm({
      title: "添加组件",
      content: `确定要添加 ${lcControl.name} 组件吗？`,
      onConfirm: (hide) => {
        if (!data.value) return
        const w = data.value.win.w * props.parent.w
        const h = data.value.win.h * props.parent.h
        if (type === 'vol') {
          data.value[type] = draftVol(w, h, props.parent)
        } else if (type === 'title') {
          data.value[type] = draftTitle(w, h, props.parent)
        } else if (type === 'vector') {
          data.value[type] = draftVector(w, h, props.parent)
        } else if (type === 'oscillogram') {
          data.value[type] = draftOscillogram(w, h, props.parent)
        } else if (type === 'meta') {
          data.value[type] = draftMeta(w, h, props.parent)
        } else if (type === 'alarm') {
          data.value[type] = draftAlarm(w, h, props.parent)
        }
        displayRef.value?.selectComponent(type, 0)
        hide()
      },
    });
  }
}

function onClose() {
  opened.value = false
  callbacks?.update?.(data.value)
  pause()
}
</script>
<template>
  <VModal
    id="layout-cell-modal"
    size="big"
    noclose
    :open="opened"
    :title="`窗口序号${index}`"
    cancel-label="取消"
    dialog-class="layout-cell-setting"
    @close="onClose"
  >
    <template #content>
      <div class="lc-body">
        <div ref="layoutContainer" class="lc-display" :style="{'--padding': `${PADDING}px`}">
          <div v-if="!data?.timer" class="lc-controls">
            <div
              v-for="lc in lcControls"
              :key="lc.key"
              class="lc-control"
              :class="data?.[lc.key] && 'selected'"
              role="button"
              tabindex="-1"
              @click.prevent="toggleComponent(lc.key)"
              @keyup.enter.prevent="toggleComponent(lc.key)"
            >
              <i v-if="lc.iconType === 'feather'" class="iconify" :data-icon="`feather:${lc.icon}`" aria-hidden="true" />
              <i v-else :class="`fas fa-${lc.icon}`" aria-hidden="true" />
              {{ lc.name }}
              <div v-if="!data?.[lc.key]" class="add-mask"><i aria-hidden="true" class="fas fa-plus" /></div>
              <div v-else class="add-mask"><i aria-hidden="true" class="fas fa-times" /></div>
            </div>
          </div>
          <div v-else :style="{ height: `${PADDING}px` }" />
          <div class="lc-display-inner" :style="{ width: `${bounding.w}px`, height: `${bounding.h}px` }">
            <layout-cell-display v-if="data" ref="displayRef" v-model="data" :bound="bounding" @active="activeComponent" />
          </div>
        </div>
        <div class="lc-divider" />
        <div class="lc-setting is-scrollable" data-role="LayoutSetting">
          <layout-cell-setting v-if="data" v-model="data" :type="activeType" :base="base" :bound="bounding" @reset="resetCell" @delete="deleteCell" />
        </div>
      </div>
    </template>
    <template #action />
  </VModal>
</template>
<style lang="scss">
.layout-cell-setting {
  &.modal.is-big .modal-content {
    max-width: 1280px;
  }

  .modal-card-body {
    border-bottom-left-radius: var(--radius-large);
    border-bottom-right-radius: var(--radius-large);
  }

  .lc-body {
    --available-height: 200px;
    display: flex;
    flex-wrap: nowrap;

    .lc-display {
      flex: 1;
      border: 1px dashed var(--dark-sidebar-light-20);
      border-radius: 4px;
      overflow: hidden;
      padding: 0 var(--padding)  var(--padding);
      display: flex;
      flex-direction: column;
      align-items: center;
      justify-content: center;

      .lc-controls {
        display: flex;
        justify-content: center;
        align-items: center;
        padding: var(--padding);
        gap: 16px;

        .lc-control {
          width: 128px;
          height: 40px;
          border-radius: 8px;
          background-color: var(--dark-sidebar-light-2);
          font-size: 0.875rem;
          gap: 8px;
          display: flex;
          justify-content: center;
          align-items: center;
          position: relative;
          overflow: hidden;
          transition: all 0.3s;

          &:hover {
            cursor: pointer;
            .add-mask {
              opacity: 1;
            }
          }

          .add-mask {
            position: absolute;
            opacity: 0;
            transition: opacity 0.3s;
            top: 0;
            left: 0;
            width: 100%;
            height: 100%;
            background-color: rgba(255, 255, 255, 0.1);
            backdrop-filter: blur(1px);
            display: flex;
            justify-content: center;
            align-items: center;

            .fas {
              font-size: 1.5rem;
              color: #fff;
            }
          }

          &.selected {
            background-color: var(--primary);
            color: var(--primary--light-color);
          }
        }
      }

      .lc-display-inner {
        position: relative;

        &:before {
          content: '';
          position: absolute;
          inset: 0;
          border: 1px solid rgba(255, 255, 255, .5);
          z-index: 100;
          pointer-events: none;
        }
      }
    }

    .lc-divider {
      width: 16px;
    }

    .lc-setting {
      width: 240px;
      border: 1px dashed var(--dark-sidebar-light-20);
      border-radius: 4px;
      overflow: hidden;
      max-height: var(--available-height);

      &.is-scrollable {
        overflow: auto;
      }
    }
  }

  .modal-card-foot {
    display: none;
  }
}
</style>
