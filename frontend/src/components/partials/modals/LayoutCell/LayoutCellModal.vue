<script setup lang="ts">
import { useNotyf } from "@src/composable/useNotyf";
import { draftVol, draftTitle } from '@src/components/pages/app/MtvLayouts/utils';
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
  if (!activeType.value || !data.value) return
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
  }
  activeComponent(type)
}

function onClose() {
  opened.value = false
  callbacks?.update?.(data.value)
  pause()
}
</script>
<template>
  <VModal
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
              <i class="iconify" data-icon="feather:bar-chart-2" aria-hidden="true" />
              音柱
              <div class="add-mask"><i aria-hidden="true" class="fas fa-plus" /></div>
            </div>
            <div class="lc-control" role="button" tabindex="-1" @click.prevent="addComponent('title')" @keyup.enter.prevent="addComponent('title')">
              <i class="iconify" data-icon="feather:calendar" aria-hidden="true" />
              窗口名称
              <div class="add-mask"><i aria-hidden="true" class="fas fa-plus" /></div>
            </div>
          </div>
          <div v-else :style="{ height: `${PADDING}px` }" />
          <div class="lc-display-inner" :style="{ width: `${bounding.w}px`, height: `${bounding.h}px` }">
            <layout-cell-display v-if="data" ref="displayRef" v-model="data" :bound="bounding" @active="activeComponent" />
          </div>
        </div>
        <div class="lc-divider" />
        <div class="lc-setting" data-role="LayoutSetting">
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
          width: 120px;
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
        border: 1px solid rgba(255, 255, 255, .5);
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
      padding: 16px;
    }
  }

  .modal-card-foot {
    display: none;
  }
}
</style>
