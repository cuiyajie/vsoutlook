<script lang="ts" setup>
import { useTemplate } from "/@src/stores/template";

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
  router.push(`/app/template/${tmpld.value.id}`)
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
          DPDKCPU 核心数
        </h4>
        <span>
          {{ tmpld.requirement.dpdkCpu || '无' }}
        </span>
      </div>
      <div class="info-block-line">
        <h4 class="dark-inverted">
          内存 (GB)
        </h4>
        <span>
          {{ tmpld.requirement.memory || '无' }}
        </span>
      </div>
      <!-- <div class="info-block-line">
        <h4 class="dark-inverted">
          CPU核心
        </h4>
        <span>
          {{ tmpld.requirement.cpuCore || '无' }}
        </span>
      </div> -->
      <div class="info-block-line">
        <h4 class="dark-inverted">
          大叶内存 (GB)
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
