<!-- eslint-disable vue/prop-name-casing -->
<script setup lang="ts">
import { useUserSession } from "@src/stores/userSession";

const usStore = useUserSession()
const settings = computed(() => usStore.settings)

const font = defineModel({
  default: '',
  local: true,
})

defineProps<{
  openDirection?: string
}>()

</script>
<template>
  <Multiselect
    v-model="font"
    class="font-select"
    placeholder="选择字体"
    no-options-text="未设置字体，请在系统设置中设置"
    :searchable="false"
    :open-direction="openDirection"
    :can-deselect="false"
    :can-clear="false"
    :style="{'--ms-max-height': '90px'}"
    :options="settings.mv_template_fonts"
  >
    <template #singlelabel="{ value }">
      <div class="multiselect-single-label">
        <span class="select-label-text">
          <span class="option-item">
            <i class="lnir lnir-text" aria-hidden="true"></i>
            <span :style="{fontFamily: value.value}">{{ value.label }}</span>
          </span>
        </span>
      </div>
    </template>
    <template #option="{ option }">
      <span class="select-option-text">
        <span class="option-item">
          <span :style="{fontFamily: option.value}">{{ option.label }}</span>
        </span>
      </span>
    </template>
  </Multiselect>
</template>
<style lang="scss">
.multiselect.font-select {
  .multiselect-single-label {
    padding-left: 8px;
    padding-right: 8px;

    .option-item i {
      margin-right: 6px;
    }
  }
}
</style>
