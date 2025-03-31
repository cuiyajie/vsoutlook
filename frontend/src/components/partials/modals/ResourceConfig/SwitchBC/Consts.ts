import { audio_workmodes } from '../SwitchShare/Consts'
import { used_signal_types } from '../Utilties/Consts_V1'

export const bcswt_used_signal_types = used_signal_types.filter((v) => v.key === 0)
export const levels = [1]

export const bcswt_audio_workmodes = audio_workmodes.filter((v) => v.key === 1)

export const def_api_params = () => [
  {
    label: '播出控制接口',
    api_name: 'bc_control',
    service_port: 9012,
    checked: true,
  },
]
