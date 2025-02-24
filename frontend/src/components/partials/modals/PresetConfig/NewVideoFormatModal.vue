<script setup lang="ts">
import { toTypedSchema } from "@vee-validate/zod";
import { useForm } from "vee-validate";
import { z } from "zod";
import { useNotyf } from "@src/composable/useNotyf";
import { VideoFormatTypes, VideoResolutions, VideoFrameRates, VideoGammas, VideoColorSpaces, VideoEncodes, defVideoFormat } from './Consts';
import { useUserSession } from "@src/stores/userSession";

const usStore = useUserSession();
const videoFormats = computed(() => usStore.settings.video_formats || []);
const notyf = useNotyf();
const opened = ref(false);
const indexRef = ref(-1);
let callbacks: any = {}
const form = ref<VideoFormat>(defVideoFormat())

const zodSchema = z.object({
  name: z.string({ required_error: "请输入模板名称" })
    .nonempty("请输入模板名称")
    .regex(/^[\u4e00-\u9fa5a-zA-Z0-9-]+$/, { message: "请输入由中英文和横杠组成的名称" })
    .trim()
    .refine(value => {
      if (indexRef.value >= 0) return true
      return !videoFormats.value.some((item) => item.name === value)
    }, { message: "模板名称已存在" }),
  format: z.string({ required_error: "请选择格式类型" }),
  resolution: z.string({ required_error: "请选择分辨率" }),
  frameRate: z.string({ required_error: "请选择视频频率" }),
  gamma: z.string({ required_error: "请选择视频伽马" }),
  colorSpace: z.string({ required_error: "请选择视频色域" }),
  encodeFormat: z.string({ required_error: "请选择编码格式" }),
  bitRate: z.string({ required_error: "请输入视频码率" })
    .nonempty("请输入视频码率")
    .regex(/^\d+[kKmMgGtTpP]?bps$/, { message: "请输入合法的视频码率，比如10Mbps" })
    .trim(),
  bframe: z.number({ required_error: "请输入B帧数量", invalid_type_error: "请输入B帧数量" })
    .int()
    .min(0, { message: "B帧数量不能小于0" }),
  gop: z.number({ required_error: "请输入GOP长度", invalid_type_error: "请输入GOP长度" })
    .int()
});
const validationSchema = toTypedSchema(zodSchema);
const { handleSubmit, setFieldValue } = useForm({ validationSchema });

const loading = ref(false);
const handleCommit = handleSubmit(async () => {
  if (loading.value) return;

  loading.value = true;
  let value: string
  if (indexRef.value < 0) {
    value = JSON.stringify([
      form.value,
      ...videoFormats.value
    ])
  } else {
    const formats = videoFormats.value
    formats[indexRef.value] = form.value
    value = JSON.stringify(formats)
  }
  const res = await usStore.$updateSettings({ key: 'video_formats', value });
  loading.value = false;
  if (res) {
    opened.value = false;
    notyf.success(`${indexRef.value < 0 ? '创建' : '保存'}视频格式成功`);
    callbacks.success?.();
  } else {
    notyf.error(`${indexRef.value < 0 ? '创建' : '保存'}视频格式失败`);
  }
});

const resolutionList = ref(VideoResolutions);
const resolutionRef = ref<any>(null)
function onSearchInput(e: KeyboardEvent) {
  const input = e.target as HTMLInputElement
  if (!/^\d+\*\d+$/.test(input.value.trim())) {
    input.value = ''
    resolutionRef.value.blur()
    return
  }
  const existed = resolutionList.value.find((item) => item.includes(input.value))
  if (!existed) {
    resolutionList.value.unshift(input.value)
    resolutionRef.value.select(input.value)
    resolutionRef.value.blur()
  }
}

function syncForm() {
  for (let key of Object.keys(form.value) as (keyof VideoFormat)[]) {
    setFieldValue(key, form.value[key])
  }
}

syncForm()

useListener(Signal.OpenNewVideoFormat, (payload: {
  _callback?: any,
  index?: number,
}) => {
  opened.value = true
  indexRef.value = payload.index === undefined ? -1 : payload.index
  if (indexRef.value < 0) {
    form.value = defVideoFormat()
  } else {
    form.value = {
      ...videoFormats.value[indexRef.value]
    }
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
    size="medium"
    :open="opened"
    title="视频格式模版"
    actions="right"
    cancel-label="取消"
    dialog-class="preset-modal"
    @submit.prevent="handleCommit"
    @close="opened = false"
  >
    <template #content>
      <div class="modal-form">
        <div class="columns is-multiline">
          <div class="column is-6">
            <VField id="name" v-slot="{ field }">
              <VLabel>模版名称</VLabel>
              <VControl>
                <VInput
                  v-model="form.name"
                  type="text"
                  placeholder="支持中英文与横杠"
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
            <VField id="format" v-slot="{ field }">
              <VLabel>格式类型</VLabel>
              <VControl>
                <Multiselect
                  v-model="form.format"
                  :options="VideoFormatTypes"
                  placeholder="选择格式类型"
                  @change="(val: any) => field?.setValue(val)"
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
            <VField id="resolution" v-slot="{ field }">
              <VLabel>分辨率</VLabel>
              <VControl>
                <Multiselect
                  ref="resolutionRef"
                  v-model="form.resolution"
                  :searchable="true"
                  :options="VideoResolutions"
                  placeholder="自定义分辨率参数"
                  @change="(val: any) => field?.setValue(val)"
                  @keyup.enter.prevent="onSearchInput($event)"
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
            <VField id="frameRate" v-slot="{ field }">
              <VLabel>视频频率</VLabel>
              <VControl>
                <Multiselect
                  v-model="form.frameRate"
                  :options="VideoFrameRates"
                  placeholder="选择格式类型"
                  @change="(val: any) => field?.setValue(val)"
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
            <VField id="gamma" v-slot="{ field }">
              <VLabel>视频伽马</VLabel>
              <VControl>
                <Multiselect
                  v-model="form.gamma"
                  :options="VideoGammas"
                  placeholder="选择视频伽马"
                  @change="(val: any) => field?.setValue(val)"
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
            <VField id="colorSpace" v-slot="{ field }">
              <VLabel>视频色域</VLabel>
              <VControl>
                <Multiselect
                  v-model="form.colorSpace"
                  :options="VideoColorSpaces"
                  placeholder="选择视频色域"
                  @change="(val: any) => field?.setValue(val)"
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
            <VField id="encodeFormat" v-slot="{ field }">
              <VLabel>编码格式</VLabel>
              <VControl>
                <Multiselect
                  v-model="form.encodeFormat"
                  :options="VideoEncodes"
                  placeholder="选择编码格式"
                  @change="(val: any) => field?.setValue(val)"
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
            <VField id="bitRate" v-slot="{ field }">
              <VLabel>视频码率</VLabel>
              <VControl>
                <VInput
                  v-model="form.bitRate"
                  type="text"
                  placeholder="例如: 10Mbps"
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
            <VField id="bframe" v-slot="{ field }">
              <VLabel>B帧数量</VLabel>
              <VControl>
                <VInputNumber
                  v-model="form.bframe"
                  centered
                  :min="0"
                  :step="1"
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
            <VField id="gop" v-slot="{ field }">
              <VLabel>GOP长度</VLabel>
              <VControl>
                <VInputNumber
                  v-model="form.gop"
                  centered
                  :min="0"
                  :step="1"
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
.modal.preset-modal {

  .modal-content,
  .modal-card,
  .modal-card-body {
    overflow: visible;
  }
}
</style>
