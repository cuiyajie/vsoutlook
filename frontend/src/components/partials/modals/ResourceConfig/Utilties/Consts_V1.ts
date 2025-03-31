export const used_signal_types = [
  {
    key: 0,
    label: '仅使用无压缩及浅压缩信号（SMPTE 相关格式信号）',
  },
  {
    key: 1,
    label: '仅使用深压缩信号',
  },
  {
    key: 2,
    label: '同时使用无压缩，浅压缩及深压缩信号',
  },
]

export const def_nic_detail = () => ({
  numa_id: -1,
  '2110-7_m_local_ip': '10.1.1.61',
  nic_name_m: 'ens01',
  '2110-7_m_receive': true,
  '2110-7_m_send': true,
  '2110-7_b_local_ip': '10.1.101.61',
  nic_name_b: 'ens02',
  '2110-7_b_receive': true,
  '2110-7_b_send': true,
  nicIndex: -1,
  id: 'nicid',
})

export const def_player_params = () => ({
  videoformat_name: '',
  audioformat_name: '',
  smpte_params: {
    nic_index: -1,
    ipstream_master: {
      v_src_address: '10.1.1.64:30000',
      v_dst_address: '232.0.64.1:30000',
      a_src_address: '10.1.1.64:30000',
      a_dst_address: '232.1.64.1:30000',
    },
    ipstream_backup: {
      v_src_address: '10.1.101.64:30000',
      v_dst_address: '232.0.164.1:30000',
      a_src_address: '10.1.101.64:30000',
      a_dst_address: '232.1.164.1:30000',
    },
  },
  stream_params: {
    url: 'SRT://192.168.1.64:1234',
    codec_dev: 'gpu',
    codec_dev_index: 0,
    quality: 2,
    video_enc_driver: 0,
    audio_enc_driver: 0,
  },
})

export const codec_dev_types = ['gpu', 'cpu']

export const quality_types = [
  { value: 0, label: '速度最快，质量最差' },
  { value: 1, label: '速度较快，质量较差' },
  { value: 2, label: '速度中等，质量中等' },
  { value: 3, label: '速度较慢，质量较好' },
  { value: 4, label: '速度最慢，质量最好' },
]

export type IndexedParmas<T extends object> = { index: number } & T
export type IndexedType<T> = { index: number; value: T }

export interface IApiParams {
  api_name: string
  service_port: number
  checked: boolean
  label: string
}
export type NicDetail = ReturnType<typeof def_nic_detail>
export type IndexedNicDetail = IndexedParmas<NicDetail>
export type PlayerParams = ReturnType<typeof def_player_params>
