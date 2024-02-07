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
import { useFetch } from "/@src/composable/useFetch"
import * as dic from "/@src/utils/enums-dic"
import { valert } from "../utils/dialog"

export const useDevices = defineStore('device', () => {
  const devices = ref<DeviceDetail[]>([])
  const $fetch  = useFetch()

  async function $fetchList() {
    const res = await $fetch('/api/device/list')
    if (res && res.devices) {
      devices.value = res.devices.map((d: any) => {
        if (!Array.isArray(d.podsStatus) || d.podsStatus.length === 0) {
          d.status = dic.DeviceStatus.Unavailable
        } else if (d.podsStatus.every((p: any) => p.status === dic.DeviceStatus.Running)) {
          d.status = dic.DeviceStatus.Running
        } else if (d.podsStatus.every((p: any) => [dic.DeviceStatus.Pending, dic.DeviceStatus.Running].includes(p.status))) {
          d.status = dic.DeviceStatus.Pending
        } else if (d.podsStatus.some((p: any) => p.status === dic.DeviceStatus.Terminating)) {
          d.status = dic.DeviceStatus.Terminating
        } else if (d.podsStatus.some((p: any) => p.status === dic.DeviceStatus.Failed)) {
          d.status = dic.DeviceStatus.Failed
        } else {
          d.status = dic.DeviceStatus.Unavailable
        }
        d.statusInfo = dic.DeviceStatusDic[d.status]
        return d
      })
    }
  }

  async function $deploy(name: string, tmpl: string, body: any) {
    const res = await $fetch('/api/device/create', {
      body: {
        name,
        tmpl,
        body: JSON.stringify(body)
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
    const res = await $fetch('/api/cluster/podPhase', {
      body: { device: device.name }
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
    const res = await $fetch('/api/cluster/podPhase', {
      body: { device: device.name }
    })
    if (res && res.code === 0) {
      if (typeof res.data === 'object') {
        const podKey = Object.keys(res.data).find(k => k.includes(device.name))
        if (podKey) {
          const pod = res.data[podKey]
          if (pod === 'Running') {
            const res = await $fetch('/api/cluster/pod.delete', {
              body: {
                pod: podKey
              }
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

  async function $stop(id: string) {
    return await $fetch('/api/device/stop', {
      body: { id }
    })
  }

  async function $start(id: string) {
    return await $fetch('/api/device/start', {
      body: { id }
    })
  }

  function getById(id: string) {
    return devices.value.find(d => d.id === id)
  }

  return {
    devices,
    $stop,
    $start,
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
