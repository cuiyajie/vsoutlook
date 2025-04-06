<script lang="ts" setup>
import { defaultTimerColor } from '../utils'
import { clamp } from '@vueuse/core'
import { CLockDisplayType, DateDisplayType } from '@src/utils/enums'
import roundToNearestEven from '@src/utils/round-to-nearest-even'

const props = defineProps<{
  modelValue: LayoutDataItem
  bound: LayoutDimension
  base: LayoutDimension
}>()

const emit = defineEmits<{
  (e: 'update:model-value', value: LayoutDataItem): void
  (e: 'toggle-show-date', value: boolean): void
}>()

const tx = computed({
  get: () =>
    props.modelValue?.timer
      ? roundToNearestEven(props.modelValue.timer.x * props.base.w)
      : 0,
  set: (v: number) => {
    if (props.modelValue?.timer) {
      emit('update:model-value', {
        ...props.modelValue,
        timer: { ...props.modelValue.timer, x: v / props.base.w },
      })
    }
  },
})

const ty = computed({
  get: () =>
    props.modelValue?.timer
      ? roundToNearestEven(props.modelValue.timer.y * props.base.h)
      : 0,
  set: (v: number) => {
    if (props.modelValue?.timer) {
      emit('update:model-value', {
        ...props.modelValue,
        timer: { ...props.modelValue.timer, y: v / props.base.h },
      })
    }
  },
})

const fz = computed({
  get: () =>
    props.modelValue?.timer
      ? Math.round(props.modelValue.timer.fontSize * props.base.w)
      : 0,
  set: (v: number) => {
    if (props.modelValue?.timer) {
      const realV = clamp(v, 0, Math.floor(props.base.w / 5))
      emit('update:model-value', {
        ...props.modelValue,
        timer: { ...props.modelValue.timer, fontSize: realV / props.base.w },
      })
    }
  },
})

const fa = computed({
  get: () => props.modelValue?.timer?.fontFamily || '',
  set: (v: string) => {
    if (props.modelValue?.timer) {
      emit('update:model-value', {
        ...props.modelValue,
        timer: { ...props.modelValue.timer, fontFamily: v },
      })
    }
  },
})

const color = computed({
  get: () => props.modelValue?.timer?.color || defaultTimerColor,
  set: (v: string) => {
    if (props.modelValue?.timer) {
      emit('update:model-value', {
        ...props.modelValue,
        timer: { ...props.modelValue.timer, color: v },
      })
    }
  },
})

const showdate = computed({
  get: () => Boolean(props.modelValue?.timer?.showDate),
  set: (v: boolean) => {
    if (props.modelValue?.timer) {
      emit('update:model-value', {
        ...props.modelValue,
        timer: { ...props.modelValue.timer, showDate: v },
      })
      if (v !== props.modelValue.timer.showDate) {
        emit('toggle-show-date', v)
      }
    }
  },
})

const clockDisplayType = computed({
  get: () =>
    props.modelValue?.timer?.clockDisplayType
      ? CLockDisplayType.LED
      : CLockDisplayType.Digital,
  set: (v: number) => {
    if (props.modelValue?.timer) {
      const related =
        v === CLockDisplayType.LED ? { dateDisplayType: DateDisplayType.Hyphen } : {}
      emit('update:model-value', {
        ...props.modelValue,
        timer: { ...props.modelValue.timer, clockDisplayType: v, ...related },
      })
    }
  },
})

const datetype = computed<DateDisplayType>({
  get: () =>
    props.modelValue?.timer?.dateDisplayType
      ? DateDisplayType.Hyphen
      : DateDisplayType.Chinese,
  set: (v: number) => {
    if (props.modelValue?.timer) {
      emit('update:model-value', {
        ...props.modelValue,
        timer: { ...props.modelValue.timer, dateDisplayType: v },
      })
    }
  },
})

