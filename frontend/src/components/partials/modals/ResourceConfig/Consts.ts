export const formats = [
  { key: "1920_1080_1_25_sdr_bt709", value: "1920*1080 50i(sdr/bt709)" },
  { key: "1920_1080_0_50_hlg_bt2020", value: "1920*1080 50p(hdr/bt2020)" },
  { key: "3840_2160_0_50_hlg_bt2020", value: "3840*2160 50p(hdr/bt2020)" },
];

export const formatKeys = formats.map(f => f.key)

export const formatMap = formats.reduce((acc, cur) => {
  acc[cur.key] = cur.value;
  return acc;
}, {} as Record<string, string>);

export const v_protocols = ["st2110-20", "st2110-22"];
export const v_compression_format = "jpeg-xs"
export const v_width = ["1920", "3840"]
export const v_compression_ratio = ["5:1", "8:1"]
export const val_codec = ["encoder", "decoder"]
export const val_udx = ["upscale", "downscale"]
export const val_switch_keytype = ["external_key", "internal_key"]

export const defs = {
  format: formatKeys[0],
  v_protocol: v_protocols[0],
  a_protocol: "st2110-30",
}

export const base_config = {
  "g_2022-7": true,
  v_protocol: defs.v_protocol,
}

export const global_config = {
  "g_2022-7": true
}

export const ipstream_video = {
  v_src_address: "10.1.1.61:30000",
  v_dst_address: "232.0.61.1:30000",
}

export const ipstream = {
  ...ipstream_video,
  a_src_address: "10.1.1.61:30000",
  a_dst_address: "232.0.61.1:30000"
}

export const compression_format = {
  v_compression_format: null,
  v_compression_ratio: null
}

export const base_video_format = {
  v_width: "",
  v_height: "",
  v_interlaced: "",
  v_fps: "",
  v_gamma: "",
  v_gamut: "",
}

export const videoformat = {
  ...base_video_format,
  ...compression_format
}

export const audioformat = {
  a_channels_number: 16,
  a_bits: 24,
  a_frequency: 48000
};

export const mv_config = {
  mv_template: '',
  pips_number: 4,
  screenindex: 0,
};

export const pip_params = {
  pip_name: '',
  pip_video_index: 0,
  tallyindex: 0
};

export const switch_key = {
  key_number: 2,
  ext_key_number: 1
}

export const switch_key_parmas_externnal = {
  keytype: 'external_key',
  ext_key_index: 0
}

export const switch_key_params_internal = {
  keytype: 'internal_key',
  int_key: {
    name: "",
    top_x: 0,
    top_y: 0,
    width: 3840,
    height: 2160
	}
}

export const switch_key_params = {
  ...ipstream_video,
  key_p4_port: 1
}

export const switch_fill_params = {
  ...ipstream_video,
  fill_p4_port: 3
}

export const switch_video_params = {
  sw_index: 1,
  sw_displayname: "cam1"
}

export const switch_out_params = {
  ...ipstream_video,
  v_p4inport: 1,
  v_p4outport: 2
}

export const bcswitch_out_params = {
  ...ipstream
}

export const switch_keyfill_params = {
  key_src_address: "",
  key_dst_address: "",
  key_p4_port: 1,
  fill_src_address: "",
  fill_dst_address: "",
  fill_p4_port: 1
}

export const bcswitch_keyfill_params = {
  key_src_address: "",
  key_dst_address: "",
  fill_src_address: "",
  fill_dst_address: ""
}

export const nmos_config = {
  nmos_enable: true,
  rds_server_url: "192.168.1.112:8010",
  name: "61stream1",
  http_port: 9008,
  log_level:  0,
  log_path:  "/opt/vsomediasoftware/logs/",
  domain:  "local.",
  host_addresses: "192.168.1.64"
}

export const ssm_address = {
  index: 0,
	from: "232.0.0.0",
  to: "232.255.255.255"
}

export const def_codec_input = () => ({
  ...base_config,
  ipstream_master: { ...ipstream },
  ipstream_backup: { ...ipstream },
  videoformat: { ...videoformat },
  audioformat: { ...audioformat },
})

export const def_codec_output = () => ({
  ...global_config,
  params: {
    v_protocol: "",
    ipstream_master: { ...ipstream },
    ipstream_backup: { ...ipstream },
    videoformat: { ...videoformat },
    audioformat: { ...audioformat },
  }
})

export const def_udx_input = () => ({
  ...base_config,
  ipstream_master: { ...ipstream },
  ipstream_backup: { ...ipstream },
  videoformat: { ...videoformat },
  audioformat: { ...audioformat },
})

export const def_udx_output_params = () => ({
  v_protocol: "",
  ipstream_master: { ...ipstream },
  ipstream_backup: { ...ipstream },
  videoformat: { ...videoformat },
  audioformat: { ...audioformat },
})

export const def_mv_input_params = () => ({
  v_protocol: "",
  ipstream_master: { ...ipstream },
  ipstream_backup: { ...ipstream },
  videoformat: { ...videoformat },
  audioformat: { ...audioformat },
})

export const def_mv_output_params = () => ({
  ...mv_config,
  v_protocol: v_protocols[0],
  pip_params: [],
  ipstream_master: { ...ipstream_video },
  ipstream_backup: { ...ipstream_video },
  videoformat: { ...base_video_format }
})

export const def_mv_pip_params = (idx: number) => ({
  ...pip_params,
  pip_name: `cam${idx + 1}`,
  pip_video_index: idx,
  tallyindex: idx
})

export const def_switch_key = () => ({
  ...switch_key,
  key_params: []
})

export const def_switch_key_params = () => ([
  { ...switch_key_parmas_externnal },
  { ...switch_key_params_internal }
])

