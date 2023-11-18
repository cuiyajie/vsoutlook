<script lang="ts" setup>
import { useTmplDragging } from "/@src/stores/tmplDragging";
const dragContext = useTmplDragging();

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
  bus.trigger(Signal.OpenResourceConfig, {
    tmpl: data,
    container: target.closest("[role=dropzone]")?.getAttribute('data-name'),
  });
}
</script>
<template>
  <!-- eslint-disable vuejs-accessibility/aria-role -->
  <div class="container-list">
    <VCard
      role="dropzone"
      radius="rounded"
      data-id="c1"
      data-name="刀箱1"
      @dragenter="onDragEnter"
      @dragleave="onDragLeave"
      @dragover="onDragover"
      @drop="onDrop"
    >
      <h3>刀箱1</h3>
    </VCard>
    <VCard
      role="dropzone"
      radius="rounded"
      data-id="c2"
      data-name="刀箱2"
      @dragenter="onDragEnter"
      @dragleave="onDragLeave"
      @dragover="onDragover"
      @drop="onDrop"
    >
      <h3>刀箱2</h3>
    </VCard>
    <VCard
      role="dropzone"
      radius="rounded"
      data-id="c3"
      data-name="刀箱3"
      @dragenter="onDragEnter"
      @dragleave="onDragLeave"
      @dragover="onDragover"
      @drop="onDrop"
    >
      <h3>刀箱3</h3>
    </VCard>
    <VCard
      role="dropzone"
      radius="rounded"
      data-id="c4"
      data-name="刀箱4"
      @dragenter="onDragEnter"
      @dragleave="onDragLeave"
      @dragover="onDragover"
      @drop="onDrop"
    >
      <h3>刀箱4</h3>
    </VCard>
  </div>
</template>
<style lang="scss">
.container-list {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(200px, 1fr));
  gap: 16px;
  padding-top: 16px;

  .l-card {
    width: auto;

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
