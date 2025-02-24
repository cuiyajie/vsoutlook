<script setup lang="ts">
import { toTypedSchema } from "@vee-validate/zod";
import { useForm } from "vee-validate";
import { z } from "zod";
import { useNotyf } from "@src/composable/useNotyf";
import { useUserSession } from "@src/stores/userSession";
import { isValidNumberRanges } from './Utils';
import Draggable from "vuedraggable";
import { VideoReviewKeys, defVideoReviewRules, defAudioReviewRules, AudioReviewKeys } from './Consts';

type TechReviewForm = Pick<TechReview, 'name' | 'anyChannels' | 'allChannels'>

const usStore = useUserSession();
const techReviews = computed(() => usStore.settings.tech_reviews || []);
const notyf = useNotyf();
const opened = ref(false);
const indexRef = ref(-1);
let callbacks: any = {}
const form = ref<TechReviewForm>({
  name: '',
  anyChannels: '',
  allChannels: '',
})
const videoReviewKeys = ref([])
const videoReviewRules = ref(defVideoReviewRules())
const audioReviewKeys = ref([])
const audioReviewRules = ref(defAudioReviewRules())

const zodSchema = z.object({
  name: z.string({ required_error: "请输入模板名称" })
    .nonempty("请输入模板名称")
    .trim()
    .refine(value => {
      if (indexRef.value >= 0) return true
      return !techReviews.value.some((item) => item.name === value)
    }, { message: "模板名称已存在" }),
  anyChannels: z.string({ required_error: "请输入符合任意声道静音的声道范围" })
    .nonempty("请输入符合任意声道静音的声道范围")
    .trim()
    .refine(value => {
      return isValidNumberRanges(value)
    }, { message: "请输入例如 1-3,4,5 的范围" }),
  allChannels: z.string({ required_error: "请输入符合所有声道静音的声道范围" })
    .nonempty("请输入符合所有声道静音的声道范围")
    .trim()
    .refine(value => {
      return isValidNumberRanges(value)
    }, { message: "请输入例如 1-3,4,5 的范围" }),
});
const validationSchema = toTypedSchema(zodSchema);
const { handleSubmit, setFieldValue } = useForm({ validationSchema });

const loading = ref(false);
const handleCommit = handleSubmit(async () => {
  if (loading.value) return;

  loading.value = true;
  let value: string
  const formValue = {
    ...form.value,
    videoRules: videoReviewKeys.value.map((key: string) => videoReviewRules.value[key]),
    audioRules: audioReviewKeys.value.map((key: string) => audioReviewRules.value[key]),
  } as TechReview
  if (indexRef.value < 0) {
    value = JSON.stringify([
      formValue,
      ...techReviews.value
    ])
  } else {
    const formats = techReviews.value
    formats[indexRef.value] = formValue
    value = JSON.stringify(formats)
  }
  const res = await usStore.$updateSettings({ key: 'tech_reviews', value });
  loading.value = false;
  if (res) {
    opened.value = false;
    notyf.success(`${indexRef.value < 0 ? '创建' : '保存'}视音频计审模板成功`);
    callbacks.success?.();
  } else {
    notyf.error(`${indexRef.value < 0 ? '创建' : '保存'}视音频计审模板失败`);
  }
});

function syncForm() {
  for (let key of Object.keys(form.value) as (keyof TechReviewForm)[]) {
    setFieldValue(key, form.value[key])
  }
}

syncForm()

