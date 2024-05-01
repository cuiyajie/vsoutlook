<script setup lang="ts">
import { useViewWrapper } from '@src/stores/viewWrapper'

const viewWrapper = useViewWrapper()
viewWrapper.setPageTitle('系统配置')

useHead({
  title: `系统配置 - ${__SITE_NAME__}`,
})
</script>

<template>
  <div class="page-content-inner">
    <Configuration />
  </div>
</template>
