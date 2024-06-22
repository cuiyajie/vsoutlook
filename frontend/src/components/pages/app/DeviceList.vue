<script setup lang="ts">
import type {
  VFlexTableWrapperSortFunction,
  VFlexTableWrapperFilterFunction,
} from '@src/components/base/table/VFlexTableWrapper.vue'
import { useDevices } from '@src/stores/device'
import { useClustNode } from '@src/stores/node'
import { useTemplate } from '@src/stores/template'

const deviceStore = useDevices()
const nodeStore = useClustNode();
const tmplStore = useTemplate();
const devices = computed(() => deviceStore.devices)
const loading = ref(false)

const tmplNameSorter: VFlexTableWrapperSortFunction<DeviceDetail> = ({ order, a, b }) => {
  if (order === 'asc') {
    return a.tmplName.localeCompare(b.tmplName)
  } else if (order === 'desc') {
    return b.tmplName.localeCompare(a.tmplName)
  }
  return 0
}

const tmplNameFilter: VFlexTableWrapperFilterFunction<DeviceDetail> = ({ searchTerm, row }) => {
  if (!searchTerm) return true
  return row.tmplName.toLocaleLowerCase().includes(searchTerm.toLocaleLowerCase()) || row.appName.toLocaleLowerCase().includes(searchTerm.toLocaleLowerCase())
}

const nameFilter: VFlexTableWrapperFilterFunction<DeviceDetail> = ({ searchTerm, row }) => {
  if (!searchTerm) return true
  return row.name.toLocaleLowerCase().includes(searchTerm.toLocaleLowerCase())
}

const nodeSorter: VFlexTableWrapperSortFunction<DeviceDetail> = ({ order, a, b }) => {
  if (order === 'asc') {
    return a.node.localeCompare(b.node)
  } else if (order === 'desc') {
    return b.node.localeCompare(a.node)
  }
  return 0
}

const nodeFilter: VFlexTableWrapperFilterFunction<DeviceDetail> = ({ searchTerm, row }) => {
  if (!searchTerm) return true
  return row.node.toLocaleLowerCase().includes(searchTerm.toLocaleLowerCase())
}

const updateAtSorter: VFlexTableWrapperSortFunction<DeviceDetail> = ({ order, a, b }) => {
  if (order === 'asc') {
    return a.updatedAt - b.updatedAt
  } else if (order === 'desc') {
    return b.updatedAt - a.updatedAt
  }
  return 0
}


const formatDate = (t: number) => new Date(t * 1000).toLocaleString('zh-CN', {
  year: 'numeric',
  month: '2-digit',
  day: '2-digit',
  hour: '2-digit',
  minute: '2-digit',
  second: '2-digit'
})

const columns = {
  icon: {
    label: '图标',
    align: 'center'
  },
  name: {
    label: '设备名称',
    grow: true,
    searchable: true,
    sortable: false,
    filter: nameFilter
  },
  tmplName: {
    label: '应用名称',
    grow: false,
    searchable: true,
    sortable: true,
    sort: tmplNameSorter,
    filter: tmplNameFilter
  },
  node: {
    label: '运行节点',
    grow: false,
    searchable: true,
    sortable: true,
    sort: nodeSorter,
    filter: nodeFilter
  },
  updatedAt: {
    label: '更新时间',
    grow: false,
    sortable: true,
    align: 'center',
    sort: updateAtSorter,
  },
  status: {
    label: '状态',
    sortable: true,
    searchable: true,
    align: 'center'
  },
  // input: {
  //   label: '入信号',
  //   sortable: true,
  //   searchable: true,
  //   align: 'center'
  // },
  // output: {
  //   label: '出信号',
  //   sortable: true,
  //   searchable: true,
  //   align: 'center'
  // },
  // ptp: {
  //   label: 'PTP',
  //   sortable: true,
  //   searchable: true,
  //   align: 'center'
  // },
  action: {
    label: '操作',
    align: 'center'
  }
} as const

async function refresh () {
  loading.value = true
  await deviceStore.$fetchList()
  loading.value = false
}

function deployApp() {
  nodeStore.$fetchList()
  tmplStore.$fetchList();
  bus.trigger(Signal.OpenResourceConfig, {
    callbacks: {
      success: () => {
        refresh()
      }
    }
  });
}

refresh()

</script>

