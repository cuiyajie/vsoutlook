<script lang="ts" setup>

const props = defineProps<{
  modelValue: LayoutDataItem,
  bound: LayoutDimension,
  base: LayoutDimension
}>()

const emit = defineEmits<{
  (e: 'reset'): void,
  (e: 'update:model-value', value: LayoutDataItem): void
  (e: 'toggle-show-date', value: boolean): void
}>()

const data = computed({
  get: () => props.modelValue,
  set: (v: LayoutDataItem) => emit('update:model-value', v)
})

const isTimer = computed(() => Boolean(props.modelValue.timer))
const isText = computed(() => Boolean(props.modelValue.text))
</script>
<template>
  <div class="layout-panel is-sticky">
    <div class="layout-header mb-4">
      <h3 class="title is-5">
        属性设置
      </h3>
      <a
        role="button"
        class="action-link"
        tabindex="0"
        @keydown.space.prevent="emit('reset')"
        @click.prevent="emit('reset')"
      >重置</a>
    </div>
    <div class="layout-form">
      <timer-setting v-if="isTimer" v-model="data" :bound="props.bound" :base="props.base" @toggle-show-date="emit('toggle-show-date', $event)" />
      <text-setting v-else-if="isText" v-model="data" :bound="props.bound" :base="props.base" />
      <win-setting v-else v-model="data" :bound="props.bound" :base="props.base" />
    </div>
  </div>
</template>
<style lang="scss">
.layout-panel {

  &.is-sticky {

    .layout-header {
      position: sticky;
      top: 0;
      background: var(--dark-sidebar-light-6);
      padding-top: 20px;
      padding-bottom: 20px;
      margin-bottom: 0 !important;
      transition: none;
      z-index: 10;
    }
  }

  .layout-header {
    display: flex;
    justify-content: space-between;
    padding-left: 20px;
    padding-right: 20px;

    h3.title {
      margin-bottom: 0;
    }
  }

  .layout-form {
    padding: 0 20px 20px;
    font-size: 12px;
    color: var(--dark-dark-text);

    input {
      font-size: inherit;
    }

    &.half {
      label {
        width: 10px;
      }
    }

    div {
      color: inherit;
    }

    .layout-row-inner {
      display: flex;
      align-items: center;
      gap: 16px;
      padding: 8px 0;
      flex-wrap: wrap;
    }

    .layout-cell {
      display: flex;
      align-items: center;
      gap: 8px;

      input[type='number'] {
        -moz-appearance: textfield;
      }

      input[type='number']::-webkit-outer-spin-button,
      input[type='number']::-webkit-inner-spin-button {
        -webkit-appearance: none;
      }

      input[type='text'],
      input[type='number'] {
        background-color: inherit;
        border: none;
        border-bottom: 1px solid #fff;
        line-height: 140%;
        border-radius: 0;
        width: 48px;

        &:focus-visible {
          outline: none;
        }
      }

      &.fullwidth {
        width: 100%;

        input[type='text'] {
          flex: 1;
        }
      }

      .checkbox {
        padding: 0;
        margin-inline-start: 16px;
      }

      .checkbox input + span {
        top: 0;
      }

      .radio {
        padding: 0.2rem 0;
        font-size: 0.85rem;

        input + span {
          width: 1rem;
          height: 1rem;

          &:after {
            top: 50%;
          }
        }
      }
    }
  }

  .layout-row {

    & + .layout-row {
      margin-top: 12px;
    }
  }

  .layout-section {

    .section-header {
      font-weight: 500;
      padding: 16px 0;
      margin-inline-start: -8px;
    }
  }
}

.vc-display .vc-color-input input,
.vc-display .vc-alpha-input__inner {
  color: #fff !important;
}

.vc-display .vc-input-toggle:hover {
  background-color: transparent !important;
}
</style>
