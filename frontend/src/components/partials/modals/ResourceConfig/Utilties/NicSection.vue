<!-- eslint-disable vue/prop-name-casing -->
<script setup lang="ts">
import { useNotyf } from "@src/composable/useNotyf";
import { type NicDetail, def_nic_detail } from './Consts_V1'

const notyf = useNotyf();

const props = defineProps<{
  nics: NicInfo[]
}>()

const mv = defineModel<NicDetail[]>({
  default: [],
  local: true,
})

function add() {
  if (mv.value.length >= props.nics.length) {
    notyf.error('网卡数量已达上限')
    return
  }
  mv.value.push(def_nic_detail())
}

function remove(idx: number) {
  mv.value.splice(idx, 1)
}

function isExist(nicId: string) {
  return mv.value.some((nd) => nd.id === nicId)
}
</script>
<template>
  <div class="form-outer has-mt-20">
    <div class="form-header">
      <div class="form-header-inner">
        <div class="left">
          <h4>收发流网卡列表</h4>
        </div>
        <button
          class="button is-circle is-dark-outlined"
          @keydown.space.prevent="add"
          @click.prevent="add"
        >
          <span class="icon is-small">
            <i aria-hidden="true" class="iconify" data-icon="feather:plus" />
          </span>
        </button>
      </div>
    </div>
    <div class="form-body">
      <TransitionGroup name="list">
        <!--Fieldset-->
        <NicSelect
          v-for="(nd, nidx) in mv"
          :key="nidx"
          v-model="mv[nidx]"
          :index="nidx"
          :is-last="nidx === mv.length - 1"
          :nics="nics"
          :is-exist="isExist"
          @remove="remove(nidx)"
        />
      </TransitionGroup>
    </div>
    <div v-if="mv.length === 0" class="is-empty-list">暂无网卡设置</div>
  </div>
</template>ss
