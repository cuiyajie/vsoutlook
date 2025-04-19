<script setup lang="ts">
import { confirm } from "@src/utils/dialog";
import { useClustNode } from "@src/stores/node";
import { useNotyf } from "@src/composable/useNotyf";
import Draggable from "vuedraggable";

const opened = ref(false);
const node = ref<ClustNode | null>(null);
const page = ref(1);
const nodeStore = useClustNode();
const notyf = useNotyf();

function fetchList() {
  if (!node.value) return;
  nodeStore.$getNodeNics(node.value.id).then((res) => {
    if (res) {
      node.value = res;
    }
  });
}

function create() {
  bus.trigger(Signal.OpenNodeEdit, {
    node: node.value,
    callbacks: {
      success: () => {
        fetchList();
      },
    },
  });
}

function remove(idx: number) {
  if (!node.value) return;
  confirm({
    title: "确认",
    content: `确定要删除收发流网卡 ${idx} 吗？`,
    onConfirm: async (hide) => {
      if (!node.value) return;
      const nic = node.value.nics[idx];
      nodeStore.$deleteNic(nic.id, node.value.id).then((res) => {
        if (res?.code === 0) {
          if (node.value) {
            node.value.nics.splice(idx, 1);
          }
          notyf.success(`删除收发流网卡 ${idx} 成功`);
        } else if (res?.error) {
          notyf.error(res.message);
        }
      });
      hide();
    },
  });
}

function edit(idx: number) {
  bus.trigger(Signal.OpenNodeEdit, {
    node: node.value,
    index: idx,
    callbacks: {
      success: () => {
        fetchList();
      },
    },
  });
}

let snapshot: NicInfo[] = [];
function getSnapshot() {
  if (!node.value) return;
  snapshot = [...node.value.nics];
}

function recorder({ oldIndex, newIndex }: { oldIndex: number; newIndex: number }) {
  if (!node.value) return;
  const target = snapshot[oldIndex]
  nodeStore.$reorderNic(node.value.id, target.id, snapshot[newIndex].position)
}

useListener(Signal.OpenNodeResourceList, async (_node: ClustNode) => {
  opened.value = true;
  node.value = _node;
  fetchList();
});
</script>
<template>
  <VModal
    size="big"
    :open="opened"
    :title="`编辑节点资源 - ${node?.ip}`"
    actions="center"
    cancel-label="关闭"
    dialog-class="resource-dialog"
    @close="opened = false"
  >
    <template #header-action>
      <div class="header-action">
        <VButton color="primary" icon="fas fa-plus" elevated @click="create">
          新建
        </VButton>
      </div>
    </template>
    <template #content>
      <div v-if="node">
        <div class="datatable-wrapper">
          <div class="table-container">
            <table class="table datatable-table is-fullwidth">
              <thead>
                <th></th>
                <th align="center">序号</th>
                <th align="center">主路网口</th>
                <th align="center">收发开关</th>
                <th align="center">备路网口</th>
                <th align="center">收发开关</th>
                <th align="center">CPU 隔离核</th>
                <th align="center">DMA</th>
                <th align="center">VF</th>
                <th>操作</th>
              </thead>
              <Draggable
                v-model="node.nics"
                tag="tbody"
                handle=".drag-area"
                ghost-class="row-ghost"
                chosen-class="row-chosen"
                item-key="id"
                @start="getSnapshot"
                @end="recorder"
              >
                <template #item="{ element: nic, index: cidx }: { element: NicInfo, index: number }">
                  <tr>
                    <td>
                      <div class="drag-area">
                        <VIconButton icon="feather:menu" circle />
                      </div>
                    </td>
                    <td>{{ cidx }}</td>
                    <td>{{ nic.nicNameMain }}</td>
                    <td>
                      <div class="is-list">
                        <VControl subcontrol class="switch-2">
                          <VSwitchBlock
                            :model-value="nic.receiveMain"
                            color="info"
                            label="收"
                          />
                        </VControl>
                        <VControl subcontrol class="switch-2">
                          <VSwitchBlock
                            :model-value="nic.sendMain"
                            color="danger"
                            label="发"
                          />
                        </VControl>
                      </div>
                    </td>
                    <td>{{ nic.nicNameBackup }}</td>
                    <td>
                      <div class="is-list">
                        <VControl subcontrol class="switch-2">
                          <VSwitchBlock
                            :model-value="nic.receiveBackup"
                            color="info"
                            label="收"
                          />
                        </VControl>
                        <VControl subcontrol class="switch-2">
                          <VSwitchBlock
                            :model-value="nic.sendBackup"
                            color="danger"
                            label="发"
                          />
                        </VControl>
                      </div>
                    </td>
                    <td>{{ nic.coreList }}</td>
                    <td>{{ nic.dmaList }}</td>
                    <td>{{ nic.vfCount }}</td>
                    <td>
                      <PresetListDropdown
                        @add="create"
                        @edit="edit(cidx)"
                        @remove="remove(cidx)"
                      />
                    </td>
                  </tr>
                </template>
              </Draggable>
            </table>
          </div>
          <VPlaceholderPage
            v-if="node.nics.length === 0"
            title=""
            subtitle="暂无收发流网卡配置"
          />
        </div>

        <!--Table Pagination-->
        <VFlexPagination
          v-if="node.nics.length > 5"
          v-model:current-page="page"
          :item-per-page="6"
          :total-items="node.nics.length"
          :max-links-displayed="7"
          no-router
          class="mt-4"
        />
      </div>
    </template>
  </VModal>
</template>
<style lang="scss">
.resource-dialog {
  &.modal.is-big .modal-content {
    max-width: 1280px;
  }

  .modal-card-body {
    min-height: 400px;
  }

  .header-action {
    display: flex;
    align-items: center;
    justify-content: flex-end;
    flex: 1;
    margin-inline-end: 60px;
  }
}

.is-dark {
  .resource-dialog.modal {
    .table {
      background: var(--dark-sidebar-light-2);
    }

    .datatable-table {
      thead {
        th:first-child {
          width: 80;
        }

        th:last-child {
          width: 80px;
        }
      }

      tr {
        &.row-ghost {
          background: var(--dark-sidebar-light-8) !important;
        }

        &.row-chosen {
          background: var(--dark-sidebar-light-2);
          border: 1px solid var(--dark-sidebar-light-12);
          border-radius: 0.75rem;
        }
      }
    }
  }
}
</style>
