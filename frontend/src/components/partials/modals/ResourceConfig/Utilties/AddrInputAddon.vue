<script lang="ts" setup>
const addr = defineModel<string>({
  default: "",
  local: true,
});

const host = computed({
  get: () => addr.value.split(":")[0],
  set: (value: string) => {
    addr.value = `${value}:${port.value}`
  }
})

const port = computed({
  get: () => parseInt(addr.value.split(":")[1]),
  set: (value: number) => {
    addr.value = `${host.value}:${value}`
  }
})

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
