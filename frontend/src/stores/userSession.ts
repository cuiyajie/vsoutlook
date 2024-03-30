import { acceptHMRUpdate, defineStore } from 'pinia'
import { ref, computed } from 'vue'
import { useFetch } from "/@src/composable/useFetch"

export type UserData = {
  id: string,
  name: string
}

function parseJsonString<T>(json: string, def: T) {
  let result: T = def
  try {
    result = JSON.parse(json || "")
  } catch (error) {
    result = def
  }
  return result
}

export const useUserSession = defineStore('userSession', () => {
  const user = ref<Partial<UserData>>()
  const settings = ref<Partial<Settings>>({})
  const loading = ref(true)
  const $fetch  = useFetch()

  const isLoggedIn = computed(() => !!user.value?.id)

  function setUser(newUser: Partial<UserData>) {
    user.value = newUser
  }

  function setSettings(newSettings: Partial<Record<keyof Settings, string>>) {
    if (newSettings['auto_save_container_config']) {
      settings.value.auto_save_container_config = String(newSettings.auto_save_container_config) === 'true'
      delete newSettings.auto_save_container_config
    }
    if (newSettings['authorization_services']) {
      settings.value.authorization_services = parseJsonString(newSettings.authorization_services, [])
      delete newSettings.authorization_services
    }
    if (newSettings['mv_template_list']) {
      settings.value.mv_template_list = parseJsonString(newSettings.mv_template_list, [])
      delete newSettings.mv_template_list
    }
    settings.value = Object.assign({}, settings.value, newSettings)
  }

  function setLoading(newLoading: boolean) {
    loading.value = newLoading
  }

  async function $updateSettings(newSettings: any) {
    const res = await $fetch('/api/settings/update', {
      body: newSettings
    })
    if (res?.settings) {
      const s = res.settings
      setSettings({ [s.key]: s.value } as any)
    }
  }

  async function logoutUser() {
    user.value = undefined
  }

  return {
    user,
    isLoggedIn,
    loading,
    logoutUser,
    settings,
    setUser,
    setLoading,
    setSettings,
    $updateSettings
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
  import.meta.hot.accept(acceptHMRUpdate(useUserSession, import.meta.hot))
}
