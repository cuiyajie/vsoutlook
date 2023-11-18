<script setup lang="ts">
import { useViewWrapper } from "/@src/stores/viewWrapper";
import { useTemplate } from "/@src/stores/template";

const route = useRoute();
const viewWrapper = useViewWrapper();
viewWrapper.setPageTitle("应用详情");
const tmplStore = useTemplate();

useHead({
  title: `应用详情 - ${__SITE_NAME__}`,
});

const tmpl = computed(() => tmplStore.getById(route.params.id));
</script>

<template>
  <div class="page-content-inner">
    <VBreadcrumb
      with-icons
      separator="arrow"
      :items="[
        {
          label: 'Vuero',
          hideLabel: true,
          icon: 'feather:home',
          to: '/app',
        },
        {
          label: '设备管理',
          to: '/app/template',
        },
        {
          label: tmpl?.name || '新建应用',
          to: `/app/template/${route.params.id}`,
        },
      ]"
    />
    <TemplateDetail />
  </div>
</template>
