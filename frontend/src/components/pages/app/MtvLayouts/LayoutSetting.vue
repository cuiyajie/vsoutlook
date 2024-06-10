<script lang="ts" setup>
import { ColorPicker } from "vue3-colorpicker";
import "vue3-colorpicker/style.css";
import { draftVol, draftTitle, defaultTimerColor } from './utils';
import { clamp } from "@vueuse/core";

const props = defineProps<{
  modelValue: LayoutDataItem,
  bound: LayoutDimension,
  base: LayoutDimension
}>()

const emit = defineEmits<{
  (e: 'reset'): void,
  (e: 'update:model-value', value: LayoutDataItem): void
}>()

const winw = computed({
  get: () => Math.round(props.modelValue.win.w * props.base.w),
  set: (v: number) => emit('update:model-value', { ...props.modelValue, win: { ...props.modelValue.win, w: v / props.base.w } })
})

const winh = computed({
  get: () => Math.round(props.modelValue.win.h * props.base.h),
  set: (v: number) => emit('update:model-value', { ...props.modelValue, win: { ...props.modelValue.win, h: v / props.base.h } })
})

const winx = computed({
  get: () => Math.round(props.modelValue.win.x * props.base.w),
  set: (v: number) => emit('update:model-value', { ...props.modelValue, win: { ...props.modelValue.win, x: v / props.base.w } })
})

const winy = computed({
  get: () => Math.round(props.modelValue.win.y * props.base.h),
  set: (v: number) => emit('update:model-value', { ...props.modelValue, win: { ...props.modelValue.win, y: v / props.base.h } })
})

const reservedTitle = ref<LayoutDataItem['title']>(null)
const title = computed({
  get: () => Boolean(props.modelValue.title),
  set: (v: boolean) => {
    if (v) {
      reservedTitle.value = reservedTitle.value || draftTitle(props.modelValue.win.w * props.bound.w, props.modelValue.win.h * props.bound.h, props.bound)
    } else {
      reservedTitle.value = props.modelValue.title
    }
    emit('update:model-value', { ...props.modelValue, title: v ? reservedTitle.value : null})
  }
})

const reservedVol = ref<LayoutDataItem['vol']>(null)
const vol = computed({
  get: () => Boolean(props.modelValue.vol),
  set: (v: boolean) => {
    if (v) {
      reservedVol.value = reservedVol.value || draftVol(props.modelValue.win.h * props.bound.h, props.bound)
    } else {
      reservedVol.value = props.modelValue.vol
    }
    emit('update:model-value', { ...props.modelValue, vol: v ? reservedVol.value : null })
  }
})

const isTimer = computed(() => Boolean(props.modelValue.timer))

const tx = computed({
  get: () => props.modelValue?.timer ? Math.round(props.modelValue.timer.x * props.base.w) : 0,
  set: (v: number) => {
    if (props.modelValue?.timer) {
      emit('update:model-value', { ...props.modelValue, timer: { ...props.modelValue.timer, x: v / props.base.w } })
    }
  }
})

const ty = computed({
  get: () => props.modelValue?.timer ? Math.round(props.modelValue.timer.y * props.base.h) : 0,
  set: (v: number) => {
    if (props.modelValue?.timer) {
      emit('update:model-value', { ...props.modelValue, timer: { ...props.modelValue.timer, y: v / props.base.h } })
    }
  }
})

const fz = computed({
  get: () => props.modelValue?.timer ? Math.round(props.modelValue.timer.fontSize * props.base.w) : 0,
  set: (v: number) => {
    if (props.modelValue?.timer) {
      const realV = clamp(v, 0, Math.floor(props.base.w / 5))
      emit('update:model-value', { ...props.modelValue, timer: { ...props.modelValue.timer, fontSize: realV / props.base.w } })
    }
  }
})

const fa = computed({
  get: () => props.modelValue?.timer?.fontFamily || '',
  set: (v: string) => {
    if (props.modelValue?.timer) {
      emit('update:model-value', { ...props.modelValue, timer: { ...props.modelValue.timer, fontFamily: v } })
    }
  }
})

