<script setup lang="ts">
import { useVueFlow, type FlowExportObject } from "@vue-flow/core";
import { useTemplate } from "@src/stores/template";
import { useTemplateType } from "@src/stores/templateType";
import { useNotyf } from "@src/composable/useNotyf";

const route = useRoute();
const params = route.params as { id: string }
const tmplTypeStore = useTemplateType();
const tmplStore = useTemplate();
const { toObject } = useVueFlow();
const locked = ref(true);
const tmpl = ref<any>(null)
const disabled = computed(() => !tmpl.value?.id || locked.value)
const notyf = useNotyf();

function createNewTmpl() {
  bus.trigger(Signal.OpenNewTemplate, {
    success: (tmpl: TemplateData) => {
      tmplStore.navigate(tmpl.id)
    }
  });
}

function updateTmplMeta() {
  if (tmpl.value?.id) {
    bus.trigger(Signal.OpenTmplMetaEdit, {
      tmpl: tmpl.value,
      callbacks: {
        success: (_tmpl: TemplateData) => {
          tmpl.value = _tmpl
        }
      }
    });
  }
}

const flowObject = ref<FlowExportObject | null>(null)
function tabChange(tab: string) {
  if (tab === 'resource') {
    flowObject.value = toObject()
  }
}

const specsData = ref<TmplRequirement & { description: string }>({
  cpu: '',
  cpuNum: 0,
  cpuCore: '',
  hugePage: 0,
  memory: 0,
  disk: '',
  gpu: '',
  inputBand: '',
  outputBand: '',
  network: '',
  chart: '',
  description: '',
  logLevel: 0,
  repairRecvFrame: true,
  repairSendFrame: true,
  hostNetwork: false,
  utfOffset: 37,
  recvAVFrameNodeCount: 2,
  sendAVFrameNodeCount: 2,
  recvFrameCount: 2,
  maxRateMbpsByCore: 9000,
  receiveSessions: 18,
  shm: 0,
  nicCount: 1,
  nicConfig: [],
})

tmplTypeStore.$fetchList();
tmplStore.$fetchList();

watch(() => params?.id, async (nv, ov) => {
  if (nv && nv !== ov) {
    tmpl.value = await tmplStore.$getTmplById(nv)
    if (tmpl.value) {
      const requirement = tmpl.value.requirement as TmplRequirement
      if (requirement.nicCount && requirement.nicConfig) {
        if (requirement.nicConfig.length < requirement.nicCount) {
          requirement.nicConfig = Array.from({ length: requirement.nicCount }, () => ({ dpdkCpu: 1, dma: 1 }))
        } else {
          requirement.nicCount = requirement.nicConfig.length
        }
      }
      specsData.value = Object.assign({}, requirement, { description: tmpl.value.description })
      if (tmpl.value.flow) {
        flowObject.value = JSON.parse(tmpl.value.flow)
      }
    }
  }
}, { immediate: true });

async function save() {
  if (specsData.value.cpuNum === 0) {
    notyf.error('CPU 总核心数不能为 0')
    return
  }
  const dpdkCpu = specsData.value.nicConfig.reduce((acc, cur) => acc + cur.dpdkCpu, 0)
  if (dpdkCpu === 0) {
    notyf.error('DPDK CPU 核心不能为 0')
    return
  }
  if (specsData.value.cpuNum <= dpdkCpu) {
    notyf.error('CPU 总核心数不能小于等于 DPDK CPU 核心')
    return
  }
  if (tmpl.value?.id) {
    const { description, ...requirement } = specsData.value
    flowObject.value = toObject()
    const res = await tmplStore.$updateTmpl(tmpl.value.id, {
      flow: JSON.stringify(flowObject.value),
      requirement,
      description
    })
    if (res) {
      notyf.success('保存应用成功')
    }
  }
}

onMounted(() => {
  if (!params?.id) {
    locked.value = false
    createNewTmpl()
  }
})
</script>

