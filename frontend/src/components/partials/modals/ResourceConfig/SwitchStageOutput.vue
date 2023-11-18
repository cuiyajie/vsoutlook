<script setup lang="ts">
import { type Config2, formatMap, formatKeys, vProtocols } from './Consts';
import { useFormat, useProtocolDC } from './Utils';

const mv = defineModel<Config2>({
  default: {} as any,
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
  useBackup: boolean
}>()

const opened = ref(false)

const V_Format = useFormat(mv, formatKeys[2]);
useProtocolDC(mv)

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
            <div class="column is-4">
              <VField>
                <VLabel>视频流协议</VLabel>
                <VControl>
                  <VSelect
                    v-model="mv.V_Protocol"
                    class="is-rounded"
                  >
                    <VOption
                      v-for="vp in vProtocols"
                      :key="vp"
                      :value="vp"
                    >
                      {{ vp }}
                    </VOption>
                  </VSelect>
                </VControl>
              </VField>
            </div>
            <div class="column is-4">
              <VField>
                <VLabel>主视频流组播地址（含端口）</VLabel>
                <VControl>
                  <VInput
                    v-model="mv.V_M_Address"
                  />
                </VControl>
              </VField>
            </div>
            <div class="column is-4">
              <VField>
                <VLabel>主音频流组播地址（含端口）</VLabel>
                <VControl>
                  <VInput
                    v-model="mv.A_M_Address"
                  />
                </VControl>
              </VField>
            </div>
            <Transition name="fade-slow">
              <div
                v-if="useBackup"
                class="column is-4"
              >
                <VField>
                  <VLabel>备视频流组播地址（含端口）</VLabel>
                  <VControl>
                    <VInput
                      v-model="mv.V_B_Address"
                    />
                  </VControl>
                </VField>
              </div>
            </Transition>
            <Transition name="fade-slow">
              <div
                v-if="useBackup"
                class="column is-4"
              >
                <VField>
                  <VLabel>备音频流组播地址（含端口）</VLabel>
                  <VControl>
                    <VInput
                      v-model="mv.A_B_Address"
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
                    :model-value="formatMap[V_Format]"
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
        v-if="isOpen"
        class="form-fieldset-nested is-tail"
      >
        <div class="form-fieldset">
          <div class="fieldset-heading">
            <h5>视频编码参数</h5>
          </div>
          <div class="columns is-multiline">
            <div class="column is-6">
              <VField>
                <VLabel>编码格式</VLabel>
                <VControl>
                  <VInput
                    v-model="mv.V_DecFormat"
                    readonly
                  />
                </VControl>
              </VField>
            </div>
            <div class="column is-6">
              <VField>
                <VLabel>压缩比</VLabel>
                <VControl>
                  <VInput
                    v-model="mv.V_CompressionRatio"
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
