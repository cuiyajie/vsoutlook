<script setup lang="ts">
import { toTypedSchema } from "@vee-validate/zod";
import { useForm } from "vee-validate";
import { z } from "zod";
import { useNotyf } from "@src/composable/useNotyf";
import { useUserSession } from "@src/stores/userSession";
import { isValidNumberRanges } from "./Utils";
import Draggable from "vuedraggable";
import {
  VideoReviewKeyName,
  AudioReviewKeyName,
  VideoReviewKeys,
  AudioReviewKeys,
} from "./Consts";

type AuditAlarmForm = Pick<AuditAlarmRule, "rule_name"> & AuditAlarmRule["av_alarm"]["audio_mute_rule"];

const usStore = useUserSession();
const techReviews = computed(() => usStore.settings.tech_reviews || []);
const auditAlarmRules = computed(() => usStore.settings.audit_alarm_rules || []);
const notyf = useNotyf();
const opened = ref(false);
const indexRef = ref(-1);
let callbacks: any = {};
const form = ref<AuditAlarmForm>({
  rule_name: "",
  condition_any: "",
  condition_all: "",
});

const techReviewName = ref("");
const techReview = computed(() => techReviews.value.find((tr) => tr.name === techReviewName.value));
const videoReviewRules = ref<TVideoRuleKey[]>([])
const audioReviewRules = ref<TAudioRuleKey[]>([])
watch(techReview, () => {
  videoReviewRules.value = VideoReviewKeys.filter((key) => techReview.value?.[key])
  audioReviewRules.value = AudioReviewKeys.filter((key) => techReview.value?.[key])
}, { immediate: true })

const zodSchema = z.object({
  rule_name: z
    .string({ required_error: "请输入模板名称" })
    .nonempty("请输入模板名称")
    .trim()
    .refine(
      (value) => {
        if (indexRef.value >= 0) return true;
        return !auditAlarmRules.value.some((item) => item.rule_name === value);
      },
      { message: "模板名称已存在" }
    ),
  condition_any: z
    .string({ required_error: "请输入符合任意声道静音的声道范围" })
    .nonempty("请输入符合任意声道静音的声道范围")
    .trim()
    .refine(
      (value) => {
        return isValidNumberRanges(value);
      },
      { message: "请输入例如 1-3,4,5 的范围" }
    ),
  condition_all: z
    .string({ required_error: "请输入符合所有声道静音的声道范围" })
    .nonempty("请输入符合所有声道静音的声道范围")
    .trim()
    .refine(
      (value) => {
        return isValidNumberRanges(value);
      },
      { message: "请输入例如 1-3,4,5 的范围" }
    ),
});
const validationSchema = toTypedSchema(zodSchema);
const { handleSubmit, setFieldValue } = useForm({ validationSchema });

const loading = ref(false);
const handleCommit = handleSubmit(async () => {
  if (loading.value) return;

  loading.value = true;
  let value: string;
  const formValue = {
    rule_name: form.value.rule_name,
    av_alarm: {
      audit_template_name: techReviewName.value,
      video_alarm_priority: videoReviewRules.value.map((rule, i) => ({
        itemname: rule,
        priority: i + 1,
      })),
      audio_alarm_priority: audioReviewRules.value.map((rule, i) => ({
        itemname: rule,
        priority: i + 1,
      })),
      audio_mute_rule: {
        condition_any: form.value.condition_any,
        condition_all: form.value.condition_all,
      },
    },
  } as AuditAlarmRule;
  if (indexRef.value < 0) {
    value = JSON.stringify([formValue, ...auditAlarmRules.value]);
  } else {
    const formats = auditAlarmRules.value;
    formats[indexRef.value] = formValue;
    value = JSON.stringify(formats);
  }
  const res = await usStore.$updateSettings({ key: "audit_alarm_rules", value });
  loading.value = false;
  if (res) {
    opened.value = false;
    notyf.success(`${indexRef.value < 0 ? "创建" : "保存"}报警规则模板成功`);
    callbacks.success?.();
  } else {
    notyf.error(`${indexRef.value < 0 ? "创建" : "保存"}报警规则模板失败`);
  }
});

function syncForm() {
  for (let key of Object.keys(form.value) as (keyof AuditAlarmForm)[]) {
    setFieldValue(key, form.value[key]);
  }
}

syncForm();

