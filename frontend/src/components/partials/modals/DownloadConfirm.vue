<script setup lang="ts">
import { toTypedSchema } from "@vee-validate/zod";
import dayjs from "dayjs";
import { useForm } from "vee-validate";
import { z } from "zod";

const opened = ref(false);
const fileName = ref('');
const nameRef = ref<any>(null);
let callbacks: any = {}

useListener(Signal.OpenDownloadConfig, (p: { callbacks?: any, tmpl: TemplateData }) => {
  opened.value = true;
  fileName.value = `${p.tmpl.typeCategory}_${dayjs().format('YYYYMMDDHHmmss')}`;
  nextTick(() => {
    nameRef.value?.field?.setValue(fileName.value)
  })
  callbacks = p.callbacks || {}
});

const zodSchema = z.object({
  name: z.string({
    required_error: "请输入文件名称",
  }),
});
const validationSchema = toTypedSchema(zodSchema);
const { handleSubmit } = useForm({ validationSchema });

const handleSave = handleSubmit(async () => {
  if (!fileName.value) return;
  opened.value = false;
  callbacks?.success?.(fileName.value);
});
</script>
<template>
  <VModal
    is="form"
    method="post"
    novalidate
    size="small"
    :open="opened"
    title="下载配置文件"
    actions="right"
    cancel-label="取消"
    @submit.prevent="handleSave"
    @close="opened = false"
  >
    <template #content>
      <div class="modal-form">
        <VField
          id="name"
          ref="nameRef"
          v-slot="{ field }"
          label="文件名称 *"
        >
          <VControl>
            <VInput
              v-model="fileName"
              type="text"
              placeholder="请输入配置文件名称"
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
    </template>
    <template #action>
      <VButton
        type="submit"
        color="primary"
        raised
      >
        保存
      </VButton>
    </template>
  </VModal>
</template>
