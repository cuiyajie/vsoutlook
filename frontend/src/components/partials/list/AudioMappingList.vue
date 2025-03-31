<script setup lang="ts">
import { useUserSession } from '@src/stores/userSession'
import { confirm } from '@src/utils/dialog'
import { numberRangeToArray } from '../modals/PresetConfig/Utils'

const usStore = useUserSession()
const audioMappings = computed(() => usStore.settings.audio_mappings || [])
const parsedAudioMappings = computed(() =>
  audioMappings.value.map((mapping) => ({
    ...mapping,
    muteArray: numberRangeToArray(mapping.mute_channels),
  }))
)
const page = ref(1)

function create() {
  bus.trigger(Signal.OpenNewAudioMapping, {})
}

function remove(idx: number) {
  const mapping = audioMappings.value[idx]
  confirm({
    title: '确认',
    content: `确定要删除音频映射模版 ${mapping.name} 吗？`,
    onConfirm: async (hide) => {
      audioMappings.value.splice(idx, 1)
      usStore.$updateSettings({
        key: 'audio_mappings',
        value: JSON.stringify(audioMappings.value),
      })
      hide()
    },
  })
}

function edit(idx: number) {
  bus.trigger(Signal.OpenNewAudioMapping, {
    index: idx,
  })
}
</script>

<template>
  <div>
    <div class="datatable-toolbar">
      <VButtons>
        <VButton color="primary" icon="fas fa-plus" elevated @click="create">
          新建
        </VButton>
      </VButtons>
    </div>
    <div class="datatable-wrapper">
      <div class="table-container">
        <table class="table datatable-table is-fullwidth mapping-table">
          <thead>
            <th align="center">模板名称</th>
            <th align="center">声道复制</th>
            <th align="center">声道静音</th>
            <th>操作</th>
          </thead>
          <tbody>
            <tr v-for="(mapping, aidx) in parsedAudioMappings" :key="mapping.name">
              <td>{{ mapping.name }}</td>
              <td>
                <div class="mapping-cols scrollbar-hide">
                  <div class="mapping-col-heading">
                    <div>源声道</div>
                    <div>目标声道</div>
                  </div>
                  <div
                    v-for="(channel, ci) in mapping.copy_channels"
                    :key="ci"
                    class="mapping-col with-border"
                  >
                    <div class="mapping-cell is-source">{{ channel.src_channel }}</div>
                    <div class="mapping-cell">{{ channel.dst_channel }}</div>
                  </div>
                </div>
              </td>
              <td>
                <div class="mapping-cols scrollbar-hide">
                  <div class="mapping-col-heading">静音声道</div>
                  <div v-for="mute in mapping.muteArray" :key="mute" class="mapping-col">
                    <div class="mapping-cell">{{ mute }}</div>
                  </div>
                </div>
              </td>
              <td>
                <PresetListDropdown
                  @add="create"
                  @edit="edit(aidx)"
                  @remove="remove(aidx)"
                />
              </td>
            </tr>
          </tbody>
        </table>
      </div>
      <VPlaceholderPage
        v-if="audioMappings.length === 0"
        title=""
        subtitle="暂无音频映射模版"
      />
    </div>

    <!--Table Pagination-->
    <VFlexPagination
      v-if="audioMappings.length > 5"
      v-model:current-page="page"
      :item-per-page="6"
      :total-items="audioMappings.length"
      :max-links-displayed="7"
      no-router
      class="mt-4"
    />
  </div>
  <NewAudioMappingModal />
</template>
<style lang="scss">
.datatable-table.mapping-table {
  table-layout: fixed;

  thead {
    th:first-child {
      width: 180px;
    }

    th:nth-child(2),
    th:nth-child(3) {
      width: calc(50% - 130px);
      overflow-x: auto;
    }

    th:last-child {
      width: 80px;
    }
  }

  .scrollbar-hide {
    -ms-overflow-style: none;
    scrollbar-width: none;

    &::-webkit-scrollbar {
      display: none;
    }
  }
}
</style>
