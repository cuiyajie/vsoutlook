<script lang="ts" setup>
import { type InjectSwitchTemplateToggle, type SwitchInputVideoParams, def_switch_template_toggle } from './Consts'
import { type IndexedNicDetail } from '../Utilties/Consts_V1'

const mv = defineModel<SwitchInputVideoParams>({
  default: {} as SwitchInputVideoParams,
  local: true,
})

defineProps<{
  index: number
  isLast: boolean
}>()

const usedSignalType = inject<Ref<number>>('switch_used_signal_type', ref(0))
const { videoSelected, videoUnSelected, audioSelected, audioUnSelected } = inject<InjectSwitchTemplateToggle>('switch_template_toggle', def_switch_template_toggle())
const nics = inject<IndexedNicDetail[]>('switch_nics', [])

const emit = defineEmits<{
  (e: 'remove'): void
}>()

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
      <h4>第 {{ index + 1 }} 路视频输入信号源&nbsp;&nbsp;索引: {{ index }}</h4>
      <div class="collapse-icons">
        <div
          class="collapse-icon is-close-hidden"
          role="button"
          tabindex="0"
          @click.prevent.stop="emit('remove')"
          @keydown.space.prevent.stop="emit('remove')"
        >
          <VIcon icon="feather:x" />
        </div>
        <div class="collapse-icon">
          <VIcon icon="feather:chevron-down" />
        </div>
      </div>
    </div>
    <expand-transition>
      <div v-if="opened">
        <div class="form-fieldset-nested-3 seperator">
          <div class="columns is-multiline">
            <div class="column is-4">
              <VField>
                <VLabel>视频格式名称</VLabel>
                <VControl>
                  <VideoFormatSelect v-model="mv.videoformat_name" :used-signal-type="usedSignalType" @video-selected="videoSelected($event)" @video-unselected="videoUnSelected($event)" />
                </VControl>
              </VField>
            </div>
            <div class="column is-4">
              <VField>
                <VLabel>输入视频是否强制使用指定的色域</VLabel>
                <VControl class="in-form">
                  <VSwitchBlock
                    v-model="mv.force_use_videoformat"
                    color="primary"
                  />
                </VControl>
              </VField>
            </div>
            <div class="column is-4">
              <VField>
                <VLabel>音频格式名称</VLabel>
                <VControl>
                  <AudioFormatSelect v-model="mv.audioformat_name" :videoformat="mv.videoformat_name" @audio-selected="audioSelected($event)" @audio-unselected="audioUnSelected($event)" />
                </VControl>
              </VField>
            </div>
            <div class="column is-6">
              <VField>
                <VLabel>信源ID</VLabel>
                <VControl>
                  <VLabel class="is-static">{{ mv.signal_id }}</VLabel>
                </VControl>
              </VField>
            </div>
            <div class="column is-6">
              <VField>
                <VLabel>信源名称</VLabel>
                <VControl>
                  <VInput v-model="mv.signal_name" />
                </VControl>
              </VField>
            </div>
          </div>
        </div>
        <PlayerParamsForm v-model="mv" :nics="nics" :show-nic="false" smpte="receive" />
      </div>
    </expand-transition>
  </div>
</template>
