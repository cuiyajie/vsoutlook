<script lang="ts" setup>
defineProps<{
  device: DeviceVerbose;
}>();
</script>
<template>
  <div class="kv-info">
    <div class="info-block-inner">
      <div class="title-wrap">
        <h3 class="dark-inverted">
          {{ device.tmpl.name }}
        </h3>
        <a
          class="action-link"
          tabindex="0"
        >修改</a>
      </div>
      <div class="info-block-line">
        <h4 class="dark-inverted">
          设备状态
        </h4>
        <VTag
          rounded
          :color="device.statusInfo.color"
        >
          {{ device.statusInfo.text }}
        </VTag>
      </div>
      <div class="info-block-line">
        <h4 class="dark-inverted">
          管理IP
        </h4>
        <span>//xx/xx/xx</span>
      </div>
      <div class="info-block-line">
        <h4 class="dark-inverted">
          专有参数1
        </h4>
        <span>ssh://xx.xxx</span>
      </div>
      <div class="info-block-line">
        <h4 class="dark-inverted">
          专有参数2
        </h4>
        <span>ssh://xx.xxx</span>
      </div>
      <div class="info-block-line">
        <h4 class="dark-inverted">
          信号输入1
        </h4>
        <span>//xx/xx/xx</span>
      </div>
      <div class="info-block-line">
        <h4 class="dark-inverted">
          信号输出1
        </h4>
        <span>ssh://xx.xxx</span>
      </div>
    </div>
  </div>
</template>
