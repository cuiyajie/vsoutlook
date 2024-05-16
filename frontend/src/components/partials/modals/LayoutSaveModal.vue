<script setup lang="ts">
import { toTypedSchema } from "@vee-validate/zod";
import { useForm } from "vee-validate";
import { z } from "zod";
import { useNotyf } from "@src/composable/useNotyf";
import { useLayout } from "@src/stores/layout";

const notyf = useNotyf();
const opened = ref(false);
const layout = ref({} as Layout);
const layoutName = ref('')
const layoutUtils = useLayout();
let callbacks: any = {}

const initialValues = ref({} as Layout)
useListener(Signal.LayoutSaveAs, (p: any) => {
  opened.value = true;
  layout.value = p?.layout || {};
  initialValues.value = { ...layout.value }
  callbacks = p?._callbacks || {}
  nextTick(() => {
  })
});

const zodSchema = z.object({
  name: z.string({
    required_error: "布局名称不能为空",
  }),
});
const validationSchema = toTypedSchema(zodSchema);
const { handleSubmit } = useForm({ validationSchema, initialValues });

const loading = ref(false);
const handleSave = handleSubmit(async () => {
  if (loading.value) return;

  loading.value = true;
  const res = await layoutUtils.$duplicate(layout.value.id, layoutName.value);
  loading.value = false;
  if (res?.layout) {
    opened.value = false;
    notyf.success("布局复制成功");
    callbacks?.success?.(res.layout);
  } else {
    notyf.error(res?.error ? res.message : "布局复制失败");
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
    title="复制布局"
    actions="right"
    cancel-label="取消"
    @submit.prevent="handleSave"
    @close="opened = false"
  >
    <template #content>
      <div class="modal-form">
        <VField
          id="name"
          v-slot="{ field }"
          label="布局名称"
        >
          <VControl>
            <VInput
              v-model="layoutName"
              type="text"
              placeholder="布局名称"
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
