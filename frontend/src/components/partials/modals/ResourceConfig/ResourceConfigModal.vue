<script setup lang="ts">
import { toTypedSchema } from "@vee-validate/zod";
import { useForm } from "vee-validate";
import { z } from "zod";
import { useNotyf } from "/@src/composable/useNotyf";
import CodecForm from "./CodecForm.vue";
import UdxForm from "./UdxForm.vue";
import MVForm from "./MVForm.vue";
import SwitchForm from "./SwitchForm.vue";

const notyf = useNotyf();
const opened = ref(true);

useListener(Signal.OpenResourceConfig, (p: { tmpl: TemplateData, nodeName: string }) => {
  opened.value = true;
  tmpl.value = p.tmpl;
  nodeName.value = p.nodeName;
});

const deviceName = ref("");
const tmpl = ref<TemplateData | null>(null)
const nodeName = ref("")

const zodSchema = z.object({
  deviceName: z.string({
    required_error: "请输入设备名称",
  })
});
const validationSchema = toTypedSchema(zodSchema);
const { handleSubmit } = useForm({ validationSchema });

const loading = ref(false)
const addInstance = handleSubmit(async () => {
  console.log(compRef.value?.getValue())
  if (true || loading.value) return;

  loading.value = true;
  // TODO: add instance
  loading.value = false;
  if (true) {
    opened.value = false;
    notyf.success("部署应用成功");
  }
});

const compRef = ref<InstanceType<typeof CodecForm> | null>(null)
const TmplComponent = computed(() => {
  const b = "切换台"
  switch (b) {
  // switch (tmpl.value?.typeName) {
    case "编解码":
      return CodecForm;
    case "上下变换":
      return UdxForm;
    case "多画面":
      return MVForm;
    case "切换台":
      return SwitchForm;
    default:
      return CodecForm;
  }
})
</script>
<template>
  <VModal
    is="form"
    method="post"
    novalidate
    noscroll
    noclose
    size="big"
    :open="opened"
    title="启动设备"
    actions="right"
    cancel-label="取消"
    @submit.prevent="opened = false"
    @close="opened = false"
  >
    <template #content>
      <div class="modal-form">
        <VField
          id="deviceName"
          v-slot="{ field }"
          label="设备名称*"
          horizontal
          class="device-name"
        >
          <VControl fullwidth>
            <VInput
              v-model="deviceName"
              type="text"
              placeholder="设备名称"
            />
            <Transition name="fade-slow">
              <p
                v-if="field?.errorMessage"
                class="help is-danger"
              >
                {{ field.errorMessage }}
              </p>
            </Transition>
          </VControl>
        </VField>
        <TmplComponent ref="compRef" />
      </div>
    </template>
    <template #action>
      <VButton
        type="submit"
        color="primary"
        raised
        @click="addInstance"
      >
        启动
      </VButton>
    </template>
  </VModal>
</template>
<style lang="scss">
.device-name {
  &.field.is-horizontal .field-label .label {
    font-size: 1.2rem;
    font-weight: 600;
    text-align: left;
    margin-left: 20px;
  }
}
</style>
