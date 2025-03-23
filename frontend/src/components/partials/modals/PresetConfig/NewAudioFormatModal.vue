<script setup lang="ts">
import { toTypedSchema } from '@vee-validate/zod'
import { useForm } from 'vee-validate'
import { z } from 'zod'
import { useNotyf } from '@src/composable/useNotyf'
import {
  AudioBits,
  AudioCompressionFormats,
  AudioPacketTimeUs,
  defAudioFormat,
} from './Consts'
import { useUserSession } from '@src/stores/userSession'

const usStore = useUserSession()
const audioFormats = computed(() => usStore.settings.audio_formats || [])
const notyf = useNotyf()
const opened = ref(false)
const indexRef = ref(-1)
let callbacks: any = {}
const form = ref<AudioFormat>(defAudioFormat())

const zodSchema = z.object({
  name: z
    .string({ required_error: '请输入模板名称' })
    .nonempty('请输入模板名称')
    .regex(/^[\u4e00-\u9fa5a-zA-Z0-9-]+$/, { message: '请输入由中英文和横杠组成的名称' })
    .trim()
    .refine(
      (value) => {
        if (indexRef.value >= 0) return true
        return !audioFormats.value.some((item) => item.name === value)
      },
      { message: '模板名称已存在' }
    ),
  channels_number: z
    .number({ required_error: '请输入声道数量', invalid_type_error: '请输入声道数量' })
    .int()
    .min(0, { message: '声道数量不能小于0' }),
  bits: z.number({ required_error: '请选择量化 bit' }),
  frequency: z.number({ required_error: '请选择采样率' }),
  compression_format: z.string({ required_error: '请选择编码格式' }),
  packet_time_us: z.number({ required_error: '请选择发包间隔' }).int(),
  bitrate_bps: z
    .number({ required_error: '请输入音频码率' })
    .int(),
  compression_subtype: z.string(),
})
const validationSchema = toTypedSchema(zodSchema)
const { handleSubmit, setFieldValue } = useForm({ validationSchema })

const loading = ref(false)
const handleCommit = handleSubmit(async () => {
  if (loading.value) return

  loading.value = true
  let value: string
  if (indexRef.value < 0) {
    value = JSON.stringify([form.value, ...audioFormats.value])
  } else {
    const formats = audioFormats.value
    formats[indexRef.value] = form.value
    value = JSON.stringify(formats)
  }
  const res = await usStore.$updateSettings({ key: 'audio_formats', value })
  loading.value = false
  if (res) {
    opened.value = false
    notyf.success(`${indexRef.value < 0 ? '创建' : '保存'}音频格式成功`)
    callbacks.success?.()
  } else {
    notyf.error(`${indexRef.value < 0 ? '创建' : '保存'}音频格式失败`)
  }
})

function syncForm() {
  for (let key of Object.keys(form.value) as (keyof AudioFormat)[]) {
    setFieldValue(key, form.value[key])
  }
}

syncForm()

useListener(Signal.OpenNewAudioFormat, (payload: { _callback?: any; index?: number }) => {
  opened.value = true
  indexRef.value = payload.index === undefined ? -1 : payload.index
  if (indexRef.value < 0) {
    form.value = defAudioFormat()
  } else {
    form.value = {
      ...audioFormats.value[indexRef.value],
    }
  }

  callbacks = payload._callback || {}
  syncForm()
})
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
                <VInput v-model="form.name" type="text" placeholder="支持中英文与横杠" />
                <p v-if="field?.errorMessage" class="help is-danger">
                  {{ field.errorMessage }}
                </p>
              </VControl>
            </VField>
          </div>
          <div class="column is-6">
            <VField id="channels_number" v-slot="{ field }">
              <VLabel>声道数量</VLabel>
              <VControl>
                <VInputNumber v-model="form.channels_number" centered :min="1" :step="1" />
                <p v-if="field?.errorMessage" class="help is-danger">
                  {{ field.errorMessage }}
                </p>
              </VControl>
            </VField>
          </div>
          <div class="column is-6">
            <VField id="bits" v-slot="{ field }">
              <VLabel>量化 bit</VLabel>
              <VControl>
                <Multiselect
                  v-model="form.bits"
                  :options="AudioBits"
                  placeholder="选择量化 bit"
                  @change="(val: any) => field?.setValue(val)"
                />
                <p v-if="field?.errorMessage" class="help is-danger">
                  {{ field.errorMessage }}
                </p>
              </VControl>
            </VField>
          </div>
          <div class="column is-6">
            <VField id="frequency" v-slot="{ field }">
              <VLabel>采样率</VLabel>
              <VControl>
                <VInputNumber v-model="form.frequency" centered :min="0" :step="1" />
                <p v-if="field?.errorMessage" class="help is-danger">
                  {{ field.errorMessage }}
                </p>
              </VControl>
            </VField>
          </div>
          <div class="column is-6">
            <VField id="compression_format" v-slot="{ field }">
              <VLabel>压缩格式</VLabel>
              <VControl>
                <Multiselect
                  v-model="form.compression_format"
                  :options="AudioCompressionFormats"
                  placeholder="选择压缩格式"
                  @change="(val: any) => field?.setValue(val)"
                />
                <p v-if="field?.errorMessage" class="help is-danger">
                  {{ field.errorMessage }}
                </p>
              </VControl>
            </VField>
          </div>
          <div class="column is-6">
            <VField id="compression_subtype" v-slot="{ field }">
              <VLabel>子类型</VLabel>
              <VControl>
                <VInput v-model="form.compression_subtype" type="text" placeholder="压缩子类型" />
                <p v-if="field?.errorMessage" class="help is-danger">
                  {{ field.errorMessage }}
                </p>
              </VControl>
            </VField>
          </div>
          <div class="column is-6">
            <VField id="packet_time_us" v-slot="{ field }">
              <VLabel>发包间隔</VLabel>
              <VControl>
                <Multiselect
                  v-model="form.packet_time_us"
                  :options="AudioPacketTimeUs"
                  placeholder="选择发包间隔"
                  @change="(val: any) => field?.setValue(val)"
                />
                <p v-if="field?.errorMessage" class="help is-danger">
                  {{ field.errorMessage }}
                </p>
              </VControl>
            </VField>
          </div>
          <div class="column is-6">
            <VField id="bitrate_bps" v-slot="{ field }" label="音频码率" addons class="is-input-number">
              <VControl expanded>
                <VInputNumber v-model="form.bitrate_bps" centered :min="0" :step="1" />
                <p v-if="field?.errorMessage" class="help is-danger">
                  {{ field.errorMessage }}
                </p>
              </VControl>
              <VControl>
                <VButton static>bps</VButton>
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
<style lang="scss"></style>
