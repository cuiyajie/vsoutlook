<script setup lang="ts">
import { useViewWrapper } from "@src/stores/viewWrapper";

const viewWrapper = useViewWrapper();
viewWrapper.setPageTitle("新建应用");

useHead({
  title: `新建应用 - ${__SITE_NAME__}`,
});

const siteName = __SITE_NAME__
</script>

<template>
  <div class="page-content-inner">
    <VBreadcrumb
      with-icons
      separator="arrow"
      :items="[
        {
          label: siteName,
          hideLabel: true,
          icon: 'feather:home',
          to: '/app',
        },
        {
          label: '应用商店',
          to: '/app/template',
        },
        {
          label: '未命名应用',
          to: '/app/template/new',
        },
      ]"
    />
    <TemplateDetail />
  </div>
</template>
