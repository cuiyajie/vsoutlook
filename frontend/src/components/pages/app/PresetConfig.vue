<script lang="ts" setup></script>
<template>
  <VTabs
    class="pc-tabs"
    selected="videoFormat"
    align="centered"
    :tabs="[
      { label: '视频格式模板', value: 'videoFormat' },
      { label: '音频格式模板', value: 'audioFormat' },
      { label: '技审模板', value: 'techPreview' },
      { label: '报警规则模板', value: 'auditAlarm' },
      { label: '音频映射模板', value: 'audioMapping' },
    ]"
  >
    <template #tab="{ activeValue }">
      <p v-if="activeValue === 'videoFormat'">
        <VideoFormatList />
      </p>
      <p v-else-if="activeValue === 'audioFormat'">
        <AudioFormatList />
      </p>
      <p v-else-if="activeValue === 'techPreview'">
        <TechReviewList />
      </p>
      <p v-else-if="activeValue === 'auditAlarm'">
        <AuditAlarmList />
      </p>
      <p v-else-if="activeValue === 'audioMapping'">
        <AudioMappingList />
      </p>
    </template>
  </VTabs>
</template>
<style lang="scss">
.pc-tabs {
  > .tabs-inner .tabs li {
    min-width: 160px;
  }
}
</style>
