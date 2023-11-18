<script setup lang="ts">
import { type Config5 } from "./Consts";

const mv = defineModel<Config5>({
  default: {} as any,
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
      <h4>第{{ index }}路输入参数</h4>
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
            <div class="column is-8">
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
                <VLabel>主视频流P4交换机输入端口</VLabel>
                <VControl>
                  <VInputNumber
                    v-model="mv.V_M_P4_Port"
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
                class="column is-8"
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
                  <VLabel>备视频流P4交换机输入端口</VLabel>
                  <VControl>
                    <VInputNumber
                      v-model="mv.V_B_P4_Port"
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
    <Transition name="fade-show">
      <div
        v-show="opened"
        class="form-fieldset-nested is-tail"
      >
        <div class="form-fieldset">
          <div class="fieldset-heading">
            <h5>切换台映射</h5>
          </div>
          <div class="columns is-multiline">
            <div class="column is-4">
              <VField>
                <VLabel>切换台输入序号</VLabel>
                <VControl>
                  <VInputNumber
                    v-model="mv.SW_Index"
                    centered
                    :min="0"
                    :step="1"
                  />
                </VControl>
              </VField>
            </div>
            <div class="column is-8">
              <VField>
                <VLabel>切换台显示名称</VLabel>
                <VControl>
                  <VInput
                    v-model="mv.SW_DisPlayName"
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
@import "/@src/scss/abstracts/all";
@import "/@src/scss/components/forms-outer";
</style>
