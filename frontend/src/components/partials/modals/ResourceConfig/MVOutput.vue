<!-- eslint-disable vue/prop-name-casing -->
<script setup lang="ts">
import { useFormat, getFormat } from './Utils';
import { formatMap, type MVOutputParamsType, pip_params } from './Consts';

const mv = defineModel<MVOutputParamsType>({
  default: {} as MVOutputParamsType,
  local: true,
});

defineProps<{
  index: number,
  isLast: boolean,
  useBackup: boolean,
  m_local_ip: string,
  b_local_ip: string,
}>()

const opened = ref(false)

const format = useFormat(mv.value, getFormat(mv.value));

const pips = ref<Array<{
  index: number,
  value: Ref<typeof pip_params>
}>>([]);

watch(() => mv.value.pips_number, (nv) => {
  const len = pips.value.length
  const params: any[] = mv.value.pip_params
  if (nv < len) {
    pips.value = pips.value.slice(0, nv);
  } else if (nv > len) {
    pips.value = [
      ...pips.value,
      ...Array.from({ length: nv - len }, (_, i) => ({
        index: len + i + 1,
        value: params[len + i] ? ref(params[len + i]) : ref({ ...pip_params })
      }))
    ]
  }
}, { immediate: true });
</script>
<template>
  <div
    class="form-fieldset collapse-form"
    :class="!isLast && 'seperator'"
    :open="opened || undefined"
  >
    <div
      class="fieldset-heading collapse-header"
      tabindex="0"
      role="button"
      @keydown.space.prevent="opened = !opened"
      @click.prevent="opened = !opened"
    >
      <h4>第{{ index }}路输出参数</h4>
      <div class="collapse-icon">
        <VIcon icon="feather:chevron-down" />
      </div>
    </div>
    <Transition name="fade-slow">
      <div
        v-show="opened"
        class="form-fieldset-nested is-tail"
      >
        <div class="form-fieldset">
          <div class="fieldset-heading">
            <h5>输出布局</h5>
          </div>
          <div class="columns is-multiline">
            <div class="column is-8">
              <VField>
                <VLabel>布局模板</VLabel>
                <VControl>
                  <div class="file has-name vso-form-upload">
                    <label class="file-label">
                      <input
                        class="file-input"
                        type="file"
                        name="resume"
                      >
                      <span class="file-cta">
                        <span class="file-icon">
                          <i class="fas fa-cloud-upload-alt" />
                        </span>
                        <span class="file-label"> 选择文件 </span>
                      </span>
                      <span class="file-name light-text"> {{ mv.mv_template }} </span>
                    </label>
                  </div>
                </VControl>
              </VField>
            </div>
            <div class="column is-4">
              <VField>
                <VLabel>窗口数量</VLabel>
                <VControl>
                  <VSelect
                    v-model="mv.pips_number"
                    class="is-rounded"
                  >
                    <VOption :value="4">
                      4
                    </VOption>
                    <VOption :value="9">
                      9
                    </VOption>
                    <VOption :value="10">
                      10
                    </VOption>
                  </VSelect>
                </VControl>
              </VField>
            </div>
          </div>
          <div
            v-for="pip in pips"
            :key="pip.index"
            class="columns is-multiline"
          >
            <div class="column is-6">
              <VField>
                <VLabel>窗口{{ pip.index }}名称</VLabel>
                <VInput
                  v-model="pip.value.pip_name"
                />
              </VField>
            </div>
            <div class="column is-6">
              <VField>
                <VLabel>窗口{{ pip.index }}关联输入视频序号</VLabel>
                <VInputNumber
                  v-model="pip.value.pip_video_index"
                  centered
                  :min="0"
                  :step="1"
                />
              </VField>
            </div>
          </div>
        </div>
      </div>
    </Transition>
    <Transition name="fade-slow">
      <div
        v-show="opened"
        class="form-fieldset-nested"
      >
        <div class="form-fieldset seperator">
          <div class="fieldset-heading">
            <h5>IP流参数</h5>
          </div>
          <div class="columns is-multiline">
            <div class="column is-6">
              <VField>
                <VLabel>视频流协议</VLabel>
                <VControl>
                  <VInput
                    v-model="mv.v_protocol"
                    readonly
                  />
                </VControl>
              </VField>
            </div>
            <div class="column is-6" />
            <div class="column is-6">
              <VField>
                <VLabel>主视频流组播源IP（含端口）</VLabel>
                <VControl>
                  <VInput
                    v-model="mv.ipstream_master.v_src_address"
                  />
                </VControl>
              </VField>
            </div>
            <div class="column is-6">
              <VField>
                <VLabel>主视频流组播目标IP（含端口）</VLabel>
                <AddrAddon
                  v-model="mv.ipstream_master.v_dst_address"
                  :host="m_local_ip"
                />
              </VField>
            </div>
            <Transition name="fade-slow">
              <div
                v-if="useBackup"
                class="column is-6"
              >
                <VField>
                  <VLabel>备视频流组播源IP（含端口）</VLabel>
                  <VControl>
                    <VInput
                      v-model="mv.ipstream_backup.v_src_address"
                    />
                  </VControl>
                </VField>
              </div>
            </Transition>
            <Transition name="fade-slow">
              <div
                v-if="useBackup"
                class="column is-6"
              >
                <VField>
                  <VLabel>备视频流组播目标IP（含端口）</VLabel>
                  <AddrAddon
                    v-model="mv.ipstream_backup.v_dst_address"
                    :host="b_local_ip"
                  />
                </VField>
              </div>
            </Transition>
          </div>
        </div>
      </div>
    </Transition>
    <Transition name="fade-slow">
      <div
        v-show="opened"
        class="form-fieldset-nested is-tail"
      >
        <div class="form-fieldset">
          <div class="fieldset-heading">
            <h5>视频参数</h5>
          </div>
          <div class="columns is-multiline">
            <div class="column is-6">
              <VField>
                <VLabel>视频格式</VLabel>
                <VControl>
                  <VInput
                    v-model="formatMap[format]"
                    readonly
                  />
                </VControl>
              </VField>
            </div>
          </div>
        </div>
      </div>
    </Transition>
  </div>
</template>