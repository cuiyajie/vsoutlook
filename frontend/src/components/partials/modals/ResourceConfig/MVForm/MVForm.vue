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
  type IndexedType,
} from '../Utilties/Consts_V1'
import { handle, watchNmosName } from '../Utilties/Utils';
import pick from 'lodash-es/pick'
import merge from 'lodash-es/merge'
import mvData from '@src/data/vscomponent/mv.json'
import { useUsedFormat } from '../Utilties/Composables';
import { useUserSession } from "@src/stores/userSession"
import { handleVideoFormat, handleAudioFormat, handleNicList, unwrap, wrap, handleApiParams, checkApiParams, checkPlayerParams, checkNicDetails } from '../Utilties/Utils_V1';
import { def_mv_input_param, def_mv_output_param, type MVInputItemParam, type MVOutputItemParam, def_api_params } from './Consts';
import { checkInputParam, handleAuditAlarmRule, handleAuditAvTemplate, handleInputParams, handleOutputParams } from './Utils';
import { useNicList } from '../Utilties/Composable';

const props = defineProps<{
  name: string,
  nics: NicInfo[],
  requiredment?: TmplRequirement
}>()

const mv = defineModel<{
  moudle: string
  av_log_level: number
  input_number: number
  output_number: number
  used_signal_type: number
  nmos: NMosConfigType
  ssm_address_range: SSMAddressType[]
}>({
  default: {
    moudle: 'mv',
    av_log_level: 5,
    input_number: 4,
    output_number: 1,
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
const avTemplates = computed(() => usStore.settings.tech_reviews || [])
const auditAlarmRules = computed(() => usStore.settings.audit_alarm_rules || [])

mv.value = pick(mvData, [
  'moudle',
  'input_number',
  'output_number',
  'av_log_level',
  'used_signal_type',
  'nmos',
  'ssm_address_range',
  'authorization_service',
])

const apiParams = ref<IApiParams[]>(def_api_params())

const { nicDetails, indexedNicDetails } = useNicList(props)

const templates = computed(() => usStore.settings.mv_template_list || [])
function mvTemplateIsValid(tmpl: string) {
  return !!templates.value.find((v) => v.path === tmpl)
}

const inputs = ref<IndexedType<MVInputItemParam>[]>([])
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
            index: len + i,
            value: reactive(def_mv_input_param()),
          }
        })
      )
    }
  },
  { immediate: true }
)

const outputs = ref<IndexedType<MVOutputItemParam>[]>([])
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
            index: len + i,
            value: reactive(def_mv_output_param()),
          }
        })
      )
    }
    outputs.value.forEach((v) => {
      if (v.value.out_mv_template && !mvTemplateIsValid(v.value.out_mv_template)) {
        v.value.out_mv_template = ''
      }
    })
  },
  { immediate: true }
)

const [videoFormatEnum, videoSelected, videoUnSelected] = useUsedFormat()
const [audioFormatEnum, audioSelected, audioUnSelected] = useUsedFormat()
const [avTmplEnum, avTmplSelected, avTmplUnSelected] = useUsedFormat()
const [auditRuleEnum, auditRuleSelected, auditRuleUnSelected] = useUsedFormat()

watchNmosName(() => props.name, mv)

function getValue() {
  const result = {
    ...handle(mv.value),
    api_params: handleApiParams(apiParams.value),
    videoformat_enum: videoFormatEnum.value.map(vfn => handleVideoFormat(vfn, videoFormats.value)).filter(v => v),
    audioformat_enum: audioFormatEnum.value.map(afn => handleAudioFormat(afn, audioFormats.value)).filter(a => a),
    audit_av_template_enum: avTmplEnum.value
      .map(tmpl => handleAuditAvTemplate(tmpl, avTemplates.value))
      .filter(a => a),
    audit_alarm_rule_enum: auditRuleEnum.value.map(rule => handleAuditAlarmRule(rule, auditAlarmRules.value)),
    input: wrap({
      input_params: inputs.value.map((input, i) => handleInputParams(input.value, videoFormats.value, i))
    }, 'in_'),
    output: wrap({
      out_params: outputs.value.map((output, i) => {
        let v = output.value
        const optRef = outputRefs.value[i]
        if (optRef?.getValue) {
          v = optRef.getValue()
        }
        return handleOutputParams(v, videoFormats.value, i)
      })
    }, 'out_'),
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
  ])
  mv.value.output_number = 1

  apiParams.value = checkApiParams(def_api_params(), data.api_params)

  nicDetails.value = checkNicDetails(data.nic_list, props.nics)

  const _ipData = unwrap(data.input, 'in_')
  nextTick(() => {
    inputs.value.forEach((iptv, idx) => {
      iptv.value = checkInputParam(
        merge(def_mv_input_param(), _ipData.input_params[idx]),
        videoFormats.value,
        audioFormats.value,
        auditAlarmRules.value
      )
    })
  })

  const _opData = unwrap(data.output, 'out_')
  nextTick(() => {
    outputs.value.forEach((optv, idx) => {
      optv.value = checkPlayerParams(
        merge(def_mv_output_param(), _opData.params[idx]),
        videoFormats.value,
        audioFormats.value
      )
      if (
        optv.value.out_mv_template &&
        !mvTemplateIsValid(optv.value.out_mv_template)
      ) {
        optv.value.out_mv_template = ''
      }
    })
  })
}

