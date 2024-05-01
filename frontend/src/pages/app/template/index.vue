<script setup lang="ts">
import { useViewWrapper } from "@src/stores/viewWrapper";

const viewWrapper = useViewWrapper();
viewWrapper.setPageTitle("应用商店");

useHead({
  title: `应用商店 - ${__SITE_NAME__}`,
});
</script>

<template>
  <div class="page-content-inner">
    <TemplateList />
  </div>
</template>
