<script lang="ts" setup>
import { useSwitchInputSignal } from '../Utilties/Composables';
import { type IndexedNicDetail } from '../Utilties/Consts_V1';
import { type SwitchBusLevelParams, bus_signal_types, bus_signal_types_map, type SwitchInputKeyParams, type SwitchInputVideoParams, def_switch_bus_level_input_params, type SwitchBusKeyParams, type SwitchBusMeParams, def_switch_bus_level_sub_params } from './Consts'

const mv = defineModel<SwitchBusLevelParams>({
  default: {} as SwitchBusLevelParams,
  local: true,
})

const props = defineProps<{
  index: number
  isLast: boolean
  level: number
  inputKeys: SwitchInputKeyParams[]
  inputVideos: SwitchInputVideoParams[]
  busKeys: SwitchBusKeyParams[]
  busMes: SwitchBusMeParams[]
  busLevels: SwitchBusLevelParams[]
}>()

const nics = inject<IndexedNicDetail[]>('switch_nics', [])

const busSignalTypes = computed(() => {
  return bus_signal_types.filter(
    bst => bst.value === 'video' ||
    (bst.value === 'bus_output' && props.level === 2 && mv.value.level === 1)
  )
})

watch(() => mv.value.bus_input.bus_input_number, (nv) => {
  const len = mv.value.bus_input.input_list.length
  if (nv < len) {
    mv.value.bus_input.input_list = mv.value.bus_input.input_list.slice(0, nv)
  } else if (nv > len) {
    mv.value.bus_input.input_list = mv.value.bus_input.input_list.concat(
      Array.from({ length: nv - len }, (_, i) => def_switch_bus_level_input_params(i + len))
    )
  }
}, { immediate: true })

watch(() => mv.value.bus_input.nic_index, () => {
  const core = `rx#${mv.value.bus_input.nic_index}`
  mv.value.pgm_bus.v_core = core
  mv.value.pgm_bus.a_core = core
  mv.value.pvw_bus.v_core = core
  mv.value.pvw_bus.a_core = core
}, { immediate: true })

const [selectedSignal, signalSelected, signalUnSelected] = useSwitchInputSignal()
provide('switch_input_signal_select', { selectedSignal, signalSelected, signalUnSelected })

function addPgmSubBus() {
  mv.value.pgm_bus.sub_bus.push(def_switch_bus_level_sub_params(mv.value.pgm_bus.sub_bus.length))
}

function removePgmSubBus() {
  mv.value.pgm_bus.sub_bus.pop()
}

function addPvwSubBus() {
  mv.value.pvw_bus.sub_bus.push(def_switch_bus_level_sub_params(mv.value.pvw_bus.sub_bus.length))
}

