<script setup lang="ts">
import { getFormat, useFormat, useProtocolDC } from './Utils';
import { formats, v_protocols, type SwitchInputParamsType } from './Consts';

const mv = defineModel<SwitchInputParamsType>({
  default: {} as SwitchInputParamsType,
  local: true,
});

const format = useFormat(mv.value, getFormat(mv.value));
useProtocolDC(mv.value);

</script>
<template>
  <div class="form-fieldset">
    <div class="fieldset-heading">
      <h4>全局参数</h4>
    </div>
    <div class="columns is-multiline">
      <div class="column is-3">
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
      <div class="column is-3">
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
      <div v-if="mv.videoformat.v_compression_format" class="column is-6">
        <VField>
          <VLabel>视频编码格式</VLabel>
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
          <VLabel>视频压缩比</VLabel>
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
</template>
