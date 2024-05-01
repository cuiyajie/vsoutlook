<script lang="ts" setup>
import { useSystems } from "@src/stores/system";

const sysStore = useSystems()

const createNewSystem = () => {
  bus.trigger(Signal.OpenNewSystem)
};

</script>
<template>
  <div class="sys-list card-widget">
    <div class="title-wrap">
      <h3>
        系统列表
      </h3>
      <button
        class="button is-circle is-dark-outlined"
        @click="createNewSystem"
      >
        <span class="icon is-small">
          <i
            aria-hidden="true"
            class="iconify"
            data-icon="feather:plus"
          />
        </span>
      </button>
    </div>
    <ul class="mt-4">
      <li
        v-for="sys in sysStore.systems"
        :key="sys.id"
      >
        <span class="sys-title">
          {{ sys.name }}
        </span>
        <SystemDropdown />
      </li>
    </ul>
  </div>
</template>
<style lang="scss">
.sys-list {

  li {
    height: 36px;
    display: flex;
    justify-content: flex-start;
    align-items: center;
    border-inline-start: 2px solid transparent;
    cursor: pointer;

    &.is-active {
      .sys-title {
        font-weight: 600;
        color: var(--primary);
      }
    }

    .sys-title {
      font-family: var(--font);
      display: block;
      width: 100%;
      font-size: 0.9rem;
      font-weight: 500;
      color: var(--light-text);

      &:hover,
      &:focus {
        color: var(--dark-text);
      }
    }

    .dropdown.is-dots {
      display: none;
    }

    &:hover {
      .dropdown.is-dots {
        display: inline-flex;
      }
    }

    .dropdown.is-dots .is-trigger {
      width: 20px;
      height: 20px;

      > span {
        height: 20px;
        display: flex;
        align-items: center;
      }

      svg {
        width: 14px;
        height: 14px;
      }
    }
  }
}
</style>
