<script setup lang="ts">
import { useUserSession } from "@src/stores/userSession";
import { confirm } from "@src/utils/dialog";

const usStore = useUserSession();
const videoFormats = computed(() => usStore.settings.video_formats || []);
const page = ref(1)

function create() {
  bus.trigger(Signal.OpenNewVideoFormat, {})
}

function remove(idx: number) {
  const vft = videoFormats.value[idx]
  confirm({
    title: "确认",
    content: `确定要删除视频格式 ${vft.name} 吗？`,
    onConfirm: async (hide) => {
      videoFormats.value.splice(idx, 1)
      usStore.$updateSettings({
        key: 'video_formats',
        value: JSON.stringify(videoFormats.value)
      })
      hide()
    },
  });
}

function edit(idx: number) {
  bus.trigger(Signal.OpenNewVideoFormat, {
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
        <table class="table datatable-table is-fullwidth">
          <thead>
            <th align="center">模板名称</th>
            <th align="center">格式类型</th>
            <th align="center">分辨率</th>
            <th align="center">帧率</th>
            <th align="center">伽马</th>
            <th align="center">色域</th>
            <th align="center">编码格式</th>
            <th align="center">视频码率</th>
            <th align="center">B 帧数量</th>
            <th align="center">GOP 长度</th>
            <th>操作</th>
          </thead>
          <tbody>
            <tr
              v-for="(vft, vidx) in videoFormats"
              :key="vft.name"
            >
              <td>{{ vft.name }}</td>
              <td>{{ vft.format }}</td>
              <td>{{ vft.resolution }}</td>
              <td>{{ vft.frameRate }}</td>
              <td>{{ vft.gamma }}</td>
              <td>{{ vft.colorSpace }}</td>
              <td>{{ vft.encodeFormat }}</td>
              <td>{{ vft.bitRate }}</td>
              <td>{{ vft.bframe }}</td>
              <td>{{ vft.gop }}</td>
              <td>
                <PresetListDropdown
                  @add="create"
                  @edit="edit(vidx)"
                  @remove="remove(vidx)"
                />
              </td>
            </tr>
          </tbody>
        </table>
      </div>
      <VPlaceholderPage
        v-if="videoFormats.length === 0"
        title=""
        subtitle="暂无视频格式数据"
      />
    </div>

    <!--Table Pagination-->
    <VFlexPagination
      v-if="videoFormats.length > 5"
      v-model:current-page="page"
      :item-per-page="6"
      :total-items="videoFormats.length"
      :max-links-displayed="7"
      no-router
      class="mt-4"
    />
  </div>
  <NewVideoFormatModal />
</template>
<style lang="scss">
.is-navbar {
  .datatable-toolbar {
    padding-top: 30px;
  }
}

.datatable-toolbar {
  display: flex;
  align-items: center;
  margin-bottom: 20px;

  &.is-reversed {
    flex-direction: row-reverse;
  }

  .field {
    margin-bottom: 0;

    .control {
      .button {
        color: var(--light-text);

        &:hover,
        &:focus {
          background: var(--primary);
          border-color: var(--primary);
          color: var(--primary--color-invert);
        }
      }
    }
  }

  .buttons {
    margin-left: auto;
    margin-bottom: 0;

    .v-button {
      margin-bottom: 0;
    }
  }
}

.is-dark {
  .datatable-toolbar {
    .field {
      .control {
        .button {
          background: var(--dark-sidebar) !important;
          color: var(--light-text);

          &:hover,
          &:focus {
            background: var(--primary) !important;
            border-color: var(--primary) !important;
            color: var(--smoke-white) !important;
          }
        }
      }
    }
  }
}

.datatable-wrapper {
  width: 100%;

  .table-container {
    overflow: visible;
  }

  .datatable-container {
    background: var(--white);
    border: none !important;
    overflow-x: auto;

    .table,
    table {
      width: 100%;
    }

    &::-webkit-scrollbar {
      height: 8px !important;
    }

    &::-webkit-scrollbar-thumb {
      border-radius: 10px !important;
      background: rgb(0 0 0 / 20%) !important;
    }
  }
}

.datatable-table {
  border: 1px solid var(--fade-grey);
  border-collapse: collapse;
  border-radius: 0.75rem;
  text-align: center;

  thead th {
    border-width: 1px;
  }

  th {
    padding: 16px 20px;
    font-family: var(--font-alt);
    font-size: 0.8rem;
    color: var(--dark-text);
    text-transform: uppercase;
    border: 1px solid var(--fade-grey);
    font-weight: 600;

    &:last-child {
      text-align: right;
    }
  }

  td {
    font-family: var(--font);
    font-weight: 400;
    vertical-align: middle;
    padding: 12px 20px;
    border-bottom: 1px solid var(--fade-grey);
    text-align: center;

    &:last-child {
      text-align: right;
    }

    &.datatables-empty {
      opacity: 0;
    }
  }

  .light-text {
    color: var(--light-text);
  }

  .flex-media {
    display: flex;
    align-items: center;

    .meta {
      margin-left: 10px;
      line-height: 1.3;

      span {
        display: block;
        font-size: 0.8rem;
        color: var(--light-text);
        font-family: var(--font);

        &:first-child {
          font-family: var(--font-alt);
          color: var(--dark-text);
        }
      }
    }
  }

  .row-action {
    display: flex;
    justify-content: flex-end;
  }

  .checkbox {
    padding: 0;
  }

  .product-photo {
    width: 80px;
    height: 80px;
    object-fit: contain;
  }

  .file-icon {
    width: 46px;
    height: 46px;
    object-fit: contain;
  }

  .drinks-icon {
    display: block;
    max-width: 48px;
    border-radius: var(--radius-rounded);
    border: 1px solid var(--fade-grey);
  }

  .negative-icon,
  .positive-icon {
    svg {
      height: 16px;
      width: 16px;
    }
  }

  .positive-icon {
    .iconify {
      color: var(--success);

      * {
        stroke-width: 4px;
      }
    }
  }

  .negative-icon {
    &.is-danger {
      .iconify {
        color: var(--danger) !important;
      }
    }

    .iconify {
      color: var(--light-text);

      * {
        stroke-width: 4px;
      }
    }
  }

  .price {
    color: var(--dark-text);
    font-weight: 500;

    &::before {
      content: '$';
    }

    &.price-free {
      color: var(--light-text);
    }
  }

  .status {
    display: flex;
    align-items: center;

    &.is-available {
      i {
        color: var(--success);
      }
    }

    &.is-busy {
      i {
        color: var(--danger);
      }
    }

    &.is-offline {
      i {
        color: var(--light-text);
      }
    }

    i {
      margin-right: 8px;
      font-size: 8px;
    }

    span {
      font-family: var(--font);
      font-size: 0.9rem;
      color: var(--light-text);
    }
  }
}

.is-dark {
  .datatable-wrapper {
    .datatable-container {
      border-color: var(--dark-sidebar-light-12);
      background: var(--dark-sidebar-light-6);
    }
  }

  .datatable-table {
    border-color: var(--dark-sidebar-light-12);

    th,
    td {
      border-color: var(--dark-sidebar-light-12);
      color: var(--dark-dark-text);
    }

    .drinks-icon {
      border-color: var(--dark-sidebar-light-12);
    }
  }
}
</style>
