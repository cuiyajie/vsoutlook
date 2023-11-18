<script lang="ts" setup></script>
<template>
  <VTabs
    class="gc-tabs"
    selected="users"
    align="centered"
    :tabs="[
      { label: '通用', value: 'general' },
      { label: '用户', value: 'users' }
    ]"
  >
    <template #tab="{ activeValue }">
      <p v-if="activeValue === 'general'">
        <GeneralConfig />
      </p>
      <p v-else-if="activeValue === 'users'">
        <UserList />
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
