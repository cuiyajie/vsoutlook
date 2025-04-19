import merge from 'lodash-es/merge'
import { def_player_params, type PlayerParams } from '../Utilties/Consts_V1'
import { checkPlayerParams, handlePlayerParams, unwrap } from '../Utilties/Utils_V1'
import {
  type SwitchInputKeyParams,
  type SwitchInput,
  type SwitchPanel,
  type SwitchInputVideoParams,
  type SwitchBus,
  type SwitchBusKeyParams,
  type SwitchBusMeParams,
  type SwitchBusLevelParams,
  type SwitchOut,
  type SwitchOutParams,
  type SwitchOutAudioParams,
  type SwitchOutMvParams,
  type TallyParams,
  type KeyFillPlayerParams,
  def_switch_bus,
  def_switch_bus_key_params,
  def_switch_bus_level_input_params,
  def_switch_bus_level_params,
  def_switch_bus_level_sub_params,
  def_switch_input,
  def_switch_input_key_params,
  def_switch_input_video_params,
  def_switch_out,
  def_switch_out_params,
  def_switch_bus_me_params,
  def_tally,
  def_tally_screen,
  def_switch_panel_params,
  def_switch_panel_url,
} from './Consts'
import type bcswtData from '@src/data/vscomponent/bcswitch.json'

export function handleTally(params: TallyParams) {
  return params.map((p, pidx) => {
    const screens = p.screens.map((s) => {
      const { pvw_checked, pgm_checked, ...rest } = s
      if (!pvw_checked) {
        delete (rest as any).pvw_tally_index
      }
      if (!pgm_checked) {
        delete (rest as any).pgm_tally_index
      }
      rest.signal_config = rest.signal_config.map((sc, scidx) => ({
        ...sc,
        index: scidx,
      }))
      return rest
    })
    return {
      screens,
      level: pidx,
    }
  })
}

export function handleSwitchPanel(params: SwitchPanel) {
  return params.map((p, pidx) => ({
    level_index: pidx,
    panel_type: p.panel_type,
    panel_url: p.panel_url.map((u, uidx) => ({
      subpanel_index: uidx,
      url: u.url,
    })),
  }))
}

function handleKeyFillParams(params: KeyFillPlayerParams, vfs: VideoFormat[]) {
  const { videoformat_name, smpte_params, stream_params, ...rest } =
    params
  const vf = vfs.find((v) => v.name === videoformat_name)
  const result = { videoformat_name, ...rest } as any
  if (!vf?.protocol) return result
  const showParams = ['st2110-20', 'st2110-22'].includes(vf.protocol)
  if (showParams) {
    result.smpte_params = smpte_params
  } else {
    result.stream_params = stream_params
  }
  return result
}

export function handleSwitchInputKey(params: SwitchInputKeyParams[], vfs: VideoFormat[]) {
  return params.map((p, pidx) => {
    let result: any = {}
    if (p.key_type === 'ext_key') {
      result = handleKeyFillParams(p, vfs)
      delete result.file_name
      delete result.url
    }
    if (p.key_type === 'int_key') {
      result.file_name = p.file_name
    }
    if (p.key_type === 'h5_key') {
      result.url = p.url
    }
    return {
      key_type: p.key_type,
      signal_id: p.signal_id,
      signal_name: p.signal_name,
      videoformat_name: p.videoformat_name,
      ...result,
      index: pidx,
    }
  })
}

export function handleSwitchInputVideo(
  params: SwitchInputVideoParams[],
  vfs: VideoFormat[]
) {
  return params.map((p, pidx) => {
    const result = handlePlayerParams(p, vfs)
    if (result.smpte_params) {
      delete result.smpte_params.nic_index
    }
    return {
      ...result,
      index: pidx,
    }
  })
}

export function handleSwitchInput(params: SwitchInput, vfs: VideoFormat[]) {
  return {
    key: handleSwitchInputKey(params.key, vfs),
    video: handleSwitchInputVideo(params.video, vfs),
  }
}

export function handleBusKey(params: SwitchBusKeyParams[], usedSignalType: number) {
  return params.map((p, pidx) => {
    const busInput: any = { ...p.bus_input }
    if (usedSignalType === 1) {
      delete busInput.nic_index
    }
    const keyParams: any = { keytype: p.key_params.keytype }
    if (p.key_params.keytype !== 'luma_key') {
      Object.assign(keyParams, p.key_params)
    }
    return {
      index: pidx,
      bus_id: p.bus_id,
      bus_name: p.bus_name,
      bus_input: busInput,
      key_params: keyParams,
      out_signal: p.out_signal,
    }
  })
}

export function handleBusMe(params: SwitchBusMeParams[]) {
  return params.map((p, pidx) => ({
    ...p,
    index: pidx,
  }))
}

