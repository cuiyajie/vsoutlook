<script lang="ts" setup>
import { type IndexedNicDetail } from '../Utilties/Consts_V1';
import { type SwitchBusKeyParams, bus_signal_types, bus_signal_types_map, bus_key_types, type SwitchInputKeyParams, type SwitchInputVideoParams, type SwitchBusLevelParams } from './Consts'

const mv = defineModel<SwitchBusKeyParams>({
  default: {} as SwitchBusKeyParams,
  local: true,
})

defineProps<{
  index: number
  isLast: boolean
  inputKeys: SwitchInputKeyParams[]
  inputVideos: SwitchInputVideoParams[]
  busLevels: SwitchBusLevelParams[]
}>()

const usedSignalType = inject<Ref<number>>('switch_used_signal_type', ref(0))
const nics = inject<IndexedNicDetail[]>('switch_nics', [])

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
      <h4>第 {{ index + 1 }} 路键总线&nbsp;&nbsp;索引: {{ index }}</h4>
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
            <div class="column is-6">
              <VField>
                <VLabel>总线ID</VLabel>
                <VControl>
                  <VLabel class="is-static">{{ mv.bus_id }}</VLabel>
                </VControl>
              </VField>
            </div>
            <div class="column is-6">
              <VField>
                <VLabel>总线名称</VLabel>
                <VControl>
                  <VInput v-model="mv.bus_name" />
                </VControl>
              </VField>
            </div>
          </div>
        </div>
        <div class="form-fieldset-nested-1 seperator">
          <div class="fieldset-heading">
            <h4>总线输入源</h4>
          </div>
          <div class="columns is-multiline">
            <div v-if="usedSignalType !== 1" class="column is-6">
              <VField>
                <VLabel>使用的网卡序号</VLabel>
                <VControl>
                  <NicDetailSelect v-model="mv.bus_input.nic_index" :nics="nics" />
                </VControl>
              </VField>
            </div>
            <div class="column" :class="usedSignalType === 1 ? 'is-4' : 'is-6'">
              <VField>
                <VLabel>信号类型</VLabel>
                <VControl>
                  <VSelect v-model="mv.bus_input.signal_type" class="is-rounded">
                    <VOption v-for="bst in bus_signal_types" :key="bst.value" :value="bst.value">{{ bst.label }}</VOption>
                  </VSelect>
                </VControl>
              </VField>
            </div>
            <div class="column" :class="usedSignalType === 1 ? 'is-4' : 'is-6'">
              <VField>
                <VLabel>信号名称</VLabel>
                <VControl>
                  <SwitchInputSignalSelect
                    v-model:id="mv.bus_input.signal_id"
                    v-model:name="mv.bus_input.signal_name"
                    :title="bus_signal_types_map.get(mv.bus_input.signal_type)?.label || '信号名称'"
                    :type="mv.bus_input.signal_type"
                    :input-keys="inputKeys"
                    :input-videos="inputVideos"
                    :bus-levels="busLevels"
                    :level-index="-1"
                  />
                </VControl>
              </VField>
            </div>
            <div class="column" :class="usedSignalType === 1 ? 'is-4' : 'is-6'">
              <VField>
                <VLabel>信号ID</VLabel>
                <VControl>
                  <VLabel class="is-static">{{ mv.bus_input.signal_id || '未配置' }}</VLabel>
                </VControl>
              </VField>
            </div>
          </div>
        </div>
        <div class="form-fieldset-nested-1 seperator">
          <div class="fieldset-heading">
            <h4>键参数</h4>
          </div>
          <div class="columns is-multiline">
            <div class="column is-6">
              <VField>
                <VLabel>键类型</VLabel>
                <VControl>
                  <VSelect v-model="mv.key_params.keytype" class="is-rounded">
                    <VOption v-for="bkt in bus_key_types" :key="bkt.value" :value="bkt.value">{{ bkt.label }}</VOption>
                  </VSelect>
                </VControl>
              </VField>
            </div>
            <expand-transition>
              <div v-if="mv.key_params.keytype !== 'luma_key'" class="column is-6">
                <VField>
                  <VLabel>显示区域</VLabel>
                  <VControl>
                    <VLabel class="is-static">请依次设置顶点坐标、宽度和高度</VLabel>
                  </VControl>
                </VField>
              </div>
            </expand-transition>
            <expand-transition>
              <div v-if="mv.key_params.keytype !== 'luma_key'" class="column is-3">
                <VField>
                  <VLabel>顶点X坐标</VLabel>
                  <VControl>
                    <VInputNumber v-model="mv.key_params.top_x" centered :step="1" />
                  </VControl>
                </VField>
              </div>
            </expand-transition>
            <expand-transition>
              <div v-if="mv.key_params.keytype !== 'luma_key'" class="column is-3">
                <VField>
                  <VLabel>顶点Y坐标</VLabel>
                  <VControl>
                    <VInputNumber v-model="mv.key_params.top_y" centered :step="1" />
                  </VControl>
                </VField>
              </div>
            </expand-transition>
            <expand-transition>
              <div v-if="mv.key_params.keytype !== 'luma_key'" class="column is-3">
                <VField>
                  <VLabel>宽度</VLabel>
                  <VControl>
                    <VInputNumber v-model="mv.key_params.w" centered :min="0" :step="1" />
                  </VControl>
                </VField>
              </div>
            </expand-transition>
            <expand-transition>
              <div v-if="mv.key_params.keytype !== 'luma_key'" class="column is-3">
                <VField>
                  <VLabel>高度</VLabel>
                  <VControl>
                    <VInputNumber v-model="mv.key_params.h" centered :min="0" :step="1" />
                  </VControl>
                </VField>
              </div>
            </expand-transition>
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
