import { InputSignal, OutputSignal, PtpStatus } from "./enums"


export const DeviceStatusDic = {
  'deployed': { text: '已部署', color: 'primary' },
  'running': { text: '运行中', color: 'solid' }
}

// export const DeviceStatusDic = {
//   [DeviceStatus.Normal]: { text: '正常', color: 'primary' },
//   [DeviceStatus.Shutdown]: { text: '关机', color: 'solid' }
// }

export const InputSignalDic = {
  [InputSignal.Normal]: { text: '正常', color: 'blue' },
  [InputSignal.None]: { text: '无信号', color: 'danger' }
}

export const OutputSignalDic = {
  [OutputSignal.Normal]: { text: '正常', color: 'blue' },
  [OutputSignal.None]: { text: '无信号', color: 'danger' }
}

export const PtpStatusDic = {
  [PtpStatus.Locked]: { text: '锁定', color: 'warning' },
  [PtpStatus.Unlocked]: { text: '未锁定', color: 'purple' }
}

export const DeviceStatusList =
  Object.keys(DeviceStatusDic).map(key => ({ key, value: DeviceStatusDic[key] }))

export const InputSignalList =
  Object.keys(InputSignalDic).map(key => ({ key: Number(key), value: InputSignalDic[Number(key) as InputSignal] }))

export const OutputSignalList =
  Object.keys(OutputSignalDic).map(key => ({ key: Number(key), value: OutputSignalDic[Number(key) as OutputSignal] }))

export const PtpStatusList =
  Object.keys(PtpStatusDic).map(key => ({ key: Number(key), value: PtpStatusDic[Number(key) as PtpStatus] }))
