<script setup lang="ts">
import type { Component } from 'vue'
import { useI18n } from 'vue-i18n'
import { FocusTrap } from 'focus-trap-vue'

export type VModalSize = 'small' | 'medium' | 'large' | 'big'
export type VModalAction = 'center' | 'right'

export interface VModalEmits {
  (e: 'close'): void
}
export interface VModalProps {
  title: string
  is?: string | Component
  size?: VModalSize
  actions?: VModalAction
  open?: boolean
  rounded?: boolean
  noscroll?: boolean
  noclose?: boolean
  tabs?: boolean
  cancelLabel?: string
  dialogClass?: string
}

defineOptions({
  inheritAttrs: false,
})

const emit = defineEmits<VModalEmits>()
const props = withDefaults(defineProps<VModalProps>(), {
  is: 'div',
  size: undefined,
  actions: undefined,
  cancelLabel: undefined,
  dialogClass: '',
})

const { t } = useI18n()
const wasOpen = ref(false)
const cancelLabel = computed(() => props.cancelLabel || t('cancel-label'))

const checkScroll = () => {
  if (props.noscroll && props.open) {
    if (!import.meta.env.SSR) {
      document.documentElement.classList.add('no-scroll')
    }
    wasOpen.value = true
  } else if (wasOpen.value && props.noscroll && !props.open) {
    if (!import.meta.env.SSR) {
      document.documentElement.classList.remove('no-scroll')
    }
    wasOpen.value = false
  }
}

watchEffect(checkScroll)
onUnmounted(() => {
  if (!import.meta.env.SSR) {
    document.documentElement.classList.remove('no-scroll')
  }
})
</script>

<template>
  <Teleport to="body">
    <FocusTrap v-if="open" :initial-focus="() => ($refs.closeButton as any)?.el">
      <component
        :is="is"
        role="dialog"
        aria-modal="true"
        :class="[open && 'is-active', size && `is-${size}`, dialogClass]"
        class="modal v-modal"
        v-bind="$attrs"
      >
        <div
          class="modal-background v-modal-close"
          tabindex="-1"
          role="button"
          @keydown.space.prevent="() => noclose === false && emit('close')"
          @click="() => noclose === false && emit('close')"
        />
        <div class="modal-content">
          <div class="modal-card">
            <header class="modal-card-head">
              <h3>{{ title }}</h3>
              <slot name="header-action" />
              <button
                ref="closeButton"
                type="button"
                class="v-modal-close ml-auto"
                aria-label="close"
                tabindex="0"
                @keydown.space.prevent="emit('close')"
                @click="emit('close')"
              >
                <i aria-hidden="true" class="iconify" data-icon="feather:x" />
              </button>
            </header>
            <div class="modal-card-body" :class="[props.tabs && 'has-tabs']">
              <div class="inner-content">
                <slot name="content" />
              </div>
            </div>
            <div
              class="modal-card-foot"
              :class="[
                actions === 'center' && 'is-centered',
                actions === 'right' && 'is-end',
              ]"
            >
              <slot name="cancel" :close="() => emit('close')">
                <a
                  tabindex="0"
                  role="button"
                  class="button v-button v-modal-close"
                  :class="[rounded && 'is-rounded']"
                  @keydown.space.prevent="emit('close')"
                  @click="emit('close')"
                >
                  {{ cancelLabel }}
                </a>
              </slot>
              <slot name="action" :close="() => emit('close')" />
            </div>
          </div>
        </div>
      </component>
    </FocusTrap>
  </Teleport>
</template>

<style lang="scss">
.modal {
  transition: all 0.5s;

  &.is-overflow-visible {
    .modal-content,
    .modal-card,
    .modal-card-body {
      overflow: visible;
    }
  }

  &.is-big {
    .modal-content {
      width: 100%;
      max-width: 840px;

      .modal-card {
        width: 100%;
      }
    }
  }

  &.is-large {
    .modal-content {
      width: 100%;
      max-width: 720px;

      .modal-card {
        width: 100%;
      }
    }
  }

  &.is-medium {
    .modal-content {
      width: 100%;
      max-width: 640px;

      .modal-card {
        width: 100%;
      }
    }
  }

  &.is-small {
    .modal-content {
      width: 100%;
      max-width: 420px;

      .modal-card {
        width: 100%;
      }
    }
  }

  .modal-content {
    transition: all 0.4s;
  }
}

