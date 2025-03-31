<script lang="ts" setup>
import { useUserSession } from "@src/stores/userSession";
import { type IndexedNicDetail } from '../Utilties/Consts_V1';
import { type MgOutputItemParams, convert_mode, dynamic_mode, luma_gain } from './Consts';

const mv = defineModel<MgOutputItemParams>({
  default: [],
  local: true,
});

const props = defineProps<{
  inputVideoFormat: string,
  usedSignalType: number,
  nics: IndexedNicDetail[],
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
  (e: 'delete'): void,
}>()

const usStore = useUserSession()
const videoFormats = computed(() => usStore.settings.video_formats || [])

const convertMode = computed(() => {
  const convert_mode_1 = convert_mode.filter(cm => cm.value !== 'dynamic')
  const vf = videoFormats.value.find(vf => vf.name === mv.value.videoformat_name)
  if (!vf?.gamma) return convert_mode_1
  const ivf = videoFormats.value.find(vf => vf.name === props.inputVideoFormat)
  if (!ivf?.gamma) return convert_mode_1
  if (['hlg', 'pq'].includes(ivf.gamma) && vf.gamma === 'sdr') {
    return [ ...convert_mode ]
  }
  return convert_mode_1
})

watch(convertMode, (nv) => {
  if (nv) {
    const cmode = mv.value.hdr_convert_params.convert_mode
    if (!nv.find(cm => cm.value === cmode)) {
      mv.value.hdr_convert_params.convert_mode = nv[0].value
    }
  }
}, { immediate: true })

function deleteItem() {
  mv.value.videoformat_name = ''
  mv.value.audioformat_name = ''
  emit('delete')
}

const opened = ref(false)

watch(() => props.defaultOpen, () => {
  opened.value = props.defaultOpen
}, { immediate: true })

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
      <h4>第 {{ index + 1 }} 路输出参数</h4>
      <div class="header-button-list">
        <a
          v-if="deletable"
          role="button"
          class="action-link"
          tabindex="0"
          @keydown.space.prevent.stop="deleteItem"
          @click.prevent.stop="deleteItem"
        >删除</a>
        <div class="collapse-icon">
          <VIcon icon="feather:chevron-down" />
        </div>
      </div>
    </div>

    <expand-transition>
      <div v-if="opened" class="form-fieldset-nested-2">
        <div class="form-fieldset-nested-3 seperator">
          <div class="columns is-multiline">
            <div class="column is-6">
              <VField>
                <VLabel>输出视频格式名称</VLabel>
                <VControl>
                  <VideoFormatSelect v-model="mv.videoformat_name" :used-signal-type="usedSignalType" @video-selected="emit('video-selected', $event)" @video-unselected="emit('video-unselected', $event)" />
                </VControl>
              </VField>
            </div>
            <div class="column is-6">
              <VField>
                <VLabel>输出音频格式名称</VLabel>
                <VControl>
                  <AudioFormatSelect v-model="mv.audioformat_name" @audio-selected="emit('audio-selected', $event)" @audio-unselected="emit('audio-unselected', $event)" />
                </VControl>
              </VField>
            </div>
          </div>
        </div>
        <div class="form-fieldset seperator">
          <div class="fieldset-heading">
            <h4>伽马/色域变换设置</h4>
          </div>
          <div class="columns is-multiline">
            <div class="column is-6">
              <VField>
                <VLabel>变换模式</VLabel>
                <VControl>
                  <VSelect v-model="mv.hdr_convert_params.convert_mode" class="is-rounded">
                    <VOption v-for="cm in convertMode" :key="cm.value" :value="cm.value">
                      {{ cm.label }}
                    </VOption>
                  </VSelect>
                </VControl>
              </VField>
            </div>
            <div class="column is-6">
              <VField>
                <VLabel>lut文件名</VLabel>
                <VControl>
                  <VInput v-model="mv.hdr_convert_params.lut_filename" type="text" />
                </VControl>
              </VField>
            </div>
            <expand-transition>
              <div v-if="mv.hdr_convert_params.convert_mode === 'dynamic'" class="column is-6">
                <VField>
                  <VLabel>动态变换模式</VLabel>
                  <VControl>
                    <VSelect v-model="mv.hdr_convert_params.dynamic_mode" class="is-rounded">
                      <VOption v-for="dm in dynamic_mode" :key="dm.value" :value="dm.value">
                        {{ dm.label }}
                      </VOption>
                    </VSelect>
                  </VControl>
                </VField>
              </div>
            </expand-transition>
            <expand-transition>
              <div v-if="mv.hdr_convert_params.convert_mode === 'dynamic'" class="column is-6">
                <VField>
                  <VLabel>亮度增益</VLabel>
                  <VControl>
                    <VSelect v-model="mv.hdr_convert_params.luma_gain" class="is-rounded">
                      <VOption v-for="lg in luma_gain" :key="lg" :value="lg">
                        {{ lg }}
                      </VOption>
                    </VSelect>
                  </VControl>
                </VField>
              </div>
            </expand-transition>
          </div>
        </div>
        <PlayerParamsForm v-model="mv" :nics="nics" />
      </div>
    </expand-transition>
  </div>
</template>
<style lang="scss" scoped>
.header-button-list {
  display: flex;
  justify-items: flex-end;
  align-items: center;
  gap: 24px;

  > button {
    width: 30px;
    height: 30px;
  }
}
</style>
