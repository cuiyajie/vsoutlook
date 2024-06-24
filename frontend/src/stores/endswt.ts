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
    try {
      endSwtInfo.value = await $fetch('/api/endswt/target', { timeout: 15 * 1000 }).then(res => {
        if (res?.data) {
          return res.data
        } else {
          return def_endswt_info()
        }
      })
    } catch {
      endSwtInfo.value = def_endswt_info()
    }
  }

  async function $switchTarget(v: number) {
    if (!usStore.settings.endswt_api) return
    try {
      endSwtInfo.value = {
        target: 0,
        source: v,
        status: 0,
      }
      const res = await $fetch('/api/endswt/switch', {
        body: {
          target: 0,
          source: v,
        },
        timeout: 15 * 1000
      })
      if (res.error) {
        notyf.error('Failed to switch target: ' + res.message)
        endSwtInfo.value = def_endswt_info()
      } else {
        if (timer) {
          clearTimeout(timer)
          timer = null
        }
        $startQuery()
      }
    } catch {
      notyf.error('请求失败，请检查对外 API 地址设置')
      endSwtInfo.value = def_endswt_info()
    }
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
