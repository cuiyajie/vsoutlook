<script lang="ts" setup>
import { DeviceStatus } from '/@src/utils/enums'
defineProps<{
  device: DeviceVerbose
}>()
</script>
<template>
  <VDropdown
    icon="feather:more-vertical"
    spaced
    right
    class="device-list-dropdown"
  >
    <template #content="{ close }">
      <a
        href="#"
        role="menuitem"
        class="dropdown-item is-media"
        @click="close()"
      >
        <div class="icon">
          <i
            class="iconify"
            data-icon="feather:cpu"
            aria-hidden="true"
          />
        </div>
        <div class="meta">
          <span>配置</span>
        </div>
      </a>

      <hr class="dropdown-divider">

      <a
        href="#"
        role="menuitem"
        class="dropdown-item is-media"
        @click="close()"
      >
        <div class="icon">
          <i
            class="iconify"
            :data-icon="`feather:${device.status === DeviceStatus.Normal ? 'stop-circle' : 'arrow-right-circle'}`"
            aria-hidden="true"
          />
        </div>
        <div class="meta">
          <span>{{ device.status === DeviceStatus.Normal ? '关闭' : '开启' }}</span>
        </div>
      </a>

      <hr class="dropdown-divider">

      <a
        href="#"
        role="menuitem"
        class="dropdown-item is-media"
        @click="close()"
      >
        <div class="icon">
          <i
            class="iconify"
            data-icon="feather:refresh-cw"
            aria-hidden="true"
          />
        </div>
        <div class="meta">
          <span>刷新</span>
        </div>
      </a>

      <hr class="dropdown-divider">

      <a
        href="#"
        role="menuitem"
        class="dropdown-item is-media"
        @click="close()"
      >
        <div class="icon">
          <i
            class="iconify"
            data-icon="feather:trash-2"
            aria-hidden="true"
          />
        </div>
        <div class="meta">
          <span>删除</span>
        </div>
      </a>
    </template>
  </VDropdown>
</template>
<style lang="scss">
.dropdown.is-spaced.device-list-dropdown {
  .dropdown-menu {
    min-width: 110px;
  }
}
</style>
