interface TemplateType {
  id: string,
  name: string,
  icon: string
}

interface TmplRequirement {
  cpu: string,
  cpuNum: string,
  cpuCore: string,
  hugePage: string,
  memory: string,
  disk: string,
  gpu: string,
  inputBand: string,
  outputBand: string,
  network: string,
  chart: string,
  logLevel: number,
  repairRecvFrame: boolean,
  repairSendFrame: boolean,
  dmaList: string,
  hostNetwork: boolean,
  utfOffset: number,
	recvAVFrameNodeCount: number,
	sendAVFrameNodeCount: number,
	recvframeCnt: number,
  maxRateMbpsByCore: number
}

interface TemplateData {
  id: string,
  name: string,
  type: string,
  typeName: string,
  description: string,
  listed: boolean,
  requirement: TmplRequirement,
  flow: any,
  mirrorType?: number,
  version?: number,
  virtualMirror?: string,
  containerMirror?: string
}

interface TemplateDataVerbose extends TemplateData {
  tmplType: TemplateType,
}

interface AbilityComp {
  id: string,
  name: string,
  type: string,
  groupKey: string,
  groupName: string
}

interface DeviceInfo {
  id: string,
  tmplId: string,
  status: number,
  input: number,
  output: number,
  ptp: number
}

interface StatusInfo {
  text: string,
  color: string,
  icon?: string
}

interface DeviceVerbose extends DeviceInfo {
  tmpl: TemplateDataVerbose,
  statusInfo: StatusInfo,
  inputInfo: StatusInfo,
  outputInfo: StatusInfo,
  ptpInfo: StatusInfo
}

interface DeviceDetail {
	id: string,
	name: string,
	tmplID: string,
	tmplName: string,
	tmplTypeID: string,
	tmplTypeName: string,
  tmplTypeIcon: string,
  release: string,
	revision: string,
	chartVersion: string,
	appVersion: string,
	status: string,
  statusInfo: any,
  updated: string,
  phase: string
}

interface SystemInfo {
  id: string,
  name: string,
  devices: string[]
}

interface SystemVerbose extends SystemInfo {
  deviceInfos: DeviceVerbose[]
}

interface UserData {
  id: string,
  name: string,
  role: number
}

interface ClustNode {
  id: string,
  ip: string,
  info: {
    allocatable: { cpu: { value: string }, memory: { value: string } },
    current: { cpu: { value: string }, memory: { value: string } },
  }
}

interface ClustDevice {
  id: string,
  node: string,
  release: string
}
