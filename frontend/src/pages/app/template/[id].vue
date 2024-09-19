<script setup lang="ts">
import { useViewWrapper } from "@src/stores/viewWrapper";
import { useTemplate } from "@src/stores/template";

const route = useRoute();
const viewWrapper = useViewWrapper();
viewWrapper.setPageTitle("应用详情");
const tmplStore = useTemplate();

useHead({
  title: `应用详情 - ${__SITE_NAME__}`,
});

const tmpl = computed(() => tmplStore.getById((route.params as any).id));
const breadcrumbItems = computed(() => [
  {
    label: __SITE_NAME__,
    hideLabel: true,
    icon: "feather:home",
    to: "/app",
  },
  route.query.from === "resource" ? {
    label: "资源管理",
    to: "/app/resource",
  } : {
    label: "应用商店",
    to: "/app/template",
  },
  {
    label: tmpl.value?.name || "新建应用",
    to: `/app/template/${(route.params as any).id}`,
  },
]);
</script>

<template>
  <div class="page-content-inner">
    <VBreadcrumb
      with-icons
      separator="arrow"
      :items="breadcrumbItems"
    />
    <TemplateDetail />
  </div>
</template>
