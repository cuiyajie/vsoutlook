<script setup lang="ts">
const opened = ref(false);
const draftTmpl = ref({} as TemplateData);

const tmplFile = ref<File>();

function onFileSelect(e: Event) {
  const files = (e.target as HTMLInputElement).files;
  if (files?.[0]) {
    tmplFile.value = files[0];
  }
}

useListener(Signal.OpenTemplateImport, () => {
  opened.value = true;
});
</script>
<template>
  <VModal
    is="form"
    method="post"
    novalidate
    size="medium"
    :open="opened"
    title="导入本地应用"
    actions="right"
    cancel-label="取消"
    @submit.prevent="opened = false"
    @close="opened = false"
  >
    <template #content>
      <div class="form-outer modal-form">
        <div class="form-body">
          <div
            class="form-fieldset"
            style="padding-top: 0"
          >
            <VField label="选择导入的应用文件 *">
              <VControl>
                <div class="file has-name vso-form-upload">
                  <label class="file-label">
                    <input
                      class="file-input"
                      type="file"
                      name="resume"
                      @change="onFileSelect"
                    >
                    <span class="file-cta">
                      <span class="file-icon">
                        <i class="fas fa-cloud-upload-alt" />
                      </span>
                      <span class="file-label"> 选择文件 </span>
                    </span>
                    <span class="file-name light-text"> {{ tmplFile?.name }} </span>
                  </label>
                </div>
              </VControl>
            </VField>
          </div>
          <!--Fieldset-->
          <div class="form-fieldset">
            <div class="fieldset-heading">
              <h4>应用信息</h4>
            </div>

            <div class="columns is-multiline">
              <div class="column is-6">
                <VField label="设备名称 *">
                  <VControl>
                    <VInput
                      v-model="draftTmpl.name"
                      type="text"
                      placeholder="设备名称"
                    />
                  </VControl>
                </VField>
              </div>
              <div class="column is-6">
                <TmplTypeSelect
                  v-model="draftTmpl.type"
                  label="应用类型 *"
                />
              </div>
              <div class="column is-6">
                <VField>
                  <VLabel>版本</VLabel>
                  <VControl>
                    <VInput
                      model-value="V1"
                      type="text"
                      readonly
                    />
                  </VControl>
                </VField>
              </div>
              <div class="column is-6">
                <VField>
                  <VLabel>镜像类型</VLabel>
                  <VControl>
                    <VInput
                      model-value="虚拟镜像"
                      type="text"
                      readonly
                    />
                  </VControl>
                </VField>
              </div>
            </div>
          </div>
          <!--Fieldset-->
          <div class="form-fieldset">
            <div class="fieldset-heading">
              <h4>所需资源</h4>
            </div>

            <div class="columns is-multiline">
              <div class="column is-4">
                <VField>
                  <VLabel>CPU主频</VLabel>
                  <VControl>
                    <VInput
                      type="text"
                      placeholder=""
                    />
                  </VControl>
                </VField>
              </div>
              <div class="column is-4">
                <VField>
                  <VLabel>CPU核数</VLabel>
                  <VControl>
                    <VInput
                      type="text"
                      placeholder=""
                    />
                  </VControl>
                </VField>
              </div>
              <div class="column is-4">
                <VField>
                  <VLabel>CPU核数</VLabel>
                  <VControl>
                    <VInput
                      type="text"
                      placeholder=""
                    />
                  </VControl>
                </VField>
              </div>
              <div class="column is-4">
                <VField>
                  <VLabel>内存</VLabel>
                  <VControl>
                    <VInput
                      type="text"
                      placeholder=""
                    />
                  </VControl>
                </VField>
              </div>
              <div class="column is-4">
                <VField>
                  <VLabel>GPU</VLabel>
                  <VControl>
                    <VInput
                      type="text"
                      placeholder=""
                    />
                  </VControl>
                </VField>
              </div>
              <div class="column is-4">
                <VField>
                  <VLabel>硬盘</VLabel>
                  <VControl>
                    <VInput
                      type="text"
                      placeholder=""
                    />
                  </VControl>
                </VField>
              </div>
              <div class="column is-4">
                <VField>
                  <VLabel>信号输入宽带</VLabel>
                  <VControl>
                    <VInput
                      type="text"
                      placeholder=""
                    />
                  </VControl>
                </VField>
              </div>
              <div class="column is-4">
                <VField>
                  <VLabel>信号输出宽带</VLabel>
                  <VControl>
                    <VInput
                      type="text"
                      placeholder=""
                    />
                  </VControl>
                </VField>
              </div>
              <div class="column is-4">
                <VField>
                  <VLabel>光纤网卡端口数</VLabel>
                  <VControl>
                    <VInput
                      type="text"
                      placeholder=""
                    />
                  </VControl>
                </VField>
              </div>
              <div class="column is-12">
                <VField>
                  <VLabel>备注说明</VLabel>
                  <VControl>
                    <VTextarea
                      class="textarea"
                      rows="4"
                      autocomplete="off"
                      autocapitalize="off"
                      spellcheck="true"
                    />
                  </VControl>
                </VField>
              </div>
            </div>
          </div>
        </div>
      </div>
    </template>
    <template #action>
      <VButton
        type="submit"
        color="primary"
        raised
        @click="tmplUtils.add(draftTmpl)"
      >
        保存
      </VButton>
    </template>
  </VModal>
</template>
<style lang="scss" scoped>
.modal-form {
  .form-fieldset {
    padding: 15px 0;

    .fieldset-heading {
      margin-bottom: 16px;
    }
  }
}
</style>
