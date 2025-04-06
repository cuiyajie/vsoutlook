<script setup lang="ts">
import {
  def_switch_input_key_params,
  def_switch_input_video_params,
  type SwitchInputVideoParams,
  type SwitchInput,
  type SwitchInputKeyParams
} from './Consts'

const mv = defineModel<SwitchInput>({
  default: {} as SwitchInput,
  local: true,
})

const keyParams = computed({
  get: () => mv.value.key,
  set: (value) => mv.value.key = value
})
const videoParams = computed({
  get: () => mv.value.video,
  set: (value) => mv.value.video = value
})

function addKeyParam() {
  keyParams.value.push(def_switch_input_key_params(keyParams.value.length))
}

function removeKeyParam(idx: number) {
  keyParams.value.splice(idx, 1)
}

function addVideoParam() {
  videoParams.value.push(def_switch_input_video_params(videoParams.value.length))
}

function removeVideoParam(idx: number) {
  videoParams.value.splice(idx, 1)
}

</script>
<template>
  <div class="form-outer has-mt-20">
    <div class="form-header is-sticky">
      <div class="form-header-inner">
        <div class="left">
          <h4>键输入信号源配置 {{ keyParams.length > 0 ? `- ${keyParams.length} 路` : '' }}</h4>
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
        <SwitchInputKeyParams
          v-for="(keyParam, kidx) in keyParams"
          :key="`input-key-${keyParam.index}`"
          v-model="keyParams[kidx]"
          :index="kidx"
          :is-last="kidx === keyParams.length - 1"
          @remove="removeKeyParam(kidx)"
        />
      </TransitionGroup>
      <div v-if="keyParams.length === 0" class="is-empty-list">暂时还没有配置键输入信号源</div>
    </div>
  </div>
  <div class="form-outer has-mb-20">
    <div class="form-header is-sticky">
      <div class="form-header-inner">
        <div class="left">
          <h4>视频输入信号源配置 {{ videoParams.length > 0 ? `- ${videoParams.length} 路` : '' }}</h4>
        </div>
        <button
          class="button is-circle is-dark-outlined"
          @keydown.space.prevent="addVideoParam"
          @click.prevent="addVideoParam"
        >
          <span class="icon is-small">
            <i aria-hidden="true" class="iconify" data-icon="feather:plus" />
          </span>
        </button>
      </div>
    </div>
    <div class="form-body">
      <TransitionGroup name="list">
        <SwitchInputVideoParams
          v-for="(videoParam, vidx) in videoParams"
          :key="`input-video-${videoParam.index}`"
          v-model="videoParams[vidx]"
          :index="vidx"
          :is-last="vidx === videoParams.length - 1"
          @remove="removeVideoParam(vidx)"
        />
      </TransitionGroup>
      <div v-if="videoParams.length === 0" class="is-empty-list">暂时还没有配置视频输入信号源</div>
    </div>
  </div>
</template>
