<script setup lang="ts">
import { useTemplate } from '@src/stores/template'
import { toTypedSchema } from '@vee-validate/zod'
import { useForm } from 'vee-validate'
import { z } from 'zod'
import { useNotyf } from '@src/composable/useNotyf'

const notyf = useNotyf()
const opened = ref(false)
const draftTmpl = ref({} as TemplateData)
const tmplUtils = useTemplate()
let callbacks: any = {}

useListener(Signal.OpenNewTemplate, (_callbacks?: any) => {
  opened.value = true
  callbacks = _callbacks || {}
})

const zodSchema = z.object({
  tmplType: z
    .string({ required_error: '请选择应用类型' })
    .trim()
    .nonempty('请选择应用类型'),
  name: z.string({ required_error: '请输入应用名称' }).trim().nonempty('请输入应用名称'),
})
const validationSchema = toTypedSchema(zodSchema)
const { handleSubmit } = useForm({ validationSchema })

const loading = ref(false)
const handleCreate = handleSubmit(async () => {
  if (loading.value) return

  loading.value = true
  const ntmpl = await tmplUtils.$addTmpl({
    name: draftTmpl.value.name,
    type: draftTmpl.value.type,
  })
  loading.value = false
  if (ntmpl) {
    opened.value = false
    notyf.success('新建应用成功')
    callbacks?.success?.(ntmpl)
  } else {
    notyf.error('新建应用失败')
  }
})
</script>
<template>
  <VModal
    is="form"
    method="post"
    novalidate
    size="small"
    :open="opened"
    title="新建应用"
    actions="right"
    cancel-label="取消"
    dialog-class="is-overflow-visible"
    @submit.prevent="handleCreate"
    @close="opened = false"
  >
    <template #content>
      <div class="modal-form">
        <TmplTypeSelect
          v-model="draftTmpl.type"
          label="应用类型 *"
          form-id="tmplType"
          validate
        />
        <VField id="name" v-slot="{ field }" label="设备名称 *">
          <VControl>
            <VInput v-model="draftTmpl.name" type="text" placeholder="设备名称" />
            <p v-if="field?.errorMessage" class="help is-danger">
              {{ field.errorMessage }}
            </p>
          </VControl>
        </VField>
      </div>
    </template>
    <template #action>
      <VButton type="submit" color="primary" raised> 保存 </VButton>
    </template>
  </VModal>
</template>
