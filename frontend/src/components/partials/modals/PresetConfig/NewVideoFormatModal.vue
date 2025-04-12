<script setup lang="ts">
import { toTypedSchema } from '@vee-validate/zod'
import { useForm } from 'vee-validate'
import { z } from 'zod'
import { useNotyf } from '@src/composable/useNotyf'
import {
  VideoFormatProtocols,
  VideoGammas,
  VideoGamuts,
  VideoCompressionFormats,
  VideoCompressionSubtypes,
  VideoCompressionRatios,
  VideoFormatFpsList,
  defVideoFormat,
  defaultFpsStr,
} from './Consts'
import { useUserSession } from '@src/stores/userSession'
import { handleVideoForm, parseFps, stringifyFps } from './Utils'

type VideoFormatForm = VideoFormat & { fpsStr: string };

const usStore = useUserSession()
const videoFormats = computed(() => usStore.settings.video_formats || [])
const notyf = useNotyf()
const opened = ref(false)
const indexRef = ref(-1)
let callbacks: any = {}
const form = ref<VideoFormatForm>(Object.assign(defVideoFormat(), { fpsStr: defaultFpsStr }))

const zodSchema = z.object({
  name: z
    .string({ required_error: '请输入模板名称' })
    .nonempty('请输入模板名称')
    .regex(/^[\u4e00-\u9fa5a-zA-Z0-9-]+$/, { message: '请输入由中英文和横杠组成的名称' })
    .trim()
    .refine(
      (value) => {
        if (indexRef.value >= 0) return true
        return !videoFormats.value.some((item) => item.name === value)
      },
      { message: '模板名称已存在' }
    ),
  protocol: z.string({ required_error: '请选择格式类型' }),
  width: z.number({ required_error: '请输入视频宽度', invalid_type_error: '请输入视频宽度' }).int(),
  height: z.number({ required_error: '请输入视频高度', invalid_type_error: '请输入视频高度' }).int(),
  interlaced: z.boolean(),
  fps: z.number(),
  fpsStr: z.string().refine((value) => VideoFormatFpsList.includes(value), { message: '请选择视频帧率' }),
  gamma: z.string({ required_error: '请选择视频伽马' }),
  gamut: z.string({ required_error: '请选择视频色域' }),
  compression_format: z.string().optional(),
  compression_subtype: z.string().optional(),
  compression_ratio: z.string().optional(),
  bitrate_bps: z
    .number({ required_error: '请输入视频码率' })
    .int()
    .optional(),
  gop_b_frames: z
    .number({ invalid_type_error: '请输入B帧数量' })
    .int()
    .optional(),
  gop_length: z
    .number({ invalid_type_error: '请输入GOP长度' })
    .int()
    .optional(),
})
const validationSchema = toTypedSchema(zodSchema)
const { handleSubmit, setFieldValue } = useForm({ validationSchema })

const loading = ref(false)
const handleCommit = handleSubmit(async () => {
  if (loading.value) return

  loading.value = true
  let value: string
  const formValue = handleVideoForm(form.value) as VideoFormat
  if (indexRef.value < 0) {
    value = JSON.stringify([formValue, ...videoFormats.value])
  } else {
    const formats = videoFormats.value
    formats[indexRef.value] = formValue
    value = JSON.stringify(formats)
  }
  const res = await usStore.$updateSettings({ key: 'video_formats', value })
  loading.value = false
  if (res) {
    opened.value = false
    notyf.success(`${indexRef.value < 0 ? '创建' : '保存'}视频格式成功`)
    callbacks.success?.()
  } else {
    notyf.error(`${indexRef.value < 0 ? '创建' : '保存'}视频格式失败`)
  }
})

function syncForm() {
  for (let key of Object.keys(form.value) as (keyof VideoFormat)[]) {
    setFieldValue(key, form.value[key])
  }
}

function onFpsStrChange(val: string) {
  const { fps, interlaced } = parseFps(val)
  setFieldValue('fpsStr', val)
  setFieldValue('fps', fps)
  setFieldValue('interlaced', interlaced)
  form.value.fps = fps
  form.value.interlaced = interlaced
}

