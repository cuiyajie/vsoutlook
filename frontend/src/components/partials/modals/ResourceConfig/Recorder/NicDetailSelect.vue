<script lang="ts" setup>
import { type IndexedNicDetail } from '../Utilties/Consts_V1';

const mv = defineModel<number>({
  default: -1,
  local: true,
});

defineProps<{
  nics: IndexedNicDetail[]
}>()

</script>
<template>
  <Multiselect
    v-model="mv"
    placeholder="选择收发流网卡"
    value-prop="index"
    label="label"
    :searchable="false"
    :can-deselect="false"
    :can-clear="false"
    :max-height="145"
    no-options-text="暂时没有配置收发流网卡"
    :options="nics"
  >
    <template #singlelabel="{ value }">
      <div class="multiselect-single-label">
        <span class="select-label-text">
          第 {{ value.index + 1 }} 块收发流网卡
        </span>
      </div>
    </template>
    <template #option="{ option }">
      <span class="select-option-text">
        第 {{ option.index + 1 }} 块收发流网卡
      </span>
    </template>
  </Multiselect>
</template>
