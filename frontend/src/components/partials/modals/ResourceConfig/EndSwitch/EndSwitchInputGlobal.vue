<script setup lang="ts">
import { getFormat, useFormat, useProtocolDC } from '../Utilties/Utils'
import {
  formats,
  v_protocols,
  type EndSwitchInputParamsType,
  type BCSwitchInputParamsType,
} from '../Utilties/Consts'

const mv = defineModel<EndSwitchInputParamsType | BCSwitchInputParamsType>({
  default: {} as EndSwitchInputParamsType,
  local: true,
})

const format = useFormat(mv, getFormat(mv.value))
useProtocolDC(mv)
</script>
<template>
  <div class="form-fieldset seperator">
    <div class="fieldset-heading">
      <h4>全局参数</h4>
    </div>
    <div class="form-fieldset-nested-3">
      <div class="form-fieldset seperator">
        <div class="columns is-multiline">
          <div class="column is-3">
            <VField>
              <VLabel>启用2022-7备份</VLabel>
              <VControl>
                <VSwitchBlock v-model="mv['g_2022-7']" color="primary" />
              </VControl>
            </VField>
          </div>
        </div>
      </div>
    </div>
    <div class="form-fieldset-nested-3">
      <div class="form-fieldset seperator">
        <div class="fieldset-heading">
          <h5>视频参数</h5>
        </div>
        <div class="columns is-multiline">
          <div class="column is-3">
            <VField>
              <VLabel>视频流协议</VLabel>
              <VControl>
                <VSelect v-model="mv.v_protocol" class="is-rounded">
                  <VOption v-for="vp in v_protocols" :key="vp" :value="vp">
                    {{ vp }}
                  </VOption>
                </VSelect>
              </VControl>
            </VField>
          </div>
          <div class="column is-6">
            <VField>
              <VLabel>视频格式</VLabel>
              <VControl>
                <VSelect v-model="format" class="is-rounded">
                  <VOption v-for="f in formats" :key="f.key" :value="f.key">
                    {{ f.value }}
                  </VOption>
                </VSelect>
              </VControl>
            </VField>
          </div>
        </div>
      </div>
    </div>
    <div
      v-if="mv.videoformat.v_compression_format || mv.videoformat.v_compression_ratio"
      class="form-fieldset-nested-3"
    >
      <div class="form-fieldset seperator">
        <div class="fieldset-heading">
          <h5>视频编码参数</h5>
        </div>
        <div class="columns is-multiline">
          <div v-if="mv.videoformat.v_compression_format" class="column is-6">
            <VField>
              <VLabel>视频编码格式</VLabel>
              <VControl>
                <VInput v-model="mv.videoformat.v_compression_format" readonly />
              </VControl>
            </VField>
          </div>
          <div v-if="mv.videoformat.v_compression_ratio" class="column is-6">
            <VField>
              <VLabel>视频压缩比</VLabel>
              <VControl>
                <VInput v-model="mv.videoformat.v_compression_ratio" readonly />
              </VControl>
            </VField>
          </div>
        </div>
      </div>
    </div>
    <div class="form-fieldset-nested-3 is-tail">
      <div class="form-fieldset">
        <div class="fieldset-heading">
          <h5>音频参数</h5>
        </div>
        <div class="columns is-multiline">
          <div class="column is-4">
            <VField>
              <VLabel>声道数</VLabel>
              <VControl>
                <VInputNumber
                  v-model="mv.audioformat.a_channels_number"
                  centered
                  :min="0"
                  :max="64"
                  :step="1"
                />
              </VControl>
            </VField>
          </div>
          <div class="column is-4">
            <VField>
              <VLabel>量化比特</VLabel>
              <VControl>
                <VSelect v-model="mv.audioformat.a_bits" class="is-rounded">
                  <VOption :value="16"> 16 </VOption>
                  <VOption :value="24"> 24 </VOption>
                  <VOption :value="32"> 32 </VOption>
                </VSelect>
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
  </div>
</template>
