<script setup lang="ts">
const modelValue = defineModel<TmplRequirement & { description: string }>({
  default: {
    cpu: "",
    cpuNum: 0,
    cpuCore: "",
    hugePage: 0,
    memory: 0,
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
    hostNetwork: false,
    utfOffset: 37,
    recvAVFrameNodeCount: 2,
    sendAVFrameNodeCount: 2,
    recvFrameCount: 2,
    imageTag: "",
    maxRateMbpsByCore: 9000,
    receiveSessions: 18,
    shm: 0,
    nicCount: 0,
    nicConfig: [],
    coreType: 'txrx',
    txShare: 1,
    rxShare: 1,
  },
  local: true,
});

const onNicCountChange = (e: Event) => {
  const len = +(e.target as HTMLSelectElement).value;
  if (len > modelValue.value.nicConfig.length) {
    modelValue.value.nicConfig = modelValue.value.nicConfig.concat(
      Array.from({ length: len - modelValue.value.nicConfig.length }, () => ({
        dpdkCpu: 1,
        sdkCpu: 1,
        dma: 1,
      }))
    );
  } else {
    modelValue.value.nicConfig = modelValue.value.nicConfig.slice(0, len);
  }
};

const coreOptions = [
  { label: "DPDK核心列表", value: "dpdk" },
  { label: "应用收发流核心列表", value: "txrx" },
  { label: "不指定", value: "none" },
]
</script>

