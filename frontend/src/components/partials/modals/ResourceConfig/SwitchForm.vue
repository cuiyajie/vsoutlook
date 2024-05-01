<script setup lang="ts">
import { def_switch_input_params, def_switch_output_params, def_switch_input_video_params, type SwitchInputParamsType, global_config, def_switch_bus_params, type SwitchBusParamsType, type NMosConfigType, type SSMAddressType, nmos_config, ssm_address, type SwitchKeyType, def_switch_input_key_params, def_switch_input_fill_params, def_switch_bus_keyfill_params, def_switch_key_params, def_switch_key } from './Consts';
import { handle, unwrap, watchNmosName, wrap } from "./Utils";
import pick from 'lodash-es/pick'
import merge from 'lodash-es/merge'
import switchData from '@src/data/vscomponent/switch.json'

const props = defineProps<{
  name: string
}>()

const mv = defineModel<{
  moudle: string,
  input_number: number,
  tallyserver_url: string,
  p4server_url: string,
  hw_panel_url: string,
  physic_nic_port0_ip: string,
  physic_nic_port1_ip: string,
  "2110-7_m_local_ip": string,
  "2110-7_b_local_ip": string,
  nmos: NMosConfigType,
  ssm_address_range: SSMAddressType[],
}>({
  default: {
    moudle: "switch",
    input_number: 10,
    tallyserver_url: "",
    p4server_url: "",
    hw_panel_url: "",
    physic_nic_port0_ip: "",
    physic_nic_port1_ip: "",
    "2110-7_m_local_ip": "",
    "2110-7_b_local_ip": "",
    nmos: { ...nmos_config },
    ssm_address_range: [{ ...ssm_address }],
  },
  local: true,
});
const advanceOpened = ref(false)

mv.value = pick(switchData, ['moudle', 'input_number', 'tallyserver_url', 'p4server_url', 'hw_panel_url', 'physic_nic_port0_ip', 'physic_nic_port1_ip', '2110-7_m_local_ip', '2110-7_b_local_ip', 'nmos', 'ssm_address_range', 'authorization_service'])

watchNmosName(() => props.name, mv)

const keyData = switchData.key
const key = ref<SwitchKeyType>({ ...keyData })
const keyParams = ref<any[]>([])

watch(() => key.value.key_number, nv => {
  const len = keyParams.value.length
  const params = keyData.key_params
  if (nv < len) {
    keyParams.value = keyParams.value.slice(0, nv);
  } else if (nv > len) {
    const defKeyParams = def_switch_key_params()
    keyParams.value = keyParams.value.concat(
      Array.from({ length: nv - len }, (_, i) => {
        const idx = len + i
        const defKeyParam = idx === 1 ? defKeyParams[1] : defKeyParams[0]
        return {
          index: idx + 1,
          value: params[idx] ? ref(params[idx]) : ref({ ...defKeyParam })
        }
      })
    );
  }

}, { immediate: true })

const ipData = unwrap(switchData.input, 'in_')
const input = ref<SwitchInputParamsType>(def_switch_input_params())
input.value = ipData
const inputs = ref<any[]>([]);
const inputKeys = ref<any[]>([]);
const inputFills = ref<any[]>([]);

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
          value: params[len + i] ? ref(params[len + i]) : ref(def_switch_input_video_params(len + i))
        }
      })
    );
  }
}, { immediate: true });

const busData = unwrap(switchData.bus, 'bus_in_')
const bus = ref<SwitchBusParamsType>(def_switch_bus_params())
bus.value = busData
const masterBusKeyFills = ref<any[]>([]);
const backupBusKeyFills = ref<any[]>([]);

watch(() => key.value.ext_key_number, (nv) => {
  const len = inputKeys.value.length
  const ikParams = ipData.input_key_params
  const ifParams = ipData.input_fill_params
  const kbmParams = busData.keyfill_bus_master
  const kbbParams = busData.keyfill_bus_backup
  if (nv < len) {
    inputKeys.value = inputKeys.value.slice(0, nv);
    inputFills.value = inputFills.value.slice(0, nv);
    masterBusKeyFills.value = masterBusKeyFills.value.slice(0, nv);
    backupBusKeyFills.value = backupBusKeyFills.value.slice(0, nv);
  } else if (nv > len) {
    inputKeys.value = inputKeys.value.concat(
      Array.from({ length: nv - len }, (_, i) => {
        return {
          index: len + i + 1,
          value: ikParams[len + i] ? ref(ikParams[len + i]) : ref(def_switch_input_key_params())
        }
      })
    );
    inputFills.value = inputFills.value.concat(
      Array.from({ length: nv - len }, (_, i) => {
        return {
          index: len + i + 1,
          value: ifParams[len + i] ? ref(ifParams[len + i]) : ref(def_switch_input_fill_params())
        }
      })
    );
    masterBusKeyFills.value = masterBusKeyFills.value.concat(
      Array.from({ length: nv - len }, (_, i) => {
        return {
          index: len + i + 1,
          value: kbmParams[len + i] ? ref(kbmParams[len + i]) : ref(def_switch_bus_keyfill_params())
        }
      })
    );
    backupBusKeyFills.value = backupBusKeyFills.value.concat(
      Array.from({ length: nv - len }, (_, i) => {
        return {
          index: len + i + 1,
          value: kbbParams[len + i] ? ref(kbbParams[len + i]) : ref(def_switch_bus_keyfill_params())
        }
      })
    );
  }
}, { immediate: true });

