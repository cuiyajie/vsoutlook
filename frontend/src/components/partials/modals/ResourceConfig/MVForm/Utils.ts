import { checkPlayerParams, handlePlayerParams } from '../Utilties/Utils_V1'
import { type MVInputItemParam, type MVOutputItemParam } from './Consts'

export function handleAuditAvTemplate(name: string, avTemplates: TechReview[]) {
  const tmpl = avTemplates.find((tmpl) => tmpl.name === name)
  if (!tmpl) return null
  return tmpl
}

export function handleAuditAlarmRule(
  ruleName: string,
  auditAlarmRules: AuditAlarmRule[]
) {
  const rule = auditAlarmRules.find((r) => r.rule_name === ruleName)
  if (!rule) return null
  return rule
}

export function checkInputParam(
  params: any,
  vfs: VideoFormat[],
  afs: AudioFormat[],
  rules: AuditAlarmRule[]
) {
  const result = checkPlayerParams(params, vfs, afs)
  const rule = rules.find((r) => r.rule_name === result.audit_alarm_rule_name)
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
  const result = handlePlayerParams(input, vfs, 'receive')
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
