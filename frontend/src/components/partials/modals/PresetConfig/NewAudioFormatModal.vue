<script setup lang="ts">
import { toTypedSchema } from "@vee-validate/zod";
import { useForm } from "vee-validate";
import { z } from "zod";
import { useNotyf } from "@src/composable/useNotyf";
import { AudioQuantBits, DefaultAudioQuantBit, AudioSampleRates, DefaultAudioSampleRate, AudioEncodes, DefaultAudioEncode, AudioTxInterval, DefaultAudioTxInterval, defAudioFormat } from './Consts';
import { useUserSession } from "@src/stores/userSession";

const usStore = useUserSession();
const audioFormats = computed(() => usStore.settings.audio_formats || []);
const notyf = useNotyf();
const opened = ref(false);
const indexRef = ref(-1);
let callbacks: any = {}
const form = ref<AudioFormat>(defAudioFormat())

const zodSchema = z.object({
  name: z.string({ required_error: "请输入模板名称" })
    .nonempty("请输入模板名称")
    .regex(/^[\u4e00-\u9fa5a-zA-Z0-9-]+$/, { message: "请输入由中英文和横杠组成的名称" })
    .trim()
    .refine(value => {
      if (indexRef.value >= 0) return true
      return !audioFormats.value.some((item) => item.name === value)
    }, { message: "模板名称已存在" }),
  channels: z.number({ required_error: "请输入声道数量", invalid_type_error: "请输入声道数量" })
    .int()
    .min(0, { message: "声道数量不能小于0" }),
  quantBits: z.number({ required_error: "请选择量化 bit" }),
  sampleRate: z.number({ required_error: "请选择采样率" }),
  encodeFormat: z.string({ required_error: "请选择编码格式" }),
  txInterval: z.number({ required_error: "请选择发包间隔" }),
  bitRate: z.string({ required_error: "请输入音频码率" })
    .nonempty("请输入音频码率")
    .regex(/^\d+[kKmMgGtTpP]?bps$/, { message: "请输入合法的音频码率，比如100000bps" })
    .trim(),
  childType: z.string().default('-')
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
      ...audioFormats.value
    ])
  } else {
    const formats = audioFormats.value
    formats[indexRef.value] = form.value
    value = JSON.stringify(formats)
  }
  const res = await usStore.$updateSettings({ key: 'audio_formats', value });
  loading.value = false;
  if (res) {
    opened.value = false;
    notyf.success(`${indexRef.value < 0 ? '创建' : '保存'}音频格式成功`);
    callbacks.success?.();
  } else {
    notyf.error(`${indexRef.value < 0 ? '创建' : '保存'}音频格式失败`);
  }
});

function syncForm() {
  for (let key of Object.keys(form.value) as (keyof AudioFormat)[]) {
    setFieldValue(key, form.value[key])
  }
}

syncForm()

useListener(Signal.OpenNewAudioFormat, (payload: {
  _callback?: any,
  index?: number,
}) => {
  opened.value = true
  indexRef.value = payload.index === undefined ? -1 : payload.index
  if (indexRef.value < 0) {
    form.value = defAudioFormat()
  } else {
    form.value = {
      ...audioFormats.value[indexRef.value]
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
    title="音频格式模板"
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
            <VField id="channels" v-slot="{ field }">
              <VLabel>声道数量</VLabel>
              <VControl>
                <VInputNumber
                  v-model="form.channels"
                  centered
                  :min="1"
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
            <VField id="quantBits" v-slot="{ field }">
              <VLabel>量化 bit</VLabel>
              <VControl>
                <Multiselect
                  v-model="form.quantBits"
                  :options="AudioQuantBits"
                  placeholder="选择量化 bit"
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
            <VField id="sampleRate" v-slot="{ field }">
              <VLabel>采样率</VLabel>
              <VControl>
                <Multiselect
                  v-model="form.sampleRate"
                  :options="AudioSampleRates"
                  placeholder="选择采样率"
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
                  :options="AudioEncodes"
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
            <VField id="gamma" v-slot="{ field }">
              <VLabel>发包间隔</VLabel>
              <VControl>
                <Multiselect
                  v-model="form.txInterval"
                  :options="AudioTxInterval"
                  placeholder="选择发包间隔"
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
              <VLabel>音频码率</VLabel>
              <VControl>
                <VInput
                  v-model="form.bitRate"
                  type="text"
                  placeholder="例如: 100000bps"
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
</style>
