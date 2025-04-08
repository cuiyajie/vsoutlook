import { acceptHMRUpdate, defineStore } from 'pinia'
import { ref, computed } from 'vue'
import { useFetch } from '@src/composable/useFetch'
import { defaultMvFont } from '@src/utils/constants/config'

export type UserData = {
  id: string
  name: string
}

function parseJsonString<T>(json: string, def: T) {
  let result: T = def
  try {
    result = JSON.parse(json || '')
  } catch (error) {
    result = def
  }
  return result
}

export const useUserSession = defineStore('userSession', () => {
  const user = ref<Partial<UserData>>()
  const settings = ref<Partial<Settings>>({})
  const loading = ref(true)
  const $fetch = useFetch()

  const isLoggedIn = computed(() => !!user.value?.id)

  function setUser(newUser: Partial<UserData>) {
    user.value = newUser
  }

  function setSettings(newSettings: Partial<Record<keyof Settings, string>>) {
    if (newSettings['auto_save_container_config']) {
      settings.value.auto_save_container_config =
        String(newSettings.auto_save_container_config) === 'true'
      delete newSettings.auto_save_container_config
    }
    if (newSettings['mv_template_fonts']) {
      settings.value.mv_template_fonts = parseJsonString<any>(
        newSettings.mv_template_fonts,
        []
      )
      if (settings.value.mv_template_fonts) {
        const idx = settings.value.mv_template_fonts.findIndex(
          (font) => font === defaultMvFont
        )
        if (idx !== -1) {
          settings.value.mv_template_fonts.splice(idx, 1)
        }
        settings.value.mv_template_fonts.unshift(defaultMvFont)
      }
      delete newSettings.mv_template_fonts
    }
    ;(
      [
        'authorization_services',
        'mv_template_list',
        'endswt_titles',
        'endswt_panel_types',
        'lut_upscale_names',
        'lut_downscale_names',
        'video_formats',
        'audio_formats',
        'audio_mappings',
        'tech_reviews',
        'audit_alarm_rules',
      ] as Array<keyof Settings>
    ).forEach((key) => {
      if (newSettings[key]) {
        settings.value[key] = parseJsonString<any>(newSettings[key]!, [])
        delete newSettings[key]
      }
    })
    settings.value = Object.assign({}, settings.value, newSettings)
  }

  function updateMtvSettings(mtvList: string) {
    settings.value = {
      ...settings.value,
      mv_template_list: parseJsonString(mtvList, settings.value.mv_template_list || []),
    }
  }

  function setLoading(newLoading: boolean) {
    loading.value = newLoading
  }

  async function $updateSettings(newSettings: any) {
    const res = await $fetch('/api/settings/update', {
      body: newSettings,
    })
    if (res?.settings) {
      const s = res.settings
      setSettings({ [s.key]: s.value } as any)
    }
    return res?.settings
  }

  async function logoutUser() {
    user.value = undefined
  }

  const endSwitchTitles = computed(() =>
    settings.value.endswt_titles?.length === 5
      ? settings.value.endswt_titles
      : Array.from({ length: 5 }).map(() => '')
  )

  return {
    user,
    isLoggedIn,
    loading,
    logoutUser,
    settings,
    setUser,
    setLoading,
    setSettings,
    $updateSettings,
    endSwitchTitles,
    updateMtvSettings,
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
