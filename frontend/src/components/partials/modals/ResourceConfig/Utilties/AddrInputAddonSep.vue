<script lang="ts" setup>
const host = defineModel<string>('host', {
  default: "",
  local: true,
});

const port = defineModel<number>('port', {
  default: 30000,
  local: true,
});

defineProps<{
  label: string
}>()

</script>
<template>
  <VField :label="label" addons class="api-addons is-input-number">
    <VControl expanded>
      <VInput v-model="host" type="text" class="input" placeholder="接口名称" />
    </VControl>
    <VControl>
      <VButton class="addon-conn" static>:</VButton>
    </VControl>
    <VControl>
      <VInputNumber
        v-model="port"
        centered
        :min="0"
        :step="1"
      />
    </VControl>
  </VField>
</template>
