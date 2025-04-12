export function isValidNumberRanges(input: string): boolean {
  // 分割字符串为多个子范围
  const ranges = input.split(',')

  for (const range of ranges) {
    const trimmed = range.trim()

    // 忽略空字符串
    if (trimmed === '') continue

    // 分割范围为起始和结束数字
    const parts = trimmed.split('-')

    // 多个 '-' 或无效分割时返回 false
    if (parts.length > 2) return false

    const start = parseInt(parts[0], 10)
    const end = parts.length === 2 ? parseInt(parts[1], 10) : start

    // 验证数字有效性及顺序
    if (isNaN(start) || (parts.length === 2 && isNaN(end))) return false
    if (parts.length === 2 && end < start) return false
  }

  return true
}

export function numberRangeToArray(input: string = ''): number[] {
  // 原始转换逻辑（与之前一致）
  const ranges = input.split(',')
  const result: number[] = []
  for (const range of ranges) {
    const trimmed = range.trim()
    if (trimmed.includes('-')) {
      const [start, end] = trimmed.split('-').map(Number)
      for (let num = start; num <= end; num++) {
        result.push(num)
      }
    } else {
      result.push(Number(trimmed))
    }
  }

  // 使用 Set 去重（自动去除重复项）
  const uniqueResult = [...new Set(result)]

  // 排序（Set 不保证顺序，需显式排序）
  return uniqueResult.sort((a, b) => a - b)
}

export function stringifyFps(fps: number, interlaced: boolean): string {
  return interlaced ? `${fps}I` : `${fps}P`
}

export function parseFps(fpsStr: string): { fps: number; interlaced: boolean } {
  const match = fpsStr.match(/^(\d+\.?\d*)([IP])$/)
  if (!match) {
    throw new Error('Invalid FPS format')
  }
  const fps = parseFloat(match[1])
  const interlaced = match[2] === 'I'
  return { fps, interlaced }
}

export function handleVideoForm(
  form: VideoFormat & { fpsStr: string }
): Partial<VideoFormat> {
  // eslint-disable-next-line @typescript-eslint/no-unused-vars
  const { fpsStr, ...rest } = form
  const result: Partial<VideoFormat> = rest
  if (form.protocol === 'st2110-20' || form.protocol === 'st2110-22') {
    if (form.protocol === 'st2110-20') {
      delete result.compression_format
    }
    delete result.compression_subtype
    delete result.bitrate_bps
    delete result.gop_b_frames
    delete result.gop_length
  }
  if (form.protocol !== 'st2110-22') {
    delete result.compression_ratio
  }
  if (form.compression_format === 'speedhq') {
    delete result.bitrate_bps
    delete result.gop_b_frames
    delete result.gop_length
  }
  if (form.compression_subtype === '无') {
    result.compression_subtype = ''
  }
  return result
}

export function handleAudioForm(form: AudioFormat): Partial<AudioFormat> {
  const result: Partial<AudioFormat> = { ...form }
  if (form.compression_format === 'pcm') {
    delete result.bitrate_bps
  }
  if (form.compression_format === 'aac') {
    delete result.packet_time_us
  }
  return result
}
