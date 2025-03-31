<!-- eslint-disable vue/prop-name-casing -->
<script setup lang="ts">
import { useFormat, getFormat, useProtocolDC } from '../Utilties/Utils'
import {
  formats,
  type MVOutputParamsType,
  def_mv_pip_params,
  type SwitchMVPipParamsType,
} from '../Utilties/Consts'
import { useUserSession } from '@src/stores/userSession'

const usStore = useUserSession()
const templates = computed(() => usStore.settings.mv_template_list || [])

const mv = defineModel<MVOutputParamsType>({
  default: {} as MVOutputParamsType,
  local: true,
})

defineProps<{
  index: number
  isLast: boolean
  useBackup: boolean
  m_local_ip: string
  b_local_ip: string
}>()

const opened = ref(false)

const format = useFormat(mv, getFormat(mv.value))
useProtocolDC(mv)

const pips = ref<
  Array<{
    index: number
    value: Ref<SwitchMVPipParamsType>
  }>
>([])

watch(
  () => mv.value.pips_number,
  (nv) => {
    const len = pips.value.length
    const params: any[] = mv.value.pip_params
    if (nv < len) {
      pips.value = pips.value.slice(0, nv)
    } else if (nv > len) {
      pips.value = [
        ...pips.value,
        ...Array.from({ length: nv - len }, (_, i) => ({
          index: len + i + 1,
          value: params[len + i] ? ref(params[len + i]) : ref(def_mv_pip_params(len + i)),
        })),
      ]
    }
  },
  { immediate: true }
)

watch(
  () => mv.value.pip_params,
  () => {
    pips.value = mv.value.pip_params.map((v, i) => ({
      index: i + 1,
      value: ref(v),
    }))
  },
  { immediate: true }
)

defineExpose({
  getValue() {
    return {
      ...mv.value,
      pip_params: pips.value.map((v) => v.value),
    }
  },
})
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
      <h4>第{{ index }}路输出参数</h4>
      <div class="collapse-icon">
        <VIcon icon="feather:chevron-down" />
      </div>
    </div>
    <expand-transition>
      <div v-show="opened" class="form-fieldset-nested is-tail">
        <div class="form-fieldset">
          <div class="fieldset-heading">
            <h5>输出布局</h5>
          </div>
          <div class="columns is-multiline">
            <div class="column is-12">
              <VField v-slot="{ field }">
                <VLabel>布局模板</VLabel>
                <Multiselect
                  v-model="mv.mv_template"
                  placeholder="选择应用类型"
                  value-prop="path"
                  label="name"
                  :max-height="145"
                  :options="templates"
                  @change="(val: any) => field?.setValue(val)"
                >
                  <template #singlelabel="{ value }">
                    <div class="multiselect-single-label">
                      <span class="select-label-text">
                        {{ value.name }} ({{ value.path }})
                      </span>
                    </div>
                  </template>
                  <template #option="{ option }">
                    <div class="mv-tmpl-dropdown">
                      <div>
                        {{ option.name }}
                      </div>
                      <div>
                        {{ option.path }}
                      </div>
                    </div>
                  </template>
                </Multiselect>
              </VField>
            </div>
            <div class="column is-6">
              <VField>
                <VLabel>窗口数量</VLabel>
                <VControl>
                  <VSelect v-model="mv.pips_number" class="is-rounded">
                    <VOption :value="4"> 4 </VOption>
                    <VOption :value="6"> 6 </VOption>
                    <VOption :value="8"> 8 </VOption>
                    <VOption :value="9"> 9 </VOption>
                    <VOption :value="10"> 10 </VOption>
                  </VSelect>
                </VControl>
              </VField>
            </div>
            <div class="column is-6">
              <VField>
                <VLabel>tally-屏幕索引</VLabel>
                <VControl>
                  <VInputNumber v-model="mv.screenindex" centered :min="0" :step="1" />
                </VControl>
              </VField>
            </div>
          </div>
          <div v-for="pip in pips" :key="pip.index" class="columns is-multiline">
            <div class="column is-4">
              <VField>
                <VLabel>窗口{{ pip.index }}名称</VLabel>
                <VInput v-model="pip.value.pip_name" />
              </VField>
            </div>
            <div class="column is-4">
              <VField>
                <VLabel>窗口{{ pip.index }}关联输入视频序号</VLabel>
                <VInputNumber
                  v-model="pip.value.pip_video_index"
                  centered
                  :min="0"
                  :step="1"
                />
              </VField>
            </div>
            <div class="column is-4">
              <VField>
                <VLabel>窗口{{ pip.index }}tally索引</VLabel>
                <VInputNumber
                  v-model="pip.value.tallyindex"
                  centered
                  :min="0"
                  :step="1"
                />
              </VField>
            </div>
          </div>
        </div>
      </div>
    </expand-transition>
    <expand-transition>
      <div v-show="opened" class="form-fieldset-nested">
        <div class="form-fieldset seperator">
          <div class="fieldset-heading">
            <h5>IP流参数</h5>
          </div>
          <div class="columns is-multiline">
            <div class="column is-6">
              <VField>
                <VLabel>视频流协议</VLabel>
                <VControl>
                  <VInput v-model="mv.v_protocol" readonly />
                </VControl>
              </VField>
            </div>
            <div class="column is-6" />
            <div class="column is-6">
              <VField>
                <VLabel>主视频流组播源IP（含端口）</VLabel>
                <AddrAddon
                  v-model="mv.ipstream_master.v_src_address"
                  :host="m_local_ip"
                />
              </VField>
            </div>
            <div class="column is-6">
              <VField>
                <VLabel>主视频流组播目标IP（含端口）</VLabel>
                <VControl>
                  <VInput v-model="mv.ipstream_master.v_dst_address" />
                </VControl>
              </VField>
            </div>
            <expand-transition>
              <div v-if="useBackup" class="column is-6">
                <VField>
                  <VLabel>备视频流组播源IP（含端口）</VLabel>
                  <AddrAddon
                    v-model="mv.ipstream_backup.v_src_address"
                    :host="b_local_ip"
                  />
                </VField>
              </div>
            </expand-transition>
            <expand-transition>
              <div v-if="useBackup" class="column is-6">
                <VField>
                  <VLabel>备视频流组播目标IP（含端口）</VLabel>
                  <VControl>
                    <VInput v-model="mv.ipstream_backup.v_dst_address" />
                  </VControl>
                </VField>
              </div>
            </expand-transition>
          </div>
        </div>
      </div>
    </expand-transition>
    <expand-transition>
      <div v-show="opened" class="form-fieldset-nested is-tail">
        <div class="form-fieldset">
          <div class="fieldset-heading">
            <h5>视频参数</h5>
          </div>
          <div class="columns is-multiline">
            <div class="column is-6">
              <VField>
                <VLabel>视频格式</VLabel>
                <VControl>
                  <VSelect v-model="format" class="is-rounded">
                    <VOption v-for="f in formats" :key="f.key" :value="f.key">
                      {{ f.value }}
                    </VOption>
                  </VSelect>
                </VControl>
              </VField>
            </div>
          </div>
        </div>
      </div>
    </expand-transition>
  </div>
</template>
<style lang="scss">
.mv-tmpl-dropdown {
  > div {
    &:first-child {
    }

    &:last-child {
      font-size: 0.9em;
      color: var(--light-text);
      font-family: var(--font);
    }
  }
}
</style>
