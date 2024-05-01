<script setup lang="ts">
import { useUserSession } from '@src/stores/userSession'
import { useFetch } from "@src/composable/useFetch";

const userSession = useUserSession()
const router = useRouter()
const $fetch = useFetch()
const currUser = computed(() => userSession.user)

async function logout() {
  await $fetch('/api/logout')
  userSession.logoutUser()
  router.push('/login')
}
</script>

<template>
  <VDropdown
    right
    spaced
    class="user-dropdown profile-dropdown"
  >
    <template #button="{ toggle }">
      <a
        role="button"
        tabindex="0"
        class="is-trigger dropdown-trigger"
        aria-haspopup="true"
        @keydown.space.prevent="toggle"
        @click="toggle"
      >
        <VAvatar picture="/images/avatars/placeholder.jpg" />
      </a>
    </template>

    <template #content>
      <div class="dropdown-head">
        <VAvatar
          size="large"
          picture="/images/avatars/placeholder.jpg"
        />

        <div class="meta">
          <span>{{ currUser?.name }}</span>
        </div>
      </div>

      <hr class="dropdown-divider">

      <div class="dropdown-item is-button">
        <VButton
          class="logout-button"
          icon="feather:log-out"
          color="primary"
          role="menuitem"
          raised
          fullwidth
          @click="logout"
        >
          Logout
        </VButton>
      </div>
    </template>
  </VDropdown>
</template>

<style lang="scss">
.profile-dropdown {

  > a {
    width: 48px;
    height: 48px;
  }


  &.dropdown.is-spaced:hover .is-trigger,
  &.dropdown.is-spaced.is-active .is-trigger {
    border: 1px solid var(--primary);
    border-radius: 9999px;
  }

}

.main-sidebar .sidebar-inner .bottom-menu .profile-dropdown .dropdown-menu .dropdown-content {

  .dropdown-head {
    margin-bottom: 0;
  }

  .dropdown-divider {
    margin: 0;
  }
}

</style>
