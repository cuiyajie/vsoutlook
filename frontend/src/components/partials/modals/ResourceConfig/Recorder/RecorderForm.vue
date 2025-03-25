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
import rcData from '@src/data/vscomponent/recorder.json'
import { useUsedFormat } from '../Utilties/Composables';
import { useUserSession } from "@src/stores/userSession"
import { handleVideoFormat, handleAudioFormat, handlePlayerParams, handleNicList, unwrap, wrap, checkPlayerParams } from '../Utilties/Utils_V1';

const props = defineProps<{
  name: string,
  nics: NicInfo[]
}>()

const mv = defineModel<{
  moudle: string
  av_log_level: number
  used_signal_type: number
  nmos: NMosConfigType
  ssm_address_range: SSMAddressType[]
  api_params: IApiParams[],
  recoder_params: { in_nic_index: number },
}>({
  default: {
    moudle: 'recorder',
    av_log_level: 5,
    used_signal_type: 0,
    nmos: { ...nmos_config },
    ssm_address_range: [{ ...ssm_address }],
    api_params: [],
    recoder_params: { in_nic_index: -1 },
  },
  local: true,
})
const advanceOpened = ref(false)

const usStore = useUserSession()
const videoFormats = computed(() => usStore.settings.video_formats || [])
const audioFormats = computed(() => usStore.settings.audio_formats || [])

mv.value = pick(rcData, [
  'moudle',
  'av_log_level',
  'used_signal_type',
  'nmos',
  'ssm_address_range',
  'authorization_service',
  'api_params',
  'recoder_params'
])

const nicDetails = ref<NicDetail[]>([])
const indexedNicDetails = computed(() => {
  return nicDetails.value
    .filter(nic => nic.nicIndex >= 0)
    .map((nic, idx) => ({ ...nic, index: idx }))
})

const playerParams = ref(def_player_params())

const [videoFormatEnum, videoSelected, videoUnSelected] = useUsedFormat()
const [audioFormatEnum, audioSelected, audioUnSelected] = useUsedFormat()

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

function setValue(data: typeof rcData) {
  mv.value = pick(data, [
    'moudle',
    'av_log_level',
    'used_signal_type',
    'nmos',
    'ssm_address_range',
    'authorization_service',
    'api_params',
    'recoder_params'
  ])
  const _playerParams = unwrap(data.player_params, 'out_')
  playerParams.value = checkPlayerParams(merge(def_player_params(), _playerParams), videoFormats.value, audioFormats.value)
  nicDetails.value = data.nic_list.map((nic: any) => {
    const nicIndex = props.nics.findIndex(n => n.nicNameMain === nic.nic_name_m && n.nicNameBackup === nic.nic_name_b)
    if (nicIndex !== -1) {
      return { ...nic, nicIndex, id: props.nics[nicIndex].id }
    }
    return { ...nic, nicIndex: -1, id: '' }
  })
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
                label="录制控制接口"
              />
            </div>
            <div class="column is-6">
              <AddrInputAddonSep
                v-model:host="mv.api_params[1].api_name"
                v-model:port="mv.api_params[1].service_port"
                label="播放控制接口"
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
        <div class="form-fieldset">
          <div class="fieldset-heading">
            <h4>录制参数</h4>
          </div>
          <div class="columns is-multiline">
            <div class="column is-6">
              <VField>
                <VLabel>smpte收流网卡序号</VLabel>
                <VControl>
                  <NicDetailSelect v-model="mv.recoder_params.in_nic_index" :nics="indexedNicDetails" />
                </VControl>
              </VField>
            </div>
          </div>
        </div>
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
        <div class="form-outer">
          <div class="form-header">
            <div class="form-header-inner">
              <div class="left">
                <h4>播放参数</h4>
              </div>
            </div>
          </div>
          <div class="form-body">
            <!--Fieldset-->
            <div class="form-fieldset seperator">
              <div class="columns is-multiline">
                <div class="column is-6">
                  <VField>
                    <VLabel>输出视频格式名称</VLabel>
                    <VControl>
                      <VideoFormatSelect v-model="playerParams.videoformat_name" :used-signal-type="mv.used_signal_type" @video-selected="videoSelected" @video-unselected="videoUnSelected" />
                    </VControl>
                  </VField>
                </div>
                <div class="column is-6">
                  <VField>
                    <VLabel>输出音频格式名称</VLabel>
                    <VControl>
                      <AudioFormatSelect v-model="playerParams.audioformat_name" @audio-selected="audioSelected" @audio-unselected="audioUnSelected" />
                    </VControl>
                  </VField>
                </div>
              </div>
            </div>
            <PlayerParamsForm v-model="playerParams" :nics="indexedNicDetails" />
          </div>
        </div>
      </div>
    </div>
  </form>
</template>
<style lang="scss"></style>
