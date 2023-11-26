<script lang="ts" setup>
import { useTmplDragging } from "/@src/stores/tmplDragging";
import { useClustNode } from "/@src/stores/node";
import { useResourceCharts } from "/@src/composable/useResourceChart";
import ApexChart from 'vue3-apexcharts'

const dragContext = useTmplDragging();
const nodeStore = useClustNode();
const { cpuOptions, memoryOptions } = useResourceCharts()

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
  });
}

function parseCpu(info: ClustNode['info']) {
  return [Math.round((parseInt(info.current?.cpu?.value) || 0) * 100 / (parseInt(info.allocatable?.cpu?.value) || 1)) ]
}

function parseMemory(info: ClustNode['info']) {
  return [Math.round((parseInt(info.current?.memory?.value) || 0) * 100 / (parseInt(info.allocatable?.memory?.value) || 1)) ]
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
      <div class="chart-wrap ml-6">
        <ApexChart
          :height="cpuOptions.chart.height"
          :type="cpuOptions.chart.type"
          :series="parseCpu(node.info)"
          :options="cpuOptions"
        />
      </div>
      <div class="chart-wrap">
        <ApexChart
          :height="memoryOptions.chart.height"
          :type="memoryOptions.chart.type"
          :series="parseMemory(node.info)"
          :options="memoryOptions"
        />
      </div>
    </VCard>
  </TransitionGroup>
</template>
<style lang="scss">
.container-list {
  padding-top: 8px;

  .l-card {
    width: auto;
    padding: 8px 20px;
    display: flex;
    align-items: center;
    margin-top: 16px;

    .block-meta {
      margin-left: 16px;

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

    .chart-wrap {
      width: 200px;
      transform: translateY(-8px);
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
