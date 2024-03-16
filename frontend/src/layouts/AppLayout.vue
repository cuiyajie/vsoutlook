<script setup lang="ts">
import type { SidebarTheme } from '/@src/components/navigation/desktop/Sidebar.vue'
import { useViewWrapper } from '/@src/stores/viewWrapper'

const props = withDefaults(
  defineProps<{
    theme?: SidebarTheme
    defaultSidebar?: string
    closeOnChange?: boolean
    openOnMounted?: boolean
    nowrap?: boolean
  }>(),
  {
    defaultSidebar: 'dashboard',
    theme: 'default',
  }
)

const viewWrapper = useViewWrapper()
const route = useRoute()
const isMobileSidebarOpen = ref(false)
const isDesktopSidebarOpen = ref(props.openOnMounted)
const activeMobileSubsidebar = ref(props.defaultSidebar)

function switchSidebar(id: string) {
  if (id === activeMobileSubsidebar.value) {
    isDesktopSidebarOpen.value = !isDesktopSidebarOpen.value
  } else {
    isDesktopSidebarOpen.value = true
    activeMobileSubsidebar.value = id
  }
}

/**
 * watchPostEffect callback will be executed each time dependent reactive values has changed
 */
watchPostEffect(() => {
  viewWrapper.setPushed(isDesktopSidebarOpen.value ?? false)
})
watch(
  () => route.fullPath,
  () => {
    isMobileSidebarOpen.value = false

    if (props.closeOnChange && isDesktopSidebarOpen.value) {
      isDesktopSidebarOpen.value = false
    }
  }
)
</script>

<template>
  <div class="sidebar-layout is-narrow">
    <div class="app-overlay" />

    <!-- Mobile navigation -->
    <MobileNavbar
      :is-open="isMobileSidebarOpen"
      @toggle="isMobileSidebarOpen = !isMobileSidebarOpen"
    >
      <template #brand>
        <RouterLink
          to="/"
          class="navbar-item is-brand"
        >
          <Logo
            width="38"
            height="38"
          />
        </RouterLink>

        <div class="brand-end" />
      </template>
    </MobileNavbar>

    <!-- Mobile sidebar links -->
    <MobileSidebar
      :is-open="isMobileSidebarOpen"
      @toggle="isMobileSidebarOpen = !isMobileSidebarOpen"
    >
      <template #links>
        <li>
          <RouterLink to="/app">
            <i
              aria-hidden="true"
              class="iconify"
              data-icon="feather:home"
            />
          </RouterLink>
        </li>
      </template>
    </MobileSidebar>

    <!-- Mobile subsidebar links -->
    <Transition name="slide-x">
      <DashboardsMobileSubsidebar
        v-if="isMobileSidebarOpen && activeMobileSubsidebar === 'dashboard'"
      />
    </Transition>

    <Sidebar
      :theme="props.theme"
      :is-open="isDesktopSidebarOpen"
    >
      <template #links>
        <!-- Dashboards -->
        <!-- <li>
          <a
            :class="[activeMobileSubsidebar === 'dashboard' && 'is-active']"
            data-content="Dashboards"
            tabindex="0"
            role="button"
            @keydown.space.prevent="switchSidebar('dashboard')"
            @click="switchSidebar('dashboard')"
          >
            <i
              aria-hidden="true"
              class="iconify sidebar-svg"
              data-icon="feather:home"
            />
            <span>主页</span>
          </a>
        </li> -->
      </template>

      <template #bottom-links>
        <!-- Profile Dropdown -->
        <li>
          <UserProfileDropdown />
        </li>
      </template>
    </Sidebar>

    <Transition name="slide-x">
      <KeepAlive>
        <DashboardsSubsidebar
          v-if="isDesktopSidebarOpen && activeMobileSubsidebar === 'dashboard'"
          @close="isDesktopSidebarOpen = false"
        />
      </KeepAlive>
    </Transition>

    <VViewWrapper>
      <VPageContentWrapper>
        <template v-if="props.nowrap">
          <slot />
        </template>
        <VPageContent
          v-else
          class="is-relative"
        >
          <div class="page-title has-text-centered">
            <!-- Sidebar Trigger -->
            <div
              class="vuero-hamburger nav-trigger push-resize"
              tabindex="0"
              role="button"
              @keydown.space.prevent="isDesktopSidebarOpen = !isDesktopSidebarOpen"
              @click="isDesktopSidebarOpen = !isDesktopSidebarOpen"
            >
              <span class="menu-toggle has-chevron">
                <span
                  :class="[isDesktopSidebarOpen && 'active']"
                  class="icon-box-toggle"
                >
                  <span class="rotate">
                    <i
                      aria-hidden="true"
                      class="icon-line-top"
                    />
                    <i
                      aria-hidden="true"
                      class="icon-line-center"
                    />
                    <i
                      aria-hidden="true"
                      class="icon-line-bottom"
                    />
                  </span>
                </span>
              </span>
            </div>

            <div class="title-wrap">
              <h1 class="title is-4">
                {{ viewWrapper.pageTitle }}
              </h1>
            </div>

            <!-- <Toolbar class="desktop-toolbar" /> -->
          </div>

          <slot />
        </VPageContent>
      </VPageContentWrapper>
    </VViewWrapper>
  </div>
</template>
<style lang="scss">
.sidebar-layout.is-narrow {
  overflow: auto;

  .sidebar-panel {
    width: 160px;

    .inner {
      width: 160px;
    }
  }

  .view-wrapper.is-pushed-full {
    margin-inline-start: 160px;
    width: calc(100% - 160px);
    min-width: 1200px;
    overflow: auto;
  }

  .page-content-wrapper {
    max-width: 1920px;
  }
}
</style>
