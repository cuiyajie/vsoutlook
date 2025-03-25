<script lang="ts" setup>
import { type AuditAlarmRule } from './Consts';

const mv = defineModel<string>({
  default: '',
  local: true,
});

defineProps<{
  rules: AuditAlarmRule[]
}>()

const emit = defineEmits<{
  (e: 'audit-rule-selected', name: string): void,
  (e: 'audit-rule-unselected', name: string): void
}>()

watch<string>(() => mv.value, (nv, ov) => {
  if (nv === ov) return
  if (nv) {
    emit('audit-rule-selected', nv)
  } else {
    emit('audit-rule-unselected', ov)
  }
})

</script>
<template>
  <Multiselect
    v-model="mv"
    placeholder="选择技审报警规则模板"
    value-prop="rule_name"
    label="rule_name"
    :searchable="false"
    :can-deselect="false"
    :can-clear="false"
    :max-height="145"
    no-options-text="暂时没有配置技审报警规则模板"
    :options="rules"
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
    </template>
  </Multiselect>
</template>
