<script setup lang="ts">
import { useTemplateType } from "@src/stores/templateType";
import { useTemplate } from "@src/stores/template";
import { useTmplDragging } from "@src/stores/tmplDragging";
import { MirrorType } from "@src/utils/enums";

const tmplTypeStore = useTemplateType();
const tmplStore = useTemplate();
const dragContext = useTmplDragging();

const modelValue = defineModel<TemplateData | null>({
  default: null,
});

const groups = computed(() => {
  return tmplStore.tmpls.reduce(
    ([cg, vg], tmpl) => {
      if (tmpl.mirrorType === MirrorType.Container) {
        if (!cg[tmpl.type]) {
          cg[tmpl.type] = [];
        }
        cg[tmpl.type].push(tmpl);
      } else if (tmpl.mirrorType === MirrorType.Virtual) {
        if (!vg[tmpl.type]) {
          vg[tmpl.type] = [];
        }
        vg[tmpl.type].push(tmpl);
      }
      return [cg, vg];
    },
    [{} as Record<string, any[]>, {} as Record<string, any[]>]
  );
});

const containerTypes = ref<TemplateType[]>([]);
const virtualTypes = ref<TemplateType[]>([]);
const containerOpen = ref<Record<string, boolean>>({});
const virtualOpen = ref<Record<string, boolean>>({});

watchEffect(() => {
  containerTypes.value = Object.keys(groups.value[0]).map(
    (id) => tmplTypeStore.getById(id)!
  );
  virtualTypes.value = Object.keys(groups.value[1]).map(
    (id) => tmplTypeStore.getById(id)!
  );

  containerOpen.value = containerTypes.value.reduce((acc, type) => {
    acc[type.id] = true;
    return acc;
  }, {} as Record<string, boolean>);
  virtualOpen.value = virtualTypes.value.reduce((acc, type) => {
    acc[type.id] = true;
    return acc;
  }, {} as Record<string, boolean>);
});

function onDragStart(event: DragEvent, tmpl: TemplateData) {
  if (event.dataTransfer) {
    event.dataTransfer.setData("application/device-template", JSON.stringify(tmpl));
    event.dataTransfer.effectAllowed = "move";
  }
}

function onDragEnd() {
  dragContext.clearDropzoneEffect();
}
</script>

<template>
  <TmplGroupList filter-listed>
    <template #listItem="{ tmpl, isSubmenu }">
      <a
        :class="[{ 'is-active': tmpl.id === modelValue?.id }, isSubmenu && 'is-submenu']"
        draggable="true"
        role="menuitem"
        tabindex="0"
        @click="modelValue = tmpl"
        @keydown.enter="modelValue = tmpl"
        @dragstart="onDragStart($event, tmpl)"
        @dragend="onDragEnd"
      >
        <span>{{ tmpl.name }}</span>
      </a>
    </template>
  </TmplGroupList>
</template>
<style lang="scss">
.tmpl-tabs {
  &.tabs-wrapper.is-slider {
    > .tabs-inner > .tabs {
      width: 100%;
      max-width: none;

      li a {
        text-align: center;
        line-height: 27px;
      }
    }
  }
}
</style>
