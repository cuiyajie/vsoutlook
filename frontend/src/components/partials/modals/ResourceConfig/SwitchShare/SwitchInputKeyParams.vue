<script lang="ts" setup>
import { type InjectSwitchTemplateToggle, type SwitchInputKeyParams, def_switch_template_toggle } from './Consts'
import { type IndexedNicDetail } from '../Utilties/Consts_V1'
import { key_types } from '../SwitchShare/Consts'

const mv = defineModel<SwitchInputKeyParams>({
  default: {} as SwitchInputKeyParams,
  local: true,
})

defineProps<{
  index: number
  isLast: boolean
}>()

const usedSignalType = inject<Ref<number>>('switch_used_signal_type', ref(0))
const { videoSelected, videoUnSelected } = inject<InjectSwitchTemplateToggle>('switch_template_toggle', def_switch_template_toggle())
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
      <h4>第 {{ index + 1 }} 路键输入信号源&nbsp;&nbsp;索引: {{ index }}</h4>
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
      <div v-show="opened">
        <div :class="mv.key_type === 'ext_key' ? 'form-fieldset-nested-3 seperator' : 'form-fieldset-nested-4'">
          <div class="columns is-multiline">
            <div class="column is-6">
              <VField>
                <VLabel>键类型</VLabel>
                <VControl>
                  <VSelect v-model="mv.key_type" class="is-rounded">
                    <VOption v-for="kt in key_types" :key="kt.value" :value="kt.value">{{ kt.label }}</VOption>
                  </VSelect>
                </VControl>
              </VField>
            </div>
            <div class="column is-6">
              <VField>
                <VLabel>视频格式名称</VLabel>
                <VControl>
                  <VideoFormatSelect v-model="mv.videoformat_name" :used-signal-type="usedSignalType" @video-selected="videoSelected($event)" @video-unselected="videoUnSelected($event)" />
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
            <expand-transition>
              <div v-if="mv.key_type === 'int_key'" class="column is-6">
                <VField>
                  <VLabel>內键文件名称（含路径）</VLabel>
                  <VControl>
                    <VInput v-model="mv.file_name" />
                  </VControl>
                </VField>
              </div>
            </expand-transition>
            <expand-transition>
              <div v-if="mv.key_type === 'h5_key'" class="column is-6">
                <VField>
                  <VLabel>H5地址</VLabel>
                  <VControl>
                    <VInput v-model="mv.url" />
                  </VControl>
                </VField>
              </div>
            </expand-transition>
          </div>
        </div>
        <expand-transition>
          <div v-if="mv.key_type === 'ext_key'">
            <PlayerParamsForm v-model="mv" :nics="nics" :smpte="mv.key_type === 'ext_key' ? 'receive' : 'send'" />
          </div>
        </expand-transition>
      </div>
    </expand-transition>
  </div>
</template>
