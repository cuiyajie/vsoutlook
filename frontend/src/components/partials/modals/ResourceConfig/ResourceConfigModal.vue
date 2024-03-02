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
import { useClustNode } from "/@src/stores/node";
import isCIDR from "/@src/utils/cidr";

const tmplStore = useTemplate();
const deviceStore = useDevices()
const nodeStore = useClustNode();

const notyf = useNotyf();
const opened = ref(false);

const deviceName = ref("");
const deviceInput = ref<any>(null);
const tmpl = ref<TemplateData | null>(null)
const node = ref<ClustNode | null>(null)
const device = ref<ClustDevice | null>(null)
const fileInput = ref<HTMLInputElement | null>(null)
const inited = ref(false)
let callbacks: any = {}

const tmpls = computed(() => tmplStore.tmpls)
const nodes = computed(() => nodeStore.nodes)

function onTmplSelect(field?: any, val?: any) {
  field?.setValue(val)
  tmpl.value = tmpls.value.find(t => t.id === val) || null
  if (val && !tmpl.value?.requirement) {
    tmplStore.$getTmplById(val).then(res => {
      tmpl.value = res
    })
  }
}

function onNodeSelect(field?: any, val?: any) {
  field?.setValue(val);
  node.value = nodes.value.find(n => n.id === val) || null
}

useListener(Signal.OpenResourceConfig, (p?: { tmpl: TemplateData, node: ClustNode, device: ClustDevice, callbacks: any }) => {
  opened.value = true;
  inited.value = !!p
  if (!p) return
  if (!p.tmpl?.requirement) {
    tmplStore.$getTmplById(p.tmpl.id).then(res => {
      tmpl.value = res
    })
  }
  tmpl.value = p.tmpl;
  node.value = p.node;
  device.value = p.device;
  if (p.device?.name) {
    deviceName.value = p.device?.name || ""
  } else {
    deviceName.value = ""
  }
  callbacks = p.callbacks || {}
  nextTick(() => {
    deviceInput.value?.field?.setValue(deviceName.value)
    if (p.device?.config) {
      const val = JSON.parse(p.device.config)
      compRef.value?.setValue(val)
    }
  })
});

const validationSchema = computed(() => {
  const rules: any = {
    deviceName: z.string({
      required_error: "请输入设备名称",
    }).refine(val => /^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$/.test(val), "请输入合法的设备名称")
    .refine(val => val.length <= 53, "设备名称不能大于53个字符"),
  };
  if (!inited.value) {
    rules.tmpl = z.string({
      required_error: "请选择应用",
    }).nullable().refine(v => v !== null, "请选择应用");
    rules.node = z.string({
      required_error: "请选择容器",
    }).nullable().refine(v => v !== null, "请选择容器");
  }
  return toTypedSchema(z.object(rules))
})
const { handleSubmit } = useForm({ validationSchema });

const dgi = computed(() => {
  const isCreated = device.value != null;
  return {
    created: isCreated,
    confirmTitle: isCreated ? "更新设备配置" : "部署设备",
    confirmContent: isCreated ? `确定要更新设备 ${device.value?.name} 的配置吗？` : `确定要将应用 ${tmpl.value?.name} 部署到 ${node.value?.id}(${node.value?.ip}) 吗？`,
    confirmMsg: isCreated ? "更新设备配置成功" : "部署设备成功",
    title: `${isCreated ? "更新设备配置" : "启动设备"} - ${tmplConfig.value.name}`,
    submitText: isCreated ? "更新" : "启动",
    nodeName: isCreated ? device.value?.node : node.value?.id,
  };
})

