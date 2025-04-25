<script lang="ts" setup>
import { type MVPipParams, type MVOutputItemParam, def_mv_pip_params } from './Consts'
import { type IndexedType, type IndexedNicDetail } from '../Utilties/Consts_V1'
import { useUserSession } from '@src/stores/userSession'

const usStore = useUserSession()
const templates = computed(() => usStore.settings.mv_template_list || [])

const mv = defineModel<MVOutputItemParam>({
  default: {} as MVOutputItemParam,
  local: true,
})

const props = defineProps<{
  usedSignalType: number,
  nics: IndexedNicDetail[],
  pipsNumber: number,
  index: number,
  isLast?: boolean,
  deletable?: boolean,
  defaultOpen?: boolean,
}>()

const emit = defineEmits<{
  (e: 'video-selected', name: string): void,
  (e: 'video-unselected', name: string): void,
  (e: 'audio-selected', name: string): void,
  (e: 'audio-unselected', name: string): void,
}>()

const opened = ref(false)

const pips = ref<IndexedType<MVPipParams>[]>([])

watch(() => props.pipsNumber, (newNumber) => {
  const len = pips.value.length
  const newParams = mv.value.pip_params
  if (newNumber < len) {
    pips.value = pips.value.slice(0, newNumber)
  } else if (newNumber > len) {
    pips.value = [
      ...pips.value,
      ...Array.from({ length: newNumber - len }, (_, i) => ({
        index: len + i,
        value: newParams[len + i] ? reactive(newParams[len + i]) : reactive(def_mv_pip_params(len + i)),
      })),
    ]
  }
}, { immediate: true })

watch(
  () => mv.value.pip_params,
  (newParams) => {
    pips.value = newParams.map((v, i) => ({
      index: i + 1,
      value: reactive(v),
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
      <h4>第 {{ index + 1 }} 路输出参数&nbsp;&nbsp;（序号: {{ index }}）</h4>
      <div class="collapse-icon">
        <VIcon icon="feather:chevron-down" />
      </div>
    </div>

    <expand-transition>
      <div v-show="opened" class="form-fieldset-nested-2">
        <div class="form-fieldset-nested-3 seperator">
          <div class="fieldset-heading">
            <h5>输出布局</h5>
          </div>
          <div class="columns is-multiline">
            <div class="column is-12">
              <VField>
                <VLabel>布局模板</VLabel>
                <Multiselect
                  v-model="mv.out_mv_template"
                  placeholder="选择应用类型"
                  value-prop="path"
                  label="name"
                  :style="{'--ms-max-height': '245px'}"
                  :options="templates"
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
                  <VInput :model-value="pipsNumber" readonly />
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
                <VLabel>窗口{{ pip.index + 1 }}名称&nbsp;&nbsp;索引: {{ pip.index }}</VLabel>
                <VInput v-model="pip.value.out_pip_name" />
              </VField>
            </div>
            <div class="column is-4">
              <VField>
                <VLabel>窗口{{ pip.index + 1 }} 关联输入视频序号</VLabel>
                <VInputNumber
                  v-model="pip.value.out_pip_video_index"
                  centered
                  :min="0"
                  :step="1"
                />
              </VField>
            </div>
            <div class="column is-4">
              <VField>
                <VLabel>窗口{{ pip.index }} tally索引</VLabel>
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
        <div class="form-fieldset-nested-1 seperator">
          <div class="columns is-multiline">
            <div class="column is-4">
              <VField>
                <VLabel>输出视频格式名称</VLabel>
                <VControl>
                  <VideoFormatSelect v-model="mv.videoformat_name" :used-signal-type="usedSignalType" @video-selected="emit('video-selected', $event)" @video-unselected="emit('video-unselected', $event)" />
                </VControl>
              </VField>
            </div>
            <div class="column is-4">
              <VField>
                <VLabel>输出音频格式名称</VLabel>
                <VControl>
                  <AudioFormatSelect v-model="mv.audioformat_name" :videoformat="mv.videoformat_name" @audio-selected="emit('audio-selected', $event)" @audio-unselected="emit('audio-unselected', $event)" />
                </VControl>
              </VField>
            </div>
            <div class="column is-4">
              <VField>
                <VLabel>音频直通源</VLabel>
                <VControl>
                  <Multiselect
                    v-model="mv.audio_src_pip_index"
                    placeholder="选择音频直通源"
                    value-prop="index"
                    :searchable="false"
                    :can-deselect="false"
                    :can-clear="false"
                    :style="{'--ms-max-height': '245px'}"
                    no-options-text="暂时没有配置音频直通源"
                    :options="pips"
                  >
                    <template #singlelabel="{ value }">
                      <div class="multiselect-single-label">
                        <span class="select-label-text">
                          {{ value.value.out_pip_name }}
                        </span>
                      </div>
                    </template>
                    <template #option="{ option }">
                      <span class="select-option-text">
                        {{ option.value.out_pip_name }}
                      </span>
                    </template>
                  </Multiselect>
                </VControl>
              </VField>
            </div>
          </div>
        </div>
        <PlayerParamsForm v-model="mv" :nics="nics" />
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
      margin-top: 6px;
      text-decoration: underline;
      letter-spacing: 0.025em;
    }
  }
}
</style>
