<script lang="ts" setup>
import {
  type NMosConfigType,
  nmos_config,
  ssm_address,
  type SSMAddressType,
} from '../Utilties/Consts'
import {
  used_signal_types,
  type NicDetail,
} from '../Utilties/Consts_V1'
import {
  def_mg_player_params,
  def_mg_output_item_params,
} from './Consts'
import { handle, watchNmosName } from '../Utilties/Utils'
import pick from 'lodash-es/pick'
import merge from 'lodash-es/merge'
import mgData from '@src/data/vscomponent/media_gateway.json'
import { useUsedFormat } from '../Utilties/Composables';
import { useUserSession } from "@src/stores/userSession"
import { handleVideoFormat, handleAudioFormat, handleNicList, wrap, unwrap, checkPlayerParams } from '../Utilties/Utils_V1';
import { handleInputParams, handleOutputParams } from "./Utils"

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
}>({
  default: {
    moudle: 'media_gateway',
    av_log_level: 5,
    used_signal_type: 0,
    nmos: { ...nmos_config },
    ssm_address_range: [{ ...ssm_address }],
  },
  local: true,
})
const advanceOpened = ref(false)

const usStore = useUserSession()
const videoFormats = computed(() => usStore.settings.video_formats || [])
const audioFormats = computed(() => usStore.settings.audio_formats || [])

mv.value = pick(mgData, [
  'moudle',
  'av_log_level',
  'used_signal_type',
  'nmos',
  'ssm_address_range',
  'authorization_service',
])

const nicDetails = ref<NicDetail[]>([])
const indexedNicDetails = computed(() => {
  return nicDetails.value
    .filter(nic => nic.nicIndex >= 0)
    .map((nic, idx) => ({ ...nic, index: idx }))
})

const inputParams = ref(def_mg_player_params())
const outputParams = ref(Array.from({ length: 2 }, () => {
  return ref(def_mg_output_item_params())
}))
const showSecondOutput = ref(false)

const [videoFormatEnum, videoSelected, videoUnSelected] = useUsedFormat()
const [audioFormatEnum, audioSelected, audioUnSelected] = useUsedFormat()

watchNmosName(() => props.name, mv)

function getValue() {
  const result = {
    ...handle(mv.value),
    videoformat_enum: videoFormatEnum.value.map(vfn => handleVideoFormat(vfn, videoFormats.value)).filter(v => v),
    audioformat_enum: audioFormatEnum.value.map(afn => handleAudioFormat(afn, audioFormats.value)).filter(a => a),
    input: wrap(handleInputParams(inputParams.value, videoFormats.value), 'in_'),
    output: wrap({
      out_params: outputParams.value.slice(0, showSecondOutput.value ? 2 : 1).map(o => handleOutputParams(o.value, videoFormats.value))
    }, 'out_')
  }
  if (mv.value.used_signal_type !== 1) {
    result.nic_list = handleNicList(indexedNicDetails.value)
  }
  return result
}

function setValue(data: typeof mgData) {
  mv.value = pick(data, [
    'moudle',
    'av_log_level',
    'used_signal_type',
    'nmos',
    'ssm_address_range',
    'authorization_service',
  ])
  const _input = unwrap(data.input, 'in_')
  inputParams.value = checkPlayerParams(merge(def_mg_player_params(), _input), videoFormats.value, audioFormats.value)
  const _output = unwrap(data.output.out_params, 'out_')
  showSecondOutput.value = _output.length > 1
  outputParams.value = Array.from({ length: 2 }, (_, idx) => {
    return checkPlayerParams(merge(def_mg_output_item_params(), _output[idx]), videoFormats.value, audioFormats.value)
  })
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
          <NicSection v-if="mv.used_signal_type !== 1" v-model="nicDetails" :nics="nics" />
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
        <MediaGateInput
          v-model="inputParams"
          :used-signal-type="mv.used_signal_type"
          :nics="indexedNicDetails"
          @video-selected="videoSelected"
          @video-unselected="videoUnSelected"
          @audio-selected="audioSelected"
          @audio-unselected="audioUnSelected"
        />
        <div class="form-outer">
          <div class="form-header">
            <div class="form-header-inner">
              <div class="left">
                <h4>输出参数</h4>
              </div>
            </div>
          </div>
          <div class="form-body">
            <MediaGateOutputParams
              v-for="(_, pidx) in outputParams.slice(0, showSecondOutput ? 2 : 1)"
              :key="pidx"
              v-model="outputParams[pidx]"
              :index="pidx"
              :used-signal-type="mv.used_signal_type"
              :input-video-format="inputParams.videoformat_name || ''"
              :nics="indexedNicDetails"
              :is-last="pidx === 1"
              :default-open="showSecondOutput && pidx === 1"
              :deletable="showSecondOutput && pidx === 1"
              @video-selected="videoSelected"
              @video-unselected="videoUnSelected"
              @audio-selected="audioSelected"
              @audio-unselected="audioUnSelected"
              @delete="showSecondOutput = false"
            />
            <div v-if="!showSecondOutput" class="form-action-buttons">
              <div class="left">
                <VButton
                  class="is-rounded"
                  color="primary"
                  raised
                  @click="showSecondOutput = !showSecondOutput"
                >
                  添加第2路输出参数
                </VButton>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </form>
</template>
<style lang="scss" scoped>
.form-action-buttons {
  display: flex;
  justify-content: center;
  align-items: center;
  margin: 32px 0;
}
</style>
