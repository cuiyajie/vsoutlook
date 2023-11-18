<script lang="ts" setup>
</script>
<template>
  <VDropdown
    icon="feather:more-vertical"
    spaced
    right
    class="sys-list-dropdown"
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
.dropdown.is-spaced.sys-list-dropdown {
  .dropdown-menu {
    min-width: 110px;
  }

  .dropdown-item {
    padding: 0 0.5rem;
  }
}
</style>