const color = computed({
  get: () => props.modelValue?.timer?.color || defaultTimerColor,
  set: (v: string) => {
    if (props.modelValue?.timer) {
      emit('update:model-value', { ...props.modelValue, timer: { ...props.modelValue.timer, color: v } })
    }
  }
})

const showdate = computed({
  get: () => Boolean(props.modelValue?.timer?.showDate),
  set: (v: boolean) => {
    if (props.modelValue?.timer) {
      emit('update:model-value', { ...props.modelValue, timer: { ...props.modelValue.timer, showDate: v } })
    }
  }
})

const newline = computed({
  get: () => props.modelValue?.timer?.dateNewLine ? 1 : 0,
  set: (v: number) => {
    if (props.modelValue?.timer) {
      emit('update:model-value', { ...props.modelValue, timer: { ...props.modelValue.timer, dateNewLine: v } })
    }
  }
})

const datetype = computed({
  get: () => props.modelValue?.timer?.dateDisplayType ? 1 : 0,
  set: (v: number) => {
    if (props.modelValue?.timer) {
      emit('update:model-value', { ...props.modelValue, timer: { ...props.modelValue.timer, dateDisplayType: v } })
    }
  }
})

const showframe = computed({
  get: () => props.modelValue?.timer?.showFrame ? 1 : 0,
  set: (v: number) => {
    if (props.modelValue?.timer) {
      emit('update:model-value', { ...props.modelValue, timer: { ...props.modelValue.timer, showFrame: v } })
    }
  }
})

const time24 = computed({
  get: () => props.modelValue?.timer?.time24 || 24,
  set: (v: number) => {
    if (props.modelValue?.timer) {
      emit('update:model-value', { ...props.modelValue, timer: { ...props.modelValue.timer, time24: v } })
    }
  }
})

</script>
<template>
  <div class="layout-panel" :class="{ 'is-timer': isTimer }">
    <div class="layout-header mb-4">
      <h3 class="title is-5">
        属性设置
      </h3>
      <a
        role="button"
        class="action-link"
        tabindex="0"
        @keydown.space.prevent="emit('reset')"
        @click.prevent="emit('reset')"
      >重置</a>
    </div>
    <div class="layout-form">
      <template v-if="isTimer">
        <section class="layout-row">
          <div>位置</div>
          <div class="layout-row-inner">
            <div class="layout-cell half">
              <label for="tx">x</label>
              <input id="winx" v-model="tx" type="number" min="0" :max="base.w">
            </div>
            <div class="layout-cell half">
              <label for="ty">y</label>
              <input id="winy" v-model="ty" type="number" min="0" :max="base.h">
            </div>
          </div>
        </section>
        <section class="layout-row">
          <div class="layout-row-inner">
            <div class="layout-cell">
              <div>字体大小</div>
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
              <div>颜色</div>
              <color-picker v-model:pureColor="color" format="rgb" shape="square" theme="black" picker-type="chrome" />
            </div>
          </div>
        </section>
        <section class="layout-row">
          <div class="layout-row-inner">
            <div class="layout-cell">
              <div>是否显示日期</div>
              <VCheckbox v-model="showdate" />
            </div>
          </div>
        </section>
        <section class="layout-row">
          <div>是否换行</div>
          <div class="layout-row-inner">
            <div class="layout-cell">
              <VRadio
                v-model="newline"
                :value="0"
                label="一行显示"
                name="newline"
                color="primary"
              />
            </div>
            <div class="layout-cell">
              <VRadio
                v-model="newline"
                :value="1"
                label="两行显示"
                name="newline"
                color="primary"
              />
            </div>
          </div>
        </section>
        <section class="layout-row">
          <div>显示形式</div>
          <div class="layout-row-inner">
            <div class="layout-cell">
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
        <section class="layout-row">
          <div>时间显示帧</div>
          <div class="layout-row-inner">
            <div class="layout-cell">
              <VRadio
                v-model="showframe"
                :value="0"
                label="HH:MM:SS:FF"
                name="showframe"
                color="primary"
              />
            </div>
            <div class="layout-cell">
              <VRadio
                v-model="showframe"
                :value="1"
                label="HH:MM:SS"
                name="showframe"
                color="primary"
              />
            </div>
          </div>
        </section>
        <section class="layout-row">
          <div>小时制</div>
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
      </template>
      <template v-else>
        <section class="layout-row">
          <div>尺寸</div>
          <div class="layout-row-inner">
            <div class="layout-cell half">
              <label for="winw">w</label>
              <input id="winw" v-model="winw" type="number" min="0" :max="base.w">
            </div>
            <div class="layout-cell half">
              <label for="winh">h</label>
              <input id="winh" v-model="winh" type="number" min="0" :max="base.h">
            </div>
          </div>
        </section>
        <section class="layout-row">
          <div>位置</div>
          <div class="layout-row-inner">
            <div class="layout-cell half">
              <label for="winx">x</label>
              <input id="winx" v-model="winx" type="number" min="0" :max="base.w">
            </div>
            <div class="layout-cell half">
              <label for="winy">y</label>
              <input id="winy" v-model="winy" type="number" min="0" :max="base.h">
            </div>
          </div>
        </section>
        <section class="layout-row">
          <div class="layout-row-inner">
            <div class="layout-cell">
              <div>显示名称</div>
              <VCheckbox v-model="title" />
            </div>
          </div>
        </section>
        <section class="layout-row">
          <div class="layout-row-inner">
            <div class="layout-cell">
              <div>显示音柱</div>
              <VCheckbox v-model="vol" />
            </div>
          </div>
        </section>
      </template>
    </div>
  </div>
