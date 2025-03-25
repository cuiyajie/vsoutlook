<!-- eslint-disable vue/prop-name-casing -->
<script setup lang="ts">
import { type AuditAlarmRule, def_audit_alarm_rule } from './Consts';

const mv = defineModel<AuditAlarmRule[]>({
  default: [],
  local: true,
})

function add() {
  mv.value.push(def_audit_alarm_rule())
}

function remove(idx: number) {
  mv.value.splice(idx, 1)
}

function hasError(name: string, field: 'name' | 'tmpl', index: number) {
  if (field === 'tmpl' && mv.value.some((nd) => nd.av_alarm.audit_template_name === name)) {
    return '该技审规则模板已添加'
  }
  if (field === 'name' && mv.value.some((nd, nidx) => nd.rule_name === name && nidx !== index)) {
    return '规则名称重复'
  }
  return ''
}
</script>
<template>
  <div class="form-outer">
    <div class="form-header">
      <div class="form-header-inner">
        <div class="left">
          <h4>技审报警规则列表</h4>
        </div>
        <button
          class="button is-circle is-dark-outlined"
          @keydown.space.prevent="add"
          @click.prevent="add"
        >
          <span class="icon is-small">
            <i aria-hidden="true" class="iconify" data-icon="feather:plus" />
          </span>
        </button>
      </div>
    </div>
    <div v-if="mv.length" class="form-body">
      <!--Fieldset-->
      <AuditAlarmRuleItem
        v-for="(rule, ridx) in mv"
        :key="ridx"
        v-model="mv[ridx]"
        :index="ridx"
        :is-last="ridx === mv.length - 1"
        :has-error="hasError"
        @remove="remove(ridx)"
      />
    </div>
    <div v-else class="is-empty-list">暂无设置技审报警规则</div>
  </div>
</template>
