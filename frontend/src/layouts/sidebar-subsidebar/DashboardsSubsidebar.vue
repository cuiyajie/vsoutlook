<script setup lang="ts">
// const emit = defineEmits(['close'])
</script>

<template>
  <div class="sidebar-panel is-generic">
    <!-- <div class="subpanel-header">
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
    </div> -->
    <div
      class="inner"
      data-simplebar
    >
      <ul>
        <li>
          <RouterLink to="/app/template">
            <i
              class="mr-1 fas fa-shopping-bag"
              aria-hidden="true"
            />
            应用商店
          </RouterLink>
        </li>
        <li>
          <RouterLink to="/app/devices">
            <i
              class="mr-1 fas fa-layer-group"
              aria-hidden="true"
            />
            设备管理
          </RouterLink>
        </li>
        <li>
          <RouterLink to="/app/resource">
            <i
              class="mr-1 fas fa-server"
              aria-hidden="true"
            />
            资源管理
          </RouterLink>
        </li>
        <li>
          <RouterLink to="/app/config">
            <i
              class="mr-1 fas fa-users-cog"
              aria-hidden="true"
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
