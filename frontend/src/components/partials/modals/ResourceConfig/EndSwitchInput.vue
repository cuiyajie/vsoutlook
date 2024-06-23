<script setup lang="ts">
import { type EndSwitchInputVideoParamsType } from "./Consts";

const mv = defineModel<EndSwitchInputVideoParamsType>({
  default: {} as EndSwitchInputVideoParamsType,
  local: true,
});

defineProps<{
  index: number,
  isLast: boolean,
  useBackup: boolean
}>()

const opened = ref(false)
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
      <h4>第{{ index }}路视频输入参数</h4>
      <div class="collapse-icon">
        <VIcon icon="feather:chevron-down" />
      </div>
    </div>
    <Transition name="fade-show">
      <div
        v-show="opened"
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
                    v-model="mv.ipstream_master.v_src_address"
                  />
                </VControl>
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
            <Transition name="fade-slow">
              <div
                v-if="useBackup"
                class="column is-6"
              >
                <VField>
                  <VLabel>备视频流组播源IP（含端口）</VLabel>
                  <VControl>
                    <VInput
                      v-model="mv.ipstream_backup.v_src_address"
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
                  <VLabel>备视频流组播目标IP（含端口）</VLabel>
                  <VControl>
                    <VInput
                      v-model="mv.ipstream_backup.v_dst_address"
                    />
                  </VControl>
                </VField>
              </div>
            </Transition>
            <div class="column is-6">
              <VField>
                <VLabel>显示名称</VLabel>
                <VControl>
                  <VInput
                    v-model="mv.sw_displayname"
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
<style lang="scss">
@import "@src/scss/abstracts/all";
@import "@src/scss/components/forms-outer";
</style>
