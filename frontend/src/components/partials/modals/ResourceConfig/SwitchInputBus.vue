<script setup lang="ts">
import { type SwitchInputBusParamsType } from "./Consts";

const mv = defineModel<SwitchInputBusParamsType>({
  default: {} as SwitchInputBusParamsType,
  local: true,
});

defineProps<{
  useBackup: boolean
}>()

const video_opened = ref(false)
const keyfill_opened = ref(false)
</script>
<template>
  <div
    class="form-fieldset collapse-form seperator"
    :open="video_opened || undefined"
  >
    <div
      class="fieldset-heading collapse-header"
      tabindex="0"
      role="button"
      @keydown.space.prevent="video_opened = !video_opened"
      @click.prevent="video_opened = !video_opened"
    >
      <h4>视频输入母线</h4>
      <div class="collapse-icon">
        <VIcon icon="feather:chevron-down" />
      </div>
    </div>
    <Transition name="fade-show">
      <div
        v-show="video_opened"
        class="form-fieldset"
      >
        <div class="columns is-multiline">
          <div class="column is-4">
            <VField>
              <VLabel>主视频母线a组播源IP（含端口）</VLabel>
              <VControl>
                <VInput
                  v-model="mv.video_bus_master[0].v_src_address"
                />
              </VControl>
            </VField>
          </div>
          <div class="column is-4">
            <VField>
              <VLabel>主视频母线a组播目标IP（含端口）</VLabel>
              <VControl>
                <VInput
                  v-model="mv.video_bus_master[0].v_dst_address"
                />
              </VControl>
            </VField>
          </div>
          <div class="column is-4">
            <VField>
              <VLabel>主视频母线a使用的前置SDN输出端口</VLabel>
              <VControl>
                <VInputNumber
                  v-model="mv.video_bus_master[0].v_p4_port"
                  centered
                  :min="0"
                  :step="1"
                />
              </VControl>
            </VField>
          </div>
          <div class="column is-4">
            <VField>
              <VLabel>主视频母线b组播源IP（含端口）</VLabel>
              <VControl>
                <VInput
                  v-model="mv.video_bus_master[1].v_src_address"
                />
              </VControl>
            </VField>
          </div>
          <div class="column is-4">
            <VField>
              <VLabel>主视频母线b组播目标IP（含端口）</VLabel>
              <VControl>
                <VInput
                  v-model="mv.video_bus_master[1].v_dst_address"
                />
              </VControl>
            </VField>
          </div>
          <div class="column is-4">
            <VField>
              <VLabel>主视频母线b使用的前置SDN输出端口</VLabel>
              <VControl>
                <VInputNumber
                  v-model="mv.video_bus_master[1].v_p4_port"
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
              class="column is-4"
            >
              <VField>
                <VLabel>备视频母线a组播源IP（含端口）</VLabel>
                <VControl>
                  <VInput
                    v-model="mv.video_bus_backup[0].v_src_address"
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
                <VLabel>备视频母线a组播目标IP（含端口）</VLabel>
                <VControl>
                  <VInput
                    v-model="mv.video_bus_backup[0].v_dst_address"
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
                <VLabel>备视频母线a使用的前置SDN输出端口</VLabel>
                <VControl>
                  <VInputNumber
                    v-model="mv.video_bus_backup[0].v_p4_port"
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
              class="column is-4"
            >
              <VField>
                <VLabel>备视频母线b组播源IP（含端口）</VLabel>
                <VControl>
                  <VInput
                    v-model="mv.video_bus_backup[1].v_src_address"
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
                <VLabel>备视频母线b组播目标IP（含端口）</VLabel>
                <VControl>
                  <VInput
                    v-model="mv.video_bus_backup[1].v_dst_address"
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
                <VLabel>备视频母线b使用的前置SDN输出端口</VLabel>
                <VControl>
                  <VInputNumber
                    v-model="mv.video_bus_backup[1].v_p4_port"
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
    </Transition>
  </div>
  <div
    class="form-fieldset collapse-form"
    :open="keyfill_opened || undefined"
  >
    <div
      class="fieldset-heading collapse-header"
      tabindex="0"
      role="button"
      @keydown.space.prevent="keyfill_opened = !keyfill_opened"
      @click.prevent="keyfill_opened = !keyfill_opened"
    >
      <h4>键输入母线</h4>
      <div class="collapse-icon">
        <VIcon icon="feather:chevron-down" />
      </div>
    </div>
    <Transition name="fade-show">
      <div
        v-show="keyfill_opened"
        class="form-fieldset"
      >
        <div class="columns is-multiline">
          <div class="column is-4">
            <VField>
              <VLabel>主key母线组播源IP（含端口）</VLabel>
              <VControl>
                <VInput
                  v-model="mv.keyfill_bus_master.key_src_address"
                />
              </VControl>
            </VField>
          </div>
          <div class="column is-4">
            <VField>
              <VLabel>主key母线组播目标IP（含端口）</VLabel>
              <VControl>
                <VInput
                  v-model="mv.keyfill_bus_master.key_dst_address"
                />
              </VControl>
            </VField>
          </div>
          <div class="column is-4">
            <VField>
              <VLabel>主key母线使用的前置SDN输出端口</VLabel>
              <VControl>
                <VInputNumber
                  v-model="mv.keyfill_bus_master.key_p4_port"
                  centered
                  :min="0"
                  :step="1"
                />
              </VControl>
            </VField>
          </div>
          <div class="column is-4">
            <VField>
              <VLabel>主fill母线组播源IP（含端口）</VLabel>
              <VControl>
                <VInput
                  v-model="mv.keyfill_bus_master.fill_src_address"
                />
              </VControl>
            </VField>
          </div>
          <div class="column is-4">
            <VField>
              <VLabel>主fill母线组播目标IP（含端口）</VLabel>
              <VControl>
                <VInput
                  v-model="mv.keyfill_bus_master.fill_dst_address"
                />
              </VControl>
            </VField>
          </div>
          <div class="column is-4">
            <VField>
              <VLabel>主fill母线使用的前置SDN输出端口</VLabel>
              <VControl>
                <VInputNumber
                  v-model="mv.keyfill_bus_master.fill_p4_port"
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
              class="column is-4"
            >
              <VField>
                <VLabel>备key母线组播源IP（含端口）</VLabel>
                <VControl>
                  <VInput
                    v-model="mv.keyfill_bus_backup.key_src_address"
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
                <VLabel>备key母线组播目标IP（含端口）</VLabel>
                <VControl>
                  <VInput
                    v-model="mv.keyfill_bus_backup.key_dst_address"
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
                <VLabel>备key母线使用的前置SDN输出端口</VLabel>
                <VControl>
                  <VInputNumber
                    v-model="mv.keyfill_bus_backup.key_p4_port"
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
              class="column is-4"
            >
              <VField>
                <VLabel>备fill母线组播源IP（含端口）</VLabel>
                <VControl>
                  <VInput
                    v-model="mv.keyfill_bus_backup.fill_src_address"
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
                <VLabel>备fill母线组播目标IP（含端口）</VLabel>
                <VControl>
                  <VInput
                    v-model="mv.keyfill_bus_backup.fill_dst_address"
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
                <VLabel>备fill母线使用的前置SDN输出端口</VLabel>
                <VControl>
                  <VInputNumber
                    v-model="mv.keyfill_bus_backup.fill_p4_port"
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
    </Transition>
  </div>
</template>
<style lang="scss">
@import "/@src/scss/abstracts/all";
@import "/@src/scss/components/forms-outer";
</style>
