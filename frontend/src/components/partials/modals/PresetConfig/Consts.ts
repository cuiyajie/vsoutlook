export const VideoFormatTypes = [
  '文件录制',
  'ST2110-20',
  'ST2110-22',
  'NDI',
  'NVI',
  'SRT',
  'TS'
]
export const DefaultVideoFormatType = '文件录制'

export const VideoResolutions = [
  '720*576',
  '1920*1080',
  '3840*2160',
]
export const DefaultVideoResolution = '1920*1080'

export const VideoFrameRates = [
  '50I',
  '50P',
  '59.94I',
  '59.94P',
  '60I',
  '60P',
]
export const DefaultVideoFrameRate = '59.94I'

export const VideoGammas = [
  'SDR',
  'HLG',
  'PQ',
]
export const DefaultVideoGamma = 'SDR'

export const VideoColorSpaces = [
  'BT709',
  'BT2020',
]
export const DefaultVideoColorSpace = 'BT2020'

export const VideoEncodes = [
  'JPEG-XS',
  'XAVC-4K-INTRA',
  'Prores422',
  'H264',
  'H265',
  'MPEG2',
]
export const DefaultVideoEncode = 'JPEG-XS'

export const defVideoFormat = () => ({
  name: '',
  format: DefaultVideoFormatType,
  resolution: DefaultVideoResolution,
  frameRate: DefaultVideoFrameRate,
  gamma: DefaultVideoGamma,
  colorSpace: DefaultVideoColorSpace,
  encodeFormat: DefaultVideoEncode,
  bitRate: '10Mbps',
  bframe: 2,
  gop: 50
})


export const AudioQuantBits = [
  16,
  24,
  32,
]
export const DefaultAudioQuantBit = 24

export const AudioSampleRates = [
  48000,
]
export const DefaultAudioSampleRate = 48000

export const AudioEncodes = [
  'PCM',
  'AAC',
]
export const DefaultAudioEncode = 'PCM'

export const AudioTxInterval = [
  125,
  1000,
]
export const DefaultAudioTxInterval = 125

export const defAudioFormat = () => ({
  name: '',
  channels: 8,
  quantBits: DefaultAudioQuantBit,
  sampleRate: DefaultAudioSampleRate,
  encodeFormat: DefaultAudioEncode,
  txInterval: DefaultAudioTxInterval,
  bitRate: '128000bps',
  childType: '-'
})

export const VideoReviewKeys = ['黑场检测', '静帧检测', 'YUV超标', '视频丢失']
export const defVideoReviewRules: () => { [key in typeof VideoReviewKeys[number]]: VideoReviewRule } = () => ({
  '黑场检测': { name: '黑场检测', threshold: 0.997, duration: 1 },
  '静帧检测': { name: '静帧检测', threshold: 0.996, duration: 50 },
  'YUV超标': { name: 'YUV超标', threshold: 0.01, duration: 150 },
  '视频丢失': { name: '视频丢失', missingDuration: 100 },
})

export const AudioReviewKeys = ['音频静音', '音频过高', '音频过低', '音频丢失']
export const defAudioReviewRules: () => { [key in typeof AudioReviewKeys[number]]: AudioReviewRule } = () => ({
  '音频静音': { name: '音频静音', channels: '1-5', duration: 250 },
  '音频过高': { name: '音频过高', channels: '1-5', duration: 100 },
  '音频过低': { name: '音频过低', channels: '1-5', duration: 100, threshold: -60 },
  '音频丢失': { name: '音频丢失', missingDuration: 1000 },
})
