<script lang="ts" setup>
import { useUserSession } from "@src/stores/userSession";
import { VideoFormatPrefixMap } from '@src/components/partials/modals/PresetConfig/Consts';
import { changeProtocol } from '../Utilties/Utils_V1';
import { codec_dev_types, quality_types } from "../Utilties/Consts_V1";
import { def_keyfill_player_params, type KeyFillPlayerParams } from "./Consts";

const mv = defineModel<KeyFillPlayerParams>({
  default: def_keyfill_player_params(),
  local: true,
})

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
    mv.value.stream_params.key_params.url = changeProtocol(mv.value.stream_params.key_params.url, prefix)
    mv.value.stream_params.fill_params.url = changeProtocol(mv.value.stream_params.fill_params.url, prefix)
  }
}, { immediate: true })

const subSections: {
  key: 'key_params' | 'fill_params',
  label: string
}[] = [{
  key: 'key_params',
  label: 'Key 信号'
}, {
  key: 'fill_params',
  label: 'Fill 信号'
}]

</script>
<template>
  <!--Fieldset-->
  <expand-transition>
    <div v-if="mv.videoformat_name && showParams" class="form-fieldset" style="padding-bottom: 0;">
      <div class="fieldset-heading">
        <h4>smpte收流参数</h4>
      </div>
      <div v-for="sec in subSections" :key="sec.key">
        <div class="fieldset-subheading" :class="sec.key ==='fill_params' && 'is-sibling'">
          <h5>{{ sec.label }}</h5>
        </div>
        <div class="form-fieldset-nested-6">
          <div class="form-fieldset seperator">
            <div class="fieldset-heading">
              <h5>主路参数</h5>
            </div>
            <div class="columns is-multiline">
              <div class="column is-6">
                <AddrInputAddon v-model="mv.smpte_params[sec.key].ipstream_master.v_src_address" label="视频源地址（含端口）" />
              </div>
              <div class="column is-6">
                <AddrInputAddon v-model="mv.smpte_params[sec.key].ipstream_master.v_dst_address" label="视频组播地址（含端口）" />
              </div>
              <div class="column is-6">
                <AddrInputAddon v-model="mv.smpte_params[sec.key].ipstream_master.a_src_address" label="音频源地址（含端口）" />
              </div>
              <div class="column is-6">
                <AddrInputAddon v-model="mv.smpte_params[sec.key].ipstream_master.a_dst_address" label="音频组播地址（含端口）" />
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
                <AddrInputAddon v-model="mv.smpte_params[sec.key].ipstream_backup.v_src_address" label="视频源地址（含端口）" />
              </div>
              <div class="column is-6">
                <AddrInputAddon v-model="mv.smpte_params[sec.key].ipstream_backup.v_dst_address" label="视频组播地址（含端口）" />
              </div>
              <div class="column is-6">
                <AddrInputAddon v-model="mv.smpte_params[sec.key].ipstream_backup.a_src_address" label="音频源地址（含端口）" />
              </div>
              <div class="column is-6">
                <AddrInputAddon v-model="mv.smpte_params[sec.key].ipstream_backup.a_dst_address" label="音频组播地址（含端口）" />
              </div>
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
      <div v-for="sec in subSections" :key="sec.key">
        <div class="fieldset-subheading" :class="sec.key ==='fill_params' && 'is-sibling'">
          <h5>{{ sec.label }}</h5>
        </div>
        <div class="columns is-multiline">
          <div class="column is-6">
            <AddrAddonPrefix v-model="mv.stream_params[sec.key].url" label="发流地址" placeholder="请输入发流地址" />
          </div>
          <div class="column is-3">
            <VField>
              <VLabel>编码设备类型</VLabel>
              <VControl>
                <VSelect
                  v-model="mv.stream_params[sec.key].codec_dev"
                  class="is-rounded"
                >
                  <VOption v-for="cdt in codec_dev_types" :key="cdt" :value="cdt">{{ cdt }}</VOption>
                </VSelect>
              </VControl>
            </VField>
          </div>
          <div class="column is-3">
            <VField>
              <VLabel>编码设备索引</VLabel>
              <VControl>
                <VInputNumber
                  v-model="mv.stream_params[sec.key].codec_dev_index"
                  centered
                  :min="0"
                  :step="1"
                />
              </VControl>
            </VField>
          </div>
        </div>
      </div>
    </div>
  </expand-transition>
</template>
