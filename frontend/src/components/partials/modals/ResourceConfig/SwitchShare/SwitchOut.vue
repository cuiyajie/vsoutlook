<script setup lang="ts">
import {
  def_switch_out_params,
  type SwitchInputVideoParams,
  type SwitchBusLevelParams,
  type SwitchOut,
  type SwitchOutParams
} from './Consts'

const mv = defineModel<SwitchOut>({
  default: {} as SwitchOut,
  local: true,
})

defineProps<{
  level: number
  inputVideos: SwitchInputVideoParams[]
  busLevels: SwitchBusLevelParams[]
}>()

const audioMode = inject<Ref<number>>('switch_audio_mode')
const outParms = computed({
  get: () => mv.value.out_params,
  set: (v) => mv.value.out_params = v
})

watch(() => mv.value.out_number, (nv) => {
  const len = outParms.value.length
  if (nv < len) {
    outParms.value = outParms.value.slice(0, nv)
  } else if (nv > len) {
    outParms.value = outParms.value.concat(
      Array.from({ length: nv - len }, (_, i) => def_switch_out_params(len + i))
    )
  }
}, { immediate: true })

const audioOpen = ref(false)
const mvOpen = ref(false)

</script>
<template>
  <div class="form-outer has-mt-20">
    <div class="form-body">
      <div class="form-fieldset-nested-1" :class="outParms.length > 0 && 'seperator'">
        <div class="columns is-multiline">
          <div class="column is-6">
            <VField>
              <VLabel>输出信号数量（最多 {{ level * 3 }} 路）</VLabel>
              <VControl>
                <VInputNumber v-model="mv.out_number" centered :min="0" :step="1" :max="level * 3" />
              </VControl>
            </VField>
          </div>
        </div>
      </div>
      <TransitionGroup name="list">
        <SwitchOutParams
          v-for="(outParam, oidx) in outParms"
          :key="`out-${outParam.index}`"
          v-model="outParms[oidx]"
          :index="oidx"
          :is-last="oidx === outParms.length - 1"
          :input-videos="inputVideos"
          :bus-levels="busLevels"
        />
      </TransitionGroup>
    </div>
  </div>
  <div
    v-if="audioMode === 2"
    class="form-outer collapse-form"
    :open="audioOpen || undefined"
  >
    <div
      class="form-header collapse-header"
      :style="{'border-bottom-width': audioOpen ? '1px' : '0'}"
      tabindex="0"
      role="button"
      @keydown.space.prevent="audioOpen = !audioOpen"
      @click.prevent="audioOpen = !audioOpen"
    >
      <div class="form-header-inner">
        <div class="left">
          <h4>音频输出设置</h4>
        </div>
      </div>
      <div class="collapse-icon">
        <VIcon icon="feather:chevron-down" />
      </div>
    </div>
    <expand-transition>
      <div v-if="audioOpen" class="form-body">
        <SwitchOutAudioParams v-model="mv.audio_output_params" :input-videos="inputVideos" :out-params="outParms" />
      </div>
    </expand-transition>
  </div>
  <div
    class="form-outer has-mb-20 collapse-form"
    :open="mvOpen || undefined"
  >
    <div
      class="form-header collapse-header"
      :style="{'border-bottom-width': mvOpen ? '1px' : '0'}"
      tabindex="0"
      role="button"
      @keydown.space.prevent="mvOpen = !mvOpen"
      @click.prevent="mvOpen = !mvOpen"
    >
      <div class="form-header-inner">
        <div class="left">
          <h4>多画分输出设置</h4>
        </div>
      </div>
      <div class="collapse-icon">
        <VIcon icon="feather:chevron-down" />
      </div>
    </div>
    <expand-transition>
      <div v-if="mvOpen" class="form-body">
        <SwitchOutMvParams v-model="mv.mv_output_params" />
      </div>
    </expand-transition>
  </div>
</template>
