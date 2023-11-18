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
import * as data from '/@src/data/system-templates'
import { useDevices } from './device'

export const useSystems = defineStore('system', () => {
  const systems = ref<SystemVerbose[]>([])
  const deviceStore = useDevices()

  watchEffect(() => {
    systems.value = data.systemList.map(sys => {
      return {
        ...sys,
        deviceInfos: sys.devices.map(id => deviceStore.getById(id))
      } as SystemVerbose
    })
  })

  return {
    systems
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
  import.meta.hot.accept(acceptHMRUpdate(useSystems, import.meta.hot))
}
