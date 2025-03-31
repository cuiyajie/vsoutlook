<script lang="ts" setup>
import { useTmplDragging } from "@src/stores/tmplDragging";
import { useClustNode } from "@src/stores/node";
import { useResourceCharts } from "@src/composable/useResourceChart";
import ApexChart from 'vue3-apexcharts'

const dragContext = useTmplDragging();
const nodeStore = useClustNode();
const { cpuOptions, memoryOptions, nnrOptions, ntrOptions } = useResourceCharts()

const nodes = computed(() => nodeStore.nodes)
const loading = ref(false)

function onDragEnter(event: DragEvent) {
  const target = event.target as HTMLElement;
  const dropzone = target.closest("[role=dropzone]") as HTMLElement;
  if (dropzone && dropzone.dataset.id) {
    if (dragContext.dragoverId) {
      if (dragContext.dragoverId !== dropzone.dataset.id) {
        dragContext.clearDropzoneEffect();
        dragContext.setDragoverId(dropzone.dataset.id);
        dropzone.classList.add("drag-over");
      }
    } else {
      dragContext.setDragoverId(dropzone.dataset.id);
      dropzone.classList.add("drag-over");
    }
  }
}

function onDragLeave(event: DragEvent) {
  const target = event.target as HTMLElement;
  const dropzone = target.closest("[role=dropzone]");
  if (!dropzone) {
    dragContext.clearDropzoneEffect();
  }
}

function onDragover(event: DragEvent) {
  event.preventDefault();

  if (event.dataTransfer) {
    event.dataTransfer.dropEffect = "move";
  }
}

function onDrop(event: DragEvent) {
  if (!event.dataTransfer) return;
  const jsonData = event.dataTransfer.getData("application/device-template");
  const data = JSON.parse(jsonData) as TemplateData;
  if (!data.id) return;
  const target = event.target as HTMLElement;
  if (!target) return
  const nid = target.closest("[role=dropzone]")?.getAttribute('data-id');
  bus.trigger(Signal.OpenResourceConfig, {
    tmpl: data,
    node: nodes.value.find(n => n.id === nid),
    callbacks: {
      success: () => {
        nodeStore.$getById(nid!)
      }
    }
  });
}

setTimeout(() => {
  const type = new URLSearchParams(location.search).get('type')
  if (!type) return
  bus.trigger(Signal.OpenResourceConfig, {
    tmpl: { typeCategory: type } as TemplateData,
    node: nodes.value[0]
  });
}, 1000)

function parseChartData(info: ClustNode, key: keyof ClustNode['allocatable']) {
  return info.allocated ? [Math.round((info.allocated[key] || 0) * 100 / (info.allocatable?.[key] || 1))] : [-1]
}

function edit(node: ClustNode) {
  bus.trigger(Signal.OpenNodeResourceList, node);
}

(async () => {
  loading.value = true
  await nodeStore.$fetchList()
  loading.value = false
  nodeStore.$startQuery()
})()