useListener(Signal.OpenNewAuditAlarm, (payload: { _callback?: any; index?: number }) => {
  opened.value = true;
  indexRef.value = payload.index === undefined ? -1 : payload.index;
  const tr = auditAlarmRules.value[indexRef.value];
  if (tr) {
    form.value = {
      rule_name: tr.rule_name,
      condition_any: tr.av_alarm.audio_mute_rule.condition_any,
      condition_all: tr.av_alarm.audio_mute_rule.condition_all,
    };
    techReviewName.value = tr.av_alarm.audit_template_name;
    videoReviewRules.value = tr.av_alarm.video_alarm_priority.map((rule) => rule.itemname as TVideoRuleKey);
    audioReviewRules.value = tr.av_alarm.audio_alarm_priority.map((rule) => rule.itemname as TAudioRuleKey);
  } else {
    form.value = {
      rule_name: "",
      condition_any: "",
      condition_all: "",
    };
    techReviewName.value = "";
    videoReviewRules.value = [];
    audioReviewRules.value = [];
  }

  callbacks = payload._callback || {};
  syncForm();
});
</script>
<template>
  <VModal
    is="form"
    method="post"
    novalidate
    size="big"
    :open="opened"
    title="报警规则模版"
    actions="right"
    cancel-label="取消"
    dialog-class="preset-modal audit-alarm-modal"
    @submit.prevent="handleCommit"
    @close="opened = false"
  >
    <template #content>
      <div class="modal-form review-form">
        <div class="columns is-multiline">
          <div class="column is-6">
            <VField id="rule_name" v-slot="{ field }">
              <VLabel>模版名称</VLabel>
              <VControl>
                <VInput v-model="form.rule_name" type="text" placeholder="请输入模版名称" />
                <p v-if="field?.errorMessage" class="help is-danger">
                  {{ field.errorMessage }}
                </p>
              </VControl>
            </VField>
          </div>
          <div class="column is-6">
            <VField>
              <VLabel>视音频技审模板</VLabel>
              <VControl>
                <Multiselect
                  v-model="techReviewName"
                  class="tippy-select"
                  placeholder="选择视音频技审参数模板"
                  no-options-text="还没有配置视音频技审参数模板"
                  value-prop="name"
                  label="name"
                  :searchable="false"
                  :style="{'--ms-max-height': '245px'}"
                  :options="techReviews"
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
                    <TechReviewTooltip :tech-review="option">
                      <i class="iconify" data-icon="feather:alert-circle" aria-hidden="true"></i>
                    </TechReviewTooltip>
                  </template>
                </Multiselect>
              </VControl>
            </VField>
          </div>
        </div>
        <div class="columns is-multiline">
          <div class="column is-6">
            <VField id="condition_any" v-slot="{ field }">
              <VLabel>任意声道静音</VLabel>
              <VControl>
                <VInput
                  v-model="form.condition_any"
                  type="text"
                  placeholder="请输入形如 1-3,7 声道范围"
                />
                <p v-if="field?.errorMessage" class="help is-danger">
                  {{ field.errorMessage }}
                </p>
              </VControl>
            </VField>
          </div>
          <div class="column is-6">
            <VField id="condition_all" v-slot="{ field }">
              <VLabel>所有声道静音</VLabel>
              <VControl>
                <VInput
                  v-model="form.condition_all"
                  type="text"
                  placeholder="请输入形如 6,8-16 声道范围"
                />
                <p v-if="field?.errorMessage" class="help is-danger">
                  {{ field.errorMessage }}
                </p>
              </VControl>
            </VField>
          </div>
        </div>
        <expand-transition>
          <div v-if="videoReviewRules.length > 0" class="columns is-multiline">
            <div class="column is-12">
              <VField id="videoRules" horizontal>
                <VLabel>视频报警规则</VLabel>
              </VField>
            </div>
            <div class="column is-12">
              <div class="page-content-inner">
                <!--List-->
                <div class="list-view list-view-v1 list-review">
                  <div class="list-view-inner">
                    <Draggable
                      v-model="videoReviewRules"
                      :component-data="{
                        tag: 'div',
                        type: 'transition-group',
                        name: 'list-complete',
                      }"
                      handle=".drag-area"
                      item-key="value"
                    >
                      <template #item="{ element, index }: { element: TVideoRuleKey, index: number }">
                        <div class="list-view-item">
                          <div class="list-view-item-inner">
                            <div class="index-no">级别 {{ index + 1 }}</div>
                            <div class="meta-name">{{ VideoReviewKeyName[element] }}</div>
                            <VField v-if="techReview?.[element].threshold_percentage !== undefined" horizontal>
                              <VLabel>阈值：{{ techReview[element].threshold_percentage }}</VLabel>
                            </VField>
                            <VField
                              v-if="techReview?.[element].duration_frames !== undefined"
                              horizontal
                              addons
                            >
                              <VLabel>持续时长：{{ techReview[element].duration_frames }} 帧</VLabel>
                            </VField>
                            <VField
                              v-if="techReview?.[element].duration_ms !== undefined"
                              horizontal
                              addons
                            >
                              <VLabel>时长：{{ techReview[element].duration_ms }} ms</VLabel>
                            </VField>
                            <div className="drag-area">
                              <VIconButton icon="feather:menu" circle />
                            </div>
                          </div>
                        </div>
                      </template>
                    </Draggable>
                  <!--Item-->
                  </div>
                </div>
              </div>
            </div>
          </div>
        </expand-transition>
        <expand-transition>
          <div v-if="audioReviewRules.length > 0" class="columns is-multiline">
            <div class="column is-12">
              <VField id="audioRules" horizontal>
                <VLabel>音频报警规则</VLabel>
              </VField>
            </div>
            <div class="column is-12">
              <div class="page-content-inner">
                <!--List-->
                <div class="list-view list-view-v1 list-review list-audio-review">
                  <div class="list-view-inner">
                    <Draggable
                      v-model="audioReviewRules"
                      :component-data="{
                        tag: 'div',
                        type: 'transition-group',
                        name: 'list-complete',
                      }"
                      handle=".drag-area"
                      item-key="value"
                    >
                      <template #item="{ element, index }: { element: TAudioRuleKey, index: number }">
                        <div
                          class="list-view-item"
                          :class="element === 'audio_low' && 'list-view-too-low'"
                        >
                          <div class="list-view-item-inner">
                            <div class="index-no">级别 {{ index + 1 }}</div>
                            <div class="meta-name">{{ AudioReviewKeyName[element] }}</div>
                            <VField v-if="techReview?.[element].detect_channels !== undefined" horizontal>
                              <VLabel>声道：{{ techReview[element].detect_channels }}</VLabel>
                            </VField>
                            <VField v-if="techReview?.[element].threshold_dbfs !== undefined" horizontal>
                              <VLabel>阈值：{{ techReview[element].threshold_dbfs }}</VLabel>
                            </VField>
                            <VField
                              v-if="techReview?.[element].duration_frames !== undefined"
                              horizontal
                              addons
                            >
                              <VLabel>持续时长：{{ techReview[element].duration_frames }} 帧</VLabel>
                            </VField>
                            <VField
                              v-if="techReview?.[element].duration_ms !== undefined"
                              horizontal
                              addons
                            >
                              <VLabel>时长：{{ techReview[element].duration_ms }} ms</VLabel>
                            </VField>
                            <div className="drag-area">
                              <VIconButton icon="feather:menu" circle />
                            </div>
                          </div>
                        </div>
                      </template>
                    </Draggable>
                  <!--Item-->
                  </div>
                </div>
              </div>
            </div>
          </div>
        </expand-transition>
      </div>
    </template>
    <template #action>
      <VButton
        type="submit"
        color="primary"
        raised
        :disabled="loading"
        :loading="loading"
      >
        保存
      </VButton>
    </template>
  </VModal>
