<!-- eslint-disable vue/prop-name-casing -->
<script setup lang="ts">
import {
  type UdxOutputParamsType,
  v_protocols,
  formats,
  formatKeys,
} from '../Utilties/Consts'
import { useProtocolDC, watchModeVFormat } from '../Utilties/Utils'

const mv = defineModel<UdxOutputParamsType>({
  default: {} as UdxOutputParamsType,
  local: true,
})

const props = defineProps<{
  mode: string
  title: string
  type: 'output1' | 'output2'
  toggleTitle?: string
  isLast?: boolean
  useBackup: boolean
  m_local_ip: string
  b_local_ip: string
}>()

const OPEN = defineModel('OPEN', {
  default: false,
  local: true,
})

const format = defineModel('format', {
  default: formatKeys[1],
  local: true,
  required: true,
})

const opened = ref(false)

const isOpen = computed(() => opened.value && (!props.toggleTitle || OPEN.value))

const shouldFormats = ref([...formats])
watchModeVFormat(() => props.mode, shouldFormats, props.type)

useProtocolDC(mv)
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
      <h4>{{ title }}</h4>
      <div class="collapse-icon">
        <VIcon icon="feather:chevron-down" />
      </div>
    </div>
    <div v-if="opened && toggleTitle" class="columns is-multiline">
      <div class="column is-12">
        <VField horizontal>
          <VLabel>{{ toggleTitle }}</VLabel>
          <VControl>
            <VSwitchBlock v-model="OPEN" color="primary" />
          </VControl>
        </VField>
      </div>
    </div>
    <expand-transition>
      <div v-if="isOpen" class="form-fieldset-nested">
        <div class="form-fieldset seperator">
          <div class="fieldset-heading">
            <h5>IP流参数</h5>
          </div>
          <div class="columns is-multiline">
            <div class="column is-6">
              <VField>
                <VLabel>视频流协议</VLabel>
                <VControl>
                  <VSelect v-model="mv.v_protocol" class="is-rounded">
                    <VOption v-for="p in v_protocols" :key="p" :value="p">
                      {{ p }}
                    </VOption>
                  </VSelect>
                </VControl>
              </VField>
            </div>
            <div class="column is-6" />
            <div class="column is-6">
              <VField>
                <VLabel>主视频流组播源IP（含端口）</VLabel>
                <AddrAddon
                  v-model="mv.ipstream_master.v_src_address"
                  :host="m_local_ip"
                />
              </VField>
            </div>
            <div class="column is-6">
              <VField>
                <VLabel>主视频流组播目标IP（含端口）</VLabel>
                <VControl>
                  <VInput v-model="mv.ipstream_master.v_dst_address" />
                </VControl>
              </VField>
            </div>
            <div class="column is-6">
              <VField>
                <VLabel>主音频流组播源IP（含端口）</VLabel>
                <AddrAddon
                  v-model="mv.ipstream_master.a_src_address"
                  :host="m_local_ip"
                />
              </VField>
            </div>
            <div class="column is-6">
              <VField>
                <VLabel>主音频流组播目标IP（含端口）</VLabel>
                <VControl>
                  <VInput v-model="mv.ipstream_master.a_dst_address" />
                </VControl>
              </VField>
            </div>
            <expand-transition>
              <div v-if="useBackup" class="column is-6">
                <VField>
                  <VLabel>备视频流组播源IP（含端口）</VLabel>
                  <AddrAddon
                    v-model="mv.ipstream_backup.v_src_address"
                    :host="b_local_ip"
                  />
                </VField>
              </div>
            </expand-transition>
            <expand-transition>
              <div v-if="useBackup" class="column is-6">
                <VField>
                  <VLabel>备视频流组播目标IP（含端口）</VLabel>
                  <VControl>
                    <VInput v-model="mv.ipstream_backup.v_dst_address" />
                  </VControl>
                </VField>
              </div>
            </expand-transition>
            <expand-transition>
              <div v-if="useBackup" class="column is-6">
                <VField>
                  <VLabel>备音频流组播源IP（含端口）</VLabel>
                  <AddrAddon
                    v-model="mv.ipstream_backup.a_src_address"
                    :host="b_local_ip"
                  />
                </VField>
              </div>
            </expand-transition>
            <expand-transition>
              <div v-if="useBackup" class="column is-6">
                <VField>
                  <VLabel>备音频流组播目标IP（含端口）</VLabel>
                  <VControl>
                    <VInput v-model="mv.ipstream_backup.a_dst_address" />
                  </VControl>
                </VField>
              </div>
            </expand-transition>
          </div>
        </div>
      </div>
    </expand-transition>
    <expand-transition>
      <div v-if="isOpen" class="form-fieldset-nested">
        <div class="form-fieldset seperator">
          <div class="fieldset-heading">
            <h5>视频参数</h5>
          </div>
          <div class="columns is-multiline">
            <div class="column is-6">
              <VField>
                <VLabel>视频格式</VLabel>
                <VControl>
                  <VSelect v-model="format" class="is-rounded">
                    <VOption v-for="f in shouldFormats" :key="f.key" :value="f.key">
                      {{ f.value }}
                    </VOption>
                  </VSelect>
                </VControl>
              </VField>
            </div>
          </div>
        </div>
      </div>
    </expand-transition>
    <expand-transition>
      <div
        v-if="
          isOpen &&
            (mv.videoformat.v_compression_format || mv.videoformat.v_compression_ratio)
        "
        class="form-fieldset-nested"
      >
        <div class="form-fieldset">
          <div class="fieldset-heading">
            <h5>视频编码参数</h5>
          </div>
          <div class="columns is-multiline">
            <div v-if="mv.videoformat.v_compression_format" class="column is-6">
              <VField>
                <VLabel>编码格式</VLabel>
                <VControl>
                  <VInput v-model="mv.videoformat.v_compression_format" readonly />
                </VControl>
              </VField>
            </div>
            <div v-if="mv.videoformat.v_compression_ratio" class="column is-6">
              <VField>
                <VLabel>压缩比</VLabel>
                <VControl>
                  <VInput v-model="mv.videoformat.v_compression_ratio" readonly />
                </VControl>
              </VField>
            </div>
          </div>
        </div>
      </div>
    </expand-transition>
    <expand-transition>
      <div v-if="isOpen" class="form-fieldset-nested is-tail">
        <div class="form-fieldset">
          <div class="fieldset-heading">
            <h5>音频参数</h5>
          </div>
          <div class="columns is-multiline">
            <div class="column is-4">
              <VField>
                <VLabel>声道数</VLabel>
                <VControl>
                  <VInput v-model="mv.audioformat.a_channels_number" readonly />
                </VControl>
              </VField>
            </div>
            <div class="column is-4">
              <VField>
                <VLabel>量化比特</VLabel>
                <VControl>
                  <VInput v-model="mv.audioformat.a_bits" readonly />
                </VControl>
              </VField>
            </div>
            <div class="column is-4">
              <VField>
                <VLabel>采样率</VLabel>
                <VControl>
                  <VInput v-model="mv.audioformat.a_frequency" readonly />
                </VControl>
              </VField>
            </div>
          </div>
        </div>
      </div>
    </expand-transition>
  </div>
</template>
