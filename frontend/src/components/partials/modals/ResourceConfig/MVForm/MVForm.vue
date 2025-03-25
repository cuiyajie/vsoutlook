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
  type IndexedType,
} from '../Utilties/Consts_V1'
import { handle, watchNmosName } from '../Utilties/Utils';
import pick from 'lodash-es/pick'
import merge from 'lodash-es/merge'
import mvData from '@src/data/vscomponent/mv.json'
import { useUsedFormat } from '../Utilties/Composables';
import { useUserSession } from "@src/stores/userSession"
import { handleVideoFormat, handleAudioFormat, handlePlayerParams, handleNicList, unwrap, wrap, checkPlayerParams } from '../Utilties/Utils_V1';
import { def_mv_input_param, def_mv_output_param, type AuditAlarmRule, type MVInputItemParam, type MVOutputItemParam } from './Consts';
import { checkInputParam, handleAuditAlarmRule, handleAuditAvTemplate, handleInputParams, handleOutputParams, techReview2AudioTmpl } from './Utils';

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
const audits = computed(() => techReview2AudioTmpl(usStore.settings.tech_reviews || []));

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

const auditAlarmRules = ref<AuditAlarmRule[]>([])
const indexedAuditAlarmRules = computed(() => {
  return auditAlarmRules.value
    .filter(rule => rule.rule_name && rule.av_alarm.audit_template_name)
    .map((rule, idx) => ({ ...rule, index: idx }))
})
function getAudioAlarmRule(ruleName: string) {
  return indexedAuditAlarmRules.value.find(rule => rule.rule_name === ruleName)
}

const [videoFormatEnum, videoSelected, videoUnSelected] = useUsedFormat()
const [audioFormatEnum, audioSelected, audioUnSelected] = useUsedFormat()
const [auditRuleEnum, auditRuleSelected, auditRuleUnSelected] = useUsedFormat()

watchNmosName(() => props.name, mv)

function getValue() {
  const result = {
    ...handle(mv.value),
    videoformat_enum: videoFormatEnum.value.map(vfn => handleVideoFormat(vfn, videoFormats.value)).filter(v => v),
    audioformat_enum: audioFormatEnum.value.map(afn => handleAudioFormat(afn, audioFormats.value)).filter(a => a),
    audit_av_template_enum: auditRuleEnum.value
      .map(atn => handleAuditAvTemplate(atn, indexedAuditAlarmRules.value, audits.value.auditAvTemplates))
      .filter(a => a),
    audit_alarm_rule_enum: indexedAuditAlarmRules.value.map(rule => handleAuditAlarmRule(rule)),
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
    'api_params',
  ])
  mv.value.output_number = 1

  nicDetails.value = data.nic_list.map((nic: any) => {
    const nicIndex = props.nics.findIndex(n => n.nicNameMain === nic.nic_name_m && n.nicNameBackup === nic.nic_name_b)
    if (nicIndex !== -1) {
      return { ...nic, nicIndex, id: props.nics[nicIndex].id }
    }
    return { ...nic, nicIndex: -1, id: '' }
  })

  auditAlarmRules.value = data.audit_alarm_rule_enum.map((rule: any) => {
    const parma = audits.value.auditAlarmRuleMap[rule.av_alarm.audit_template_name]
    return parma ? rule : null
  }).filter(a => a)

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
              <AddrInputAddonSep
                v-model:host="mv.api_params[0].api_name" v-model:port="mv.api_params[0].service_port"
                label="tally通知接口"
              />
            </div>
            <div class="column is-6">
              <AddrInputAddonSep
                v-model:host="mv.api_params[1].api_name" v-model:port="mv.api_params[1].service_port"
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
        <AuditAlarmRuleSection v-model="indexedAuditAlarmRules" />
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
                <h4>使用的技审报警规则列表</h4>
              </div>
            </div>
          </div>
          <div class="form-body">
            <VTags v-if="auditRuleEnum.length > 0" class="format-item-container">
              <AuditAvTemplateTooltip v-for="rule in auditRuleEnum" :key="rule" :name="getAudioAlarmRule(rule)?.av_alarm.audit_template_name || ''">
                <VTag color="danger" :label="rule" curved outlined />
              </AuditAvTemplateTooltip>
            </VTags>
            <div v-else class="is-empty-list">暂时还没有选择技审报警规则</div>
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
              :rules="indexedAuditAlarmRules"
              :used-signal-type="mv.used_signal_type"
              @video-selected="videoSelected"
              @video-unselected="videoUnSelected"
              @audio-selected="audioSelected"
              @audio-unselected="audioUnSelected"
              @audit-rule-selected="auditRuleSelected"
              @audit-rule-unselected="auditRuleUnSelected"
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
              :rules="indexedAuditAlarmRules"
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
