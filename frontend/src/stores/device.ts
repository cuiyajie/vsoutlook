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
import { useRlsFetch } from "../composable/useRlsFetch"
import { valert } from "../utils/dialog"

export const useDevices = defineStore('device', () => {
  const devices = ref<DeviceDetail[]>([])
  const notyf = useNotyf()
  const $fetch  = useFetch()
  const $rfetch = useRlsFetch()

  async function $fetchList() {
    const res = await $fetch('/api/device/list')
    if (res && res.devices) {
      devices.value = res.devices.map((d: any) => {
        d.updated = new Date(Date.parse(d.updated)).toLocaleString('zh')
        d.statusInfo = dic.DeviceStatusDic[d.phase] || dic.DeviceStatusDic['Unavailable']
        return d
      })
    }
  }

  async function $deploy(name: string, tmpl: string, body: any) {
    const res = await $fetch('/api/device/create', {
      body: {
        name,
        tmpl,
        body: JSON.stringify({ params: body })
      }
    })
    if (res && res.device) {
      return { result: 'success' }
    } else {
      return { result: 'error', message: res.message }
    }
  }

  async function $updateConfig(name: string, deviceId: string, body: any) {
    const res = await $fetch('/api/device/update', {
      body: {
        name,
        deviceId,
        body: JSON.stringify({ params: body })
      }
    })
    if (res && res.device) {
      return { result: 'success' }
    } else {
      return { result: 'error', message: res.message }
    }
  }

  async function $remove(id: string) {
    return await $fetch('/api/device/delete', {
      body: { id }
    })
  }

  async function $showContainer(device: DeviceDetail) {
    const res = await $rfetch(`/api/namespaces/default/releases/${device.name}/podPhase`, {
      method: 'GET',
    })
    if (res && res.code === 0) {
      if (typeof res.data === 'object') {
        const podKey = Object.keys(res.data).find(k => k.includes(device.name))
        console.log(podKey)
        if (podKey) {
          valert({
            title: '容器信息',
            content: podKey
          })
        }
      }
    }
  }

  async function $reboot(device: DeviceDetail) {
    const res = await $rfetch(`/api/namespaces/default/releases/${device.name}/podPhase`, {
      method: 'GET',
    })
    if (res && res.code === 0) {
      if (typeof res.data === 'object') {
        const podKey = Object.keys(res.data).find(k => k.includes(device.name))
        if (podKey) {
          const pod = res.data[podKey]
          if (pod === 'Running') {
            const res = await $rfetch(`/api/namespaces/default/pods/${podKey}/delete`, {
              method: 'GET',
            }).catch(() => {
              return { code: 500, error: "重启设备失败" }
            })
            if (res && res.code === 0) {
              return { result: 'success' }
            } else {
              return { result: 'error', message: res.error }
            }
          } else {
            return { result: 'error', message: '设备不在线' }
          }
        }
      }
    } else {
      return { result: 'error', message: res.error }
    }
  }

  function shutdown() {
    notyf.dismissAll()
    notyf.info('即将支持，敬请期待')
  }

  function start() {
    notyf.dismissAll()
    notyf.info('即将支持，敬请期待')
  }

  function getById(id: string) {
    return devices.value.find(d => d.id === id)
  }

  return {
    devices,
    shutdown,
    start,
    getById,
    $fetchList,
    $deploy,
    $updateConfig,
    $remove,
    $reboot,
    $showContainer
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
