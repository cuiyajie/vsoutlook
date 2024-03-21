<!-- eslint-disable vue/prop-name-casing -->
<script setup lang="ts">
import { type CodecOutputType, formatMap, formatKeys } from './Consts';
import { useProtocolDC } from "./Utils";

const mv = defineModel<CodecOutputType>({
  default: {} as CodecOutputType,
  local: true,
});

const params = defineModel('params', {
  default: {} as CodecOutputType['params'],
  local: true,
  required: true
})

const format = defineModel('format', {
  default: formatKeys[0],
  local: true,
  required: true
})

useProtocolDC(params)

defineProps<{
  m_local_ip: string,
  b_local_ip: string,
}>()
</script>
<template>
  <div class="form-body">
    <div class="form-fieldset seperator">
      <div class="fieldset-heading">
        <h5>IP流参数</h5>
      </div>
      <div class="columns is-multiline">
        <div class="column is-6">
          <VField>
            <VLabel>启用2022-7备份</VLabel>
            <VControl>
              <VSwitchBlock
                v-model="mv['g_2022-7']"
                color="primary"
              />
            </VControl>
          </VField>
        </div>
        <div class="column is-4">
          <VField>
            <VLabel>视频流协议</VLabel>
            <VControl>
              <VInput
                v-model="mv.params.v_protocol"
                readonly
              />
            </VControl>
          </VField>
        </div>
        <div class="column is-6">
          <VField>
            <VLabel>主视频流组播源IP（含端口）</VLabel>
            <AddrAddon
              v-model="mv.params.ipstream_master.v_src_address"
              :host="m_local_ip"
            />
          </VField>
        </div>
        <div class="column is-6">
          <VField>
            <VLabel>主视频流组播目标IP（含端口）</VLabel>
            <VControl>
              <VInput
                v-model="mv.params.ipstream_master.v_dst_address"
              />
            </VControl>
          </VField>
        </div>
        <div class="column is-6">
          <VField>
            <VLabel>主音频流组播源IP（含端口）</VLabel>
            <AddrAddon
              v-model="mv.params.ipstream_master.a_src_address"
              :host="m_local_ip"
            />
          </VField>
        </div>
        <div class="column is-6">
          <VField>
            <VLabel>主音频流组播目标IP（含端口）</VLabel>
            <VControl>
              <VInput
                v-model="mv.params.ipstream_master.a_dst_address"
              />
            </VControl>
          </VField>
        </div>
        <Transition name="fade-slow">
          <div
            v-if="mv['g_2022-7']"
            class="column is-6"
          >
            <VField>
              <VLabel>备视频流组播源IP（含端口）</VLabel>
              <AddrAddon
                v-model="mv.params.ipstream_backup.v_src_address"
                :host="b_local_ip"
              />
            </VField>
          </div>
        </Transition>
        <Transition name="fade-slow">
          <div
            v-if="mv['g_2022-7']"
            class="column is-6"
          >
            <VField>
              <VLabel>备视频流组播目标IP（含端口）</VLabel>
              <VControl>
                <VInput
                  v-model="mv.params.ipstream_backup.v_dst_address"
                />
              </VControl>
            </VField>
          </div>
        </Transition>
        <Transition name="fade-slow">
          <div
            v-if="mv['g_2022-7']"
            class="column is-6"
          >
            <VField>
              <VLabel>备音频流组播源IP（含端口）</VLabel>
              <AddrAddon
                v-model="mv.params.ipstream_backup.a_src_address"
                :host="b_local_ip"
              />
            </VField>
          </div>
        </Transition>
        <Transition name="fade-slow">
          <div
            v-if="mv['g_2022-7']"
            class="column is-6"
          >
            <VField>
              <VLabel>备音频流组播目标IP（含端口）</VLabel>
              <VControl>
                <VInput
                  v-model="mv.params.ipstream_backup.a_dst_address"
                />
              </VControl>
            </VField>
          </div>
        </Transition>
      </div>
    </div>
    <div class="form-fieldset">
      <div class="fieldset-heading">
        <h5>视频参数</h5>
      </div>
      <div class="columns is-multiline">
        <div class="column is-6">
          <VField>
            <VLabel>视频格式</VLabel>
            <VControl>
              <VInput
                :model-value="formatMap[format]"
                readonly
              />
            </VControl>
          </VField>
        </div>
      </div>
    </div>
    <div v-if="mv.params.videoformat.v_compression_format || mv.params.videoformat.v_compression_ratio" class="form-fieldset">
      <div class="fieldset-heading">
        <h5>视频编码参数</h5>
      </div>
      <div class="columns is-multiline">
        <div v-if="mv.params.videoformat.v_compression_format" class="column is-6">
          <VField>
            <VLabel>编码格式</VLabel>
            <VControl>
              <VInput
                v-model="mv.params.videoformat.v_compression_format"
                readonly
              />
            </VControl>
          </VField>
        </div>
        <div v-if="mv.params.videoformat.v_compression_ratio" class="column is-6">
          <VField>
            <VLabel>压缩比</VLabel>
            <VControl>
              <VInput
                v-model="mv.params.videoformat.v_compression_ratio"
                readonly
              />
            </VControl>
          </VField>
        </div>
      </div>
    </div>
    <div class="form-fieldset">
      <div class="fieldset-heading">
        <h5>音频参数</h5>
      </div>
      <div class="columns is-multiline">
        <div class="column is-4">
          <VField>
            <VLabel>声道数</VLabel>
            <VControl>
              <VInput
                v-model="mv.params.audioformat.a_channels_number"
                readonly
              />
            </VControl>
          </VField>
        </div>
        <div class="column is-4">
          <VField>
            <VLabel>量化比特</VLabel>
            <VControl>
              <VInput
                v-model="mv.params.audioformat.a_bits"
                readonly
              />
            </VControl>
          </VField>
        </div>
        <div class="column is-4">
          <VField>
            <VLabel>采样率</VLabel>
            <VControl>
              <VInput
                v-model="mv.params.audioformat.a_frequency"
                readonly
              />
            </VControl>
          </VField>
        </div>
      </div>
    </div>
  </div>
</template>
