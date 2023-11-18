<script setup lang="ts">
import { useFormat, useProtocolDC } from './Utils';
import { type Config4, formats, vProtocols } from './Consts'

const mv = defineModel<Config4>({
  default: {} as any,
  local: true,
});

const V_Format = useFormat(mv);
useProtocolDC(mv);
</script>
<template>
  <div class="columns is-multiline">
    <div class="column is-12">
      <VField horizontal>
        <VLabel>启用2022-7备份</VLabel>
        <VControl>
          <VSwitchBlock
            v-model="mv['2022-7']"
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
            v-model="mv.V_Protocol"
            class="is-rounded"
          >
            <VOption
              v-for="vp in vProtocols"
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
            v-model="V_Format"
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
    <div class="column is-6">
      <VField>
        <VLabel>视频编码格式</VLabel>
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
        <VLabel>视频压缩比</VLabel>
        <VControl>
          <VInput
            v-model="mv.V_CompressionRatio"
            readonly
          />
        </VControl>
      </VField>
    </div>
  </div>
</template>
