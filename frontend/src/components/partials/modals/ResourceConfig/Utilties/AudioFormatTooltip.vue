<script lang="ts" setup>
import { useUserSession } from "@src/stores/userSession";
import { defAudioFormat } from '../../PresetConfig/Consts';

const props = defineProps<{
  value?: AudioFormat,
  name?: string
}>()

const usStore = useUserSession()
const audioFormats = computed(() => usStore.settings.audio_formats || [])

const format = computed(() => {
  if (props.value) return props.value
  if (props.name) {
    return audioFormats.value.find(f => f.name === props.name) || defAudioFormat()
  }
  return defAudioFormat()
})

const tippyAppendTo = () => document.body
</script>
<template>
  <tippy
    animation="scale"
    theme="light"
    placement="right"
    :arrow="true"
    arrow-type="round"
    :interactive="true"
    :append-to="tippyAppendTo"
  >
    <template #content>
      <div class="kv-info tippy-detail">
        <div class="info-block-inner is-dark-bordered-12">
          <div class="info-block-line">
            <h4 class="dark-inverted">
              模板名称
            </h4>
            <span>
              {{ format.name }}
            </span>
          </div>
          <div class="info-block-line">
            <h4 class="dark-inverted">
              声道数量
            </h4>
            <span>
              {{ format.channels_number }}
            </span>
          </div>
          <div class="info-block-line">
            <h4 class="dark-inverted">
              量化 bit
            </h4>
            <span>
              {{ format.bits }}
            </span>
          </div>
          <div class="info-block-line">
            <h4 class="dark-inverted">
              采样率
            </h4>
            <span>
              {{ format.frequency }}
            </span>
          </div>
          <div class="info-block-line">
            <h4 class="dark-inverted">
              压缩格式
            </h4>
            <span>
              {{ format.compression_format }}
            </span>
          </div>
          <div v-if="format.compression_subtype" class="info-block-line">
            <h4 class="dark-inverted">
              子类型
            </h4>
            <span>
              {{ format.compression_subtype }}
            </span>
          </div>
          <div class="info-block-line">
            <h4 class="dark-inverted">
              发包间隔
            </h4>
            <span>
              {{ format.packet_time_us }}
            </span>
          </div>
          <div class="info-block-line">
            <h4 class="dark-inverted">
              音频码率
            </h4>
            <span>
              {{ format.bitrate_bps }} bps
            </span>
          </div>
        </div>
      </div>
    </template>
    <slot></slot>
  </tippy>
</template>
