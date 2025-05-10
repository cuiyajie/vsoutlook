<script lang="ts" setup>
import {
  type NMosConfigType,
  nmos_config,
  ssm_address,
  type SSMAddressType,
} from '../Utilties/Consts'
import {
  type IApiParams,
} from '../Utilties/Consts_V1'
import { handle, watchNmosName } from '../Utilties/Utils';
import { def_api_params } from './Consts';
import pick from 'lodash-es/pick'
import bcswtData from '@src/data/vscomponent/bcswitch.json'
import { useUserSession } from "@src/stores/userSession"
import { useUsedFormat, useNicList, useSwitchPanel } from '../Utilties/Composables';
import { def_switch_bus, def_switch_input, def_switch_out, def_switch_panel } from '../SwitchShare/Consts';
import { checkApiParams, checkNicDetails, handleApiParams, handleAudioFormat, handleAudioMapping, handleNicList, handleVideoFormat, wrap } from '../Utilties/Utils_V1';
import { checkSwitchData, handleSwitchBus, handleSwitchInput, handleSwitchOut, handleSwitchPanel } from '../SwitchShare/Utils';

const props = defineProps<{
  name: string,
  nics: NicInfo[],
  requiredment?: TmplRequirement
}>()

const mv = defineModel<{
  moudle: string
  av_log_level: number
  gpu_index: number
  in_copy_thr_count: number
  used_signal_type: number
  level: number
  audio_workmode: number
  tally_notify: boolean
  nmos: NMosConfigType
  ssm_address_range: SSMAddressType[]
}>({
  default: {
    moudle: 'switch-v2',
    av_log_level: 5,
    gpu_index: 0,
    in_copy_thr_count: 8,
    used_signal_type: 0,
    level: 1,
    audio_workmode: 1,
    tally_notify: false,
    nmos: { ...nmos_config },
    ssm_address_range: [{ ...ssm_address }]
  },
  local: true,
})
const fullOpened = ref(true)
const advanceOpened = ref(false)

const usStore = useUserSession()
const videoFormats = computed(() => usStore.settings.video_formats || [])
const audioFormats = computed(() => usStore.settings.audio_formats || [])
const audioMappings = computed(() => usStore.settings.audio_mappings || [])

mv.value = pick(bcswtData, [
  'moudle',
  'av_log_level',
  'gpu_index',
  'used_signal_type',
  'in_copy_thr_count',
  'level',
  'audio_workmode',
  'tally_notify',
  'nmos',
  'ssm_address_range',
  'authorization_service',
])

const apiParams = ref<IApiParams[]>(def_api_params())

const { nicDetails, indexedNicDetails } = useNicList(props)

const [videoFormatEnum, videoSelected, videoUnSelected] = useUsedFormat()
const [audioFormatEnum, audioSelected, audioUnSelected] = useUsedFormat()
const [audioMappingEnum, audioMappingSelected, audioMappingUnSelected] = useUsedFormat()

provide('switch_type_category', 'bcswt')
provide('switch_template_toggle', {
  videoSelected,
  videoUnSelected,
  audioSelected,
  audioUnSelected,
  audioMappingSelected,
  audioMappingUnSelected
})
provide('switch_nics', indexedNicDetails)
provide('switch_used_signal_type', computed(() => mv.value.used_signal_type))
provide('switch_audio_mode', computed(() => mv.value.audio_workmode))

const panel = ref(def_switch_panel(mv.value.level))
useSwitchPanel(() => mv.value.level, panel)

const input = ref(def_switch_input())
const bus = ref(def_switch_bus(mv.value.level))
const out = ref(def_switch_out(mv.value.level))

watchNmosName(() => props.name, mv)

function getValue() {
  const result = {
    ...handle(mv.value),
    api_params: handleApiParams(apiParams.value),
    panel_params: handleSwitchPanel(panel.value),
    videoformat_enum: videoFormatEnum.value.map(vfn => handleVideoFormat(vfn, videoFormats.value)).filter(v => v),
    input: wrap(handleSwitchInput(input.value, videoFormats.value), 'in_'),
    bus: handleSwitchBus(bus.value, mv.value.used_signal_type, 'bcswt'),
    output: wrap(handleSwitchOut(out.value, videoFormats.value, mv.value.audio_workmode, mv.value.used_signal_type, 'bcswt'), 'out_'),
  }
  if (mv.value.audio_workmode !== 0) {
    result.audioformat_enum = audioFormatEnum.value.map(afn => handleAudioFormat(afn, audioFormats.value)).filter(a => a)
    result.audio_mapping_enum = audioMappingEnum.value.map(amfn => handleAudioMapping(amfn, audioMappings.value)).filter(a => a)
  }
  if (mv.value.used_signal_type !== 1) {
    result.nic_list = handleNicList(indexedNicDetails.value)
  }
  return result
}

