<script setup lang="ts">
import { useVFieldContext } from "/@src/composable/useVFieldContext";

const isNaN = Number.isNaN || window.isNaN;
const REGEXP_NUMBER = /^-?(?:\d+|\d+\.\d+|\.\d+)(?:[eE][-+]?\d+)?$/;
const REGEXP_DECIMALS = /\.\d*(?:0|9){10}\d*$/;

const normalizeDecimalNumber = (value: number, times = 100000000000) =>
  REGEXP_DECIMALS.test(String(value)) ? Math.round(value * times) / times : value;

export interface VInputNumberProps {
  raw?: boolean;
  trueValue?: boolean;
  falseValue?: boolean;
  readonly?: boolean;
  max?: number;
  min?: number;
  step?: number;
  centered?: boolean;
}

const modelValue = defineModel<any>({
  default: "",
  local: true,
});
const props = withDefaults(defineProps<VInputNumberProps>(), {
  modelValue: "",
  trueValue: true,
  falseValue: false,
  max: Infinity,
  min: -Infinity,
  step: 1,
});

const { field, id } = useVFieldContext({
  create: false,
  help: "VInputNumber",
});

const internal = computed({
  get() {
    if (field?.value) {
      return field.value.value;
    } else {
      return modelValue.value;
    }
  },
  set(value: any) {
    let newValue = typeof value !== "number" ? parseFloat(value) : value;
    if (!isNaN(newValue)) {
      if (props.min <= props.max) {
        newValue = Math.min(props.max, Math.max(props.min, newValue));
      }
    }
    if (field?.value) {
      field.value.setValue(newValue);
    }
    modelValue.value = newValue;
  },
});

const classes = computed(() => {
  if (props.raw) return [];

  return ["input", "v-input", "v-input-number"];
});

const increasable = computed(() => {
  return isNaN(internal.value) || internal.value < props.max;
});

const decreasable = computed(() => {
  return isNaN(internal.value) || internal.value > props.min;
});

function decrease() {
  if (decreasable.value) {
    if (isNaN(internal.value)) {
      internal.value = 0;
    }
    internal.value = normalizeDecimalNumber(internal.value - props.step);
  }
}

function increase() {
  if (increasable.value) {
    if (isNaN(internal.value)) {
      internal.value = 0;
    }
    internal.value = normalizeDecimalNumber(internal.value + props.step);
  }
}

function paste(event: ClipboardEvent) {
  const clipboardData = event.clipboardData || (window as any).clipboardData;

  if (clipboardData && !REGEXP_NUMBER.test(clipboardData.getData("text"))) {
    event.preventDefault();
  }
}
</script>

<template>
  <div class="v-input-number__wrapper" :class="[centered && 'is-center']">
    <input
      :id="id"
      v-model="internal"
      :class="classes"
      type="number"
      :name="id"
      :min="min"
      :max="max"
      :step="step"
      :true-value="props.trueValue"
      :false-value="props.falseValue"
      :readonly="readonly"
      @change="field?.handleChange"
      @blur="field?.handleBlur"
      @paste="paste"
    />
    <label
      class="v-input-number__button v-input-number__button--minus"
      role="button"
      tabindex="-1"
      :disabled="readonly || !decreasable"
      @click.prevent="decrease"
    >
      <i class="fas fa-minus" aria-hidden="true"></i>
    </label>
    <label
      class="v-input-number__button v-input-number__button--plus"
      role="button"
      tabindex="-1"
      :disabled="readonly || !increasable"
      @click.prevent="increase"
    >
      <i class="fas fa-plus" aria-hidden="true"></i>
    </label>
  </div>
</template>
