<script lang="ts" setup>
import { useUserSession } from '@src/stores/userSession'
import { def_switch_template_toggle, type InjectSwitchTemplateToggle, type SwitchOutMvParams } from './Consts'
import { type IndexedNicDetail } from '../Utilties/Consts_V1'

const usStore = useUserSession()
const templates = computed(() => usStore.settings.mv_template_list || [])

const mv = defineModel<SwitchOutMvParams>({
  default: {} as SwitchOutMvParams,
  local: true,
})

const usedSignalType = inject<Ref<number>>('switch_used_signal_type', ref(0))
const {
  videoSelected, videoUnSelected,
  audioSelected, audioUnSelected
} = inject<InjectSwitchTemplateToggle>('switch_template_toggle', def_switch_template_toggle())
const nics = inject<IndexedNicDetail[]>('switch_nics', [])
const audioMode = inject<Ref<number>>('switch_audio_mode')

</script>
<template>
  <div class="form-fieldset-nested-1">
    <div class="columns is-multiline">
      <div class="column is-6">
        <VField>
          <VLabel>是否开启多画分</VLabel>
          <VControl class="in-form">
            <VSwitchBlock
              v-model="mv.mv_is_open"
              color="primary"
              :disabled="usedSignalType !== 1"
            />
          </VControl>
        </VField>
      </div>
      <expand-transition>
        <div v-if="mv.mv_is_open" class="column is-6">
          <VField>
            <VLabel>多画分布局模板</VLabel>
            <VControl>
              <Multiselect
                v-model="mv.out_mv_template"
                placeholder="多画分布局模板"
                value-prop="path"
                label="name"
                :max-height="145"
                :options="templates"
              >
                <template #singlelabel="{ value }">
                  <div class="multiselect-single-label">
                    <span class="select-label-text">
                      {{ value.name }} ({{ value.path }})
                    </span>
                  </div>
                </template>
                <template #option="{ option }">
                  <div class="mv-tmpl-dropdown">
                    <div>
                      {{ option.name }}
                    </div>
                    <div>
                      {{ option.path }}
                    </div>
                  </div>
                </template>
              </Multiselect>
            </VControl>
          </VField>
        </div>
      </expand-transition>
      <expand-transition>
        <div v-if="mv.mv_is_open" class="column is-6">
          <VField>
            <VLabel>视频格式名称</VLabel>
            <VControl>
              <VideoFormatSelect v-model="mv.videoformat_name" :used-signal-type="usedSignalType" @video-selected="videoSelected" @video-unselected="videoUnSelected" />
            </VControl>
          </VField>
        </div>
      </expand-transition>
      <expand-transition>
        <div v-if="mv.mv_is_open" class="column is-6">
          <VField>
            <VLabel>音频格式名称</VLabel>
            <VControl>
              <AudioFormatSelect v-model="mv.audioformat_name" :videoformat="mv.videoformat_name" @audio-selected="audioSelected" @audio-unselected="audioUnSelected" />
            </VControl>
          </VField>
        </div>
      </expand-transition>
    </div>
    <expand-transition>
      <div v-if="mv.mv_is_open">
        <PlayerParamsForm v-model="mv" :nics="nics" :show-audio="audioMode === 1" />
      </div>
    </expand-transition>
  </div>
</template>
