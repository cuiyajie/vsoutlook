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

export const useTemplateType = defineStore('templateType', () => {
  const tmplTypes = ref<TemplateType[]>([])
  const $fetch = useFetch()

  function getById(id: string) {
    return tmplTypes.value.find(t => t.id === id)
  }

  async function $fetchList() {
    const res = await $fetch('/api/tmpl_type/list')
    if (res.types) {
      tmplTypes.value = res.types.map((v: any) => {
        v.icon = `/images/tmpl/${v.icon}`
        return v
      })
    }
  }

  return {
    tmplTypes,
    $fetchList,
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
  import.meta.hot.accept(acceptHMRUpdate(useTemplateType, import.meta.hot))
}
