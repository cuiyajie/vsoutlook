<script lang="ts" setup>
import { TallyType } from "@src/utils/enums";
import { useColorPickerRef } from "./useColorPickerRef";

const props = defineProps<{
  modelValue: LayoutTitle,
  base: LayoutDimension
}>()

const emit = defineEmits<{
  (e: 'update:model-value', value: LayoutTitle): void,
}>()

const fz = computed({
  get: () => Math.round(props.modelValue.fontSize * props.base.w),
  set: (v: number) => emit('update:model-value', { ...props.modelValue, fontSize: v / props.base.w })
})

const fa = computed({
  get: () => props.modelValue.fontFamily,
  set: (v: string) => emit('update:model-value', { ...props.modelValue, fontFamily: v })
})

const color = computed({
  get: () => props.modelValue.color,
  set: (v: string) => emit('update:model-value', { ...props.modelValue, color: v })
})

const tx = computed({
  get: () => Math.round(props.modelValue.x * props.base.w),
  set: (v: number) => emit('update:model-value', { ...props.modelValue, x: v / props.base.w })
})

const ty = computed({
  get: () => Math.round(props.modelValue.y * props.base.h),
  set: (v: number) => emit('update:model-value', { ...props.modelValue, y: v / props.base.h })
})

const tw = computed({
  get: () => Math.round(props.modelValue.w * props.base.w),
  set: (v: number) => emit('update:model-value', { ...props.modelValue, w: v / props.base.w })
})

const th = computed({
  get: () => Math.round(props.modelValue.h * props.base.h),
  set: (v: number) => emit('update:model-value', { ...props.modelValue, h: v / props.base.h })
})

const vertical = computed({
  get: () => props.modelValue.isVertial,
  set: (v: boolean) => emit('update:model-value', { ...props.modelValue, isVertial: v })
})

const bgColor = computed({
  get: () => props.modelValue.bgColor,
  set: (v: string) => emit('update:model-value', { ...props.modelValue, bgColor: v })
})

const tallyType = computed({
  get: () => props.modelValue.tallyType,
  set: (v: TallyType) => emit('update:model-value', { ...props.modelValue, tallyType: v })
})

const tallyBgColor = computed({
  get: () => props.modelValue.tallyBgColor,
  set: (v: string) => emit('update:model-value', { ...props.modelValue, tallyBgColor: v })
})

const tallyBorder = computed({
  get: () => Math.round(props.modelValue.tallyBorderWidth * props.base.w),
  set: (v: number) => emit('update:model-value', { ...props.modelValue, tallyBorderWidth: v / props.base.w })
})

const pickerRef = useColorPickerRef()

</script>
<template>
  <div class="layout-form">
    <section class="layout-row">
      <div class="layout-row-inner">
        <div class="layout-cell">
          <div>字号</div>
          <input v-model="fz" type="number" min="0" :max="Math.floor(props.base.w / 5)">
        </div>
      </div>
    </section>
    <section class="layout-row">
      <div class="layout-row-inner">
        <div class="layout-cell fullwidth">
          <div>字体</div>
          <input v-model="fa" type="text">
        </div>
      </div>
    </section>
    <section class="layout-row">
      <div class="layout-row-inner">
        <div class="layout-cell">
          <div>字体颜色</div>
          <color-picker v-model:pureColor="color" format="rgb" shape="square" theme="black" picker-type="chrome" :picker-container="pickerRef" />
        </div>
      </div>
    </section>
    <section class="layout-row">
      <div>背景框样式</div>
      <div class="layout-row-inner">
        <div class="layout-cell">
          <VRadio v-model="vertical" :value="false" label="横向" name="vertical" color="primary" />
        </div>
        <div class="layout-cell">
          <VRadio v-model="vertical" :value="true" label="竖向" name="vertical" color="primary" />
        </div>
      </div>
    </section>
    <section class="layout-row">
      <div>背景位置</div>
      <div class="layout-row-inner">
        <div class="layout-cell half">
          <label for="tx">x</label>
          <input id="tx" v-model="tx" type="number" min="0" :max="base.w">
        </div>
        <div class="layout-cell half">
          <label for="ty">y</label>
          <input id="ty" v-model="ty" type="number" min="0" :max="base.h">
        </div>
      </div>
    </section>
    <section class="layout-row">
      <div>背景宽高</div>
      <div class="layout-row-inner">
        <div class="layout-cell half">
          <label for="tw">w</label>
          <input id="tw" v-model="tw" type="number" min="0" :max="base.w">
        </div>
        <div class="layout-cell half">
          <label for="th">h</label>
          <input id="th" v-model="th" type="number" min="0" :max="base.h">
        </div>
      </div>
    </section>
    <section class="layout-row">
      <div class="layout-row-inner">
        <div class="layout-cell">
          <div>背景框颜色</div>
          <color-picker v-model:pureColor="bgColor" format="rgb" shape="square" theme="black" picker-type="chrome" :picker-container="pickerRef" />
        </div>
      </div>
    </section>
    <section class="layout-row">
      <div>Tally方式</div>
      <div class="layout-row-inner">
        <div class="layout-cell">
          <VRadio v-model="tallyType" :value="TallyType.Text" label="文字" name="tallyType" color="primary" />
        </div>
        <div class="layout-cell">
          <VRadio v-model="tallyType" :value="TallyType.Border" label="边框" name="tallyType" color="primary" />
        </div>
        <div class="layout-cell">
          <VRadio v-model="tallyType" :value="TallyType.Light" label="Tally灯" name="tallyType" color="primary" />
        </div>
      </div>
    </section>
    <section v-if="tallyType === TallyType.Border" class="layout-row">
      <div class="layout-row-inner">
        <div class="layout-cell">
          <div>边框宽度</div>
          <input v-model="tallyBorder" type="number" min="0" :max="Math.floor(props.base.w / 5)">
        </div>
      </div>
    </section>
    <section v-if="tallyType === TallyType.Light" class="layout-row">
      <div class="layout-row-inner">
        <div class="layout-cell">
          <div>左/右灯背景色</div>
          <color-picker v-model:pureColor="tallyBgColor" format="rgb" shape="square" theme="black" picker-type="chrome" :picker-container="pickerRef" />
        </div>
      </div>
      <div class="layout-row-inner">
        <div class="layout-cell fullwidth">
          <div>tally灯宽度</div>
          <input id="tallyw" v-model="th" type="number" min="0" :max="base.h">
        </div>
      </div>
    </section>
  </div>
</template>
<style lang="scss">
.layout-panel {
}
</style>
