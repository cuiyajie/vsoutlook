<script lang="ts" setup>
const port = defineModel<string>({
  default: "",
  local: true,
});

defineProps<{
  host: string
}>()
</script>
<template>
  <VField addons class="ip-addr">
    <VControl>
      <VButton static>
        {{ host }}
      </VButton>
    </VControl>
    <VControl expanded>
      <VInput v-model="port" />
    </VControl>
  </VField>
</template>
