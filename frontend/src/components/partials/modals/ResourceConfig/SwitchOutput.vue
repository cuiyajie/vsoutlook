<script setup lang="ts">
import { type SwitchOutputParamsType, formatMap, formatKeys, v_protocols } from './Consts';
import { useFormat, useProtocolDC } from './Utils';

const mv = defineModel<SwitchOutputParamsType>({
  default: {} as SwitchOutputParamsType,
  local: true,
});

const OPEN = defineModel('OPEN', {
  default: false,
  local: true
})

const props = defineProps<{
  title: string,
  toggleTitle?: string,
  isLast?: boolean,
  useBackup: boolean,
  m_local_ip: string,
  b_local_ip: string
}>()

const opened = ref(false)

const format = useFormat(mv.value, formatKeys[2]);
useProtocolDC(mv.value)

const isOpen = computed(() => opened.value && (!props.toggleTitle || OPEN.value))
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
      <h4>{{ title }}</h4>
      <div class="collapse-icon">
        <VIcon icon="feather:chevron-down" />
      </div>
    </div>
    <div
      v-if="opened && toggleTitle"
      class="columns is-multiline"
    >
      <div class="column is-12">
        <VField horizontal>
          <VLabel>{{ toggleTitle }}</VLabel>
          <VControl>
            <VSwitchBlock
              v-model="OPEN"
              color="primary"
            />
          </VControl>
        </VField>
      </div>
    </div>
    <Transition name="fade-slow">
      <div
        v-if="isOpen"
        class="form-fieldset-nested"
      >
        <div class="form-fieldset seperator">
          <div class="fieldset-heading">
            <h5>IP流参数</h5>
          </div>
          <div class="columns is-multiline">
            <div class="column is-6">
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
                  <VInput
                    v-model="mv.ipstream_master.v_dst_address"
                  />
                </VControl>
              </VField>
            </div>
            <div class="column is-6">
              <VField>
                <VLabel>主视频流前置SDN输入端口</VLabel>
                <VControl>
                  <VInputNumber
                    v-model="mv.ipstream_master.v_p4inport"
                    centered
                    :min="0"
                    :step="1"
                  />
                </VControl>
              </VField>
            </div>
            <div class="column is-6">
              <VField>
                <VLabel>主视频流前置SDN输出端口</VLabel>
                <VControl>
                  <VInputNumber
                    v-model="mv.ipstream_master.v_p4outport"
                    centered
                    :min="0"
                    :step="1"
                  />
                </VControl>
              </VField>
            </div>
            <Transition name="fade-slow">
              <div
                v-if="useBackup"
                class="column is-6"
              >
                <VField>
                  <VLabel>备视频流组播源IP（含端口）</VLabel>
                  <AddrAddon
                    v-model="mv.ipstream_backup.v_src_address"
                    :host="b_local_ip"
                  />
                </VField>
              </div>
            </Transition>
            <Transition name="fade-slow">
              <div
                v-if="useBackup"
                class="column is-6"
              >
                <VField>
                  <VLabel>备视频流组播目标IP（含端口）</VLabel>
                  <VControl>
                    <VInput
                      v-model="mv.ipstream_backup.v_dst_address"
                    />
                  </VControl>
                </VField>
              </div>
            </Transition>
            <Transition name="fade-slow">
              <div
                v-if="useBackup"
                class="column is-6"
              >
                <VField>
                  <VLabel>备视频流前置SDN输入端口</VLabel>
                  <VControl>
                    <VInputNumber
                      v-model="mv.ipstream_backup.v_p4inport"
                      centered
                      :min="0"
                      :step="1"
                    />
                  </VControl>
                </VField>
              </div>
            </Transition>
            <Transition name="fade-slow">
              <div
                v-if="useBackup"
                class="column is-6"
              >
                <VField>
                  <VLabel>备视频流前置SDN输入端口</VLabel>
                  <VControl>
                    <VInputNumber
                      v-model="mv.ipstream_backup.v_p4outport"
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
    <Transition name="fade-slow">
      <div
        v-if="isOpen"
        class="form-fieldset-nested"
      >
        <div class="form-fieldset seperator">
          <div class="fieldset-heading">
            <h5>视频参数</h5>
          </div>
          <div class="columns is-multiline">
            <div class="column is-6">
              <VField>
                <VLabel>视频格式</VLabel>
                <VControl>
                  <VInput
                    :model-value="formatMap[format]"
                    readonly
                  />
                </VControl>
              </VField>
            </div>
          </div>
        </div>
      </div>
    </Transition>
    <Transition name="fade-slow">
      <div
        v-if="isOpen && (mv.videoformat.v_compression_format || mv.videoformat.v_compression_ratio)"
        class="form-fieldset-nested is-tail"
      >
        <div class="form-fieldset">
          <div class="fieldset-heading">
            <h5>视频编码参数</h5>
          </div>
          <div class="columns is-multiline">
            <div v-if="mv.videoformat.v_compression_format" class="column is-6">
              <VField>
                <VLabel>编码格式</VLabel>
                <VControl>
                  <VInput
                    v-model="mv.videoformat.v_compression_format"
                    readonly
                  />
                </VControl>
              </VField>
            </div>
            <div v-if="mv.videoformat.v_compression_ratio" class="column is-6">
              <VField>
                <VLabel>压缩比</VLabel>
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
      </div>
    </Transition>
  </div>
</template>