<template>
  <div class="columns is-multiline page-template">
    <div class="column is-2">
      <VCard radius="rounded" class="card-widget">
        <div class="title-wrap mb-4">
          <h3>应用列表</h3>
          <button class="button is-circle is-dark-outlined" @click="createNewTmpl">
            <span class="icon is-small">
              <i aria-hidden="true" class="iconify" data-icon="feather:plus" />
            </span>
          </button>
        </div>
        <TmplGroupList />
      </VCard>
    </div>
    <div class="column is-8">
      <VCard radius="rounded" class="tmpl-dashboard">
        <div class="title is-5 mb-6 tmpl-title">
          <h3>
            {{ tmpl?.name || "未命名应用" }}
            <span
              class="icon ml-2"
              style="cursor: pointer"
              role="button"
              tabindex="-1"
              @keyup.space.prevent="updateTmplMeta"
              @click.prevent="updateTmplMeta"
            >
              <i aria-hidden="true" class="fas fa-edit" />
            </span>
          </h3>

          <VButtons>
            <VButton color="primary" raised :disabled="disabled" @click="save">
              <span class="icon">
                <i aria-hidden="true" class="fas fa-save" />
              </span>
              <span>保存</span>
            </VButton>
            <VButton raised :disabled="disabled">
              <span class="icon">
                <i aria-hidden="true" class="fas fa-share-square" />
              </span>
              <span>另存为</span>
            </VButton>
            <VButton raised :disabled="disabled">
              <span class="icon">
                <i class="fas fa-cloud-upload-alt" aria-hidden="true" />
              </span>
              <span>发布到应用商店</span>
            </VButton>
            <VButton raised :disabled="disabled">
              <span class="icon">
                <i class="fas fa-caret-square-right" aria-hidden="true" />
              </span>
              <span>生成程序</span>
            </VButton>
          </VButtons>
        </div>
        <VField class="lock-switch">
          <VControl>
            <VSwitchBlock
              v-model="locked"
              :label="locked ? '已锁定' : '可编辑'"
              thin
              color="primary"
            />
          </VControl>
        </VField>
        <VTabs
          class="dashboard-tabs"
          :class="[locked && 'is-locked']"
          slider
          type="rounded"
          selected="resource"
          align="centered"
          :tabs="[
            { label: '能力组件拼装', value: 'module' },
            { label: '资源需求', value: 'resource' },
          ]"
          @update:selected="tabChange"
        >
          <template #tab="{ activeValue }">
            <KeepAlive>
              <div v-if="activeValue === 'module'" class="flow-container">
                <AbilityFlow :flow="flowObject" />
              </div>
              <p v-else-if="activeValue === 'resource'">
                <TmplSpecsForm v-model="specsData" />
              </p>
            </KeepAlive>
          </template>
        </VTabs>
      </VCard>
    </div>
    <div class="column is-2">
      <VCard radius="rounded">
        <h3 class="title is-5 mb-4">能力组件</h3>
        <AbilityComponent :locked="locked" />
      </VCard>
    </div>
    <NewTemplateModal />
    <TmplMetaEdit />
  </div>
</template>
<style lang="scss">
.tmpl-dashboard {
  position: relative;

  .lock-switch {
    position: absolute;
    right: 24px;
  }

  .tmpl-title {
    display: flex;
    justify-content: space-between;
    align-items: center;
  }

  .flow-container {
    height: 100%;
    background-color: var(--dark-sidebar-light-2);
    border-radius: 16px;
    overflow: hidden;
  }

  .form-layout .form-outer {
    border-radius: 16px;
  }
}

.dashboard-tabs {
  height: calc(100% - 3rem - 38px);
  &.is-locked {
    .flow-container,
    .form-layout {
      position: relative;
      &:before {
        position: absolute;
        content: "";
        cursor: not-allowed;
        top: 0;
        left: 0;
        right: 0;
        bottom: 0;
        z-index: 10;
        border-radius: 16px;
        background-color: rgba(255, 255, 255, 0.1);
      }
    }
  }
  &.tabs-wrapper.is-slider {
    > .tabs-inner > .tabs {
      max-width: 280px;
      height: 40px;
    }

    .tab-naver {
      height: 38px;
    }
  }

  .tab-content {
    height: calc(100% - 60px);
  }
}

.page-template {
  min-height: calc(100vh - 150px);
  .column {
    padding: 0.25rem;

    .l-card {
      height: 100%;
    }
  }
}

.is-dark {
  .tmpl-dashboard .form-layout .form-outer {
    background-color: var(--dark-sidebar-light-4);
  }
}
</style>
