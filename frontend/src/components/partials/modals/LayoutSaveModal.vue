<script setup lang="ts">
import { toTypedSchema } from "@vee-validate/zod";
import { useForm } from "vee-validate";
import { z } from "zod";
import { useNotyf } from "@src/composable/useNotyf";
import { useLayout } from "@src/stores/layout";

const notyf = useNotyf();
const opened = ref(false);
const layout = ref({} as Layout);
const layoutUtils = useLayout();
const locationRef = ref<any>(null);
let callbacks: any = {}

useListener(Signal.LayoutSaveAs, (p: any) => {
  opened.value = true;
  layout.value = p?.layout || {};
  callbacks = p?._callbacks || {}
  nextTick(() => {
  })
});

const zodSchema = z.object({
  location: z.string({
    required_error: "请输入存储位置",
  }),
});
const validationSchema = toTypedSchema(zodSchema);
const { handleSubmit } = useForm({ validationSchema });

const loading = ref(false);
const handleSave = handleSubmit(async () => {
  if (loading.value) return;

  loading.value = true;
  const nlayout = await layoutUtils.$updateLocation(layout.value.id, layout.value.location);
  loading.value = false;
  if (nlayout) {
    opened.value = false;
    notyf.success("布局存储成功");
    callbacks?.success?.(nlayout);
  } else {
    notyf.error("布局存储失败");
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
    title="另存到"
    actions="right"
    cancel-label="取消"
    @submit.prevent="handleSave"
    @close="opened = false"
  >
    <template #content>
      <div class="modal-form">
        <VField
          label="布局名称"
        >
          <VControl>
            <VInput
              v-model="layout.name"
              readonly
              type="text"
            />
          </VControl>
        </VField>
        <VField
          id="location"
          ref="locationRef"
          v-slot="{ field }"
          label="存储位置"
        >
          <VControl>
            <VInput
              v-model="layout.location"
              type="text"
              placeholder="存储位置"
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
