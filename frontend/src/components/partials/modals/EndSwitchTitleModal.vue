<script setup lang="ts">
import { toTypedSchema } from "@vee-validate/zod";
import { useForm } from "vee-validate";
import { z } from "zod";
import { useNotyf } from "@src/composable/useNotyf";
import { useUserSession } from "@src/stores/userSession";

const usStore = useUserSession();
const titles = computed(() => usStore.endSwitchTitles);
const notyf = useNotyf();
const opened = ref(false);
const titleIndex = ref(0);
const title = ref("");

useListener(Signal.OpenEndSwtTitle, (idx: number) => {
  opened.value = true
  titleIndex.value = idx
  title.value = titles.value[idx] || (idx === 0 ? '4K 频道1末级切换' : `信号${idx}`)
  setFieldValue("title", title.value)
});

const formTitle = computed(() => titleIndex.value === 0 ? '末级切换面板' : `信号${titleIndex.value}`);

const zodSchema = z.object({
  title: z.string({ required_error: `请输入标题内容` })
    .trim()
    .nonempty(`请输入标题内容`)
});
const validationSchema = toTypedSchema(zodSchema);
const { handleSubmit, setFieldValue } = useForm({ validationSchema });

const loading = ref(false);
const handleCommit = handleSubmit(async () => {
  if (loading.value) return;

  loading.value = true;
  titles.value[titleIndex.value] = title.value;
  const res = await usStore.$updateSettings({ key: 'endswt_titles', value: JSON.stringify(titles.value) });
  loading.value = false;
  if (res) {
    opened.value = false;
    notyf.success("保存设置成功");
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
    title="设置标题"
    actions="right"
    cancel-label="取消"
    @submit.prevent="handleCommit"
    @close="opened = false"
  >
    <template #content>
      <div class="modal-form">
        <VField
          id="title"
          v-slot="{ field }"
          :label="`${formTitle}标题 *`"
        >
          <VControl>
            <VInput
              v-model="title"
              type="text"
              placeholder="标题内容"
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
