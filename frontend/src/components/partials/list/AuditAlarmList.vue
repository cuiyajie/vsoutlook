<script setup lang="ts">
import { useUserSession } from "@src/stores/userSession";
import { confirm } from "@src/utils/dialog";
import { VideoReviewKeyName, AudioReviewKeyName } from '../modals/PresetConfig/Consts';

const usStore = useUserSession();
const auditAlarmRules = computed(() => usStore.settings.audit_alarm_rules || []);
const page = ref(1)

function create() {
  bus.trigger(Signal.OpenNewAuditAlarm, {})
}

function remove(idx: number) {
  const tr = auditAlarmRules.value[idx]
  confirm({
    title: "确认",
    content: `确定要删除报警规则模版 ${tr.rule_name} 吗？`,
    onConfirm: async (hide) => {
      auditAlarmRules.value.splice(idx, 1)
      usStore.$updateSettings({
        key: 'audit_alarm_rules',
        value: JSON.stringify(auditAlarmRules.value)
      })
      hide()
    },
  });
}

function edit(idx: number) {
  bus.trigger(Signal.OpenNewAuditAlarm, {
    index: idx
  })
}

</script>

<template>
  <div>
    <div class="datatable-toolbar">
      <VButtons>
        <VButton
          color="primary"
          icon="fas fa-plus"
          elevated
          @click="create"
        >
          新建
        </VButton>
      </VButtons>
    </div>
    <div class="datatable-wrapper">
      <div class="table-container">
        <table class="table datatable-table is-fullwidth audit-alarm-table">
          <thead>
            <th align="center">模版名称</th>
            <th align="center">技审模板名称</th>
            <th align="center">视频报警规则</th>
            <th align="center">音频报警规则</th>
            <th align="center">任意声道静音</th>
            <th align="center">所有声道静音</th>
            <th>操作</th>
          </thead>
          <tbody>
            <tr
              v-for="(tr, tidx) in auditAlarmRules"
              :key="tr.rule_name"
            >
              <td>{{ tr.rule_name }}</td>
              <td>{{ tr.av_alarm.audit_template_name }}</td>
              <td>
                {{ tr.av_alarm.video_alarm_priority.map((v) => VideoReviewKeyName[v.itemname as TVideoRuleKey]).join("&nbsp; --> &nbsp;") }}
              </td>
              <td>
                {{ tr.av_alarm.audio_alarm_priority.map((v) => AudioReviewKeyName[v.itemname as TAudioRuleKey]).join("&nbsp; --> &nbsp;") }}
              </td>
              <td>{{ tr.av_alarm.audio_mute_rule.condition_any }}</td>
              <td>{{ tr.av_alarm.audio_mute_rule.condition_all }}</td>
              <td>
                <PresetListDropdown
                  @add="create"
                  @edit="edit(tidx)"
                  @remove="remove(tidx)"
                />
              </td>
            </tr>
          </tbody>
        </table>
      </div>
      <VPlaceholderPage
        v-if="auditAlarmRules.length === 0"
        title=""
        subtitle="暂无报警规则模版"
      />
    </div>

    <!--Table Pagination-->
    <VFlexPagination
      v-if="auditAlarmRules.length > 5"
      v-model:current-page="page"
      :item-per-page="6"
      :total-items="auditAlarmRules.length"
      :max-links-displayed="7"
      no-router
      class="mt-4"
    />
  </div>
  <NewAuditAlarmModal />
</template>
<style lang="scss">
.datatable-table.audit-alarm-table {

  td {
    font-size: 0.8rem;

    > span {
      display: block;

      & + span {
        margin-top: 0.25rem;
      }
    }
  }
}
</style>
