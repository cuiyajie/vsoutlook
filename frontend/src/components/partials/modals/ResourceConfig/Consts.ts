export const formats = [
  { key: "1920_1080_1_25_sdr_bt709", value: "1920*1080 50i(sdr/bt709)" },
  { key: "1920_1080_0_50_hlg_bt2020", value: "1920*1080 50p(hdr/bt2020)" },
  { key: "3840_2160_0_50_hlg_bt2020", value: "3840*2160 50i(hdr/bt2020)" },
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
  v_src_address: "",
  v_dst_address: "",
}

export const ipstream = {
  ...ipstream_video,
  a_src_address: "",
  a_dst_address: ""
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
  pips_number: 4
};

export const pip_params = {
  pip_name: '',
  pip_video_index: 0
};

export const switch_key_params = {
  ...ipstream_video,
  Key_P4_port: 1
}

export const switch_fill_params = {
  ...ipstream_video,
  fill_P4_port: 3
}

export const switch_video_params = {
  in_sw_index: 1,
  in_sw_displayname: "cam1"
}

export const switch_out_params = {
  ...ipstream_video,
  v_p4inport: 1,
  v_p4outport: 2
}

export const switch_keyfill_params = {
  key_src_address: "",
  key_dst_address: "",
  key_p4_port: 1,
  fill_src_address: "",
  fill_dst_address: "",
  fill_p4_port: 1
}

export const nmos_config = {
  nmos_enable: true,
  rds_server_url: "192.168.1.112:8010",
  name: "61stream1",
  http_port: 9008,
  log_level:  0,
  log_path:  "/opt/vsomediasoftware/logs/",
  domain:  "local."
}

export const ssm_address = {
  index: 0,
	from: "232.0.0.0",
  to: "232.255.255.255"
}

export const auth_service = {
  index: 0,
	ip: "232.0.0.0",
  port: 8000
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

export const def_switch_input_params = () => ({
  ...base_config,
  videoformat: { ...videoformat },
  input_key_params: {
    ipstream_master: { ...switch_key_params },
    ipstream_backup: { ...switch_key_params },
  },
  input_fill_params: {
    ipstream_master: { ...switch_fill_params },
    ipstream_backup: { ...switch_fill_params },
  },
  input_video_params: []
})

export const def_switch_input_video_params = () => ({
  ipstream_master: { ...ipstream_video, v_P4_port: 1 },
  ipstream_backup: { ...ipstream_video, v_P4_port: 2 },
  ...switch_video_params
})

export const def_switch_input_bus_params = () => ({
  video_bus_master: [{ ...ipstream_video, v_P4_port: 1 }, { ...ipstream_video, v_P4_port: 1 }],
  video_bus_backup: [{ ...ipstream_video, v_P4_port: 2 }, { ...ipstream_video, v_P4_port: 2 }],
  keyfill_bus_master: {
    ...switch_keyfill_params
  },
  keyfill_bus_backup: {
    ...switch_keyfill_params
  }
})

export const def_switch_output_params = () => ({
  v_protocol: "",
  ipstream_master: { ...switch_out_params },
  ipstream_backup: { ...switch_out_params },
  videoformat: { ...videoformat },
})

export type CodecInputType = ReturnType<typeof def_codec_input>
export type CodecOutputType = ReturnType<typeof def_codec_output>
export type UdxInputType = ReturnType<typeof def_udx_input>
export type UdxOutputParamsType = ReturnType<typeof def_udx_output_params>
export type MVInputParamsType = ReturnType<typeof def_mv_input_params>
export type MVOutputParamsType = ReturnType<typeof def_mv_output_params>
export type SwitchInputParamsType = ReturnType<typeof def_switch_input_params>
export type SwitchInputBusParamsType = ReturnType<typeof def_switch_input_bus_params>
export type SwitchInputVideoParamsType = ReturnType<typeof def_switch_input_video_params>
export type SwitchOutputParamsType = ReturnType<typeof def_switch_output_params>
export type NMosConfigType = typeof nmos_config
export type SSMAddressType = typeof ssm_address
export type AuthServiceType = typeof auth_service
