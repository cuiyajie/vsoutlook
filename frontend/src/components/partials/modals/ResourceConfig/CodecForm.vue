<script setup lang="ts">
import { getFormat, handle, unwrap, useFormat, watchInput, watchNmosName, wrap } from './Utils'
import { def_codec_input, def_codec_output, v_protocols, val_codec, type NMosConfigType, nmos_config, type SSMAddressType, ssm_address } from './Consts';
import pick from 'lodash-es/pick'
import merge from 'lodash-es/merge'
import codecData from '/@src/data/vscomponent/codec.json'

const props = defineProps<{
  name: string
}>()

const mv = defineModel<{
  moudle: string;
  mode: string;
  "2110-7_m_local_ip": string;
  "2110-7_b_local_ip": string;
  nmos: NMosConfigType,
  ssm_address_range: SSMAddressType[]
}>({
  default: {
    moudle: "codec",
    mode: val_codec[0],
    "2110-7_m_local_ip": "",
    "2110-7_b_local_ip": "",
    nmos: { ...nmos_config },
    ssm_address_range: [{ ...ssm_address }],
  },
  local: true,
});
const advanceOpened = ref(false)

mv.value = pick(codecData, ['moudle', 'mode', '2110-7_m_local_ip', '2110-7_b_local_ip', 'nmos', 'ssm_address_range', 'authorization_service'])

const input = ref(def_codec_input());
input.value = unwrap(codecData.input, 'in_')
const inputFormat = useFormat(input, getFormat(input.value))

const output = ref(def_codec_output());
output.value = unwrap(codecData.output, 'out_')
const outputParams = ref(output.value.params)
const outputFormat = useFormat(outputParams, getFormat(output.value.params))

watchNmosName(() => props.name, mv.value)

watch(inputFormat, nv => {
  outputFormat.value = nv
}, { immediate: true })

watch(outputParams, nv => {
  output.value.params = nv
}, { deep: true })

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
  const mip = mv.value['2110-7_m_local_ip']
  const bip = mv.value['2110-7_b_local_ip']
  return {
    ...handle(mv.value),
    input: wrap(input.value, 'in_', input.value['g_2022-7']),
    output: wrap(output.value, 'out_', output.value['g_2022-7'], false, mip, bip)
  }
}

function setValue(data: typeof codecData) {
  mv.value = pick(data, ['moudle', 'mode', '2110-7_m_local_ip', '2110-7_b_local_ip', 'nmos', 'ssm_address_range', 'authorization_service'])
  input.value = merge(def_codec_input(), unwrap(data.input, 'in_'))
  output.value = merge(def_codec_output(), unwrap(data.output, 'out_'))
  outputParams.value = output.value.params
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
            v-model:params="outputParams"
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
