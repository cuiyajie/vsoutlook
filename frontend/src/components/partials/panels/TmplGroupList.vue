<script setup lang="ts">
import { useTemplateType } from "/@src/stores/templateType";
import { useTemplate } from "/@src/stores/template";

const props = defineProps<{
  filterListed?: boolean;
}>();

const tmplTypeStore = useTemplateType();
const tmplStore = useTemplate();
const open = ref<Record<string, boolean>>({});

const tmplList = computed(() => {
  if (props.filterListed) {
    return tmplStore.tmpls.filter((tmpl) => tmpl.listed);
  } else {
    return tmplStore.tmpls;
  }
});

const tmplGroup = computed(() => {
  return tmplList.value.reduce((acc, tmpl) => {
    if (!acc[tmpl.type]) {
      acc[tmpl.type] = [];
    }
    acc[tmpl.type].push(tmpl);
    return acc;
  }, {} as Record<string, any[]>);
});

const tmplTypes = computed(() => {
  return tmplTypeStore.tmplTypes.filter((type) => {
    return tmplGroup.value[type.id];
  });
});

watch(
  () => tmplTypeStore.tmplTypes,
  () => {
    tmplTypeStore.tmplTypes.forEach((type) => {
      open.value[type.id] = true;
    });
  },
  { immediate: true }
);
</script>

<template>
  <VTabs
    class="tmpl-tabs"
    slider
    type="rounded"
    selected="all"
    align="centered"
    :tabs="[
      { label: '所有', value: 'all' },
      { label: '分组', value: 'group' },
    ]"
  >
    <template #tab="{ activeValue }">
      <div v-if="activeValue === 'all'">
        <div class="tmpl-list">
          <ul>
            <VCollapseLinks
              class="plain-list"
              open
            >
              <slot
                v-for="tmpl in tmplList"
                :key="tmpl.id"
                name="listItem"
                v-bind="{ tmpl }"
              >
                <RouterLink :to="`/app/template/${tmpl.id}`">
                  <span>{{ tmpl.name }}</span>
                </RouterLink>
              </slot>
            </VCollapseLinks>
          </ul>
        </div>
      </div>
      <div v-else-if="activeValue === 'group'">
        <div class="tmpl-list">
          <ul>
            <VCollapseLinks
              v-for="type in tmplTypes"
              :key="type.id"
              v-model:open="open[type.id]"
            >
              <template #header>
                <VIconWrap
                  size="small"
                  :picture="type.icon"
                  class="mr-2 is-tt-icon"
                />
                {{ type.name }}
                <i
                  aria-hidden="true"
                  class="iconify rtl-hidden"
                  data-icon="feather:chevron-right"
                />
                <i
                  aria-hidden="true"
                  class="iconify ltr-hidden"
                  data-icon="feather:chevron-left"
                />
              </template>
              <slot
                v-for="tmpl in tmplGroup[type.id]"
                :key="tmpl.id"
                name="listItem"
                v-bind="{ tmpl, isSubmenu: true }"
              >
                <RouterLink
                  :to="`/app/template/${tmpl.id}`"
                  class="is-submenu"
                >
                  <span>{{ tmpl.name }}</span>
                </RouterLink>
              </slot>
            </VCollapseLinks>
          </ul>
        </div>
      </div>
    </template>
  </VTabs>
</template>
