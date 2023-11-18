<script setup lang="ts">
const opened = ref(false);
const draftSys = ref({} as SystemInfo);

useListener(Signal.OpenNewSystem, () => {
  opened.value = true;
});
</script>
<template>
  <VModal
    is="form"
    method="post"
    novalidate
    size="small"
    :open="opened"
    title="新建系统"
    actions="right"
    cancel-label="取消"
    @submit.prevent="opened = false"
    @close="opened = false"
  >
    <template #content>
      <div class="modal-form">
        <VField>
          <VControl>
            <VInput
              v-model="draftSys.name"
              type="text"
              placeholder="系统名称"
            />
          </VControl>
        </VField>
      </div>
    </template>
    <template #action>
      <VButton
        type="submit"
        color="primary"
        raised
        @click=""
      >
        保存
      </VButton>
    </template>
  </VModal>
</template>
