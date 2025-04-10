export const VideoFormatProtocols = [
  { value: 'file', label: '文件录制' },
  { value: 'st2110-20', label: '无压缩信号（st2110-20）' },
  { value: 'st2110-22', label: '浅压缩信号（st2110-22）' },
  { value: 'ndi', label: 'ndi信号' },
  { value: 'nvi', label: 'nvi信号' },
  { value: 'srt', label: 'srt信号' },
  { value: 'ts', label: 'ts信号（Over UDP）' },
  { value: 'rtmp', label: 'RTMP信号' },
]
export const VideoFormatProtocolsMap = VideoFormatProtocols.reduce(
  (acc, cur) => {
    acc[cur.value] = cur.label
    return acc
  },
  {} as { [key: string]: string }
)
export const DefaultVideoFormatProtocol = 'file'
export const VideoFormatPrefixMap: Record<string, string> = {
  ndi: 'ndi',
  nvi: 'nvi',
  srt: 'srt',
  ts: 'udp',
  rtmp: 'rtmp',
}

export const VideoFormatFpsList = [
  '25P',
  '29.97P',
  '30P',
  '50I',
  '50P',
  '59.94I',
  '59.94P',
  '60I',
  '60P',
]

export const VideoGammas = ['sdr', 'hlg', 'pq']
export const DefaultVideoGamma = 'hlg'

export const VideoGamuts = ['bt709', 'bt2020']
export const DefaultVideoGamut = 'bt2020'

export const VideoCompressionFormats = [
  'jpeg-xs',
  'xavc-4k-intra',
  'prores422',
  'h264',
  'h265',
  'mpeg2',
  'full ndi',
  'nvi-hf',
  'suvc',
  'avs',
  'avs2',
  'avs3',
]
export const DefaultVideoCompressionFormat = 'jpeg-xs'

export const VideoCompressionSubtypes = [
  '无',
  'high',
  'base',
  'main',
  'main10',
  'avs+',
  'main8',
  'cbg_c300',
  'cbg_c480',
  'vt_prores422',
  'vt_prores422_hq',
]
export const DefaultVideoCompressionSubtype = 'base'

export const VideoCompressionRatios = ['5:1', '8:1', '10:1', '12:1', '16:1']
export const DefaultVideoCompressionRatio = '8:1'

export const defaultFpsStr = '50P'
export const defVideoFormat = () => ({
  name: '',
  protocol: DefaultVideoFormatProtocol,
  width: 1920,
  height: 1080,
  interlaced: false,
  fps: 50,
  gamma: DefaultVideoGamma,
  gamut: DefaultVideoGamut,
  compression_format: DefaultVideoCompressionFormat,
  compression_subtype: DefaultVideoCompressionSubtype,
  compression_ratio: DefaultVideoCompressionRatio,
  bitrate_bps: 0,
  gop_b_frames: 0,
  gop_length: 0,
})

export const AudioBits = [16, 24, 32]
export const DefaultAudioBit = 24

export const AudioCompressionFormats = ['pcm', 'aac']
export const DefaultAudioCompressionFormat = 'pcm'

export const AudioPacketTimeUs = [125, 1000]
export const DefaultAudioPacketTimeUs = 125

export const AudioFrequencies = [48000]
export const DefaultAudioFrequency = 48000

export const defAudioFormat = () => ({
  name: '',
  channels_number: 8,
  bits: DefaultAudioBit,
  frequency: DefaultAudioFrequency,
  compression_format: DefaultAudioCompressionFormat,
  compression_subtype: '',
  packet_time_us: DefaultAudioPacketTimeUs,
  bitrate_bps: 128000,
})

export const VideoReviewKeys: TVideoRuleKey[] = [
  'video_black_frame',
  'video_static_frame',
  'video_yuv',
  'video_lost',
]
export const defVideoReviewKeys = () =>
  VideoReviewKeys.reduce(
    (acc, cur) => {
      acc[cur] = false
      return acc
    },
    {} as Record<TVideoRuleKey, boolean>
  )
export const VideoReviewKeyName: Record<TVideoRuleKey, string> = {
  video_black_frame: '黑场检测',
  video_static_frame: '静帧检测',
  video_yuv: 'YUV超标',
  video_lost: '视频丢失',
}
export const defVideoReviewRules: () => {
  [key in TVideoRuleKey]: VideoReviewRule
} = () => ({
  video_black_frame: {
    threshold_percentage: 0.997,
    duration_frames: 1,
  },
  video_static_frame: {
    threshold_percentage: 0.996,
    duration_frames: 50,
  },
  video_yuv: {
    threshold_percentage: 0.01,
    duration_frames: 150,
  },
  video_lost: { duration_ms: 100 },
})

export const AudioReviewKeys: TAudioRuleKey[] = [
  'audio_mute',
  'audio_high',
  'audio_low',
  'audio_lost',
]
export const defAudioReviewKeys = () =>
  AudioReviewKeys.reduce(
    (acc, cur) => {
      acc[cur] = false
      return acc
    },
    {} as Record<TAudioRuleKey, boolean>
  )
export const AudioReviewKeyName: Record<TAudioRuleKey, string> = {
  audio_mute: '音频静音',
  audio_high: '音频过高',
  audio_low: '音频过低',
  audio_lost: '音频丢失',
}
export const defAudioReviewRules: () => {
  [key in TAudioRuleKey]: AudioReviewRule
} = () => ({
  audio_mute: {
    detect_channels: '1-5',
    duration_frames: 250,
  },
  audio_high: {
    detect_channels: '1-5',
    duration_frames: 100,
  },
  audio_low: {
    detect_channels: '1-5',
    duration_frames: 100,
    threshold_dbfs: -60,
  },
  audio_lost: { duration_ms: 1000 },
})

export const defTechReview = () => ({
  name: '',
  ...defVideoReviewRules(),
  ...defAudioReviewRules(),
})

export const defAudioMapping = () => ({
  name: 'mapping-1',
  mute_channels: '2,3-6',
  copy_channels: [
    {
      src_channel: 0,
      dst_channel: 7,
    },
    {
      src_channel: 0,
      dst_channel: 8,
    },
  ],
})