async function prepareParams() {
  if (tmpl.value?.id && node.value?.id) {
    // const val = compRef.value?.getValue()
    const params: any = {}
    const val = compRef.value?.getValue()
    if (!val || !isCIDR(`${val['2110-7_m_local_ip']}/24`)) {
      return { error: "主网卡地址格式错误" }
    }
    if (!val || !isCIDR(`${val['2110-7_b_local_ip']}/24`)) {
      return { error: "备网卡地址格式错误" }
    }
    params.primaryVFAddress = `${val['2110-7_m_local_ip']}/24`
    params.secondaryVFAddress = `${val['2110-7_b_local_ip']}/24`
    params.configFile = val

    const rq = tmpl.value.requirement
    params.targetNode = dgi.value.nodeName
    params.cpu = +rq.cpuNum || 1
    params.memory = +rq.memory || 1
    params.hugepages = +rq.hugePage || 6
    params.shm = rq.shm || 0

    // params.configFilePath = '/opt/vsomediasoftware/config/vsompconfiginfo-web.json'
    // params.hostNetwork = rq.hostNetwork
    params.logLevel = rq.logLevel
    params.MaxRateMbpsByCore = rq.maxRateMbpsByCore
    params.RXsessioncnt = rq.receiveSessions
    // params.repairRecvFrame = +rq.repairRecvFrame
    // params.repairSendFrame = +rq.repairSendFrame
    // params.utfOffset = rq.utfOffset
    // params.recvAVFrameNodeCount = rq.recvAVFrameNodeCount
    // params.sendAVFrameNodeCount = rq.sendAVFrameNodeCount
    // params.recvframeCnt = rq.recvframeCnt
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
        pro = deviceStore.$deploy(deviceName.value, tmpl.value!.id, node.value!.id, params)
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
const tmplConfig = computed(() => {
  switch (tmpl.value?.typeName) {
    case "编解码":
      return {
        name: '编解码器',
        component: CodecForm
      };
    case "上下变换":
      return {
        name: '上下变换',
        component: UdxForm
      };
    case "多画面":
      return {
        name: '多画面',
        component: MVForm
      };
    case "切换台":
      return {
        name: '切换台',
        component: SwitchForm
      };
    default:
      return {
        name: '编解码器',
        component: CodecForm
      };
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
    dialog-class="resource-config-modal"
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
          class="device-name"
        >
          <VControl fullwidth>
            <VInput
              v-model="deviceName"
              type="text"
              placeholder="分段以字母数字开头，可包含连字符(-)，字母数字结尾，可由点分隔的多段组成，例如:my-example.name123"
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
        <div v-if="!inited" class="columns">
          <div class="column is-6 mt-4">
            <VField
              id="tmpl"
              v-slot="{ field }"
              label="选择应用*"
            >
              <VControl fullwidth>
                <Multiselect
                  placeholder="选择应用"
                  value-prop="id"
                  label="name"
                  :max-height="145"
                  :options="tmpls"
                  @change="(val: any) => onTmplSelect(field, val)"
                >
                  <template #singlelabel="{ value }">
                    <div class="multiselect-single-label">
                      <span class="select-label-text">
                        {{ value.name }}
                      </span>
                    </div>
                  </template>
                  <template #option="{ option }">
                    <span class="select-label-text">
                      {{ option.name }}
                    </span>
                  </template>
                </Multiselect>
                <p
                  v-if="field?.errorMessage"
                  class="help is-danger"
                >
                  {{ field.errorMessage }}
                </p>
              </VControl>
            </VField>
          </div>
          <div class="column is-6 mt-4">
            <VField
              id="node"
              v-slot="{ field }"
              label="选择容器*"
            >
              <VControl fullwidth>
                <Multiselect
                  placeholder="选择容器"
                  value-prop="id"
                  label="name"
                  :max-height="145"
                  :options="nodes"
                  @change="(val: any) => onNodeSelect(field, val)"
                >
                  <template #singlelabel="{ value }">
                    <div class="multiselect-single-label">
                      <span class="select-label-text">
                        {{ value.id }}
                      </span>
                    </div>
                  </template>
                  <template #option="{ option }">
                    <span class="select-label-text">
                      {{ option.id }}
                    </span>
                  </template>
                </Multiselect>
                <p
                  v-if="field?.errorMessage"
                  class="help is-danger"
                >
                  {{ field.errorMessage }}
                </p>
              </VControl>
            </VField>
          </div>
        </div>
        <tmplConfig.component v-if="tmpl" ref="compRef" :name="deviceName" :requiredment="tmpl?.requirement" />
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

.v-modal.resource-config-modal {
  .modal-card-body {
    min-height: 400px;
  }

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
