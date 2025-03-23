<script lang="ts" setup>
import {
  type NMosConfigType,
  nmos_config,
  ssm_address,
  type SSMAddressType,
} from '../Utilties/Consts'
import {
  type IApiParams,
  used_signal_types,
  type NicDetail,
  def_player_params,
} from '../Utilties/Consts_V1'
import { handle, watchNmosName } from '../Utilties/Utils';
import pick from 'lodash-es/pick'
import merge from 'lodash-es/merge'
import mvData from '@src/data/vscomponent/mv.json'
import { useUsedFormat } from '../Utilties/Composables';
import { useUserSession } from "@src/stores/userSession"
import { handleVideoFormat, handleAudioFormat, handlePlayerParams, handleNicList, unwrap, wrap } from '../Utilties/Utils_V1';
import { def_mv_input_param, def_mv_output_param } from "./Consts";

const props = defineProps<{
  name: string,
  nics: NicInfo[]
}>()

const mv = defineModel<{
  moudle: string
  av_log_level: number
  input_number: number
  output_number: number
  used_signal_type: number
  nmos: NMosConfigType
  ssm_address_range: SSMAddressType[]
  api_params: IApiParams[],
}>({
  default: {
    moudle: 'mv',
    av_log_level: 5,
    input_number: 4,
    output_number: 1,
    used_signal_type: 0,
    nmos: { ...nmos_config },
    ssm_address_range: [{ ...ssm_address }],
    api_params: [],
  },
  local: true,
})
const advanceOpened = ref(false)

const usStore = useUserSession()
const videoFormats = computed(() => usStore.settings.video_formats || [])
const audioFormats = computed(() => usStore.settings.audio_formats || [])

mv.value = pick(mvData, [
  'moudle',
  'input_number',
  'output_number',
  'av_log_level',
  'used_signal_type',
  'nmos',
  'ssm_address_range',
  'authorization_service',
  'api_params',
])

const nicDetails = ref<NicDetail[]>([])
const indexedNicDetails = computed(() => {
  return nicDetails.value
    .filter(nic => nic.nicIndex >= 0)
    .map((nic, idx) => ({ ...nic, index: idx }))
})

const templates = computed(() => usStore.settings.mv_template_list || [])
function mvTemplateIsValid(tmpl: string) {
  return !!templates.value.find((v) => v.path === tmpl)
}


const ipData = unwrap(mvData.input, 'in_')
const input = ref<{ 'g_2022-7': boolean }>({ 'g_2022-7': false })
input.value = pick(ipData, ['g_2022-7'])

const inputs = ref<any[]>([])
watch(
  () => mv.value.input_number,
  (nv) => {
    const len = inputs.value.length
    if (nv < len) {
      inputs.value = inputs.value.slice(0, nv)
    } else if (nv > len) {
      inputs.value = inputs.value.concat(
        Array.from({ length: nv - len }, (_, i) => {
          return {
            index: len + i + 1,
            value: ipData.input_params[len + i]
              ? ref(ipData.input_params[len + i])
              : ref(def_mv_input_param()),
          }
        })
      )
    }
  },
  { immediate: true }
)

const opData = unwrap(mvData.output, 'out_')
const outputs = ref<any[]>([])
const outputRefs = ref<any[]>([])
watch(
  () => mv.value.output_number,
  (nv) => {
    const len = outputs.value.length
    if (nv < len) {
      outputs.value = outputs.value.slice(0, nv)
    } else if (nv > len) {
      outputs.value = outputs.value.concat(
        Array.from({ length: nv - len }, (_, i) => {
          return {
            index: len + i + 1,
            value: opData.params[len + i]
              ? ref(opData.params[len + i])
              : ref(def_mv_output_param()),
          }
        })
      )
    }
    outputs.value.forEach((v) => {
      if (v.value.mv_template && !mvTemplateIsValid(v.value.mv_template)) {
        v.value.mv_template = ''
      }
    })
  },
  { immediate: true }
)

const [videoFormatEnum, videoSelected, videoUnSelected] = useUsedFormat()
const [audioFormatEnum, audioSelected, audioUnSelected] = useUsedFormat()
const [auditAvTemplateEnum, auditAvSelected, auditAvUnSelected] = useUsedFormat()

watchNmosName(() => props.name, mv)

function getValue() {
  const result = {
    ...handle(mv.value),
    videoformat_enum: videoFormatEnum.value.map(vfn => handleVideoFormat(vfn, videoFormats.value)).filter(v => v),
    audioformat_enum: audioFormatEnum.value.map(afn => handleAudioFormat(afn, audioFormats.value)).filter(a => a),
    player_params: wrap(handlePlayerParams(playerParams.value, videoFormats.value), 'out_'),
  }
  if (mv.value.used_signal_type !== 1) {
    result.nic_list = handleNicList(indexedNicDetails.value)
  }
  return result
}

function setValue(data: typeof mvData) {
  mv.value = pick(data, [
    'moudle',
    'input_number',
    'output_number',
    'av_log_level',
    'used_signal_type',
    'nmos',
    'ssm_address_range',
    'authorization_service',
    'api_params',
  ])
}

defineExpose({
  getValue,
  setValue,
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
                <VLabel>需要使用的信号类型</VLabel>
                <VControl>
                  <VSelect
                    v-model="mv.used_signal_type"
                    class="is-rounded"
                  >
                    <VOption v-for="ust in used_signal_types" :key="ust.key" :value="ust.key">{{ ust.label }}</VOption>
                  </VSelect>
                </VControl>
              </VField>
            </div>
          </div>
        </div>
        <div class="form-fieldset">
          <div class="fieldset-heading">
            <h4>接口设置</h4>
          </div>
          <div class="columns is-multiline">
            <div class="column is-6">
              <AddrInputAddonSep
                v-model:host="mv.api_params[0].api_name"
                v-model:port="mv.api_params[0].service_port"
                label="tally通知接口"
              />
            </div>
            <div class="column is-6">
              <AddrInputAddonSep
                v-model:host="mv.api_params[1].api_name"
                v-model:port="mv.api_params[1].service_port"
                label="动态控制接口"
              />
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
        <Transition name="fade-slow">
          <NicSection v-if="mv.used_signal_type !== 1" v-model="nicDetails" class="has-mb-20" :nics="nics" />
        </Transition>
        <div class="form-outer">
          <div class="form-header">
            <div class="form-header-inner">
              <div class="left">
                <h4>使用的视频格式列表</h4>
              </div>
            </div>
          </div>
          <div class="form-body">
            <VTags v-if="videoFormatEnum.length > 0" class="format-item-container">
              <VideoFormatTooltip v-for="vformat in videoFormatEnum" :key="vformat" :name="vformat">
                <VTag color="blue" :label="vformat" curved outlined />
              </VideoFormatTooltip>
            </VTags>
            <div v-else class="is-empty-list">暂时还没有选择视频格式</div>
          </div>
        </div>
        <div class="form-outer">
          <div class="form-header">
            <div class="form-header-inner">
              <div class="left">
                <h4>使用的音频格式列表</h4>
              </div>
            </div>
          </div>
          <div class="form-body">
            <VTags v-if="audioFormatEnum.length > 0" class="format-item-container">
              <AudioFormatTooltip v-for="aformat in audioFormatEnum" :key="aformat" :name="aformat">
                <VTag color="green" :label="aformat" curved outlined />
              </AudioFormatTooltip>
            </VTags>
            <div v-else class="is-empty-list">暂时还没有选择音频格式</div>
          </div>
        </div>
      </div>
    </div>
  </form>
</template>
<style lang="scss"></style>
