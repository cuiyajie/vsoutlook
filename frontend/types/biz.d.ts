/* eslint-disable @typescript-eslint/consistent-type-imports */
interface TemplateType {
  id: string
  name: string
  icon: string
  category: string
}

interface TmplNicConfig {
  dpdkCpu: number
  dma: number
}

interface TmplRequirement {
  cpu: string
  cpuNum: number
  cpuCore: string
  hugePage: number
  memory: number
  disk: string
  gpu: string
  inputBand: string
  outputBand: string
  network: string
  chart: string
  logLevel: number
  repairRecvFrame: boolean
  repairSendFrame: boolean
  hostNetwork: boolean
  utfOffset: number
  recvAVFrameNodeCount: number
  sendAVFrameNodeCount: number
  recvFrameCount: number
  maxRateMbpsByCore: number
  receiveSessions: number
  nicCount: number
  nicConfig: TmplNicConfig[]
  shm: number
}

interface TemplateData {
  id: string
  name: string
  type: string
  typeName: string
  typeCategory:
    | 'codec'
    | 'udx'
    | 'multiv'
    | 'swt'
    | 'endswt'
    | 'recorder'
    | 'media_gateway'
    | 'mv'
    | 'makeswt'
    | 'nmswt'
    | 'bcswt'
  description: string
  listed: boolean
  requirement: TmplRequirement
  flow: any
  mirrorType?: number
  version?: number
  virtualMirror?: string
  containerMirror?: string
}

interface TemplateDataVerbose extends TemplateData {
  tmplType: TemplateType
}

interface AbilityComp {
  id: string
  name: string
  type: string
  groupKey: string
  groupName: string
}

interface DeviceInfo {
  id: string
  tmplId: string
  status: number
  input: number
  output: number
  ptp: number
}

interface StatusInfo {
  text: string
  color: string
  icon?: string
}

interface DeviceVerbose extends DeviceInfo {
  tmpl: TemplateDataVerbose
  statusInfo: StatusInfo
  inputInfo: StatusInfo
  outputInfo: StatusInfo
  ptpInfo: StatusInfo
}

interface DeviceDetail {
  id: string
  name: string
  tmplID: string
  tmplName: string
  tmplTypeID: string
  tmplTypeName: string
  tmplTypeIcon: string
  tmplTypeCategory: string
  appName: string
  podsStatus: [{ name: string; status: string }]
  targetNode: string
  updatedAt: number
  status: string
  statusInfo: any
  config: string
  node: string
  nodeIP: string
}

interface SystemInfo {
  id: string
  name: string
  devices: string[]
}

interface SystemVerbose extends SystemInfo {
  deviceInfos: DeviceVerbose[]
}

interface UserData {
  id: string
  c
  name: string
  role: number
}

interface NicInfo {
  id: string
  nodeId: string
  nicNameMain: string
  nicNameBackup: string
  receiveMain: boolean
  receiveBackup: boolean
  sendMain: boolean
  sendBackup: boolean
  coreList: string
  dmaList: string
  vfCount: number
}

interface ClustNode {
  id: string
  ip: string
  running: string[]
  stopped: string[]
  nics: NicInfo[]
  allocatable: {
    cpu: number
    memory: number
    networkReceiveRate: number
    networkTransmitRate: number
  }
  allocated: {
    cpu: number
    memory: number
    networkReceiveRate: number
    networkTransmitRate: number
  }
}

interface NMosNode {
  id: string
  api: {
    endpoints: {
      host: string
      port: number
      protocol: string
    }[]
    versions: string[]
  }
  caps: any
  clocks: {
    name: string
    ref_type: string
  }[]
  description: string
  hostname: string
  href: string
  interfaces: {
    chassis_id: string
    name: string
    port_id: string
  }[]
  label: string
  services: any[]
  tags: any
  version: string
}

interface ClustDevice {
  id: string
  node: string
  name: string
  config: string
}

type Settings = {
  rds_server_url: string
  authorization_services: SettingAuthService[]
  auto_save_container_config: boolean
  mv_template_list: SettingMvTemplate[]
  mv_template_font: string
  mv_template_fonts: string[]
  endswt_api: string
  endswt_titles: string[]
  endswt_panel_types: SettingEndSwtPanelType[]
  lut_upscale_names: SettingLutScaleName[]
  lut_downscale_names: SettingLutScaleName[]
  video_formats: VideoFormat[]
  audio_formats: AudioFormat[]
  audio_mappings: AudioMapping[]
  tech_reviews: TechReview[]
  audit_alarm_rules: AuditAlarmRule[]
}

type SettingAuthService = {
  ip: string
  port: number
}

type SettingEndSwtPanelType = {
  type: string
}

type SettingLutScaleName = {
  name: string
  remark: string
}

type SettingMvTemplate = {
  name: string
  path: string
}