syncForm()

const subTypes = computed(() => {
  if (form.value.compression_format === 'speedhq') {
    return VideoCompressionSubtypes.slice(-2)
  } else {
    return VideoCompressionSubtypes.slice(0, -2)
  }
})

watch(() => form.value.compression_format, () => {
  const isSpeedHqSubTypes = VideoCompressionSubtypes.slice(-2).includes(form.value.compression_subtype)
  if (form.value.compression_format === 'speedhq' && !isSpeedHqSubTypes) {
    setFieldValue('compression_subtype', '')
    form.value.compression_subtype = ''
  } else if (form.value.compression_format !== 'speedhq' && isSpeedHqSubTypes) {
    setFieldValue('compression_subtype', '')
    form.value.compression_subtype = ''
  }
})

const showAdditionalOptions = computed(() => {
  return !['st2110-20','st2110-22'].includes(form.value.protocol) && form.value.compression_format !== 'speedhq'
})

useListener(Signal.OpenNewVideoFormat, (payload: { _callback?: any; index?: number }) => {
  opened.value = true
  indexRef.value = payload.index === undefined ? -1 : payload.index
  if (indexRef.value < 0) {
    form.value = Object.assign(defVideoFormat(), { fpsStr: defaultFpsStr })
  } else {
    const v = videoFormats.value[indexRef.value]
    form.value = {
      ...defVideoFormat(),
      ...v,
      fpsStr: stringifyFps(v.fps, v.interlaced),
      compression_subtype: v.compression_subtype === '' ? '无' : v.compression_subtype,
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
    size="big"
    :open="opened"
    title="视频格式模版"
    actions="right"
    cancel-label="取消"
    dialog-class="preset-modal preset-modal is-overflow-visible"
    @submit.prevent="handleCommit"
    @close="opened = false"
  >
    <template #content>
      <div class="modal-form">
        <div class="columns is-multiline">
          <div class="column is-4">
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
          <div class="column is-4">
            <VField id="format" v-slot="{ field }">
              <VLabel>格式类型</VLabel>
              <VControl>
                <Multiselect
                  v-model="form.protocol"
                  :options="VideoFormatProtocols"
                  value-prop="value"
                  label="label"
                  placeholder="选择格式类型"
                  @change="(val: any) => field?.setValue(val)"
                />
                <p v-if="field?.errorMessage" class="help is-danger">
                  {{ field.errorMessage }}
                </p>
              </VControl>
            </VField>
          </div>
          <div class="column is-4">
            <VField label="分辨率" addons class="resolution-addons is-input-number">
              <VControl id="width" v-slot="{ field }" expanded>
                <VInputNumber
                  v-model="form.width"
                  centered
                  :min="0"
                  :step="1"
                  :show-control="false"
                />
                <p v-if="field?.errorMessage" class="help is-danger">
                  {{ field.errorMessage }}
                </p>
              </VControl>
              <VControl>
                <VButton class="addon-conn" static>x</VButton>
              </VControl>
              <VControl id="height" v-slot="{ field }" expanded>
                <VInputNumber
                  v-model="form.height"
                  centered
                  :min="0"
                  :step="1"
                  :show-control="false"
                />
                <p v-if="field?.errorMessage" class="help is-danger">
                  {{ field.errorMessage }}
                </p>
              </VControl>
            </VField>
          </div>
          <div class="column is-4">
            <VField id="fpsStr" v-slot="{ field }">
              <VLabel>视频帧率</VLabel>
              <VControl>
                <Multiselect
                  v-model="form.fpsStr"
                  :options="VideoFormatFpsList"
                  placeholder="选择视频帧率"
                  :style="{'--ms-max-height': '180px'}"
                  @change="onFpsStrChange"
                />
                <p v-if="field?.errorMessage" class="help is-danger">
                  {{ field.errorMessage }}
                </p>
              </VControl>
            </VField>
          </div>
          <div class="column is-4">
            <VField id="gamma" v-slot="{ field }">
              <VLabel>视频伽马</VLabel>
              <VControl>
                <Multiselect
                  v-model="form.gamma"
                  :options="VideoGammas"
                  placeholder="选择视频伽马"
                  @change="(val: any) => field?.setValue(val)"
                />
                <p v-if="field?.errorMessage" class="help is-danger">
                  {{ field.errorMessage }}
                </p>
              </VControl>
            </VField>
          </div>
          <div class="column is-4">
            <VField id="gamut" v-slot="{ field }">
              <VLabel>视频色域</VLabel>
              <VControl>
                <Multiselect
                  v-model="form.gamut"
                  :options="VideoGamuts"
                  placeholder="选择视频色域"
                  @change="(val: any) => field?.setValue(val)"
                />
                <p v-if="field?.errorMessage" class="help is-danger">
                  {{ field.errorMessage }}
                </p>
              </VControl>
            </VField>
          </div>
          <expand-transition>
            <div v-if="form.protocol !== 'st2110-20'" class="column is-4">
              <VField id="compression_format" v-slot="{ field }">
                <VLabel>压缩格式</VLabel>
                <VControl>
                  <Multiselect
                    v-model="form.compression_format"
                    :options="VideoCompressionFormats"
                    placeholder="选择压缩格式"
                    @change="(val: any) => field?.setValue(val)"
                  />
                  <p v-if="field?.errorMessage" class="help is-danger">
                    {{ field.errorMessage }}
                  </p>
                </VControl>
              </VField>
            </div>
          </expand-transition>
          <expand-transition>
            <div v-if="!['st2110-20','st2110-22'].includes(form.protocol)" class="column is-4">
              <VField id="compression_format" v-slot="{ field }">
                <VLabel>压缩子类型</VLabel>
                <VControl>
                  <Multiselect
                    v-model="form.compression_subtype"
                    :options="subTypes"
                    placeholder="选择压缩子类型"
                    @change="(val: any) => field?.setValue(val)"
                  />
                  <p v-if="field?.errorMessage" class="help is-danger">
                    {{ field.errorMessage }}
                  </p>
                </VControl>
              </VField>
            </div>
          </expand-transition>
          <expand-transition>
            <div v-if="form.protocol === 'st2110-22'" class="column is-4">
              <VField id="compression_format" v-slot="{ field }">
                <VLabel>压缩比</VLabel>
                <VControl>
                  <Multiselect
                    v-model="form.compression_ratio"
                    :options="VideoCompressionRatios"
                    placeholder="选择压缩比"
                    @change="(val: any) => field?.setValue(val)"
                  />
                  <p v-if="field?.errorMessage" class="help is-danger">
                    {{ field.errorMessage }}
                  </p>
                </VControl>
              </VField>
            </div>
          </expand-transition>
          <expand-transition>
            <div v-if="showAdditionalOptions" class="column is-4">
              <VField id="bitrate_bps" v-slot="{ field }" label="视频码率" addons class="is-input-number">
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
          </expand-transition>
          <expand-transition>
            <div v-if="showAdditionalOptions" class="column is-4">
              <VField id="gop_b_frames" v-slot="{ field }">
                <VLabel>B帧数量</VLabel>
                <VControl>
                  <VInputNumber v-model="form.gop_b_frames" centered :min="0" :step="1" />
                  <p v-if="field?.errorMessage" class="help is-danger">
                    {{ field.errorMessage }}
                  </p>
                </VControl>
              </VField>
            </div>
          </expand-transition>
          <expand-transition>
            <div v-if="showAdditionalOptions" class="column is-4">
              <VField id="gop_length" v-slot="{ field }">
                <VLabel>GOP长度</VLabel>
                <VControl>
                  <VInputNumber v-model="form.gop_length" centered :min="0" :step="1" />
                  <p v-if="field?.errorMessage" class="help is-danger">
                    {{ field.errorMessage }}
                  </p>
                </VControl>
              </VField>
            </div>
          </expand-transition>
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
