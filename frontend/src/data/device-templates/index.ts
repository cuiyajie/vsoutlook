import { DeviceStatus, InputSignal, MirrorType, OutputSignal, PtpStatus } from '/@src/utils/enums'

export const deviceTemplates = [
  { id: '1', name: 'JPEG-XS解码器', description: '', type: '1', version: 1, mirrorType: MirrorType.Container },
  { id: '2', name: 'PNG-XS解码器', description: '', type: '1', version: 1, mirrorType: MirrorType.Container },
  { id: '3', name: 'WEBP-XS解码器', description: '', type: '1', version: 1, mirrorType: MirrorType.Container },
  { id: '4', name: 'GIF-XS解码器', description: '', type: '1', version: 1, mirrorType: MirrorType.Virtual },
  { id: '5', name: '16路高清多画面', description: '', type: '2', version: 1, mirrorType: MirrorType.Virtual },
  { id: '6', name: '32路4K多画面', description: '', type: '2', version: 1, mirrorType: MirrorType.Virtual }
]

export const templateTypes = [
  { id: '1',  name: '编码器', icon: '/images/avatars/svg/vuero-1.svg'},
  { id: '2',  name: '多画面', icon: '/images/avatars/svg/vuero-2.svg'}
]

export const abilityComponents = [
  { id: '1', name: '2110流发送', groupType: 'input', groupName: '发送模块' },
  { id: '2', name: '2111流发送', groupType: 'input', groupName: '发送模块' },
  { id: '3', name: '2112流发送', groupType: 'input', groupName: '发送模块' },
  { id: '4', name: '2110流接收', groupType: 'output', groupName: '接收模块' },
  { id: '5', name: '2111流接收', groupType: 'output', groupName: '接收模块' },
  { id: '6', name: '2112流接收', groupType: 'output', groupName: '接收模块' },
  { id: '7', name: '色域转换', groupType: 'default', groupName: '转码模块' },
  { id: '8', name: '上变换', groupType: 'default', groupName: '转码模块' },
  { id: '9', name: '下变换', groupType: 'default', groupName: '转码模块' },
  { id: '10', name: '声音转换', groupType: 'default', groupName: '转码模块' },
]

export const devices = [
  { id: '1', tmplId: '1', status: DeviceStatus.Normal, input: InputSignal.Normal, output: OutputSignal.None, ptp: PtpStatus.Locked },
  { id: '2', tmplId: '2', status: DeviceStatus.Normal, input: InputSignal.None, output: OutputSignal.Normal, ptp: PtpStatus.Unlocked },
  { id: '3', tmplId: '2', status: DeviceStatus.Shutdown, input: InputSignal.None, output: OutputSignal.None, ptp: PtpStatus.Unlocked },
  { id: '4', tmplId: '4', status: DeviceStatus.Normal, input: InputSignal.Normal, output: OutputSignal.None, ptp: PtpStatus.Locked },
  { id: '5', tmplId: '5', status: DeviceStatus.Normal, input: InputSignal.None, output: OutputSignal.Normal, ptp: PtpStatus.Unlocked },
  { id: '6', tmplId: '6', status: DeviceStatus.Shutdown, input: InputSignal.None, output: OutputSignal.None, ptp: PtpStatus.Unlocked },
  { id: '7', tmplId: '1', status: DeviceStatus.Normal, input: InputSignal.Normal, output: OutputSignal.None, ptp: PtpStatus.Locked },
  { id: '8', tmplId: '2', status: DeviceStatus.Normal, input: InputSignal.None, output: OutputSignal.Normal, ptp: PtpStatus.Unlocked },
  { id: '9', tmplId: '2', status: DeviceStatus.Shutdown, input: InputSignal.None, output: OutputSignal.None, ptp: PtpStatus.Unlocked },
  { id: '10', tmplId: '4', status: DeviceStatus.Normal, input: InputSignal.Normal, output: OutputSignal.None, ptp: PtpStatus.Locked },
  { id: '11', tmplId: '6', status: DeviceStatus.Normal, input: InputSignal.None, output: OutputSignal.Normal, ptp: PtpStatus.Unlocked },
  { id: '12', tmplId: '5', status: DeviceStatus.Shutdown, input: InputSignal.None, output: OutputSignal.None, ptp: PtpStatus.Unlocked },
  { id: '13', tmplId: '1', status: DeviceStatus.Normal, input: InputSignal.Normal, output: OutputSignal.None, ptp: PtpStatus.Locked },
  { id: '14', tmplId: '2', status: DeviceStatus.Normal, input: InputSignal.None, output: OutputSignal.Normal, ptp: PtpStatus.Unlocked },
  { id: '15', tmplId: '2', status: DeviceStatus.Shutdown, input: InputSignal.None, output: OutputSignal.None, ptp: PtpStatus.Unlocked },
  { id: '16', tmplId: '4', status: DeviceStatus.Normal, input: InputSignal.Normal, output: OutputSignal.None, ptp: PtpStatus.Locked },
  { id: '17', tmplId: '5', status: DeviceStatus.Normal, input: InputSignal.None, output: OutputSignal.Normal, ptp: PtpStatus.Unlocked },
  { id: '18', tmplId: '6', status: DeviceStatus.Shutdown, input: InputSignal.None, output: OutputSignal.None, ptp: PtpStatus.Unlocked },
  { id: '19', tmplId: '1', status: DeviceStatus.Normal, input: InputSignal.Normal, output: OutputSignal.None, ptp: PtpStatus.Locked },
  { id: '20', tmplId: '2', status: DeviceStatus.Normal, input: InputSignal.None, output: OutputSignal.Normal, ptp: PtpStatus.Unlocked },
  { id: '21', tmplId: '2', status: DeviceStatus.Shutdown, input: InputSignal.None, output: OutputSignal.None, ptp: PtpStatus.Unlocked },
  { id: '22', tmplId: '4', status: DeviceStatus.Normal, input: InputSignal.Normal, output: OutputSignal.None, ptp: PtpStatus.Locked },
  { id: '23', tmplId: '5', status: DeviceStatus.Normal, input: InputSignal.None, output: OutputSignal.Normal, ptp: PtpStatus.Unlocked },
  { id: '24', tmplId: '6', status: DeviceStatus.Shutdown, input: InputSignal.None, output: OutputSignal.None, ptp: PtpStatus.Unlocked },
]
