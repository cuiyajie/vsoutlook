<script lang="ts" setup>
import { type SwitchBusMeParams } from './Consts'

const mv = defineModel<SwitchBusMeParams>({
  default: {} as SwitchBusMeParams,
  local: true,
})

defineProps<{
  index: number
  isLast: boolean
}>()

const emit = defineEmits<{
  (e: 'remove'): void
}>()

const opened = ref(false)
</script>
<template>
  <div
    class="form-fieldset collapse-form"
    :class="!isLast && 'seperator'"
    :open="opened || undefined"
  >
    <div
      class="fieldset-heading collapse-header"
      tabindex="0"
      role="button"
      @keydown.space.prevent="opened = !opened"
      @click.prevent="opened = !opened"
    >
      <h4>第 {{ index + 1 }} 路ME总线&nbsp;&nbsp;索引: {{ index }}</h4>
      <div class="collapse-icons">
        <div
          class="collapse-icon is-close-hidden"
          role="button"
          tabindex="0"
          @click.prevent.stop="emit('remove')"
          @keydown.space.prevent.stop="emit('remove')"
        >
          <VIcon icon="feather:x" />
        </div>
        <div class="collapse-icon">
          <VIcon icon="feather:chevron-down" />
        </div>
      </div>
    </div>
    <expand-transition>
      <div v-show="opened">
        <div class="form-fieldset-nested-3 seperator">
          <div class="columns is-multiline">
            <div class="column is-4">
              <VField>
                <VLabel>总线ID</VLabel>
                <VControl>
                  <VLabel class="is-static">{{ mv.bus_id }}</VLabel>
                </VControl>
              </VField>
            </div>
            <div class="column is-4">
              <VField>
                <VLabel>总线名称</VLabel>
                <VControl>
                  <VInput v-model="mv.bus_name" />
                </VControl>
              </VField>
            </div>
            <div class="column is-4">
              <VField>
                <VLabel>auto模式时长（帧）</VLabel>
                <VControl>
                  <VInputNumber v-model="mv.auto_duration" centered :min="0" :step="1" />
                </VControl>
              </VField>
            </div>
          </div>
        </div>
        <div class="form-fieldset-nested-5">
          <div class="fieldset-heading">
            <h4>总线输出</h4>
          </div>
          <div class="columns is-multiline">
            <div class="column is-6">
              <VField>
                <VLabel>信号名称</VLabel>
                <VControl>
                  <VInput v-model="mv.out_signal.signal_name" />
                </VControl>
              </VField>
            </div>
            <div class="column is-6">
              <VField>
                <VLabel>信号ID</VLabel>
                <VControl>
                  <VLabel class="is-static">{{ mv.out_signal.signal_id }}</VLabel>
                </VControl>
              </VField>
            </div>
          </div>
        </div>
      </div>
    </expand-transition>
  </div>
</template>
