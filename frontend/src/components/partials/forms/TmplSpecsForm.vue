<script setup lang="ts">
const modelValue = defineModel<any>({
  default: {
    cpu: "",
    cpuNum: "",
    cpuCore: '',
    hugePage: '',
    memory: "",
    disk: "",
    gpu: "",
    inputBand: "",
    outputBand: "",
    network: "",
    chart: "",
    description: "",
    logLevel: 0,
    repairRecvFrame: true,
    repairSendFrame: true,
    dmaList: "",
    hostNetwork: false,
    utfOffset: 37
  },
  local: true,
});

const opened = ref(false)
</script>

<template>
  <form
    method="post"
    novalidate
    class="form-layout specs-form"
    @submit.prevent=""
  >
    <div class="form-outer">
      <div class="form-body">
        <!--Fieldset-->
        <div class="form-fieldset">
          <div class="columns is-multiline">
            <div class="column is-4">
              <VField>
                <VLabel>CPU主频</VLabel>
                <VControl>
                  <VInput
                    v-model="modelValue.cpu"
                    type="text"
                    placeholder=""
                  />
                </VControl>
              </VField>
            </div>
            <div class="column is-4">
              <VField>
                <VLabel>CPU核数</VLabel>
                <VControl>
                  <VInput
                    v-model="modelValue.cpuNum"
                    type="text"
                    placeholder=""
                  />
                </VControl>
              </VField>
            </div>
            <div class="column is-4">
              <VField>
                <VLabel>CPU核心</VLabel>
                <VControl>
                  <VInput
                    v-model="modelValue.cpuCore"
                    type="text"
                    placeholder=""
                  />
                </VControl>
              </VField>
            </div>
            <div class="column is-4">
              <VField>
                <VLabel>大页数(Huge Page)</VLabel>
                <VControl>
                  <VInput
                    v-model="modelValue.hugePage"
                    type="text"
                    placeholder=""
                  />
                </VControl>
              </VField>
            </div>
            <div class="column is-4">
              <VField>
                <VLabel>内存</VLabel>
                <VControl>
                  <VInput
                    v-model="modelValue.memory"
                    type="text"
                    placeholder=""
                  />
                </VControl>
              </VField>
            </div>
            <div class="column is-4">
              <VField>
                <VLabel>GPU</VLabel>
                <VControl>
                  <VInput
                    v-model="modelValue.gpu"
                    type="text"
                    placeholder=""
                  />
                </VControl>
              </VField>
            </div>
            <div class="column is-4">
              <VField>
                <VLabel>硬盘</VLabel>
                <VControl>
                  <VInput
                    v-model="modelValue.disk"
                    type="text"
                    placeholder=""
                  />
                </VControl>
              </VField>
            </div>
            <div class="column is-4">
              <VField>
                <VLabel>信号输入宽带</VLabel>
                <VControl>
                  <VInput
                    v-model="modelValue.inputBand"
                    type="text"
                    placeholder=""
                  />
                </VControl>
              </VField>
            </div>
            <div class="column is-4">
              <VField>
                <VLabel>信号输出宽带</VLabel>
                <VControl>
                  <VInput
                    v-model="modelValue.outputBand"
                    type="text"
                    placeholder=""
                  />
                </VControl>
              </VField>
            </div>
            <div class="column is-4">
              <VField>
                <VLabel>光纤网卡端口数</VLabel>
                <VControl>
                  <VInput
                    v-model="modelValue.network"
                    type="text"
                    placeholder=""
                  />
                </VControl>
              </VField>
            </div>
            <div class="column is-4">
              <VField>
                <VLabel>Chart包名</VLabel>
                <VControl>
                  <VInput
                    v-model="modelValue.chart"
                    type="text"
                    placeholder=""
                  />
                </VControl>
              </VField>
            </div>
            <div class="column is-12">
              <VField>
                <VLabel>备注说明</VLabel>
                <VControl>
                  <VTextarea
                    v-model="modelValue.description"
                    class="textarea"
                    rows="4"
                    autocomplete="off"
                    autocapitalize="off"
                    spellcheck="true"
                  />
                </VControl>
              </VField>
            </div>
          </div>
        </div>
        <div
          class="form-fieldset collapse-form"
          :open="opened || undefined"
        >
          <div
            class="fieldset-heading collapse-header"
            tabindex="0"
            role="button"
            @keydown.space.prevent="opened = !opened"
            @click.prevent="opened = !opened"
          >
            <h4>Debug配置</h4>
            <div class="collapse-icon">
              <VIcon icon="feather:chevron-down" />
            </div>
          </div>
          <Transition name="fade-slow">
            <div
              v-show="opened"
              class="form-fieldset"
            >
              <div class="columns is-multiline">
                <div class="column is-6">
                  <VField>
                    <VLabel>日志级别</VLabel>
                    <VControl>
                      <VInputNumber
                        v-model="modelValue.logLevel"
                        centered
                        :min="0"
                        :max="5"
                        :step="1"
                      />
                    </VControl>
                  </VField>
                </div>
                <div class="column is-2">
                  <VField>
                    <VLabel>共享主机网络</VLabel>
                    <VControl>
                      <VSwitchBlock
                        v-model="modelValue.hostNetwork"
                        color="primary"
                      />
                    </VControl>
                  </VField>
                </div>
                <div class="column is-2">
                  <VField>
                    <VLabel>FR-In</VLabel>
                    <VControl>
                      <VSwitchBlock
                        v-model="modelValue.RepairRecvFrame"
                        color="primary"
                      />
                    </VControl>
                  </VField>
                </div>
                <div class="column is-2">
                  <VField>
                    <VLabel>FR-Out</VLabel>
                    <VControl>
                      <VSwitchBlock
                        v-model="modelValue.RepairSendFrame"
                        color="primary"
                      />
                    </VControl>
                  </VField>
                </div>
                <div class="column is-6">
                  <VField>
                    <VLabel>DMA List</VLabel>
                    <VControl>
                      <VInput
                        v-model="modelValue.DMAList"
                        type="text"
                        placeholder=""
                      />
                    </VControl>
                  </VField>
                </div>
                <div class="column is-6">
                  <VField>
                    <VLabel>UTC Offset</VLabel>
                    <VControl>
                      <VInputNumber
                        v-model="modelValue.utfOffset"
                        centered
                        :step="1"
                      />
                    </VControl>
                  </VField>
                </div>
              </div>
            </div>
          </Transition>
        </div>
      </div>
    </div>
  </form>
