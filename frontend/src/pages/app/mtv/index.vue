<script setup lang="ts">
import { useViewWrapper } from "@src/stores/viewWrapper";

const viewWrapper = useViewWrapper();
viewWrapper.setPageTitle("多画面布局");

useHead({
  title: `多画面布局 - ${__SITE_NAME__}`,
});
</script>

<template>
  <div class="page-content-inner" style="margin-top: -20px;">
    <MultiTvLayout />
  </div>
</template>
