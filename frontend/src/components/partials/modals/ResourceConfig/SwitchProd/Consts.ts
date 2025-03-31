import { audio_workmodes } from '../SwitchShare/Consts'
import { used_signal_types } from '../Utilties/Consts_V1'

export const makeswt_used_signal_types = used_signal_types.filter((v) => v.key === 0)
export const levels = [1, 2]

export const makeswt_audio_workmodes = [...audio_workmodes]

export const def_api_params = () => [
  {
    label: '动态控制接口',
    api_name: 'dynamic_config',
    service_port: 9010,
    checked: true,
  },
  {
    label: '宏定义接口',
    api_name: 'macro_config',
    service_port: 9011,
    checked: true,
  },
]
