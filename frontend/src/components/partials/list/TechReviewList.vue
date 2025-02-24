<script setup lang="ts">
import { useUserSession } from "@src/stores/userSession";
import { confirm } from "@src/utils/dialog";
import { VideoReviewKeys, AudioReviewKeys } from "@src/components/partials/modals/PresetConfig/Consts";

const usStore = useUserSession();
const techReviews = computed(() => usStore.settings.tech_reviews || []);
const techReviewsRef = computed(() => {
  return techReviews.value.map((tr: TechReview) => {
    const videoRules: Record<string, VideoReviewRule> = {}
    const audioRules: Record<string, AudioReviewRule> = {}
    tr.videoRules?.forEach((vr) => {
      videoRules[vr.name] = vr
    })
    tr.audioRules?.forEach((ar) => {
      audioRules[ar.name] = ar
    })
    return {
      ...tr,
      videoRules,
      audioRules
    }
  })
})
const page = ref(1)

function create() {
  bus.trigger(Signal.OpenNewTechReview, {})
}

function remove(idx: number) {
  const tr = techReviews.value[idx]
  confirm({
    title: "确认",
    content: `确定要删除视音频技审模版 ${tr.name} 吗？`,
    onConfirm: async (hide) => {
      techReviews.value.splice(idx, 1)
      usStore.$updateSettings({
        key: 'tech_reviews',
        value: JSON.stringify(techReviews.value)
      })
      hide()
    },
  });
}

function edit(idx: number) {
  bus.trigger(Signal.OpenNewTechReview, {
    index: idx
  })
}

</script>

<template>
  <div>
    <div class="datatable-toolbar">
      <VButtons>
        <VButton
          color="primary"
          icon="fas fa-plus"
          elevated
          @click="create"
        >
          新建
        </VButton>
      </VButtons>
    </div>
    <div class="datatable-wrapper">
      <div class="table-container">
        <table class="table datatable-table is-fullwidth tech-review-table">
          <thead>
            <th align="center">模版名称</th>
            <th v-for="vk in VideoReviewKeys" :key="vk" align="center">{{ vk }}</th>
            <th v-for="ak in AudioReviewKeys" :key="ak" align="center">{{ ak }}</th>
            <th>操作</th>
          </thead>
          <tbody>
            <tr
              v-for="(tr, tidx) in techReviewsRef"
              :key="tr.name"
            >
              <td>{{ tr.name }}</td>
              <td v-for="vk in VideoReviewKeys" :key="vk">
                <template v-if="tr.videoRules?.[vk]">
                  <span v-if="tr.videoRules[vk].threshold">阈值 {{ tr.videoRules[vk].threshold }}</span>
                  <span v-if="tr.videoRules[vk].duration">时长 {{ tr.videoRules[vk].duration }} 帧</span>
                  <span v-if="tr.videoRules[vk].missingDuration">时长 {{ tr.videoRules[vk].missingDuration }} ms</span>
                </template>
                <span v-else>-</span>
              </td>
              <td v-for="ak in AudioReviewKeys" :key="ak">
                <template v-if="tr.audioRules?.[ak]">
                  <span v-if="tr.audioRules[ak].channels">声道 {{ tr.audioRules[ak].channels }}</span>
                  <span v-if="tr.audioRules[ak].duration">时长 {{ tr.audioRules[ak].duration }} 帧</span>
                  <span v-if="tr.audioRules[ak].threshold">阈值 {{ tr.audioRules[ak].threshold }}</span>
                  <span v-if="tr.audioRules[ak].missingDuration">时长 {{ tr.audioRules[ak].missingDuration }} ms</span>
                </template>
                <span v-else>-</span>
              </td>
              <td>
                <PresetListDropdown
                  @add="create"
                  @edit="edit(tidx)"
                  @remove="remove(tidx)"
                />
              </td>
            </tr>
          </tbody>
        </table>
      </div>
      <VPlaceholderPage
        v-if="techReviews.length === 0"
        title=""
        subtitle="暂无视音频技审模版"
      />
    </div>

    <!--Table Pagination-->
    <VFlexPagination
      v-if="techReviews.length > 5"
      v-model:current-page="page"
      :item-per-page="6"
      :total-items="techReviews.length"
      :max-links-displayed="7"
      no-router
      class="mt-4"
    />
  </div>
  <NewTechReviewModal />
</template>
<style lang="scss">
.datatable-table.tech-review-table {

  td {
    font-size: 0.8rem;

    > span {
      display: block;

      & + span {
        margin-top: 0.25rem;
      }
    }
  }
}
</style>
