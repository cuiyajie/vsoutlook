<script setup lang="ts">
import { useUserSession } from "@src/stores/userSession";
import { confirm } from "@src/utils/dialog";

const usStore = useUserSession();
const audioFormats = computed(() => usStore.settings.audio_formats || []);
const page = ref(1)

function create() {
  bus.trigger(Signal.OpenNewAudioFormat, {})
}

function remove(idx: number) {
  const adf = audioFormats.value[idx]
  confirm({
    title: "确认",
    content: `确定要删除视频格式 ${adf.name} 吗？`,
    onConfirm: async (hide) => {
      audioFormats.value.splice(idx, 1)
      usStore.$updateSettings({
        key: 'audio_formats',
        value: JSON.stringify(audioFormats.value)
      })
      hide()
    },
  });
}

function edit(idx: number) {
  bus.trigger(Signal.OpenNewAudioFormat, {
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
            <th align="center">声道数量</th>
            <th align="center">量化 bit</th>
            <th align="center">采样率</th>
            <th align="center">编码格式</th>
            <th align="center">子类型</th>
            <th align="center">发包间隔</th>
            <th align="center">音频码率</th>
            <th>操作</th>
          </thead>
          <tbody>
            <tr
              v-for="(adf, aidx) in audioFormats"
              :key="adf.name"
            >
              <td>{{ adf.name }}</td>
              <td>{{ adf.channels }}</td>
              <td>{{ adf.quantBits }}</td>
              <td>{{ adf.sampleRate }}</td>
              <td>{{ adf.encodeFormat }}</td>
              <td>{{ adf.childType }}</td>
              <td>{{ adf.txInterval }}</td>
              <td>{{ adf.bitRate }}</td>
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
        v-if="audioFormats.length === 0"
        title=""
        subtitle="暂无音频格式数据"
      />
    </div>

    <!--Table Pagination-->
    <VFlexPagination
      v-if="audioFormats.length > 5"
      v-model:current-page="page"
      :item-per-page="6"
      :total-items="audioFormats.length"
      :max-links-displayed="7"
      no-router
      class="mt-4"
    />
  </div>
  <NewAudioFormatModal />
</template>
<style lang="scss">
</style>
