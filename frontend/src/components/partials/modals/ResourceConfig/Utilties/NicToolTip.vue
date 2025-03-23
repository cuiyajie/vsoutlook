<script lang="ts" setup>
defineProps<{
  nic: NicInfo
}>()

const tippyAppendTo = () => document.body
</script>
<template>
  <tippy
    animation="scale"
    theme="light"
    placement="right"
    :arrow="true"
    arrow-type="round"
    :interactive="true"
    :append-to="tippyAppendTo"
  >
    <template #content>
      <div class="kv-info tippy-detail">
        <div class="info-block-inner is-dark-bordered-12">
          <div class="info-block-line">
            <h4 class="dark-inverted">
              主路网口
            </h4>
            <span>
              {{ nic.nicNameMain }}
            </span>
          </div>
          <div class="info-block-line">
            <h4 class="dark-inverted">
              主路收开关
            </h4>
            <span class="switch-tag" :class="nic.receiveMain && 'is-receive-open'"></span>
          </div>
          <div class="info-block-line">
            <h4 class="dark-inverted">
              主路发开关
            </h4>
            <span class="switch-tag" :class="nic.sendMain && 'is-send-open'"></span>
          </div>
          <div class="info-block-line">
            <h4 class="dark-inverted">
              备路网口
            </h4>
            <span>
              {{ nic.nicNameBackup }}
            </span>
          </div>
          <div class="info-block-line">
            <h4 class="dark-inverted">
              备路收开关
            </h4>
            <span class="switch-tag" :class="nic.receiveBackup && 'is-receive-open'"></span>
          </div>
          <div class="info-block-line">
            <h4 class="dark-inverted">
              备路发开关
            </h4>
            <span class="switch-tag" :class="nic.sendBackup && 'is-send-open'"></span>
          </div>
          <div class="info-block-line">
            <h4 class="dark-inverted">
              CPU 隔离核
            </h4>
            <span>
              {{ nic.coreList }}
            </span>
          </div>
          <div class="info-block-line">
            <h4 class="dark-inverted">
              DMA
            </h4>
            <span>
              {{ nic.dmaList }}
            </span>
          </div>
          <div class="info-block-line">
            <h4 class="dark-inverted">
              VF Count
            </h4>
            <span>
              {{ nic.vfCount }}
            </span>
          </div>
        </div>
      </div>
    </template>
    <slot></slot>
  </tippy>
</template>
