<script setup lang="ts">
import { z } from "zod";
import { useNotyf } from "@src/composable/useNotyf";
import { useClustNode } from "@src/stores/node";
import { toTypedSchema } from "@vee-validate/zod";
import { useForm } from "vee-validate";

type NicInfoForm = Omit<NicInfo, "id" | "nodeId" | "position">;

const defNicInfo = (): NicInfoForm => ({
  nicNameMain: "",
  nicNameBackup: "",
  receiveMain: true,
  receiveBackup: true,
  sendMain: true,
  sendBackup: true,
  coreList: "",
  txRxCoreList: "",
  dmaList: "",
  vfCount: 4,
});

const notyf = useNotyf();
const opened = ref(false);
const nodeStore = useClustNode();
const indexRef = ref(-1);
const node = ref<ClustNode | null>(null);
const form = ref<NicInfoForm>(defNicInfo());
let callbacks: any = {};

const loading = ref(false);
const zodSchema = z.object({
  nicNameMain: z.string({ required_error: "请输入主路网口" }).trim().nonempty("请输入主路网口"),
  receiveMain: z.boolean().optional(),
  sendMain: z.boolean().optional(),
  nicNameBackup: z.string({ required_error: "请输入备路网口" }).trim().nonempty("请输入备路网口"),
  receiveBackup: z.boolean().optional(),
  sendBackup: z.boolean().optional(),
  coreList: z
    .string({ required_error: "请输入DPDK核心列表" })
    .nonempty("请输入DPDK核心列表")
    .refine(
      (val) => /^(\d+|(\d+-\d+))(,(\d+|(\d+-\d+)))*$/.test(val),
      "请输入合法的DPDK核心列表"
    ),
  txRxCoreList: z
    .string({ required_error: "请输入应用收发流核心列表" })
    .nonempty("请输入应用收发流核心列表")
    .refine(
      (val) => /^(\d+|(\d+-\d+))(,(\d+|(\d+-\d+)))*$/.test(val),
      "请输入合法的应用收发流核心列表"
    ),
  dmaList: z
    .string()
    .trim(),
  vfCount: z.number().optional(),
});
const validationSchema = toTypedSchema(zodSchema);
const { handleSubmit, setFieldValue } = useForm({ validationSchema });

const handleEdit = handleSubmit(async () => {
  if (loading.value || !node.value?.id) return;

  loading.value = true;

  let pro: Promise<any>;
  if (indexRef.value < 0) {
    pro = nodeStore.$createNic(
      {
        ...form.value,
      },
      node.value.id
    );
  } else {
    const nic = node.value.nics[indexRef.value];
    pro = nodeStore.$updateNic({
      ...form.value,
      id: nic.id,
      nodeId: node.value.id,
      position: nic.position,
    });
  }
  const res = await pro;
  loading.value = false;
  if (res?.code === 0) {
    opened.value = false;
    notyf.success(
      indexRef.value < 0 ? "创建节点网卡成功" : `保存网卡 ${indexRef.value} 信息成功`
    );
  } else if (res.error) {
    notyf.error(
      indexRef.value < 0 ? res.message : `网卡 ${indexRef.value} ${res.message}`
    );
  }
  callbacks.success?.();
});

function syncForm() {
  for (let key of Object.keys(form.value) as (keyof NicInfoForm)[]) {
    setFieldValue(key, form.value[key]);
  }
}

syncForm();

