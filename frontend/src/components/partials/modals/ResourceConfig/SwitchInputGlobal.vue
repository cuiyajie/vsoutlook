<script setup lang="ts">
import { getFormat, useFormat, useProtocolDC } from './Utils';
import { formats, v_protocols, type SwitchInputParamsType } from './Consts';

const mv = defineModel<SwitchInputParamsType>({
  default: {} as SwitchInputParamsType,
  local: true,
});

const format = useFormat(mv.value, getFormat(mv.value));
useProtocolDC(mv.value);

const key_opened = ref(false)
const fill_opened = ref(false)
</script>
<template>
  <div class="form-fieldset">
    <div class="fieldset-heading">
      <h4>全局参数</h4>
    </div>
    <div class="columns is-multiline">
      <div class="column is-3">
        <VField>
          <VLabel>启用2022-7备份</VLabel>
          <VControl>
            <VSwitchBlock
              v-model="mv['g_2022-7']"
              color="primary"
            />
          </VControl>
        </VField>
      </div>
      <div class="column is-3">
        <VField>
          <VLabel>视频流协议</VLabel>
          <VControl>
            <VSelect
              v-model="mv.v_protocol"
              class="is-rounded"
            >
              <VOption
                v-for="vp in v_protocols"
                :key="vp"
                :value="vp"
              >
                {{ vp }}
              </VOption>
            </VSelect>
          </VControl>
        </VField>
      </div>
      <div class="column is-6">
        <VField>
          <VLabel>视频格式</VLabel>
          <VControl>
            <VSelect
              v-model="format"
              class="is-rounded"
            >
              <VOption
                v-for="f in formats"
                :key="f.key"
                :value="f.key"
              >
                {{ f.value }}
              </VOption>
            </VSelect>
          </VControl>
        </VField>
      </div>
      <div class="column is-6">
        <VField>
          <VLabel>视频编码格式</VLabel>
          <VControl>
            <VInput
              v-model="mv.videoformat.v_compression_format"
              readonly
            />
          </VControl>
        </VField>
      </div>
      <div class="column is-6">
        <VField>
          <VLabel>视频压缩比</VLabel>
          <VControl>
            <VInput
              v-model="mv.videoformat.v_compression_ratio"
              readonly
            />
          </VControl>
        </VField>
      </div>
    </div>
  </div>
  <div
    class="form-fieldset collapse-form seperator"
    :open="key_opened || undefined"
  >
    <div
      class="fieldset-heading collapse-header"
      tabindex="0"
      role="button"
      @keydown.space.prevent="key_opened = !key_opened"
      @click.prevent="key_opened = !key_opened"
    >
      <h4>第1路key输入参数</h4>
      <div class="collapse-icon">
        <VIcon icon="feather:chevron-down" />
      </div>
    </div>
    <Transition name="fade-show">
      <div
        v-show="key_opened"
        class="form-fieldset-nested"
      >
        <div class="form-fieldset seperator">
          <div class="fieldset-heading">
            <h5>IP流参数</h5>
          </div>
          <div class="columns is-multiline">
            <div class="column is-6">
              <VField>
                <VLabel>主视频流组播源IP（含端口）</VLabel>
                <VControl>
                  <VInput
                    v-model="mv.input_key_params.ipstream_master.v_src_address"
                  />
                </VControl>
              </VField>
            </div>
            <div class="column is-6">
              <VField>
                <VLabel>主视频流组播目标IP（含端口）</VLabel>
                <VControl>
                  <VInput
                    v-model="mv.input_key_params.ipstream_master.v_dst_address"
                  />
                </VControl>
              </VField>
            </div>
            <div class="column is-6">
              <VField>
                <VLabel>主视频流IO输入端口</VLabel>
                <VControl>
                  <VInputNumber
                    v-model="mv.input_key_params.ipstream_master.Key_P4_port"
                    centered
                    :min="0"
                    :step="1"
                  />
                </VControl>
              </VField>
            </div>
            <div class="column is-6" />
            <Transition name="fade-slow">
              <div
                v-if="mv['g_2022-7']"
                class="column is-6"
              >
                <VField>
                  <VLabel>备视频流组播源IP（含端口）</VLabel>
                  <VControl>
                    <VInput
                      v-model="mv.input_key_params.ipstream_backup.v_src_address"
                    />
                  </VControl>
                </VField>
              </div>
            </Transition>
            <Transition name="fade-slow">
              <div
                v-if="mv['g_2022-7']"
                class="column is-6"
              >
                <VField>
                  <VLabel>备视频流组播目标IP（含端口）</VLabel>
                  <VControl>
                    <VInput
                      v-model="mv.input_key_params.ipstream_backup.v_dst_address"
                    />
                  </VControl>
                </VField>
              </div>
            </Transition>
            <Transition name="fade-slow">
              <div
                v-if="mv['g_2022-7']"
                class="column is-6"
              >
                <VField>
                  <VLabel>备视频流IO输入端口</VLabel>
                  <VControl>
                    <VInputNumber
                      v-model="mv.input_key_params.ipstream_backup.Key_P4_port"
                      centered
                      :min="0"
                      :step="1"
                    />
                  </VControl>
                </VField>
              </div>
            </Transition>
          </div>
        </div>
      </div>
    </Transition>
  </div>
  <div
    class="form-fieldset collapse-form seperator"
    :open="fill_opened || undefined"
  >
    <div
      class="fieldset-heading collapse-header"
      tabindex="0"
      role="button"
      @keydown.space.prevent="fill_opened = !fill_opened"
      @click.prevent="fill_opened = !fill_opened"
    >
      <h4>第1路fill输入参数</h4>
      <div class="collapse-icon">
        <VIcon icon="feather:chevron-down" />
      </div>
    </div>
    <Transition name="fade-show">
      <div
        v-show="fill_opened"
        class="form-fieldset-nested"
      >
        <div class="form-fieldset seperator">
          <div class="fieldset-heading">
            <h5>IP流参数</h5>
          </div>
          <div class="columns is-multiline">
            <div class="column is-6">
              <VField>
                <VLabel>主视频流组播源IP（含端口）</VLabel>
                <VControl>
                  <VInput
                    v-model="mv.input_fill_params.ipstream_master.v_src_address"
                  />
                </VControl>
              </VField>
            </div>
            <div class="column is-6">
              <VField>
                <VLabel>主视频流组播目标IP（含端口）</VLabel>
                <VControl>
                  <VInput
                    v-model="mv.input_fill_params.ipstream_master.v_dst_address"
                  />
                </VControl>
              </VField>
            </div>
            <div class="column is-6">
              <VField>
                <VLabel>主视频流IO输入端口</VLabel>
                <VControl>
                  <VInputNumber
                    v-model="mv.input_fill_params.ipstream_master.fill_P4_port"
                    centered
                    :min="0"
                    :step="1"
                  />
                </VControl>
              </VField>
            </div>
            <div class="column is-6" />
            <Transition name="fade-slow">
              <div
                v-if="mv['g_2022-7']"
                class="column is-6"
              >
                <VField>
                  <VLabel>备视频流组播源IP（含端口）</VLabel>
                  <VControl>
                    <VInput
                      v-model="mv.input_fill_params.ipstream_backup.v_src_address"
                    />
                  </VControl>
                </VField>
              </div>
            </Transition>
            <Transition name="fade-slow">
              <div
                v-if="mv['g_2022-7']"
                class="column is-6"
              >
                <VField>
                  <VLabel>备视频流组播目标IP（含端口）</VLabel>
                  <VControl>
                    <VInput
                      v-model="mv.input_fill_params.ipstream_backup.v_dst_address"
                    />
                  </VControl>
                </VField>
              </div>
            </Transition>
            <Transition name="fade-slow">
              <div
                v-if="mv['g_2022-7']"
                class="column is-6"
              >
                <VField>
                  <VLabel>备视频流IO输入端口</VLabel>
                  <VControl>
                    <VInputNumber
                      v-model="mv.input_fill_params.ipstream_backup.fill_P4_port"
                      centered
                      :min="0"
                      :step="1"
                    />
                  </VControl>
                </VField>
              </div>
            </Transition>
          </div>
        </div>
      </div>
    </Transition>
  </div>
</template>
