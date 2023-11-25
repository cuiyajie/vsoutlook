<script setup lang="ts">
import { ref } from "vue";
import type { ConfirmDialogOptions } from "/@src/utils/dialog";
const opened = ref(false);

const defaultOptions: ConfirmDialogOptions = {
  title: "提醒",
  cancelText: "取消",
  confirmText: "确认",
  onConfirm: (hide: () => void) => {},
};
const dopts = ref<ConfirmDialogOptions>(defaultOptions);

useListener(Signal.OpenConfirmDialog, (payload: any) => {
  dopts.value = { ...defaultOptions, ...payload };
  opened.value = true;
});

function handleConfirm() {
  dopts.value.onConfirm?.(() => {
    opened.value = false;
  });
}
</script>

<template>
  <VModal
    :open="opened"
    actions="center"
    size="small"
    :title="dopts.title"
    :cancel-label="dopts.cancelText"
    dialog-class="confirm-dialog"
    @close="opened = false"
  >
    <template #content>
      <div
        v-if="dopts.html"
        v-html="dopts.html"
      />
      <h3 v-else>
        {{ dopts.content }}
      </h3>
    </template>
    <template #action>
      <VButton
        color="primary"
        raised
        @click="handleConfirm"
      >
        {{ dopts.confirmText }}
      </VButton>
    </template>
  </VModal>
</template>
<style lang="scss">
.modal.v-modal {
  &.confirm-dialog {
    z-index: 1000 !important;
  }
}
</style>
