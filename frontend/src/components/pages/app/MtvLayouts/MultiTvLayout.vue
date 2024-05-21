<script lang="ts" setup>
import { useNotyf } from "@src/composable/useNotyf";
import { confirm } from "@src/utils/dialog";
import { useElementBounding, useElementSize } from '@vueuse/core'
import { DefaultLayouts, draftTitle, draftVol, draftTimer, ds2db, db2ds, resizeTitle, resizeVol, resizeTimer } from './utils';
import { useLayout } from "@src/stores/layout";
import { LayoutSize } from "@src/utils/enums";
import IsEqual from "lodash-es/isEqual";
import { type WatchStopHandle } from "vue"

const notyf = useNotyf();
const layoutSamples = ref<[number, number][][]>(DefaultLayouts);
const currLayoutSample = ref<[number, number][] | null>(null);
const cellCount = computed(() => currLayoutSample.value ? currLayoutSample.value.reduce((acc, [row, column]) => acc + row * column, 0) : 0)

const layoutStore = useLayout();
const layouts = computed(() => layoutStore.layouts)
const router = useRouter()
const route = useRoute()
const currLayout = ref<Layout | null>(null)
const originDataset = ref<LayoutDataItem[]>([])
layoutStore.$fetchList().then(() => {
  if (route.query.layout) {
    currLayout.value = layouts.value.find(l => l.id === route.query.layout) || null
    parseLayoutContent()
  } else if (layouts.value.length > 0) {
    selectLayout(layouts.value[layouts.value.length - 1])
  }
})

watch(() => route.query.layout, () => {
  if (route.query.layout) {
    currLayout.value = layouts.value.find(l => l.id === route.query.layout) || null
    parseLayoutContent()
  }
})

const base = computed(() => {
  if (currLayout.value?.size === LayoutSize.FK) {
    return { w: 3840, h: 2160 }
  } else {
    return { w: 1920, h: 1080 }
  }
})

const contentRef = ref<HTMLElement | null>(null)
const { top, update: updateBounding } = useElementBounding(contentRef)
const layoutsRef = ref<HTMLElement | null>(null)
const { width: containerW, height: containerH } = useElementSize(layoutsRef)
const bounding = computed(() => {
  const ratio = containerW.value / containerH.value
  if (ratio > 16 / 9) {
    return {
      w: containerH.value * 16 / 9,
      h: containerH.value
    }
  } else {
    return {
      w: containerW.value,
      h: containerW.value * 9 / 16
    }
  }
})
const changed = ref(false)
const dataset = ref<LayoutDataItem[]>([]);

let unwatch: WatchStopHandle | null = null

const beforeUnloadHandler = (event: BeforeUnloadEvent) => {
  // Recommended
  event.preventDefault();

  // Included for legacy support, e.g. Chrome/Edge < 119
  event.returnValue = true;
};

watch(changed, () => {
  if (changed.value) {
    window.addEventListener('beforeunload', beforeUnloadHandler)
  } else {
    window.removeEventListener('beforeunload', beforeUnloadHandler)
  }
}, { immediate: true })

const routeGuard = () => {
  let answer = true
  if (changed.value) {
    answer = window.confirm('您有未保存的布局设置，确定要离开么？')
  }
  return answer
}

onBeforeRouteLeave(routeGuard)

function parseLayoutContent() {
  if (currLayout.value) {
    currLayoutSample.value = null
    const str = JSON.stringify(db2ds(currLayout.value.content || [], base.value))
    dataset.value = JSON.parse(str)
    originDataset.value = JSON.parse(str)
    unwatch?.()
    changed.value = false
    activeCell.value = null
    unwatch = watch(dataset, () => {
      changed.value = !IsEqual(dataset.value, originDataset.value)
    }, { deep: true })
  }
}

function selectLayout(ly: Layout) {
  const answer = routeGuard()
  if (!answer) return
  currLayout.value = ly
  router.push({ query: { layout: ly.id } })
  parseLayoutContent()
}

