<script setup lang="ts">
import { formats, type Config1, formatKeys } from './Consts';

const mv = defineModel<Config1>({
  default: {} as any,
  local: true,
});

const format = defineModel('format', {
  default: formatKeys[0],
  local: true,
  required: true
});
</script>
<template>
  <div class="form-body">
    <div class="form-fieldset seperator">
      <div class="fieldset-heading">
        <h5>IP流参数</h5>
      </div>
      <div class="columns is-multiline">
        <div class="column is-4">
          <VField>
            <VLabel>视频流协议</VLabel>
            <VControl>
              <VInput
                v-model="mv.V_Protocol"
                readonly
              />
            </VControl>
          </VField>
        </div>
        <div class="column is-4">
          <VField>
            <VLabel>音频流协议</VLabel>
            <VControl>
              <VInput
                v-model="mv.A_Protocol"
                readonly
              />
            </VControl>
          </VField>
        </div>
        <div class="column is-4">
          <VField>
            <VLabel>主视频流组播地址（含端口）</VLabel>
            <VControl>
              <VInput
                v-model="mv.V_M_Address"
              />
            </VControl>
          </VField>
        </div>
        <div class="column is-4">
          <VField>
            <VLabel>主音频流组播地址（含端口）</VLabel>
            <VControl>
              <VInput
                v-model="mv.A_M_Address"
              />
            </VControl>
          </VField>
        </div>
        <Transition name="fade-slow">
          <div
            v-if="mv['2022-7']"
            class="column is-4"
          >
            <VField>
              <VLabel>备视频流组播地址（含端口）</VLabel>
              <VControl>
                <VInput
                  v-model="mv.V_B_Address"
                />
              </VControl>
            </VField>
          </div>
        </Transition>
        <Transition name="fade-slow">
          <div
            v-if="mv['2022-7']"
            class="column is-4"
          >
            <VField>
              <VLabel>备音频流组播地址（含端口）</VLabel>
              <VControl>
                <VInput
                  v-model="mv.A_B_Address"
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
                  v-for="f in formats"
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
    <div class="form-fieldset">
      <div class="fieldset-heading">
        <h5>视频编码参数</h5>
      </div>
      <div class="columns is-multiline">
        <div class="column is-6">
          <VField>
            <VLabel>编码格式</VLabel>
            <VControl>
              <VInput
                v-model="mv.V_DecFormat"
                readonly
              />
            </VControl>
          </VField>
        </div>
        <div class="column is-6">
          <VField>
            <VLabel>压缩比</VLabel>
            <VControl>
              <VInput
                v-model="mv.V_CompressionRatio"
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
                v-model="mv.A_Channels_Number"
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
                v-model="mv.A_Bits"
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
                v-model="mv.A_Frequency"
                readonly
              />
            </VControl>
          </VField>
        </div>
      </div>
    </div>
  </div>
</template>