onUnmounted(() => {
  nodeStore.$stopQuery()
})
</script>
<template>
  <VPlaceloadWrap
    v-if="loading"
    class="columns is-multiline pt-2"
  >
    <div class="column is-12 mt-4">
      <VPlaceload
        height="72px"
        width="100%"
      />
    </div>
    <div class="column is-12 mt-4">
      <VPlaceload
        height="72px"
        width="100%"
      />
    </div>
    <div class="column is-12 mt-4">
      <VPlaceload
        height="72px"
        width="100%"
      />
    </div>
  </VPlaceloadWrap>
  <!-- eslint-disable vuejs-accessibility/aria-role -->
  <TransitionGroup
    v-if="!loading"
    name="list"
    tag="div"
    class="container-list"
  >
    <VCard
      v-for="node in nodes"
      :key="node.id"
      role="dropzone"
      radius="rounded"
      :data-id="node.id"
      @dragenter="onDragEnter"
      @dragleave="onDragLeave"
      @dragover="onDragover"
      @drop="onDrop"
    >
      <VIconWrap
        size="large"
        picture="/images/cluster.svg"
      />
      <div class="block-meta">
        <h5>{{ node.id }}</h5>
        <span>IP: {{ node.ip }}</span>
      </div>
      <div class="info-meta">
        <div class="chart-wrap">
          <ApexChart
            :height="cpuOptions.chart.height"
            :type="cpuOptions.chart.type"
            :series="parseChartData(node, 'cpu')"
            :options="cpuOptions"
          />
        </div>
        <div class="chart-wrap">
          <ApexChart
            :height="memoryOptions.chart.height"
            :type="memoryOptions.chart.type"
            :series="parseChartData(node, 'memory')"
            :options="memoryOptions"
          />
        </div>
        <div class="chart-wrap">
          <ApexChart
            :height="nnrOptions.chart.height"
            :type="nnrOptions.chart.type"
            :series="parseChartData(node, 'networkReceiveRate')"
            :options="nnrOptions"
          />
        </div>
        <div class="chart-wrap">
          <ApexChart
            :height="ntrOptions.chart.height"
            :type="ntrOptions.chart.type"
            :series="parseChartData(node, 'networkTransmitRate')"
            :options="ntrOptions"
          />
        </div>
      </div>
      <div class="app-block">
        <div class="app-list running">
          <h6>运行中</h6>
          <div v-for="run in node.running" :key="run" class="app-item">{{ run }}</div>
        </div>
        <div class="app-list stopped">
          <h6>已暂停</h6>
          <div v-for="stp in node.stopped" :key="stp" class="app-item">{{ stp }}</div>
        </div>
      </div>
      <div class="meta-right">
        <div class="buttons">
          <VButton
            color="primary"
            raised
            @keydown.space.prevent="edit(node)"
            @click.prevent="edit(node)"
          >
            编辑
          </VButton>
        </div>
      </div>
    </VCard>
  </TransitionGroup>
  <NodeResourceListModal />
  <NodeEditModal />
</template>
<style lang="scss">
.container-page {

  .columns {
    margin-top: 0;
    margin-inline-start: 0;
    margin-inline-end: 0;
  }
}

.container-list {

  .l-card {
    width: auto;
    padding: 8px 20px;
    display: flex;
    align-items: center;
    margin-top: 16px;
    overflow-x: auto;

    .block-meta {
      margin-left: 16px;
      flex: 0 0 160px;

      h5 {
        font-size: 1.2rem;
        font-weight: 600;
        margin-bottom: 4px;
      }

      span {
        font-size: 0.9rem;
        color: var(--dark-dark-text);
      }
    }

    .icon-wrap.is-large {
      border-radius: 10px;

      img {
        width: 60%;
        border-radius: 0 !important;
      }
    }

    .info-meta {
      flex: 0 0 auto;
      display: flex;
      align-items: center;
      gap: 12px;
      margin-left: 20px;
      margin-right: 20px;
    }

    .chart-wrap {
      width: 90px;
    }

    .app-block {
      flex: 1;
      min-width: 200px;
      display: flex;
      align-items: flex-start;
    }

    .app-list {
      flex: 0 0 90px;
      font-size: 0.75rem;
      padding-left: 12px;
      padding-right: 12px;
      min-height: 90px;

      h6, .app-item {
        color: inherit;
      }

      &.running {
        color: var(--success);
      }

      &.stopped {
        color: var(--danger);
      }
    }

    .meta-right {
      flex: 0 0 auto;
    }

    &.drag-over {
      background-color: var(--fade-grey-light-2);
      color: var(--primary);
      border-color: var(--primary);
    }
  }
}

.is-dark {
  .container-list {
    .l-card.drag-over {
      background-color: var(--dark-sidebar-light-2);
    }
  }
}
</style>
