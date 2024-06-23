<script setup lang="ts">
import { toTypedSchema } from "@vee-validate/zod";
import { useForm } from "vee-validate";
import { z } from "zod";
import { useNotyf } from "@src/composable/useNotyf";
import { useUserSession } from "@src/stores/userSession";

const usStore = useUserSession();
const settings = computed(() => usStore.settings);
const notyf = useNotyf();
const opened = ref(false);
const apiServerPort = ref("");
let callbacks: any = {}

useListener(Signal.OpenApiServerPort, (_callback: any) => {
  opened.value = true
  callbacks = _callback || {}
  setFieldValue("apiServerPort", settings.value.endswt_api || '')
});

const zodSchema = z.object({
  apiServerPort: z.string({ required_error: "请输入对外API IP及端口" })
    .trim()
    .nonempty("请输入对外API IP及端口")
    .regex(/^(?:(?:25[0-5]|2[0-4]\d|1\d{2}|[1-9]?\d)\.){3}(?:25[0-5]|2[0-4]\d|1\d{2}|[1-9]?\d):([0-9]+)$/, { message: "请输入正确的IP地址" })
});
const validationSchema = toTypedSchema(zodSchema);
const { handleSubmit, setFieldValue } = useForm({ validationSchema });

const loading = ref(false);
const handleCommit = handleSubmit(async () => {
  if (loading.value) return;

  loading.value = true;
  const res = await usStore.$updateSettings({ key: 'endswt_api', value: apiServerPort.value });
  loading.value = false;
  if (res) {
    opened.value = false;
    notyf.success("保存设置成功");
    callbacks.success?.();
  } else {
    notyf.error("保存设置失败");
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
    title="末级切换API对外地址"
    actions="right"
    cancel-label="取消"
    @submit.prevent="handleCommit"
    @close="opened = false"
  >
    <template #content>
      <div class="modal-form">
        <VField
          id="apiServerPort"
          v-slot="{ field }"
          label="对外API IP及端口 *"
        >
          <VControl>
            <VInput
              v-model="apiServerPort"
              type="text"
              placeholder="例如: 232.0.61.1:30000"
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
