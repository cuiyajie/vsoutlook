<script setup lang="ts">
const emit = defineEmits(['close'])
</script>

<template>
  <div class="sidebar-panel is-generic">
    <div class="subpanel-header">
      <h3 class="no-mb">
        主页
      </h3>
      <div
        class="panel-close"
        tabindex="0"
        role="button"
        @keydown.space.prevent="emit('close')"
        @click="emit('close')"
      >
        <i
          aria-hidden="true"
          class="iconify"
          data-icon="feather:x"
        />
      </div>
    </div>
    <div
      class="inner"
      data-simplebar
    >
      <ul>
        <li>
          <RouterLink to="/app/template">
            <i
              aria-hidden="true"
              class="lnil lnil-page pr-2"
            />
            应用商店
          </RouterLink>
        </li>
        <li>
          <RouterLink to="/app/devices">
            <i
              aria-hidden="true"
              class="lnil lnil-cloud-computing pr-2"
            />
            设备管理
          </RouterLink>
        </li>
        <li>
          <RouterLink to="/app/resource">
            <i
              aria-hidden="true"
              class="lnil lnil-database pr-2"
            />
            资源管理
          </RouterLink>
        </li>
        <li>
          <RouterLink to="/app/config">
            <i
              aria-hidden="true"
              class="lnil lnil-cogs pr-2"
            />
            系统配置
          </RouterLink>
        </li>
      </ul>
    </div>
  </div>
</template>

<style lang="scss">
@import '/@src/scss/layout/sidebar-panel';
</style>
