<script lang="ts" setup>
const addr = defineModel<string>({
  default: "",
  local: true,
});

const prefix = computed(() => {
  return addr.value.split('://')[0] + '://'
})

const host = computed({
  get: () => addr.value.split('://')[1],
  set: (v: string) => {
    addr.value = prefix.value + v
  }
})

defineProps<{
  label: string,
  placeholder: string,
}>()
</script>
<template>
  <VField :label="label" addons>
    <VControl>
      <VButton static>{{ prefix }}</VButton>
    </VControl>
    <VControl expanded>
      <VInput v-model="host" type="text" class="input" :placeholder="placeholder" />
    </VControl>
  </VField>
</template>