function initLayout() {
  dataset.value = Array.from({ length: cellCount.value }, () => ({
    win: { x: 0, y: 0, w: 0, h: 0 },
    title: draftTitle(base.value.w, base.value.h, bounding.value),
    vol: draftVol(base.value.h, bounding.value),
    timer: null
  }))
}

function redrawLayout() {
  nextTick(() => {
    const guideTable = layoutsRef.value?.querySelector('[data-role="guide-table"]') as HTMLElement
    if (!guideTable) return
    const origin = guideTable.getBoundingClientRect()
    const cells = guideTable.querySelectorAll('[data-role="guide-cell"]')
    const { w, h } = bounding.value
    cells.forEach((cell, index) => {
      const { x, y, width, height } = cell.getBoundingClientRect()
      dataset.value[index].win = { x: (x - origin.x) / w, y: (y - origin.y) / h, w: width / w, h: height / h }
      dataset.value[index].title = draftTitle(width, height, bounding.value)
      dataset.value[index].vol = draftVol(height, bounding.value)
      dataset.value[index].timer = null
    })
    originDataset.value = JSON.parse(JSON.stringify(dataset.value))
  })
}

const activeCell = ref<IndexedLayoutRect | null>(null)
const ads = computed<LayoutDataItem | null>({
  get: () => activeCell.value ? dataset.value[activeCell.value.index] : null,
  set: (v) => {
    if (!activeCell.value || !v) return
    const { w, h } = bounding.value
    activeCell.value = {
      ...activeCell.value,
      x: v.win.x * w,
      y: v.win.y * h,
      w: v.win.w * w,
      h: v.win.h * h
    }
    resizeComponents(v.win, dataset.value[activeCell.value.index].win)
    dataset.value[activeCell.value.index] = v
  }
})

function resizeComponents(nv: LayoutDimension, ov: LayoutDimension) {
  if (nv && ov && (nv.w !== ov.w || nv.h !== ov.h)) {
    const rw = nv.w / ov.w
    const rh = nv.h / ov.h
    if (activeCell.value) {
      const _ads = dataset.value[activeCell.value.index]
      Object.assign(dataset.value[activeCell.value.index], {
        title: _ads.title ? resizeTitle(_ads.title, rw, rh) : null,
        vol: _ads.vol ? resizeVol(_ads.vol, rw, rh) : null,
        timer: _ads.timer ? resizeTimer(_ads.timer, rw, rh) : null
      })
    }
  }
}

function resetActiveCell() {
  if (!activeCell.value) return
  const aidx = activeCell.value.index
  dataset.value[aidx] = JSON.parse(JSON.stringify(originDataset.value[aidx]))
  selectCell(dataset.value[aidx], aidx)
}

function handleActiveCell(x: number, y: number, pw: number, ph: number) {
  if (!activeCell.value) return
  activeCell.value = {
    ...activeCell.value,
    x,
    y,
    w: pw || activeCell.value.w,
    h: ph || activeCell.value.h
  }
  const { w, h } = bounding.value
  const ac = activeCell.value
  const nv = {
    x: ac.x / w,
    y: ac.y / h,
    w: ac.w / w,
    h: ac.h / h
  }
  resizeComponents(nv, dataset.value[activeCell.value.index].win)
  dataset.value[activeCell.value.index].win = nv
}

function selectCell(di: LayoutDataItem, didx: number) {
  const { w, h } = bounding.value
  activeCell.value = {
    index: didx,
    x: di.win.x * w,
    y: di.win.y * h,
    w: di.win.w * w,
    h: di.win.h * h
  }
}

function deleteCell(didx: number) {
  if (activeCell.value?.index === didx) {
    activeCell.value = null
  }
  dataset.value.splice(didx, 1)
}

function editCell(didx: number) {
  bus.trigger(Signal.OpenLayoutCellSetting, {
    item: dataset.value[didx],
    callbacks: {
      update: (v: LayoutDataItem) => {
        dataset.value[didx] = v
      }
    }
  })
}

