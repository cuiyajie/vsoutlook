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
const nameRef = ref<any>(null);
const sizeRef = ref<any>(null);
let callbacks: any = {}

useListener(Signal.OpenNewLayout, (p: any) => {
  opened.value = true;
  const def = Object.assign({ size: LayoutSize.FK, name: '' }, p?.layout)
  draftLayout.value = def
  callbacks = p?._callbacks || {}
  nextTick(() => {
    nameRef.value?.field?.setValue(def.name)
    sizeRef.value?.field?.setValue(def.size)
  })
});

const zodSchema = z.object({
  name: z.string({
    required_error: "请输入布局名称",
  }),
  size: z.number({
    required_error: "请输入屏幕大小",
  }),
});
const validationSchema = toTypedSchema(zodSchema);
const { handleSubmit } = useForm({ validationSchema });

const loading = ref(false);
const handleCreate = handleSubmit(async () => {
  if (loading.value) return;

  loading.value = true;
  const nlayout = await (draftLayout.value.id ? layoutUtils.$updateLayout({
    id: draftLayout.value.id,
    name: draftLayout.value.name,
    size: draftLayout.value.size,
  }) : layoutUtils.$addLayout({
    name: draftLayout.value.name,
    size: draftLayout.value.size,
  }));
  loading.value = false;
  const act = draftLayout.value.id ? "修改" : "新建";
  if (nlayout) {
    opened.value = false;
    notyf.success(`${act}布局成功`);
    callbacks?.success?.(nlayout);
  } else {
    notyf.error(`${act}布局失败`);
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
          ref="nameRef"
          v-slot="{ field }"
          label="布局名称 *"
        >
          <VControl>
            <VInput
              v-model="draftLayout.name"
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
        <VField
          id="size"
          ref="sizeRef"
          v-slot="{ field }"
          label="屏幕大小 *"
        >
          <VControl>
            <VSelect v-model="draftLayout.size">
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
        type="submit"
        color="primary"
        raised
      >
        保存
      </VButton>
    </template>
  </VModal>
</template>
