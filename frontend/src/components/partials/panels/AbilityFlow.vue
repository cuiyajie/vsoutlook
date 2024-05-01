<script lang="ts" setup>
import { VueFlow, useVueFlow, Position, type GraphNode, type FlowExportObject } from "@vue-flow/core";
import { Background, BackgroundVariant } from "@vue-flow/background";
import { MiniMap } from "@vue-flow/minimap";
import genId from "@src/utils/id-gen";

const props = defineProps<{
  flow: FlowExportObject | null
}>()

const defaultConfig: Record<string, any> =  {
  input: {
    sourcePosition: Position.Right,
  },
  output: {
    targetPosition: Position.Left,
  },
  default: {
    sourcePosition: Position.Right,
    targetPosition: Position.Left,
  }
}

const { findNode, onConnect, addEdges, setEdges, addNodes, setNodes, setTransform, project, vueFlowRef } = useVueFlow();

onConnect((params) => addEdges(params));

watch(() => props.flow, (nv) => {
  if (nv) {
    const [x = 0, y = 0] = nv.position
    setNodes(nv.nodes)
    setEdges(nv.edges)
    setTransform({ x, y, zoom: nv.zoom || 0 })
  }
}, { immediate: true })

function onDragover(event: DragEvent) {
  event.preventDefault();

  if (event.dataTransfer) {
    event.dataTransfer.dropEffect = "move";
  }
}

function onDrop(event: DragEvent) {
  if (!vueFlowRef.value || !event.dataTransfer) return;
  const jsonData = event.dataTransfer.getData("application/vueflow");
  const { left, top } = vueFlowRef.value.getBoundingClientRect();
  const data = JSON.parse(jsonData)

  const position = project({
    x: event.clientX - left,
    y: event.clientY - top,
  });

  const newNode = {
    id: genId(),
    type: data.type,
    position,
    label: data.name,
    ...defaultConfig[data.type]
  };

  addNodes([newNode]);

  // align node position after drop, so it's centered to the mouse
  nextTick(() => {
    const node = findNode(newNode.id);
    if (!node) return;
    const stop = watch(
      () => node.dimensions,
      (dimensions) => {
        if (dimensions.width > 0 && dimensions.height > 0) {
          node.position = {
            x: node.position.x - node.dimensions.width / 2,
            y: node.position.y - node.dimensions.height / 2,
          };
          stop();
        }
      },
      { deep: true, flush: "post" }
    );
  });
}

function nodeStorColor(node: GraphNode) {
  return {
    'input': 'var(--warning)',
    'output': 'var(--danger)',
    'default': 'var(--primary)',
  }[node.type]
}
</script>
<template>
  <div class="dndflow" @drop="onDrop">
    <VueFlow @dragover="onDragover" class="dark">
      <MiniMap :node-stroke-color="nodeStorColor" node-color="transparent" />
      <Background :variant="BackgroundVariant.Dots" />
    </VueFlow>
  </div>
</template>
<style lang="scss">
.app-wrapper {
  --vf-node-bg: transparent;
  --vf-node-text: var (--light-text);

  .vue-flow__node-default,
  .vue-flow__node-input,
  .vue-flow__node-output {
    padding: 8px 16px;
    width: auto;
  }

  .vue-flow__node-default {
    --vf-node-color: var(--primary);
  }
  .vue-flow__node-input {
    --vf-node-color: var(--warning);
  }
  .vue-flow__node-output {
    --vf-node-color: var(--danger);
  }
}
.dndflow {
  flex-direction: column;
  display: flex;
  height: 100%;
  min-height: inherit;

  .vue-flow {
    min-height: inherit;
  }
}
.dndflow aside {
  color: #fff;
  font-weight: 700;
  border-right: 1px solid #eee;
  padding: 15px 10px;
  font-size: 12px;
  background: rgba(16, 185, 129, 0.75);
  -webkit-box-shadow: 0px 5px 10px 0px rgba(0, 0, 0, 0.3);
  box-shadow: 0 5px 10px #0000004d;
}
.dndflow aside .nodes > * {
  margin-bottom: 10px;
  cursor: grab;
  font-weight: 500;
  -webkit-box-shadow: 5px 5px 10px 2px rgba(0, 0, 0, 0.25);
  box-shadow: 5px 5px 10px 2px #00000040;
}
.dndflow aside .description {
  margin-bottom: 10px;
}
.dndflow .vue-flow-wrapper {
  flex-grow: 1;
  height: 100%;
}
@media screen and (min-width: 640px) {
  .dndflow {
    flex-direction: row;
  }
  .dndflow aside {
    min-width: 25%;
  }
}
@media screen and (max-width: 639px) {
  .dndflow aside .nodes {
    display: flex;
    flex-direction: row;
    gap: 5px;
  }
}
</style>
