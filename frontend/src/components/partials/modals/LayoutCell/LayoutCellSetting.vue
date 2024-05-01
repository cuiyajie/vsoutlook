<script lang="ts" setup>
import { ColorPicker } from "vue3-colorpicker";
import "vue3-colorpicker/style.css";

const props = defineProps<{
  modelValue: LayoutDataItem,
  type: LayoutProps | null,
  base: LayoutDimension
}>()

const emit = defineEmits<{
  (e: 'reset'): void,
  (e: 'update:model-value', value: LayoutDataItem): void
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

const actl = computed({
  get: () => ac.value && 'text' in ac.value ? ac.value.text : '',
  set: (v: string) => {
    if (ac.value) ac.value = { ...ac.value, text: v }
  }
})

const acolor = computed({
  get: () => ac.value && 'color' in ac.value ? ac.value.color : '',
  set: (v: string) => {
    if (ac.value) ac.value = { ...ac.value, color: v }
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
      <a
        v-if="ac"
        role="button"
        class="action-link"
        tabindex="0"
        @keydown.space.prevent="emit('reset')"
        @click.prevent="emit('reset')"
      >重置</a>
    </div>
    <div v-if="ac" class="layout-form">
      <section class="layout-row">
        <div>尺寸</div>
        <div class="layout-row-inner">
          <div v-if="type !== 'vol'" class="layout-cell half">
            <label for="acw">w</label>
            <input v-model="acw" name="acw" type="number" min="0" :max="base.w">
          </div>
          <div class="layout-cell half">
            <label for="ach">h</label>
            <input v-model="ach" name="ach" type="number" min="0" :max="base.h">
          </div>
        </div>
      </section>
      <section class="layout-row">
        <div>位置</div>
        <div class="layout-row-inner">
          <div class="layout-cell half">
            <label for="acx">X</label>
            <input v-model="acx" name="acx" type="number" min="0" :max="base.w">
          </div>
          <div class="layout-cell half">
            <label for="acy">Y</label>
            <input v-model="acy" name="acy" type="number" min="0" :max="base.h">
          </div>
        </div>
      </section>
      <section v-if="type === 'title' || type === 'timer'" class="layout-row">
        <div class="layout-row-inner">
          <div class="layout-cell">
            <div>字体</div>
            <input v-model="acfz" type="number" min="0">
          </div>
        </div>
      </section>
      <section v-if="type === 'title'" class="layout-row">
        <div class="layout-row-inner">
          <div class="layout-cell">
            <div>内容</div>
            <input v-model="actl" type="text" style="width: auto;">
          </div>
        </div>
      </section>
      <section v-if="type === 'timer'" class="layout-row">
        <div class="layout-row-inner">
          <div class="layout-cell">
            <div>颜色</div>
            <color-picker v-model:pureColor="acolor" format="rgb" shape="square" theme="black" picker-type="chrome" />
          </div>
        </div>
      </section>
      <section v-if="type === 'vol'" class="layout-row">
        <div>音柱</div>
        <div class="layout-row-inner">
          <div class="layout-cell half">
            <label for="acow">柱宽</label>
            <input v-model="acow" name="acow" type="number" min="0" :max="base.w">
          </div>
          <div class="layout-cell half">
            <label for="acgp">间距</label>
            <input v-model="acgp" name="acgp" type="number" min="0" :max="base.w">
          </div>
          <div class="layout-cell half">
            <label for="acnt">数量</label>
            <input v-model="acnt" name="acnt" type="number" min="0">
          </div>
        </div>
      </section>
    </div>
  </div>
</template>
<style lang="scss">
.layout-panel {
}

.vc-display .vc-color-input input,
.vc-display .vc-alpha-input__inner {
  color: #fff !important;
}

.vc-display .vc-input-toggle:hover {
  background-color: transparent !important;
}
</style>
