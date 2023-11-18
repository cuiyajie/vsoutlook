<script lang="ts" setup></script>
<template>
  <VDropdown icon="feather:more-vertical" spaced right class="tmpl-page-dropdown">
    <template #content="{ close }">
      <a href="#" role="menuitem" class="dropdown-item is-media" @click="close()">
        <div class="icon">
          <i class="iconify" data-icon="feather:save" aria-hidden="true"></i>
        </div>
        <div class="meta">
          <span>保存</span>
        </div>
      </a>

      <a href="#" role="menuitem" class="dropdown-item is-media" @click="close()">
        <div class="icon">
          <i class="iconify" data-icon="feather:share" aria-hidden="true"></i>
        </div>
        <div class="meta">
          <span>另存为</span>
        </div>
      </a>

      <hr class="dropdown-divider" />

      <a href="#" role="menuitem" class="dropdown-item is-media" @click="close()">
        <div class="icon">
          <i class="iconify" data-icon="feather:activity" aria-hidden="true"></i>
        </div>
        <div class="meta">
          <span>创建程序</span>
        </div>
      </a>

      <hr class="dropdown-divider" />

      <a href="#" role="menuitem" class="dropdown-item is-media" @click="close()">
        <div class="icon">
          <i class="iconify" data-icon="feather:upload-cloud" aria-hidden="true"></i>
        </div>
        <div class="meta">
          <span>导入虚机镜像</span>
        </div>
      </a>

      <a href="#" role="menuitem" class="dropdown-item is-media" @click="close()">
        <div class="icon">
          <i class="iconify" data-icon="feather:server" aria-hidden="true"></i>
        </div>
        <div class="meta">
          <span>创建容器镜像</span>
        </div>
      </a>
    </template>
  </VDropdown>
</template>
<style lang="scss">
.dropdown.is-spaced.tmpl-page-dropdown {
  .dropdown-menu {
    min-width: 160px;
  }
}
</style>