useListener(Signal.OpenNewTechReview, (payload: {
  _callback?: any,
  index?: number,
}) => {
  opened.value = true
  indexRef.value = payload.index === undefined ? -1 : payload.index
  const tr = techReviews.value[indexRef.value]
  if (tr) {
    form.value = {
      name: tr.name,
      anyChannels: tr.anyChannels,
      allChannels: tr.allChannels,
    }
    videoReviewKeys.value = tr.videoRules.map((rule) => rule.name)
    videoReviewRules.value = tr.videoRules.reduce((acc, rule) => {
      acc[rule.name] = rule
      return acc
    }, defVideoReviewRules())
    audioReviewKeys.value = tr.audioRules.map((rule) => rule.name)
    audioReviewRules.value = tr.audioRules.reduce((acc, rule) => {
      acc[rule.name] = rule
      return acc
    }, defAudioReviewRules())
  } else {
    form.value = {
      name: '',
      anyChannels: '',
      allChannels: '',
    }
    videoReviewKeys.value = []
    videoReviewRules.value = defVideoReviewRules()
    audioReviewKeys.value = []
    audioReviewRules.value = defAudioReviewRules()
  }

  callbacks = payload._callback || {}
  syncForm()
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
                <VInput
                  v-model="form.name"
                  type="text"
                  placeholder="请输入模版名称"
                />
                <p
                  v-if="field?.errorMessage"
                  class="help is-danger"
                >
                  {{ field.errorMessage }}
                </p>
              </VControl>
            </VField>
          </div>
          <div class="column is-12">
            <VField id="videoRules" horizontal>
              <VLabel>视频技审及规则</VLabel>
              <VField class="is-flex">
                <VControl v-for="rk in VideoReviewKeys" :key="rk" raw subcontrol>
                  <VCheckbox v-model="videoReviewKeys" :value="rk" :label="rk" color="primary" />
                </VControl>
              </VField>
            </VField>
          </div>
          <div class="column is-12">
            <div class="page-content-inner">
              <!--List-->
              <div class="list-view list-view-v1 list-review">
                <!--List Empty Search Placeholder -->
                <VPlaceholderPage
                  :class="[videoReviewKeys.length !== 0 && 'is-hidden']"
                  title=""
                  subtitle="暂时没有视频技审及规则"
                />

                <div class="list-view-inner">
                  <Draggable
                    v-model="videoReviewKeys"
                    :component-data="{
                      tag: 'div',
                      type: 'transition-group',
                      name: 'list-complete'
                    }"
                    handle=".drag-area"
                    item-key="value"
                  >
                    <template #item="{element, index}">
                      <div class="list-view-item">
                        <div class="list-view-item-inner">
                          <div class="index-no">级别 {{ index + 1 }}</div>
                          <div class="meta-name">{{ element }}</div>
                          <VField v-if="videoReviewRules[element].threshold" horizontal>
                            <VLabel>阈值：</VLabel>
                            <VControl>
                              <VInputNumber
                                v-model="videoReviewRules[element].threshold"
                                :min="0"
                                :step="0.001"
                                placeholder="请输入阈值"
                              />
                            </VControl>
                          </VField>
                          <VField v-if="videoReviewRules[element].duration" horizontal addons>
                            <VLabel>持续时长：</VLabel>
                            <VControl>
                              <VInputNumber
                                v-model="videoReviewRules[element].duration"
                                :min="0"
                                :step="1"
                                placeholder="请输入持续时长"
                              />
                            </VControl>
                            <VControl>
                              <VButton static>帧</VButton>
                            </VControl>
                          </VField>
                          <VField v-if="videoReviewRules[element].missingDuration" horizontal addons>
                            <VLabel>时长：</VLabel>
                            <VControl>
                              <VInputNumber
                                v-model="videoReviewRules[element].missingDuration"
                                :min="0"
                                :step="0.001"
                                placeholder="请输入时长"
                              />
                            </VControl>
                            <VControl>
                              <VButton static>ms</VButton>
                            </VControl>
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
          <div class="column is-12">
            <VField id="audioRules" horizontal>
              <VLabel>音频技审及规则</VLabel>
              <VField class="is-flex">
                <VControl v-for="ak in AudioReviewKeys" :key="ak" raw subcontrol>
                  <VCheckbox v-model="audioReviewKeys" :value="ak" :label="ak" color="primary" />
                </VControl>
              </VField>
            </VField>
          </div>
          <div class="column is-12">
            <div class="page-content-inner">
              <!--List-->
              <div class="list-view list-view-v1 list-review list-audio-review">
                <!--List Empty Search Placeholder -->
                <VPlaceholderPage
                  :class="[audioReviewKeys.length !== 0 && 'is-hidden']"
                  title=""
                  subtitle="暂时没有音频技审及规则"
                />

                <div class="list-view-inner">
                  <Draggable
                    v-model="audioReviewKeys"
                    :component-data="{
                      tag: 'div',
                      type: 'transition-group',
                      name: 'list-complete'
                    }"
                    handle=".drag-area"
                    item-key="value"
                  >
                    <template #item="{element, index}">
                      <div class="list-view-item" :class="element === '音频过低' && 'list-view-too-low'">
                        <div class="list-view-item-inner">
                          <div class="index-no">级别 {{ index + 1 }}</div>
                          <div class="meta-name">{{ element }}</div>
                          <VField v-if="audioReviewRules[element].channels" horizontal>
                            <VLabel>声道：</VLabel>
                            <VControl>
                              <VInput
                                v-model="audioReviewRules[element].channels"
                                type="text"
                                placeholder="请输入声道范围"
                              />
                            </VControl>
                          </VField>
                          <VField v-if="audioReviewRules[element].threshold" horizontal>
                            <VLabel>阈值：</VLabel>
                            <VControl>
                              <VInputNumber
                                v-model="audioReviewRules[element].threshold"
                                :min="0"
                                :step="0.001"
                                placeholder="请输入阈值"
                              />
                            </VControl>
                          </VField>
                          <VField v-if="audioReviewRules[element].duration" horizontal addons>
                            <VLabel>持续时长：</VLabel>
                            <VControl>
                              <VInputNumber
                                v-model="audioReviewRules[element].duration"
                                :min="0"
                                :step="1"
                                placeholder="请输入持续时长"
                              />
                            </VControl>
                            <VControl>
                              <VButton static>帧</VButton>
                            </VControl>
                          </VField>
                          <VField v-if="audioReviewRules[element].missingDuration" horizontal addons>
                            <VLabel>时长：</VLabel>
                            <VControl>
                              <VInputNumber
                                v-model="audioReviewRules[element].missingDuration"
                                :min="0"
                                :step="0.001"
                                placeholder="请输入时长"
                              />
                            </VControl>
                            <VControl>
                              <VButton static>ms</VButton>
                            </VControl>
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
          <div class="column is-6">
            <VField id="anyChannels" v-slot="{ field }">
              <VLabel>任意声道静音</VLabel>
              <VControl>
                <VInput
                  v-model="form.anyChannels"
                  type="text"
                  placeholder="请输入形如 1-3,7 声道范围"
                />
                <p
                  v-if="field?.errorMessage"
                  class="help is-danger"
                >
                  {{ field.errorMessage }}
                </p>
              </VControl>
            </VField>
          </div>
          <div class="column is-6">
            <VField id="allChannels" v-slot="{ field }">
              <VLabel>所有声道静音</VLabel>
              <VControl>
                <VInput
                  v-model="form.allChannels"
                  type="text"
                  placeholder="请输入形如 6,8-16 声道范围"
                />
                <p
                  v-if="field?.errorMessage"
                  class="help is-danger"
                >
                  {{ field.errorMessage }}
                </p>
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
@import '@src/scss/abstracts/all';

.modal.preset-modal.tech-review-modal {

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
      gap: 16px;

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
