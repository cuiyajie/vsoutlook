<script lang="ts" setup>
import { useTemplate } from "@src/stores/template";

const tmplStore = useTemplate();
const router = useRouter()
const props = defineProps<{
  tmpl: TemplateData;
}>();

const tmpld = ref<TemplateData>({ requirement: {} } as TemplateData);
watch(() => props.tmpl, async (nv, ov) => {
  if (nv && nv.id !== ov?.id) {
    tmpld.value = await tmplStore.$getTmplById(nv.id)
  }
}, { immediate: true })

function viewTmpl() {
  if (!tmpld.value.id) return
  router.push(`/app/template/${tmpld.value.id}?from=resource`)
}
</script>
<template>
  <div class="kv-info">
    <div class="info-block-inner">
      <div class="title-wrap">
        <h3 class="dark-inverted">
          {{ tmpld.name }}
        </h3>
        <a
          role="button"
          class="action-link"
          tabindex="0"
          @keydown.space.prevent="viewTmpl"
          @click.prevent="viewTmpl"
        >编辑</a>
      </div>
      <div class="info-block-line">
        <h4 class="dark-inverted">
          类型
        </h4>
        <span>{{ tmpld.typeName }}</span>
      </div>
    </div>
    <div class="info-block-inner is-dark-bordered-12">
      <div class="title-wrap">
        <h3 class="dark-inverted">
          容器所需资源
        </h3>
      </div>
      <div class="info-block-line">
        <h4 class="dark-inverted">
          CPU 总核心数
        </h4>
        <span>
          {{ tmpld.requirement.cpuNum || '无' }}
        </span>
      </div>
      <div class="info-block-line">
        <h4 class="dark-inverted">
          需要使用DPDK的网卡数量
        </h4>
        <span>
          {{ tmpld.requirement.nicCount || '无' }}
        </span>
      </div>
      <template v-if="tmpld.requirement.nicConfig?.length">
        <div v-for="(nic, nidx) in tmpld.requirement.nicConfig" :key="nidx" class="info-block-line">
          <h4 class="dark-inverted">
            第 {{ nidx + 1 }} 块网卡
          </h4>
          <span>
            {{ `${nic.dpdkCpu} x Dpdk Cpu, ${nic.dma} x DMA` }}
          </span>
        </div>
      </template>
      <div class="info-block-line">
        <h4 class="dark-inverted">
          内存 (GB)
        </h4>
        <span>
          {{ tmpld.requirement.memory || '无' }}
        </span>
      </div>
      <div class="info-block-line">
        <h4 class="dark-inverted">
          大页内存 (GB)
        </h4>
        <span>
          {{ tmpld.requirement.hugePage || '无' }}
        </span>
      </div>
      <div class="info-block-line">
        <h4 class="dark-inverted">
          共享内存 (GB)
        </h4>
        <span>
          {{ tmpld.requirement.shm || '无' }}
        </span>
      </div>
      <div class="info-block-line">
        <h4 class="dark-inverted">
          日志级别
        </h4>
        <span>
          {{ tmpld.requirement.logLevel || '无' }}
        </span>
      </div>
      <div class="info-block-line">
        <h4 class="dark-inverted">
          单核心最大处理能力(Mb)
        </h4>
        <span>
          {{ tmpld.requirement.maxRateMbpsByCore || '无' }}
        </span>
      </div>
      <div class="info-block-line">
        <h4 class="dark-inverted">
          收流会话数
        </h4>
        <span>
          {{ tmpld.requirement.receiveSessions || '无' }}
        </span>
      </div>
      <div class="info-block-line">
        <h4 class="dark-inverted">
          收流会话数
        </h4>
        <span>
          {{ tmpld.requirement.receiveSessions || '无' }}
        </span>
      </div>
      <div class="info-block-line">
        <h4 class="dark-inverted">
          SendAVFrameNodeCount
        </h4>
        <span>
          {{ tmpld.requirement.sendAVFrameNodeCount || '无' }}
        </span>
      </div>
      <div class="info-block-line">
        <h4 class="dark-inverted">
          RecvFrameCount
        </h4>
        <span>
          {{ tmpld.requirement.recvFrameCount || '无' }}
        </span>
      </div>
      <div class="info-block-line">
        <h4 class="dark-inverted">
          Image Tag
        </h4>
        <span>
          {{ tmpld.requirement.imageTag || '无' }}
        </span>
      </div>
      <!-- <div class="info-block-line">
        <h4 class="dark-inverted">
          Chart包名
        </h4>
        <span>
          {{ tmpld.requirement.chart || '无' }}
        </span>
      </div> -->
      <div class="info-block-line multi">
        <h4 class="dark-inverted">
          说明
        </h4>
        <span>{{ tmpld.description || '无' }}</span>
      </div>
    </div>
  </div>
</template>
