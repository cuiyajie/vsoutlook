<script setup lang="ts">
import { formats, formatKeys, type UdxInputType, v_protocols } from './Consts';
import { watchModeVFormat, useProtocolDC } from './Utils';

const mv = defineModel<UdxInputType>({
  default: {} as UdxInputType,
  local: true,
});

const props = defineProps<{
  mode: string,
}>()

const format = defineModel('format', {
  default: formatKeys[0],
  local: true,
  required: true
});

const shouldFormats = ref([...formats])
watchModeVFormat(() => props.mode, shouldFormats, 'input')

useProtocolDC(mv)

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
        <div class="column is-6">
          <VField>
            <VLabel>视频流协议</VLabel>
            <VControl>
              <VSelect
                v-model="mv.v_protocol"
                class="is-rounded"
              >
                <VOption
                  v-for="vp in v_protocols"
                  :key="vp"
                  :value="vp"
                >
                  {{ vp }}
                </VOption>
              </VSelect>
            </VControl>
          </VField>
        </div>
        <div class="column is-6">
          <VField>
            <VLabel>主视频流组播源IP（含端口）</VLabel>
            <VControl>
              <VInput
                v-model="mv.ipstream_master.v_src_address"
              />
            </VControl>
          </VField>
        </div>
        <div class="column is-6">
          <VField>
            <VLabel>主视频流组播目标IP（含端口）</VLabel>
            <VControl>
              <VInput
                v-model="mv.ipstream_master.v_dst_address"
              />
            </VControl>
          </VField>
        </div>
        <div class="column is-6">
          <VField>
            <VLabel>主音频流组播源IP（含端口）</VLabel>
            <VControl>
              <VInput
                v-model="mv.ipstream_master.a_src_address"
              />
            </VControl>
          </VField>
        </div>
        <div class="column is-6">
          <VField>
            <VLabel>主音频流组播目标IP（含端口）</VLabel>
            <VControl>
              <VInput
                v-model="mv.ipstream_master.a_dst_address"
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
              <VControl>
                <VInput
                  v-model="mv.ipstream_backup.v_src_address"
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
              <VLabel>备视频流组播目标IP（含端口）</VLabel>
              <VControl>
                <VInput
                  v-model="mv.ipstream_backup.v_dst_address"
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
              <VControl>
                <VInput
                  v-model="mv.ipstream_backup.a_src_address"
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
              <VLabel>备音频流组播目标IP（含端口）</VLabel>
              <VControl>
                <VInput
                  v-model="mv.ipstream_backup.a_dst_address"
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
              <VSelect
                v-model="format"
                class="is-rounded"
              >
                <VOption
                  v-for="f in shouldFormats"
                  :key="f.key"
                  :value="f.key"
                >
                  {{ f.value }}
                </VOption>
              </VSelect>
            </VControl>
          </VField>
        </div>
      </div>
    </div>
    <div v-if="mv.videoformat.v_compression_format || mv.videoformat.v_compression_ratio" class="form-fieldset">
      <div class="fieldset-heading">
        <h5>视频编码参数</h5>
      </div>
      <div class="columns is-multiline">
        <div v-if="mv.videoformat.v_compression_format" class="column is-6">
          <VField>
            <VLabel>编码格式</VLabel>
            <VControl>
              <VInput
                v-model="mv.videoformat.v_compression_format"
                readonly
              />
            </VControl>
          </VField>
        </div>
        <div v-if="mv.videoformat.v_compression_ratio" class="column is-6">
          <VField>
            <VLabel>压缩比</VLabel>
            <VControl>
              <VInput
                v-model="mv.videoformat.v_compression_ratio"
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
              <VSelect
                v-model="mv.audioformat.a_bits"
                class="is-rounded"
              >
                <VOption :value="16">
                  16
                </VOption>
                <VOption :value="24">
                  24
                </VOption>
                <VOption :value="32">
                  32
                </VOption>
              </VSelect>
            </VControl>
          </VField>
        </div>
        <div class="column is-4">
          <VField>
            <VLabel>采样率</VLabel>
            <VControl>
              <VInput
                v-model="mv.audioformat.a_frequency"
                readonly
              />
            </VControl>
          </VField>
        </div>
      </div>
    </div>
  </div>
</template>
