import { acceptHMRUpdate, defineStore } from 'pinia'
import { ref, computed } from 'vue'

export type UserData = {
  id: string,
  name: string
}

export const useUserSession = defineStore('userSession', () => {
  const user = ref<Partial<UserData>>()
  const loading = ref(true)

  const isLoggedIn = computed(() => !!user.value?.id)

  function setUser(newUser: Partial<UserData>) {
    user.value = newUser
  }

  function setLoading(newLoading: boolean) {
    loading.value = newLoading
  }

  async function logoutUser() {
    user.value = undefined
  }

  return {
    user,
    isLoggedIn,
    loading,
    logoutUser,
    setUser,
    setLoading,
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
