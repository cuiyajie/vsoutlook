<script setup lang="ts">
import { getFormat, unwrap, useFormat, watchInput, wrap } from './Utils'
import { def_codec_input, def_codec_output, v_protocols, val_codec } from './Consts';
import pick from 'lodash-es/pick'
import codecData from '/@src/data/vscomponent/codec.json'

const mv = defineModel<{
  moudle: string;
  mode: string;
  nmos_devname: string;
  "2110-7_m_local_ip": string;
  "2110-7_b_local_ip": string;
}>({
  default: {
    moudle: "codec",
    mode: val_codec[0],
    nmos_devname: "",
    "2110-7_m_local_ip": "",
    "2110-7_b_local_ip": ""
  },
  local: true,
});

mv.value = pick(codecData, ['moudle', 'mode', 'nmos_devname', '2110-7_m_local_ip', '2110-7_b_local_ip'])

const moduleName = ref('编解码器')
const input = ref(def_codec_input());
input.value = unwrap(codecData.input, 'in_')
const inputFormat = useFormat(input.value, getFormat(input.value))

const output = ref(def_codec_output());
output.value = unwrap(codecData.output, 'out_')
const outputFormat = useFormat(output.value.params, getFormat(output.value.params))

watch(inputFormat, nv => {
  outputFormat.value = nv
}, { immediate: true })

watchInput('audioformat', input.value, [output.value.params], { deep: true })

watch(() => mv.value.mode, () => {
  if (mv.value.mode === val_codec[0]) {
    input.value.v_protocol = v_protocols[0]
    output.value.params.v_protocol = v_protocols[1]
  } else {
    input.value.v_protocol = v_protocols[1]
    output.value.params.v_protocol = v_protocols[0]
  }
}, { immediate: true })

function getValue() {
  return {
    ...mv.value,
    input: wrap(input.value, 'in_'),
    output: wrap(output.value, 'out_')
  }
}

defineExpose({
  getValue
})
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
                    v-model="mv.mode"
                    class="is-rounded"
                  >
                    <VOption :value="val_codec[0]">
                      编码器
                    </VOption>
                    <VOption :value="val_codec[1]">
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
                    v-model="mv.nmos_devname"
                  />
                </VControl>
              </VField>
            </div>
            <div class="column is-6">
              <VField>
                <VLabel>2022-7主路收发网口IP</VLabel>
                <VControl>
                  <VInput
                    v-model="mv['2110-7_m_local_ip']"
                  />
                </VControl>
              </VField>
            </div>
            <div class="column is-6">
              <VField>
                <VLabel>2022-7备路收发网口IP</VLabel>
                <VControl>
                  <VInput
                    v-model="mv['2110-7_b_local_ip']"
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
          <CodecInput
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
          <CodecOutput
            v-model="output"
            v-model:format="outputFormat"
            :m_local_ip="mv['2110-7_m_local_ip']"
            :b_local_ip="mv['2110-7_b_local_ip']"
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
