<script setup lang="ts">
import { useUserSession } from "@src/stores/userSession";
import { confirm } from "@src/utils/dialog";
import { VideoFormatProtocolsMap } from "../modals/PresetConfig/Consts"
import { stringifyFps } from "../modals/PresetConfig/Utils"

const usStore = useUserSession();
const videoFormats = computed(() => usStore.settings.video_formats || []);
const page = ref(1);

function create() {
  bus.trigger(Signal.OpenNewVideoFormat, {});
}

function remove(idx: number) {
  const vft = videoFormats.value[idx];
  confirm({
    title: "确认",
    content: `确定要删除视频格式 ${vft.name} 吗？`,
    onConfirm: async (hide) => {
      videoFormats.value.splice(idx, 1);
      usStore.$updateSettings({
        key: "video_formats",
        value: JSON.stringify(videoFormats.value),
      });
      hide();
    },
  });
}

function edit(idx: number) {
  bus.trigger(Signal.OpenNewVideoFormat, {
    index: idx,
  });
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
        <table class="table datatable-table video-format-table is-fullwidth">
          <thead>
            <th align="center">模板名称</th>
            <th align="center">格式类型</th>
            <th align="center">分辨率</th>
            <th align="center">帧率</th>
            <th align="center">伽马</th>
            <th align="center">色域</th>
            <th align="center">压缩格式</th>
            <th align="center">压缩子类型</th>
            <th align="center">压缩比</th>
            <th align="center">码率</th>
            <th align="center">B帧数量</th>
            <th align="center">GOP长度</th>
            <th>操作</th>
          </thead>
          <tbody>
            <tr v-for="(vft, vidx) in videoFormats" :key="vft.name">
              <td>{{ vft.name }}</td>
              <td>{{ VideoFormatProtocolsMap[vft.protocol] }}</td>
              <td>{{ `${vft.width}x${vft.height}` }}</td>
              <td>{{ stringifyFps(vft.fps, vft.interlaced) }}</td>
              <td>{{ vft.gamma }}</td>
              <td>{{ vft.gamut }}</td>
              <td>{{ vft.compression_format }}</td>
              <td>{{ vft.compression_subtype || '-' }}</td>
              <td>{{ vft.compression_ratio || '-' }}</td>
              <td>{{ vft.bitrate_bps || '-' }}</td>
              <td>{{ vft.gop_b_frames || '-' }}</td>
              <td>{{ vft.gop_length || '-' }}</td>
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
.datatable-table.video-format-table {
  font-size: 0.875rem;

  th {
    padding: 16px 8px;
  }

  td {
    padding: 12px 8px;
  }
}
</style>
