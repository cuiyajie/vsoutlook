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
import { useFetch } from "/@src/composable/useFetch";
// import * as nodeMock from "/@src/data/node-mock"

const Timeout = 15 * 1000
export const useClustNode = defineStore('clustNode', () => {
  const nodes = ref<ClustNode[]>([])
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

  async function $update(id: string, core: string, dma: string, vf: number) {
    const res = await $fetch('/api/cluster/node.update', {
      body: { id, core, dma, vf }
    })
    if (res?.code === 0) {
      const tnode = nodes.value.find(node => node.id === id)
      if (tnode) {
        tnode.coreList = core
        tnode.dmaList = dma
        tnode.vfCount = vf
      }
    }
    return res
  }

  async function $fetchById(id: string) {
    const res1 = await $fetch('/api/cluster/node.detail', {
      body: { id }
    })
    if (res1 && res1.code === 0) {
      const { coreList, dmaList, node, vfCount, running, stopped } = res1.data
      return {
        id: node?.nodeName,
        ip: node?.nodeIP,
        coreList,
        dmaList,
        vfCount,
        allocatable: node?.allocatable,
        allocated: node?.allocated,
        running,
        stopped
      }
    }
    return {
      id: "",
      ip: "",
      coreList: "",
      dmaList: "",
      running: [],
      stopped: []
    }
  }

  async function $getById(id: string) {
    const node = await $fetchById(id)
    const tnode = nodes.value.find(n => n.id === id)
    if (tnode) {
      Object.assign(tnode, node)
    }
  }

  async function $fetchList() {
    // nodes.value =  nodeMock.nodes
    const res = await $fetch('/api/cluster/nodes')
    if (res && res.code === 0) {
      nodes.value = await Promise.all(res.data.map(async (n: { nodeName: string, nodeIP: string }) => {
        const node = await $fetchById(n.nodeName)
        return Object.assign(node, {
          id: n.nodeName,
          ip: n.nodeIP
        })
      }))
    }
  }

  return {
    nodes,
    $fetchList,
    $startQuery,
    $stopQuery,
    $update,
    $getById
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
