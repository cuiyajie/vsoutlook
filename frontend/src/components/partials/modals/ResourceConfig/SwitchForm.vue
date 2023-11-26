<script setup lang="ts">
import { def_switch_input_params, def_switch_output_params, def_switch_input_video_params, type SwitchInputParamsType, global_config, def_switch_input_bus_params, type SwitchInputBusParamsType } from './Consts';
import { unwrap, wrap } from "./Utils";
import pick from 'lodash-es/pick'
import switchData from '/@src/data/vscomponent/switch.json'

const mv = defineModel<{
  moudle: string,
  input_number: number,
  tallyserver_url: string,
  p4server_url: string,
  hw_panel_url: string,
  nmos_devname: string,
  physic_nic_port0_IP: string,
  physic_nic_port1_IP: string,
  "2110-7_m_local_ip": string,
  "2110-7_b_local_ip": string
}>({
  default: {
    moudle: "switch",
    input_number: 10,
    tallyserver_url: "",
    p4server_url: "",
    hw_panel_url: "",
    nmos_devname: "",
    physic_nic_port0_IP: "",
    physic_nic_port1_IP: "",
    "2110-7_m_local_ip": "",
    "2110-7_b_local_ip": ""
  },
  local: true,
});

const moduleName = ref('切换台')
mv.value = pick(switchData, ['moudle', 'input_number', 'tallyserver_url', 'p4server_url', 'hw_panel_url', 'nmos_devname', 'physic_nic_port0_IP', 'physic_nic_port1_IP', '2110-7_m_local_ip', '2110-7_b_local_ip'])

const ipData = unwrap(switchData.input, 'in_')
const input = ref<SwitchInputParamsType>(def_switch_input_params())
input.value = ipData
const inputBus = ref<SwitchInputBusParamsType>(def_switch_input_bus_params())
inputBus.value = unwrap(switchData.bus, 'bus_in_')
const inputs = ref<any[]>([]);

watch(() => mv.value.input_number, (nv) => {
  const len = inputs.value.length
  const params = ipData.input_video_params
  if (nv < len) {
    inputs.value = inputs.value.slice(0, nv);
  } else if (nv > len) {
    inputs.value = inputs.value.concat(
      Array.from({ length: nv - len }, (_, i) => {
        return {
          index: len + i + 1,
          value: params[len + i] ? ref(params[len + i]) : ref(def_switch_input_video_params())
        }
      })
    );
  }
}, { immediate: true });

const opData= unwrap(switchData.output, 'out_')
const output = ref<any>({ ...global_config })
output.value = pick(opData, Object.keys(global_config))
const outputs = Array.from({ length: 3 }, (_, i) => {
  if (i === 0) {
    return opData.pgm_params ? ref(opData.pgm_params) : ref(def_switch_output_params())
  } else if (i === 1) {
    return opData.pvw_params ? ref(opData.pvw_params) : ref(def_switch_output_params())
  } else if (i === 2) {
    return opData.clean_params ? ref(opData.clean_params) : ref(def_switch_output_params())
  }
});
const OUT_3_OPEN = ref(false)
OUT_3_OPEN.value = opData.clean_params.clean_is_open


function getValue() {
  const mip = mv.value['2110-7_m_local_ip']
  const bip = mv.value['2110-7_b_local_ip']
  const useb = output.value['g_2022-7']
  return {
    ...mv.value,
    input: {
      ...wrap(input.value, 'in_', input.value['g_2022-7']),
      input_params: inputs.value.map(ipt => wrap(ipt.value, 'in_', input.value['g_2022-7'], true, mip, bip))
    },
    bus: wrap(inputBus.value, 'bus_in_', input.value['g_2022-7']),
    output: {
      ...wrap(output.value, 'out_'),
      out_pgm_params: wrap(outputs[0].value, 'out_', useb, false, mip, bip),
      out_pvw_params: wrap(outputs[1].value, 'out_', useb, false, mip, bip),
      out_clean_params: {
        out_clean_is_open: OUT_3_OPEN.value,
        ...(OUT_3_OPEN.value ? wrap(outputs[2].value, 'out_', useb, false, mip, bip) : {})
      },
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
                    v-model="mv.input_number"
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
                    v-model="mv.tallyserver_url"
                  />
                </VControl>
              </VField>
            </div>
            <div class="column is-6">
              <VField>
                <VLabel>IO通讯地址</VLabel>
                <VControl>
                  <VInput
                    v-model="mv.p4server_url"
                  />
                </VControl>
              </VField>
            </div>
            <div class="column is-6">
              <VField>
                <VLabel>硬件控制面板通讯地址</VLabel>
                <VControl>
                  <VInput
                    v-model="mv.hw_panel_url"
                  />
                </VControl>
              </VField>
            </div>
            <div class="column is-6">
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
                <VLabel>拉流物理网卡端口0IP</VLabel>
                <VControl>
                  <VInput
                    v-model="mv.physic_nic_port0_IP"
                  />
                </VControl>
              </VField>
            </div>
            <div class="column is-6">
              <VField>
                <VLabel>拉流物理网卡端口1IP</VLabel>
                <VControl>
                  <VInput
                    v-model="mv.physic_nic_port1_IP"
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
                <h4>输入信源参数</h4>
              </div>
            </div>
          </div>
          <div class="form-body">
            <!--Fieldset-->
            <SwitchInputGlobal
              v-model="input"
            />
            <SwitchInput
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
                <h4>内部母线输入参数</h4>
              </div>
            </div>
          </div>
          <div class="form-body">
            <!--Fieldset-->
            <SwitchInputBus
              v-model="inputBus"
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
                  <VField horizontal>
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
            <SwitchOutput
              v-model="outputs[0]"
              title="PGM参数"
              :use-backup="output['g_2022-7']"
              :m_local_ip="mv['2110-7_m_local_ip']"
              :b_local_ip="mv['2110-7_b_local_ip']"
            />
            <SwitchOutput
              v-model="outputs[1]"
              title="PVW参数"
              :use-backup="output['g_2022-7']"
              :m_local_ip="mv['2110-7_m_local_ip']"
              :b_local_ip="mv['2110-7_b_local_ip']"
            />
            <SwitchOutput
              v-model="outputs[2]"
              v-model:OPEN="OUT_3_OPEN"
              title="Clean参数"
              toggle-title="是否启用Clean输出"
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

.form-layout {
  margin: 0 auto;

  .form-fieldset {
    max-width: none;
  }
}
</style>
