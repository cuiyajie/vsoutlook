import { acceptHMRUpdate, defineStore } from 'pinia'
import { useFetch } from '@src/composable/useFetch';
import { useNotyf } from "@src/composable/useNotyf";
import dayjs from "dayjs";

export const useBackup = defineStore('backupApi', () => {
  const notyf = useNotyf()
  const $fetch = useFetch()

  async function $exportData(params: { name: string, table: string }[]) {
    const tables = params.map(p => p.table)
    $fetch('/api/backup/export', {
      body: { tables },
    }).then(res => {
      if (res?.data) {
        const blob = new Blob([JSON.stringify(res.data, null, 2)], { type: 'application/json' })
        const url = URL.createObjectURL(blob)
        const a = document.createElement('a')
        a.href = url
        a.download = `${params.length === 0 ? "[空数据]" : params.map(p => p.name).join('_')}_${dayjs().format('YYYYMMDDHHmmss')}.json`
        a.click()
        URL.revokeObjectURL(url)
      } else {
        notyf.error(res.message)
      }
    })
  }

  function $importData(jsonFile: File, skip: boolean) {
    return new Promise(resolve => {
      const reader = new FileReader()
      reader.onload = () => {
        $fetch('/api/backup/import', {
          body: {
            file: reader.result as string,
            skip,
          },
        }).then(res => {
          resolve(null)
          if (res?.data) {
            notyf.success('数据导入成功')
          } else {
            notyf.error(res.message)
          }
        })
      }
      reader.readAsText(jsonFile)
    })
  }

  return {
    $exportData,
    $importData,
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
  import.meta.hot.accept(acceptHMRUpdate(useBackup, import.meta.hot))
}
