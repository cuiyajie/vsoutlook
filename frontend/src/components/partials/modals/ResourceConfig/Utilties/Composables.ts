import { type WatchSource } from 'vue'
import { type NicDetail, def_nic_detail } from './Consts_V1'
import { type PlayerParams } from './Consts_V1'
import {
  def_switch_panel_params,
  def_tally_config,
  type SwitchPanel,
  type TallyParams,
} from '../SwitchShare/Consts'

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

export function useSmpteParams(params: Ref<PlayerParams>, nicDetails: Ref<NicDetail[]>) {
  const sync = (p: PlayerParams, nic: NicDetail) => {
    if (nic) {
      p.smpte_params.ipstream_master.v_src_address = nic['2110-7_m_local_ip']
      p.smpte_params.ipstream_master.a_src_address = nic['2110-7_m_local_ip']
      p.smpte_params.ipstream_backup.v_src_address = nic['2110-7_b_local_ip']
      p.smpte_params.ipstream_backup.a_src_address = nic['2110-7_b_local_ip']
    }
  }
  watch(
    () => params.value.smpte_params.nic_index,
    (nv) => {
      sync(params.value, nicDetails.value[nv])
    },
    { immediate: true }
  )
  watch(
    nicDetails,
    (nv) => {
      sync(params.value, nv[params.value.smpte_params.nic_index])
    },
    { immediate: true, deep: true }
  )
}

export function useUsedFormat() {
  const formatEnumMap = ref<Record<string, number>>({})
  const formatEnum = computed(() => Object.keys(formatEnumMap.value))
  function onSelected(name: string) {
    formatEnumMap.value[name] = (formatEnumMap.value[name] || 0) + 1
  }

  function onUnselected(name: string) {
    if (formatEnumMap.value[name] === 1) {
      delete formatEnumMap.value[name]
    } else {
      formatEnumMap.value[name] -= 1
    }
  }

  return [formatEnum, onSelected, onUnselected] as const
}

export function useTally(level: WatchSource<number>, tally: Ref<TallyParams>) {
  watch(
    level,
    (nv) => {
      const len = tally.value.length
      if (nv > len) {
        tally.value = Array.from(
          { length: nv },
          (_, i) => tally.value[i] || def_tally_config(i)
        )
      } else if (nv < len) {
        tally.value = tally.value.slice(0, nv)
      }
    },
    { immediate: true }
  )
}

export function useSwitchPanel(level: WatchSource<number>, panel: Ref<SwitchPanel>) {
  watch(
    level,
    (nv) => {
      const len = panel.value.length
      if (nv > len) {
        panel.value = Array.from(
          { length: nv },
          (_, i) => panel.value[i] || def_switch_panel_params(i)
        )
      } else if (nv < len) {
        panel.value = panel.value.slice(0, nv)
      }
    },
    { immediate: true }
  )
}

export function useSwitchBus() {
  const busSet = ref<Set<string>>(new Set())
  function onSelected(id: string) {
    busSet.value.add(id)
  }
  function onUnselected(id: string) {
    busSet.value.delete(id)
  }
  return [busSet, onSelected, onUnselected] as const
}
