<!-- eslint-disable vue/prop-name-casing -->
<script setup lang="ts">
import { type SSMAddressType, ssm_address } from './Consts'

const mv = defineModel<SSMAddressType[]>({
  default: [],
  local: true,
})

function add() {
  mv.value.push({ ...ssm_address })
}

function remove(idx: number) {
  mv.value.splice(idx, 1)
}
</script>
<template>
  <div class="form-fieldset">
    <div class="fieldset-heading flex">
      <h4>SSM组播源地址范围</h4>
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

    <div v-for="(ssm, idx) in mv" :key="idx" class="columns is-multiline">
      <div class="column is-1">
        <VField class="index-field">
          <VLabel>序号</VLabel>
          <VControl>
            <VLabel>{{ idx }}</VLabel>
          </VControl>
        </VField>
      </div>
      <div class="column is-5">
        <VField>
          <VLabel>起始地址</VLabel>
          <VControl>
            <VInput v-model="ssm.from" />
          </VControl>
        </VField>
      </div>
      <div class="column is-5">
        <VField>
          <VLabel>结束地址</VLabel>
          <VControl>
            <VInput v-model="ssm.to" />
          </VControl>
        </VField>
      </div>
      <div class="column is-1">
        <VField>
          <VLabel>&nbsp;</VLabel>
          <VControl>
            <VIconButton
              color="warning"
              light
              raised
              circle
              icon="feather:x"
              @click="remove(idx)"
            />
          </VControl>
        </VField>
      </div>
    </div>
  </div>
</template>
