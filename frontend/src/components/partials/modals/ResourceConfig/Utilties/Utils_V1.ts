import { type PlayerParams, type IndexedNicDetail } from './Consts_V1'

export function changeProtocol(url: string, newProtocol: string): string {
  try {
    const parsedUrl = new URL(url)
    // 替换协议部分（保留末尾的冒号）
    parsedUrl.protocol = newProtocol + ':'
    return parsedUrl.toString()
  } catch (error) {
    console.error('Invalid URL:', error)
    return url // 返回原始 URL 作为兜底
  }
}

export function handleVideoFormat(name: string, vfs: VideoFormat[]) {
  const vf = vfs.find((v) => v.name === name)
  if (!vf) return null
  const cloned = { ...vf } as Partial<VideoFormat>
  if (!cloned.compression_subtype) {
    delete cloned.compression_subtype
  }
  if (!cloned.compression_ratio) {
    delete cloned.compression_ratio
  }
  if (!cloned.bitrate_bps) {
    delete cloned.bitrate_bps
  }
  if (!cloned.gop_b_frames) {
    delete cloned.gop_b_frames
  }
  if (!cloned.gop_length) {
    delete cloned.gop_length
  }
  return cloned
}

export function handleAudioFormat(name: string, afs: AudioFormat[]) {
  const af = afs.find((a) => a.name === name)
  if (!af) return null
  const cloned = { ...af } as Partial<AudioFormat>
  if (!cloned.compression_subtype) {
    delete cloned.compression_subtype
  }
  if (!cloned.bitrate_bps) {
    delete cloned.bitrate_bps
  }
  return cloned
}

export function handlePlayerParams<T extends PlayerParams>(
  params: T,
  vfs: VideoFormat[]
) {
  const { videoformat_name, audioformat_name, smpte_params, stream_params, ...rest } =
    params
  const vf = vfs.find((v) => v.name === videoformat_name)
  const result = { videoformat_name, audioformat_name, ...rest } as any
  if (!vf?.protocol) return result
  const showParams = ['st2110-20', 'st2110-22'].includes(vf.protocol)
  if (showParams) {
    result.smpte_params = smpte_params
  } else {
    result.stream_params = stream_params
  }
  return result
}

export function handleNicList(nics: IndexedNicDetail[]) {
  return nics.map((nic) => {
    // eslint-disable-next-line @typescript-eslint/no-unused-vars
    const { index, ...rest } = nic
    return rest
  })
}

export function checkPlayerParams(params: any, vfs: VideoFormat[], afs: AudioFormat[]) {
  const vf = vfs.find((v) => v.name === params.videoformat_name)
  if (!vf) {
    params.videoformat_name = ''
  }
  const af = afs.find((a) => a.name === params.audioformat_name)
  if (!af) {
    params.audioformat_name = ''
  }
  return params
}

export function wrap(src: any, prefix: string) {
  const dst: any = {}
  for (const key in src) {
    let wrapKey = key
    if (
      ['v_src_address', 'v_dst_address', 'a_src_address', 'a_dst_address'].includes(key)
    ) {
      wrapKey = `${prefix}${key}`
    }
    if (src[key] instanceof Array) {
      dst[wrapKey] = src[key].map((v: any, index: number) =>
        Object.assign(wrap(v, prefix), { index })
      )
    } else if (typeof src[key] === 'object') {
      dst[wrapKey] = wrap(src[key], prefix)
      if (Object.keys(dst[wrapKey]).length === 0) {
        dst[wrapKey] = null
      }
    } else {
      dst[wrapKey] = src[key]
    }
  }
  return dst
}

export function unwrap(src: any, prefix: string) {
  const dst: any = {}
  for (const key in src) {
    const unwrapKey = key.replace(prefix, '')
    if (src[key] instanceof Array) {
      dst[unwrapKey] = src[key].map((v: any, index: number) =>
        Object.assign(unwrap(v, prefix), { index })
      )
    } else if (typeof src[key] === 'object') {
      if (src[key] === null) {
        dst[unwrapKey] = null
      } else {
        dst[unwrapKey] = unwrap(src[key], prefix)
      }
    } else {
      dst[unwrapKey] = src[key]
    }
  }
  return dst
}
