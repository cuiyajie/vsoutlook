interface TemplateType {
  id: string,
  name: string,
  icon: string,
  category: string
}

interface TmplRequirement {
  cpu: string,
  cpuNum: number,
  dpdkCpu: number,
  cpuCore: string,
  hugePage: number,
  memory: number,
  disk: string,
  gpu: string,
  inputBand: string,
  outputBand: string,
  network: string,
  chart: string,
  logLevel: number,
  repairRecvFrame: boolean,
  repairSendFrame: boolean,
  hostNetwork: boolean,
  utfOffset: number,
	recvAVFrameNodeCount: number,
	sendAVFrameNodeCount: number,
	recvframeCnt: number,
  maxRateMbpsByCore: number,
  receiveSessions: number,
  shm: number
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
	appName: string,
	podsStatus: [{ name: string, status: string }],
  targetNode: string,
  updatedAt: number,
  status: string,
  statusInfo: any,
  config: string,
  node: string
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

interface ClustNodeInfo {
  coreList: string,
  dmaList: string,
  vfCount: number,
}

interface ClustNode extends ClustNodeInfo {
  id: string,
  ip: string,
  running: string[],
  stopped: string[],
  allocatable: {
    cpu: number,
    memory: number,
    networkReceiveRate: number,
    networkTransmitRate: number
  },
  allocated: {
    cpu: number,
    memory: number,
    networkReceiveRate: number,
    networkTransmitRate: number
  }
}

interface ClustDevice {
  id: string,
  node: string,
  name: string,
  config: string
}
