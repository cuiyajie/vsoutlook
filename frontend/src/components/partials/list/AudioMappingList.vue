<script setup lang="ts">
import { useUserSession } from "@src/stores/userSession";
import { confirm } from "@src/utils/dialog";
import { numberRangeToArray } from "../modals/PresetConfig/Utils";

const usStore = useUserSession();
const audioMappings = computed(() => usStore.settings.audio_mappings || []);
const parsedAudioMappings = computed(() => audioMappings.value.map(mapping => ({
  ...mapping,
  muteArray: numberRangeToArray(mapping.mute)
})))
const page = ref(1)

function create() {
  bus.trigger(Signal.OpenNewAudioMapping, {})
}

function remove(idx: number) {
  const mapping = audioMappings.value[idx]
  confirm({
    title: "确认",
    content: `确定要删除音频映射模版 ${mapping.name} 吗？`,
    onConfirm: async (hide) => {
      audioMappings.value.splice(idx, 1)
      usStore.$updateSettings({
        key: 'audio_mappings',
        value: JSON.stringify(audioMappings.value)
      })
      hide()
    },
  });
}

function edit(idx: number) {
  bus.trigger(Signal.OpenNewAudioMapping, {
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
        <table class="table datatable-table is-fullwidth mapping-table">
          <thead>
            <th align="center">模板名称</th>
            <th align="center">声道复制</th>
            <th align="center">声道静音</th>
            <th>操作</th>
          </thead>
          <tbody>
            <tr
              v-for="(mapping, aidx) in parsedAudioMappings"
              :key="mapping.name"
            >
              <td>{{ mapping.name }}</td>
              <td>
                <div class="mapping-cols scrollbar-hide">
                  <div class="mapping-col-heading">
                    <div>源声道</div>
                    <div>目标声道</div>
                  </div>
                  <div v-for="(copy, ci) in mapping.copys" :key="ci" class="mapping-col with-border">
                    <div class="mapping-cell is-source">{{ copy[0] }}</div>
                    <div class="mapping-cell">{{ copy[1] }}</div>
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


  .mapping-cols {
    display: flex;
    align-items: center;
    text-align: left;
    font-size: 0.875rem;
    width: 100%;
    overflow-x: auto;
    padding: 8px 0;

    .mapping-col {
      padding-left: 8px;
      padding-right: 8px;

      & + .mapping-col.with-border {
        border-left: 1px solid var(--dark-sidebar-light-12);
      }

      .mapping-cell + .mapping-cell {
        margin-top: 8px;
      }
    }

    .mapping-col-heading {
      margin-right: 12px;
      line-height: 20px;
      flex: 0 0 auto;

      > div + div {
        margin-top: 8px;
      }
    }

    .mapping-cell {
      width: 18px;
      height: 18px;
      display: flex;
      align-items: center;
      justify-content: center;
      background-color: red;
      border-radius: 4px;
      box-shadow: 2px 2px 8px 0 rgba(0, 0, 0, 0.75);
      color: #000;

      &.is-source {
        background-color: #66FF00;
      }
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
