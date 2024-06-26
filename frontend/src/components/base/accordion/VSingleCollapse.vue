<script setup lang="ts">
export interface VSingleCollapseProps {
  title: string;
  withChevron?: boolean;
}

withDefaults(defineProps<VSingleCollapseProps>(), {
  title: "",
  withChevron: true,
});

const opened = ref(true);
const toggle = () => {
  opened.value = !opened.value;
};
</script>

<template>
  <details
    :class="[withChevron && 'has-chevron', !withChevron && 'has-plus']"
    :open="opened"
    class="collapse"
  >
    <summary
      class="collapse-header"
      tabindex="0"
      role="button"
      @keydown.space.prevent="toggle"
      @click.prevent="toggle"
    >
      <h3>{{ title }}</h3>
      <div class="collapse-head-info">
        <div class="collapse-icon">
          <VIcon v-if="withChevron" icon="feather:chevron-down" />
          <VIcon v-else-if="!withChevron" icon="feather:plus" />
        </div>
      </div>
    </summary>
    <div class="collapse-content">
      <slot />
    </div>
  </details>
</template>

<style lang="scss">
@import "@src/scss/abstracts/all";

.collapse {
  @include vuero-s-card;

  padding: 0;
  margin-bottom: 1.5rem;

  &.has-plus {
    &[open] {
      .collapse-header {
        .collapse-icon {
          transform: rotate(calc(var(--transform-direction) * 45deg));
        }
      }

      .collapse-content {
        display: block;
      }
    }
  }

  &.has-chevron {
    &[open] {
      .collapse-header {
        .collapse-icon {
          transform: rotate(calc(var(--transform-direction) * 180deg));
        }
      }

      .collapse-content {
        display: block;
      }
    }
  }

  &[open] {
    .collapse-icon {
      border-color: var(--fade-grey-dark-3) !important;
      box-shadow: var(--light-box-shadow);
    }
  }

  .collapse-header {
    position: relative;
    display: flex;
    align-items: center;
    justify-content: space-between;
    height: 60px;
    padding: 0 20px;
    cursor: pointer;

    h3 {
      font-family: var(--font-alt);
      font-size: 0.9rem;
      font-weight: 600;
      color: var(--dark-text);
    }

    .collapse-head-info {
      display: flex;
      align-items: center;
      justify-content: flex-end;
    }

    .collapse-icon {
      display: flex;
      justify-content: center;
      align-items: center;
      height: 30px;
      width: 30px;
      background: var(--white);
      border-radius: var(--radius-rounded);
      border: 1px solid transparent;
      transition: all 0.3s; // transition-all test

      > span {
        display: block;
        height: 16px;
        width: 16px;
      }

      svg {
        height: 16px;
        width: 16px;
        color: var(--light-text);
      }
    }
  }

  .collapse-content {
    display: none;
    padding: 0 20px 20px;
    color: var(--light-text);
    font-family: var(--font);

    p:not(:last-child) {
      margin-bottom: 12px;
    }
  }
}

.is-dark {
  .collapse {
    @include vuero-card--dark;

    &[open] {
      .collapse-header {
        .collapse-icon {
          background: var(--dark-sidebar-light-2);
          border-color: var(--dark-sidebar-light-4) !important;
        }
      }
    }

    .collapse-header {
      h3 {
        color: var(--dark-dark-text);
      }

      .collapse-icon {
        background: var(--dark-sidebar-light-6);
        border-color: var(--dark-sidebar-light-6);
      }
    }
  }
}
</style>
