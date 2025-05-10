<script lang="ts" setup>
import { useUserSession } from "@src/stores/userSession";
import { type PlayerParams, def_player_params, type IndexedNicDetail, codec_dev_types, quality_types } from './Consts_V1';
import { VideoFormatPrefixMap } from '@src/components/partials/modals/PresetConfig/Consts';
import { changeProtocol } from './Utils_V1';
import { useSmpteParams } from "./Composables";

const mv = defineModel<PlayerParams>({
  default: def_player_params(),
  local: true,
})

const props = withDefaults(defineProps<{
  nics: IndexedNicDetail[]
  showAudio?: boolean
  showNic?: boolean
  smpte?: 'receive' | 'send'
}>(), {
  showAudio: true,
  showNic: true,
  smpte: 'send'
})

const nicsRef = computed(() => props.nics)

if (props.smpte === 'send') {
  useSmpteParams(mv, nicsRef)
}

watch(() => mv.value.smpte_params.nic_index, () => {
  const prefix = props.smpte === 'send' ? 'tx' : 'rx'
  const core = `${prefix}#${mv.value.smpte_params.nic_index}`
  mv.value.smpte_params.v_core = core
  mv.value.smpte_params.a_core = core
}, { immediate: true })

const usStore = useUserSession()
const videoFormats = computed(() => usStore.settings.video_formats || [])

const showParams = computed(() => {
  const vf = videoFormats.value.find(vf => vf.name === mv.value.videoformat_name)
  if (vf) return ['st2110-20', 'st2110-22'].includes(vf.protocol)
  return false
})

watch(() => mv.value.videoformat_name, (nv) => {
  const vf = videoFormats.value.find(vf => vf.name === nv)
  if (!vf?.protocol) return
  const prefix = VideoFormatPrefixMap[vf.protocol]
  if (prefix) {
    mv.value.stream_params.url = changeProtocol(mv.value.stream_params.url, prefix)
  }
}, { immediate: true })

</script>
<template>
  <!--Fieldset-->
  <expand-transition>
    <div v-if="mv.videoformat_name && showParams" class="form-fieldset">
      <div class="fieldset-heading">
        <h4>smpte{{ smpte === 'send' ? '发流' : '收流' }}参数</h4>
      </div>
      <div v-if="showNic" class="form-fieldset-nested-3">
        <div class="form-fieldset-nested-3 seperator">
          <div class="columns is-multiline">
            <div class="column is-6">
              <VField>
                <VLabel>{{ smpte === 'send' ? '发流' : '收流' }}网卡序号</VLabel>
                <VControl>
                  <NicDetailSelect v-model="mv.smpte_params.nic_index" :nics="nics" />
                </VControl>
              </VField>
            </div>
          </div>
        </div>
      </div>
      <div class="form-fieldset-nested-6">
        <div class="form-fieldset seperator">
          <div class="fieldset-heading">
            <h5>主路参数</h5>
          </div>
          <div class="columns is-multiline">
            <div class="column is-6">
              <AddrInputAddon v-model="mv.smpte_params.ipstream_master.v_src_address" label="视频源地址（含端口）" />
            </div>
            <div class="column is-6">
              <AddrInputAddon v-model="mv.smpte_params.ipstream_master.v_dst_address" label="视频组播地址（含端口）" />
            </div>
            <div v-if="showAudio" class="column is-6">
              <AddrInputAddon v-model="mv.smpte_params.ipstream_master.a_src_address" label="音频源地址（含端口）" />
            </div>
            <div v-if="showAudio" class="column is-6">
              <AddrInputAddon v-model="mv.smpte_params.ipstream_master.a_dst_address" label="音频组播地址（含端口）" />
            </div>
            <div class="column is-6">
              <AddrInputAddon v-model="mv.smpte_params.ipstream_master['40_src_address']" label="辅助数据源地址（含端口）" />
            </div>
            <div class="column is-6">
              <AddrInputAddon v-model="mv.smpte_params.ipstream_master['40_dst_address']" label="辅助数据组播地址（含端口）" />
            </div>
          </div>
        </div>
      </div>
      <div class="form-fieldset-nested-6 is-tail">
        <div class="form-fieldset-nested-4">
          <div class="fieldset-heading">
            <h5>备路参数</h5>
          </div>
          <div class="columns is-multiline">
            <div class="column is-6">
              <AddrInputAddon v-model="mv.smpte_params.ipstream_backup.v_src_address" label="视频源地址（含端口）" />
            </div>
            <div class="column is-6">
              <AddrInputAddon v-model="mv.smpte_params.ipstream_backup.v_dst_address" label="视频组播地址（含端口）" />
            </div>
            <div v-if="showAudio" class="column is-6">
              <AddrInputAddon v-model="mv.smpte_params.ipstream_backup.a_src_address" label="音频源地址（含端口）" />
            </div>
            <div v-if="showAudio" class="column is-6">
              <AddrInputAddon v-model="mv.smpte_params.ipstream_backup.a_dst_address" label="音频组播地址（含端口）" />
            </div>
            <div class="column is-6">
              <AddrInputAddon v-model="mv.smpte_params.ipstream_backup['40_src_address']" label="辅助数据源地址（含端口）" />
            </div>
            <div class="column is-6">
              <AddrInputAddon v-model="mv.smpte_params.ipstream_backup['40_dst_address']" label="辅助数据组播地址（含端口）" />
            </div>
          </div>
        </div>
      </div>
    </div>
  </expand-transition>
  <expand-transition>
    <div v-if="mv.videoformat_name && !showParams" class="form-fieldset">
      <div class="fieldset-heading">
        <h4>深压缩流发流参数</h4>
      </div>
      <div class="columns is-multiline">
        <div class="column is-4">
          <AddrAddonPrefix v-model="mv.stream_params.url" label="发流地址" placeholder="请输入发流地址" />
        </div>
        <div class="column is-4">
          <VField>
            <VLabel>编码设备类型</VLabel>
            <VControl>
              <VSelect
                v-model="mv.stream_params.codec_dev"
                class="is-rounded"
              >
                <VOption v-for="cdt in codec_dev_types" :key="cdt" :value="cdt">{{ cdt }}</VOption>
              </VSelect>
            </VControl>
          </VField>
        </div>
        <div class="column is-4">
          <VField>
            <VLabel>编码质量</VLabel>
            <VControl>
              <VSelect
                v-model="mv.stream_params.quality"
                class="is-rounded"
              >
                <VOption v-for="quat in quality_types" :key="quat.value" :value="quat.value">{{ quat.label }}</VOption>
              </VSelect>
            </VControl>
          </VField>
        </div>
      </div>
    </div>
  </expand-transition>
</template>