function removePvwSubBus() {
  mv.value.pvw_bus.sub_bus.pop()
}

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
      <h4>第 {{ index + 1 }} 级输出总线&nbsp;&nbsp;索引: {{ index }}</h4>
      <div class="collapse-icon">
        <VIcon icon="feather:chevron-down" />
      </div>
    </div>
    <expand-transition>
      <div v-show="opened">
        <div class="form-fieldset-nested-3 seperator" style="margin-bottom: 30px;">
          <div class="fieldset-heading">
            <h4>输入信源</h4>
          </div>
          <div class="columns is-multiline">
            <div class="column is-6">
              <VField>
                <VLabel>输入信源数量</VLabel>
                <VControl>
                  <VInputNumber v-model="mv.bus_input.bus_input_number" centered :min="0" :step="1" />
                </VControl>
              </VField>
            </div>
            <div class="column is-6">
              <VField>
                <VLabel>使用的网卡序号</VLabel>
                <VControl>
                  <NicDetailSelect v-model="mv.bus_input.nic_index" :nics="nics" />
                </VControl>
              </VField>
            </div>
          </div>
          <transition-group name="list">
            <div
              v-for="(bipt, bidx) in mv.bus_input.input_list"
              :key="bidx"
              :class="{
                'form-fieldset-nested-4': bidx === 0,
                'form-fieldset-nested-5': bidx !== 0
              }"
            >
              <div class="fieldset-heading is-nested">
                <h5>第{{ bidx + 1 }}路输入信号&nbsp;&nbsp;索引: {{ bidx }}</h5>
              </div>
              <div class="columns is-multiline" style="padding-left: 16px;">
                <div class="column is-6">
                  <VField>
                    <VLabel>切换台输入序号</VLabel>
                    <VControl>
                      <VInput v-model="bipt.input_index" readonly centered />
                    </VControl>
                  </VField>
                </div>
                <div class="column is-6">
                  <VField>
                    <VLabel>信号类型</VLabel>
                    <VControl>
                      <VSelect v-model="bipt.signal_type" class="is-rounded">
                        <VOption
                          v-for="bst in busSignalTypes"
                          :key="bst.value"
                          :value="bst.value"
                        >
                          {{ bst.label }}
                        </VOption>
                      </VSelect>
                    </VControl>
                  </VField>
                </div>
                <div class="column is-6">
                  <VField>
                    <VLabel>输入信号源名称</VLabel>
                    <VControl>
                      <SwitchInputSignalSelect
                        v-model:id="bipt.signal_id"
                        v-model:name="bipt.signal_name"
                        :title="bus_signal_types_map.get(bipt.signal_type)?.label || '输入信号源名称'"
                        :type="bipt.signal_type"
                        :input-keys="inputKeys"
                        :input-videos="inputVideos"
                        :bus-levels="busLevels"
                        :level-index="mv.level"
                        :emit-signal="true"
                        :unique-key-signal="true"
                        :unique-signal="true"
                      />
                    </VControl>
                  </VField>
                </div>
                <div class="column is-6">
                  <VField>
                    <VLabel>输入信号源ID</VLabel>
                    <VControl>
                      <VLabel class="is-static">{{ bipt.signal_id || '未配置' }}</VLabel>
                    </VControl>
                  </VField>
                </div>
              </div>
            </div>
          </transition-group>
        </div>
        <div class="form-fieldset-nested-3 seperator" style="margin-bottom: 30px;">
          <div class="fieldset-heading flex-between">
            <h4>PGM总线设置</h4>
            <VButton
              class="is-rounded"
              color="primary"
              raised
              @click="addPgmSubBus"
            >
              添加子总线
            </VButton>
          </div>
          <div class="form-fieldset-nested-4" style="padding-left: 16px;">
            <transition-group name="list">
              <SwitchSubBus
                v-for="(sub, subIdx) in mv.pgm_bus.sub_bus"
                :key="subIdx"
                v-model="mv.pgm_bus.sub_bus[subIdx]"
                :index="subIdx"
                :is-last="subIdx === mv.pgm_bus.sub_bus.length - 1"
                :bus-keys="busKeys"
                :bus-mes="busMes"
                type="pgm"
                @remove="removePgmSubBus"
              ></SwitchSubBus>
            </transition-group>
          </div>
          <div class="form-fieldset-nested-5" style="padding-left: 16px;">
            <div class="fieldset-heading">
              <h5>PGM总线clean输出信号</h5>
            </div>
            <div class="columns is-multiline">
              <div class="column is-1">
                <VField>
                  <VLabel>索引</VLabel>
                  <VControl>
                    <VLabel class="is-static">{{ mv.pgm_bus.out_signal[0].index }}</VLabel>
                  </VControl>
                </VField>
              </div>
              <div class="column is-3">
                <VField>
                  <VLabel>输出信号类型</VLabel>
                  <VControl>
                    <VLabel class="is-static">{{ mv.pgm_bus.out_signal[0].signal_type }}</VLabel>
                  </VControl>
                </VField>
              </div>
              <div class="column is-4">
                <VField>
                  <VLabel>输出信号名称</VLabel>
                  <VControl>
                    <VLabel class="is-static">{{ mv.pgm_bus.out_signal[0].signal_name }}</VLabel>
                  </VControl>
                </VField>
              </div>
              <div class="column is-4">
                <VField>
                  <VLabel>输出信号ID</VLabel>
                  <VControl>
                    <VLabel class="is-static">{{ mv.pgm_bus.out_signal[0].signal_id }}</VLabel>
                  </VControl>
                </VField>
              </div>
            </div>
          </div>
          <div class="form-fieldset-nested-5" style="padding-left: 16px;">
            <div class="fieldset-heading">
              <h5>PGM总线末级输出信号</h5>
            </div>
            <div class="columns is-multiline">
              <div class="column is-1">
                <VField>
                  <VLabel>索引</VLabel>
                  <VControl>
                    <VLabel class="is-static">{{ mv.pgm_bus.out_signal[1].index }}</VLabel>
                  </VControl>
                </VField>
              </div>
              <div class="column is-3">
                <VField>
                  <VLabel>输出信号类型</VLabel>
                  <VControl>
                    <VLabel class="is-static">{{ mv.pgm_bus.out_signal[1].signal_type }}</VLabel>
                  </VControl>
                </VField>
              </div>
              <div class="column is-4">
                <VField>
                  <VLabel>输出信号名称</VLabel>
                  <VControl>
                    <VLabel class="is-static">{{ mv.pgm_bus.out_signal[1].signal_name }}</VLabel>
                  </VControl>
                </VField>
              </div>
              <div class="column is-4">
                <VField>
                  <VLabel>输出信号ID</VLabel>
                  <VControl>
                    <VLabel class="is-static">{{ mv.pgm_bus.out_signal[1].signal_id }}</VLabel>
                  </VControl>
                </VField>
              </div>
            </div>
          </div>
        </div>
        <div class="form-fieldset-nested-4">
          <div class="fieldset-heading flex-between">
            <h4>PVW总线设置</h4>
            <VButton
              class="is-rounded"
              color="primary"
              raised
              @click="addPvwSubBus"
            >
              添加子总线
            </VButton>
          </div>
          <div class="form-fieldset-nested-4" style="padding-left: 16px;">
            <transition-group name="list">
              <SwitchSubBus
                v-for="(sub, subIdx) in mv.pvw_bus.sub_bus"
                :key="subIdx"
                v-model="mv.pvw_bus.sub_bus[subIdx]"
                :index="subIdx"
                :is-last="subIdx === mv.pvw_bus.sub_bus.length - 1"
                :bus-keys="busKeys"
                :bus-mes="busMes"
                type="pvw"
                @remove="removePvwSubBus"
              ></SwitchSubBus>
            </transition-group>
          </div>
          <div class="form-fieldset-nested-5" style="padding-left: 16px;">
            <div class="fieldset-heading">
              <h5>PVW总线末级输出信号</h5>
            </div>
            <div class="columns is-multiline">
              <div class="column is-1">
                <VField>
                  <VLabel>索引</VLabel>
                  <VControl>
                    <VLabel class="is-static">{{ mv.pvw_bus.out_signal[0].index }}</VLabel>
                  </VControl>
                </VField>
              </div>
              <div class="column is-3">
                <VField>
                  <VLabel>输出信号类型</VLabel>
                  <VControl>
                    <VLabel class="is-static">{{ mv.pvw_bus.out_signal[0].signal_type }}</VLabel>
                  </VControl>
                </VField>
              </div>
              <div class="column is-4">
                <VField>
                  <VLabel>输出信号名称</VLabel>
                  <VControl>
                    <VLabel class="is-static">{{ mv.pvw_bus.out_signal[0].signal_name }}</VLabel>
                  </VControl>
                </VField>
              </div>
              <div class="column is-4">
                <VField>
                  <VLabel>输出信号ID</VLabel>
                  <VControl>
                    <VLabel class="is-static">{{ mv.pvw_bus.out_signal[0].signal_id }}</VLabel>
                  </VControl>
                </VField>
              </div>
            </div>
          </div>
        </div>
      </div>
    </expand-transition>
  </div>
</template>
