<script setup lang="ts">
import { useNotyf } from "@src/composable/useNotyf";
import { draftVol, draftTitle, draftVector, draftOscillogram, draftMeta, draftAlarm } from '@src/components/pages/app/MtvLayouts/utils';
import type { GlobalComponents } from "vue";

const props = defineProps<{
  index: number,
  parent: LayoutDimension,
  base: LayoutDimension
}>()

const opened = ref(false);
const PADDING = 16
let callbacks: any = null

const notyf = useNotyf();
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
  data.value = JSON.parse(JSON.stringify(originData.value))
}

function deleteCell() {
  if (!activeType.value || !data.value || activeType.value === 'win') return
  data.value[activeType.value] = null
  activeComponent(null)
  displayRef.value?.clearComponent()
}

function addComponent(type: LayoutProps) {
  if (data.value?.[type]) {
    notyf.error('已存在该类型组件')
    return
  }
  if (!data.value) return
  const w = data.value.win.w * props.parent.w
  const h = data.value.win.h * props.parent.h
  if (type === 'vol') {
    data.value[type] = draftVol(h, props.parent)
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
            <div class="lc-control" role="button" tabindex="-1" @click.prevent="addComponent('vol')" @keyup.enter.prevent="addComponent('vol')">
              <i class="iconify" data-icon="feather:columns" aria-hidden="true" />
              UV表
              <div class="add-mask"><i aria-hidden="true" class="fas fa-plus" /></div>
            </div>
            <div class="lc-control" role="button" tabindex="-1" @click.prevent="addComponent('title')" @keyup.enter.prevent="addComponent('title')">
              <i class="iconify" data-icon="feather:airplay" aria-hidden="true" />
              窗口名称及Tally
              <div class="add-mask"><i aria-hidden="true" class="fas fa-plus" /></div>
            </div>
            <div class="lc-control" role="button" tabindex="-1" @click.prevent="addComponent('vector')" @keyup.enter.prevent="addComponent('title')">
              <i class="fas fa-chart-pie" aria-hidden="true" />
              色度矢量图
              <div class="add-mask"><i aria-hidden="true" class="fas fa-plus" /></div>
            </div>
            <div class="lc-control" role="button" tabindex="-1" @click.prevent="addComponent('oscillogram')" @keyup.enter.prevent="addComponent('title')">
              <i class="fas fa-sun" aria-hidden="true" />
              亮度波形图
              <div class="add-mask"><i aria-hidden="true" class="fas fa-plus" /></div>
            </div>
            <div class="lc-control" role="button" tabindex="-1" @click.prevent="addComponent('alarm')" @keyup.enter.prevent="addComponent('title')">
              <i class="fas fa-bell" aria-hidden="true" />
              报警
              <div class="add-mask"><i aria-hidden="true" class="fas fa-plus" /></div>
            </div>
            <div class="lc-control" role="button" tabindex="-1" @click.prevent="addComponent('meta')" @keyup.enter.prevent="addComponent('title')">
              <i class="fas fa-file-video" aria-hidden="true" />
              视频信息
              <div class="add-mask"><i aria-hidden="true" class="fas fa-plus" /></div>
            </div>
          </div>
          <div v-else :style="{ height: `${PADDING}px` }" />
          <div class="lc-display-inner" :style="{ width: `${bounding.w}px`, height: `${bounding.h}px` }">
            <layout-cell-display v-if="data" ref="displayRef" v-model="data" :bound="bounding" @active="activeComponent" />
          </div>
        </div>
        <div class="lc-divider" />
        <div class="lc-setting is-scrollable" data-role="LayoutSetting">
          <layout-cell-setting v-if="data" v-model="data" :type="activeType" :base="base" @reset="resetCell" @delete="deleteCell" />
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
