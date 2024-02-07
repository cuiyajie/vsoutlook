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

  async function $fetchList() {
    // nodes.value =  nodeMock.nodes
    const res = await $fetch('/api/cluster/nodes')
    if (res && res.code === 0) {
      nodes.value = res.data.map((n: any) => ({
        id: n.nodeName,
        ip: n.nodeIP
      }))
    }
  }

  return {
    nodes,
    $fetchList,
    $startQuery,
    $stopQuery
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
