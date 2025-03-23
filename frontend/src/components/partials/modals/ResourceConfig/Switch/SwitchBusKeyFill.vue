<script setup lang="ts">
import { type SwitchBusKeyfillParamsType } from '../Utilties/Consts'

const mv = defineModel<SwitchBusKeyfillParamsType>({
  default: {} as SwitchBusKeyfillParamsType,
  local: true,
})

defineProps<{
  name: '主' | '备'
  index: number
  isLast: boolean
}>()

const keyfill_opened = ref(false)
</script>
<template>
  <div
    class="form-fieldset collapse-form"
    :class="!isLast && 'seperator'"
    :open="keyfill_opened || undefined"
  >
    <div
      class="fieldset-heading collapse-header"
      tabindex="0"
      role="button"
      @keydown.space.prevent="keyfill_opened = !keyfill_opened"
      @click.prevent="keyfill_opened = !keyfill_opened"
    >
      <h4>第{{ index }}路{{ name }}键输入母线</h4>
      <div class="collapse-icon">
        <VIcon icon="feather:chevron-down" />
      </div>
    </div>
    <Transition name="fade-show">
      <div v-show="keyfill_opened" class="form-fieldset">
        <div class="columns is-multiline">
          <div class="column is-4">
            <VField>
              <VLabel>{{ name }}key母线组播源IP（含端口）</VLabel>
              <VControl>
                <VInput v-model="mv.key_src_address" />
              </VControl>
            </VField>
          </div>
          <div class="column is-4">
            <VField>
              <VLabel>{{ name }}key母线组播目标IP（含端口）</VLabel>
              <VControl>
                <VInput v-model="mv.key_dst_address" />
              </VControl>
            </VField>
          </div>
          <div class="column is-4">
            <VField>
              <VLabel>{{ name }}key母线使用的前置SDN输出端口</VLabel>
              <VControl>
                <VInputNumber v-model="mv.key_p4_port" centered :min="0" :step="1" />
              </VControl>
            </VField>
          </div>
          <div class="column is-4">
            <VField>
              <VLabel>{{ name }}fill母线组播源IP（含端口）</VLabel>
              <VControl>
                <VInput v-model="mv.fill_src_address" />
              </VControl>
            </VField>
          </div>
          <div class="column is-4">
            <VField>
              <VLabel>{{ name }}fill母线组播目标IP（含端口）</VLabel>
              <VControl>
                <VInput v-model="mv.fill_dst_address" />
              </VControl>
            </VField>
          </div>
          <div class="column is-4">
            <VField>
              <VLabel>{{ name }}fill母线使用的前置SDN输出端口</VLabel>
              <VControl>
                <VInputNumber v-model="mv.fill_p4_port" centered :min="0" :step="1" />
              </VControl>
            </VField>
          </div>
        </div>
      </div>
    </Transition>
  </div>
</template>
<style lang="scss">
@import '@src/scss/abstracts/all';
@import '@src/scss/components/forms-outer';
</style>
