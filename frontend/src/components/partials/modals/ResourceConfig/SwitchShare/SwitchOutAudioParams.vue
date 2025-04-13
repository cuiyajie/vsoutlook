<script lang="ts" setup>
import { type IndexedNicDetail } from '../Utilties/Consts_V1';
import { def_switch_template_toggle, type SwitchOutAudioParams, type InjectSwitchTemplateToggle, type SwitchInputVideoParams, type SwitchOutParams } from './Consts'

const mv = defineModel<SwitchOutAudioParams>({
  default: {} as SwitchOutAudioParams,
  local: true,
})

const props = defineProps<{
  inputVideos: SwitchInputVideoParams[],
  outParams: SwitchOutParams[],
}>()

const { audioSelected, audioUnSelected } = inject<InjectSwitchTemplateToggle>('switch_template_toggle', def_switch_template_toggle())
const nics = inject<IndexedNicDetail[]>('switch_nics', [])
const usedSignalType = inject<Ref<number>>('switch_used_signal_type', ref(0))

const showSmpte = computed(() => {
  let show = false
  props.outParams.forEach(outParam => {
    if (outParam.source_signal.signal_name === 'level0-pvwout-final') {
      show = true
    }
  })
  return show
})

</script>
<template>
  <div class="form-fieldset-nested-1">
    <div class="columns is-multiline">
      <div class="column is-6">
        <VField>
          <VLabel>直通音频源名称</VLabel>
          <VControl>
            <SwitchInputSignalSelect
              v-model:id="mv.pgm_src_signal_id"
              v-model:name="mv.src_signal_name"
              title="直通音频源"
              type="video"
              :input-keys="[]"
              :input-videos="inputVideos"
              :bus-levels="[]"
              :level-index="-1"
            />
          </VControl>
        </VField>
      </div>
      <div class="column is-6">
        <VField>
          <VLabel>直通音频源ID</VLabel>
          <VControl>
            <VLabel class="is-static">{{ mv.pgm_src_signal_id || '未配置' }}</VLabel>
          </VControl>
        </VField>
      </div>
      <div v-if="usedSignalType !== 1" class="column is-6">
        <VField>
          <VLabel>音频输入使用的网卡序号</VLabel>
          <VControl>
            <NicDetailSelect v-model="mv.in_nic_index" :nics="nics" />
          </VControl>
        </VField>
      </div>
      <div class="column is-6">
        <VField>
          <VLabel>输出音频格式名称</VLabel>
          <VControl>
            <AudioFormatSelect v-model="mv.audioformat_name" :videoformat="mv.videoformat_name" @audio-selected="audioSelected" @audio-unselected="audioUnSelected" />
          </VControl>
        </VField>
      </div>
    </div>
    <expand-transition>
      <div v-if="showSmpte" class="form-fieldset is-tail">
        <div class="fieldset-heading">
          <h4>smpte发流参数</h4>
        </div>
        <div class="form-fieldset-nested-3">
          <div class="columns is-multiline">
            <div class="column is-6">
              <VField>
                <VLabel>发流网卡序号</VLabel>
                <VControl>
                  <NicDetailSelect v-model="mv.smpte_params.nic_index" :nics="nics" />
                </VControl>
              </VField>
            </div>
          </div>
        </div>
        <div class="form-fieldset-nested-3">
          <div class="fieldset-heading is-nested">
            <h5>主路参数</h5>
          </div>
          <div class="columns is-multiline">
            <div class="column is-6">
              <AddrInputAddon v-model="mv.smpte_params.ipstream_master.a_src_address" label="音频源地址（含端口）" />
            </div>
            <div class="column is-6">
              <AddrInputAddon v-model="mv.smpte_params.ipstream_master.a_src_address" label="音频组播地址（含端口）" />
            </div>
          </div>
        </div>
        <div class="form-fieldset-nested-3">
          <div class="fieldset-heading is-nested">
            <h5>备路参数</h5>
          </div>
          <div class="columns is-multiline">
            <div class="column is-6">
              <AddrInputAddon v-model="mv.smpte_params.ipstream_backup.a_src_address" label="音频源地址（含端口）" />
            </div>
            <div class="column is-6">
              <AddrInputAddon v-model="mv.smpte_params.ipstream_backup.a_src_address" label="音频组播地址（含端口）" />
            </div>
          </div>
        </div>
      </div>
    </expand-transition>
  </div>
</template>
