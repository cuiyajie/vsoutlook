<!-- eslint-disable vue/prop-name-casing -->
<script setup lang="ts">
import { useClustNode } from "@src/stores/node";

const nodeStore = useClustNode()
const nmosNodes = computed(() => nodeStore.nmosNodes)

const mv = defineModel({
  default: '',
  local: true,
})

function onNMosNodeSelect(id: string) {
  const nmos = nmosNodes.value.find(n => n.id === id)
  mv.value = nmos?.href || ''
}

const selectRef = ref<any>(null)

watch(() => mv.value, (nv) => {
  const nmos = nmosNodes.value.find(n => n.href === nv)
  selectRef.value?.select(nmos?.id || '')
}, { immediate: true })

</script>
<template>
  <Multiselect
    ref="selectRef"
    class="nmos-select"
    placeholder="选择NMos节点"
    no-options-text="未查询到NMos节点，请检查RDS服务是否正常"
    value-prop="id"
    label="label"
    :searchable="false"
    :can-deselect="false"
    :can-clear="false"
    :style="{'--ms-max-height': '245px'}"
    :options="nmosNodes"
    @select="onNMosNodeSelect"
  >
    <template #singlelabel="{ value }: { value: NMosNode }">
      <div class="multiselect-single-label">
        <span class="select-label-text">
          <span class="option-item">
            <i class="iconify" data-icon="feather:disc" aria-hidden="true"></i>
            {{ value.label }}
          </span>
          <span class="option-item">
            <i class="iconify" data-icon="feather:server" aria-hidden="true"></i>
            {{ value.hostname }}
          </span>
          <span class="option-item">
            <i class="iconify" data-icon="feather:link-2" aria-hidden="true"></i>
            {{ value.href }}
          </span>
        </span>
      </div>
    </template>
    <template #option="{ option }">
      <span class="select-option-text">
        <span class="option-item">
          <i class="iconify" data-icon="feather:disc" aria-hidden="true"></i>
          {{ option.label }}
        </span>
        <span class="option-item">
          <i class="iconify" data-icon="feather:server" aria-hidden="true"></i>
          {{ option.hostname }}
        </span>
        <span class="option-item">
          <i class="iconify" data-icon="feather:link-2" aria-hidden="true"></i>
          {{ option.href }}
        </span>
      </span>
      <NMosTooltip :nmos="option">
        <i class="iconify" data-icon="feather:alert-circle" aria-hidden="true"></i>
      </NMosTooltip>
    </template>
  </Multiselect>
</template>
<style lang="scss">
.nmos-select {

  .multiselect-option {
    padding-right: 48px;

    span[data-v-tippy] {
      display: flex;
      align-items: center;
    }
  }

  .select-label-text,
  .select-option-text {
    flex: 1;
    display: flex;
    align-items: center;
    gap: 24px;

    .option-item {
      display: flex;
      align-items: center;
      gap: 6px;
    }
  }
}
</style>
