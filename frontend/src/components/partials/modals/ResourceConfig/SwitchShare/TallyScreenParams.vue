<script lang="ts" setup>
import { type TallyScreenParams, def_tally_signal } from './Consts'

const mv = defineModel<TallyScreenParams>({
  default: {} as TallyScreenParams,
  local: true,
})

defineProps<{
  index: number
  isLast: boolean
}>()

const emit = defineEmits<{
  (e: 'tally-checked'): void
}>()

watch(() => mv.value.signal_number, (nv) => {
  const len = mv.value.signal_config.length
  if (nv < len) {
    mv.value.signal_config = mv.value.signal_config.slice(0, nv)
  } else if (nv > len) {
    mv.value.signal_config = mv.value.signal_config.concat(
      Array.from({ length: nv - len }, (_, i) => {
        return def_tally_signal(len + i)
      })
    )
  }
}, { immediate: true })

</script>
<template>
  <div class="form-fieldset-nested-5 collapse-form">
    <div class="fieldset-heading collapse-header is-nested">
      <h4>多画面 {{ index + 1 }} 设置</h4>
    </div>
    <div :class="isLast ? 'form-fieldset-nested-4' : 'form-fieldset-nested-3 seperator'">
      <div class="columns is-multiline">
        <div class="column is-4">
          <VField>
            <VLabel>屏幕ID</VLabel>
            <VControl>
              <VInputNumber v-model="mv.screen_id" centered :min="0" :step="1" />
            </VControl>
          </VField>
        </div>
        <div class="column is-4">
          <VField>
            <VLabel>通讯地址</VLabel>
            <VControl>
              <VInput v-model="mv.address" />
            </VControl>
          </VField>
        </div>
        <div class="column is-4">
          <VField>
            <VLabel>当前多画面信源数量</VLabel>
            <VControl>
              <VInputNumber v-model="mv.signal_number" centered :min="0" :step="1" />
            </VControl>
          </VField>
        </div>
        <div class="column is-4">
          <VField class="field-check" label="PGM tally索引">
            <VControl raw subcontrol class="check-control">
              <VCheckbox
                v-model="mv.pgm_checked"
                color="primary"
                @change="event => mv.pvw_checked = !Boolean((event.target as HTMLInputElement).value)"
              />
            </VControl>
            <VControl>
              <VInputNumber v-model="mv.pgm_tally_index" centered :min="0" :step="1" />
            </VControl>
          </VField>
        </div>
        <div class="column is-4">
          <VField class="field-check" label="PVW tally索引">
            <VControl raw subcontrol class="check-control">
              <VCheckbox
                v-model="mv.pvw_checked"
                color="primary"
                @change="event => mv.pgm_checked = !Boolean((event.target as HTMLInputElement).value)"
              />
            </VControl>
            <VControl>
              <VInputNumber v-model="mv.pvw_tally_index" centered :min="0" :step="1" />
            </VControl>
          </VField>
        </div>
      </div>
      <expand-transition>
        <div v-if="mv.signal_number > 0" class="form-fieldset-nested-5">
          <div class="fieldset-heading" style="margin-bottom: 0 !important;">
            <h5>信源配置</h5>
          </div>
          <div class="form-body signal-container">
            <transition-group name="list">
              <TallySignalParams
                v-for="(signalParam, signalidx) in mv.signal_config"
                :key="`input-signal-${signalParam.index}`"
                v-model="mv.signal_config[signalidx]"
                :index="signalidx"
                :is-last="signalidx === mv.signal_config.length - 1"
              />
            </transition-group>
          </div>
        </div>
      </expand-transition>
    </div>
  </div>
</template>
<style lang="scss" scoped>
.signal-container {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  column-gap: 30px;
}
</style>