// const dragHandle = ref<HTMLElement | null>(null)
// onClickOutside(dragHandle, (event) => {
//   if (!activeCell.value) return null
//   const target = event.target as HTMLElement
//   if (target.closest('.vdr, [role=dialog].layout-cell-setting, [data-role=LayoutSetting]')) return
//   activeCell.value = null
// })

function createWinCell(timer: boolean) {
  const { w, h } = bounding.value
  const cell: LayoutDataItem = {
    win: { x: 0.25, y: 0.25, w: 0.5, h: 0.5 },
    title: null,
    timer: null,
    vol: null
  }
  if (timer) {
    cell.timer = draftTimer(w * 0.5, bounding.value)
  } else {
    cell.title = draftTitle(w * 0.5, h * 0.5, bounding.value)
    cell.vol = draftVol(h * 0.5, bounding.value)
  }
  return cell
}

function createWin() {
  if (!currLayoutSample.value && dataset.value.length === 0) {
    notyf.error('请先选择布局')
    return
  }
  const cell = createWinCell(false)
  dataset.value.push(cell)
  selectCell(cell, dataset.value.length - 1)
}

function createTimerWin() {
  if (!currLayoutSample.value && dataset.value.length === 0) {
    notyf.error('请先选择布局')
    return
  }
  if (dataset.value.filter(d => d.timer).length > 0) {
    notyf.error('只能添加一个时间窗口')
    return
  }
  const cell = createWinCell(true)
  dataset.value.push(cell)
  selectCell(cell, dataset.value.length - 1)
}

function selectLayoutSample(ly: [number, number][]) {
  currLayoutSample.value = ly
  initLayout()
  redrawLayout()
}

function confirmSelectLayout(ly: [number, number][]) {
  if (currLayoutSample.value || dataset.value.length > 0) {
    confirm({
      title: "选择布局",
      content: "确认选择该布局替换现有布局么？",
      size: 'small',
      onConfirm: async (hide) => {
        hide()
        selectLayoutSample(ly)
      },
    });
  } else {
    selectLayoutSample(ly)
  }
}

watch(currLayout, (newV, oldV) => {
  if (!oldV && newV) {
    setTimeout(() => {
      updateBounding()
    }, 200)
  }
}, { immediate: true })

function createLayout() {
  bus.trigger(Signal.OpenNewLayout)
}

function saveAs() {
  if (!currLayout.value) return
  bus.trigger(Signal.LayoutSaveAs, { layout: currLayout.value })
}

function editLayout(ly: Layout) {
  bus.trigger(Signal.OpenNewLayout, { layout: ly })
}

function deleteLayout(ly: Layout) {
  confirm({
    title: "删除布局",
    content: "确认删除？",
    size: 'small',
    onConfirm: async (hide) => {
      const res = await layoutStore.$deleteLayout(ly.id)
      if (res === 'success') {
        notyf.success('删除成功')
        hide()
      } else {
        notyf.error(res || '删除失败')
      }
    },
  });
}

async function saveLayoutData() {
  if (!currLayout.value) return
  const res = await layoutStore.$updateContent(currLayout.value.id, ds2db(dataset.value, base.value))
  if (res) {
    notyf.success('保存布局成功')
    currLayout.value = res
    parseLayoutContent()
  } else {
    notyf.error('保存布局失败')
  }
}

