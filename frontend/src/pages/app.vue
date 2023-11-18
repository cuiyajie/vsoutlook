<route lang="yaml">
meta:
  requiresAuth: true
</route>
<template>
  <AppLayout :open-on-mounted="true">
    <!-- Content Wrapper -->
    <RouterView v-slot="{ Component }">
      <Transition name="fade-fast" mode="out-in">
        <component :is="Component" />
      </Transition>
    </RouterView>
    <ConfirmModal />
  </AppLayout>
</template>
