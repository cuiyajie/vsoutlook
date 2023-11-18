<script setup lang="ts">
import { useFormat, watchInput } from './Utils'
import { videoConfig, audioConfig, defs, vProtocols } from './Consts';

const mv = defineModel<{
  Module: string;
  Mode: string;
  NMOS_DevName: string;
}>({
  default: {
    Module: "EnCoder/Decoder",
    Mode: "Encoder",
    NMOS_DevName: ""
  },
  local: true,
});

const moduleName = ref('编解码器')
const input = ref<any>({
  '2022-7': true,
  ...videoConfig,
  ...audioConfig,
  V_Protocol: defs.vProtocol,
});
const inputFormat = useFormat(input)

const output = ref<any>({
  '2022-7': true,
  ...videoConfig,
  ...audioConfig,
  V_Protocol: vProtocols[1],
});
const outputFormat = useFormat(output)
watch(inputFormat, nv => {
  outputFormat.value = nv
}, { immediate: true })

watchInput([
  'A_Channels_Number',
  'A_Bits',
  'A_Frequency'
], input, [output])

watch(() => mv.value.Mode, () => {
  if (mv.value.Mode === 'Encoder') {
    input.value.V_Protocol = vProtocols[0]
    output.value.V_Protocol = vProtocols[1]
  } else {
    input.value.V_Protocol = vProtocols[1]
    output.value.V_Protocol = vProtocols[0]
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
                    <VOption value="Encoder">
                      编码器
                    </VOption>
                    <VOption value="Decoder">
                      解码器
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
          <DecoderInput
            v-model="input"
            v-model:format="inputFormat"
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
          <DecoderOutput
            v-model="output"
            v-model:format="outputFormat"
          />
        </div>
      </div>
    </div>
  </form>
</template>

<style lang="scss">
@import "/@src/scss/abstracts/all";
@import "/@src/scss/components/forms-outer";
</style>
