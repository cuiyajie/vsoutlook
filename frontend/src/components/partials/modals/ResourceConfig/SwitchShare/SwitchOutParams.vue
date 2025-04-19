<script setup lang="ts">
import { type SwitchOutParams, switch_out_types, bus_signal_types_map, bus_signal_types, type SwitchInputVideoParams, type SwitchBusLevelParams, bus_sub_types } from './Consts'
import { type InjectSwitchTemplateToggle, def_switch_template_toggle } from '../SwitchShare/Consts'
import { type IndexedNicDetail } from '../Utilties/Consts_V1'

const mv = defineModel<SwitchOutParams>({
  default: {} as SwitchOutParams,
  local: true,
})

defineProps<{
  index: number
  isLast: boolean
  inputVideos: SwitchInputVideoParams[]
  busLevels: SwitchBusLevelParams[]
}>()

const usedSignalType = inject<Ref<number>>('switch_used_signal_type', ref(0))
const {
  videoSelected, videoUnSelected,
  audioSelected, audioUnSelected,
  audioMappingSelected, audioMappingUnSelected
} = inject<InjectSwitchTemplateToggle>('switch_template_toggle', def_switch_template_toggle())
const nics = inject<IndexedNicDetail[]>('switch_nics', [])
const switchTypeCategory = inject('switch_type_category')
const audioMode = inject<Ref<number>>('switch_audio_mode')

const signalTypes = computed(() => {
  if (mv.value.out_type !== 'aux') {
    return bus_signal_types.filter(bst => bst.value === 'bus_output')
  }
  return bus_signal_types.filter(bst => bst.value !== 'key')
})

const switchOutTypes = computed(() => {
  return switch_out_types.filter(sot => sot !== 'aux' || (sot === 'aux' && switchTypeCategory !== 'bcswt'))
})

const opened = ref(false)
</script>
<template>
  <div
    class="form-fieldset collapse-form"
    :class="!isLast && 'seperator'"
    :open="opened || undefined"
  >
    <div
      class="fieldset-heading collapse-header"
      tabindex="0"
      role="button"
      @keydown.space.prevent="opened = !opened"
      @click.prevent="opened = !opened"
    >
      <h4>第 {{ index + 1 }} 路输出信号&nbsp;&nbsp;索引: {{ index }}</h4>
      <div class="collapse-icon">
        <VIcon icon="feather:chevron-down" />
      </div>
    </div>
    <expand-transition>
      <div v-show="opened">
        <div class="form-fieldset-nested-3">
          <div class="columns is-multiline">
            <div class="column is-6">
              <VField>
                <VLabel>输出类型</VLabel>
                <VControl>
                  <VSelect v-model="mv.out_type" class="is-rounded">
                    <VOption v-for="sot in switchOutTypes" :key="sot" :value="sot">{{ sot }}</VOption>
                  </VSelect>
                </VControl>
              </VField>
            </div>
            <div class="column is-6">
              <VField>
                <VLabel>输出信号名称</VLabel>
                <VControl>
                  <VInput v-model="mv.out_name" />
                </VControl>
              </VField>
            </div>
            <div class="column is-6">
              <VField>
                <VLabel>输出视频格式名称</VLabel>
                <VControl>
                  <VideoFormatSelect v-model="mv.videoformat_name" :used-signal-type="usedSignalType" @video-selected="videoSelected" @video-unselected="videoUnSelected" />
                </VControl>
              </VField>
            </div>
            <div v-if="audioMode !== 0" class="column is-6">
              <VField>
                <VLabel>输出音频格式名称</VLabel>
                <VControl>
                  <AudioFormatSelect v-model="mv.audioformat_name" :videoformat="mv.videoformat_name" @audio-selected="audioSelected" @audio-unselected="audioUnSelected" />
                </VControl>
              </VField>
            </div>
            <div v-if="switchTypeCategory !== 'nmswt' && audioMode !== 0" class="column is-6">
              <VField class="field-check" label="使用的音频映射模板">
                <VControl raw subcontrol class="check-control">
                  <VCheckbox v-model="mv.mapping_checked" color="primary" circle />
                </VControl>
                <VControl>
                  <AudioMappingSelect v-model="mv.audio_mapping_name" @mapping-selected="audioMappingSelected" @mapping-unselected="audioMappingUnSelected" />
                </VControl>
              </VField>
            </div>
          </div>
        </div>
        <div class="form-fieldset">
          <div class="fieldset-heading">
            <h4>使用的信号源</h4>
          </div>
          <div class="columns is-multiline">
            <div class="column is-4">
              <VField>
                <VLabel>信源类型</VLabel>
                <VControl>
                  <VSelect v-model="mv.source_signal.signal_type" class="is-rounded">
                    <VOption v-for="st in signalTypes" :key="st.value" :value="st.value">{{ st.label }}</VOption>
                  </VSelect>
                </VControl>
              </VField>
            </div>
            <div class="column is-4">
              <VField>
                <VLabel>信源名称</VLabel>
                <VControl>
                  <SwitchOutSignalSelect
                    v-model:id="mv.source_signal.signal_id"
                    v-model:name="mv.source_signal.signal_name"
                    :title="bus_signal_types_map.get(mv.source_signal.signal_type)?.label || '信源名称'"
                    :type="mv.source_signal.signal_type"
                    :out-type="mv.out_type"
                    :input-videos="inputVideos"
                    :bus-levels="busLevels"
                  />
                </VControl>
              </VField>
            </div>
            <div class="column is-4">
              <VField>
                <VLabel>信源ID</VLabel>
                <VControl>
                  <VLabel class="is-static">{{ mv.source_signal.signal_id || '未配置' }}</VLabel>
                </VControl>
              </VField>
            </div>
          </div>
        </div>
        <PlayerParamsForm v-model="mv" :nics="nics" :show-audio="audioMode !== 0" />
      </div>
    </expand-transition>
  </div>
</template>
<style lang="scss">
@import '@src/scss/abstracts/all';
@import '@src/scss/components/forms-outer';
</style>
