<script lang="ts" setup>
import { type SwitchPanelUrl } from './Consts'

const mv = defineModel<SwitchPanelUrl>({
  default: {} as SwitchPanelUrl,
  local: true,
})

defineProps<{
  index: number
  isLast: boolean
  deletable: boolean
}>()

const emit = defineEmits<{
  (e: 'remove'): void
}>()

</script>
<template>
  <div class="form-fieldset-nested-5 collapse-form">
    <div class="fieldset-heading collapse-header" style="height: 30px; margin-bottom: 16px !important;">
      <h5>第{{ index + 1 }}个子面板通讯地址</h5>
      <div class="collapse-icons">
        <div
          v-if="deletable"
          class="collapse-icon is-close-hidden"
          role="button"
          tabindex="0"
          @click.prevent.stop="emit('remove')"
          @keydown.space.prevent.stop="emit('remove')"
        >
          <VIcon icon="feather:x" />
        </div>
      </div>
    </div>
    <div class="form-fieldset-nested-4">
      <div class="columns is-multiline">
        <div class="column is-12">
          <VField>
            <VControl>
              <VInput v-model="mv.url" />
            </VControl>
          </VField>
        </div>
      </div>
    </div>
  </div>
</template>

