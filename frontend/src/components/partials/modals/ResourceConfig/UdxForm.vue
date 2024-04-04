<script setup lang="ts">
import { watchInput, useFormat, unwrap, getFormat, wrap, watchNmosName, handle } from './Utils';
import { formatKeys, type UdxInputType, def_udx_output_params, val_udx, def_udx_input, global_config, type NMosConfigType, type SSMAddressType, nmos_config, ssm_address } from './Consts';
import pick from 'lodash-es/pick'
import merge from 'lodash-es/merge'
import udxData from '/@src/data/vscomponent/udx.json'

const props = defineProps<{
  name: string
}>()

const mv = defineModel<{
  moudle: string,
  mode: string,
  "2110-7_m_local_ip": string,
  "2110-7_b_local_ip": string,
  nmos: NMosConfigType,
  ssm_address_range: SSMAddressType[],
}>({
  default: {
    moudle: "udx",
    mode: "upscale",
    "2110-7_m_local_ip": "",
    "2110-7_b_local_ip": "",
    nmos: { ...nmos_config },
    ssm_address_range: [{ ...ssm_address }],
  },
  local: true,
});
const advanceOpened = ref(false)

mv.value = pick(udxData, ['moudle', 'mode', '2110-7_m_local_ip', '2110-7_b_local_ip', 'nmos', 'ssm_address_range', 'authorization_service'])

const input = ref<UdxInputType>(def_udx_input());
input.value = unwrap(udxData.input, 'in_')
const inputFormat = useFormat(input, getFormat(input.value))

const OUT_2_OPEN = ref(false)
const output = ref<typeof global_config>({ ...global_config })
const opData = unwrap(udxData.output, 'out_')
output.value = pick(opData, ['g_2022-7'])
OUT_2_OPEN.value = opData.params.length > 1
const outputs = ref<any[]>([])
outputs.value = Array.from({ length: 2 }, (_, i) => {
  return opData.params[i] ? ref(opData.params[i]) : ref(def_udx_output_params())
});
const output1Format = useFormat(outputs.value[0], getFormat(opData.params[0]))
const output2Format = useFormat(outputs.value[1], getFormat(opData.params[1]))

watchNmosName(() => props.name, mv)
watchInput('audioformat', input, outputs.value.map(o => o), { deep: true })

watch(() => mv.value.mode, () => {
  if (mv.value.mode === val_udx[0]) {
    inputFormat.value = inputFormat.value !== formatKeys[2] ? inputFormat.value : formatKeys[0]
    output1Format.value = formatKeys[2]
    output2Format.value = formatKeys[1]
  } else {
    inputFormat.value = formatKeys[2]
    output1Format.value = output1Format.value !== formatKeys[2] ? output1Format.value : formatKeys[1]
    output2Format.value = output2Format.value !== formatKeys[2] ? output2Format.value : formatKeys[1]
  }
}, { immediate: true })

function getValue() {
  const mip = mv.value['2110-7_m_local_ip']
  const bip = mv.value['2110-7_b_local_ip']
  const useb = output.value['g_2022-7']
  return {
    ...handle(mv.value),
    input: wrap(input.value, 'in_', input.value['g_2022-7']),
    output: {
      ...wrap(output.value, 'out_'),
      out_params: outputs.value.map((o, i) => Object.assign(wrap(unref(o), 'out_', useb, false, mip, bip), { index: i }))
        .slice(0, (OUT_2_OPEN.value && mv.value.mode !== val_udx[1]) ? 2 : 1)
    }
  }
}

function setValue(data: typeof udxData) {
  mv.value = pick(data, ['moudle', 'mode', '2110-7_m_local_ip', '2110-7_b_local_ip', 'nmos', 'ssm_address_range', 'authorization_service'])
  input.value = merge(def_udx_input(), unwrap(data.input, 'in_'))
  inputFormat.value = getFormat(input.value)

  const _opData = unwrap(data.output, 'out_')
  output.value = pick(_opData, ['g_2022-7'])
  OUT_2_OPEN.value = _opData.params.length > 1
  outputs.value.forEach((o, i) => {
    o.value = merge(def_udx_output_params(), _opData.params[i])
  })
  output1Format.value = getFormat(outputs.value[0].value)
  output2Format.value = getFormat(outputs.value[1].value)
}

defineExpose({
  getValue,
  setValue
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
            <div class="column is-6">
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
            <div class="column is-6">
              <VField>
                <VLabel>对外IP</VLabel>
                <VControl>
                  <VInput
                    v-model="mv.nmos.host_addresses"
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
        <div
          class="form-fieldset collapse-form form-outer"
          :open="advanceOpened || undefined"
        >
          <div
            class="fieldset-heading collapse-header"
            tabindex="0"
            role="button"
            @keydown.space.prevent="advanceOpened = !advanceOpened"
            @click.prevent="advanceOpened = !advanceOpened"
          >
            <h4>高级配置</h4>
            <div class="collapse-icon">
              <VIcon icon="feather:chevron-down" />
            </div>
          </div>
          <Transition name="fade-slow">
            <div v-show="advanceOpened" class="form-body">
              <NMosConfig v-model="mv.nmos" class="seperator" />
              <SSMAddressRange v-model="mv.ssm_address_range" class="seperator" />
            </div>
          </Transition>
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
                <div class="column is-12">
                  <VField class="is-horizontal">
                    <VLabel>启用2022-7备份</VLabel>
                    <VControl>
                      <VSwitchBlock
                        v-model="output['g_2022-7']"
                        color="primary"
                      />
                    </VControl>
                  </VField>
                </div>
              </div>
            </div>
            <UdxOutput
              v-model="outputs[0]"
              v-model:format="output1Format"
              :mode="mv.mode"
              title="第一路输出参数"
              type="output1"
              :use-backup="output['g_2022-7']"
              :m_local_ip="mv['2110-7_m_local_ip']"
              :b_local_ip="mv['2110-7_b_local_ip']"
            />
            <UdxOutput
              v-if="mv.mode === val_udx[0]"
              v-model="outputs[1]"
              v-model:format="output2Format"
              v-model:OPEN="OUT_2_OPEN"
              :mode="mv.mode"
              type="output2"
              title="第二路输出参数"
              toggle-title="是否启用第二路输出"
              is-last
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
