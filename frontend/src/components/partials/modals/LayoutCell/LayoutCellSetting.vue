<script lang="ts" setup>

const props = defineProps<{
  modelValue: LayoutDataItem,
  type: LayoutProps | null,
  base: LayoutDimension
}>()

const emit = defineEmits<{
  (e: 'reset'): void,
  (e: 'delete'): void,
  (e: 'update:model-value', value: LayoutDataItem): void,
}>()

const ac = computed<LayoutTitle | LayoutVol | LayoutTimer | null>({
  get: () => props.type ? props.modelValue[props.type] : null,
  set: (v: LayoutDataItem[LayoutProps]) => props.type && emit('update:model-value', { ...props.modelValue, [props.type]: v })
})

const acx = computed({
  get: () => ac.value ? Math.round(ac.value.x * props.base.w) : 0,
  set: (v: number) => {
    if (ac.value) ac.value = { ...ac.value, x: v / props.base.w }
  }
})

const acy = computed({
  get: () => ac.value ? Math.round(ac.value.y * props.base.h) : 0,
  set: (v: number) => {
    if (ac.value) ac.value = { ...ac.value, y: v / props.base.h }
  }
})

const acw = computed({
  get: () => ac.value && 'w' in ac.value ? Math.round(ac.value.w * props.base.w) : 0,
  set: (v: number) => {
    if (ac.value) ac.value = { ...ac.value, w: v / props.base.w }
  }
})

const ach = computed({
  get: () => ac.value ? Math.round(ac.value.h * props.base.h) : 0,
  set: (v: number) => {
    if (ac.value) ac.value = { ...ac.value, h: v / props.base.h }
  }
})

const acfz = computed({
  get: () => ac.value && 'fontSize' in ac.value ? Math.round(ac.value.fontSize * props.base.w) : 0,
  set: (v: number) => {
    if (ac.value) ac.value = { ...ac.value, fontSize: v / props.base.w }
  }
})

const acow = computed({
  get: () => ac.value && 'one_w' in ac.value ? Math.round(ac.value.one_w * props.base.w) : 0,
  set: (v: number) => {
    if (ac.value) ac.value = { ...ac.value, one_w: v / props.base.w }
  }
})

const acgp = computed({
  get: () => ac.value && 'g' in ac.value ? Math.round(ac.value.g * props.base.w) : 0,
  set: (v: number) => {
    if (ac.value) ac.value = { ...ac.value, g: v / props.base.w }
  }
})

const acnt = computed({
  get: () => ac.value && 'cnt' in ac.value ? ac.value.cnt : 0,
  set: (v: number) => {
    if (ac.value) ac.value = { ...ac.value, cnt: v }
  }
})

</script>
<template>
  <div className="layout-panel">
    <div class="layout-header mb-4">
      <h3 class="title is-5">属性设置</h3>
      <div v-if="ac">
        <a
          v-if="type !== 'timer'"
          role="button"
          class="action-link is-warning mr-4"
          tabindex="0"
          @keydown.space.prevent="emit('delete')"
          @click.prevent="emit('delete')"
        >删除</a>
        <a
          role="button"
          class="action-link"
          tabindex="0"
          @keydown.space.prevent="emit('reset')"
          @click.prevent="emit('reset')"
        >重置</a>
      </div>
    </div>
    <div v-if="ac" class="layout-form">
      <section class="layout-row">
        <div>尺寸</div>
        <div class="layout-row-inner">
          <div v-if="type !== 'vol'" class="layout-cell half">
            <label for="acw">w</label>
            <input id="acw" v-model="acw" type="number" min="0" :max="base.w">
          </div>
          <div class="layout-cell half">
            <label for="ach">h</label>
            <input id="ach" v-model="ach" type="number" min="0" :max="base.h">
          </div>
        </div>
      </section>
      <section class="layout-row">
        <div>位置</div>
        <div class="layout-row-inner">
          <div class="layout-cell half">
            <label for="acx">X</label>
            <input id="acx" v-model="acx" type="number" min="0" :max="base.w">
          </div>
          <div class="layout-cell half">
            <label for="acy">Y</label>
            <input id="acy" v-model="acy" type="number" min="0" :max="base.h">
          </div>
        </div>
      </section>
      <section v-if="type === 'title'" class="layout-row">
        <div class="layout-row-inner">
          <div class="layout-cell">
            <div>字体</div>
            <input v-model="acfz" type="number" min="0">
          </div>
        </div>
      </section>
      <section v-if="type === 'vol'" class="layout-row">
        <div>音柱</div>
        <div class="layout-row-inner">
          <div class="layout-cell half">
            <label for="acow">柱宽</label>
            <input id="acow" v-model="acow" type="number" min="0" :max="base.w">
          </div>
          <div class="layout-cell half">
            <label for="acgp">间距</label>
            <input id="acgp" v-model="acgp" type="number" min="0" :max="base.w">
          </div>
          <div class="layout-cell half">
            <label for="acnt">数量</label>
            <input id="acnt" v-model="acnt" type="number" min="0">
          </div>
        </div>
      </section>
    </div>
  </div>
</template>
<style lang="scss">
.layout-panel {
}
</style>
