<script setup lang="ts">
import { useViewWrapper } from "@src/stores/viewWrapper";

const viewWrapper = useViewWrapper();
viewWrapper.setPageTitle("主页");

useHead({
  title: `主页 - ${__SITE_NAME__}`,
});
</script>

<template>
  <div class="page-content-inner">
    <Dashboard />
  </div>
</template>