export function handleBusLevel(params: SwitchBusLevelParams[]) {
  return params.map((p, pidx) => {
    const { bus_input, pgm_bus, pvw_bus } = p
    const pgm_sub_bus = pgm_bus.sub_bus.map((s, sidx) => ({
      bus_type: s.bus_type,
      bus_id: s.bus_id,
      index: sidx,
    }))
    const pvw_sub_bus = pvw_bus.sub_bus.map((s, sidx) => ({
      bus_type: s.bus_type,
      bus_id: s.bus_id,
      index: sidx,
    }))
    return {
      level: pidx,
      bus_input: {
        ...bus_input,
        input_list: bus_input.input_list.map((b, bidx) => ({
          ...b,
          index: bidx,
          input_index: bidx + 1,
        }))
      },
      pgm_bus: {
        ...pgm_bus,
        sub_bus: pgm_sub_bus,
      },
      pvw_bus: {
        ...pvw_bus,
        sub_bus: pvw_sub_bus,
      },
    }
  })
}

export function handleSwitchBus(
  params: SwitchBus,
  usedSignalType: number,
  category: string
) {
  if (category === 'bcswt') {
    return {
      key_bus: handleBusKey(params.key_bus, usedSignalType),
      level_bus: handleBusLevel(params.level_bus),
    }
  } else {
    return {
      key_bus: handleBusKey(params.key_bus, usedSignalType),
      me_bus: handleBusMe(params.me_bus),
      level_bus: handleBusLevel(params.level_bus),
    }
  }
}

function omitAudioParams<
  T extends {
    ipstream_master: PlayerParams['smpte_params']['ipstream_master']
    ipstream_backup: PlayerParams['smpte_params']['ipstream_backup']
  },
>(params: T) {
  return {
    ...params,
    ipstream_master: {
      v_dst_address: params.ipstream_master.v_dst_address,
      v_src_address: params.ipstream_master.v_src_address,
    },
    ipstream_backup: {
      v_dst_address: params.ipstream_backup.v_dst_address,
      v_src_address: params.ipstream_backup.v_src_address,
    },
  }
}

export function handleOutParams(
  params: SwitchOutParams[],
  vfs: VideoFormat[],
  audioMode: number,
  category: string
) {
  return params.map((p, pidx) => {
    const result = handlePlayerParams(p, vfs)
    if (audioMode === 0) {
      delete result.audioformat_name
      if (result.smpte_params) {
        result.smpte_params = omitAudioParams(result.smpte_params)
      }
    }
    if (category === 'bcswt' || audioMode === 0 || !result.mapping_checked) {
      delete result.audio_mapping_name
    }
    delete result.mapping_checked
    return {
      ...result,
      index: pidx,
    }
  })
}

function filterAudioParams<
  T extends {
    ipstream_master: PlayerParams['smpte_params']['ipstream_master']
    ipstream_backup: PlayerParams['smpte_params']['ipstream_backup']
  },
>(params: T) {
  return {
    ...params,
    ipstream_master: {
      a_dst_address: params.ipstream_master.a_dst_address,
      a_src_address: params.ipstream_master.a_src_address,
    },
    ipstream_backup: {
      a_dst_address: params.ipstream_backup.a_dst_address,
      a_src_address: params.ipstream_backup.a_src_address,
    },
  }
}

export function handleAudioOutputParams(
  params: SwitchOutAudioParams,
  usedSignalType: number,
  showSmpte = false
) {
  return {
    src_signal_name: params.src_signal_name,
    pgm_src_signal_id: params.pgm_src_signal_id,
    ...(usedSignalType !== 1 ? { in_nic_index: params.in_nic_index } : {}),
    audioformat_name: params.audioformat_name,
    ...(showSmpte
      ? {
          smpte_params: filterAudioParams({
            ipstream_backup: params.smpte_params.ipstream_backup,
            ipstream_master: params.smpte_params.ipstream_master,
          }),
        }
      : {}),
  }
}

export function handleMVOutputParams(
  params: SwitchOutMvParams,
  vfs: VideoFormat[],
  audioMode: number
) {
  if (!params.mv_is_open) return undefined
  const result = handlePlayerParams(params, vfs)
  return {
    ...result,
    smpte_params:
      audioMode === 1 ? result.smpte_params : filterAudioParams(result.smpte_params),
  }
}

export function handleSwitchOut(
  params: SwitchOut,
  vfs: VideoFormat[],
  audioMode: number,
  usedSignalType: number,
  category: string
) {
  let showSmpte = false
  params.out_params.forEach((outParam) => {
    if (outParam.source_signal.signal_name === 'level0-pvwout-final') {
      showSmpte = true
    }
  })
  return {
    out_number: params.out_number,
    out_params: handleOutParams(params.out_params, vfs, audioMode, category),
    ...(audioMode === 2
      ? {
          audio_output_params: handleAudioOutputParams(
            params.audio_output_params,
            usedSignalType,
            showSmpte
          ),
        }
      : {}),
    mv_output_params: handleMVOutputParams(params.mv_output_params, vfs, audioMode),
  }
}

