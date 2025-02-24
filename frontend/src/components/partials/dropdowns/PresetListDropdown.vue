<script lang="ts" setup>
const emit = defineEmits<{
  (e: 'add'): void,
  (e: 'remove'): void,
  (e: 'edit'): void,
}>()
</script>
<template>
  <VDropdown
    icon="feather:more-vertical"
    spaced
    right
    class="preset-dropdown"
  >
    <template #content="{ close }">
      <a
        href="#"
        role="menuitem"
        class="dropdown-item is-media"
        @click="close(); emit('add')"
      >
        <div class="icon">
          <i
            class="iconify"
            data-icon="feather:plus"
            aria-hidden="true"
          />
        </div>
        <div class="meta">
          <span>添加</span>
        </div>
      </a>
      <a
        href="#"
        role="menuitem"
        class="dropdown-item is-media"
        @click="close(); emit('edit')"
      >
        <div class="icon">
          <i
            class="iconify"
            data-icon="feather:edit"
            aria-hidden="true"
          />
        </div>
        <div class="meta">
          <span>编辑</span>
        </div>
      </a>

      <hr class="dropdown-divider">

      <a
        href="#"
        role="menuitem"
        class="dropdown-item is-media is-danger"
        @click="close(); emit('remove')"
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
.dropdown.preset-dropdown {
  .dropdown-menu {
    min-width: 120px;
  }
}

.is-dark {
  .dropdown.preset-dropdown {
    .dropdown-item.is-media {

      &.is-danger {
        .icon svg,
        &:hover .icon svg,
        .meta span:first-child {
          color: var(--danger) !important;
        }
      }

      &:hover .meta span:first-child {
        color: var(--primary);
      }
    }
  }
}
</style>
