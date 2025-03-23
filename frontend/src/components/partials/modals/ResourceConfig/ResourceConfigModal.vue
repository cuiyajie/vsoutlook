<script setup lang="ts">
import { toTypedSchema } from '@vee-validate/zod'
import { useForm } from 'vee-validate'
import { z } from 'zod'
import { useNotyf } from '@src/composable/useNotyf'
import CodecForm from './Codec/CodecForm.vue'
import UdxForm from './Udx/UdxForm.vue'
import MVForm_V1 from './MVForm_V1/MVForm_V1.vue'
import MVForm from './MVForm/MVForm.vue'
import SwitchForm from './Switch/SwitchForm.vue'
import BCSwitchForm from './BCSwitch/BCSwitchForm.vue'
import EndSwitchForm from './EndSwitch/EndSwitchForm.vue'
import RecorderForm from './Recorder/RecorderForm.vue'
import MediaGateForm from "./MediaGate/MediaGateForm.vue"
import { confirm } from '@src/utils/dialog'
import { useDevices } from '@src/stores/device'
import { useTemplate } from '@src/stores/template'
import downloadJsonFile from '@src/utils/download-json'
import { useClustNode } from '@src/stores/node'
import { useUserSession } from '@src/stores/userSession'
import dayjs from 'dayjs'

const tmplStore = useTemplate()
const deviceStore = useDevices()
const nodeStore = useClustNode()
const usStore = useUserSession()

const notyf = useNotyf()
const opened = ref(false)

const deviceName = ref('')
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
  tmpl.value = tmpls.value.find((t) => t.id === val) || null
}

function onNodeSelect(field?: any, val?: any) {
  field?.setValue(val)
  node.value = nodes.value.find((n) => n.id === val) || null
}

useListener(
  Signal.OpenResourceConfig,
  (p?: { tmpl: TemplateData; node: ClustNode; device: DeviceDetail; callbacks: any }) => {
    opened.value = true
    callbacks = p?.callbacks || {}
    inited.value = (!p?.tmpl || !p?.node) && !p?.device
    if (p?.device) {
      device.value = p.device
      deviceName.value = p.device?.name || ''
      tmpl.value = tmpls.value.find((t) => t.id === p.device.tmplID) || null
      node.value = nodes.value.find((n) => n.id === p.device.node) || null
    } else {
      tmpl.value = p?.tmpl || null
      node.value = p?.node || null
      deviceName.value = ''
      device.value = null
    }
    if (tmpl.value?.id) {
      tmplStore.$getTmplById(tmpl.value.id).then(res => {
        if (res) {
          tmpl.value = res
        }
      })
    }
    if (node.value) {
      nodeStore.$getNodeNics(node.value.id)
    }
    nodeStore.$fetchNMos()
    nextTick(() => {
      setValues({
        deviceName: deviceName.value,
        tmpl: tmpl.value?.id || null,
        node: node.value?.id || null,
      })
      if (p?.device?.config) {
        const val = JSON.parse(p.device.config)
        compRef.value?.setValue(val)
      }
    })
  }
)

const validationSchema = computed(() => {
  const rules: any = {
    deviceName: z
      .string({
        required_error: '请输入设备名称',
      })
      .refine(
        (val) =>
          /^[A-Za-z0-9]([-A-Za-z0-9]*[A-Za-z0-9])?(\.[A-Za-z0-9]([-A-Za-z0-9]*[A-Za-z0-9])?)*$/.test(
            val
          ),
        '请输入合法的设备名称'
      )
      .refine((val) => val.length <= 53, '设备名称不能大于53个字符'),
  }
  if (!inited.value) {
    rules.tmpl = z
      .string({
        required_error: '请选择应用',
      })
      .nullable()
      .refine((v) => v !== null, '请选择应用')
    rules.node = z
      .string({
        required_error: '请选择容器',
      })
      .nullable()
      .refine((v) => v !== null, '请选择容器')
  }
  return toTypedSchema(z.object(rules))
})
const { handleSubmit, setValues } = useForm({ validationSchema })

const dgi = computed(() => {
  const isCreated = device.value != null
  return {
    created: isCreated,
    confirmTitle: isCreated ? '更新设备配置' : '部署设备',
    confirmContent: isCreated
      ? `确定要更新设备 ${device.value?.name} 的配置吗？`
      : `确定要将 ${deviceName.value} 设备部署到 ${node.value?.id} (${node.value?.ip}) 吗？`,
    confirmMsg: isCreated ? '更新设备配置成功' : '部署设备成功',
    title: `${isCreated ? '更新设备配置' : '启动设备'}${
      tmpl.value ? ` - ${tmpl.value.name}` : ''
    }`,
    submitText: isCreated ? '更新' : '启动',
    nodeName: isCreated ? device.value?.node : node.value?.id,
  }
})

function checkRequestParams(params: any) {
  if (tmpl.value?.typeCategory === 'multiv') {
    const output = params?.output?.out_params || []
    if (output.some((opt: any) => !opt.out_mv_template)) {
      notyf.error('多画面 输出参数 - 输出布局 - 布局模板 不能为空')
      return false
    }
  }
  const nicCount = tmpl.value?.requirement.nicCount || 0
  if (nicCount > 0) {
    const nics = params?.nic_list || []
    if (nics.length !== nicCount) {
      notyf.error(`${tmpl.value?.name} 需要配置 ${nicCount} 个网卡`)
      return false
    }
  }
  return true
}

