<script setup lang="ts">
import { useTemplate } from "@src/stores/template";
import { toTypedSchema } from "@vee-validate/zod";
import { useForm } from "vee-validate";
import { z } from "zod";
import { useNotyf } from "@src/composable/useNotyf";

const notyf = useNotyf();
const opened = ref(false);
const tmpl = ref<TemplateData | null>(null);
const tmplUtils = useTemplate();
let callbacks: any = {}

useListener(Signal.OpenTmplMetaEdit, (p: { callbacks?: any, tmpl: TemplateData }) => {
  opened.value = true;
  tmpl.value = { ...p.tmpl };
  setFieldValue("name", p.tmpl.name)
  callbacks = p.callbacks || {}
});

const zodSchema = z.object({
  name: z.string({ required_error: "请输入应用名称" }).trim().nonempty("请输入应用名称")
});
const validationSchema = toTypedSchema(zodSchema);
const { handleSubmit, setFieldValue } = useForm({ validationSchema });

const loading = ref(false);
const handleSave = handleSubmit(async () => {
  if (loading.value || !tmpl.value?.id) return;

  loading.value = true;
  const ntmpl = await tmplUtils.$updateTmpl(tmpl.value?.id, {
    name: tmpl.value.name
  });
  loading.value = false;
  if (ntmpl) {
    opened.value = false;
    notyf.success("修改应用成功");
    callbacks?.success?.(ntmpl);
  } else {
    notyf.error("修改应用失败");
  }
});
</script>
<template>
  <VModal
    is="form"
    method="post"
    novalidate
    size="small"
    :open="opened"
    title="修改应用"
    actions="right"
    cancel-label="取消"
    @submit.prevent="handleSave"
    @close="opened = false"
  >
    <template #content>
      <div v-if="tmpl" class="modal-form">
        <VField
          id="name"
          v-slot="{ field }"
          label="设备名称 *"
        >
          <VControl>
            <VInput
              v-model="tmpl.name"
              type="text"
              placeholder="设备名称"
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
