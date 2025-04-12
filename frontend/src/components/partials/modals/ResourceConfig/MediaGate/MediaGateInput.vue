<script lang="ts" setup>
import { type IndexedNicDetail } from '../Utilties/Consts_V1';
import { def_mg_player_params, type MgInputParams } from './Consts';

const mv = defineModel<MgInputParams>({
  default: def_mg_player_params(),
  local: true,
});

defineProps<{
  usedSignalType: number,
  nics: IndexedNicDetail[]
}>()

const emit = defineEmits<{
  (e: 'video-selected', name: string): void,
  (e: 'video-unselected', name: string): void,
  (e: 'audio-selected', name: string): void,
  (e: 'audio-unselected', name: string): void
}>()
</script>
<template>
  <div class="form-outer">
    <div class="form-header">
      <div class="form-header-inner">
        <div class="left">
          <h4>输入参数</h4>
        </div>
      </div>
    </div>
    <div class="form-body">
      <!--Fieldset-->
      <div class="form-fieldset" :class="mv.videoformat_name && 'seperator'">
        <div class="columns is-multiline">
          <div class="column is-4">
            <VField>
              <VLabel>视频格式名称</VLabel>
              <VControl>
                <VideoFormatSelect v-model="mv.videoformat_name" :used-signal-type="usedSignalType" @video-selected="emit('video-selected', $event)" @video-unselected="emit('video-unselected', $event)" />
              </VControl>
            </VField>
          </div>
          <div class="column is-4">
            <VField>
              <VLabel>输入视频元数据是否强制使用指定的色域</VLabel>
              <VControl class="in-form">
                <VSwitchBlock
                  v-model="mv.force_use_videoformat"
                  color="primary"
                />
              </VControl>
            </VField>
          </div>
          <div class="column is-4">
            <VField>
              <VLabel>音频格式名称</VLabel>
              <VControl>
                <AudioFormatSelect v-model="mv.audioformat_name" :videoformat="mv.videoformat_name" @audio-selected="emit('audio-selected', $event)" @audio-unselected="emit('audio-selected', $event)" />
              </VControl>
            </VField>
          </div>
        </div>
      </div>
      <PlayerParamsForm key="input" v-model="mv" :nics="nics" smpte="receive" />
    </div>
  </div>
</template>
