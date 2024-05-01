<script setup lang="ts">
import { useViewWrapper } from "@src/stores/viewWrapper";

const viewWrapper = useViewWrapper();
viewWrapper.setPageTitle("资源配置");

useHead({
  title: `资源配置 - ${__SITE_NAME__}`,
});
</script>

<template>
  <div class="page-content-inner">
    <AllocResource />
  </div>
</template>
