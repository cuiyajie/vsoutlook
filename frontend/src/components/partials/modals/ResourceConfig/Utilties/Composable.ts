import { type NicDetail, def_nic_detail } from './Consts_V1'

export function useNicList(props: { nics: NicInfo[]; requiredment?: TmplRequirement }) {
  const nicDetails = ref<NicDetail[]>([])
  const indexedNicDetails = computed(() => {
    return nicDetails.value
      .filter((nic) => nic.nicIndex >= 0)
      .map((nic, idx) => ({ ...nic, index: idx }))
  })

  watch(
    [() => props.nics, () => props.requiredment],
    () => {
      if (props.requiredment?.nicCount && props.requiredment.nicCount > 0) {
        nicDetails.value = props.nics
          .slice(0, props.requiredment.nicCount)
          .map((nic, idx) => ({
            ...def_nic_detail(),
            nic_name_m: nic.nicNameMain,
            '2110-7_m_receive': nic.receiveMain,
            '2110-7_m_send': nic.sendMain,
            nic_name_b: nic.nicNameBackup,
            '2110-7_b_receive': nic.receiveBackup,
            '2110-7_b_send': nic.sendBackup,
            nicIndex: idx,
            id: nic.id,
          }))
      } else {
        nicDetails.value = []
      }
    },
    { immediate: true }
  )

  return {
    nicDetails,
    indexedNicDetails,
  }
}
