<script setup lang="ts">
import { z } from "zod";
import { useNotyf } from "/@src/composable/useNotyf";
import { useClustNode } from "/@src/stores/node";
import { toTypedSchema } from "@vee-validate/zod";
import { useForm } from "vee-validate";

const notyf = useNotyf();
const opened = ref(false);
const nodeInfo = ref<{ coreList: string }>({ coreList: "" });
const node = ref<ClustNode | null>(null);
const nodeStore = useClustNode();
const coreListRef = ref<any>(null);

useListener(Signal.OpenNodeEdit, (_node: ClustNode) => {
  opened.value = true;
  node.value = _node;
  nodeInfo.value.coreList = _node.coreList
  nextTick(() => {
    coreListRef.value?.field?.setValue(nodeInfo.value.coreList)
  })
});

const loading = ref(false);
const zodSchema = z.object({
  coreList: z.string({
    required_error: "请输入隔离核心数",
  }).refine(val => /^(\d+|(\d+-\d+))(,(\d+|(\d+-\d+)))*$/.test(val), "请输入合法的隔离核心数")
});
const validationSchema = toTypedSchema(zodSchema);
const { handleSubmit } = useForm({ validationSchema });

const handleEdit = handleSubmit(async () => {
  if (loading.value) return;

  loading.value = true;
  const res = await nodeStore.$update(node.value!.id, nodeInfo.value.coreList);
  loading.value = false;
  if (res?.code === 0) {
    opened.value = false;
    notyf.success("保存 Node 信息成功");
  } else if (res.error) {
    notyf.error(res.message);
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
    title="编辑 Node"
    actions="right"
    cancel-label="取消"
    @submit.prevent="handleEdit"
    @close="opened = false"
  >
    <template #content>
      <div class="modal-form">
        <VField
          id="coreList"
          ref="coreListRef"
          v-slot="{ field }"
          label="隔离核心数 *"
        >
          <VControl>
            <VInput
              v-model="nodeInfo.coreList"
              type="text"
              placeholder="例如: 2-31,32"
            />
            <Transition name="fade-slow">
              <p
                v-if="field?.errorMessage"
                class="help is-danger mt-3"
              >
                {{ field.errorMessage }}
              </p>
            </Transition>
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