</template>

<style lang="scss">
@import "/@src/scss/abstracts/all";
@import "/@src/scss/components/forms-outer";

.specs-form.form-layout {
  max-width: 740px;
  margin: 0 auto;

  .form-outer .form-body {
    padding: 10px 20px;
  }

  .form-fieldset {
    padding: 10px 0;
  }

  .collapse-header {
    display: flex;
    align-items: center;
    justify-content: space-between;
    cursor: pointer;

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

  .collapse-form {
    .fieldset-heading {
      margin-bottom: 0 !important;
      transition: 0.3s ease-in-out;
    }

    &[open] {
      .collapse-header {
        .collapse-icon {
          transform: rotate(calc(var(--transform-direction) * 180deg));
        }
      }

      .fieldset-heading {
        margin-bottom: 20px !important;
      }
    }

    &[open] {
      .collapse-icon {
        border-color: var(--fade-grey-dark-3) !important;
        box-shadow: var(--light-box-shadow);
      }
    }
  }
}

.is-dark {
  .collapse-form {
    &[open] {
      .collapse-header {
        .collapse-icon {
          background: var(--dark-sidebar-light-2);
          border-color: var(--dark-sidebar-light-4) !important;
        }
      }
    }

    .collapse-header {
      .collapse-icon {
        background: var(--dark-sidebar-light-6);
        border-color: var(--dark-sidebar-light-6);
      }
    }
  }
}
</style>
