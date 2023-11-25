<script setup lang="ts">
import { watchInput, useFormat, watchFormat2, unwrap, getFormat, wrap } from './Utils';
import { formatKeys, type UdxInputType, def_udx_output_params, val_udx, def_udx_input, global_config } from './Consts';
import pick from 'lodash-es/pick'
import udxData from '/@src/data/vscomponent/udx.json'

const mv = defineModel<{
  moudle: string,
  mode: string,
  nmos_devname: string,
  "2110-7_m_local_ip": string,
  "2110-7_b_local_ip": string,
}>({
  default: {
    moudle: "udx",
    mode: "upscale",
    nmos_devname: "",
    "2110-7_m_local_ip": "",
    "2110-7_b_local_ip": "",
  },
  local: true,
});

mv.value = pick(udxData, ['moudle', 'mode', 'nmos_devname', '2110-7_m_local_ip', '2110-7_b_local_ip'])
const moduleName = ref("上下变换")

const input = ref<UdxInputType>(def_udx_input());
input.value = unwrap(udxData.input, 'in_')
const inputFormat = useFormat(input.value, getFormat(input.value))

const OUT_2_OPEN = ref(false)
const output = ref<typeof global_config>({ ...global_config })
const opData = unwrap(udxData.output, 'out_')
output.value = pick(opData, ['g_2022-7', 'g_local_ip1', 'g_local_ip2'])
OUT_2_OPEN.value = opData.params.length > 1
const outputs = Array.from({ length: 2 }, (_, i) => {
  return opData.params[i] ? ref(opData.params[i]) : ref(def_udx_output_params())
});
const outputFormat = ref(getFormat(opData.params[0]))
watchFormat2(outputFormat, outputs.map(o => o.value))

watchInput('audioformat', input.value, outputs.map(o => o.value), { deep: true })

watch(() => mv.value.mode, () => {
  if (mv.value.mode === val_udx[0]) {
    inputFormat.value = inputFormat.value !== formatKeys[2] ? inputFormat.value : formatKeys[0]
    outputFormat.value = formatKeys[2]
  } else {
    inputFormat.value = formatKeys[2]
    outputFormat.value = formatKeys[1]
  }
}, { immediate: true })

function getValue() {
  const mip = mv.value['2110-7_m_local_ip']
  const bip = mv.value['2110-7_b_local_ip']
  const useb = output.value['g_2022-7']
  return {
    ...mv.value,
    input: wrap(input.value, 'in_', input.value['g_2022-7']),
    output: {
      ...wrap(output.value, 'out_'),
      out_params: outputs.map(o => wrap(o.value, 'out_', useb, false, mip, bip)).slice(0, OUT_2_OPEN.value ? 2 : 1)
    }
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
                    <VOption :value="val_udx[0]">
                      上变换
                    </VOption>
                    <VOption :value="val_udx[1]">
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
          <UdxInput
            v-model="input"
            v-model:format="inputFormat"
            :mode="mv.mode"
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
                <div class="column is-4">
                  <VField>
                    <VLabel>启用2022-7备份</VLabel>
                    <VControl>
                      <VSwitchBlock
                        v-model="output['g_2022-7']"
                        color="primary"
                      />
                    </VControl>
                  </VField>
                </div>
                <div class="column is-4">
                  <VField>
                    <VLabel>2022-7主路输出网口IP</VLabel>
                    <VControl>
                      <VInput
                        v-model="output['g_local_ip1']"
                      />
                    </VControl>
                  </VField>
                </div>
                <Transition name="fade-slow">
                  <div
                    v-if="output['g_2022-7']"
                    class="column is-4"
                  >
                    <VField>
                      <VLabel>2022-7备路输出网口IP</VLabel>
                      <VControl>
                        <VInput
                          v-model="output['g_local_ip2']"
                        />
                      </VControl>
                    </VField>
                  </div>
                </Transition>
              </div>
            </div>
            <UdxOutput
              v-model="outputs[0]"
              title="第一路输出参数"
              :format="outputFormat"
              :use-backup="output['g_2022-7']"
              :m_local_ip="mv['2110-7_m_local_ip']"
              :b_local_ip="mv['2110-7_b_local_ip']"
            />
            <UdxOutput
              v-model="outputs[1]"
              v-model:OPEN="OUT_2_OPEN"
              title="第二路输出参数"
              toggle-title="是否启用第二路输出"
              is-last
              :format="outputFormat"
              :use-backup="output['g_2022-7']"
              :m_local_ip="mv['2110-7_m_local_ip']"
              :b_local_ip="mv['2110-7_b_local_ip']"
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