export function checkAudioMapping(params: any, afs: AudioMapping[]) {
  const af = afs.find((a) => a.name === params.audio_mapping_name)
  if (!af) {
    params.audio_mapping_name = ''
  }
  return params
}

export function checkKeyFillParams(params: any, vfs: VideoFormat[]) {
  const vf = vfs.find((v) => v.name === params.videoformat_name)
  if (!vf) {
    params.videoformat_name = ''
  }
  return params
}

export function checkSwitchData(
  data: typeof bcswtData,
  vfs: VideoFormat[],
  afs: AudioFormat[],
  afsM: AudioMapping[]
) {
  const tallyData = data.tally_config
  let _tally = def_tally(data.level)
  if (data.tally_config && data.tally_notify) {
    _tally = Array.from({ length: tallyData.length }, (_, idx) => {
      const screens = Array.from(
        { length: tallyData[idx].screens.length },
        (_, screenIdx) => {
          const screenData = tallyData[idx].screens[screenIdx]
          const screen = merge(def_tally_screen(), screenData)
          screen.pgm_checked = screenData.pgm_tally_index !== undefined
          screen.pvw_checked = screenData.pvw_tally_index !== undefined
          return screen
        }
      )
      return {
        level: tallyData[idx].level,
        screens,
      }
    })
  }
  const panelData = data.panel_params
  const _panel = Array.from({ length: panelData.length }, (_, idx) => {
    const _panelAt = merge(def_switch_panel_params(idx), panelData[idx])
    const panelUrl = panelData[idx].panel_url
    _panelAt.panel_url = Array.from({ length: panelUrl.length }, (_, i) =>
      merge(def_switch_panel_url(i), panelUrl[i])
    )
    return _panelAt
  })
  const inputData = unwrap(data.input, 'in_')
  const _input = def_switch_input()
  _input.key = Array.from({ length: inputData.key.length }, (_, idx) => {
    const _inputKeyAt = checkKeyFillParams(inputData.key[idx], vfs)
    return merge(def_switch_input_key_params(idx), _inputKeyAt)
  })
  _input.video = Array.from({ length: inputData.video.length }, (_, idx) => {
    const _inputVideoAt = checkPlayerParams(inputData.video[idx], vfs, afs)
    return merge(def_switch_input_video_params(idx), _inputVideoAt)
  })
  const busData = data.bus
  const _bus = def_switch_bus(data.level)
  _bus.key_bus = Array.from({ length: busData.key_bus.length }, (_, idx) => {
    return merge(def_switch_bus_key_params(idx), busData.key_bus[idx])
  })
  if (busData.me_bus) {
    _bus.me_bus = Array.from({ length: busData.me_bus.length }, (_, idx) => {
      return merge(def_switch_bus_me_params(idx), busData.me_bus?.[idx])
    })
  } else {
    _bus.me_bus = []
  }
  _bus.level_bus = Array.from({ length: busData.level_bus.length }, (_, lidx) => {
    const levelBusData = busData.level_bus[lidx] as any
    const levelBus = merge(
      def_switch_bus_level_params(lidx, data.level),
      levelBusData
    ) as SwitchBusLevelParams
    levelBus.bus_input.input_list = Array.from(
      { length: levelBusData.bus_input.input_list.length },
      (_, idx) => {
        return merge(def_switch_bus_level_input_params(idx), levelBusData.bus_input.input_list[idx])
      }
    )
    levelBus.pgm_bus.sub_bus = Array.from(
      { length: levelBusData.pgm_bus.sub_bus.length },
      (subBusData: any, idx) => {
        return merge(def_switch_bus_level_sub_params(idx), subBusData)
      }
    )
    levelBus.pvw_bus.sub_bus = Array.from(
      { length: levelBusData.pvw_bus.sub_bus.length },
      (subBusData: any, idx) => {
        return merge(def_switch_bus_level_sub_params(idx), subBusData)
      }
    )
    return levelBus
  })
  const outData = unwrap(data.output, 'out_')
  const _out = merge(def_switch_out(data.level), outData) as SwitchOut
  _out.out_params = Array.from({ length: _out.out_number }, (_, idx) => {
    let result = checkPlayerParams(outData.out_params[idx], vfs, afs)
    result = checkAudioMapping(result, afsM)
    return merge(def_switch_out_params(idx), result, {
      mapping_checked: result.audio_mapping_name !== '',
    })
  })
  _out.audio_output_params = merge(def_player_params(), outData.audio_output_params)
  if (outData.mv_output_params) {
    _out.mv_output_params = merge(def_player_params(), outData.mv_output_params)
  }

  return {
    tally: _tally,
    panel: _panel,
    input: _input,
    bus: _bus,
    out: _out,
  }
}