function setValue(data: typeof bcswtData) {
  mv.value = pick(data, [
    'moudle',
    'av_log_level',
    'gpu_index',
    'used_signal_type',
    'in_copy_thr_count',
    'level',
    'audio_workmode',
    'tally_notify',
    'nmos',
    'ssm_address_range',
    'authorization_service',
  ])
  mv.value.nmos.rds_server_url = usStore.settings.rds_server_url || ''
  apiParams.value = checkApiParams(def_api_params(), data.api_params)
  nicDetails.value = checkNicDetails(data.nic_list, props.nics)
  const _data = checkSwitchData(data, videoFormats.value, audioFormats.value, audioMappings.value)
  panel.value = _data.panel
  input.value = _data.input
  bus.value = _data.bus
  out.value = _data.out
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
      <div class="form-header is-h1" :class="!fullOpened && 'border-b-none'">
        <div
          class="form-header-inner collapse-control-header"
          tabindex="-1"
          role="button"
          :open="fullOpened || undefined"
          @keydown.space.prevent="fullOpened = !fullOpened"
          @click.prevent="fullOpened = !fullOpened"
        >
          <div class="left">
            <h3>设备参数</h3>
          </div>
          <div class="collapse-icon">
            <VIcon icon="feather:chevron-down" />
          </div>
        </div>
      </div>
      <expand-transition>
        <div v-show="fullOpened" class="form-body is-nested">
          <!--Fieldset-->
          <div class="form-fieldset">
            <div class="fieldset-heading">
              <h4>通用参数</h4>
            </div>
            <div class="columns is-multiline">
              <div class="column is-6">
                <VField>
                  <VLabel>使用的显卡序号</VLabel>
                  <VControl>
                    <VInputNumber
                      v-model="mv.gpu_index"
                      centered
                      :min="0"
                      :step="1"
                    />
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
              <div v-for="(apiParam, apiIndex) in apiParams" :key="apiParam.api_name" class="column is-6">
                <ApiParamsCheck
                  v-model="apiParams[apiIndex]"
                  :label="apiParam.label"
                />
              </div>
            </div>
          </div>
          <div class="form-fieldset form-outer">
            <div class="fieldset-heading">
              <h4>操作面板设置</h4>
            </div>
            <div class="form-body">
              <SwitchPanel
                v-for="(panelConfig, pidx) in panel"
                :key="pidx"
                v-model="panel[pidx]"
                :index="pidx"
                :is-last="pidx === panel.length - 1"
              />
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
            <expand-transition>
              <div v-show="advanceOpened" class="form-body">
                <NMosConfig v-model="mv.nmos" class="seperator" />
                <SSMAddressRange v-model="mv.ssm_address_range" />
              </div>
            </expand-transition>
          </div>
          <expand-transition name="fade-slow">
            <NicSection v-if="mv.used_signal_type !== 1" v-model="nicDetails" :nics="nics" :max="requiredment?.nicCount || 0" />
          </expand-transition>
          <div class="form-outer has-mt-20">
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
          <div v-if="mv.audio_workmode !== 0" class="form-outer">
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
          <div v-if="mv.audio_workmode !== 0" class="form-outer">
            <div class="form-header">
              <div class="form-header-inner">
                <div class="left">
                  <h4>使用的音频声道映射模板列表</h4>
                </div>
              </div>
            </div>
            <div class="form-body">
              <VTags v-if="audioMappingEnum.length > 0" class="format-item-container">
                <AudioMappingTooltip v-for="aMapping in audioMappingEnum" :key="aMapping" :name="aMapping">
                  <VTag color="orange" :label="aMapping" curved outlined />
                </AudioMappingTooltip>
              </VTags>
              <div v-else class="is-empty-list">暂时还没有选择音频声道映射模板</div>
            </div>
          </div>
          <div class="form-outer">
            <div class="form-header">
              <div class="form-header-inner">
                <div class="left">
                  <h4>输入信源参数</h4>
                </div>
              </div>
            </div>
            <div class="form-body">
              <SwitchInput v-model="input" />
            </div>
          </div>
          <div class="form-outer">
            <div class="form-header">
              <div class="form-header-inner">
                <div class="left">
                  <h4>总线设置</h4>
                </div>
              </div>
            </div>
            <div class="form-body">
              <SwitchBus v-model="bus" :input-keys="input.key" :input-videos="input.video" :level="mv.level" />
            </div>
          </div>
          <div class="form-outer">
            <div class="form-header">
              <div class="form-header-inner">
                <div class="left">
                  <h4>输出设置</h4>
                </div>
              </div>
            </div>
            <div class="form-body">
              <SwitchOut v-model="out" :level="mv.level" :input-videos="input.video" :bus-levels="bus.level_bus" />
            </div>
          </div>
        </div>
      </expand-transition>
    </div>
  </form>
</template>
<style lang="scss"></style>
