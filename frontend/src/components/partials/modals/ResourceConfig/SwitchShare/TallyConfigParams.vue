<script lang="ts" setup>
import { type TallyConfigParams } from './Consts'

const mv = defineModel<TallyConfigParams>({
  default: {} as TallyConfigParams,
  local: true,
})

defineProps<{
  index: number
  isLast: boolean
}>()

const opened = ref(false)
const showSecondScreen = ref(false)

</script>
<template>
  <div
    class="form-fieldset collapse-form"
    :class="!isLast && 'seperator'"
    :open="opened || undefined"
  >
    <div
      class="fieldset-heading collapse-header"
      :style="opened ? { marginBottom: '10px !important' } : {}"
      tabindex="0"
      role="button"
      @keydown.space.prevent="opened = !opened"
      @click.prevent="opened = !opened"
    >
      <h4>第 {{ index + 1 }} 级切换台tally设置</h4>
      <div class="collapse-icon">
        <VIcon icon="feather:chevron-down" />
      </div>
    </div>
    <div v-show="opened" class="form-fieldset-nested-4">
      <TallyScreenParams
        v-for="(screenParam, screenidx) in mv.screens.slice(0, showSecondScreen ? 2 : 1)"
        :key="`input-screen-${screenidx}`"
        v-model="mv.screens[screenidx]"
        :index="screenidx"
        :is-last="screenidx === mv.screens.length - 1"
      />
      <div v-if="!showSecondScreen" class="form-action-buttons">
        <div class="left">
          <VButton
            class="is-rounded"
            color="primary"
            raised
            @click="showSecondScreen = !showSecondScreen"
          >
            添加第2个画面
          </VButton>
        </div>
      </div>
    </div>
  </div>
</template>
<style lang="scss" scoped>
.form-action-buttons {
  display: flex;
  justify-content: center;
  align-items: center;
  margin: 32px 0;
}
</style>
