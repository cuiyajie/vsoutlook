<script setup lang="ts">
import { useBackup } from "@src/stores/backup";

const tables = ref<{
  table: string,
  name: string,
  checked: boolean
}[]>([
  { table: 'tmpl_types', name: '应用类型', checked: false },
  { table: 'tmpls', name: '应用', checked: false },
  { table: 'devices', name: '设备', checked: false },
  { table: 'layouts', name: '多画面布局', checked: false },
  { table: 'settings', name: '系统设置', checked: false },
]);

const selectAll = (val: boolean) => {
  tables.value.forEach(table => table.checked = val);
}

const skipImport = ref(false);
const api = useBackup();

function exportData() {
  const selectedTables = tables.value.filter(table => table.checked);
  api.$exportData(selectedTables);
}

const fileInput = ref<HTMLInputElement | null>(null);
function importData() {
  fileInput.value?.click();
}

function onFileImported(e: Event) {
  const targetEl = e.target as HTMLInputElement;
  const file = targetEl.files?.[0];
  if (!file) return
  api.$importData(file, skipImport.value);
  targetEl.value = '';
}
</script>
<template>
  <div class="data-cards">
    <VCard radius="rounded">
      <div class="card-header2">
        <VIconWrap
          size="medium"
          icon="fas fa-file-export"
          color="light"
        />
        <h3 class="card-title">
          数据备份
        </h3>
        <VControl raw subcontrol class="check-all">
          <VCheckbox circle :model-value="tables.every(table => table.checked)" label="全选" color="primary" @update:model-value="selectAll" />
        </VControl>
        <div class="meta-right">
          <div class="buttons">
            <VButton
              color="primary"
              raised
              @keydown.space.prevent="exportData"
              @click.prevent="exportData"
            >
              导出
            </VButton>
          </div>
        </div>
      </div>
      <VField class="is-flex check-list">
        <VControl v-for="table in tables" :key="table.table" raw subcontrol>
          <VCheckbox v-model="table.checked" circle :label="table.name" color="primary" />
        </VControl>
      </VField>
    </VCard>
    <VCard radius="rounded">
      <div class="card-header2">
        <VIconWrap
          size="medium"
          icon="fas fa-file-import"
          color="light"
        />
        <h3 class="card-title">
          数据导入
        </h3>
        <div class="meta-right">
          <div class="buttons">
            <VButton
              color="primary"
              raised
              @keydown.space.prevent="importData"
              @click.prevent="importData"
            >
              导入
            </VButton>
          </div>
        </div>
      </div>
      <p class="import-note">注: 当数据出现相同的属性时，应当替换当前数据，还是保留当前数据，忽略导入数据</p>
      <VField class="radio-group">
        <VControl>
          <VRadio
            v-model="skipImport"
            :value="false"
            label="替换"
            name="outlined_radio"
            color="primary"
          />
          <VRadio
            v-model="skipImport"
            :value="true"
            label="保留"
            name="outlined_radio"
            color="primary"
          />
        </VControl>
      </VField>
      <input
        ref="fileInput"
        type="file"
        accept=".json"
        style="width: 0; height: 0; opacity: 0; position: absolute; top: -100px; left: -100px; z-index: -1;"
        @change="onFileImported"
      >
    </VCard>
  </div>
</template>

<style lang="scss">

.data-cards {

  .l-card {
    width: auto;
    max-width: 876px;
    padding: 12px 20px;
    display: flex;
    margin: 0 auto;
    flex-wrap: wrap;
    min-height: 96px;

    & + .l-card {
      margin-top: 16px;
    }

    .card-header2 {
      display: flex;
      align-items: center;
      width: 100%;
    }

    .check-all {
      margin-right: 16px;
    }

    .radio-group,
    .check-list {
      margin: 12px 0 12px -10px;
    }

    .import-note {
      font-style: italic;
      display: block;
      width: 100%;
      margin-top: 1.5rem;
    }

    .icon-wrap.is-medium {
      border-radius: 10px;

      img {
        width: 60%;
        border-radius: 0 !important;
      }
    }

    h3.card-title {
      margin-left: 16px;
      font-size: 18px;
      font-weight: 600;
      flex: 1 1 auto;
    }
  }
}
</style>