const opData= unwrap(switchData.output, 'out_')
const output = ref<any>({ ...global_config })
output.value = pick(opData, Object.keys(global_config))
const outputs = ref<any>([])
outputs.value = Array.from({ length: 3 }, (_, i) => {
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
    ...handle(mv.value),
    input: {
      ...wrap(input.value, 'in_', input.value['g_2022-7']),
      input_video_params: inputs.value.map((ipt, idx) => Object.assign(wrap(ipt.value, 'in_', input.value['g_2022-7'], true, mip, bip), { index: idx })),
      input_key_params: inputKeys.value.map((ipt, idx) => Object.assign(wrap(ipt.value, 'in_', input.value['g_2022-7'], true, mip, bip), { index: idx })),
      input_fill_params: inputFills.value.map((ipt, idx) => Object.assign(wrap(ipt.value, 'in_', input.value['g_2022-7'], true, mip, bip), { index: idx }))
    },
    key: {
      ...key.value,
      key_params: keyParams.value.map((ipt, index) => Object.assign(ipt.value, { index }))
    },
    bus: {
      ...wrap(bus.value, 'bus_in_', input.value['g_2022-7']),
      keyfill_bus_master: masterBusKeyFills.value.map((ipt, idx) => Object.assign(wrap(ipt.value, 'bus_in_', input.value['g_2022-7']), { index: idx })),
      keyfill_bus_backup: backupBusKeyFills.value.map((ipt, idx) => Object.assign(wrap(ipt.value, 'bus_in_', input.value['g_2022-7']), { index: idx }))
    },
    output: {
      ...wrap(output.value, 'out_'),
      out_pgm_params: wrap(unref(outputs.value[0]), 'out_', useb, false, mip, bip),
      out_pvw_params: wrap(unref(outputs.value[1]), 'out_', useb, false, mip, bip),
      out_clean_params: {
        out_clean_is_open: OUT_3_OPEN.value,
        ...(OUT_3_OPEN.value ? wrap(unref(outputs.value[2]), 'out_', useb, false, mip, bip) : {})
      },
    }
  }
}

