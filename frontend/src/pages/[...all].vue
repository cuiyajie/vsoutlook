<script setup lang="ts">
/**
 * This is a Vue Component that will be
 * automatically mapped to a catch allpath on vue-router (404).
 *
 * You will be able to access this page  at http://localhost:3000/non-existing-page
 *
 * Read more about routing:
 * @see /vite.config.ts
 * @see /src/router.ts
 */

import { useI18n } from 'vue-i18n'

const { t } = useI18n()

useHead({
  title: `${t('pages.not-found.page-title')} - ${__SITE_NAME__}`,
  meta: [
    {
      name: 'robots',
      content: 'noindex',
    },
  ],
})
</script>

<template>
  <LandingLayout theme="darker">
    <div class="error-container">
      <div class="error-wrapper">
        <div class="error-inner has-text-centered">
          <div class="bg-number dark-inverted">
            404
          </div>
          <img
            class="light-image"
            src="/@src/assets/placeholders/error.svg"
            alt=""
          >
          <img
            class="dark-image"
            src="/@src/assets/placeholders/error-dark.svg"
            alt=""
          >
          <h3 class="dark-inverted mt-6">
            我们找不到那个页面
          </h3>
          <p>
            看起来我们找不到那个页面。如果问题仍然存在，请再试一次或联系管理员。
          </p>
          <div class="button-wrap">
            <VButton
              color="primary"
              elevated
              @click="$router.go(-1)"
            >
              返回
            </VButton>
          </div>
        </div>
      </div>
    </div>
  </LandingLayout>
</template>

<style lang="scss">
.error-container {
  width: 100vw;
  min-height: 100vh;

  .error-nav {
    .dark-mode {
      position: absolute;
      inset-inline-end: 0;
      top: 0;
      display: inline-block;
      transform: scale(0.5);
    }
  }

  .error-wrapper {
    max-width: 840px;
    margin: 0 auto;
    padding-top: 40px;

    .error-inner {
      position: relative;
      max-width: 540px;
      margin: 0 auto;

      .bg-number {
        font-family: var(--font);
        position: absolute;
        top: -58px;
        inset-inline-start: -50px;
        inset-inline-end: 0;
        margin: 0 auto;
        font-size: 28rem;
        font-weight: 600;
        opacity: 0.15;
        z-index: 0;
      }

      img,
      svg,
      h3,
      p,
      .button-wrap {
        position: relative;
        z-index: 1;
      }

      img,
      svg {
        display: block;
        max-width: 100%;
        margin: 0 auto;
      }

      h3 {
        font-size: 1.5rem;
        font-family: var(--font-alt);
        color: var(--dark-text);
        font-weight: 600;
        margin-top: 10px;
      }

      p {
        font-family: var(--font);
        font-size: 1.1rem;
        margin-bottom: 16px;
      }

      .button-wrap {
        .button {
          min-width: 220px;
          min-height: 50px;
        }
      }
    }
  }
}

.is-dark {
  .error-container {
    .error-wrapper {
      .error-inner {
        .bg-number {
          opacity: 0.09;
        }
      }
    }
  }
}

@media only screen and (width <= 767px) {
  .error-container {
    .error-wrapper {
      padding-top: 60px;

      .error-inner {
        padding: 10px;

        .bg-number {
          top: -35px;
          inset-inline-start: -18px;
          inset-inline-end: 0;
          font-size: 13rem;
        }

        img,
        svg {
          max-width: 345px;
        }
      }
    }
  }
}
</style>
