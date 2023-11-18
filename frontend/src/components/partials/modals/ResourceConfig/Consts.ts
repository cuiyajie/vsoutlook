export const formats = [
  { key: "1920_1080_隔行_25_SDR_BT709", value: "1920*1080 50I(SDR/BT709)" },
  { key: "1920_1080_逐行_50_HLG_BT2020", value: "1920*1080 50P(HDR/BT2020)" },
  { key: "3804_2160_逐行_50_HLG_BT2020", value: "3840*2160 50I(HDR/BT2020)" },
];

export const formatKeys = formats.map(f => f.key)

export const formatMap = formats.reduce((acc, cur) => {
  acc[cur.key] = cur.value;
  return acc;
}, {} as Record<string, string>);

export const vProtocols = ["ST2110-20", "ST2110-22"];

export const defs = {
  format: formatKeys[0],
  vProtocol: vProtocols[0],
  aProtocol: "ST2110-30",
}

export const videoConfig1 = {
  V_Protocol: defs.format,
  V_M_Address: "",
  V_B_Address: "",
  V_Width: "",
  V_Height: "",
  V_Interlaced: "",
  V_FPS: "",
  V_Gamma: "",
  V_ColorGamut: "",
}

export const videoConfig2 = {
  ...videoConfig1,
  V_DecFormat: "",
  V_CompressionRatio: ""
};

export const videoConfig3 = {
  V_M_Address: '',
  V_M_P4_Port: 0,
  V_B_Address: '',
  V_B_P4_Port: 0,
  SW_Index: 0,
  SW_DisPlayName: ''
}

export const videoConfig = {
  ...videoConfig2,
  A_M_Address: "",
  A_B_Address: ""
};

export const audioConfig = {
  A_Protocol: defs.aProtocol,
  A_Channels_Number: 16,
  A_Bits: 24,
  A_Frequency: 48000
};

export const screenConfig = {
  MV_Template: '',
  PIPs_Number: 4
};

export const screenPipConfig = {
  Name: '',
  Video_Index: 0
};

export type Config1 = { '2022-7': boolean } & typeof videoConfig & typeof audioConfig

export type Config2 = typeof videoConfig & typeof audioConfig

export type Config3 = typeof videoConfig1 & typeof screenConfig

export type Config4 = { '2022-7': boolean } & typeof videoConfig2

export type Config5 = typeof videoConfig3

export type Config6 = typeof videoConfig
