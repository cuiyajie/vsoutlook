<script setup lang="ts">
import { def_endswitch_input_params, def_endswitch_output_params, def_endswitch_input_video_params, type EndSwitchInputParamsType, global_config, type NMosConfigType, type SSMAddressType, nmos_config, ssm_address } from './Consts';
import { handle, unwrap, watchInput, watchNmosName, wrap } from "./Utils";
import pick from 'lodash-es/pick'
import merge from 'lodash-es/merge'
import endSwitchData from '@src/data/vscomponent/endswitch.json'
import { useUserSession } from "@src/stores/userSession";

const props = defineProps<{
  name: string
}>()

const mv = defineModel<{
  moudle: string,
  input_number: number,
  tallyserver_url: string,
  panel_url: string,
  panel_type: string,
  "2110-7_m_local_ip": string,
  "2110-7_b_local_ip": string,
  api_server_port: number,
  nmos: NMosConfigType,
  ssm_address_range: SSMAddressType[],
}>({
  default: {
    moudle: "switch",
    input_number: 4,
    tallyserver_url: "",
    panel_url: "",
    panel_type: "",
    "2110-7_m_local_ip": "",
    "2110-7_b_local_ip": "",
    api_server_port: 7001,
    nmos: { ...nmos_config },
    ssm_address_range: [{ ...ssm_address }],
  },
  local: true,
});
const advanceOpened = ref(false)

mv.value = pick(endSwitchData, ['moudle', 'input_number', 'tallyserver_url', 'panel_url', 'panel_type', '2110-7_m_local_ip', '2110-7_b_local_ip', 'api_server_port', 'nmos', 'ssm_address_range', 'authorization_service'])

watchNmosName(() => props.name, mv)

const usStore = useUserSession();
const panelTypes = computed(() => usStore.settings.endswt_panel_types || []);

const ipData = unwrap(endSwitchData.input, 'in_')
const input = ref<EndSwitchInputParamsType>(def_endswitch_input_params())
input.value = ipData
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
          value: params[len + i] ? ref(params[len + i]) : ref(def_endswitch_input_video_params(len + i))
        }
      })
    );
  }
}, { immediate: true });

const opData= unwrap(endSwitchData.output, 'out_')
const output = ref<any>({ ...global_config, number: 1 })
output.value = pick(opData, Object.keys(global_config).concat('number'))
const outputs = ref<any[]>([])

watch(() => output.value.number, (nv) => {
  const len = outputs.value.length
  const params = opData.params
  if (nv < len) {
    outputs.value = outputs.value.slice(0, nv);
  } else if (nv > len) {
    outputs.value = outputs.value.concat(
      Array.from({ length: nv - len }, (_, i) => {
        return {
          index: len + i + 1,
          value: params[len + i] ? ref(params[len + i]) : ref(def_endswitch_output_params())
        }
      })
    );
  }
}, { immediate: true });

watchInput('audioformat', input, outputs.value.map(o => o), { deep: true })

function getValue() {
  const mip = mv.value['2110-7_m_local_ip']
  const bip = mv.value['2110-7_b_local_ip']
  const useb = output.value['g_2022-7']
  return {
    ...handle(mv.value),
    input: {
      ...wrap(input.value, 'in_', input.value['g_2022-7']),
      input_video_params: inputs.value.map((ipt, idx) => Object.assign(wrap(ipt.value, 'in_', input.value['g_2022-7'], true, mip, bip), { index: idx })),
    },
    output: {
      ...wrap(output.value, 'out_'),
      out_params: outputs.value.map((opt, idx) => Object.assign(wrap(opt.value, 'out_', useb, false, mip, bip), { index: idx })),
    }
  }
}

function setValue(data: typeof endSwitchData) {
  mv.value = pick(data, ['moudle', 'input_number', 'tallyserver_url', 'panel_url', 'panel_type', '2110-7_m_local_ip', '2110-7_b_local_ip', 'api_server_port', 'nmos', 'ssm_address_range', 'authorization_service'])
  const _ipData = unwrap(data.input, 'in_')
  input.value = merge(def_endswitch_input_params(), _ipData)
  nextTick(() => {
    const params = _ipData.input_video_params
    inputs.value.forEach((iptv, idx) => {
      iptv.value = merge(def_endswitch_input_video_params(0), params[idx])
    })
  })

  const _opData = unwrap(data.output, 'out_')
  output.value = pick(_opData, Object.keys(global_config).concat('number'))
  nextTick(() => {
    const params = _opData.params
    outputs.value.forEach((optv, idx) => {
      optv.value = merge(def_endswitch_output_params(), params[idx])
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
                <VLabel>面板通讯地址</VLabel>
                <VControl>
                  <VInput
                    v-model="mv.panel_url"
                  />
                </VControl>
              </VField>
            </div>
            <div class="column is-6">
              <VField>
                <VLabel>面板类型</VLabel>
                <VControl>
                  <VSelect
                    v-model="mv.panel_type"
                    class="is-rounded"
                  >
                    <VOption
                      v-for="pt in panelTypes"
                      :key="pt.type"
                      :value="pt.type"
                    >
                      {{ pt.type }}
                    </VOption>
                  </VSelect>
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
            <div class="column is-6">
              <VField>
                <VLabel>对外端口</VLabel>
                <VControl>
                  <VInputNumber
                    v-model="mv.api_server_port"
                    centered
                    :min="0"
                    :step="1"
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
              <SSMAddressRange v-model="mv.ssm_address_range" />
            </div>
          </Transition>
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
            <EndSwitchInputGlobal v-model="input" />
            <EndSwitchInput
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
                <div class="column is-6">
                  <VField>
                    <VLabel>输出信号数量</VLabel>
                    <VControl>
                      <VInputNumber
                        v-model="output.number"
                        centered
                        :min="0"
                        :step="1"
                      />
                    </VControl>
                  </VField>
                </div>
              </div>
            </div>
            <EndSwitchOutput
              v-for="(opt, idx) in outputs"
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
@import "@src/scss/abstracts/all";
@import "@src/scss/components/forms-outer";

.form-layout {
  margin: 0 auto;

  .form-fieldset {
    max-width: none;
  }
}
</style>
