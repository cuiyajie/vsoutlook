<script setup lang="ts">
import VField from "../../base/form/VField.vue";
import { useUserSession } from "/@src/stores/userSession";

const usStore = useUserSession();
const settings = computed(() => usStore.settings);

const def_auth_service = {
	ip: "232.0.0.0",
  port: 6001
}

const authServices = ref<{ ip: string, port: number }[]>([]);

function parseAuthService() {
  let services = []
  try {
    services = JSON.parse(settings.value.authorization_services || "")
  } catch (error) {
    services = []
  }
  authServices.value = services
}

watch(() => settings.value.authorization_services, () => {
  parseAuthService()
}, { immediate: true })

function updateProp(key: string, value?: string) {
  if (key === 'authorization_services') {
    usStore.$updateSettings({ key, value: JSON.stringify(authServices.value) })
  } else {
    usStore.$updateSettings({ key, value })
  }
}

const handleKeyDown = (event: KeyboardEvent) => {
  if (event.key === 'Enter') {
    (event.target as HTMLInputElement).blur();
  }
}

const vBlurOnEnter = {
  mounted(el: HTMLInputElement) {

    const handleBlur = (event: FocusEvent) => {
      const target = event.target as HTMLInputElement;
      const key = target.dataset.key;
      if (key) {
        updateProp(key, target.value)
      }
    }

    el.addEventListener('keydown', handleKeyDown);
    el.addEventListener('blur', handleBlur);

    (el as any)._blurOnEnterDirectiveCleanup = () => {
      el.removeEventListener('keydown', handleKeyDown);
      el.removeEventListener('blur', handleBlur);
    };
  },
  beforeUnmount(el: HTMLInputElement) {
    if ((el as any)._blurOnEnterDirectiveCleanup) {
      (el as any)._blurOnEnterDirectiveCleanup();
      delete (el as any)._blurOnEnterDirectiveCleanup;
    }
  }
};

function addAuthService() {
  authServices.value.push({ ...def_auth_service })
  updateProp('authorization_services')
}

function removeAuthService(idx: number) {
  authServices.value.splice(idx, 1)
  updateProp('authorization_services')
}

</script>

<template>
  <form
    method="post"
    novalidate
    class="form-layout is-stacked gconfig"
  >
    <div class="form-outer">
      <div class="form-body">
        <div class="form-section is-grey">
          <div class="form-section-header">
            <div class="left">
              <h3>NMOS服务</h3>
            </div>
          </div>

          <div class="form-section-inner is-horizontal flex-auto">
            <VField
              horizontal
              label="RDS服务地址(含端口)"
            >
              <VControl
                icon="feather:map-pin"
                fullwidth
              >
                <VInput
                  v-model="settings.rds_server_url"
                  v-blur-on-enter
                  type="url"
                  placeholder="192.168.1.112:8010"
                  inputmode="url"
                  data-key="rds_server_url"
                />
              </VControl>
            </VField>
          </div>
        </div>

        <div class="form-section is-grey">
          <div class="form-section-header">
            <div class="left">
              <h3>授权服务配置</h3>
            </div>
            <button
              type="button"
              class="button is-circle is-dark-outlined"
              @keydown.space.prevent="addAuthService"
              @click.prevent="addAuthService"
            >
              <span class="icon is-large">
                <i aria-hidden="true" class="iconify" data-icon="feather:plus" />
              </span>
            </button>
          </div>
          <div v-for="(auth, idx) in authServices" :key="idx" class="columns is-multiline">
            <div class="column is-1">
              <VField class="index-field">
                <VLabel>序号</VLabel>
                <VControl>
                  <VLabel>{{ idx }}</VLabel>
                </VControl>
              </VField>
            </div>
            <div class="column is-6">
              <VField id="ip">
                <VLabel>IP地址</VLabel>
                <VControl icon="feather:map-pin">
                  <input
                    v-model="auth.ip"
                    v-blur-on-enter
                    type="text"
                    :placeholder="def_auth_service.ip"
                    class="input is-primary-focus"
                    data-key="authorization_services"
                  >
                </VControl>
              </VField>
            </div>
            <div class="column is-4">
              <VField>
                <VLabel>端口</VLabel>
                <VControl>
                  <VInputNumber
                    v-model="auth.port"
                    :placeholder="def_auth_service.port"
                    centered
                    :min="0"
                    :step="1"
                    data-key="authorization_services"
                    @blur="updateProp('authorization_services')"
                    @step-click="val => $nextTick(() => updateProp('authorization_services'))"
                    @keyup.enter.stop="handleKeyDown"
                  />
                </VControl>
              </VField>
            </div>
            <div class="column is-1">
              <VField>
                <VLabel>&nbsp;</VLabel>
                <VControl>
                  <VIconButton color="warning" light raised circle icon="feather:x" @click="removeAuthService(idx)" />
                </VControl>
              </VField>
            </div>
          </div>
          <div v-if="authServices.length === 0" class="form-empty">暂时没有添加授权服务</div>
        </div>
      </div>
    </div>
  </form>
</template>

<style lang="scss">
@import "/@src/scss/abstracts/all";
@import "/@src/scss/components/forms-outer";

.is-navbar {
  .form-layout {
    margin-top: 30px;
  }
}

