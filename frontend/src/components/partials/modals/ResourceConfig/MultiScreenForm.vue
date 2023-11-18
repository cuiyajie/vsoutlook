<script setup lang="ts">
import { videoConfig, audioConfig, screenConfig, videoConfig2 } from './Consts';

const mv = defineModel<any>({
  default: {
    Module: "MV",
    Input_Number: 10,
    Output_Number: 1,
    Tally_Port: 6001,
    NMOS_DevName: "",
    "IN_G_2022-7": true,
    "OUT_G_2022-7": true
  },
  local: true,
});

const moduleName = ref('多画面')

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
          ...audioConfig,
          ...videoConfig,
        })
      }))
    );
  }
}, { immediate: true });

const outputs = ref<any[]>([]);
watch(() => mv.value.Output_Number, (nv) => {
  const len = outputs.value.length
  if (nv < len) {
    outputs.value = outputs.value.slice(0, nv);
  } else if (nv > len) {
    outputs.value = outputs.value.concat(
      Array.from({ length: nv - len }, (_, i) => ({
        index: len + i + 1,
        value: ref({
          ...videoConfig2,
          ...screenConfig
        })
      }))
    );
  }
}, { immediate: true });
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
            <div class="column is-3">
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
            <div class="column is-3">
              <VField>
                <VLabel>输入信号数量</VLabel>
                <VControl>
                  <VInputNumber
                    v-model="mv.Input_Number"
                    centered
                    :min="0"
                    :max="10"
                    :step="1"
                  />
                </VControl>
              </VField>
            </div>
            <div class="column is-3">
              <VField>
                <VLabel>输出信号数量</VLabel>
                <VControl>
                  <VSelect
                    v-model="mv.Output_Number"
                    class="is-rounded"
                  >
                    <VOption :value="1">
                      1
                    </VOption>
                    <VOption :value="2">
                      2
                    </VOption>
                  </VSelect>
                </VControl>
              </VField>
            </div>
            <div class="column is-3">
              <VField>
                <VLabel>Tally服务端口</VLabel>
                <VControl>
                  <VInputNumber
                    v-model="mv.Tally_Port"
                    centered
                    :min="0"
                    :max="65535"
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

              <div class="columns is-multiline">
                <div class="column is-12">
                  <VField horizontal>
                    <VLabel>启用2022-7备份</VLabel>
                    <VControl>
                      <VSwitchBlock
                        v-model="mv['IN_G_2022-7']"
                        color="primary"
                      />
                    </VControl>
                  </VField>
                </div>
              </div>
            </div>
            <MultiScreenInput
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
            <MultiScreenOutput
              v-for="(opt, idx) in outputs"
              :key="opt.index"
              v-model="opt.value"
              :index="opt.index"
              :is-last="idx === outputs.length - 1"
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
</style>
