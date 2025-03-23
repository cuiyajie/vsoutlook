<!-- eslint-disable vue/prop-name-casing -->
<script setup lang="ts">
import { type NMosConfigType } from './Consts'

const mv = defineModel<NMosConfigType>({
  default: {} as NMosConfigType,
  local: true,
})
</script>
<template>
  <div class="form-fieldset">
    <div class="fieldset-heading">
      <h4>nmos参数配置</h4>
    </div>

    <div class="columns is-multiline">
      <div class="column is-6">
        <VField>
          <VLabel>是否启用nmos</VLabel>
          <VControl>
            <VSwitchBlock v-model="mv.nmos_enable" color="primary" />
          </VControl>
        </VField>
      </div>
      <div class="column is-6">
        <VField>
          <VLabel>对外端口</VLabel>
          <VControl>
            <VInputNumber v-model="mv.http_port" centered :min="0" :step="1" />
          </VControl>
        </VField>
      </div>
      <div class="column is-6">
        <VField>
          <VLabel>日志级别</VLabel>
          <VControl>
            <VInputNumber v-model="mv.log_level" centered :min="0" :step="1" />
          </VControl>
        </VField>
      </div>
      <div class="column is-6">
        <VField>
          <VLabel>日志路径</VLabel>
          <VControl>
            <VInput v-model="mv.log_path" />
          </VControl>
        </VField>
      </div>
      <div class="column is-6">
        <VField>
          <VLabel>域</VLabel>
          <VControl>
            <VInput v-model="mv.domain" />
          </VControl>
        </VField>
      </div>
    </div>
  </div>
</template>
