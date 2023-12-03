<script setup lang="ts">
import { toTypedSchema } from "@vee-validate/zod";
import { useForm } from "vee-validate";
import { z } from "zod";
import { useNotyf } from "/@src/composable/useNotyf";
import CodecForm from "./CodecForm.vue";
import UdxForm from "./UdxForm.vue";
import MVForm from "./MVForm.vue";
import SwitchForm from "./SwitchForm.vue";
import { confirm } from "/@src/utils/dialog";
import { useDevices } from '/@src/stores/device'
import { useTemplate } from "/@src/stores/template";
import downloadJsonFile from '/@src/utils/download-json'

const tmplStore = useTemplate();
const deviceStore = useDevices()

const notyf = useNotyf();
const opened = ref(false);

const deviceName = ref("");
const deviceInput = ref<any>(null);
const tmpl = ref<TemplateData | null>(null)
const node = ref<ClustNode | null>(null)
const device = ref<ClustDevice | null>(null)
const fileInput = ref<HTMLInputElement | null>(null)
let callbacks: any = {}

useListener(Signal.OpenResourceConfig, (p: { tmpl: TemplateData, node: ClustNode, device: ClustDevice, callbacks: any }) => {
  opened.value = true;
  tmpl.value = p.tmpl;
  node.value = p.node;
  device.value = p.device;
  if (p.device?.release) {
    deviceName.value = p.device?.release
  } else {
    deviceName.value = ""
  }
  callbacks = p.callbacks || {}
  nextTick(() => {
    deviceInput.value?.field?.setValue(deviceName.value)
  })
});

const zodSchema = z.object({
  deviceName: z.string({
    required_error: "请输入设备名称",
  }).refine(val => /^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$/.test(val), "请输入合法的设备名称")
  .refine(val => val.length <= 53, "设备名称不能大于53个字符")
});
const validationSchema = toTypedSchema(zodSchema);
const { handleSubmit } = useForm({ validationSchema });

const dgi = computed(() => {
  const isCreated = device.value != null;
  return {
    created: isCreated,
    confirmTitle: isCreated ? "更新设备配置" : "部署设备",
    confirmContent: isCreated ? `确定要更新设备 ${device.value?.release} 的配置吗？` : `确定要将应用 ${tmpl.value?.name} 部署到 ${node.value?.id}(${node.value?.ip}) 吗？`,
    confirmMsg: isCreated ? "更新设备配置成功" : "部署设备成功",
    title: isCreated ? "更新设备配置" : "启动设备",
    submitText: isCreated ? "更新" : "启动",
    nodeName: isCreated ? device.value?.node : node.value?.id,
  };
})

async function prepareParams() {
  const tmplRes: TemplateData = await tmplStore.$getTmplById(tmpl.value!.id)
  if (tmplRes.id) {
    const val = compRef.value?.getValue()
    const params: any = {}
    const rq = tmplRes.requirement
    params.cpu = +rq.cpuNum || 1
    params.memory = +rq.memory || 1
    params.hugepage = +rq.hugePage || 6
    params.nodeName = dgi.value.nodeName
    params.cpucore = rq.cpuCore
    params.localip0 = val!['2110-7_m_local_ip']
    params.localip1 = val!['2110-7_b_local_ip']
    params.configFile = JSON.stringify(compRef.value?.getValue())
    params.configFilePath = '/opt/vsomediasoftware/config/vsompconfiginfo-web.json'
    params.hostNetwork = rq.hostNetwork
    params.logLevel = rq.logLevel
    params.repairRecvFrame = +rq.repairRecvFrame
    params.repairSendFrame = +rq.repairSendFrame
    params.dmaList = rq.dmaList || ""
    params.utfOffset = rq.utfOffset
    params.recvAVFrameNodeCount = rq.recvAVFrameNodeCount
    params.sendAVFrameNodeCount = rq.sendAVFrameNodeCount
    params.recvframeCnt = rq.recvframeCnt
    return params
  } else {
    return { error: "模板不存在" }
  }
}

const loading = ref(false)
const addInstance = handleSubmit(async () => {
  if (loading.value) return;
  confirm({
    title: dgi.value.confirmTitle,
    content: dgi.value.confirmContent,
    onConfirm: async (hide) => {
      hide()
      loading.value = true;
      const params = await prepareParams()
      if (params.error) {
        loading.value = false;
        notyf.error(params.error);
        return
      }
      let pro: Promise<any>
      if (dgi.value.created) {
        pro = deviceStore.$updateConfig(deviceName.value, device.value!.id, params)
      } else {
        pro = deviceStore.$deploy(deviceName.value, tmpl.value!.id, params)
      }
      const res = await pro
      loading.value = false;
      if (res.result === 'error') {
        notyf.error(res.message);
      } else {
        opened.value = false;
        notyf.success(dgi.value.confirmMsg);
        callbacks.success?.()
      }
    },
  });
});

function importSetting() {
  fileInput.value?.click()
}

function onFileImported(e: Event) {
  const file = (e.target as HTMLInputElement).files?.[0]
  if (file) {
    const reader = new FileReader()
    reader.onload = (e) => {
      const res = e.target?.result
      if (res) {
        try {
          const val = JSON.parse(res as string)
          compRef.value?.setValue(val)
          notyf.success("导入配置成功");
        } catch (err) {
          console.log(err)
          notyf.error("解析文件格式错误");
        }
        return
      } else {
        notyf.error("导入配置失败");
      }
    }
    reader.readAsText(file)
  } else {
    notyf.error("导入配置失败");
  }
}

function saveSetting() {
  const val = compRef.value?.getValue()
  if (val) {
    downloadJsonFile(val, `${val.moudle}.json`)
    notyf.success("保存成功");
  }
}

const compRef = ref<InstanceType<typeof CodecForm> | null>(null)
const TmplComponent = computed(() => {
  switch (tmpl.value?.typeName) {
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
    :title="dgi.title"
    actions="right"
    cancel-label="取消"
    @submit.prevent="opened = false"
    @close="opened = false"
  >
    <template #content>
      <div class="modal-form">
        <VField
          id="deviceName"
          ref="deviceInput"
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
              :readonly="dgi.created"
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
        <TmplComponent ref="compRef" />
        <input
          ref="fileInput"
          type="file"
          accept=".json"
          style="display: none;"
          @change="onFileImported"
        >
      </div>
    </template>
    <template #action>
      <VButton
        class="btn-setting-import"
        color="primary"
        raised
        @click="importSetting"
      >
        导入配置
      </VButton>
      <VButton
        class="btn-setting-save"
        color="primary"
        raised
        @click="saveSetting"
      >
        保存配置
      </VButton>
      <VButton
        type="submit"
        color="primary"
        raised
        :loading="loading"
        @click="addInstance"
      >
        {{ dgi.submitText }}
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

.v-modal {
  .btn-setting-import,
  .btn-setting-save {
    position: absolute;
  }

  .btn-setting-import {
    left: 20px;
  }

  .btn-setting-save {
    left: 138px;
  }
}
</style>
