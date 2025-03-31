<!-- eslint-disable vue/prop-name-casing -->
<script setup lang="ts">
import { useUserSession } from "@src/stores/userSession";
import { type AuditAlarmRule } from './Consts'
import { techReview2AudioTmpl } from "./Utils";
import { useNotyf } from "@src/composable/useNotyf";

const notyf = useNotyf();

const props = defineProps<{
  index: number
  isLast: boolean
  hasError: (name: string, field: 'name' | 'tmpl', index: number) => string
}>()

const emit = defineEmits<{
  (e: 'remove'): void
}>()

const mv = defineModel<AuditAlarmRule>({
  default: {} as AuditAlarmRule,
  local: true,
})

const usStore = useUserSession();
const audits = computed(() =>
  techReview2AudioTmpl(usStore.settings.tech_reviews || [])
);

const auditAvTemplates = computed(() => audits.value.auditAvTemplates);

const selectRef = ref<any>(null)

function checkRuleName(name: string) {
  const errorText = props.hasError(name, 'name', props.index)
  if (errorText) {
    notyf.error(errorText)
    mv.value.rule_name = ''
  }
}

function onAuditAvTemplateSelect(tmplName: string) {
  if (tmplName === mv.value.av_alarm.audit_template_name) return
  const errorText = props.hasError(tmplName, 'tmpl', props.index)
  if (errorText) {
    notyf.error(errorText)
    if (selectRef.value) {
      const ov = mv.value.av_alarm.audit_template_name
      if (ov) {
        selectRef.value.update(selectRef.value.getOption(ov))
      } else {
        selectRef.value.clear()
      }
    }
    return
  }
  Object.assign(mv.value.av_alarm, audits.value.auditAlarmRuleMap[tmplName])
}

watch(() => mv.value.av_alarm.audit_template_name, (nv) => {
  selectRef.value?.select(nv)
}, { immediate: true })

</script>
<template>
  <div
    class="form-fieldset collapse-form"
    :class="!isLast && 'seperator'"
    open
  >
    <div
      class="fieldset-heading collapse-header"
      tabindex="0"
      role="button"
    >
      <h4>第 {{ index + 1 }} 个技审报警规则</h4>
      <div class="collapse-icon">
        <VIconButton
          color="warning"
          light
          raised
          circle
          icon="feather:x"
          @click="emit('remove')"
        />
      </div>
    </div>
    <div class="form-fieldset-nested-4">
      <div class="columns is-multiline">
        <div class="column is-6">
          <VField>
            <VLabel>规则名称</VLabel>
            <VControl>
              <VInput v-model="mv.rule_name" type="text" placeholder="请输入名称" @blur="checkRuleName($event.target.value)" />
            </VControl>
          </VField>
        </div>
        <div class="column is-6">
          <VField>
            <VLabel>选择视音频技审参数模板</VLabel>
            <VControl>
              <Multiselect
                ref="selectRef"
                class="tippy-select"
                placeholder="选择视音频技审参数模板"
                no-options-text="还没有配置视音频技审参数模板"
                value-prop="name"
                label="name"
                :searchable="false"
                :can-deselect="false"
                :can-clear="false"
                :max-height="145"
                :options="auditAvTemplates"
                @select="onAuditAvTemplateSelect"
              >
                <template #singlelabel="{ value }">
                  <div class="multiselect-single-label">
                    <span class="select-label-text">
                      {{ value.name }}
                    </span>
                  </div>
                </template>
                <template #option="{ option }">
                  <span class="select-option-text">
                    {{ option.name }}
                  </span>
                  <AuditAvTemplateTooltip :name="option.name">
                    <i class="iconify" data-icon="feather:alert-circle" aria-hidden="true"></i>
                  </AuditAvTemplateTooltip>
                </template>
              </Multiselect>
            </VControl>
          </VField>
        </div>
      </div>
    </div>
  </div>
</template>
<style lang="scss">
@import '@src/scss/abstracts/all';
@import '@src/scss/components/forms-outer';
</style>

