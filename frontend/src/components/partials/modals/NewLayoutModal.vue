<script setup lang="ts">
import { useLayout } from "@src/stores/layout";
import { toTypedSchema } from "@vee-validate/zod";
import { useForm } from "vee-validate";
import { z } from "zod";
import { useNotyf } from "@src/composable/useNotyf";
import { LayoutSize } from '@src/utils/enums';

const notyf = useNotyf();
const opened = ref(false);
const draftLayout = ref({ size: LayoutSize.FK } as Layout);
const layoutUtils = useLayout();
let callbacks: any = {}

const zodSchema = z.object({
  name: z.string({ required_error: "请输入布局名称" }).trim().nonempty("请输入布局名称"),
  size: z.union([z.literal(LayoutSize.FK), z.literal(LayoutSize.HD)], {
    errorMap: issue => {
      if (issue.code === z.ZodIssueCode.invalid_union) {
        return { message: "请选择 4K 或 HD" }
      }
      return { message: "请选择屏幕大小" }
    }
  }),
});
const validationSchema = toTypedSchema(zodSchema);
const { handleSubmit, setValues } = useForm({ validationSchema });

useListener(Signal.OpenNewLayout, p => {
  opened.value = true;
  const def = Object.assign({ size: LayoutSize.FK, name: '' }, p?.layout)
  draftLayout.value = def
  callbacks = p?._callbacks || {}
  setValues(def)
});


const loading = ref(false);
const handleCreate = handleSubmit(async (values) => {
  if (loading.value) return;

  loading.value = true;
  const res = await (draftLayout.value.id ? layoutUtils.$updateLayout({
    id: draftLayout.value.id,
    name: values.name,
    size: values.size,
  }) : layoutUtils.$addLayout({
    name: values.name,
    size: values.size,
  }));
  loading.value = false;
  const act = draftLayout.value.id ? "修改" : "新建";
  if (res?.layout) {
    opened.value = false;
    notyf.success(`${act}布局成功`);
    callbacks?.success?.(res.layout);
  } else {
    notyf.error(res?.error ? res.message : `${act}布局失败`);
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
    title="新建多画面布局"
    actions="right"
    cancel-label="取消"
    @submit.prevent="handleCreate"
    @close="opened = false"
  >
    <template #content>
      <div class="modal-form">
        <VField
          id="name"
          v-slot="{ field }"
          label="布局名称 *"
        >
          <VControl>
            <VInput
              type="text"
              placeholder="布局名称"
            />
            <p v-if="field?.errorMessage" class="help is-danger">
              {{ field.errorMessage }}
            </p>
          </VControl>
        </VField>
        <VField
          id="size"
          v-slot="{ field }"
          label="屏幕大小 *"
        >
          <VControl>
            <VSelect>
              <VOption :value="LayoutSize.FK">
                4K
              </VOption>
              <VOption :value="LayoutSize.HD">
                HD
              </VOption>
            </VSelect>
            <p v-if="field?.errorMessage" class="help is-danger">
              {{ field.errorMessage }}
            </p>
          </VControl>
        </VField>
      </div>
    </template>
    <template #action>
      <VButton
        :loading="loading"
        type="submit"
        color="primary"
        raised
      >
        保存
      </VButton>
    </template>
  </VModal>
</template>
