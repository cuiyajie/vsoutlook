<script setup lang="ts">
const modelValue = defineModel<any>({
  default: {
    cpu: "",
    cpuNum: "",
    cpuCore: '',
    hugePage: '',
    memory: "",
    disk: "",
    gpu: "",
    inputBand: "",
    outputBand: "",
    network: "",
    description: "",
  },
  local: true,
});
</script>

<template>
  <form
    method="post"
    novalidate
    class="form-layout"
    @submit.prevent=""
  >
    <div class="form-outer">
      <div class="form-body">
        <!--Fieldset-->
        <div class="form-fieldset">
          <div class="columns is-multiline">
            <div class="column is-4">
              <VField>
                <VLabel>CPU主频</VLabel>
                <VControl>
                  <VInput
                    v-model="modelValue.cpu"
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
                    v-model="modelValue.cpuNum"
                    type="text"
                    placeholder=""
                  />
                </VControl>
              </VField>
            </div>
            <div class="column is-4">
              <VField>
                <VLabel>CPU核心</VLabel>
                <VControl>
                  <VInput
                    v-model="modelValue.cpuCore"
                    type="text"
                    placeholder=""
                  />
                </VControl>
              </VField>
            </div>
            <div class="column is-4">
              <VField>
                <VLabel>大页数(Huge Page)</VLabel>
                <VControl>
                  <VInput
                    v-model="modelValue.hugePage"
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
                    v-model="modelValue.memory"
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
                    v-model="modelValue.gpu"
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
                    v-model="modelValue.disk"
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
                    v-model="modelValue.inputBand"
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
                    v-model="modelValue.outputBand"
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
                    v-model="modelValue.network"
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
                    v-model="modelValue.description"
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
  </form>
</template>

<style lang="scss">
@import "/@src/scss/abstracts/all";
@import "/@src/scss/components/forms-outer";

.form-layout {
  max-width: 740px;
  margin: 0 auto;

  .form-outer .form-body {
    padding: 10px 20px;
  }

  .form-fieldset {
    padding: 10px 0;
  }
}
</style>
