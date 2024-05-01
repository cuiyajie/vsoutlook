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
import { useNotyf } from "@src/composable/useNotyf"
import { useFetch } from "@src/composable/useFetch"

export const useTemplate = defineStore('template', () => {
  const tmpls = ref<TemplateData[]>([])
  const router = useRouter()
  const notyf = useNotyf()
  const $fetch  = useFetch()

  async function $fetchList() {
    const res = await $fetch('/api/tmpl/list')
    if (res && res.tmpls) {
      tmpls.value = res.tmpls
    }
  }

  async function $addTmpl(tmpl: Pick<TemplateData, 'name' | 'type'>) {
    const res = await $fetch('/api/tmpl/create', { body: tmpl })
    if (res && res.tmpl) {
      tmpls.value.push(res.tmpl)
      return res.tmpl
    }
  }

  async function $deleteTmpl(id: string) {
    const res = await $fetch('/api/tmpl/delete', { body: { id } })
    if (res && res.result === 'ok') {
      tmpls.value = tmpls.value.filter(tmpl => tmpl.id !== id)
      return 'success'
    }
  }

  async function $updateTmpl(id: string, info: Partial<Pick<TemplateData, 'description' | 'flow' | 'requirement' | 'name'>>) {
    const res = await $fetch('/api/tmpl/update', { body: { id, ...info } })
    if (res && res.tmpl) {
      return res.tmpl
    }
  }

  async function $getTmplById(id: string) {
    const res = await $fetch('/api/tmpl/info', { body: { id } })
    if (res && res.tmpl) {
      tmpls.value = tmpls.value.map(tmpl => tmpl.id === id ? res.tmpl : tmpl)
      return res.tmpl
    }
    return null
  }

  async function $listed(tmpl: TemplateData) {
    const res = await $fetch('/api/tmpl/listed', { body: { id: tmpl.id, listed: !tmpl.listed } })
    if (res && res.result === 'ok') {
      tmpl.listed = !tmpl.listed
      return 'success'
    }
  }

  function getById(id: string) {
    return tmpls.value.find(t => t.id === id)
  }

  function navigate(id: string) {
    return router.push(`/app/template/${id}`)
  }

  function exportTmpl() {
    notyf.dismissAll()
    notyf.info('即将支持，敬请期待')
  }

  return {
    tmpls,
    $fetchList,
    $addTmpl,
    $deleteTmpl,
    $updateTmpl,
    $getTmplById,
    $listed,
    getById,
    navigate,
    exportTmpl
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
  import.meta.hot.accept(acceptHMRUpdate(useTemplate, import.meta.hot))
}
