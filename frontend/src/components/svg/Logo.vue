<script setup lang="ts">
defineProps<{
  width: string,
  height: string
}>()
</script>

<template>
  <img
    src="@src/assets/logo.svg"
    alt="Visual Outlook"
    :width="width"
    :height="height"
  >
</template>

<style lang="scss" scoped></style>
