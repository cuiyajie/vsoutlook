import { def_player_params } from '../Utilties/Consts_V1';

export const def_mv_input_param = () => ({
  ...def_player_params(),
  force_use_videoformat: false,
  audit_alarm_rule_name: '',
})

export const def_mv_output_param = () => ({
  ...def_player_params(),
  pip_params: [] as MVPipParams[],
})

export const pip_params = {
  pip_name: '',
  pip_video_index: 0,
  tallyindex: 0
};

export const def_mv_pip_params = (idx: number) => ({
  ...pip_params,
  pip_name: `cam${idx + 1}`,
  pip_video_index: idx,
  tallyindex: idx
})

export type MVInputItemParam = ReturnType<typeof def_mv_input_param>
export type MVOutputItemParam = ReturnType<typeof def_mv_output_param>
export type MVPipParams = ReturnType<typeof def_mv_pip_params>
