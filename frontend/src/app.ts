import { createApp as createClientApp } from 'vue'

import { createHead } from '@unhead/vue'
import { InferSeoMetaPlugin } from '@unhead/addons'
import { createPinia } from 'pinia'
import { createRouter } from './router'
import VueroApp from './VueroApp.vue'
import { ColorPicker } from 'vue3-colorpicker'
import VueTippy from 'vue-tippy'
import ExpandTransition from './components/partials/widgets/expand-transition/ExpandTransiton.vue'
import './styles'
/* import the required styles */
import '@vue-flow/core/dist/style.css'
/* import the default theme (optional) */
import '@vue-flow/core/dist/theme-default.css'
import 'vue3-colorpicker/style.css'
import 'tippy.js/dist/tippy.css'
import 'tippy.js/themes/light.css'
import 'tippy.js/animations/scale.css'

export type VueroAppContext = Awaited<ReturnType<typeof createApp>>
export type VueroPlugin = (vuero: VueroAppContext) => void | Promise<void>

const plugins = import.meta.glob<{ default: VueroPlugin }>('./plugins/*.ts')

// this is a helper function to define plugins with autocompletion
export function definePlugin(plugin: VueroPlugin) {
  return plugin
}

export async function createApp() {
  const app = createClientApp(VueroApp)
  const router = createRouter()

  const head = createHead({
    plugins: [InferSeoMetaPlugin()],
  })
  app.use(head)

  const pinia = createPinia()
  app.use(pinia)
  app.use(VueTippy, {
    directive: 'tippy', // => v-tippy
    component: 'tippy', // => <tippy/>
  })

  const vuero = {
    app,
    router,
    head,
    pinia,
  }

  app.provide('vuero', vuero)
  app.component('ColorPicker', ColorPicker)
  app.component('ExpandTransition', ExpandTransition)

  for (const path in plugins) {
    try {
      const { default: plugin } = await plugins[path]()
      await plugin(vuero)
    } catch (error) {
      console.error(`Error while loading plugin "${path}".`)
      console.error(error)
    }
  }

  // use router after plugin registration, so we can register navigation guards
  app.use(vuero.router)

  return vuero
}
