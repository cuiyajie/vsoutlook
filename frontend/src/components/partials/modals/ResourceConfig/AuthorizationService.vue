<!-- eslint-disable vue/prop-name-casing -->
<script setup lang="ts">
import { type AuthServiceType, auth_service } from './Consts';

const mv = defineModel<AuthServiceType[]>({
  default: [],
  local: true,
});

function add() {
  mv.value.push({ ...auth_service });
}

function remove(idx: number) {
  mv.value.splice(idx, 1);
}
</script>
<template>
  <div class="form-fieldset">
    <div class="fieldset-heading flex">
      <h4>授权服务配置</h4>
      <button
        class="button is-circle is-dark-outlined"
        @keydown.space.prevent="add"
        @click.prevent="add"
      >
        <span class="icon is-small">
          <i aria-hidden="true" class="iconify" data-icon="feather:plus" />
        </span>
      </button>
    </div>

    <div v-for="(auth, idx) in mv" :key="idx" class="columns is-multiline">
      <div class="column is-1">
        <VField class="index-field">
          <VLabel>索引</VLabel>
          <VControl>
            <VLabel>{{ idx }}</VLabel>
          </VControl>
        </VField>
      </div>
      <div class="column is-6">
        <VField>
          <VLabel>服务IP地址</VLabel>
          <VControl>
            <VInput v-model="auth.ip" />
          </VControl>
        </VField>
      </div>
      <div class="column is-4">
        <VField>
          <VLabel>服务端口</VLabel>
          <VControl>
            <VInputNumber
              v-model="auth.port"
              centered
              :min="0"
              :step="1"
            />
          </VControl>
        </VField>
      </div>
      <div class="column is-1">
        <VField>
          <VLabel>&nbsp;</VLabel>
          <VControl>
            <VIconButton color="warning" light raised circle icon="feather:x" @click="remove" />
          </VControl>
        </VField>
      </div>
    </div>
  </div>
</template>
