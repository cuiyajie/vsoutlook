<script lang="ts" setup>
import { useUserSession } from '@src/stores/userSession';

const mv = defineModel<string>({
  default: '',
  local: true,
});

const usStore = useUserSession()
const alarmRules = computed(() => usStore.settings.audit_alarm_rules || [])

const emit = defineEmits<{
  (e: 'audit-rule-selected', name: string): void,
  (e: 'audit-rule-unselected', name: string): void,
  (e: 'av-tmpl-selected', name: string): void,
  (e: 'av-tmpl-unselected', name: string): void,
}>()

watch<string>(() => mv.value, (nv, ov) => {
  if (nv === ov) return
  if (nv) {
    emit('audit-rule-selected', nv)
    const rule = alarmRules.value.find(r => r.rule_name === nv)
    if (rule) {
      emit('av-tmpl-selected', rule.av_alarm.audit_template_name)
    }
  } else {
    emit('audit-rule-unselected', ov)
    const rule = alarmRules.value.find(r => r.rule_name === ov)
    if (rule) {
      emit('av-tmpl-unselected', rule.av_alarm.audit_template_name)
    }
  }
})

</script>
<template>
  <Multiselect
    v-model="mv"
    class="tippy-select"
    placeholder="选择报警规则"
    value-prop="rule_name"
    label="rule_name"
    :can-deselect="false"
    :style="{'--ms-max-height': '245px'}"
    no-options-text="暂时没有配置报警规则"
    :options="alarmRules"
  >
    <template #singlelabel="{ value }">
      <div class="multiselect-single-label">
        <span class="select-label-text">
          {{ value.rule_name }}
        </span>
      </div>
    </template>
    <template #option="{ option }">
      <span class="select-option-text">
        {{ option.rule_name }}
      </span>
      <AuditAlarmRuleTooltip :value="option">
        <i class="iconify" data-icon="feather:alert-circle" aria-hidden="true"></i>
      </AuditAlarmRuleTooltip>
    </template>
  </Multiselect>
</template>