function publish(ly: Layout) {
  confirm({
    title: "推送布局",
    content: "确认推送？推送后可在部署多画面应用时选择该布局",
    size: 'small',
    onConfirm: async (hide) => {
      const res = await layoutStore.$publish(ly)
      if (res === 'success') {
        notyf.success('推送成功')
        hide()
      } else {
        notyf.error(res || '推送失败')
      }
    },
  });
}
</script>
<template>
  <div>
    <Transition name="translate-page-y">
      <VCard v-if="currLayout" radius="rounded" class="tv-header-card">
        <div v-for="(ly, id1) in layoutSamples" :key="id1" class="tv-tables mini" role="button" tabindex="-1" @click.prevent="confirmSelectLayout(ly)" @keyup.enter.prevent="selectLayoutSample(ly)">
          <div v-for="(area, id2) in ly" :key="id2" class="tv-table" :style="{ '--row': area[1], '--column': area[0], width: '100%', height: `${100 / ly.length}%`, borderTopWidth: id2 === 0 ? '1px' : '0px' }">
            <div v-for="(row, id3) in area[0] * area[1]" :key="id3" class="tv-cell" />
          </div>
          <div class="add-mask"><i aria-hidden="true" class="fas fa-plus" /></div>
        </div>
        <div class="area-sep" />
        <div class="area-btns">
          <div role="button" tabindex="-1" class="area-btn" @click.prevent="createWin" @keyup.enter.prevent="createWin">
            <span class="icon mr-2">
              <i aria-hidden="true" class="fas fa-plus" />
            </span>
            <span>新建默认窗口</span>
          </div>
          <div role="button" tabindex="-1" class="area-btn" @click.prevent="createTimerWin" @keyup.enter.prevent="createTimerWin">
            <span class="icon mr-2">
              <i aria-hidden="true" class="fas fa-plus" />
            </span>
            <span>新建时间窗口</span>
          </div>
        </div>
      </VCard>
    </Transition>
    <div ref="contentRef" class="columns is-multiline tv-page" :style="{height:`calc(100vh - ${top}px)`}">
      <div class="column is-2">
        <VCard radius="rounded">
          <h3 class="title is-5 mb-4">
            布局列表
          </h3>
          <layout-list :curr-layout="currLayout" :layouts="layouts" @select="selectLayout" @delete="deleteLayout" @publish="publish" @edit="editLayout" />
        </VCard>
      </div>
      <div class="column is-8">
        <VCard
          radius="rounded"
          class="layouts-middle"
        >
          <h3 class="title is-5 mb-0 layouts-header">
            <span>
              {{ currLayout ? currLayout.name : '布局' }}
              <span v-if="currLayout" class="ml-2" style="font-size: 1rem; font-weight: 500;">{{ currLayout.size === LayoutSize.HD ? 'HD' : '4K' }}_{{ currLayout.id }}</span>
            </span>
            <div class="buttons">
              <VButtons>
                <VButton color="primary" raised @click="createLayout">
                  <span class="icon">
                    <i aria-hidden="true" class="fas fa-plus" />
                  </span>
                  <span>新建布局</span>
                </VButton>
                <template v-if="currLayout">
                  <VButton color="primary" raised @click="saveAs">
                    <span class="icon">
                      <i aria-hidden="true" class="fas fa-copy" />
                    </span>
                    <span>复制</span>
                  </VButton>
                  <VButton color="primary" raised @click="publish(currLayout)">
                    <span class="icon">
                      <i class="fas fa-cloud-upload-alt" aria-hidden="true" />
                    </span>
                    <span>推送</span>
                  </VButton>
                  <VButton v-if="dataset.length > 0" class="is-final" color="info" raised @click="saveLayoutData">
                    <span class="icon">
                      <i aria-hidden="true" class="fas fa-save" />
                    </span>
                    <span>保存设置</span>
                  </VButton>
                </template>
              </VButtons>
            </div>
          </h3>
          <div ref="layoutsRef" class="layouts-container">
            <div v-if="!currLayout || (!currLayoutSample && dataset.length === 0)" class="is-empty">
              <h3 class="dark-inverted">{{ !currLayout ? '请选择一个布局' : '请选择布局类型添加到该区域' }}</h3>
            </div>
            <div v-else class="layouts-area layout-display" :style="{width: `${bounding.w}px`, height: `${bounding.h}px`}">
              <div v-if="currLayoutSample" class="tv-tables normal" data-role="guide-table">
                <div
                  v-for="(area, id2) in currLayoutSample"
                  :key="id2"
                  class="tv-table"
                  :style="{
                    '--row': area[1],
                    '--column': area[0],
                    width: '100%',
                    height: `${100 / currLayoutSample.length}%`,
                    borderTopWidth: id2 === 0 ? '1px' : '0px'
                  }"
                >
                  <div
                    v-for="(row, id3) in area[0] * area[1]"
                    :key="id3"
                    tabindex="0"
                    role="button"
                    class="tv-cell"
                    data-role="guide-cell"
                  />
                </div>
              </div>
              <div
                v-for="(ds, cidx) in dataset"
                :key="cidx"
                role="button"
                tabindex="0"
                class="data-cell"
                :style="{left: `${ds.win.x * 100}%`, top: `${ds.win.y * 100}%`, width: `${ds.win.w * 100}%`, height: `${ds.win.h * 100}%`}"
                :class="{actived: activeCell && activeCell.index === cidx}"
                @click.prevent="selectCell(ds, cidx)"
                @keyup.enter.prevent="selectCell(ds, cidx)"
              >
                <layout-display :no="cidx + 1" :bound="bounding" :data="ds">
                  <VIconButton color="white" outlined circle icon="feather:x" class="cell-x" @click.prevent="deleteCell(cidx)" />
                  <VIconButton color="white" outlined circle icon="feather:edit-3" class="cell-edit" @click.prevent="editCell(cidx)" />
                </layout-display>
              </div>
              <vue-draggable-resizable
                v-if="activeCell"
                :parent="true"
                :active="true"
                :prevent-deactivation="true"
                :drag-handle="`.drag-handle`"
                :max-width="bounding.w"
                :max-height="bounding.h"
                :min-width="100"
                :min-height="100"
                :w="activeCell.w"
                :h="activeCell.h"
                :x="activeCell.x"
                :y="activeCell.y"
                @resizing="handleActiveCell"
                @resize-stop="handleActiveCell"
                @dragging="handleActiveCell"
                @drag-stop="handleActiveCell"
              >
                <div ref="dragHandle" class="drag-handle">
                  <layout-display :no="activeCell.index + 1" :bound="bounding" :data="ads">
                    <VIconButton color="white" outlined circle icon="feather:x" class="cell-x" @click.prevent="deleteCell(activeCell.index)" />
                    <VIconButton color="white" outlined circle icon="feather:edit-3" class="cell-edit" @click.prevent="editCell(activeCell.index)" />
                  </layout-display>
                </div>
              </vue-draggable-resizable>
            </div>
          </div>
        </VCard>
      </div>
      <div class="column is-2" data-role="LayoutSetting">
        <VCard radius="rounded">
          <layout-setting v-if="ads" v-model="ads" :bound="bounding" :base="base" @reset="resetActiveCell" />
          <div v-else className="layout-panel">
            <div class="layout-header mb-4">
              <h3 class="title is-5">
                属性设置
              </h3>
            </div>
          </div>
        </VCard>
      </div>
    </div>
    <NewLayoutModal />
    <LayoutSaveModal />
    <LayoutCellModal :index="activeCell?.index || 0" :parent="bounding" :base="base" />
  </div>