<template>
  <form method="post" novalidate class="form-layout specs-form" @submit.prevent="">
    <div class="form-outer">
      <div class="form-body">
        <!--Fieldset-->
        <div class="form-fieldset">
          <div class="fieldset-heading">
            <h4>容器资源配置</h4>
          </div>
          <div class="form-body">
            <div class="columns is-multiline">
              <div class="column is-4">
                <VField>
                  <VLabel>CPU 总核心数</VLabel>
                  <VControl>
                    <VInputNumber v-model="modelValue.cpuNum" centered :min="0" :step="1" />
                  </VControl>
                </VField>
              </div>
              <div class="column is-4">
                <VField>
                  <VLabel>内存 (GB)</VLabel>
                  <VControl>
                    <VInputNumber v-model="modelValue.memory" centered :min="0" :step="1" />
                  </VControl>
                </VField>
              </div>
              <div class="column is-4">
                <VField>
                  <VLabel>大页内存 (GB)</VLabel>
                  <VControl>
                    <VInputNumber v-model="modelValue.hugePage" centered :min="0" :step="1" />
                  </VControl>
                </VField>
              </div>
              <div class="column is-4">
                <VField>
                  <VLabel>共享内存 (GB)</VLabel>
                  <VControl>
                    <VInputNumber v-model="modelValue.shm" centered :min="0" :step="1" />
                  </VControl>
                </VField>
              </div>
              <div class="column is-4">
                <VField>
                  <VLabel>Image Tag</VLabel>
                  <VControl>
                    <VInput v-model="modelValue.imageTag" />
                  </VControl>
                </VField>
              </div>
            </div>
          </div>
          <div class="fieldset-heading" style="margin: 20px 0">
            <h4>底层收发流核心配置</h4>
          </div>
          <div class="form-body">
            <div class="columns is-multiline">
              <div class="column is-4">
                <VField>
                  <VLabel>日志级别</VLabel>
                  <VControl>
                    <VInputNumber v-model="modelValue.logLevel" centered :min="0" :max="5" :step="1" />
                  </VControl>
                </VField>
              </div>
              <div class="column is-4">
                <VField>
                  <VLabel>单核心最大处理能力(Mb)</VLabel>
                  <VControl>
                    <VInputNumber v-model="modelValue.maxRateMbpsByCore" centered :min="0" :step="1" />
                  </VControl>
                </VField>
              </div>
              <div class="column is-4">
                <VField>
                  <VLabel>收流会话数</VLabel>
                  <VControl>
                    <VInputNumber v-model="modelValue.receiveSessions" centered :min="0" :step="1" />
                  </VControl>
                </VField>
              </div>
              <div class="column is-4">
                <VField>
                  <VLabel>SendAVFrameNodeCount</VLabel>
                  <VControl>
                    <VInputNumber v-model="modelValue.sendAVFrameNodeCount" centered :min="0" :step="1" />
                  </VControl>
                </VField>
              </div>
              <div class="column is-4">
                <VField>
                  <VLabel>RecvFrameCount</VLabel>
                  <VControl>
                    <VInputNumber v-model="modelValue.recvFrameCount" centered :min="0" :step="1" />
                  </VControl>
                </VField>
              </div>
              <div class="column is-4">
                <VField>
                  <VLabel>需要使用DPDK的网卡数量</VLabel>
                  <VControl>
                    <VSelect v-model="modelValue.nicCount" class="is-rounded" @change="onNicCountChange">
                      <VOption
                        v-for="count in Array(5)
                          .fill(0)
                          .map((v, i) => v + i)" :key="count" :value="count"
                      >
                        {{ count }}
                      </VOption>
                    </VSelect>
                  </VControl>
                </VField>
              </div>
              <template v-for="(config, cidx) in modelValue.nicConfig" :key="cidx">
                <div class="column is-1">
                  <VField>
                    <VLabel>序号</VLabel>
                    <VControl class="is-indexed" static>{{ cidx }}</VControl>
                  </VField>
                </div>
                <div class="column is-4">
                  <VField>
                    <VLabel>DPDK CPU核心数</VLabel>
                    <VControl>
                      <VInputNumber v-model="config.dpdkCpu" centered :min="0" :step="1" />
                    </VControl>
                  </VField>
                </div>
                <div class="column is-4">
                  <VField>
                    <VLabel>SDK CPU核心数</VLabel>
                    <VControl>
                      <VInputNumber v-model="config.sdkCpu" centered :min="0" :step="1" />
                    </VControl>
                  </VField>
                </div>
                <div class="column is-3">
                  <VField>
                    <VLabel>DMA数量</VLabel>
                    <VControl>
                      <VInputNumber v-model="config.dma" centered :min="0" :step="1" />
                    </VControl>
                  </VField>
                </div>
              </template>
            </div>
          </div>
          <div class="fieldset-heading" style="margin: 20px 0">
            <h4>应用收发流核心配置</h4>
          </div>
          <div class="form-body">
            <div class="columns is-multiline">
              <div class="column is-4">
                <VField>
                  <VLabel>使用的核心</VLabel>
                  <VControl>
                    <VSelect v-model="modelValue.coreType" class="is-rounded">
                      <VOption
                        v-for="option in coreOptions" :key="option.value" :value="option.value"
                      >
                        {{ option.label }}
                      </VOption>
                    </VSelect>
                  </VControl>
                </VField>
              </div>
              <div class="column is-4">
                <VField>
                  <VLabel>几个收流使用一个核心</VLabel>
                  <VControl>
                    <VInputNumber v-model="modelValue.rxShare" centered :min="1" :step="1" />
                  </VControl>
                </VField>
              </div>
              <div class="column is-4">
                <VField>
                  <VLabel>几个发流使用一个核心</VLabel>
                  <VControl>
                    <VInputNumber v-model="modelValue.txShare" centered :min="1" :step="1" />
                  </VControl>
                </VField>
              </div>
            </div>
          </div>
          <div class="fieldset-heading" style="margin: 20px 0">
            <h4>备注说明</h4>
          </div>
          <div class="form-body">
            <div class="columns is-multiline">
              <!-- <div class="column is-4">
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
            </div> -->
              <!-- <div class="column is-4">
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
            </div> -->
              <!-- <div class="column is-4">
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
            </div> -->
              <div class="column is-12">
                <VField>
                  <!-- <VLabel>备注说明</VLabel> -->
                  <VControl>
                    <VTextarea
                      v-model="modelValue.description" class="textarea" rows="4" autocomplete="off"
                      autocapitalize="off" spellcheck="true"
                    />
                  </VControl>
                </VField>
              </div>
            </div>
          </div>
        </div>
        <!-- <div
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
                <div class="column is-4">
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
                <div class="column is-4">
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
                <div class="column is-4 pl-6">
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
                <div class="column is-3">
                  <VField>
                    <VLabel>FR-In</VLabel>
                    <VControl>
                      <VSwitchBlock
                        v-model="modelValue.repairRecvFrame"
                        color="primary"
                      />
                    </VControl>
                  </VField>
                </div>
                <div class="column is-3">
                  <VField>
                    <VLabel>FR-Out</VLabel>
                    <VControl>
                      <VSwitchBlock
                        v-model="modelValue.repairSendFrame"
                        color="primary"
                      />
                    </VControl>
                  </VField>
                </div>
                <div class="column is-4">
                  <VField>
                    <VLabel>RecvAVFrameNodeCount</VLabel>
                    <VControl>
                      <VInputNumber
                        v-model="modelValue.recvAVFrameNodeCount"
                        centered
                        :min="0"
                        :step="1"
                      />
                    </VControl>
                  </VField>
                </div>
                <div class="column is-4">
                  <VField>
                    <VLabel>SendAVFrameNodeCount</VLabel>
                    <VControl>
                      <VInputNumber
                        v-model="modelValue.sendAVFrameNodeCount"
                        centered
                        :min="0"
                        :step="1"
                      />
                    </VControl>
                  </VField>
                </div>
                <div class="column is-4">
                  <VField>
                    <VLabel>Recvframe Cnt</VLabel>
                    <VControl>
                      <VInputNumber
                        v-model="modelValue.recvFrameCount"
                        centered
                        :min="0"
                        :step="1"
                      />
                    </VControl>
                  </VField>
                </div>
                <div class="column is-4">
                  <VField>
                    <VLabel>单核心最大处理能力(Mb)</VLabel>
                    <VControl>
                      <VInputNumber
                        v-model="modelValue.maxRateMbpsByCore"
                        centered
                        :min="0"
                        :step="1"
                      />
                    </VControl>
                  </VField>
                </div>
                <div class="column is-4">
                  <VField>
                    <VLabel>收流会话数</VLabel>
                    <VControl>
                      <VInputNumber
                        v-model="modelValue.receiveSessions"
                        centered
                        :min="0"
                        :step="1"
                      />
                    </VControl>
                  </VField>
                </div>
              </div>
            </div>
          </Transition>
        </div> -->
      </div>
    </div>
  </form>
</template>

<style lang="scss">
@import "@src/scss/abstracts/all";
@import "@src/scss/components/forms-outer";

.specs-form.form-layout {
  max-width: 740px;
  margin: 0 auto;
  height: 100%;

  .form-outer {
    height: 100%;
    overflow: auto;
  }

  .form-outer .form-body {
    padding: 10px 20px;
  }

  .form-fieldset {
    padding: 10px 0;
    max-width: none;
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

  .control.is-indexed {
    line-height: 36px;
    background-color: var(--dark-sidebar-light-2);
    border: 1px solid var(--dark-sidebar-light-4);
    border-radius: var(--radius);
    text-align: center;
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