const loading = ref(false)
const addInstance = handleSubmit(async () => {
  if (loading.value) return
  confirm({
    title: dgi.value.confirmTitle,
    content: dgi.value.confirmContent,
    size: 'medium',
    onConfirm: async (hide) => {
      hide()
      const params = compRef.value?.getValue()
      if (!params || !checkRequestParams(params)) return
      loading.value = true
      let pro: Promise<any>
      if (dgi.value.created) {
        pro = deviceStore.$updateConfig(deviceName.value, device.value!.id, params)
      } else {
        pro = deviceStore.$deploy(
          deviceName.value,
          tmpl.value!.id,
          node.value!.id,
          params
        )
      }
      const res = await pro
      loading.value = false
      if (res.result === 'error') {
        notyf.error(res.message)
      } else {
        opened.value = false
        notyf.success(dgi.value.confirmMsg)
        callbacks.success?.()
        if (usStore.settings.auto_save_container_config) {
          let config = null
          try {
            config = JSON.parse(res.data.config)
          } catch (err) {}
          saveSetting(config)
        }
      }
    },
  })
})

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
          notyf.success('导入配置成功')
        } catch (err) {
          console.log(err)
          notyf.error('解析文件格式错误')
        }
        return
      } else {
        notyf.error('导入配置失败')
      }
    }
    reader.readAsText(file)
  } else {
    notyf.error('导入配置失败')
  }
}

function exportSetting() {
  bus.trigger(Signal.OpenDownloadConfig, {
    tmpl: tmpl.value,
    callbacks: {
      success: (fileName: string) => {
        saveSetting('', fileName)
      },
    },
  })
}

function saveSetting(config?: string, name?: string) {
  const val = config || compRef.value?.getValue()
  const fileName =
    name ||
    `${tmpl.value?.name}_${deviceName.value || '未命名'}_${dayjs().format(
      'YYYYMMDDHHmmss'
    )}`
  if (val) {
    downloadJsonFile(val, `${fileName}.json`)
    notyf.success('保存成功')
  }
}

const compRef = ref<InstanceType<typeof CodecForm> | null>(null)
const tmplConfig = computed(() => {
  switch (tmpl.value?.typeCategory) {
    case 'codec':
      return {
        name: '编解码器',
        component: CodecForm,
      }
    case 'udx':
      return {
        name: '上下变换',
        component: UdxForm,
      }
    case 'multiv':
      return {
        name: '多画面',
        component: MVForm_V1,
      }
    case 'swt':
      return {
        name: '切换台',
        component: SwitchForm,
      }
    case 'bcswt':
      return {
        name: '播出切换台',
        component: BCSwitchForm,
      }
    case 'endswt':
      return {
        name: '末级切换',
        component: EndSwitchForm,
      }
    case 'recorder':
      return {
        name: '录放机',
        component: RecorderForm,
      }
    case 'media_gateway':
      return {
        name: '全媒体网关',
        component: MediaGateForm
      }
    case 'mv':
      return {
        name: '多画面',
        component: MVForm
      }
    default:
      return {
        name: '编解码器',
        component: CodecForm,
      }
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
        <VField id="deviceName" v-slot="{ field }" label="设备名称 *" class="device-name">
          <VControl fullwidth>
            <VInput
              v-model="deviceName"
              type="text"
              placeholder="分段以字母数字开头，可包含连字符(-)，字母数字结尾，可由点分隔的多段组成，例如:my-example.name123"
              :readonly="dgi.created"
            />
            <Transition name="fade-slow">
              <p v-if="field?.errorMessage" class="help is-danger mt-3">
                {{ field.errorMessage }}
              </p>
            </Transition>
          </VControl>
        </VField>
        <div v-if="inited" class="columns">
          <div class="column is-6 mt-4">
            <VField id="tmpl" v-slot="{ field }" label="应用类型 *">
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
                <p v-if="field?.errorMessage" class="help is-danger">
                  {{ field.errorMessage }}
                </p>
              </VControl>
            </VField>
          </div>
          <div class="column is-6 mt-4">
            <VField id="node" v-slot="{ field }" label="部署节点 *">
              <VControl fullwidth>
                <Multiselect
                  placeholder="选择节点"
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
                <p v-if="field?.errorMessage" class="help is-danger">
                  {{ field.errorMessage }}
                </p>
              </VControl>
            </VField>
          </div>
        </div>
        <tmplConfig.component
          v-if="tmpl"
          ref="compRef"
          :name="deviceName"
          :requiredment="tmpl?.requirement"
          :nics="node?.nics"
        />
        <input
          ref="fileInput"
          type="file"
          accept=".json"
          style="
            width: 0;
            height: 0;
            opacity: 0;
            position: absolute;
            top: -100px;
            left: -100px;
            z-index: -1;
          "
          @change="onFileImported"
        />
      </div>
    </template>
    <template #action>
      <VButton class="btn-setting-import" color="primary" raised @click="importSetting">
        导入配置
      </VButton>
      <VButton class="btn-setting-save" color="primary" raised @click="exportSetting">
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

  .input[readonly] {
    background-color: var(--dark-sidebar-light-10);
    border-color: var(--dark-sidebar-light-10);
  }
}

.v-modal.resource-config-modal {

  &.modal.is-big .modal-content {
    max-width: 960px;
  }

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
