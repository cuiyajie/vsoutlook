<!-- eslint-disable vue/prop-name-casing -->
<script setup lang="ts">
import { useNotyf } from "@src/composable/useNotyf";
import { type NicDetail } from './Consts_V1'

const notyf = useNotyf();

const props = defineProps<{
  index: number
  isLast: boolean
  nics: NicInfo[],
  isExist: (nicId: string) => boolean
}>()

const emit = defineEmits<{
  (e: 'remove'): void
}>()

const mv = defineModel<NicDetail>({
  default: {} as NicDetail,
  local: true,
})

const nicDisplays = computed(() => {
  return props.nics.map((nic, nidx) => ({
    label: `网卡 ${nidx}`,
    ...nic
  }))
})

const selectRef = ref<any>(null)

function onNicSelect(nicId: string) {
  if (nicId === mv.value.id) return
  const nidx = nicDisplays.value.findIndex(n => n.id === nicId)
  if (nidx !== -1) {
    const nic = nicDisplays.value[nidx]
    if (props.isExist(nicId)) {
      notyf.error(`网卡 ${nidx} 已配置`)
      if (selectRef.value) {
        if (mv.value.nicIndex !== -1) {
          selectRef.value.update(selectRef.value.getOption(mv.value.id))
        } else {
          selectRef.value.clear()
        }
      }
      return
    }
    Object.assign(mv.value, {
      nic_name_m: nic.nicNameMain,
      nic_name_b: nic.nicNameBackup,
      '2110-7_m_receive': nic.receiveMain,
      '2110-7_m_send': nic.sendMain,
      '2110-7_b_receive': nic.receiveBackup,
      '2110-7_b_send': nic.sendBackup,
      nicIndex: nidx,
      id: nicId,
    })
  } else {
    Object.assign(mv.value, {
      nic_name_m: '',
      nic_name_b: '',
      '2110-7_m_receive': false,
      '2110-7_m_send': false,
      '2110-7_b_receive': false,
      '2110-7_b_send': false,
      nicIndex: -1,
      id: '',
    })
  }
}

watch(() => mv.value.id, (nv) => {
  selectRef.value?.select(nv)
}, { immediate: true })

const selected = computed(() => mv.value.nicIndex !== -1)

onMounted(() => {
  if (mv.value.id) {
    selectRef.value?.select(mv.value.id)
  }
})

</script>
<template>
  <div
    class="form-fieldset collapse-form"
    :class="!isLast && 'seperator'"
    open
  >
    <div
      class="fieldset-heading collapse-header"
      tabindex="0"
      role="button"
    >
      <h4>第 {{ index + 1 }} 块网卡配置</h4>
      <div class="collapse-icon">
        <VIconButton
          color="warning"
          light
          raised
          circle
          icon="feather:x"
          @click="emit('remove')"
        />
      </div>
    </div>
    <div class="form-fieldset-nested-4">
      <div class="columns is-multiline">
        <div class="column is-8">
          <VField>
            <VLabel>选择网卡</VLabel>
            <VControl>
              <Multiselect
                ref="selectRef"
                class="tippy-select"
                placeholder="选择网卡"
                no-options-text="该节点没有配置网卡"
                value-prop="id"
                label="label"
                :searchable="false"
                :can-deselect="false"
                :can-clear="false"
                :style="{'--ms-max-height': '245px'}"
                :options="nicDisplays"
                @select="onNicSelect"
              >
                <template #singlelabel="{ value }">
                  <div class="multiselect-single-label">
                    <span class="select-label-text">
                      {{ value.label }}
                    </span>
                  </div>
                </template>
                <template #option="{ option }">
                  <span class="select-option-text">
                    {{ option.label }}
                  </span>
                  <NicToolTip :nic="option">
                    <i class="iconify" data-icon="feather:alert-circle" aria-hidden="true"></i>
                  </NicToolTip>
                </template>
              </Multiselect>
            </VControl>
          </VField>
        </div>
        <div class="column is-4" />
        <expand-transition>
          <div v-show="selected" class="column is-4">
            <VField>
              <VLabel>主路收发网口IP</VLabel>
              <VControl>
                <VInput v-model="mv['2110-7_m_local_ip']" />
              </VControl>
            </VField>
          </div>
        </expand-transition>
        <expand-transition>
          <div v-show="selected" class="column is-4">
            <VField>
              <VLabel>主路收发网卡名称</VLabel>
              <VControl>
                <VInput v-model="mv.nic_name_m" readonly />
              </VControl>
            </VField>
          </div>
        </expand-transition>
        <expand-transition>
          <div v-show="selected" class="column is-4">
            <VField label="主路收发网卡开关" class="switch-2-group">
              <VControl subcontrol class="switch-2">
                <VSwitchBlock v-model="mv['2110-7_m_receive']" disabled color="info" label="收" />
              </VControl>
              <VControl subcontrol class="switch-2">
                <VSwitchBlock v-model="mv['2110-7_m_send']" disabled color="danger" label="发" />
              </VControl>
            </VField>
          </div>
        </expand-transition>
        <expand-transition>
          <div v-show="selected" class="column is-4">
            <VField>
              <VLabel>备路收发网口IP</VLabel>
              <VControl>
                <VInput v-model="mv['2110-7_b_local_ip']" />
              </VControl>
            </VField>
          </div>
        </expand-transition>
        <expand-transition>
          <div v-show="selected" class="column is-4">
            <VField>
              <VLabel>备路收发网卡名称</VLabel>
              <VControl>
                <VInput v-model="mv.nic_name_b" readonly />
              </VControl>
            </VField>
          </div>
        </expand-transition>
        <expand-transition>
          <div v-show="selected" class="column is-4">
            <VField label="备路收发网卡开关" class="switch-2-group">
              <VControl subcontrol class="switch-2">
                <VSwitchBlock v-model="mv['2110-7_b_receive']" disabled color="info" label="收" />
              </VControl>
              <VControl subcontrol class="switch-2">
                <VSwitchBlock v-model="mv['2110-7_b_send']" disabled color="danger" label="发" />
              </VControl>
            </VField>
          </div>
        </expand-transition>
      </div>
    </div>
  </div>
</template>
<style lang="scss">
@import '@src/scss/abstracts/all';
@import '@src/scss/components/forms-outer';
</style>

