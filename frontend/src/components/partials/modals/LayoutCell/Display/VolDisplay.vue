<script lang="ts" setup>
const props = defineProps<{
  modelValue: LayoutDataItem,
  bound: LayoutDimension
}>()

const emit = defineEmits<{
  (e: 'update:model-value', value: LayoutDataItem): void,
  (e: 'select', index: number): void
}>()

const mv = computed({
  get: () => props.modelValue,
  set: (v: LayoutDataItem) => emit('update:model-value', v)
})

const ratiow = computed(() => props.bound.w / mv.value.win.w)
const ratioh = computed(() => props.bound.h / mv.value.win.h)

</script>
<template>
  <div class="layout-vols">
    <div
      v-for="(vol, vi) in mv.vol!.vols.slice(0, mv.vol!.len)"
      :key="vi"
      class="layout-vol"
      :style="{left: `${vol.x * ratiow}px`, top: `${vol.y * ratioh}px`, height: `${vol.h * ratioh}px`, gap: `${vol.g * ratiow}px`, opacity: mv.vol!.alpha}"
      role="button"
      tabindex="0"
      @click.prevent="emit('select', vi)"
      @keyup.enter.prevent="emit('select', vi)"
    >
      <div v-for="(_, svi) in new Array(vol.end - vol.start + 1).fill(0)" :key="svi" class="vol" :style="{width: `${vol.one_w * ratiow}px`}">{{ vol.start + svi }}</div>
    </div>
  </div>
</template>
<style lang="scss">
</style>
