import { type IndexedParmas } from '../Utilties/Consts_V1'
import { checkPlayerParams, handlePlayerParams } from '../Utilties/Utils_V1'
import {
  type AuditAvTemplate,
  type AuditAlarmRule,
  type MVInputItemParam,
  type MVOutputItemParam,
} from './Consts'

export function techReview2AudioTmpl(techReviews: TechReview[]) {
  const auditAvTemplates: AuditAvTemplate[] = []
  const auditAlarmRuleMap: Record<string, AuditAlarmRule['av_alarm']> = {}
  techReviews.forEach((techReview) => {
    const tmpl = { name: techReview.name } as AuditAvTemplate
    const alarm = {
      audit_template_name: techReview.name,
      video_alarm_priority: [],
      audio_alarm_priority: [],
      audio_mute_rule: {
        condition_any: '',
        condition_all: '',
      },
    } as AuditAlarmRule['av_alarm']
    techReview.videoRules.forEach((rule, i) => {
      const { key, ...rest } = rule
      tmpl[key] = rest
      alarm.video_alarm_priority.push({ priority: i + 1, itemname: rule.key })
    })
    techReview.audioRules.forEach((rule, i) => {
      const { key, ...rest } = rule
      tmpl[key] = rest
      alarm.audio_alarm_priority.push({ priority: i + 1, itemname: rule.key })
    })
    alarm.audio_mute_rule.condition_any = techReview.condition_any
    alarm.audio_mute_rule.condition_all = techReview.condition_all
    auditAvTemplates.push(tmpl)
    auditAlarmRuleMap[techReview.name] = alarm
  })
  return {
    auditAvTemplates,
    auditAlarmRuleMap,
  }
}

export function handleAuditAvTemplate(
  ruleName: string,
  indexedRules: IndexedParmas<AuditAlarmRule>[],
  avTemplates: ReturnType<typeof techReview2AudioTmpl>['auditAvTemplates']
) {
  const rule = indexedRules.find((rule) => rule.rule_name === ruleName)
  if (!rule) return null
  return (
    avTemplates.find((tmpl) => tmpl.name === rule.av_alarm.audit_template_name) || null
  )
}

export function handleAuditAlarmRule(rule: IndexedParmas<AuditAlarmRule>) {
  // eslint-disable-next-line @typescript-eslint/no-unused-vars
  const { index, ...rest } = rule
  return rest
}

export function checkInputParam(
  params: any,
  vfs: VideoFormat[],
  afs: AudioFormat[],
  rules: AuditAlarmRule[]
) {
  const result = checkPlayerParams(params, vfs, afs)
  const rule = rules.find(
    (r) => r.av_alarm.audit_template_name === result.audit_alarm_rule_name
  )
  if (!rule) {
    params.audit_alarm_rule_name = ''
  }
  return result
}

export function handleInputParams(
  input: MVInputItemParam,
  vfs: VideoFormat[],
  index: number
) {
  const result = handlePlayerParams(input, vfs)
  result.index = index
  return result
}

export function handleOutputParams(
  output: MVOutputItemParam,
  vfs: VideoFormat[],
  index: number
) {
  const result = handlePlayerParams(output, vfs)
  result.index = index
  return result
}