.v-modal {
  background: transparent;
  border: transparent;

  &.is-active {
    z-index: 200 !important;

    .v-modal-close {
      cursor: pointer;
    }
  }

  .v-modal-card {
    width: 100%;
    background: var(--white);
    border: 1px solid var(--fade-grey);
    border-radius: 8px;
    padding: 40px;
  }

  &::backdrop {
    background: hsl(var(--dark-sidebar-h) var(--dark-sidebar-s) var(--dark-sidebar-l));
  }

  .modal-content {
    transform: scale(1) !important;
    opacity: 1 !important;
    max-width: 540px;
    overflow-x: hidden;
    animation: fadeInDown 0.5s;
    margin: 0;
    padding: 0 10px;

    .modal-card {
      max-width: 100%;
      margin: 0 auto;

      &.is-rounded {
        border-radius: 12px;
      }

      .modal-card-head {
        display: flex;
        justify-content: space-between;
        align-items: center;
        padding: 15px 20px;
        background: var(--white);
        border-bottom-color: var(--fade-grey-dark-3);

        &.no-border {
          border-bottom-color: transparent;
        }

        h3 {
          font-family: var(--font-alt);
          color: var(--dark-text);
          font-weight: 600;
          font-size: 1rem;
        }

        .v-modal-close {
          display: flex;
          justify-content: center;
          align-items: center;
          background: none;
          border-color: transparent;
          width: 22px;
          height: 22px;
          padding: 0;

          &:hover,
          &:focus {
            svg {
              color: var(--primary);
            }
          }

          &:focus-visible {
            outline-offset: var(--accessibility-focus-outline-offset);
            outline-width: var(--accessibility-focus-outline-width);
            outline-style: var(--accessibility-focus-outline-style);
            outline-color: var(--accessibility-focus-outline-color);
          }

          svg {
            width: 22px;
            height: 22px;
            color: var(--muted-grey);
          }
        }
      }

      .modal-card-body {
        .modal-form {
          padding: 10px 0 20px;
        }

        &.has-tabs {
          padding: 0;
          overflow-x: hidden;

          .tabs {
            overflow-x: auto;

            &::-webkit-scrollbar {
              height: 5px !important;
            }

            &::-webkit-scrollbar-thumb {
              border-radius: 10px !important;
              background: rgb(0 0 0 / 20%) !important;
            }

            a {
              padding: 0.75em 1em;
            }
          }
        }
      }

      .modal-card-foot {
        background-color: var(--white);
        padding: 15px 20px;
        border-top: 1px solid var(--fade-grey-dark-3);

        &.no-border {
          border-top-color: transparent;
        }

        &.is-start {
          justify-content: flex-start !important;
        }

        &.is-centered {
          justify-content: center !important;
        }

        &.is-end {
          justify-content: flex-end !important;
        }

        .v-button {
          min-width: 110px;
        }
      }
    }
  }
}

.is-dark {
  .v-modal {
    .modal-background {
      background: rgb(101 101 104 / 80%) !important;
    }

    .modal-content {
      .modal-card {
        .modal-card-head {
          background: var(--dark-sidebar-light-6) !important;
          border-color: var(--dark-sidebar-light-12);

          h3 {
            color: var(--dark-dark-text);
          }

          .v-modal-close {
            &:hover {
              svg {
                color: var(--primary);
              }
            }
          }
        }

        .modal-card-body {
          background: var(--dark-sidebar-light-6) !important;
        }

        .modal-card-foot {
          background: var(--dark-sidebar-light-6) !important;
          border-color: var(--dark-sidebar-light-12);
        }
      }
    }
  }
}

@media screen and (width >= 769px) {
  .modal.modal-lg {
    .modal-card,
    .modal-content {
      width: 800px !important;
    }
  }

  .modal.modal-sm {
    .modal-card,
    .modal-content {
      width: 400px !important;
    }
  }
}
</style>
