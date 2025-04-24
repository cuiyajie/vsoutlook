<script setup lang="ts">
import { toTypedSchema } from "@vee-validate/zod";
import { useForm } from "vee-validate";
import { z } from "zod";
import { useNotyf } from "@src/composable/useNotyf";
import { useUserSession } from "@src/stores/userSession";
import {
  VideoReviewKeys,
  defVideoReviewRules,
  defAudioReviewRules,
  AudioReviewKeys,
  VideoReviewKeyName,
  AudioReviewKeyName,
  defVideoReviewKeys,
  defAudioReviewKeys,
} from "./Consts";

type TechReviewForm = Pick<TechReview, "name">;

const usStore = useUserSession();
const techReviews = computed(() => usStore.settings.tech_reviews || []);
const notyf = useNotyf();
const opened = ref(false);
const indexRef = ref(-1);
let callbacks: any = {};
const form = ref<TechReviewForm>({
  name: "",
});
const videoReviewKeys = ref<Record<TVideoRuleKey, boolean>>(defVideoReviewKeys());
const videoReviewRules = ref(defVideoReviewRules());
const audioReviewKeys = ref<Record<TAudioRuleKey, boolean>>(defAudioReviewKeys());
const audioReviewRules = ref(defAudioReviewRules());

const zodSchema = z.object({
  name: z
    .string({ required_error: "请输入模板名称" })
    .nonempty("请输入模板名称")
    .trim()
    .refine(
      (value) => {
        if (indexRef.value >= 0) return true;
        return !techReviews.value.some((item) => item.name === value);
      },
      { message: "模板名称已存在" }
    )
});
const validationSchema = toTypedSchema(zodSchema);
const { handleSubmit, setFieldValue } = useForm({ validationSchema });

const loading = ref(false);
const handleCommit = handleSubmit(async () => {
  if (loading.value) return;

  loading.value = true;
  let value: string;
  const formValue = {
    ...form.value,
  } as TechReview;
  VideoReviewKeys.forEach(vkey => {
    if (videoReviewKeys.value[vkey]) {
      Object.assign(formValue, { [vkey]: videoReviewRules.value[vkey] })
    }
  });
  AudioReviewKeys.forEach(akey => {
    if (audioReviewKeys.value[akey]) {
      Object.assign(formValue, { [akey]: audioReviewRules.value[akey] })
    }
  });
  if (indexRef.value < 0) {
    value = JSON.stringify([formValue, ...techReviews.value]);
  } else {
    const formats = techReviews.value;
    formats[indexRef.value] = formValue;
    value = JSON.stringify(formats);
  }
  const res = await usStore.$updateSettings({ key: "tech_reviews", value });
  loading.value = false;
  if (res) {
    opened.value = false;
    notyf.success(`${indexRef.value < 0 ? "创建" : "保存"}视音频技审模板成功`);
    callbacks.success?.();
  } else {
    notyf.error(`${indexRef.value < 0 ? "创建" : "保存"}视音频技审模板失败`);
  }
});

function syncForm() {
  for (let key of Object.keys(form.value) as (keyof TechReviewForm)[]) {
    setFieldValue(key, form.value[key]);
  }
}

syncForm();

