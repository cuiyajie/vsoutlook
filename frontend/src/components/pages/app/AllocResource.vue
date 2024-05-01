<script lang="ts" setup>
import { useTemplateType } from "@src/stores/templateType";
import { useTemplate } from "@src/stores/template";

const tmplTypeStore = useTemplateType();
const tmplStore = useTemplate();
const selectedTmpl = ref<TemplateData | null>(null);

tmplTypeStore.$fetchList();
tmplStore.$fetchList();

</script>
<template>
  <div>
    <div class="columns is-multiline resource-page">
      <div class="column is-3">
        <VCard radius="rounded">
          <h3 class="title is-5 mb-4">
            应用列表
          </h3>
          <TmplGroupTabs v-model="selectedTmpl" />
        </VCard>
        <VCard
          radius="rounded"
          class="mt-2"
        >
          <h3 class="title is-5 mb-4">
            应用属性
          </h3>
          <TmplInfo
            v-if="selectedTmpl"
            :tmpl="selectedTmpl"
          />
          <h3
            v-else
            class="dark-inverted"
          >
            未选择应用
          </h3>
        </VCard>
      </div>
      <div class="column is-9 container-page">
        <VCard
          radius="rounded"
        >
          <h3 class="title is-5 mb-0 container-title">
            容器平台
          </h3>
          <ContainerBlocks />
        </VCard>
      </div>
    </div>
    <ResourceConfigModal />
    <DownloadConfirm />
  </div>
</template>
<style lang="scss">
.resource-page {
  .column {
    padding: 0.25rem;
  }

  .tmpl-list {
    .plain-list {
      li > a {
        margin-inline-start: 12px;
        padding: 8px 12px;
        height: 32px;
      }
    }

    li.has-children ul li .is-submenu {
      margin-inline-start: 17px;
      padding: 7px 12px;
    }
  }

  .container-title {
    display: flex;
    justify-content: space-between;
    align-items: flex-start;

    .v-button {
      transform: translateY(-10px);
    }
  }
}

.is-dark .resource-page {
  .column > .l-card {
    background-color: var(--dark-sidebar-light-2);
  }
}
</style>
