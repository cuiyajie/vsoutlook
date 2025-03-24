/**
 * This is a store that hold right panel state on sidebar layouts
 * It could be one of the SidebarId
 *
 * We can import and set activeSidebar (or use toggleSidebar) anywhere in our project
 * @see /src/components/navigation/desktop/sidebar/SidebarColorCurved.vue
 * @see /src/pages/sidebar-blank-page-1.vue
 */

import { ref } from 'vue'
import { acceptHMRUpdate, defineStore } from 'pinia'
import { useFetch } from '@src/composable/useFetch'
// import * as nodeMock from "@src/data/node-mock"

const Timeout = 15 * 1000
export const useClustNode = defineStore('clustNode', () => {
  const nodes = ref<ClustNode[]>([])
  const nmosNodes = ref<NMosNode[]>([])
  const $fetch = useFetch()
  let timer: any = null

  function $startQuery() {
    timer = setInterval(() => {
      $fetchList()
    }, Timeout)
  }

  function $stopQuery() {
    if (timer) {
      clearInterval(timer)
      timer = null
    }
  }

  async function $getNodeNics(id: string) {
    const res = await $fetch('/api/cluster/nics', {
      body: { id },
    })
    if (res?.code === 0) {
      const tnode = nodes.value.find((node) => node.id === id)
      if (tnode) {
        tnode.nics = (res.data || []) as NicInfo[]
      }
      return tnode
    }
    return null
  }

  async function $createNic(nic: Omit<NicInfo, 'id' | 'nodeId'>, nodeId: string) {
    const res = await $fetch('/api/cluster/nic.create', {
      body: { nodeId, ...nic },
    })
    if (res?.code === 0) {
      const tnode = nodes.value.find((node) => node.id === nodeId)
      if (tnode) {
        tnode.nics.push(res.data as NicInfo)
      }
    }
    return res
  }

  async function $updateNic(nic: NicInfo) {
    const res = await $fetch('/api/cluster/nic.update', {
      body: { ...nic },
    })
    if (res?.code === 0) {
      const tnode = nodes.value.find((node) => node.id === nic.nodeId)
      if (tnode) {
        const tnic = tnode.nics.find((n) => n.id === nic.id)
        if (tnic) {
          Object.assign(tnic, res.data as NicInfo)
        }
      }
    }
    return res
  }

  async function $deleteNic(id: string, nodeId: string) {
    const res = await $fetch('/api/cluster/nic.delete', {
      body: { id },
    })
    if (res?.code === 0) {
      const tnode = nodes.value.find((node) => node.id === nodeId)
      if (tnode) {
        tnode.nics = tnode.nics.filter((n) => n.id !== id)
      }
    }
    return res
  }

  async function $fetchById(id: string): Promise<ClustNode> {
    const res1 = await $fetch('/api/cluster/node.detail', {
      body: { id },
    })
    if (res1 && res1.code === 0) {
      const { node, running, stopped } = res1.data
      return {
        id: node?.nodeName,
        ip: node?.nodeIP,
        nics: [],
        allocatable: node?.allocatable,
        allocated: node?.allocated,
        running,
        stopped,
      }
    }
    return {
      id: '',
      ip: '',
      nics: [],
      allocatable: {} as any,
      allocated: {} as any,
      running: [],
      stopped: [],
    }
  }

  async function $getById(id: string): Promise<ClustNode> {
    const node = await $fetchById(id)
    const tnode = nodes.value.find((n) => n.id === id)
    if (tnode) {
      Object.assign(tnode, node)
    }
    return tnode as ClustNode
  }

  async function $fetchList() {
    // nodes.value =  nodeMock.nodes
    const res = await $fetch('/api/cluster/nodes')
    if (res && res.code === 0) {
      nodes.value = await Promise.all(
        res.data.map(async (n: { nodeName: string; nodeIP: string }) => {
          const node = await $fetchById(n.nodeName)
          return Object.assign(node, {
            id: n.nodeName,
            ip: n.nodeIP,
          })
        })
      )
    }
  }

  async function $fetchNMos() {
    const res = await $fetch('/api/cluster/nmos', {
      body: {
        order: 'update',
        limit: 1000,
      },
    })
    if (res && res.code === 0) {
      nmosNodes.value = res.data
      return res.data
    } else {
      nmosNodes.value = []
      return []
    }
  }

  return {
    nodes,
    nmosNodes,
    $fetchList,
    $fetchNMos,
    $startQuery,
    $stopQuery,
    $getById,
    $getNodeNics,
    $createNic,
    $updateNic,
    $deleteNic,
  } as const
})

/**
 * Pinia supports Hot Module replacement so you can edit your stores and
 * interact with them directly in your app without reloading the page.
 *
 * @see https://pinia.esm.dev/cookbook/hot-module-replacement.html
 * @see https://vitejs.dev/guide/api-hmr.html
 */
if (import.meta.hot) {
  import.meta.hot.accept(acceptHMRUpdate(useClustNode, import.meta.hot))
}
