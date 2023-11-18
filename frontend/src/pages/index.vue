<script setup lang="ts">
import { useUserSession } from "/@src/stores/userSession";

useHead({
  title: __SITE_NAME__,
});

const userSession = useUserSession();
const router = useRouter();

(() => {
  if (userSession.isLoggedIn) {
    router.replace("/app");
  } else {
    router.push("/login");
  }
})();
</script>
<template>
  <div />
</template>
