<script setup lang="ts">
import { type AuthServiceType, type NMosConfigType, type SSMAddressType, auth_service, def_mv_input_params, def_mv_output_params, global_config, nmos_config, ssm_address } from './Consts';
import { handle, unwrap, watchNmosName, wrap } from "./Utils";
import pick from 'lodash-es/pick'
import merge from 'lodash-es/merge'
import mvData from '/@src/data/vscomponent/mv.json'

const props = defineProps<{
  name: string
}>()

const mv = defineModel<{
  moudle: string,
  input_number: number,
  output_number: number,
  tally_port: number,
  "2110-7_m_local_ip": string,
  "2110-7_b_local_ip": string,
  nmos: NMosConfigType,
  ssm_address_range: SSMAddressType[],
  authorization_service: AuthServiceType[]
}>({
  default: {
    moudle: "mv",
    input_number: 10,
    output_number: 1,
    tally_port: 6000,
    "2110-7_m_local_ip": "",
    "2110-7_b_local_ip": "",
    nmos: { ...nmos_config },
    ssm_address_range: [{ ...ssm_address }],
    authorization_service: [{ ...auth_service }]
  },
  local: true,
});

mv.value = pick(mvData, ['moudle', 'input_number', 'output_number', 'tally_port', '2110-7_m_local_ip', '2110-7_b_local_ip', 'nmos', 'ssm_address_range', 'authorization_service'])

watchNmosName(() => props.name, mv.value)

const ipData = unwrap(mvData.input, 'in_')
const input = ref<{ "g_2022-7": boolean }>({ "g_2022-7": false });
input.value = pick(ipData, ['g_2022-7'])
const inputs = ref<any[]>([]);
watch(() => mv.value.input_number, (nv) => {
  const len = inputs.value.length
  if (nv < len) {
    inputs.value = inputs.value.slice(0, nv);
  } else if (nv > len) {
    inputs.value = inputs.value.concat(
      Array.from({ length: nv - len }, (_, i) => {
        return {
          index: len + i + 1,
          value: ipData.input_params[len + i] ? ref(ipData.input_params[len + i]) : ref(def_mv_input_params())
        }
      })
    );
  }
}, { immediate: true });

const opData = unwrap(mvData.output, 'out_')
const output = ref<any>({ ...global_config });
output.value = pick(opData, Object.keys(global_config))
const outputs = ref<any[]>([]);
const outputRefs = ref<any[]>([]);
watch(() => mv.value.output_number, (nv) => {
  const len = outputs.value.length
  if (nv < len) {
    outputs.value = outputs.value.slice(0, nv);
  } else if (nv > len) {
    outputs.value = outputs.value.concat(
      Array.from({ length: nv - len }, (_, i) => {
        return {
          index: len + i + 1,
          value: opData.params[len + i] ? ref(opData.params[len + i]) : ref(def_mv_output_params())
        }
      })
    );
  }
}, { immediate: true });

function getValue() {
  const mip = mv.value['2110-7_m_local_ip']
  const bip = mv.value['2110-7_b_local_ip']
  const useb = output.value['g_2022-7']
  return {
    ...handle(mv.value),
    input: {
      ...wrap(input.value, 'in_', input.value['g_2022-7']),
      input_params: inputs.value.map(ipt => wrap(ipt.value, 'in_', input.value['g_2022-7'], true, mip, bip))
    },
    output: {
      ...wrap(output.value, 'out_'),
      out_params: outputs.value.map((o, i) => {
        let v = o.value
        const optRef = outputRefs.value[i]
        if (optRef?.getValue) {
          v = optRef.getValue()
        }
        return wrap(v, 'out_', useb, false, mip, bip)
      })
    }
  }
}

function setValue(data: typeof mvData) {
  mv.value = pick(data, ['moudle', 'input_number', 'output_number', 'tally_port', '2110-7_m_local_ip', '2110-7_b_local_ip', 'nmos', 'ssm_address_range', 'authorization_service'])
  const _ipData = unwrap(data.input, 'in_')
  input.value = pick(_ipData, ['g_2022-7'])
  nextTick(() => {
    inputs.value.forEach((iptv, idx) => {
      iptv.value = merge(def_mv_input_params(), _ipData.input_params[idx])
    })
  })

  const _opData = unwrap(data.output, 'out_')
  output.value = pick(_opData, Object.keys(global_config))
  nextTick(() => {
    outputs.value.forEach((optv, idx) => {
      optv.value = merge(def_mv_output_params(), _opData.params[idx])
    })
  })
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
            <div class="column is-4">
              <VField>
                <VLabel>输入信号数量</VLabel>
                <VControl>
                  <VInputNumber
                    v-model="mv.input_number"
                    centered
                    :min="0"
                    :max="36"
                    :step="1"
                  />
                </VControl>
              </VField>
            </div>
            <div class="column is-4">
              <VField>
                <VLabel>输出信号数量</VLabel>
                <VControl>
                  <VSelect
                    v-model="mv.output_number"
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
            <div class="column is-4">
              <VField>
                <VLabel>tally服务端口</VLabel>
                <VControl>
                  <VInputNumber
                    v-model="mv.tally_port"
                    centered
                    :min="0"
                    :max="65535"
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
        <NMosConfig v-model="mv.nmos" />
        <SSMAddressRange v-model="mv.ssm_address_range" />
        <AuthorizationService v-model="mv.authorization_service" />
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
                        v-model="input['g_2022-7']"
                        color="primary"
                      />
                    </VControl>
                  </VField>
                </div>
              </div>
            </div>
            <MVInput
              v-for="(ipt, idx) in inputs"
              :key="ipt.index"
              v-model="ipt.value"
              :index="ipt.index"
              :is-last="idx === inputs.length - 1"
              :use-backup="input['g_2022-7']"
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
            <MVOutput
              v-for="(opt, idx) in outputs"
              ref="outputRefs"
              :key="opt.index"
              v-model="opt.value"
              :index="opt.index"
              :is-last="idx === outputs.length - 1"
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
