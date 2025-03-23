<script setup lang="ts">
import genAvatar from '@src/utils/avatar-gen'
import { useTemplateType } from '@src/stores/templateType'

const tmplTypeUtils = useTemplateType()

const opened = ref(false)
const draftType = ref({} as TemplateType)
const iconFile = ref<File>()

useListener(Signal.OpenNewTmplType, () => {
  opened.value = true
})

function onFileSelect(e: Event) {
  const files = (e.target as HTMLInputElement).files
  if (files?.[0]) {
    iconFile.value = files[0]
    draftType.value.icon = genAvatar()
  }
}
</script>
<template>
  <VModal
    is="form"
    method="post"
    novalidate
    size="small"
    :open="opened"
    title="新建设备类型"
    actions="right"
    cancel-label="取消"
    @submit.prevent="opened = false"
    @close="opened = false"
  >
    <template #content>
      <div class="modal-form">
        <VField label="类型名称 *">
          <VControl>
            <VInput v-model="draftType.name" type="text" placeholder="类型名称" />
          </VControl>
        </VField>
        <VField label="类型图标 *">
          <VControl>
            <div class="file has-name vso-form-upload">
              <label class="file-label">
                <input
                  class="file-input"
                  type="file"
                  name="resume"
                  @change="onFileSelect"
                />
                <span class="file-cta">
                  <span class="file-icon">
                    <i class="fas fa-cloud-upload-alt" />
                  </span>
                  <span class="file-label"> 选择文件 </span>
                </span>
                <span class="file-name light-text"> {{ iconFile?.name }} </span>
              </label>
            </div>
          </VControl>
        </VField>
      </div>
    </template>
    <template #action>
      <VButton type="submit" color="primary" raised @click="tmplTypeUtils.add(draftType)">
        保存
      </VButton>
    </template>
  </VModal>
</template>
