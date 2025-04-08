<script lang="ts" setup>
import { useUserSession } from "@src/stores/userSession";
import { def_audit_alarm_rule, def_audit_av_template } from "./Consts";
import { VideoReviewKeyName, AudioReviewKeyName } from "@src/components/partials/modals/PresetConfig/Consts"

const props = defineProps<{
  value?: AuditAlarmRule,
  name?: string
}>()

const usStore = useUserSession();
const techReviews = computed(() => usStore.settings.tech_reviews || []);
const auditAlarmRules = computed(() => usStore.settings.audit_alarm_rules || []);

const avTemplate = computed(() => {
  return techReviews.value.find((f) => f.name === props.name) || def_audit_av_template();
});

const auditAlarmRule = computed(() => {
  if (props.value) return props.value
  if (props.name) {
    return auditAlarmRules.value.find(f => f.rule_name === props.name) || def_audit_alarm_rule()
  }
  return def_audit_alarm_rule()
})

const tippyAppendTo = () => document.body;
</script>
<template>
  <tippy
    animation="scale"
    theme="light"
    placement="right"
    :arrow="true"
    arrow-type="round"
    :interactive="true"
    :append-to="tippyAppendTo"
  >
    <template #content>
      <div class="kv-info tippy-detail">
        <div class="info-block-inner is-dark-bordered-12">
          <div class="info-block-line">
            <h4 class="dark-inverted">模板名称</h4>
            <span>
              {{ auditAlarmRule.rule_name }}
            </span>
          </div>
          <div class="info-block-line has-info">
            <h4 class="dark-inverted">视频技审</h4>
            <div>
              <span v-for="rule in auditAlarmRule.av_alarm.video_alarm_priority" :key="rule.priority" class="info-right">
                <span class="is-title">优先级 {{ rule.priority }}: {{ VideoReviewKeyName[rule.itemname as TVideoRuleKey] }}</span>
                <span v-if="avTemplate[rule.itemname as TVideoRuleKey].threshold_percentage" class="is-sub">阈值 {{ avTemplate[rule.itemname as TVideoRuleKey].threshold_percentage }}</span>
                <span v-if="avTemplate[rule.itemname as TVideoRuleKey].duration_frames" class="is-sub">时长 {{ avTemplate[rule.itemname as TVideoRuleKey].duration_frames }} 帧</span>
                <span v-if="avTemplate[rule.itemname as TVideoRuleKey].duration_ms" class="is-sub">时长 {{ avTemplate[rule.itemname as TVideoRuleKey].duration_ms }} ms</span>
              </span>
            </div>
          </div>
          <div class="info-block-line has-info">
            <h4 class="dark-inverted">音频技审</h4>
            <div>
              <span v-for="rule in auditAlarmRule.av_alarm.audio_alarm_priority" :key="rule.priority" class="info-right">
                <span class="is-title">优先级 {{ rule.priority }}: {{ AudioReviewKeyName[rule.itemname as TAudioRuleKey] }}</span>
                <span v-if="avTemplate[rule.itemname as TAudioRuleKey].detect_channels" class="is-sub">声道 {{ avTemplate[rule.itemname as TAudioRuleKey].detect_channels }}</span>
                <span v-if="avTemplate[rule.itemname as TAudioRuleKey].duration_frames" class="is-sub">时长 {{ avTemplate[rule.itemname as TAudioRuleKey].duration_frames }} 帧</span>
                <span v-if="avTemplate[rule.itemname as TAudioRuleKey].threshold_dbfs" class="is-sub">阈值 {{ avTemplate[rule.itemname as TAudioRuleKey].threshold_dbfs }}</span>
                <span v-if="avTemplate[rule.itemname as TAudioRuleKey].duration_ms" class="is-sub">时长 {{ avTemplate[rule.itemname as TAudioRuleKey].duration_ms }} ms</span>
              </span>
            </div>
          </div>
        </div>
      </div>
    </template>
    <slot></slot>
  </tippy>
</template>