</template>
<style lang="scss">
.layout-panel {

  &.is-timer {

    .layout-header {
      position: sticky;
      top: 0;
      background: var(--dark-sidebar-light-6);
      padding-top: 20px;
      padding-bottom: 16px;
      margin-bottom: 0 !important;
      transition: none;
    }
  }

  .layout-header {
    display: flex;
    justify-content: space-between;

    h3.title {
      margin-bottom: 0;
    }
  }

  .layout-form {
    padding: 20px 0;
    font-size: 14px;
    color: var(--dark-dark-text);

    &.half {
      label {
        width: 10px;
      }
    }

    div {
      color: inherit;
    }

    .layout-row-inner {
      display: flex;
      align-items: center;
      gap: 16px;
      padding: 8px 0;
      flex-wrap: wrap;
    }

    .layout-cell {
      display: flex;
      align-items: center;
      gap: 8px;

      input[type='number'] {
        -moz-appearance: textfield;
      }

      input[type='number']::-webkit-outer-spin-button,
      input[type='number']::-webkit-inner-spin-button {
        -webkit-appearance: none;
      }

      input[type='text'],
      input[type='number'] {
        background-color: inherit;
        border: none;
        border-bottom: 1px solid #fff;
        line-height: 140%;
        border-radius: 0;
        width: 48px;

        &:focus-visible {
          outline: none;
        }
      }

      &.fullwidth {
        width: 100%;

        input[type='text'] {
          flex: 1;
        }
      }

      .checkbox {
        padding: 0;
        margin-inline-start: 16px;
      }

      .checkbox input + span {
        top: 0;
      }

      .radio {
        padding: 0.2rem 0;
        font-size: 0.85rem;

        input + span {
          width: 1rem;
          height: 1rem;

          &:after {
            top: 50%;
          }
        }
      }
    }
  }

  .layout-row {

    & + .layout-row {
      margin-top: 16px;
    }
  }
}

.vc-display .vc-color-input input,
.vc-display .vc-alpha-input__inner {
  color: #fff !important;
}

.vc-display .vc-input-toggle:hover {
  background-color: transparent !important;
}
</style>
