import { def_player_params } from '../Utilties/Consts_V1'

export const convert_mode = [
  { value: 'lut', label: 'lut表' },
  { value: 'dynamic', label: '动态' },
]

export const dynamic_mode = [
  { value: 1, label: '模式1' },
  { value: 2, label: '模式2' },
]

export const def_mg_player_params = () => ({
  ...def_player_params(),
  force_use_videoformat: false,
})

export const def_mg_output_item_params = () => ({
  ...def_player_params(),
  hdr_convert_params: {
    convert_mode: 'lut',
    lut_filename: '/home/lut/aaa.lut',
    dynamic_mode: 1,
    luma_gain: 1.0,
  },
})

export const luma_gain = [1.0, 1.1, 1.2, 1.3]

export type MgInputParams = ReturnType<typeof def_mg_player_params>
export type MgOutputItemParams = ReturnType<typeof def_mg_output_item_params>
