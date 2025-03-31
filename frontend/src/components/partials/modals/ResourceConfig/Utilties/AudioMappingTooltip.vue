<script lang="ts" setup>
import { useUserSession } from "@src/stores/userSession";
import { defAudioMapping } from '../../PresetConfig/Consts';
import { numberRangeToArray } from '../../PresetConfig/Utils'

const props = defineProps<{
  value?: AudioMapping
  name?: string
}>()

const usStore = useUserSession()
const audioMappings = computed(() => usStore.settings.audio_mappings || [])

const mapping = computed(() => {
  if (props.value) return props.value
  if (props.name) {
    return audioMappings.value.find(f => f.name === props.name) || defAudioMapping()
  }
  return defAudioMapping()
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
              模版名称
            </h4>
            <span>
              {{ mapping.name }}
            </span>
          </div>
          <div class="info-block-line">
            <h4 class="dark-inverted">
              声道复制
            </h4>
            <div>
              <div class="mapping-cols scrollbar-hide">
                <div class="mapping-col-heading">
                  <div>源声道</div>
                  <div>目标声道</div>
                </div>
                <div
                  v-for="(channel, ci) in mapping.copy_channels"
                  :key="ci"
                  class="mapping-col with-border"
                >
                  <div class="mapping-cell is-source">{{ channel.src_channel }}</div>
                  <div class="mapping-cell">{{ channel.dst_channel }}</div>
                </div>
              </div>
            </div>
          </div>
          <div class="info-block-line">
            <h4 class="dark-inverted">
              静音声道
            </h4>
            <div>
              <div class="mapping-cols scrollbar-hide">
                <div class="mapping-col-heading">静音声道</div>
                <div v-for="mute in numberRangeToArray(mapping.mute_channels)" :key="mute" class="mapping-col">
                  <div class="mapping-cell">{{ mute }}</div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </template>
    <slot></slot>
  </tippy>
</template>
