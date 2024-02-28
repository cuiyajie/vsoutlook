import { InputSignal, OutputSignal, PtpStatus } from "./enums"

export enum DeviceStatus {
  Running = 'Running',
  Pending = 'Pending',
  Terminating = 'Terminating',
  Failed = 'Failed',
  Terminated = 'Terminated',
  Unavailable = 'Unavailable'
}

export const DeviceStatusDic: Record<string, { text: string, color: string }> = {
  'Running': { text: '运行中', color: 'primary' },
  'Pending': { text: '创建中', color: 'blue' },
  'Terminating': { text: '删除中', color: 'purple' },
  'Failed': { text: '启动失败', color: 'danger' },
  'Terminated': { text: '已暂停', color: 'warning' },
  'Unavailable': { text: '未知状态', color: 'light' },
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
