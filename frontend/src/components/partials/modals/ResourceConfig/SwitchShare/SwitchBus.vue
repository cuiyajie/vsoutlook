<script setup lang="ts">
import { useSwitchBus, useSwitchKeySignal } from '../Utilties/Composables';
import {
  type SwitchBus,
  def_switch_bus_me_params,
  def_switch_bus_level_params,
  def_switch_bus_key_params,
  type SwitchInputKeyParams,
  type SwitchInputVideoParams
} from './Consts'
import { merge } from 'lodash-es'

const mv = defineModel<SwitchBus>({
  default: {} as SwitchBus,
  local: true,
})

const props = defineProps<{
  level: number
  inputKeys: SwitchInputKeyParams[]
  inputVideos: SwitchInputVideoParams[]
}>()

const switchTypeCategory = inject('switch_type_category')

const [selectedBus, busSelected, busUnSelected] = useSwitchBus()
provide('switch_bus_select', { selectedBus, busSelected, busUnSelected })

const [selectedKeySignal, keySignalSelected, keySignalUnSelected] = useSwitchKeySignal()
provide('switch_key_signal_select', { selectedKeySignal, keySignalSelected, keySignalUnSelected })

const keyParams = computed({
  get: () => mv.value.key_bus,
  set: (value) => mv.value.key_bus = value
})

const meParams = computed({
  get: () => mv.value.me_bus,
  set: (value) => mv.value.me_bus = value
})

const levelParams = computed({
  get: () => mv.value.level_bus,
  set: (value) => mv.value.level_bus = value
})

watch(() => props.level, () => {
  levelParams.value = Array.from({ length: props.level }, (_, i) => merge(
    def_switch_bus_level_params(i, i),
    (levelParams.value[i] || {}),
  ))
}, { immediate: true })

function addKeyParam() {
  keyParams.value.push(def_switch_bus_key_params(keyParams.value.length))
}

function removeKeyParam(idx: number) {
  keyParams.value.splice(idx, 1)
}

function addMeParam() {
  meParams.value.push(def_switch_bus_me_params(meParams.value.length))
}

function removeMeParam(idx: number) {
  meParams.value.splice(idx, 1)
}

</script>
<template>
  <div class="form-outer has-mt-20">
    <div class="form-header is-sticky">
      <div class="form-header-inner">
        <div class="left">
          <h4>键总线设置 {{ keyParams.length > 0 ? `- ${keyParams.length} 路` : '' }}</h4>
        </div>
        <button
          class="button is-circle is-dark-outlined"
          @keydown.space.prevent="addKeyParam"
          @click.prevent="addKeyParam"
        >
          <span class="icon is-small">
            <i aria-hidden="true" class="iconify" data-icon="feather:plus" />
          </span>
        </button>
      </div>
    </div>
    <div class="form-body">
      <TransitionGroup name="list">
        <SwitchBusKeyParams
          v-for="(keyParam, kidx) in keyParams"
          :key="`input-key-${keyParam.index}`"
          v-model="keyParams[kidx]"
          :index="kidx"
          :is-last="kidx === keyParams.length - 1"
          :input-keys="inputKeys"
          :input-videos="inputVideos"
          :bus-levels="levelParams"
          @remove="removeKeyParam(kidx)"
        />
      </TransitionGroup>
      <div v-if="keyParams.length === 0" class="is-empty-list">暂时还没有配置键总线</div>
    </div>
  </div>
  <div v-if="switchTypeCategory !== 'bcswt'" class="form-outer">
    <div class="form-header">
      <div class="form-header-inner">
        <div class="left">
          <h4>ME总线设置</h4>
        </div>
        <button
          class="button is-circle is-dark-outlined"
          @keydown.space.prevent="addMeParam"
          @click.prevent="addMeParam"
        >
          <span class="icon is-small">
            <i aria-hidden="true" class="iconify" data-icon="feather:plus" />
          </span>
        </button>
      </div>
    </div>
    <div class="form-body">
      <TransitionGroup name="list">
        <SwitchBusMeParams
          v-for="(meParam, meidx) in meParams"
          :key="`input-me-${meParam.index}`"
          v-model="meParams[meidx]"
          :index="meidx"
          :is-last="meidx === meParams.length - 1"
          @remove="removeMeParam(meidx)"
        />
      </TransitionGroup>
      <div v-if="meParams.length === 0" class="is-empty-list">暂时还没有配置ME总线</div>
    </div>
  </div>
  <div class="form-outer has-mb-20">
    <div class="form-header">
      <div class="form-header-inner">
        <div class="left">
          <h4>切换台级输出总线设置 {{ levelParams.length > 0 ? `- ${levelParams.length} 级` : '' }}</h4>
        </div>
      </div>
    </div>
    <div class="form-body">
      <TransitionGroup name="list">
        <SwitchBusLevelParams
          v-for="(levelParam, levelidx) in levelParams"
          :key="`input-level-${levelParam.level}`"
          v-model="levelParams[levelidx]"
          :index="levelidx"
          :is-last="levelidx === levelParams.length - 1"
          :level="level"
          :input-keys="inputKeys"
          :input-videos="inputVideos"
          :bus-keys="keyParams"
          :bus-mes="meParams"
          :bus-levels="levelParams"
        />
      </TransitionGroup>
      <div v-if="levelParams.length === 0" class="is-empty-list">暂时还没有配置切换台级输出总线</div>
    </div>
  </div>
</template>
