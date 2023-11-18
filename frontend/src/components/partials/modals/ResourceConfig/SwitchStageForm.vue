<script setup lang="ts">
import { videoConfig, videoConfig1, videoConfig3 } from './Consts'

const mv = defineModel<any>({
  default: {
    Module: "Switch",
    Input_Number: 10,
    TallyServer_URL: "",
    P4Server_URL: "",
    HW_Panel_URL: "",
    NMOS_DevName: "",
    "OUT_G_2022-7": true
  },
  local: true,
});

const moduleName = ref('切换台')

const inputGlobal = ref<any>({
  "2022-7": true,
  ...videoConfig1
})

const inputs = ref<any[]>([]);

watch(() => mv.value.Input_Number, (nv) => {
  const len = inputs.value.length
  if (nv < len) {
    inputs.value = inputs.value.slice(0, nv);
  } else if (nv > len) {
    inputs.value = inputs.value.concat(
      Array.from({ length: nv - len }, (_, i) => ({
        index: len + i + 1,
        value: ref({
          ...videoConfig3
        })
      }))
    );
  }
}, { immediate: true });

const outputs = Array.from({ length: 3 }, () => ref({
  ...videoConfig
}));
const OUT_3_OPEN = ref(false)
</script>
<template>
  <!-- prettier-ignore -->
  <form
    method="post"
    novalidate
    class="form-layout device-form"
  >
    <div class="form-outer">
      <div class="form-header">
        <div class="form-header-inner">
          <div class="left">
            <h3>设备参数</h3>
          </div>
        </div>
      </div>
      <div class="form-body is-nested">
        <!--Fieldset-->
        <div class="form-fieldset">
          <div class="fieldset-heading">
            <h4>通用参数</h4>
          </div>

          <div class="columns is-multiline">
            <div class="column is-8">
              <VField>
                <VLabel>模块名称</VLabel>
                <VControl>
                  <VInput
                    v-model="moduleName"
                    readonly
                  />
                </VControl>
              </VField>
            </div>
            <div class="column is-4">
              <VField>
                <VLabel>输入信号数量</VLabel>
                <VControl>
                  <VInputNumber
                    v-model="mv.Input_Number"
                    centered
                    :min="0"
                    :step="1"
                  />
                </VControl>
              </VField>
            </div>
            <div class="column is-6">
              <VField>
                <VLabel>Tally服务地址</VLabel>
                <VControl>
                  <VInput
                    v-model="mv.TallyServer_URL"
                  />
                </VControl>
              </VField>
            </div>
            <div class="column is-6">
              <VField>
                <VLabel>P4 交换机通讯地址</VLabel>
                <VControl>
                  <VInput
                    v-model="mv.P4Server_URL"
                  />
                </VControl>
              </VField>
            </div>
            <div class="column is-6">
              <VField>
                <VLabel>硬件控制面板通讯地址</VLabel>
                <VControl>
                  <VInput
                    v-model="mv.HW_Panel_URL"
                  />
                </VControl>
              </VField>
            </div>
            <div class="column is-6">
              <VField>
                <VLabel>NMOS注册设备名称</VLabel>
                <VControl>
                  <VInput
                    v-model="mv.NMOS_DevName"
                  />
                </VControl>
              </VField>
            </div>
          </div>
        </div>
        <!--Fieldset-->
        <div class="form-outer">
          <div class="form-header">
            <div class="form-header-inner">
              <div class="left">
                <h4>输入参数</h4>
              </div>
            </div>
          </div>
          <div class="form-body">
            <!--Fieldset-->
            <div class="form-fieldset">
              <div class="fieldset-heading">
                <h4>全局参数</h4>
              </div>
              <SwitchStageInputGlobal v-model="inputGlobal" />
            </div>
            <SwitchStageInput
              v-for="(ipt, idx) in inputs"
              :key="ipt.index"
              v-model="ipt.value"
              :index="ipt.index"
              :is-last="idx === inputs.length - 1"
              :use-backup="mv['IN_G_2022-7']"
            />
          </div>
        </div>
        <!--Fieldset-->
        <div class="form-outer">
          <div class="form-header">
            <div class="form-header-inner">
              <div class="left">
                <h4>输出参数</h4>
              </div>
            </div>
          </div>
          <div class="form-body">
            <!--Fieldset-->
            <div class="form-fieldset seperator">
              <div class="fieldset-heading">
                <h4>全局参数</h4>
              </div>
              <div class="columns is-multiline">
                <div class="column is-12">
                  <VField horizontal>
                    <VLabel>启用2022-7备份</VLabel>
                    <VControl>
                      <VSwitchBlock
                        v-model="mv['OUT_G_2022-7']"
                        color="primary"
                      />
                    </VControl>
                  </VField>
                </div>
              </div>
            </div>
            <SwitchStageOutput
              v-model="outputs[0]"
              title="PGM参数"
              :use-backup="mv['OUT_G_2022-7']"
            />
            <SwitchStageOutput
              v-model="outputs[1]"
              title="PVW参数"
              :use-backup="mv['OUT_G_2022-7']"
            />
            <SwitchStageOutput
              v-model="outputs[3]"
              v-model:OPEN="OUT_3_OPEN"
              title="Clean参数"
              toggle-title="是否启用Clean输出"
              is-last
              :use-backup="mv['OUT_G_2022-7']"
            />
          </div>
        </div>
      </div>
    </div>
  </form>
</template>

<style lang="scss">
@import "/@src/scss/abstracts/all";
@import "/@src/scss/components/forms-outer";

.form-layout {
  margin: 0 auto;

  .form-fieldset {
    max-width: none;
  }
}
</style>
