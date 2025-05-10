import { type MgInputParams, type MgOutputItemParams } from './Consts'
import { handlePlayerParams } from '../Utilties/Utils_V1'

export function handleInputParams(input: MgInputParams, vfs: VideoFormat[]) {
  return handlePlayerParams(input, vfs, 'receive')
}

export function handleOutputParams(output: MgOutputItemParams, vfs: VideoFormat[]) {
  const result = handlePlayerParams(output, vfs)
  const hdr = { ...output.hdr_convert_params } as Partial<
    MgOutputItemParams['hdr_convert_params']
  >
  if (hdr.convert_mode === 'lut') {
    delete hdr.dynamic_mode
    delete hdr.luma_gain
  }
  result.hdr_convert_params = hdr
  return result
}
