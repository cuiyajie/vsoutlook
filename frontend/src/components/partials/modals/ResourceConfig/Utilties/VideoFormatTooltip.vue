<script lang="ts" setup>
import { useUserSession } from "@src/stores/userSession";
import { defVideoFormat, VideoFormatProtocolsMap } from '../../PresetConfig/Consts';

const props = defineProps<{
  value?: VideoFormat,
  name?: string
}>()

const usStore = useUserSession()
const videoFormats = computed(() => usStore.settings.video_formats || [])

const format = computed(() => {
  if (props.value) return props.value
  if (props.name) {
    return videoFormats.value.find(f => f.name === props.name) || defVideoFormat()
  }
  return defVideoFormat()
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
              格式类型
            </h4>
            <span>
              {{ VideoFormatProtocolsMap[format.protocol] }}
            </span>
          </div>
          <div class="info-block-line">
            <h4 class="dark-inverted">
              分辨率
            </h4>
            <span>
              {{ `${format.width}x${format.height}` }}
            </span>
          </div>
          <div class="info-block-line">
            <h4 class="dark-inverted">
              帧率
            </h4>
            <span>
              {{ format.fps }}
            </span>
          </div>
          <div class="info-block-line">
            <h4 class="dark-inverted">
              伽马
            </h4>
            <span>
              {{ format.gamma }}
            </span>
          </div>
          <div class="info-block-line">
            <h4 class="dark-inverted">
              色域
            </h4>
            <span>
              {{ format.gamut }}
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
              压缩子类型
            </h4>
            <span>
              {{ format.compression_subtype }}
            </span>
          </div>
          <div v-if="format.compression_ratio" class="info-block-line">
            <h4 class="dark-inverted">
              压缩比
            </h4>
            <span>
              {{ format.compression_ratio }}
            </span>
          </div>
          <div v-if="format.bitrate_bps" class="info-block-line">
            <h4 class="dark-inverted">
              视频码率
            </h4>
            <span>
              {{ format.bitrate_bps }} bps
            </span>
          </div>
          <div v-if="format.gop_b_frames" class="info-block-line">
            <h4 class="dark-inverted">
              B帧数量
            </h4>
            <span>
              {{ format.gop_b_frames }}
            </span>
          </div>
          <div v-if="format.gop_length" class="info-block-line">
            <h4 class="dark-inverted">
              GOP长度
            </h4>
            <span>
              {{ format.gop_length }}
            </span>
          </div>
        </div>
      </div>
    </template>
    <slot></slot>
  </tippy>
</template>
