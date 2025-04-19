import idGen from '@src/utils/id-gen'
import { def_player_params } from '../Utilties/Consts_V1'

export const audio_workmodes = [
  {
    key: 0,
    label: '不处理音频',
  },
  {
    key: 1,
    label: '音频跟随视频切换',
  },
  {
    key: 2,
    label: '音频固定直通指定音源',
  },
]

export const key_types = [
  {
    value: 'ext_key',
    label: '标准Key-Fill外键',
  },
  {
    value: 'int_key',
    label: '内键',
  },
  {
    value: 'h5_key',
    label: 'H5外键',
  },
]

export const switch_panel_types = [
  { value: 'panel_s1', label: '独立面板-1' },
  { value: 'panel_c1', label: '组合面板-1' },
]

export const bus_signal_types = [
  { value: 'video', label: '视频信号' },
  { value: 'key', label: '键信号' },
  { value: 'bus_output', label: '总线输出' },
]

export const bus_signal_types_map = new Map(
  bus_signal_types.map((bst) => [bst.value, bst])
)

export const bus_key_types = [
  { value: 'luma_key', label: '亮度键' },
  { value: 'self_key', label: '自键' },
  { value: 'line_key', label: '线性键' },
]

export const bus_sub_types = [
  { value: 'key_bus', label: '键总线' },
  { value: 'me_bus', label: 'ME总线' },
]

export const bus_sub_types_map = new Map(bus_sub_types.map((bst) => [bst.value, bst]))

export const switch_out_types = ['pgm', 'pvw', 'clean', 'aux']

export const def_tally_signal = (index: number) => ({
  index,
  level_input_index: index + 1,
  tally_index: index,
})

export const def_tally_screen = () => ({
  screen_id: 1,
  address: '192.168.1.116:6001',
  pgm_tally_index: 0,
  pvw_tally_index: 1,
  signal_number: 4,
  pgm_checked: false,
  pvw_checked: false,
  signal_config: [] as TallySignalParams[],
})

export const def_tally_config = (index: number) => ({
  level: index,
  screens: [def_tally_screen(), def_tally_screen()],
})

export const def_tally = (level: number) =>
  Array.from({ length: level }, (_, i) => def_tally_config(i))

export const def_switch_panel_url = (index: number) => ({
  subpanel_index: index,
  url: '192.168.1.110:9924',
})

export const def_switch_panel_params = (index: number) => ({
  level_index: index,
  panel_type: 'panel_s1',
  panel_url: [def_switch_panel_url(0)] as SwitchPanelUrl[],
})

export const def_switch_panel = (level: number) =>
  Array.from({ length: level }, (_, i) => def_switch_panel_params(i))

export const def_switch_input = () => ({
  key: [] as SwitchInputKeyParams[],
  video: [] as SwitchInputVideoParams[],
})

export const def_switch_input_key_params = (index: number) => ({
  index,
  key_type: 'int_key',
  signal_id: `key-${idGen(8)}`,
  signal_name: `input-key-${index}`,
  ...def_keyfill_player_params(),
  file_name: '',
  url: 'http://index.html',
})

export const def_switch_input_video_params = (index: number) => ({
  index,
  force_use_videoformat: false,
  signal_id: `video-${idGen(8)}`,
  signal_name: `input-video-${index}`,
  ...def_player_params(),
})

export const def_switch_bus = (level: number) => ({
  key_bus: [] as SwitchBusKeyParams[],
  me_bus: Array.from({ length: level }, (_, i) =>
    def_switch_bus_me_params(i)
  ) as SwitchBusMeParams[],
  level_bus: Array.from({ length: level }, (_, i) =>
    def_switch_bus_level_params(i, level)
  ) as SwitchBusLevelParams[],
})

export const def_switch_bus_key_params = (index: number) => ({
  index,
  bus_id: `keybus-${idGen(8)}`,
  bus_name: `key-${index}`,
  bus_input: {
    nic_index: -1,
    signal_type: 'video',
    signal_id: '',
    signal_name: '',
  },
  key_params: {
    keytype: 'luma_key',
    top_x: 0,
    top_y: 0,
    w: 1920,
    h: 1080,
  },
  out_signal: {
    signal_id: `keybusout-${idGen(8)}`,
    signal_name: `keybus-out-${index}`,
  },
})

export const def_switch_bus_me_params = (index: number) => ({
  index,
  bus_id: `mebus-${idGen(8)}`,
  bus_name: `me-${index}`,
  auto_duration: 100,
  out_signal: {
    signal_id: `mebusout-${idGen(8)}`,
    signal_name: `mebus-out-${index}`,
  },
})

export const def_switch_bus_level_input_params = (index: number) => ({
  index,
  input_index: index + 1,
  signal_type: 'video',
  signal_id: '',
  signal_name: '',
})

export const def_switch_bus_level_sub_params = (index: number) => ({
  index,
  bus_type: 'key_bus',
  bus_id: '',
  bus_name: '',
})

export const def_switch_bus_level_params = (index: number, level: number) => ({
  level: index,
  bus_input: {
  nic_index: -1,
  bus_input_number: level,
    input_list: Array.from({ length: level }, (_, i) =>
    def_switch_bus_level_input_params(i)
  )
  },
  pgm_bus: {
    sub_bus: [def_switch_bus_level_sub_params(0)] as SwitchBusLevelSubParams[],
    out_signal: [
      {
        index: 0,
        signal_type: 'clean',
        signal_name: `level${index}-pgmout-clean`,
        signal_id: `level${index}-pgmout-clean`,
      },
      {
        index: 1,
        signal_type: 'final',
        signal_name: `level${index}-pgmout-final`,
        signal_id: `level${index}-pgmout-final`,
      },
    ],
  },
  pvw_bus: {
    sub_bus: [def_switch_bus_level_sub_params(0)] as SwitchBusLevelSubParams[],
    out_signal: [
      {
        index: 0,
        signal_type: 'final',
        signal_name: `level${index}-pvwout-final`,
        signal_id: `level${index}-pvwout-final`,
      },
    ],
  },
})

