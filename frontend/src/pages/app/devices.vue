<script setup lang="ts">
import { useViewWrapper } from "/@src/stores/viewWrapper";

const viewWrapper = useViewWrapper();
viewWrapper.setPageTitle("设备管理");

useHead({
  title: `设备管理 - ${__SITE_NAME__}`,
});
</script>

<template>
  <div class="page-content-inner">
    <DeviceList />
  </div>
</template>
