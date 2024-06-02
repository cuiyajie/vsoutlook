<script setup lang="ts">
import { useUserSession } from "@src/stores/userSession";
import { confirm } from "@src/utils/dialog";

const usStore = useUserSession();
const settings = computed(() => usStore.settings);
const router = useRouter();

/* ------------------------ Auth Service -------------------------- */
const def_auth_service = {
	ip: "232.0.0.0",
  port: 6001
}

const authServices = ref<SettingAuthService[]>([]);

watch(() => settings.value.authorization_services, () => {
  authServices.value = settings.value.authorization_services || []
}, { immediate: true })

function addAuthService() {
  authServices.value.push({ ...def_auth_service })
  updateProp('authorization_services')
}

function removeAuthService(idx: number) {
  confirm({
    title: "确认",
    content: `确定要删除授权服务地址 ${idx} 吗？`,
    onConfirm: async (hide) => {
      authServices.value.splice(idx, 1)
      updateProp('authorization_services')
      hide()
    },
  });
}

/* ----------------------------------------------------------------- */

/* ------------------------ Mv Service -------------------------- */
// const def_mv_template = {
// 	name: "vso media software",
//   path: "/opt/vsomediasoftware/config/viewlayout.json"
// }

const mvTemplates = ref<Array<SettingMvTemplate>>([]);

watch(() => settings.value.mv_template_list, () => {
  mvTemplates.value = settings.value.mv_template_list || []
}, { immediate: true })

function addMvTemplate() {
  // mvTemplates.value.push({ ...def_mv_template })
  // updateProp('mv_template_list')
  router.push('/app/mtv')
}

function removeMvTemplate(idx: number) {
  confirm({
    title: "确认",
    content: `确定要删除多画面布局 ${mvTemplates.value[idx]?.name} 吗？`,
    onConfirm: async (hide) => {
      mvTemplates.value.splice(idx, 1)
      updateProp('mv_template_list')
      hide()
    },
  });
}

/* ----------------------------------------------------------------- */

function updateProp(key: string, value?: string) {
  if (key === 'authorization_services') {
    usStore.$updateSettings({ key, value: JSON.stringify(authServices.value) })
  } else if (key === 'mv_template_list') {
    usStore.$updateSettings({ key, value: JSON.stringify(mvTemplates.value) })
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
              <h3>通用配置</h3>
            </div>
          </div>

          <div class="form-section-inner flex-auto">
            <VField
              class="field-switch"
              horizontal
              label="自动保存容器配置"
            >
              <VControl>
                <VSwitchBlock
                  v-model="settings.auto_save_container_config"
                  color="primary"
                  @update:model-value="val => updateProp('auto_save_container_config', String(val))"
                />
              </VControl>
            </VField>
          </div>
        </div>

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

        <div class="form-section is-grey">
          <div class="form-section-header">
            <div class="left">
              <h3>多画面布局</h3>
            </div>
            <button
              type="button"
              class="button is-circle is-dark-outlined"
              @keydown.space.prevent="addMvTemplate"
              @click.prevent="addMvTemplate"
            >
              <span class="icon is-large">
                <i aria-hidden="true" class="iconify" data-icon="feather:plus" />
              </span>
            </button>
          </div>
          <div class="form-section-inner flex-auto">
            <VField
              horizontal
              label="多画面模板字体"
            >
              <VControl
                icon="lnir lnir-text"
              >
                <VInput
                  v-model="settings.mv_template_font"
                  v-blur-on-enter
                  type="url"
                  placeholder="Noto Serif CJK"
                  inputmode="url"
                  data-key="mv_template_font"
                  style="width: 300px;"
                />
              </VControl>
              <VLabel class="font-test-label" :style="{fontFamily: `'${settings.mv_template_font}'`}">字体示例，Font Sample.</VLabel>
            </VField>
          </div>
          <div class="form-section-inner row-with-button">
            <div class="left">多画面布局模板</div>
            <button
              type="button"
              class="button is-circle is-dark-outlined"
              @keydown.space.prevent="addMvTemplate"
              @click.prevent="addMvTemplate"
            >
              <span class="icon is-large">
                <i aria-hidden="true" class="iconify" data-icon="feather:plus" />
              </span>
            </button>
          </div>
          <div v-for="(mvt, idx) in mvTemplates" :key="idx" class="columns is-multiline">
            <div class="column is-1">
              <VField class="index-field">
                <VLabel>序号</VLabel>
                <VControl>
                  <VLabel>{{ idx }}</VLabel>
                </VControl>
              </VField>
            </div>
            <div class="column is-4">
              <VField id="ip">
                <VLabel>模板名称</VLabel>
                <VControl icon="feather:copy">
                  <input
                    v-model="mvt.name"
                    v-blur-on-enter
                    type="text"
                    readonly
                    class="input is-primary-focus"
                    data-key="mv_template_list"
                  >
                </VControl>
              </VField>
            </div>
            <div class="column is-6">
              <VField>
                <VLabel>模板路径及模板文件名</VLabel>
                <VControl icon="feather:paperclip">
                  <input
                    v-model="mvt.path"
                    v-blur-on-enter
                    type="text"
                    readonly
                    class="input is-primary-focus"
                    data-key="mv_template_list"
                  >
                </VControl>
              </VField>
            </div>
            <div class="column is-1">
              <VField>
                <VLabel>&nbsp;</VLabel>
                <VControl>
                  <VIconButton color="warning" light raised circle icon="feather:x" @click="removeMvTemplate(idx)" />
                </VControl>
              </VField>
            </div>
          </div>
          <div v-if="mvTemplates.length === 0" class="form-empty">暂时没有添加多画面布局模板</div>
        </div>
      </div>
    </div>
  </form>
</template>

<style lang="scss">
@import "@src/scss/abstracts/all";
@import "@src/scss/components/forms-outer";

.is-navbar {
  .form-layout {
    margin-top: 30px;
  }
}

.form-layout {
  max-width: 876px;
  margin: 0 auto;

  &.gconfig {

    .form-outer {
      border-radius: 16px;
      height: calc(100vh - 225px);
      overflow: auto;
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

    .field-switch {

      .field-body {
        transform: translateY(8px);
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

            .font-test-label {
              font-size: 1rem;
              color: var(--light-text);
              font-weight: 400;
              padding-top: 10px;
              margin-left: 20px;
            }

            &.row-with-button {
              display: flex;
              align-items: center;
              justify-content: space-between;
              padding: 30px 12px 30px 0;

              .left {
                font-family: var(--font);
                font-size: 0.9rem;
                color: var(--light-text) !important;
                font-weight: 400;
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
