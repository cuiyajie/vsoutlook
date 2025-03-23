<script setup lang="ts">
import {
  type SwitchKeyParamsType,
  val_switch_keytype,
  def_switch_external_key_params,
  def_switch_internal_key_params,
} from '../Utilties/Consts'

const mv = defineModel<SwitchKeyParamsType>({
  default: {} as SwitchKeyParamsType,
  local: true,
})

watch(
  () => mv.value.keytype,
  (nv) => {
    if (nv === val_switch_keytype[0]) {
      mv.value = def_switch_external_key_params()
    } else if (nv === val_switch_keytype[1]) {
      mv.value = def_switch_internal_key_params()
    }
  },
  { immediate: true }
)

defineProps<{
  index: number
  isLast: boolean
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
      <h4>第{{ index }}路键设置</h4>
      <div class="collapse-icon">
        <VIcon icon="feather:chevron-down" />
      </div>
    </div>
    <Transition name="fade-show">
      <div v-show="opened" class="form-fieldset-nested">
        <div class="form-fieldset">
          <div class="columns is-multiline">
            <div class="column is-6">
              <VField>
                <VLabel>键类型</VLabel>
                <VControl>
                  <VSelect v-model="mv.keytype" class="is-rounded">
                    <VOption :value="val_switch_keytype[0]"> 外键 </VOption>
                    <VOption :value="val_switch_keytype[1]"> 内键 </VOption>
                  </VSelect>
                </VControl>
              </VField>
            </div>
            <div
              v-if="mv.keytype === val_switch_keytype[0] && 'ext_key_index' in mv"
              class="column is-6"
            >
              <VField>
                <VLabel>使用的外部键信号序号</VLabel>
                <VControl>
                  <VInputNumber v-model="mv.ext_key_index" centered :min="0" :step="1" />
                </VControl>
              </VField>
            </div>
            <template v-if="mv.keytype === val_switch_keytype[1] && 'int_key' in mv">
              <div class="column is-6">
                <VField>
                  <VLabel>內键文件名</VLabel>
                  <VControl>
                    <VInput v-model="mv.int_key.name" />
                  </VControl>
                </VField>
              </div>
              <div class="column is-6">
                <VField>
                  <VLabel>显示区域顶点X坐标</VLabel>
                  <VControl>
                    <VInputNumber
                      v-model="mv.int_key.top_x"
                      centered
                      :min="0"
                      :step="1"
                    />
                  </VControl>
                </VField>
              </div>
              <div class="column is-6">
                <VField>
                  <VLabel>显示区域顶点Y坐标</VLabel>
                  <VControl>
                    <VInputNumber
                      v-model="mv.int_key.top_y"
                      centered
                      :min="0"
                      :step="1"
                    />
                  </VControl>
                </VField>
              </div>
              <div class="column is-6">
                <VField>
                  <VLabel>显示区域宽度</VLabel>
                  <VControl>
                    <VInputNumber
                      v-model="mv.int_key.width"
                      centered
                      :min="0"
                      :step="1"
                    />
                  </VControl>
                </VField>
              </div>
              <div class="column is-6">
                <VField>
                  <VLabel>显示区域高度</VLabel>
                  <VControl>
                    <VInputNumber
                      v-model="mv.int_key.height"
                      centered
                      :min="0"
                      :step="1"
                    />
                  </VControl>
                </VField>
              </div>
            </template>
          </div>
        </div>
      </div>
    </Transition>
  </div>
</template>
<style lang="scss">
@import '@src/scss/abstracts/all';
@import '@src/scss/components/forms-outer';
</style>
