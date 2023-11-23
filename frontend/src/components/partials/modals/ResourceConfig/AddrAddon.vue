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
  <VField addons>
    <VControl>
      <VButton static>
        {{ host }}
      </VButton>
    </VControl>
    <VControl>
      <VInput v-model="port" />
    </VControl>
  </VField>
</template>