function setValue(data: typeof switchData) {
  mv.value = pick(data, ['moudle', 'input_number', 'tallyserver_url', 'p4server_url', 'hw_panel_url', 'physic_nic_port0_ip', 'physic_nic_port1_ip', '2110-7_m_local_ip', '2110-7_b_local_ip', 'nmos', 'ssm_address_range', 'authorization_service'])
  const _ipData = unwrap(data.input, 'in_')
  key.value = merge(def_switch_key(), data.key)
  input.value = merge(def_switch_input_params(), _ipData)
  bus.value = merge(def_switch_bus_params(), unwrap(data.bus, 'bus_in_'))
  nextTick(() => {
    const params = _ipData.input_video_params
    inputs.value.forEach((iptv, idx) => {
      iptv.value = merge(def_switch_input_video_params(0), params[idx])
    })
    const ikParams = _ipData.input_key_params
    inputKeys.value.forEach((ipt, idx) => {
      ipt.value = merge(def_switch_input_key_params(), ikParams[idx])
    })
    const ifParams = _ipData.input_fill_params
    inputFills.value.forEach((ipt, idx) => {
      ipt.value = merge(def_switch_input_fill_params(), ifParams[idx])
    })
    const kbmParams = bus.value.keyfill_bus_master
    masterBusKeyFills.value.forEach((ipt, idx) => {
      ipt.value = merge(def_switch_bus_keyfill_params(), kbmParams[idx])
    })
    const kbbParams = bus.value.keyfill_bus_backup
    backupBusKeyFills.value.forEach((ipt, idx) => {
      ipt.value = merge(def_switch_bus_keyfill_params(), kbbParams[idx])
    })
  })

  const _opData = unwrap(data.output, 'out_')
  output.value = pick(_opData, Object.keys(global_config))
  outputs.value[0].value = merge(def_switch_output_params(), _opData.pgm_params)
  outputs.value[1].value = merge(def_switch_output_params(), _opData.pvw_params)
  OUT_3_OPEN.value = _opData.clean_params.clean_is_open
  outputs.value[2].value = merge(def_switch_output_params(), _opData.clean_params)
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
                <VLabel>Tally服务地址（含端口）</VLabel>
                <VControl>
                  <VInput
                    v-model="mv.tallyserver_url"
                  />
                </VControl>
              </VField>
            </div>
            <div class="column is-6">
              <VField>
                <VLabel>前置SDN通讯地址（含端口）</VLabel>
                <VControl>
                  <VInput
                    v-model="mv.p4server_url"
                  />
                </VControl>
              </VField>
            </div>
            <div class="column is-6">
              <VField>
                <VLabel>硬件控制面板通讯地址（含端口）</VLabel>
                <VControl>
                  <VInput
                    v-model="mv.hw_panel_url"
                  />
                </VControl>
              </VField>
            </div>
            <div class="column is-6">
              <VField>
                <VLabel>拉流物理网卡端口0 - IP</VLabel>
                <VControl>
                  <VInput
                    v-model="mv.physic_nic_port0_ip"
                  />
                </VControl>
              </VField>
            </div>
            <div class="column is-6">
              <VField>
                <VLabel>拉流物理网卡端口1 - IP</VLabel>
                <VControl>
                  <VInput
                    v-model="mv.physic_nic_port1_ip"
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
        <div class="form-outer">
          <div class="form-header">
            <div class="form-header-inner">
              <div class="left">
                <h4>键设置</h4>
              </div>
            </div>
          </div>
          <div class="form-body">
            <div class="form-fieldset">
              <div class="columns is-multiline">
                <div class="column is-6">
                  <VField>
                    <VLabel>键数量</VLabel>
                    <VControl>
                      <VInputNumber
                        v-model="key.key_number"
                        centered
                        :min="0"
                        :step="1"
                      />
                    </VControl>
                  </VField>
                </div>
                <div class="column is-6">
                  <VField>
                    <VLabel>外部输入键信号数量</VLabel>
                    <VControl>
                      <VInputNumber
                        v-model="key.ext_key_number"
                        centered
                        :min="0"
                        :step="1"
                      />
                    </VControl>
                  </VField>
                </div>
              </div>
            </div>
            <SwitchKeyParams
              v-for="(ipt, idx) in keyParams"
              :key="ipt.index"
              v-model="ipt.value"
              :index="ipt.index"
              :is-last="idx === keyParams.length - 1"
            />
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
            <div class="form-fieldset form-outer">
              <div class="fieldset-heading">
                <h4>输入key参数</h4>
              </div>
              <div class="form-fieldset-nested-2">
                <SwitchInputKey
                  v-for="(ipt, idx) in inputKeys"
                  :key="ipt.index"
                  v-model="ipt.value"
                  :index="ipt.index"
                  :is-last="idx === inputKeys.length - 1"
                  :use-backup="input['g_2022-7']"
                />
              </div>
            </div>
            <div class="form-fieldset form-outer">
              <div class="fieldset-heading">
                <h4>输入fill参数</h4>
              </div>
              <div class="form-fieldset-nested-2">
                <SwitchInputFill
                  v-for="(ipt, idx) in inputFills"
                  :key="ipt.index"
                  v-model="ipt.value"
                  :index="ipt.index"
                  :is-last="idx === inputFills.length - 1"
                  :use-backup="input['g_2022-7']"
                />
              </div>
            </div>
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
          <div class="form-body is-nested-2">
            <div class="form-fieldset form-outer">
              <div class="fieldset-heading">
                <h4>视频输入母线</h4>
              </div>
              <div class="form-fieldset-nested-2">
                <!--Fieldset-->
                <SwitchBus
                  v-model="bus"
                  :use-backup="input['g_2022-7']"
                />
              </div>
            </div>
            <div class="form-fieldset form-outer">
              <div class="fieldset-heading">
                <h4>键输入母线</h4>
              </div>
              <div class="form-fieldset-nested-2">
                <div class="seperator">
                  <SwitchBusKeyFill
                    v-for="(ipt, idx) in masterBusKeyFills"
                    :key="`master_${ipt.index}`"
                    v-model="ipt.value"
                    name="主"
                    :index="ipt.index"
                    :is-last="idx === masterBusKeyFills.length - 1"
                  />
                </div>
                <Transition name="fade-slow">
                  <div v-if="input['g_2022-7']">
                    <SwitchBusKeyFill
                      v-for="(ipt, idx) in backupBusKeyFills"
                      :key="`backup_${ipt.index}`"
                      v-model="ipt.value"
                      name="备"
                      :index="ipt.index"
                      :is-last="idx === backupBusKeyFills.length - 1"
                    />
                  </div>
                </Transition>
              </div>
            </div>
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
@import "@src/scss/abstracts/all";
@import "@src/scss/components/forms-outer";

.form-layout {
  margin: 0 auto;

  .form-fieldset {
    max-width: none;
  }
}
</style>
