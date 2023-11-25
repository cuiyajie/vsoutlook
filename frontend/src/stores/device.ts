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
import { useNotyf } from "/@src/composable/useNotyf"
import { useFetch } from "/@src/composable/useFetch"
import * as dic from "/@src/utils/enums-dic"

export const useDevices = defineStore('device', () => {
  const devices = ref<DeviceDetail[]>([])
  const notyf = useNotyf()
  const $fetch  = useFetch()

  async function $fetchList() {
    const res = await $fetch('/api/device/list')
    if (res && res.devices) {
      devices.value = res.devices.map((d: any) => {
        d.updated = new Date(d.updated).toLocaleString('zh')
        d.statusInfo = dic.DeviceStatusDic[d.status]
        return d
      })
    }
  }

  async function $deploy(name: string, tmpl: string, body: any) {
    const res = await $fetch('/api/device/create', {
      body: { name, tmpl, body: JSON.stringify(body) }
    })
    if (res && res.device) {
      return { result: 'success' }
    } else {
      return { result: 'error', message: res.message }
    }
  }

  function config() {
    notyf.dismissAll()
    notyf.info('即将支持，敬请期待')
  }

  function restart() {
    notyf.dismissAll()
    notyf.info('即将支持，敬请期待')
  }

  function shutdown() {
    notyf.dismissAll()
    notyf.info('即将支持，敬请期待')
  }

  function start() {
    notyf.dismissAll()
    notyf.info('即将支持，敬请期待')
  }

  function remove() {
    notyf.dismissAll()
    notyf.info('即将支持，敬请期待')
  }

  function refresh() {
    notyf.dismissAll()
    notyf.info('即将支持，敬请期待')
  }

  function getById(id: string) {
    return devices.value.find(d => d.id === id)
  }

  return {
    devices,
    config,
    restart,
    shutdown,
    start,
    remove,
    refresh,
    getById,
    $fetchList,
    $deploy
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
  import.meta.hot.accept(acceptHMRUpdate(useDevices, import.meta.hot))
}

function useClustFetch() {
throw new Error("Function not implemented.");
}