defineExpose({
  getValue,
  setValue,
})
</script>
<template>
  <!-- prettier-ignore -->
  <form method="post" novalidate class="form-layout device-form">
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
                  <VSelect v-model="mv.used_signal_type" class="is-rounded">
                    <VOption v-for="ust in used_signal_types" :key="ust.key" :value="ust.key">{{ ust.label }}</VOption>
                  </VSelect>
                </VControl>
              </VField>
            </div>
            <div class="column is-6">
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
          </div>
        </div>
        <div class="form-fieldset">
          <div class="fieldset-heading">
            <h4>接口设置</h4>
          </div>
          <div class="columns is-multiline">
            <div class="column is-6">
              <ApiParamsCheck
                v-model="apiParams[0]"
                label="tally通知接口"
              />
            </div>
            <div class="column is-6">
              <ApiParamsCheck
                v-model="apiParams[1]"
                label="动态控制接口"
              />
            </div>
          </div>
        </div>
        <div class="form-fieldset collapse-form form-outer" :open="advanceOpened || undefined">
          <div
            class="fieldset-heading collapse-header" tabindex="0" role="button"
            @keydown.space.prevent="advanceOpened = !advanceOpened" @click.prevent="advanceOpened = !advanceOpened"
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
        <expand-transition>
          <NicSection v-if="mv.used_signal_type !== 1" v-model="nicDetails" :nics="nics" />
        </expand-transition>
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
        <div class="form-outer">
          <div class="form-header">
            <div class="form-header-inner">
              <div class="left">
                <h4>使用的技审模板列表</h4>
              </div>
            </div>
          </div>
          <div class="form-body">
            <VTags v-if="avTmplEnum.length > 0" class="format-item-container">
              <TechReviewTooltip v-for="tmpl in avTmplEnum" :key="tmpl" :name="tmpl">
                <VTag color="orange" :label="tmpl" curved outlined />
              </TechReviewTooltip>
            </VTags>
            <div v-else class="is-empty-list">暂时还没有选择技审模板</div>
          </div>
        </div>
        <div class="form-outer">
          <div class="form-header">
            <div class="form-header-inner">
              <div class="left">
                <h4>使用的报警规则列表</h4>
              </div>
            </div>
          </div>
          <div class="form-body">
            <VTags v-if="auditRuleEnum.length > 0" class="format-item-container">
              <AuditAlarmRuleTooltip v-for="rule in auditRuleEnum" :key="rule" :name="rule">
                <VTag color="danger" :label="rule" curved outlined />
              </AuditAlarmRuleTooltip>
            </VTags>
            <div v-else class="is-empty-list">暂时还没有选择报警规则</div>
          </div>
        </div>
        <div class="form-outer">
          <div class="form-header">
            <div class="form-header-inner">
              <div class="left">
                <h4>输入参数</h4>
              </div>
            </div>
          </div>
          <div class="form-body">
            <MVInput
              v-for="(ipt, idx) in inputs"
              :key="idx"
              v-model="ipt.value"
              :index="ipt.index"
              :is-last="idx === inputs.length - 1"
              :nics="indexedNicDetails"
              :used-signal-type="mv.used_signal_type"
              @video-selected="videoSelected"
              @video-unselected="videoUnSelected"
              @audio-selected="audioSelected"
              @audio-unselected="audioUnSelected"
              @audit-rule-selected="auditRuleSelected"
              @audit-rule-unselected="auditRuleUnSelected"
              @av-tmpl-selected="avTmplSelected"
              @av-tmpl-unselected="avTmplUnSelected"
            />
          </div>
        </div>
        <div class="form-outer">
          <div class="form-header">
            <div class="form-header-inner">
              <div class="left">
                <h4>输出参数</h4>
              </div>
            </div>
          </div>
          <div class="form-body">
            <MVOutput
              v-for="(otp, idx) in outputs"
              ref="outputRefs"
              :key="idx"
              v-model="otp.value"
              :index="otp.index"
              :is-last="idx === outputs.length - 1"
              :nics="indexedNicDetails"
              :used-signal-type="mv.used_signal_type"
              :pips-number="mv.input_number"
              @video-selected="videoSelected"
              @video-unselected="videoUnSelected"
              @audio-selected="audioSelected"
              @audio-unselected="audioUnSelected"
            />
          </div>
        </div>
      </div>
    </div>
  </form>
</template>
<style lang="scss">
</style>
