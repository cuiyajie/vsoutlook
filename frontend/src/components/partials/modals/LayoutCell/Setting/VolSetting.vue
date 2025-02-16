<script lang="ts" setup>

const props = defineProps<{
  modelValue: LayoutVol,
  base: LayoutDimension
}>()

const emit = defineEmits<{
  (e: 'update:model-value', value: LayoutVol): void,
}>()

const vols = computed<LayoutSingleVol[]>({
  get: () => props.modelValue.vols,
  set: v => emit('update:model-value', { ...props.modelValue, vols: v})
})

const len = computed({
  get: () => props.modelValue.len,
  set: (v: number) => emit('update:model-value', { ...props.modelValue, len: v })
})

const alpha = computed({
  get: () => props.modelValue.alpha,
  set: (v: number) => emit('update:model-value', { ...props.modelValue, alpha: v })
})

const updateVol = (idx: number, vol: LayoutSingleVol) => {
  vols.value = vols.value.map((v, i) => i === idx ? vol : v)
}

</script>
<template>
  <div class="layout-form">
    <section class="layout-row">
      <div>音柱显示区域</div>
      <div class="layout-row-inner">
        <div v-for="(_, idx) in new Array(4).fill(0)" :key="idx" class="layout-cell">
          <VRadio v-model="len" :value="idx + 1" :label="String(idx + 1)" name="len" color="primary" />
        </div>
      </div>
    </section>
    <section class="layout-row">
      <div class="layout-row-inner">
        <div class="layout-cell fullwidth">
          <div>音柱透明度</div>
          <input v-model="alpha" type="number" min="0" :max="1" :step="0.1">
        </div>
      </div>
    </section>
    <single-vol-setting v-for="(vol, idx) in vols.slice(0, len)" :key="idx" :model-value="vol" :base="base" @update:model-value="updateVol(idx, $event)">
      <div class="section-header">窗口 {{ idx + 1 }}</div>
    </single-vol-setting>
  </div>
</template>
<style lang="scss">
.layout-panel {
}
</style>