export const def_switch_external_key_params = () => ({
  ...switch_key_parmas_externnal
})

export const def_switch_internal_key_params = () => ({
  ...switch_key_params_internal,
  int_key: { ...switch_key_params_internal.int_key }
})

export const def_switch_input_params = () => ({
  ...base_config,
  videoformat: { ...videoformat },
  input_key_params: [],
  input_fill_params: [],
  input_video_params: []
})

export const def_bcswitch_input_params = () => ({
  ...base_config,
  videoformat: { ...videoformat },
  audioformat: { ...audioformat },
  input_video_params: []
})

export const def_switch_input_key_params = () => ({
  ipstream_master: { ...switch_key_params },
  ipstream_backup: { ...switch_key_params },
})

export const def_switch_input_fill_params = () => ({
  ipstream_master: { ...switch_fill_params },
  ipstream_backup: { ...switch_fill_params },
})

export const def_switch_input_video_params = (idx: number) => ({
  ipstream_master: { ...ipstream_video, v_p4_port: 1 },
  ipstream_backup: { ...ipstream_video, v_p4_port: 2 },
  ...switch_video_params,
  sw_index: idx + 1,
  sw_displayname: `cam${idx + 1}`
})

export const def_bcswitch_input_video_params = (idx: number) => ({
  ipstream_master: { ...ipstream },
  ipstream_backup: { ...ipstream },
  ...switch_video_params,
  sw_index: idx + 1,
  sw_displayname: `cam${idx + 1}`
})

export const def_switch_bus_params = () => ({
  video_bus_master: [{ ...ipstream_video, v_p4_port: 1 }, { ...ipstream_video, v_p4_port: 1 }],
  video_bus_backup: [{ ...ipstream_video, v_p4_port: 2 }, { ...ipstream_video, v_p4_port: 2 }],
  keyfill_bus_master: [],
  keyfill_bus_backup: []
})

export const def_bcswitch_bus_params = () => ({
  video_bus_master: [{ ...ipstream }, { ...ipstream }],
  video_bus_backup: [{ ...ipstream }, { ...ipstream }],
  keyfill_bus_master: [],
  keyfill_bus_backup: []
})

export const def_switch_bus_keyfill_params = () => ({
  ...switch_keyfill_params
})

export const def_bcswitch_bus_keyfill_params = () => ({
  ...bcswitch_keyfill_params
})

export const def_switch_output_params = () => ({
  v_protocol: "",
  ipstream_master: { ...switch_out_params },
  ipstream_backup: { ...switch_out_params },
  videoformat: { ...videoformat },
})

export const def_bcswitch_output_params = () => ({
  v_protocol: "",
  ipstream_master: { ...bcswitch_out_params },
  ipstream_backup: { ...bcswitch_out_params },
  videoformat: { ...videoformat },
  audioformat: { ...audioformat },
})

export const def_endswitch_input_params = () => ({
  ...base_config,
  videoformat: { ...videoformat },
  audioformat: { ...audioformat },
  input_video_params: []
})

export const def_endswitch_input_video_params = (idx: number) => ({
  ipstream_master: { ...ipstream },
  ipstream_backup: { ...ipstream },
  sw_displayname: `cam${idx + 1}`
})

export const def_endswitch_output_params = () => ({
  v_protocol: "",
  ipstream_master: { ...bcswitch_out_params },
  ipstream_backup: { ...bcswitch_out_params },
  videoformat: { ...videoformat },
  audioformat: { ...audioformat },
})

export type CodecInputType = ReturnType<typeof def_codec_input>
export type CodecOutputType = ReturnType<typeof def_codec_output>
export type UdxInputType = ReturnType<typeof def_udx_input>
export type UdxOutputParamsType = ReturnType<typeof def_udx_output_params>
export type MVInputParamsType = ReturnType<typeof def_mv_input_params>
export type MVOutputParamsType = ReturnType<typeof def_mv_output_params>
export type SwitchKeyType = typeof switch_key
export type SwitchKeyParamsType = ReturnType<typeof def_switch_key_params>[0]
export type SwitchInputParamsType = ReturnType<typeof def_switch_input_params>
export type BCSwitchInputParamsType = ReturnType<typeof def_bcswitch_input_params>
export type SwitchInputKeyParamsType = ReturnType<typeof def_switch_input_key_params>
export type SwitchInputFillParamsType = ReturnType<typeof def_switch_input_fill_params>
export type SwitchInputVideoParamsType = ReturnType<typeof def_switch_input_video_params>
export type BCSwitchInputVideoParamsType = ReturnType<typeof def_bcswitch_input_video_params>
export type SwitchBusParamsType = ReturnType<typeof def_switch_bus_params>
export type BCSwitchBusParamsType = ReturnType<typeof def_bcswitch_bus_params>
export type SwitchBusKeyfillParamsType = ReturnType<typeof def_switch_bus_keyfill_params>
export type BCSwitchBusKeyfillParamsType = ReturnType<typeof def_bcswitch_bus_keyfill_params>
export type SwitchOutputParamsType = ReturnType<typeof def_switch_output_params>
export type BCSwitchOutputParamsType = ReturnType<typeof def_bcswitch_output_params>
export type EndSwitchInputParamsType = ReturnType<typeof def_endswitch_input_params>
export type EndSwitchInputVideoParamsType = ReturnType<typeof def_endswitch_input_video_params>
export type EndSwitchOutputParamsType = ReturnType<typeof def_endswitch_output_params>
export type SwitchMVPipParamsType = typeof pip_params
export type NMosConfigType = typeof nmos_config
export type SSMAddressType = typeof ssm_address