useListener(Signal.OpenNewTechReview, (payload: { _callback?: any; index?: number }) => {
  opened.value = true;
  indexRef.value = payload.index === undefined ? -1 : payload.index;
  const tr = techReviews.value[indexRef.value];
  if (tr) {
    form.value = { name: tr.name };
    const [videoChecked, videoRules] = VideoReviewKeys.reduce((acc, key) => {
      acc[0][key] = (tr as any)[key] !== undefined;
      if (acc[0][key]) {
        acc[1][key] = (tr as any)[key];
      }
      return acc;
    }, [
      {} as Record<TVideoRuleKey, boolean>,
      defVideoReviewRules()
    ]);
    videoReviewKeys.value = videoChecked;
    videoReviewRules.value = videoRules;
    const [audioChecked, audioRules] = AudioReviewKeys.reduce((acc, key) => {
      acc[0][key] = (tr as any)[key] !== undefined;
      if (acc[0][key]) {
        acc[1][key] = (tr as any)[key];
      }
      return acc;
    }, [
      {} as Record<TAudioRuleKey, boolean>,
      defAudioReviewRules()
    ]);
    audioReviewKeys.value = audioChecked;
    audioReviewRules.value = audioRules;
  } else {
    form.value = { name: "" };
    videoReviewKeys.value = defVideoReviewKeys();
    videoReviewRules.value = defVideoReviewRules();
    audioReviewKeys.value = defAudioReviewKeys();
    audioReviewRules.value = defAudioReviewRules();
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
    title="视音频技审模版"
    actions="right"
    cancel-label="取消"
    dialog-class="preset-modal tech-review-modal"
    @submit.prevent="handleCommit"
    @close="opened = false"
  >
    <template #content>
      <div class="modal-form review-form">
        <div class="columns is-multiline">
          <div class="column is-6">
            <VField id="name" v-slot="{ field }">
              <VLabel>模版名称</VLabel>
              <VControl>
                <VInput v-model="form.name" type="text" placeholder="请输入模版名称" />
                <p v-if="field?.errorMessage" class="help is-danger">
                  {{ field.errorMessage }}
                </p>
              </VControl>
            </VField>
          </div>
        </div>
        <div class="form-fieldset">
          <div class="fieldset-heading">
            <h5>视频技审及规则</h5>
          </div>
          <div v-for="rk in VideoReviewKeys" :key="rk" class="form-field-row">
            <VField class="is-checkbox">
              <VControl raw subcontrol>
                <VCheckbox
                  v-model="videoReviewKeys[rk]"
                  :value="rk"
                  :label="VideoReviewKeyName[rk]"
                  color="primary"
                  circle
                />
              </VControl>
            </VField>
            <VField v-if="videoReviewRules[rk].threshold_percentage !== undefined" horizontal>
              <VLabel>阈值：</VLabel>
              <VControl>
                <VInputNumber
                  v-model="videoReviewRules[rk].threshold_percentage"
                  centered
                  :min="0"
                  :step="0.001"
                  placeholder="请输入阈值"
                />
              </VControl>
            </VField>
            <VField
              v-if="videoReviewRules[rk].duration_frames !== undefined"
              horizontal
              addons
            >
              <VLabel>持续时长：</VLabel>
              <VControl>
                <VInputNumber
                  v-model="videoReviewRules[rk].duration_frames"
                  centered
                  :min="0"
                  :step="1"
                  placeholder="请输入持续时长"
                />
              </VControl>
              <VControl>
                <VButton static>帧</VButton>
              </VControl>
            </VField>
            <VField
              v-if="videoReviewRules[rk].duration_ms !== undefined"
              horizontal
              addons
            >
              <VLabel>时长：</VLabel>
              <VControl>
                <VInputNumber
                  v-model="videoReviewRules[rk].duration_ms"
                  centered
                  :min="0"
                  :step="1"
                  placeholder="请输入时长"
                />
              </VControl>
              <VControl>
                <VButton static>ms</VButton>
              </VControl>
            </VField>
          </div>
        </div>
        <div class="form-fieldset">
          <div class="fieldset-heading">
            <h5>音频技审及规则</h5>
          </div>
          <div v-for="ak in AudioReviewKeys" :key="ak" class="form-field-row">
            <VField class="is-checkbox">
              <VControl raw subcontrol>
                <VCheckbox
                  v-model="audioReviewKeys[ak]"
                  :value="ak"
                  :label="AudioReviewKeyName[ak]"
                  color="primary"
                  circle
                />
              </VControl>
            </VField>
            <VField v-if="audioReviewRules[ak].detect_channels !== undefined" horizontal>
              <VLabel>声道：</VLabel>
              <VControl>
                <VInput
                  v-model="audioReviewRules[ak].detect_channels"
                  type="text"
                  placeholder="请输入声道范围"
                />
              </VControl>
            </VField>
            <VField v-if="audioReviewRules[ak].threshold_dbfs !== undefined" horizontal>
              <VLabel>阈值：</VLabel>
              <VControl>
                <VInputNumber
                  v-model="audioReviewRules[ak].threshold_dbfs"
                  centered
                  :min="-Infinity"
                  :step="1"
                  placeholder="请输入阈值"
                />
              </VControl>
            </VField>
            <VField
              v-if="audioReviewRules[ak].duration_frames !== undefined"
              horizontal
              addons
            >
              <VLabel>持续时长：</VLabel>
              <VControl>
                <VInputNumber
                  v-model="audioReviewRules[ak].duration_frames"
                  centered
                  :min="0"
                  :step="1"
                  placeholder="请输入持续时长"
                />
              </VControl>
              <VControl>
                <VButton static>帧</VButton>
              </VControl>
            </VField>
            <VField
              v-if="audioReviewRules[ak].duration_ms !== undefined"
              horizontal
              addons
            >
              <VLabel>时长：</VLabel>
              <VControl>
                <VInputNumber
                  v-model="audioReviewRules[ak].duration_ms"
                  centered
                  :min="0"
                  :step="1"
                  placeholder="请输入时长"
                />
              </VControl>
              <VControl>
                <VButton static>ms</VButton>
              </VControl>
            </VField>
          </div>
        </div>
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

.modal.preset-modal.tech-review-modal {

  &.modal.is-big .modal-content {
    max-width: 960px;
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

  .form-fieldset {
    max-width: 100%;
  }
}

.form-field-row {
  display: flex;
  align-items: center;
  gap: 24px;
  margin-top: 24px;

  .field.is-checkbox {
    margin-bottom: 0 !important;

    .checkbox {
      padding: 0 16px 0 0;
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
}
</style>