interface Layout {
  id: string
  name: string
  // eslint-disable-next-line @typescript-eslint/consistent-type-imports
  size: import('@src/utils/enums').LayoutSize
  published: boolean
  location: string
  content: LayoutDataItem[]
  createdAt: number
  updatedAt: number
}

interface LayoutSample {
  name: string
  value: [number, number][]
}

interface LayoutDimension {
  w: number
  h: number
}

interface LayoutPos {
  x: number
  y: number
}

interface LayoutRect extends LayoutDimension, LayoutPos {}

interface IndexedLayoutRect extends LayoutRect {
  index: number
  isSecondary?: boolean
}

interface LayoutWinBorder {
  top: number
  bottom: number
  left: number
  right: number
  color: string
}

interface LayoutWin extends LayoutRect {
  border: LayoutWinBorder
  showBorder: boolean
}

interface LayoutTimer extends LayoutPos {
  fontSize: number
  fontFamily: string
  color: string
  showDate: boolean
  dateDisplayType: import('src/utils/enums').DateDisplayType
  clockDisplayType: import('src/utils/enums').CLockDisplayType
  time24: number
  fontHeight: number
  fontGap: number
  timeColor: string
  timePosition: LayoutPos
}

interface LayoutTextRect {
  fontSize: number
  fontFamily: string
  color: string
  bgColor: string
}

interface LayoutText extends LayoutRect {
  gap: number
  rect1: LayoutTextRect
  rect2: LayoutTextRect
}

interface LayoutSingleVol extends LayoutPos {
  one_w: number
  g: number
  h: number
  start: number
  end: number
}

interface LayoutVol {
  vols: LayoutSingleVol[]
  len: number
  alpha: number
}

interface LayoutTitle extends LayoutRect {
  fontSize: number
  fontFamily: string
  bgColor: string
  color: string
  isVertial: boolean
  tallyType: import('@src/utils/enums').TallyType
  tallyBorderWidth: number
  tallyBorderColor: string
  tallyBgColor: string
}

interface LayoutArea extends LayoutRect {
  alpha: number
}

interface LayoutAlarmBorder extends LayoutRect {
  width: number
  color: string
  interval: number
}

interface LayoutAlarm extends LayoutRect {
  fontSize: number
  color: string
  bgColor: string
  border: LayoutAlarmBorder
  lang: import('@src/utils/enums').AlarmLang
}

interface LayoutMeta extends LayoutRect {
  fontSize: number
  color: string
  bgColor: string
}

type LayoutProps =
  | 'win'
  | 'timer'
  | 'text'
  | 'vol'
  | 'title'
  | 'vector'
  | 'oscillogram'
  | 'alarm'
  | 'meta'

interface TypedLayoutInfo {
  type: LayoutProps
  index: number
}

interface TypedLayoutRect extends LayoutRect, TypedLayoutInfo {
  rotated?: boolean
}

interface LayoutDataItem {
  id: string
  win: LayoutWin
  timer: LayoutTimer | null
  text: LayoutText | null
  vol: LayoutVol | null
  title: LayoutTitle | null
  vector: LayoutArea | null
  oscillogram: LayoutArea | null
  alarm: LayoutAlarm | null
  meta: LayoutMeta | null
}

interface VideoFormat {
  name: string
  protocol: string
  width: number
  height: number
  interlaced: boolean
  fps: number
  gamma: string
  gamut: string
  compression_format: string
  compression_subtype: string
  compression_ratio: string
  bitrate_bps: number
  gop_b_frames: number
  gop_length: number
}

interface AudioFormat {
  name: string
  channels_number: number
  bits: number
  frequency: number
  compression_format: string
  compression_subtype: string
  packet_time_us: number
  bitrate_bps: number
}

interface AudioChannelMapping {
  src_channel: number
  dst_channel: number
}
interface AudioMapping {
  name: string
  mute_channels: string
  copy_channels: {
    src_channel: number
    dst_channel: number
  }[]
}

interface VideoReviewRule {
  threshold_percentage?: number
  duration_frames?: number
  duration_ms?: number
}

interface AudioReviewRule {
  detect_channels?: string
  duration_frames?: number
  duration_ms?: number
  threshold_dbfs?: number
}

type TVideoRuleKey =
  | 'video_black_frame'
  | 'video_static_frame'
  | 'video_yuv'
  | 'video_lost'

type TAudioRuleKey = 'audio_mute' | 'audio_high' | 'audio_low' | 'audio_lost'

type TechReview = {
  name: string
} & {
  [key in TVideoRuleKey]: VideoReviewRule
} & {
  [key in TAudioRuleKey]: AudioReviewRule
}

interface AlarmPriority {
  priority: number
  itemname: string
}

interface AuditAlarmRule {
  rule_name: string
  av_alarm: {
    audit_template_name: string
    video_alarm_priority: AlarmPriority[]
    audio_alarm_priority: AlarmPriority[]
    audio_mute_rule: {
      condition_any: string
      condition_all: string
    }
  }
}
