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
import * as data from '/@src/data/device-templates'
import { useNotyf } from "/@src/composable/useNotyf"
import { useTemplateType } from './templateType';
import { useTemplate } from './template'
import * as dic from '/@src/utils/enums-dic'

export const useDevices = defineStore('device', () => {
  const devices = ref<DeviceVerbose[]>([])
  const tmplTypeStore = useTemplateType()
  const tmplStore = useTemplate()
  const notyf = useNotyf()

  watchEffect(() => {
    devices.value = data.devices.map(d => {
      const tmpl = tmplStore.getById(d.tmplId) as TemplateDataVerbose
      if (!tmpl) return null
      const tmplType = tmplTypeStore.getById(tmpl.type) as TemplateType
      if (!tmpl) return null
      return {
        ...d,
        tmpl: {
          ...tmpl,
          type: tmplType
        },
        statusInfo: dic.DeviceStatusDic[d.status],
        inputInfo: dic.InputSignalDic[d.input],
        outputInfo: dic.OutputSignalDic[d.output],
        ptpInfo: dic.PtpStatusDic[d.ptp]
      }
    }).filter(d => d)
  })

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
    getById
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