export const def_switch_out_params = (index: number) => ({
  index,
  out_type: 'clean',
  out_name: `final-out-${index}`,
  ...def_player_params(),
  audio_mapping_name: '',
  mapping_checked: true,
  source_signal: {
    signal_id: '',
    signal_name: '',
    signal_type: 'bus_output',
  },
})

export const def_switch_out = (level: number) => ({
  out_number: level * 2,
  out_params: [] as SwitchOutParams[],
  audio_output_params: {
    src_signal_name: '',
    pgm_src_signal_id: '',
    in_nic_index: -1,
    ...def_player_params(),
  },
  mv_output_params: {
    mv_is_open: false,
    out_mv_template: '',
    ...def_player_params(),
  },
})

export const def_keyfill_player_params = () => {
  const playerParams = def_player_params()
  return {
    videoformat_name: '',
    smpte_params: {
      key_params: {
        ipstream_master: {
          ...playerParams.smpte_params.ipstream_master
        },
        ipstream_backup: {
          ...playerParams.smpte_params.ipstream_backup
        },
      },
      fill_params: {
        ipstream_master: {
          ...playerParams.smpte_params.ipstream_master
        },
        ipstream_backup: {
          ...playerParams.smpte_params.ipstream_backup
        },
      },
    },
    stream_params: {
      key_params: {
        url: playerParams.stream_params.url,
        codec_dev: playerParams.stream_params.codec_dev,
        codec_dev_index: playerParams.stream_params.codec_dev_index,
      },
      fill_params: {
        url: playerParams.stream_params.url,
        codec_dev: playerParams.stream_params.codec_dev,
        codec_dev_index: playerParams.stream_params.codec_dev_index,
      }
    }
  }
}

export type SignalParams = {
  signal_id: string
  signal_name: string
  signal_type: 'video' | 'audio'
}

export type TallySignalParams = ReturnType<typeof def_tally_signal>
export type TallyScreenParams = ReturnType<typeof def_tally_screen>
export type TallyConfigParams = ReturnType<typeof def_tally_config>
export type TallyParams = ReturnType<typeof def_tally>

export type SwitchPanelUrl = ReturnType<typeof def_switch_panel_url>
export type SwitchPanelParams = ReturnType<typeof def_switch_panel_params>
export type SwitchPanel = ReturnType<typeof def_switch_panel>

export type SwitchInput = ReturnType<typeof def_switch_input>
export type SwitchInputKeyParams = ReturnType<typeof def_switch_input_key_params>
export type SwitchInputVideoParams = ReturnType<typeof def_switch_input_video_params>
export type KeyFillPlayerParams = ReturnType<typeof def_keyfill_player_params>

export type SwitchBus = ReturnType<typeof def_switch_bus>
export type SwitchBusKeyParams = ReturnType<typeof def_switch_bus_key_params>
export type SwitchBusMeParams = ReturnType<typeof def_switch_bus_me_params>
export type SwitchBusLevelSubParams = ReturnType<typeof def_switch_bus_level_sub_params>
export type SwitchBusLevelInputParams = ReturnType<
  typeof def_switch_bus_level_input_params
>
export type SwitchBusLevelParams = ReturnType<typeof def_switch_bus_level_params>

export type SwitchOut = ReturnType<typeof def_switch_out>
export type SwitchOutParams = ReturnType<typeof def_switch_out_params>
export type SwitchOutAudioParams = SwitchOut['audio_output_params']
export type SwitchOutMvParams = SwitchOut['mv_output_params']

export const def_switch_template_toggle = () => ({
  videoSelected: () => {},
  videoUnSelected: () => {},
  audioSelected: () => {},
  audioUnSelected: () => {},
  audioMappingSelected: () => {},
  audioMappingUnSelected: () => {},
})

export interface InjectSwitchTemplateToggle {
  videoSelected: (name: string) => void
  videoUnSelected: (name: string) => void
  audioSelected: (name: string) => void
  audioUnSelected: (name: string) => void
  audioMappingSelected: (name: string) => void
  audioMappingUnSelected: (name: string) => void
}

export const def_switch_bus_select = () => ({
  selectedBus: ref(new Set<string>()),
  busSelected: () => {},
  busUnSelected: () => {},
})

export interface InjectSwitchBusSelect {
  selectedBus: Ref<Set<string>>
  busSelected: (bus: string) => void
  busUnSelected: (bus: string) => void
}

export const def_switch_input_signal_select = () => ({
  selectedSignal: ref(new Set<string>()),
  signalSelected: () => {},
  signalUnSelected: () => {},
})

export interface InjectSwitchInputSignalSelect {
  selectedSignal: Ref<Set<string>>
  signalSelected: (signal: string) => void
  signalUnSelected: (signal: string) => void
}

export const def_switch_key_signal_select = () => ({
  selectedKeySignal: ref(new Set<string>()),
  keySignalSelected: () => {},
  keySignalUnSelected: () => {},
})

export interface InjectSwitchKeySignalSelect {
  selectedKeySignal: Ref<Set<string>>
  keySignalSelected: (signal: string) => void
  keySignalUnSelected: (signal: string) => void
}

export interface InjectSwitchConfig {
  used_signal_type: number
  audio_workmode: number
}
