<script lang="ts" setup>
import { useUserSession } from "@src/stores/userSession";
import { techReview2AudioTmpl } from "./Utils";
import { def_audit_av_template } from "./Consts";
import { VideoReviewKeyName, AudioReviewKeyName } from "@src/components/partials/modals/PresetConfig/Consts"

const props = defineProps<{
  name: string;
}>();

const usStore = useUserSession();
const audits = computed(() =>
  techReview2AudioTmpl(usStore.settings.tech_reviews || [])
);

const avTemplate = computed(() => {
  return audits.value.auditAvTemplates.find((f) => f.name === props.name) || def_audit_av_template();
});

const auditAlarmRuleMap = computed(() => audits.value.auditAlarmRuleMap)

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
              {{ avTemplate.name }}
            </span>
          </div>
          <div class="info-block-line has-info">
            <h4 class="dark-inverted">视频技审</h4>
            <div>
              <span v-for="rule in auditAlarmRuleMap[avTemplate.name].video_alarm_priority" :key="rule.priority" class="info-right">
                <span class="is-title">优先级 {{ rule.priority }}: {{ VideoReviewKeyName[rule.itemname] }}</span>
                <span v-if="avTemplate[rule.itemname].threshold_percentage" class="is-sub">阈值 {{ avTemplate[rule.itemname].threshold_percentage }}</span>
                <span v-if="avTemplate[rule.itemname].duration_frames" class="is-sub">时长 {{ avTemplate[rule.itemname].duration_frames }} 帧</span>
                <span v-if="avTemplate[rule.itemname].duration_ms" class="is-sub">时长 {{ avTemplate[rule.itemname].duration_ms }} ms</span>
              </span>
            </div>
          </div>
          <div class="info-block-line has-info">
            <h4 class="dark-inverted">音频技审</h4>
            <div>
              <span v-for="rule in auditAlarmRuleMap[avTemplate.name].audio_alarm_priority" :key="rule.priority" class="info-right">
                <span class="is-title">优先级 {{ rule.priority }}: {{ AudioReviewKeyName[rule.itemname] }}</span>
                <span v-if="avTemplate[rule.itemname].detect_channels" class="is-sub">声道 {{ avTemplate[rule.itemname].detect_channels }}</span>
                <span v-if="avTemplate[rule.itemname].duration_frames" class="is-sub">时长 {{ avTemplate[rule.itemname].duration_frames }} 帧</span>
                <span v-if="avTemplate[rule.itemname].threshold_dbfs" class="is-sub">阈值 {{ avTemplate[rule.itemname].threshold_dbfs }}</span>
                <span v-if="avTemplate[rule.itemname].duration_ms" class="is-sub">时长 {{ avTemplate[rule.itemname].duration_ms }} ms</span>
              </span>
            </div>
          </div>
        </div>
      </div>
    </template>
    <slot></slot>
  </tippy>
</template>
