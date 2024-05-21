<script setup lang="ts">
import { LayoutSize } from "@src/utils/enums";

defineProps<{
  layouts: Layout[],
  currLayout: Layout | null,
}>();

const emit = defineEmits<{
  (e: 'delete', ly: Layout): void,
  (e: 'edit', ly: Layout): void,
  (e: 'publish', ly: Layout): void,
  (e: 'select', ly: Layout): void,
}>();

function selectLayout(ly: Layout) {
  emit('select', ly);
}

</script>

<template>
  <div class="layout-list">
    <div
      v-for="ly in layouts"
      :key="ly.id"
      class="layout-item"
      :class="{'is-active': currLayout?.id === ly.id,'is-publish': ly.published}"
      role="button"
      tabindex="0"
      @click.prevent="selectLayout(ly)"
      @keyup.enter.prevent="selectLayout(ly)"
    >
      <div class="item-name">{{ ly.name }}</div>
      <div class="item-type">{{ ly.size === LayoutSize.HD ? '高清' : '4K' }}</div>
      <VDropdown
        icon="feather:more-vertical"
        spaced
        right
        @click.stop=""
      >
        <template #content="{ close }">
          <a
            role="menuitem"
            href="#"
            class="dropdown-item is-media"
            @click="close(); emit('edit', ly);"
          >
            <div class="meta">
              <span>修改</span>
            </div>
          </a>
          <a
            role="menuitem"
            href="#"
            class="dropdown-item is-media"
            @click="close(); emit('publish', ly);"
          >
            <div class="meta">
              <span>推送</span>
            </div>
          </a>
          <hr class="dropdown-divider">
          <a
            role="menuitem"
            href="#"
            class="dropdown-item is-media"
            @click="close(); emit('delete', ly);"
          >
            <div class="meta">
              <span>删除</span>
            </div>
          </a>
        </template>
      </VDropdown>
    </div>
  </div>
</template>
<style lang="scss">
.layout-list {

  .layout-item {
    display: flex;
    align-items: center;
    padding-top: 0.5rem;
    padding-bottom: 0.5rem;
    padding-left: 0.75rem;
    padding-right: 0.5rem;
    background-color: var(--dark-sidebar-light-12);
    border-radius: 2px;
    margin-top: 4px;
    line-height: 1.35;
    font-size: 13px;
    position: relative;
    margin-left: -1rem;
    margin-right: -1rem;
    transition: all .2s ease-in-out;
    cursor: pointer;

    &.is-active {
      background-color: var(--dark-sidebar-light-4);
      border: 1px solid var(--primary);
    }

    &.is-publish:before {
      content: '';
      position: absolute;
      left: 0;
      top: 0;
      width: 4px;
      bottom: 0;
      background-color: var(--primary);
    }

    .item-name {
      flex: 1;
      color: var(--dark-dark-text);
    }

    .item-type {
      color: var(--blue);
      font-size: 12px;
    }

    .dropdown.is-dots .is-trigger {
      width: 24px;
      height: 24px;
      margin-left: 0.5rem;

      svg {
        width: 16px;
        height: 16px;
        margin-top: -2px;
      }
    }

    .dropdown.is-spaced .dropdown-menu {
      width: 80px;
      min-width: 0;
    }

    .dropdown.is-spaced .dropdown-item {
      font-size: 12px;
      padding: 0.6rem 1rem;
    }

    .dropdown.is-spaced .dropdown-item.is-media .meta {
      margin-inline-start: 0;

      span:first-child {
        font-size: 13px;
        font-weight: normal;
      }
    }

    .dropdown.is-spaced .dropdown-item.is-media .icon {
      width: 16px;
      height: 16px;
    }

    .dropdown-content {
      padding-top: 0.4rem;
      padding-bottom: 0.4rem;
    }

    .dropdown-divider {
      margin: 0;
    }
  }
}
</style>
