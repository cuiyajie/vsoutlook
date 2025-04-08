<script lang="ts" setup>
import { useUserSession } from "@src/stores/userSession";
import { VideoReviewKeyName, AudioReviewKeyName, VideoReviewKeys, AudioReviewKeys, defTechReview } from "./Consts";

const props = defineProps<{
  value?: TechReview,
  name?: string
}>()

const usStore = useUserSession()
const techReviews = computed(() => usStore.settings.tech_reviews || [])

const techReview = computed(() => {
  if (props.value) return props.value
  if (props.name) {
    return techReviews.value.find(f => f.name === props.name) || defTechReview()
  }
  return defTechReview()
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
              {{ techReview.name }}
            </span>
          </div>
          <div class="info-block-line has-info" style="margin-bottom: 12px; margin-top: 12px;">
            <h4 class="dark-inverted" style="font-weight: 600;">视频技审</h4>
          </div>
          <div v-for="vk in VideoReviewKeys.filter(vk => techReview[vk])" :key="vk" class="info-block-line has-info">
            <h4 class="dark-inverted">{{ VideoReviewKeyName[vk] }}</h4>
            <div>
              <span class="info-right">
                <span v-if="techReview[vk].threshold_percentage" class="is-sub">阈值 {{ techReview[vk].threshold_percentage }}</span>
                <span v-if="techReview[vk].duration_frames" class="is-sub">时长 {{ techReview[vk].duration_frames }} 帧</span>
                <span v-if="techReview[vk].duration_ms" class="is-sub">时长 {{ techReview[vk].duration_ms }} ms</span>
              </span>
            </div>
          </div>
          <div class="info-block-line has-info" style="margin-bottom: 12px;">
            <h4 class="dark-inverted" style="font-weight: 600;">音频技审</h4>
          </div>
          <div v-for="ak in AudioReviewKeys.filter(ak => techReview[ak])" :key="ak" class="info-block-line has-info">
            <h4 class="dark-inverted">{{ AudioReviewKeyName[ak] }}</h4>
            <div>
              <span class="info-right">
                <span v-if="techReview[ak].detect_channels" class="is-sub">声道 {{ techReview[ak].detect_channels }}</span>
                <span v-if="techReview[ak].duration_frames" class="is-sub">时长 {{ techReview[ak].duration_frames }} 帧</span>
                <span v-if="techReview[ak].threshold_dbfs" class="is-sub">阈值 {{ techReview[ak].threshold_dbfs }}</span>
                <span v-if="techReview[ak].duration_ms" class="is-sub">时长 {{ techReview[ak].duration_ms }} ms</span>
              </span>
            </div>
          </div>
        </div>
      </div>
    </template>
    <slot></slot>
  </tippy>
</template>
