import { ref } from 'vue'
import { acceptHMRUpdate, defineStore } from 'pinia'
import { useFetch } from '@src/composable/useFetch';
import { useNotyf } from "@src/composable/useNotyf";
import { useUserSession } from './userSession';

const QueryDelay = 5 * 1000

function def_endswt_info() {
  return {
    target: 0,
    source: -1,
    status: 0,
  }
}

export const useEndSwtApi = defineStore('endSwtApi', () => {
  const notyf = useNotyf()
  const $fetch = useFetch()
  let timer: NodeJS.Timeout | null = null
  const endSwtInfo = ref(def_endswt_info())
  const usStore = useUserSession()

  function $startQuery() {
    if (timer) {
      clearTimeout(timer)
      timer = null
    }
    $fetchList().finally(() => {
      timer = null
      timer = setTimeout(() => {
        $startQuery()
      }, QueryDelay)
    })
  }

  function $stopQuery() {
    if (timer) {
      clearTimeout(timer)
      timer = null
    }
  }

  async function $fetchList() {
    if (!usStore.settings.endswt_api) return
    endSwtInfo.value = await $fetch('/api/endswt/target').then(res => {
      if (res.error) {
        return {
          target: 0,
          source: -1,
          status: 1,
        }
      } else {
        return res.data
      }
    })
  }

  async function $switchTarget(v: number) {
    if (!usStore.settings.endswt_api) return
    const res = await $fetch('/api/endswt/switch', {
      body: {
        target: 0,
        source: v,
      }
    })
    if (res.error) {
      notyf.error('Failed to switch target: ' + res.message)
    }
    if (timer) {
      clearTimeout(timer)
      timer = null
    }
    $startQuery()
  }

  return {
    endSwtInfo,
    $fetchList,
    $startQuery,
    $stopQuery,
    $switchTarget
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
  import.meta.hot.accept(acceptHMRUpdate(useEndSwtApi, import.meta.hot))
}