const time24 = computed({
  get: () => props.modelValue?.timer?.time24 || 24,
  set: (v: number) => {
    if (props.modelValue?.timer) {
      emit('update:model-value', {
        ...props.modelValue,
        timer: { ...props.modelValue.timer, time24: v },
      })
    }
  },
})

const timeX = computed({
  get: () =>
    props.modelValue?.timer?.timePosition
      ? roundToNearestEven(props.modelValue.timer.timePosition.x * props.base.w)
      : 0,
  set: (v: number) => {
    if (props.modelValue?.timer) {
      const timePos = props.modelValue.timer.timePosition || { x: 0, y: 0 }
      emit('update:model-value', {
        ...props.modelValue,
        timer: {
          ...props.modelValue.timer,
          timePosition: {
            ...timePos,
            x: v / props.base.w,
          },
        },
      })
    }
  },
})

const timeY = computed({
  get: () =>
    props.modelValue?.timer?.timePosition
      ? roundToNearestEven(props.modelValue.timer.timePosition.y * props.base.h)
      : 0,
  set: (v: number) => {
    if (props.modelValue?.timer) {
      const timePos = props.modelValue.timer.timePosition || { x: 0, y: 0 }
      emit('update:model-value', {
        ...props.modelValue,
        timer: {
          ...props.modelValue.timer,
          timePosition: {
            ...timePos,
            y: v / props.base.h,
          },
        },
      })
    }
  },
})

const fh = computed({
  get: () =>
    props.modelValue?.timer
      ? Math.round(props.modelValue.timer.fontHeight * props.base.w)
      : 0,
  set: (v: number) => {
    if (props.modelValue?.timer) {
      const realV = clamp(v, 0, Math.floor(props.base.w / 5))
      emit('update:model-value', {
        ...props.modelValue,
        timer: { ...props.modelValue.timer, fontHeight: realV / props.base.w },
      })
    }
  },
})

const fg = computed({
  get: () =>
    props.modelValue?.timer
      ? Math.round(props.modelValue.timer.fontGap * props.base.w)
      : 0,
  set: (v: number) => {
    if (props.modelValue?.timer) {
      const realV = clamp(v, 0, Math.floor(props.base.w / 5))
      emit('update:model-value', {
        ...props.modelValue,
        timer: { ...props.modelValue.timer, fontGap: realV / props.base.w },
      })
    }
  },
})

