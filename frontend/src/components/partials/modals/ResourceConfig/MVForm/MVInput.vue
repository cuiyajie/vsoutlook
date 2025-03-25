<script lang="ts" setup>
import { type IndexedNicDetail } from '../Utilties/Consts_V1';
import { type MVInputItemParam, def_mv_input_param } from './Consts';
import { type AuditAlarmRule } from './Consts';

const mv = defineModel<MVInputItemParam>({
  default: def_mv_input_param(),
  local: true,
})

defineProps<{
  usedSignalType: number,
  nics: IndexedNicDetail[],
  rules: AuditAlarmRule[],
  index: number
  isLast: boolean
}>()

const emit = defineEmits<{
  (e: 'video-selected', name: string): void,
  (e: 'video-unselected', name: string): void,
  (e: 'audio-selected', name: string): void,
  (e: 'audio-unselected', name: string): void,
  (e: 'audit-rule-selected', name: string): void,
  (e: 'audit-rule-unselected', name: string): void
}>()

const opened = ref(false)

</script>
<template>
  <div class="form-fieldset collapse-form" :class="!isLast && 'seperator'" :open="opened || undefined">
    <div
      class="fieldset-heading collapse-header"
      tabindex="0"
      role="button"
      @keydown.space.prevent="opened = !opened"
      @click.prevent="opened = !opened"
    >
      <h4>第{{ index + 1 }}路输入参数&nbsp;&nbsp;（序号: {{ index }}）</h4>
      <div class="collapse-icon">
        <VIcon icon="feather:chevron-down" />
      </div>
    </div>
    <Transition name="fade-slow">
      <div v-show="opened" class="form-fieldset-nested-2">
        <div class="form-fieldset-nested-3 seperator">
          <div class="columns is-multiline">
            <div class="column is-6">
              <VField>
                <VLabel>视频格式名称</VLabel>
                <VControl>
                  <VideoFormatSelect v-model="mv.videoformat_name" :used-signal-type="usedSignalType" @video-selected="emit('video-selected', $event)" @video-unselected="emit('video-unselected', $event)" />
                </VControl>
              </VField>
            </div>
            <div class="column is-6">
              <VField>
                <VLabel>输入视频元数据是否强制使用指定的色域</VLabel>
                <VControl class="in-form">
                  <VSwitchBlock
                    v-model="mv.force_use_videoformat"
                    color="primary"
                  />
                </VControl>
              </VField>
            </div>
            <div class="column is-6">
              <VField>
                <VLabel>音频格式名称</VLabel>
                <VControl>
                  <AudioFormatSelect v-model="mv.audioformat_name" @audio-selected="emit('audio-selected', $event)" @audio-unselected="emit('audio-unselected', $event)" />
                </VControl>
              </VField>
            </div>
            <div class="column is-6">
              <VField>
                <VLabel>使用的技审报警规则模板</VLabel>
                <VControl>
                  <AuditAlarmRuleSelect
                    v-model="mv.audit_alarm_rule_name"
                    :rules="rules"
                    @audit-rule-selected="emit('audit-rule-selected', $event)"
                    @audit-rule-unselected="emit('audit-rule-unselected', $event)"
                  />
                </VControl>
              </VField>
            </div>
          </div>
        </div>
        <PlayerParamsForm v-model="mv" :nics="nics" />
      </div>
    </Transition>
  </div>
</template>
<style lang="scss"></style>