useListener(
  Signal.OpenNodeEdit,
  (payload: { callbacks?: any; index?: number; node?: ClustNode }) => {
    opened.value = true;
    indexRef.value = payload.index === undefined ? -1 : payload.index;
    node.value = payload.node || null;
    if (indexRef.value < 0) {
      form.value = defNicInfo();
    } else {
      const nic = node.value?.nics[indexRef.value] || defNicInfo();
      form.value = {
        nicNameMain: nic.nicNameMain,
        nicNameBackup: nic.nicNameBackup,
        receiveMain: nic.receiveMain,
        receiveBackup: nic.receiveBackup,
        sendMain: nic.sendMain,
        sendBackup: nic.sendBackup,
        coreList: nic.coreList,
        txRxCoreList: nic.txRxCoreList,
        dmaList: nic.dmaList,
        vfCount: nic.vfCount,
      };
    }

    callbacks = payload.callbacks || {};
    syncForm();
  }
);
</script>
<template>
  <VModal
    is="form"
    method="post"
    novalidate
    size="large"
    :open="opened"
    title="编辑 Node"
    actions="right"
    cancel-label="取消"
    @submit.prevent="handleEdit"
    @close="opened = false"
  >
    <template #content>
      <div class="modal-form node-edit-form">
        <div class="columns is-multiline">
          <div class="column is-6">
            <VField id="nicNameMain" v-slot="{ field }" label="主路网口 *" horizontal>
              <VControl class="input-1">
                <VInput
                  v-model="form.nicNameMain"
                  type="text"
                  placeholder="请输入网口名称"
                />
                <expand-transition>
                  <p v-if="field?.errorMessage" class="help is-danger mt-3">
                    {{ field.errorMessage }}
                  </p>
                </expand-transition>
              </VControl>
            </VField>
          </div>
          <div class="column is-6">
            <VField label="收发开关" horizontal>
              <VControl subcontrol class="switch-2">
                <VSwitchBlock v-model="form.receiveMain" color="info" label="收" />
              </VControl>
              <VControl subcontrol class="switch-2">
                <VSwitchBlock v-model="form.sendMain" color="danger" label="发" />
              </VControl>
            </VField>
          </div>
          <div class="column is-6">
            <VField id="nicNameBackup" v-slot="{ field }" label="备路网口 *" horizontal>
              <VControl class="input-1">
                <VInput
                  v-model="form.nicNameBackup"
                  type="text"
                  placeholder="请输入网口名称"
                />
                <expand-transition>
                  <p v-if="field?.errorMessage" class="help is-danger mt-3">
                    {{ field.errorMessage }}
                  </p>
                </expand-transition>
              </VControl>
            </VField>
          </div>
          <div class="column is-6">
            <VField label="收发开关" horizontal>
              <VControl subcontrol class="switch-2">
                <VSwitchBlock v-model="form.receiveBackup" color="info" label="收" />
              </VControl>
              <VControl subcontrol class="switch-2">
                <VSwitchBlock v-model="form.sendBackup" color="danger" label="发" />
              </VControl>
            </VField>
          </div>
          <div class="column is-6">
            <VField id="coreList" v-slot="{ field }" label="DPDK核心列表">
              <VControl>
                <VInput v-model="form.coreList" type="text" placeholder="例如: 2-31,32" />
                <expand-transition>
                  <p v-if="field?.errorMessage" class="help is-danger mt-3">
                    {{ field.errorMessage }}
                  </p>
                </expand-transition>
              </VControl>
            </VField>
          </div>
          <div class="column is-6">
            <VField id="txRxCoreList" v-slot="{ field }" label="应用收发流核心列表">
              <VControl>
                <VInput v-model="form.txRxCoreList" type="text" placeholder="例如: 2-31,32" />
                <expand-transition>
                  <p v-if="field?.errorMessage" class="help is-danger mt-3">
                    {{ field.errorMessage }}
                  </p>
                </expand-transition>
              </VControl>
            </VField>
          </div>
          <div class="column is-6">
            <VField id="dmaList" v-slot="{ field }" label="DMA通道列表">
              <VControl>
                <VInput
                  v-model="form.dmaList"
                  type="text"
                  placeholder="例如: 0000:00:01.0,0000:00:01.1"
                />
                <expand-transition>
                  <p v-if="field?.errorMessage" class="help is-danger mt-3">
                    {{ field.errorMessage }}
                  </p>
                </expand-transition>
              </VControl>
            </VField>
          </div>
          <div class="column is-6">
            <VField label="VF数量">
              <VControl>
                <VInputNumber v-model="form.vfCount" centered :min="1" :step="1" />
              </VControl>
            </VField>
          </div>
        </div>
      </div>
    </template>
    <template #action>
      <VButton type="submit" color="primary" raised> 保存 </VButton>
    </template>
  </VModal>
</template>
<style lang="scss">
.node-edit-form.modal-form {
  .field.is-horizontal {
    .field-label {
      flex: 0 0 auto;
      padding-top: 0;
      display: flex;
      align-items: center;
    }

    .field-body {
      flex: 1;

      > .control {
        width: 100%;

        &.input-1 {
          padding-top: 5.25px;
        }

        &.switch-2 {
          width: 40%;
        }
      }
    }
  }
}
</style>
