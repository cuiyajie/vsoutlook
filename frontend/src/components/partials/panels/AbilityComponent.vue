<script lang="ts" setup>
import { useAbilComp } from "/@src/stores/abilityComponent";

const acStore = useAbilComp();

const props = defineProps<{
  locked: boolean;
}>();
const filters = ref("");
const abilComps = computed(() => {
  return acStore.abilComps.filter((ac) => {
    return ac.name.toLowerCase().includes(filters.value.toLowerCase());
  });
});

const acGroups = computed(() => {
  const lookup = abilComps.value.reduce((acc, ac) => {
    if (!acc[ac.groupKey]) {
      acc[ac.groupKey] = [];
    }
    acc[ac.groupKey].push(ac);
    return acc;
  }, {} as Record<string, AbilityComp[]>);
  return Object.keys(lookup).map((key) => {
    return {
      key,
      name: lookup[key][0].groupName,
      abilComps: lookup[key],
    };
  });
});

function onDragStart(event: DragEvent, ac: AbilityComp) {
  if (props.locked) return;
  if (event.dataTransfer) {
    event.dataTransfer.setData("application/vueflow", JSON.stringify(ac));
    event.dataTransfer.effectAllowed = "move";
  }
}
</script>
<template>
  <div class="modules-container">
    <VControl icon="feather:search" class="mb-2">
      <input v-model="filters" class="input custom-text-filter" placeholder="搜索..." />
    </VControl>
    <VSingleCollapse v-for="group in acGroups" :key="group.key" :title="group.name">
      <div class="nodes">
        <div
          v-for="ac in group.abilComps"
          :key="ac.id"
          :class="`vue-flow__node-${ac.type}`"
          role="button"
          tabindex="0"
          :draggable="!locked"
          @dragstart="onDragStart($event, ac)"
        >
          {{ ac.name }}
        </div>
      </div>
    </VSingleCollapse>
  </div>
</template>
<style lang="scss">
.modules-container {
  .nodes {
    display: grid;
    grid-template-columns: repeat(auto-fill, minmax(100px, 1fr));
    grid-gap: 8px;
  }
}
</style>
