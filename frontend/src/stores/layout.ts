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
import { useFetch } from "@src/composable/useFetch"
import { useUserSession } from './userSession';

export const useLayout = defineStore('layout', () => {
  const layouts = ref<Layout[]>([])
  const $fetch  = useFetch()
  const userSession = useUserSession()

  async function $fetchList() {
    const res = await $fetch('/api/layout/list')
    if (res && res.layouts) {
      layouts.value = res.layouts
    }
  }

  async function $addLayout(layout: Pick<Layout, 'name' | 'size'>) {
    const res = await $fetch('/api/layout/create', { body: layout })
    if (res) {
      if (res.layout) {
        layouts.value.push(res.layout)
      }
      if (res.settings) {
        userSession.updateMtvSettings(res.settings)
      }
    }
    return res
  }

  async function $deleteLayout(id: string) {
    const res = await $fetch('/api/layout/delete', { body: { id } })
    if (res) {
      if (res.settings) {
        userSession.updateMtvSettings(res.settings)
      }
      if (res.result === 'ok') {
        layouts.value = layouts.value.filter(layout => layout.id !== id)
        return 'success'
      }
    }
    return res?.message
  }

  async function $updateLayout(layout: Pick<Layout, 'id' | 'name' | 'size'>) {
    const res = await $fetch('/api/layout/update', { body: layout })
    if (res) {
      if (res.layout) {
        const index = layouts.value.findIndex(l => l.id === layout.id)
        layouts.value[index] = res.layout
      }
      if (res.settings) {
        userSession.updateMtvSettings(res.settings)
      }
    }
    return res
  }

  async function $duplicate(id: string, name: string) {
    const res = await $fetch('/api/layout/duplicate', { body: { id, name } })
    if (res) {
      if (res.layout) {
        layouts.value.push(res.layout)
      }
      if (res.settings) {
        userSession.updateMtvSettings(res.settings)
      }
    }
    return res
  }

  async function $updateContent(id: string, content: any[]) {
    const res = await $fetch('/api/layout/update.content', { body: { id, content } })
    if (res && res.layout) {
      return res.layout
    }
  }

  async function $publish(layout: Layout) {
    const res = await $fetch('/api/layout/publish', { body: { id: layout.id } })
    if (res) {
      if (res.settings) {
        userSession.updateMtvSettings(res.settings)
      }
      if (res.result === 'ok') {
        layout.published = true
        return 'success'
      }
    }
    return res?.message
  }

  return {
    layouts,
    $fetchList,
    $addLayout,
    $deleteLayout,
    $updateLayout,
    $duplicate,
    $updateContent,
    $publish
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
  import.meta.hot.accept(acceptHMRUpdate(useLayout, import.meta.hot))
}
