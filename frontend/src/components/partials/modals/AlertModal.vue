<script setup lang="ts">
const opened = ref(false)
const info = ref<any>({})
useListener(Signal.OpenAlertDialog, (payload: any) => {
  info.value = payload
  opened.value = true;
});
</script>

<template>
  <VModal
    :title="info.title || '提醒'"
    :open="opened"
    actions="right"
    size="small"
    :cancel-label="info.cancelText || '关闭'"
    @close="opened = false"
  >
    <template #content>
      <VPlaceholderSection
        :title="info.content"
      />
    </template>
  </VModal>
</template>
<style lang="scss">
.modal.v-modal {
  &.confirm-dialog {
    z-index: 1500 !important;
  }
}
</style>
