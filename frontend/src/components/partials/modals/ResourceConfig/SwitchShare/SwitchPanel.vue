<script lang="ts" setup>
import { type SwitchPanelParams, switch_panel_types, def_switch_panel_url } from './Consts'

const mv = defineModel<SwitchPanelParams>({
  default: {} as SwitchPanelParams,
  local: true,
})

defineProps<{
  index: number
  isLast: boolean
}>()

const opened = ref(false)

function addPanelUrl() {
  if (mv.value.panel_type === 'panel_c1') {
    mv.value.panel_url.push(def_switch_panel_url(mv.value.panel_url.length))
  }
}

function removePanelUrl(pidx: number) {
  if (mv.value.panel_type === 'panel_c1') {
    mv.value.panel_url.splice(pidx, 1)
  }
}

</script>
<template>
  <div
    class="form-fieldset collapse-form"
    :class="!isLast && 'seperator'"
    :open="opened || undefined"
  >
    <div
      class="fieldset-heading collapse-header"
      :style="opened ? { marginBottom: '20px !important' } : {}"
      tabindex="0"
      role="button"
      @keydown.space.prevent="opened = !opened"
      @click.prevent="opened = !opened"
    >
      <h4>第 {{ index + 1 }} 级面板设置</h4>
      <div class="collapse-icon">
        <VIcon icon="feather:chevron-down" />
      </div>
    </div>
    <div v-show="opened" class="form-fieldset-nested-4">
      <div class="columns is-multiline">
        <div class="column is-6">
          <VField>
            <VLabel>面板类型</VLabel>
            <VControl>
              <VSelect v-model="mv.panel_type" class="is-rounded">
                <VOption v-for="spt in switch_panel_types" :key="spt.value" :value="spt.value">{{ spt.label }}</VOption>
              </VSelect>
            </VControl>
          </VField>
        </div>
      </div>
      <div
        class="form-fieldset-nested-4 seperator-top"
        :style="{ paddingTop: (mv.panel_type === 'panel_c1') ? '12px' : '0' }"
      >
        <expand-transition>
          <div v-if="mv.panel_type === 'panel_c1'" class="fieldset-heading collapse-header" style="height: 38px; margin-bottom: 0 !important;">
            <span>&nbsp;</span>
            <VButton
              v-if="mv.panel_type === 'panel_c1'"
              class="is-rounded"
              color="primary"
              raised
              @click="addPanelUrl"
            >
              添加子面板
            </VButton>
          </div>
        </expand-transition>
        <div class="panel-url-container">
          <transition-group name="list">
            <SwitchPanelUrl
              v-for="(panelUrl, pidx) in mv.panel_url"
              :key="`panel-url-${pidx}`"
              v-model="mv.panel_url[pidx]"
              :index="pidx"
              :is-last="pidx === mv.panel_url.length - 1"
              :deletable="mv.panel_type === 'panel_c1' && pidx > 0"
              @remove="removePanelUrl(pidx)"
            />
          </transition-group>
        </div>
      </div>
    </div>
  </div>
</template>
<style lang="scss" scoped>
.panel-url-container {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  column-gap: 30px;
}
</style>
