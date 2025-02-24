<script lang="ts" setup></script>
<template>
  <VTabs
    class="gc-tabs"
    selected="general"
    align="centered"
    :tabs="[
      { label: '通用', value: 'general' },
      { label: '用户', value: 'users' },
      { label: '备份', value: 'backup' },
      { label: '预置', value: 'preset' },
    ]"
  >
    <template #tab="{ activeValue }">
      <p v-if="activeValue === 'general'">
        <GeneralConfig />
      </p>
      <p v-else-if="activeValue === 'users'">
        <UserList />
      </p>
      <p v-else-if="activeValue === 'backup'">
        <BackupBlock />
      </p>
      <p v-else-if="activeValue === 'preset'">
        <PresetConfig />
      </p>
    </template>
  </VTabs>
</template>
<style lang="scss">
.gc-tabs {
  .tabs li {
    min-width: 120px;
  }
}
</style>