</template>
<style lang="scss">
.tv-header-card {
  margin-inline-start: -0.5rem;
  margin-inline-end: -0.5rem;
  margin-bottom: 1rem;
  height: 8rem;
  display: flex;
  justify-content: center;
  align-items: center;
  gap: 24px;

  &.l-card {
    padding: 0;
    width: calc(100% + 1rem);
  }

  .area-sep {
    width: 1px;
    height: 80%;
    background: var(--dark-sidebar-light-18);
  }

  .area-btns {
    display: flex;
    gap: 12px;
    flex-direction: column;
    justify-content: center;
    align-items: center;
  }

  .area-btn {
    width: 150px;
    height: 36px;
    display: flex;
    justify-content: center;
    align-items: center;
    border: 1px solid #fff;
    cursor: pointer;
    transition: background 0.3s;

    &:hover {
      background: var(--dark-sidebar-light-2);
    }
  }
}

.tv-tables {
  --border-color: #fff;

  &.mini {
    --border-color: #fff;

    width: 176px;
    height: 99px;
    position: relative;

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
      background-color: rgba(255, 255, 255, 0.3);
      backdrop-filter: blur(2px);
      display: flex;
      justify-content: center;
      align-items: center;

      .fas {
        font-size: 1.5rem;
        color: #fff;
      }
    }
  }

  &.normal {
    --border-color: var(--dark-sidebar-light-18);

    .tv-table,
    .tv-cell {
      border: 0;

      &:before {
        content: '';
        position: absolute;
        top: 0;
        left: 0;
        width: 100%;
        height: 100%;
      }
    }

    .tv-table {
      &:before {
        border-top: 1px solid var(--border-color);
        border-left: 1px solid var(--border-color);
      }
    }

    .tv-cell {
      &:before {
        border-right: 1px solid var(--border-color);
        border-bottom: 1px solid var(--border-color);
      }
    }
  }

  .tv-table {
    border-top: 1px solid var(--border-color);
    border-left: 1px solid var(--border-color);
    display: grid;
    grid-template-columns: repeat(var(--column), 1fr);

    .tv-cell {
      border-right: 1px solid var(--border-color);
      border-bottom: 1px solid var(--border-color);
      position: relative;
    }
  }
}