.form-layout {
  max-width: 740px;
  margin: 0 auto;

  &.gconfig {

    .form-outer {
      border-radius: 16px;
      overflow: hidden;
    }

    .form-empty {
      padding: 20px;
      text-align: center;
      font-style: italic;
      font-weight: 500;
      color: var(--light-text);
    }

    .index-field {

      .control > label {
        margin-inline-start: 0.54rem;
        color: white;
        line-height: 38px;
        color: var(--light-text);
      }
    }
  }

  &.is-stacked {
    .form-outer {
      .form-body {
        padding: 0;

        .form-section {
          padding: 40px;
          border-bottom: 1px solid var(--fade-grey-dark-4);

          &.is-grey {
            background: #fafafa;
          }

          .form-section-header {
            display: flex;
            align-items: center;
            justify-content: space-between;
            border-bottom: 1px solid var(--fade-grey-dark-4);
            padding-bottom: 20px;
            margin-bottom: 30px;

            .left {
              h3 {
                font-family: var(--font-alt);
                font-weight: 600;
                color: var(--dark-text);
              }
            }
          }

          .form-section-inner {
            &.is-horizontal {
              max-width: 540px;
            }

            &.flex-auto {

              .field-label {
                flex-basis: auto;
                flex-shrink: 0;
                flex-grow: 0;
              }
            }

            .field {
              &.is-horizontal {
                .field-label {
                  padding-top: 0.75em;
                }
              }
            }
          }

          .columns {
            .column {
              padding-top: 0.5rem;
              padding-bottom: 0.5rem;
            }
          }

          .field {
            .control {
              .checkbox {
                padding: 0;
                padding-inline-end: 10px;
                font-size: 0.9rem;
              }
            }
          }

          .participants {
            display: flex;
            padding-bottom: 10px;

            .v-avatar {
              margin-inline-end: 8px;
            }

            .add-participant {
              display: flex;
              justify-content: center;
              align-items: center;
              height: 40px;
              width: 40px;
              min-width: 40px;
              border-radius: var(--radius-rounded);
              border: 1.6px dashed var(--light-text);
              color: var(--light-text);
              padding: 0;
              background: none;
              margin-inline-start: 4px;
              cursor: pointer;
              transition: color 0.3s, background-color 0.3s, border-color 0.3s,
                height 0.3s, width 0.3s;

              &:hover,
              &:focus {
                border: 1.6px solid var(--primary);
                color: var(--primary);
              }

              &:focus-visible {
                outline-offset: var(--accessibility-focus-outline-offset);
                outline-width: var(--accessibility-focus-outline-width);
                outline-style: var(--accessibility-focus-outline-style);
                outline-color: var(--accessibility-focus-outline-color);
              }

              svg {
                height: 16px;
                width: 16px;
              }
            }
          }

          .color-codes {
            display: flex;
            align-items: center;
            height: 38px;

            .color-code {
              height: 14px;
              width: 14px;
              border-radius: var(--radius-rounded);
              background: var(--white);
              margin-inline-end: 10px;
              border: 3px solid var(--light-text);
              cursor: pointer;
              opacity: 0.6;
              transition: color 0.3s, background-color 0.3s, border-color 0.3s,
                height 0.3s, width 0.3s;

              &:hover,
              &:focus {
                opacity: 1;
              }

              &:focus-visible {
                outline-offset: var(--accessibility-focus-outline-offset);
                outline-width: var(--accessibility-focus-outline-width);
                outline-style: var(--accessibility-focus-outline-style);
                outline-color: var(--accessibility-focus-outline-color);
              }

              &.is-primary {
                border-color: var(--primary);

                &.is-active {
                  background: var(--primary);
                }
              }

              &.is-secondary {
                border-color: var(--secondary);

                &.is-active {
                  background: var(--secondary);
                }
              }

              &.is-info {
                border-color: var(--info);

                &.is-active {
                  background: var(--info);
                }
              }

              &.is-success {
                border-color: var(--success);

                &.is-active {
                  background: var(--success);
                }
              }

              &.is-purple {
                border-color: var(--purple);

                &.is-active {
                  background: var(--purple);
                }
              }
            }
          }

          .add-link {
            display: inline-block;
            padding: 4px 0;
            font-family: var(--font);
            font-weight: 500;
            font-size: 0.9rem;
            color: var(--primary);
          }
        }
      }
    }
  }
}

.is-dark {
  .form-layout {
    &.is-stacked {
      .form-outer {
        .form-body {
          .form-section {
            border-color: var(--dark-sidebar-light-12);

            &.is-grey {
              background: var(--dark-sidebar-light-4);
            }

            .form-section-header {
              border-color: var(--dark-sidebar-light-12);

              .left {
                h3 {
                  color: var(--dark-dark-text);
                }
              }
            }

            .form-section-inner {
              .add-link {
                color: var(--primary);
              }

              .color-codes {
                .color-code {
                  background: var(--dark-sidebar-light-6);

                  &.is-primary {
                    border-color: var(--primary);
                  }
                }
              }

              .participants {
                .add-participant {
                  &:hover {
                    border: 1.6px solid var(--primary);
                    color: var(--primary);
                  }
                }
              }
            }
          }
        }
      }
    }
  }
}

@media only screen and (width <= 767px) {
  .form-layout {
    &.is-stacked {
      .form-outer {
        .form-body {
          .is-vhidden {
            display: none !important;
          }
        }
      }

      .v-calendar-combo {
        margin: 0 !important;

        .column {
          padding-top: 0 !important;
          padding-bottom: 0 !important;

          &:first-child {
            padding-inline-start: 0 !important;
          }

          &:last-child {
            padding-inline-end: 0 !important;
          }
        }
      }
    }
  }
}
</style>
