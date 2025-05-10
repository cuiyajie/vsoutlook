import { type PlayerParams, type IndexedNicDetail, type IApiParams } from './Consts_V1'
import omit from 'lodash-es/omit'
import merge from 'lodash-es/merge'

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

export function handleApiParams(params: IApiParams[]) {
  return params.filter((p) => p.checked).map((p) => omit(p, ['checked', 'label']))
}

export function checkApiParams(params: IApiParams[], apiParams: any[]) {
  return params.map((p) => {
    const included = apiParams.find((ap) => ap.api_name === p.api_name)
    return {
      ...merge(p, included),
      checked: !!included,
    }
  })
}

export function handleVideoFormat(name: string, vfs: VideoFormat[]) {
  const vf = vfs.find((v) => v.name === name)
  if (!vf) return null
  const cloned = { ...vf } as Partial<VideoFormat>
  if (cloned.fps) {
    if (cloned.interlaced) {
      cloned.fps = cloned.fps / 2
    }
  }
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

export function handleAudioMapping(name: string, afs: AudioMapping[]) {
  const af = afs.find((a) => a.name === name)
  if (!af) return null
  return { ...af }
}

export function handleNicList(nics: IndexedNicDetail[]) {
  return nics
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

export function checkNicDetails(nicDetails: any[], nics: NicInfo[]) {
  return nicDetails.map((nic: any) => {
    const nicIndex = nics.findIndex(
      (n) => n.nicNameMain === nic.nic_name_m && n.nicNameBackup === nic.nic_name_b
    )
    if (nicIndex !== -1) {
      return { ...nic, nicIndex, id: nics[nicIndex].id }
    }
    return { ...nic, nicIndex: -1, id: '' }
  })
}

export function wrap(src: any, prefix: string) {
  const dst: any = {}
  for (const key in src) {
    let wrapKey = key
    if (
      [
        'v_src_address',
        'v_dst_address',
        'a_src_address',
        'a_dst_address',
        '40_src_address',
        '40_dst_address',
      ].includes(key)
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
    let unwrapKey = key
    if (key.startsWith(prefix) && key.match(/(src|dst)_address$/)) {
      unwrapKey = key.replace(prefix, '')
    }
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

export function handleJsonData(data: any) {
  if (data.nic_list && Array.isArray(data.nic_list)) {
    data.nic_list = data.nic_list.map((d: any) => {
      // eslint-disable-next-line @typescript-eslint/no-unused-vars
      const { nicIndex, id, ...rest } = d
      return rest
    })
  }
  return data
}

export function processFormValue(data: any) {
  data = replaceKeyInObject(data, 'smpte_params', 'smpte-params')
  data = replaceKeyInObject(data, 'sub_bus', 'sub-bus')
  return data
}

export function checkFormValue(data: any) {
  data = replaceKeyInObject(data, 'smpte-params', 'smpte_params')
  data = replaceKeyInObject(data, 'sub-bus', 'sub_bus')
  return data
}

/**
 * Recursively replaces a specific key with a new key in a nested object
 * @param obj The object to process
 * @param oldKey The key to replace
 * @param newKey The new key to use
 * @returns A new object with the replaced keys
 */
export function replaceKeyInObject(obj: any, oldKey: string, newKey: string): any {
  // Handle null or undefined
  if (obj === null || obj === undefined) {
    return obj
  }

  // Handle arrays
  if (Array.isArray(obj)) {
    return obj.map((item) => replaceKeyInObject(item, oldKey, newKey))
  }

  // Handle objects
  if (typeof obj === 'object') {
    const result: any = {}

    for (const key in obj) {
      if (Object.prototype.hasOwnProperty.call(obj, key)) {
        // If the current key matches the old key, use the new key instead
        const newKeyName = key === oldKey ? newKey : key
        // Recursively process the value
        result[newKeyName] = replaceKeyInObject(obj[key], oldKey, newKey)
      }
    }

    return result
  }

  // Return primitive values as is
  return obj
}

/**
 * Checks if a string ends with a pattern like ':number' and adds ':30000' if it doesn't
 * @param str The string to check
 * @returns The string with ':30000' appended if it doesn't already end with ':number'
 */
export function ensurePortSuffix(str: string, defaultPort = 30000): string {
  // Check if the string already ends with :[number]
  const portPattern = /:\d+$/

  if (portPattern.test(str)) {
    return str // Already has a port number
  }

  // Add the default port
  return `${str}:${defaultPort}`
}
