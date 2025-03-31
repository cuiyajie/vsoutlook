<script lang="ts" setup>

const props = withDefaults(defineProps<{
  speed?: number,
  minDuration?: number,
  maxDuration?: number
}>(), {
  speed: 1500,
  minDuration: 200,
  maxDuration: 500
})

// 动画钩子函数（带 TypeScript 类型注解）
const beforeEnter = (_el: Element) => {
  const el = _el as HTMLElement;
  el.style.height = '0';
  el.style.overflow = 'hidden';
};

const enter = (_el: Element, done: () => void) => {
  const el = _el as HTMLElement;
  const fullHeight = el.scrollHeight;
  const duration = Math.min(props.maxDuration, Math.max(props.minDuration, Math.round((fullHeight / props.speed) * 1000)));
  el.style.transition = `height ${duration}ms ease`;
  el.style.height = `${fullHeight}px`
  el.addEventListener('transitionend', () => {
    done()
  }, { once: true });
};

const afterEnter = (_el: Element) => {
  const el = _el as HTMLElement;
  el.style.removeProperty('height');
  el.style.removeProperty('overflow');
  el.style.removeProperty('transition');
};

const beforeLeave = (_el: Element) => {
  const el = _el as HTMLElement;
  el.style.height = `${el.scrollHeight}px`;
  el.style.overflow = 'hidden';
};

const leave = (_el: Element, done: () => void) => {
  const el = _el as HTMLElement;
  el.style.height = `${el.scrollHeight}px`;
  const duration = Math.min(props.maxDuration, Math.max(props.minDuration, Math.round((el.scrollHeight / props.speed) * 1000)));
  requestAnimationFrame(() => {
    el.style.transition = `height ${duration}ms linear`;
    el.style.height = '0';
    el.addEventListener('transitionend', () => {
      done()
    }, { once: true });
  })
};

const afterLeave = (_el: Element) => {
  const el = _el as HTMLElement;
  el.style.removeProperty('height');
  el.style.removeProperty('overflow');
  el.style.removeProperty('transition');
};

</script>
<template>
  <transition
    :css="false"
    @before-enter="beforeEnter"
    @enter="enter"
    @after-enter="afterEnter"
    @enter-cancelled="afterEnter"
    @before-leave="beforeLeave"
    @leave="leave"
    @after-leave="afterLeave"
    @leave-cancelled="afterLeave"
  >
    <slot></slot>
  </transition>
</template>
