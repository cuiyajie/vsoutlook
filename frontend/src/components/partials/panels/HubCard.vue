<script setup lang="ts">
export type VCardColor =
  | 'primary'
  | 'secondary'
  | 'info'
  | 'success'
  | 'warning'
  | 'danger'

const props = defineProps<{
  title?: string
  content?: string,
  link?: string,
  color: VCardColor,
}>()
</script>

<template>
  <RouterLink class="hub-card" :to="link" :class="[props.color && `is-${props.color}`]">
    <div class="hub-card-body">
      <slot name="icon" />
      <div class="inner-content">
        <h3 class="title is-5">
          {{ props.title }}
        </h3>
        <p class="light-text">
          {{ props.content }}
        </p>
      </div>
    </div>
    <!-- <div class="hub-card-foot">
      <slot />
    </div> -->
  </RouterLink>
</template>
