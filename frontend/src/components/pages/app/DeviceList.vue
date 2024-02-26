<script setup lang="ts">
import type {
  VFlexTableWrapperSortFunction,
  VFlexTableWrapperFilterFunction,
} from '/@src/components/base/table/VFlexTableWrapper.vue'
import { useDevices } from '/@src/stores/device'
import { useClustNode } from '/@src/stores/node'
import { useTemplate } from '/@src/stores/template'

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

const tmplTypeSorter: VFlexTableWrapperSortFunction<DeviceDetail> = ({ order, a, b }) => {
  if (order === 'asc') {
    return a.tmplTypeName.localeCompare(b.tmplTypeName)
  } else if (order === 'desc') {
    return b.tmplTypeName.localeCompare(a.tmplTypeName)
  }
  return 0
}

const tmplNameFilter: VFlexTableWrapperFilterFunction<DeviceDetail> = ({ searchTerm, row }) => {
  if (!searchTerm) return true
  return row.tmplName.toLocaleLowerCase().includes(searchTerm.toLocaleLowerCase())
}

const nameFilter: VFlexTableWrapperFilterFunction<DeviceDetail> = ({ searchTerm, row }) => {
  if (!searchTerm) return true
  return row.name.toLocaleLowerCase().includes(searchTerm.toLocaleLowerCase())
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
    grow: true,
    searchable: true,
    sortable: true,
    sort: tmplNameSorter,
    filter: tmplNameFilter
  },
  tmplTypeName: {
    label: '类型',
    sortable: true,
    searchable: true,
    sorter: tmplTypeSorter
  },
  updatedAt: {
    label: '更新时间',
    grow: true,
    sortable: true,
    align: 'center'
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
  bus.trigger(Signal.OpenResourceConfig);
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
          </vbuttons>
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
              size="medium"
              class="is-tt-icon"
              :picture="`/images/tmpl/${row.tmplTypeIcon}`"
            />
            <span
              v-else-if="column.key === 'name'"
              class="dark-text"
            >{{ row.name }}</span>
            <span
              v-else-if="column.key === 'tmplType'"
              class="dark-text"
            >{{ row.tmplTypeName }}</span>
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
        width: 120px;
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

    .flex-table-header span:nth-child(1),
    .flex-table-item .flex-table-cell:nth-child(1) {
      flex: 0 0 90px;
    }

    .flex-table-header span:nth-child(3),
    .flex-table-item .flex-table-cell:nth-child(3) {
    }

    .flex-table-header span:nth-child(4),
    .flex-table-item .flex-table-cell:nth-child(4) {
    }

    .flex-table-header span:nth-child(6),
    .flex-table-item .flex-table-cell:nth-child(6) {
      flex: 0 0 100px;
    }

    .flex-table-header span:nth-child(7),
    .flex-table-item .flex-table-cell:nth-child(7) {
      flex: 0 0 80px;
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