<template>
  <VFlexTableWrapper
    class="devices-flex-table"
    :columns="columns"
    :data="devices"
  >
    <!--
      Here we retrieve the internal wrapperState.
      Note that we can not destructure it
    -->
    <template #default="wrapperState">
      <!-- We can place any content inside the default slot-->
      <VFlexTableToolbar>
        <template #left>
          <!-- We can bind wrapperState.searchInput to any input -->
          <VField>
            <VControl icon="feather:search">
              <input
                v-model="wrapperState.searchInput"
                type="text"
                class="input"
                placeholder="搜索"
              >
            </VControl>
          </VField>
        </template>
        <template #right>
          <VButtons>
            <VButton
              color="primary"
              raised
              @click="deployApp"
            >
              <span class="icon">
                <i
                  class="lnir lnir-plus"
                  aria-hidden="true"
                />
              </span>
              <span>新建设备</span>
            </VButton>
            <VButton
              color="primary"
              raised
              @click="refresh"
            >
              <span class="icon">
                <i
                  class="lnir lnir-reload"
                  aria-hidden="true"
                />
              </span>
              <span>刷新列表</span>
            </VButton>
            <!-- <VButton
              color="primary"
              raised
            >
              <span class="icon">
                <i
                  aria-hidden="true"
                  class="fas fa-trash-alt"
                />
              </span>
              <span>全部删除</span>
            </VButton>
          </VButtons> -->
          </VButtons>
        </template>
      </VFlexTableToolbar>

      <VLoader
        size="small"
        translucent
        :active="loading"
      >
        <!--
        The VFlexTable "data" and "columns" props
        will be inherited from parent VFlexTableWrapper
      -->
        <VFlexTable>
          <!-- Custom "name" cell content -->
          <template #body-cell="{ row, column }">
            <VIconWrap
              v-if="column.key === 'icon'"
              v-tooltip.rounded.primary="row.tmplTypeName"
              size="medium"
              class="is-tt-icon"
              :picture="`/images/tmpl/${row.tmplTypeIcon}`"
            />
            <div
              v-else-if="column.key === 'name'"
              class="dark-text"
            >
              <span>{{ row.name }}</span>
              <div v-if="row.appName" class="subcell">( {{ row.appName }} )</div>
            </div>
            <span
              v-else-if="column.key === 'tmplName'"
              class="dark-text"
            >{{ row.tmplName }}</span>
            <div
              v-else-if="column.key === 'node'"
              class="dark-text"
            >
              <span>{{ row.node }}</span>
              <div v-if="row.nodeIP" class="subcell">( {{ row.nodeIP }} )</div>
            </div>
            <span
              v-else-if="column.key === 'updatedAt'"
              class="dark-text"
            >{{ formatDate(row.updatedAt) }}</span>
            <VTag
              v-else-if="column.key === 'status'"
              rounded
              :color="row.statusInfo.color"
            >
              {{ row.statusInfo.text }}
            </VTag>
            <!-- <VTag
            v-else-if="column.key === 'input'"
            rounded
            :color="row.inputInfo.color"
          >
            {{ row.inputInfo.text }}
          </VTag>
          <VTag
            v-else-if="column.key === 'output'"
            rounded
            :color="row.outputInfo.color"
          >
            {{ row.outputInfo.text }}
          </VTag>
          <VTag
            v-else-if="column.key === 'ptp'"
            rounded
            :color="row.ptpInfo.color"
          >
            {{ row.ptpInfo.text }}
          </VTag> -->
            <DeviceListDropdown
              v-else-if="column.key === 'action'"
              :device="row"
              @refresh="refresh"
            />
          </template>
        </VFlexTable>
        <!-- content ... --->
      </VLoader>


      <VPlaceholderPage
        :class="[(wrapperState.data.length !== 0 || loading) && 'is-hidden']"
        title="没有找到符合搜索条件的设备"
        subtitle="很遗憾，看起来我们无法找到与您输入的搜索词或条件匹配的设备。请尝试不同的搜索词或条件。"
        larger
      />

      <!-- Table Pagination with wrapperState.page binded-->
      <VFlexPagination
        v-if="wrapperState.data.length > 0 && !loading"
        v-model:currentPage="wrapperState.page"
        class="mt-6"
        :item-per-page="wrapperState.limit"
        :total-items="wrapperState.total"
        :max-links-displayed="5"
        no-router
      />
    </template>
  </VFlexTableWrapper>
  <ResourceConfigModal />
  <DownloadConfirm />
</template>
<style lang="scss">
.flex-table-wrapper.devices-flex-table {
  padding: 0;
  border: none;
  border-radius: 0;
  background: none;

  .v-loader-wrapper {
    min-height: 300px;
  }

  .flex-table {
    border: 1px solid var(--fade-grey);
    border-collapse: collapse;
    border-radius: 0.75rem;
    background: var(--dark-sidebar-light-6);

    .flex-table-header {
      border-bottom: 1px solid var(--fade-grey-dark-3);

      span:nth-child(1) {
        width: percentage(90/960);
      }

      > span {
        padding: 16px 20px;
      }

      span + span {
        border-left: 1px solid var(--fade-grey-dark-3);
      }
    }

    .flex-table-cell {
      padding: 0 16px;
    }

    .flex-table-header span,
    .flex-table-item .flex-table-cell {
      justify-content: center;

      div.subcell {
        font-size: 0.8rem;
        margin-top: 2px;
        color: var(--dark-dark-text);
      }
    }

    // 图标
    .flex-table-header span:nth-child(1),
    .flex-table-item .flex-table-cell:nth-child(1) {
      flex: 0 0 percentage(90/960);

      .col-icon {
        display: flex;
        flex-direction: column;
        align-items: center;
      }
    }

    // 名称
    .flex-table-item .flex-table-cell:nth-child(2) {
      min-width: 0;
      text-align: center;

      span {
        text-overflow: ellipsis;
        overflow: hidden;
        white-space: nowrap;
      }
    }

    // 应用名称
    .flex-table-header span:nth-child(3),
    .flex-table-item .flex-table-cell:nth-child(3) {
      flex: 0 0 percentage(120/960);
    }

    // 节点
    .flex-table-header span:nth-child(4),
    .flex-table-item .flex-table-cell:nth-child(4) {
      flex: 0 0 percentage(150/960);
      text-align: center;
    }

    // 更新时间
    .flex-table-header span:nth-child(5),
    .flex-table-item .flex-table-cell:nth-child(5) {
      flex: 0 0 percentage(180/960);
    }

    // 状态
    .flex-table-header span:nth-child(6),
    .flex-table-item .flex-table-cell:nth-child(6) {
      flex: 0 0 percentage(100/960);
    }

    // 操作
    .flex-table-header span:nth-child(7),
    .flex-table-item .flex-table-cell:nth-child(7) {
      flex: 0 0 percentage(80/960);
    }
  }
}

.is-dark {
  .flex-table-wrapper.devices-flex-table {
    .flex-table {
      border-color: var(--dark-sidebar-light-12);

      .flex-table-header {
        border-color: var(--dark-sidebar-light-12);

        span {
          border-color: var(--dark-sidebar-light-12);
        }
      }
    }
  }
}
</style>