const tcolor = computed({
  get: () => props.modelValue?.timer?.timeColor || defaultTimerColor,
  set: (v: string) => {
    if (props.modelValue?.timer) {
      emit('update:model-value', {
        ...props.modelValue,
        timer: { ...props.modelValue.timer, timeColor: v },
      })
    }
  },
})
</script>
<template>
  <section class="layout-row">
    <div>时钟样式</div>
    <div class="layout-row-inner">
      <div class="layout-cell">
        <VRadio
          v-model="clockDisplayType"
          :value="CLockDisplayType.Digital"
          label="数字时钟"
          name="newline"
          color="primary"
        />
      </div>
      <div class="layout-cell">
        <VRadio
          v-model="clockDisplayType"
          :value="CLockDisplayType.LED"
          label="LED 时钟"
          name="newline"
          color="primary"
        />
      </div>
    </div>
  </section>
  <section class="layout-row">
    <div>制式</div>
    <div class="layout-row-inner">
      <div class="layout-cell">
        <VRadio
          v-model="time24"
          :value="24"
          label="24小时制"
          name="time24"
          color="primary"
        />
      </div>
      <div class="layout-cell">
        <VRadio
          v-model="time24"
          :value="12"
          label="12小时制"
          name="time24"
          color="primary"
        />
      </div>
    </div>
  </section>
  <section class="layout-row">
    <div>日期样式</div>
    <div class="layout-row-inner">
      <div v-if="clockDisplayType === CLockDisplayType.Digital" class="layout-cell">
        <VRadio
          v-model="datetype"
          :value="0"
          label="YYYY年MM月DD日"
          name="datetype"
          color="primary"
        />
      </div>
      <div class="layout-cell">
        <VRadio
          v-model="datetype"
          :value="1"
          label="YYYY-MM-DD"
          name="datetype"
          color="primary"
        />
      </div>
    </div>
  </section>
  <template v-if="clockDisplayType === CLockDisplayType.Digital">
    <section class="layout-row">
      <div class="layout-row-inner">
        <div class="layout-cell">
          <div>字体大小</div>
          <input v-model="fz" type="number" min="0" :max="Math.floor(props.base.w / 5)" />
        </div>
      </div>
    </section>
    <section class="layout-row">
      <div class="layout-row-inner">
        <div class="layout-cell fullwidth is-select">
          <div class="select-label">字体</div>
          <font-select v-model="fa" />
        </div>
      </div>
    </section>
    <section class="layout-row">
      <div class="layout-row-inner">
        <div class="layout-cell">
          <div>颜色</div>
          <color-picker
            v-model:pureColor="color"
            format="rgb"
            shape="square"
            theme="black"
            picker-type="chrome"
          />
        </div>
      </div>
    </section>
    <section class="layout-row">
      <div>时钟位置</div>
      <div class="layout-row-inner">
        <div class="layout-cell half">
          <label for="tx">x</label>
          <input
            id="timeWinx"
            v-model="timeX"
            type="number"
            min="0"
            :max="base.w"
            step="2"
          />
        </div>
        <div class="layout-cell half">
          <label for="ty">y</label>
          <input
            id="timeWiny"
            v-model="timeY"
            type="number"
            min="0"
            :max="base.h"
            step="2"
          />
        </div>
      </div>
    </section>
    <section class="layout-row">
      <div class="layout-row-inner">
        <div class="layout-cell">
          <div>是否显示日期</div>
          <VCheckbox v-model="showdate" color="primary" circle />
        </div>
      </div>
    </section>
  </template>
  <template v-else>
    <section class="layout-row">
      <div class="layout-row-inner">
        <div class="layout-cell">
          <div>字体高度</div>
          <input v-model="fh" type="number" min="0" :max="Math.floor(props.base.w / 5)" />
        </div>
      </div>
    </section>
    <section class="layout-row">
      <div class="layout-row-inner">
        <div class="layout-cell">
          <div>数字间距</div>
          <input v-model="fg" type="number" min="0" :max="Math.floor(props.base.w / 5)" />
        </div>
      </div>
    </section>
    <section class="layout-row">
      <div class="layout-row-inner">
        <div class="layout-cell">
          <div>颜色</div>
          <color-picker
            v-model:pureColor="tcolor"
            format="rgb"
            shape="square"
            theme="black"
            picker-type="chrome"
          />
        </div>
      </div>
    </section>
    <section class="layout-row">
      <div>时钟位置</div>
      <div class="layout-row-inner">
        <div class="layout-cell half">
          <label for="tx">x</label>
          <input
            id="timeWinx"
            v-model="timeX"
            type="number"
            min="0"
            :max="base.w"
            step="2"
          />
        </div>
        <div class="layout-cell half">
          <label for="ty">y</label>
          <input
            id="timeWiny"
            v-model="timeY"
            type="number"
            min="0"
            :max="base.h"
            step="2"
          />
        </div>
      </div>
    </section>
    <section class="layout-row">
      <div class="layout-row-inner">
        <div class="layout-cell">
          <div>是否显示日期</div>
          <VCheckbox v-model="showdate" color="primary" circle />
        </div>
      </div>
    </section>
  </template>
  <section class="layout-row">
    <div>日期位置</div>
    <div class="layout-row-inner">
      <div class="layout-cell half">
        <label for="tx">x</label>
        <input id="winx" v-model="tx" type="number" min="0" :max="base.w" step="2" />
      </div>
      <div class="layout-cell half">
        <label for="ty">y</label>
        <input id="winy" v-model="ty" type="number" min="0" :max="base.h" step="2" />
      </div>
    </div>
  </section>
</template>
<style lang="scss"></style>
