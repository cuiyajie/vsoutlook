<script setup lang="ts">
import { useViewWrapper } from "@src/stores/viewWrapper";

const viewWrapper = useViewWrapper();
viewWrapper.setPageTitle("末级切换面板");

useHead({
  title: `末级切换面板 - ${__SITE_NAME__}`,
});
</script>

<template>
  <div class="page-content-inner">
    <EndSwitchTmpl />
  </div>
</template>