.tv-header,
.tv-page {
  padding-bottom: 1rem;

  .column {
    padding: 0.25rem;
    height: 100%;

    > .l-card {
      height: 100%;
    }
  }
}

.layouts-middle {
  display: flex;
  flex-direction: column;

  .layouts-header {
    flex: 0 0 auto;
  }

  .layouts-container {
    flex: 1;
    min-height: 0;
  }
}

.layouts-header {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;

  .buttons {
    display: flex;
    gap: 8px;
    margin-top: -5px;

    .button.is-final {
      margin-left: 16px;
    }
  }
}

.layouts-container {
  margin-top: 2rem;
  display: flex;
  justify-content: center;
  align-items: center;

  .is-empty {
    display: flex;
    justify-content: center;
    align-items: center;
    width: 100%;
    height: 100%;
    border: 1px dashed #ccc;
    border-radius: 8px;
  }
}

.layouts-area {
  position: relative;
  aspect-ratio: 16 / 9;

  .tv-tables {
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
  }

  .data-cell {
    position: absolute;
    background-color: rgba(0, 0, 0, 0.1);
    border: 1px solid rgba(255, 255, 255, 0.5);
    overflow: hidden;

    &.actived {
      visibility: hidden;
    }
  }

  .cell-title {
    position: absolute;
    top: 4px;
    left: 50%;
    transform: translateX(-50%);
    display: flex;
    justify-content: center;
    align-items: center;
    font-size: 0.75rem;
    color: var(--dark-inverted);
  }

  .button {
    width: 20px;
    height: 20px;
    padding: 0;
  }

  .cell-x {
    position: absolute;
    top: 4px;
    right: 4px;
  }

  .cell-edit {
    position: absolute;
    bottom: 4px;
    right: 4px;
  }
}

.layout-display {

  .layout-title {
    position: absolute;
    display: flex;
    align-items: center;
    justify-content: center;
    background-color: var(--dark-sidebar-light-10);
    text-wrap: nowrap;
  }

  .layout-vol {
    position: absolute;
    display: flex;
    align-items: center;

    .vol {
      height: 100%;
      background: linear-gradient(to top, #ff4500, #ff8c00, #00ff00);
    }
  }

  .layout-timer {
    position: absolute;
    display: flex;
    align-items: center;
    justify-content: center;
    text-align: center;
    background: var(--dark-sidebar-light-2);
    text-wrap: nowrap;
  }

  &.draggable {
    width: 100%;
    height: 100%;

    .layout-title,
    .layout-vol,
    .layout-timer {
      cursor: pointer;
      user-select: none;
    }
  }
}

.vdr {
  position: absolute;
  background-color: #000;
  border: 1px dashed var(--primary);

  &.in-modal {
    background-color: transparent;
  }

  .drag-handle {
    position: absolute;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
  }
}
</style>