</template>
<style lang="scss">
@import "@src/scss/abstracts/all";

.modal.preset-modal.audit-alarm-modal {
  .modal-content,
  .modal-card,
  .modal-card-body {
    overflow: auto;
  }
}

.modal-form.review-form {
  .field {
    &.is-horizontal {
      .label:not(:last-child) {
        margin-bottom: 0;
      }

      .label {
        display: flex;
        align-items: center;
        margin-right: 1em;
      }
    }
  }
}

.list-review {
  .page-placeholder {
    min-height: 120px;

    .placeholder-content p {
      font-size: 1rem;
    }
  }
}

.list-view-v1 {
  .list-view-item {
    @include vuero-r-card;

    margin-bottom: 16px;
    padding: 16px;

    &.list-view-too-low {
      .field.is-horizontal {
        .input {
          width: 100px;
        }

        &.has-addons {
          .field-addon-body {
            .input {
              width: 100px;
            }
          }
        }
      }
    }

    .list-view-item-inner {
      display: flex;
      align-items: center;
      font-size: 0.875rem;
      gap: 24px;

      .drag-area {
        flex: 1;
        display: flex;
        align-items: center;
        justify-content: end;

        .button:focus:not(:active) {
          box-shadow: none;
        }
      }
    }

    .field.is-horizontal {
      margin-bottom: 0 !important;

      .label {
        margin-right: 0;
      }

      .input {
        font-size: 0.875rem;
        width: 160px;

        &.v-input {
          height: 32px;
        }
      }

      .v-input-number__button {
        width: 32px;
        height: 32px;
        z-index: 10;
      }

      &.has-addons {
        .field-addon-body {
          .control:nth-child(2) .input {
            border-top-left-radius: var(--radius);
            border-bottom-left-radius: var(--radius);
          }

          .input {
            width: 120px;
          }
        }
      }

      .field-addon-body {
        > label {
          font-family: var(--font);
          font-size: 0.9rem;
          color: var(--light-text) !important;
          font-weight: 400;
        }

        .control .button {
          height: 32px;
          padding: 8px 12px;
        }
      }
    }
  }
}

.is-dark {
  .list-view-v1 {
    .list-view-item {
      @include vuero-card--dark;

      .list-view-item-inner {
      }
    }
  }
}

@media only screen and (width <= 767px) {
  .list-view-v1 {
    .list-view-item {
      .list-view-item-inner {
        position: relative;
        flex-direction: column;
      }
    }
  }
}

@media only screen and (width >= 768px) and (width <= 1024px) and (orientation: portrait) {
  .list-view-v1 {
    display: flex;
    flex-wrap: wrap;

    .list-view-item {
      margin: 10px;
      width: calc(50% - 20px);

      .list-view-item-inner {
        position: relative;
        flex-direction: column;
      }
    }
  }
}
</style>
