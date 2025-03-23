<script setup lang="ts">
import { toTypedSchema } from '@vee-validate/zod'
import { useForm } from 'vee-validate'
import { z } from 'zod'
import { useNotyf } from '@src/composable/useNotyf'
import { useUserSession } from '@src/stores/userSession'
import { isValidNumberRanges } from './Utils'

type AudioMappingForm = Pick<AudioMapping, 'name' | 'mute_channels'>

const usStore = useUserSession()
const audioMappings = computed(() => usStore.settings.audio_mappings || [])
const notyf = useNotyf()
const opened = ref(false)
const indexRef = ref(-1)
let callbacks: any = {}
const form = ref<AudioMappingForm>({
  name: '',
  mute_channels: '',
})
const audioCopys = ref<AudioChannelMapping[]>([])

function addAudioCopy() {
  audioCopys.value.unshift({
    src_channel: 1,
    dst_channel: 1,
  })
}

function removeAudioCopy(idx: number) {
  audioCopys.value.splice(idx, 1)
}

const zodSchema = z.object({
  name: z
    .string({ required_error: '请输入模板名称' })
    .nonempty('请输入模板名称')
    .regex(/^[\u4e00-\u9fa5a-zA-Z0-9-]+$/, { message: '请输入由中英文和横杠组成的名称' })
    .trim()
    .refine(
      (value) => {
        if (indexRef.value >= 0) return true
        return !audioMappings.value.some((item) => item.name === value)
      },
      { message: '模板名称已存在' }
    ),
  mute_channels: z
    .string({ required_error: '请输入静音声道范围' })
    .nonempty('请输入静音声道范围')
    .trim()
    .refine(
      (value) => {
        return isValidNumberRanges(value)
      },
      { message: '请输入例如 1-3,4,5 的范围' }
    ),
})
const validationSchema = toTypedSchema(zodSchema)
const { handleSubmit, setFieldValue } = useForm({ validationSchema })

const loading = ref(false)
const handleCommit = handleSubmit(async () => {
  if (loading.value) return

  loading.value = true
  let value: string
  const formValue = {
    ...form.value,
    channels: audioCopys.value,
  } as AudioMapping
  if (indexRef.value < 0) {
    value = JSON.stringify([formValue, ...audioMappings.value])
  } else {
    const formats = audioMappings.value
    formats[indexRef.value] = formValue
    value = JSON.stringify(formats)
  }
  const res = await usStore.$updateSettings({ key: 'audio_mappings', value })
  loading.value = false
  if (res) {
    opened.value = false
    notyf.success(`${indexRef.value < 0 ? '创建' : '保存'}音频映射模版成功`)
    callbacks.success?.()
  } else {
    notyf.error(`${indexRef.value < 0 ? '创建' : '保存'}音频映射模版失败`)
  }
})

function syncForm() {
  for (let key of Object.keys(form.value) as (keyof AudioMappingForm)[]) {
    setFieldValue(key, form.value[key])
  }
}

syncForm()

useListener(
  Signal.OpenNewAudioMapping,
  (payload: { _callback?: any; index?: number }) => {
    opened.value = true
    indexRef.value = payload.index === undefined ? -1 : payload.index
    const mapping = audioMappings.value[indexRef.value]
    if (mapping) {
      form.value = {
        name: mapping.name,
        mute_channels: mapping.mute_channels,
      }
      audioCopys.value = mapping.channels
    } else {
      form.value = {
        name: '',
        mute_channels: '',
      }
      audioCopys.value = []
    }

    callbacks = payload._callback || {}
    syncForm()
  }
)
</script>
<template>
  <VModal
    is="form"
    method="post"
    novalidate
    size="large"
    :open="opened"
    title="音频映射模版"
    actions="right"
    cancel-label="取消"
    dialog-class="preset-modal audio-mapping-modal"
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
                <VInput v-model="form.name" type="text" placeholder="请输入模版名称" />
                <p v-if="field?.errorMessage" class="help is-danger">
                  {{ field.errorMessage }}
                </p>
              </VControl>
            </VField>
          </div>
          <div class="column is-6">
            <VField id="mute_channels" v-slot="{ field }">
              <VLabel>声道静音</VLabel>
              <VControl>
                <VInput v-model="form.mute_channels" type="text" placeholder="请输入声道范围" />
                <p v-if="field?.errorMessage" class="help is-danger">
                  {{ field.errorMessage }}
                </p>
              </VControl>
            </VField>
          </div>
          <div class="column is-12">
            <div class="copy-header">
              <div class="left">声道复制</div>
              <button
                type="button"
                class="button is-circle is-dark-outlined"
                @keydown.space.prevent="addAudioCopy"
                @click.prevent="addAudioCopy"
              >
                <span class="icon is-large">
                  <i aria-hidden="true" class="iconify" data-icon="feather:plus" />
                </span>
              </button>
            </div>
            <div v-for="(ac, idx) in audioCopys" :key="idx" class="columns is-multiline">
              <div class="column is-5">
                <VField class="copy-input" horizontal>
                  <VLabel>源声道</VLabel>
                  <VControl>
                    <VInputNumber v-model="ac.src_channel" centered :step="1" />
                  </VControl>
                </VField>
              </div>
              <div class="column is-5">
                <VField class="copy-input" horizontal>
                  <VLabel>目标声道</VLabel>
                  <VControl>
                    <VInputNumber v-model="ac.dst_channel" centered :step="1" />
                  </VControl>
                </VField>
              </div>
              <div class="column is-2">
                <VField class="copy-action">
                  <VControl>
                    <VIconButton
                      color="danger"
                      light
                      raised
                      circle
                      icon="feather:x"
                      @click="removeAudioCopy(idx)"
                    />
                  </VControl>
                </VField>
              </div>
            </div>
            <div v-if="audioCopys.length === 0" class="form-empty">
              暂时没有添加音频映射模版
            </div>
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
.modal.preset-modal.audio-mapping-modal {

  .modal-content,
  .modal-card,
  .modal-card-body {
    overflow: auto;
  }

  .copy-header {
    display: flex;
    align-items: center;
    justify-content: space-between;
    margin-top: 16px;
    margin-bottom: 24px;
  }

  .field.copy-action {
    display: flex;
    align-items: center;
    justify-content: flex-end;
  }

  .field.copy-input {
    .label {
      margin-bottom: 0;
      display: flex;
      align-items: center;
      margin-right: 16px;
    }

    .input {
      width: 160px;
    }
  }

  .form-empty {
    height: 60px;
    display: flex;
    align-items: center;
    justify-content: center;
    font-size: 0.875rem;
    font-weight: 300;
  }
}
</style>
