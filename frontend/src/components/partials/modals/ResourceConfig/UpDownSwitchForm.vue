<script setup lang="ts">
import { watchInput, useFormat, watchFormat2 } from './Utils';
import { videoConfig, audioConfig, defs, formatKeys } from './Consts'

const mv = defineModel<any>({
  default: {
    Module: "UDX",
    Mode: "UpScale",
    NMOS_DevName: "",
    "OUT_G_2022-7": true
  },
  local: true,
});

const moduleName = ref("上下变换")

const input = ref<any>({
  '2022-7': true,
  ...videoConfig,
  ...audioConfig,
});
const inputFormat = useFormat(input, defs.format)

const OUT_2_OPEN = ref(false)
const outputs = Array.from({ length: 2 }, () => ref({
  ...videoConfig,
  ...audioConfig
}));
const outputFormat = ref(defs.format)
watchFormat2(outputFormat, outputs)

watchInput([
  'A_Channels_Number',
  'A_Bits',
  'A_Frequency'
], input, outputs)

watch(() => mv.value.Mode, () => {
  if (mv.value.Mode === 'UpScale') {
    outputFormat.value = formatKeys[2]
  } else {
    outputFormat.value = formatKeys[1]
  }
}, { immediate: true })
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
            <div class="column is-4">
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
                <VLabel>工作模式</VLabel>
                <VControl>
                  <VSelect
                    v-model="mv.Mode"
                    class="is-rounded"
                  >
                    <VOption value="UpScale">
                      上变换
                    </VOption>
                    <VOption value="DownScale">
                      下变换
                    </VOption>
                  </VSelect>
                </VControl>
              </VField>
            </div>
            <div class="column is-4">
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
          <UpDownSwitchInput
            v-model="input"
            v-model:format="inputFormat"
            :mode="mv.Mode"
          />
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
            <UpDownSwitchOutput
              v-model="outputs[0]"
              title="第一路输出参数"
              :format="outputFormat"
              :use-backup="mv['OUT_G_2022-7']"
            />
            <UpDownSwitchOutput
              v-model="outputs[1]"
              v-model:OPEN="OUT_2_OPEN"
              title="第二路输出参数"
              toggle-title="是否启用第二路输出"
              is-last
              :format="outputFormat"
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
