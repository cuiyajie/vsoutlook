import { def_player_params } from '../Utilties/Consts_V1'

export const def_mv_input_param = () => ({
  ...def_player_params(),
  force_use_videoformat: false,
  audit_alarm_rule_name: '',
})

export const def_mv_output_param = () => ({
  ...def_player_params(),
  audio_src_pip_index: -1,
  out_mv_template: '',
  screenindex: 0,
  pip_params: [] as MVPipParams[],
})

export const pip_params = {
  out_pip_name: '',
  out_pip_video_index: 0,
  tallyindex: 0,
}

export const def_mv_pip_params = (idx: number) => ({
  ...pip_params,
  out_pip_name: `cam${idx + 1}`,
  out_pip_video_index: idx,
  tallyindex: idx,
})

export const def_audit_av_template = () =>
  ({
    name: '',
    video_black_frame: {
      threshold_percentage: 0.997,
      duration_frames: 1,
    },
    video_static_frame: {
      threshold_percentage: 0.996,
      duration_frames: 50,
    },
    video_yuv: {
      threshold_percentage: 0.01,
      duration_frames: 150,
    },
    video_lost: {
      duration_ms: 100,
    },
    audio_mute: {
      detect_channels: '1-3,7,8',
      duration_frames: 250,
    },
    audio_high: {
      detect_channels: '1-3,7,8',
      duration_frames: 100,
    },
    audio_low: {
      detect_channels: '1-3,7,8',
      threshold_dbfs: -60.0,
      duration_frames: 250,
    },
    audio_lost: {
      duration_ms: 1000,
    },
  }) as {
    name: string
    [key: string]: any
  }

export const def_audit_alarm_rule = () => ({
  rule_name: '',
  av_alarm: {
    audit_template_name: '',
    video_alarm_priority: [
      {
        priority: 1,
        itemname: 'video_lost',
      },
      {
        priority: 2,
        itemname: 'video_black_frame',
      },
      {
        priority: 3,
        itemname: 'video_static_frame',
      },
      {
        priority: 4,
        itemname: 'video_yuv',
      },
    ],
    audio_alarm_priority: [
      {
        priority: 1,
        itemname: 'audio_lost',
      },
      {
        priority: 2,
        itemname: 'audio_mute',
      },
      {
        priority: 3,
        itemname: 'audio_high',
      },
      {
        priority: 4,
        itemname: 'audio_low',
      },
    ],
    audio_mute_rule: {
      condition_any: '1,2',
      condition_all: '1-4',
    },
  },
})

export type MVInputItemParam = ReturnType<typeof def_mv_input_param>
export type MVOutputItemParam = ReturnType<typeof def_mv_output_param>
export type MVPipParams = ReturnType<typeof def_mv_pip_params>
export type AuditAvTemplate = ReturnType<typeof def_audit_av_template>
export type AuditAlarmRule = ReturnType<typeof def_audit_alarm_rule>
