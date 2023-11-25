<template>
  <div class="onboarding-wrapper onboarding-v4">
    <div class="title-wrap">
      <Logo
        width="40"
        height="40"
      />
      <h2 class="dark-inverted">
        SENSE VISION TECHNOLOGIES
      </h2>
    </div>

    <div class="onboarding-wrap">
      <div class="onboarding-wrap-inner">
        <!--Card-->
        <RouterLink
          class="onboarding-card"
          to="/app/template"
        >
          <div class="image-container">
            <img
              class="dark-image"
              src="/@src/assets/placeholders/template.svg"
              alt=""
            >
          </div>
          <div class="onboarding-content">
            <h3>
              应用商店
            </h3>
          </div>
        </RouterLink>
        <!--Card-->
        <RouterLink
          class="onboarding-card"
          to="/app/devices"
        >
          <div class="image-container">
            <img
              class="dark-image"
              src="/@src/assets/placeholders/device.svg"
              alt=""
            >
          </div>
          <div class="onboarding-content">
            <h3>
              设备管理
            </h3>
          </div>
        </RouterLink>
        <!--Card-->
        <RouterLink
          class="onboarding-card"
          to="/app/resource"
        >
          <div class="image-container">
            <img
              class="dark-image"
              src="/@src/assets/placeholders/resource.svg"
              alt=""
            >
          </div>
          <div class="onboarding-content">
            <h3>
              资源管理
            </h3>
          </div>
        </RouterLink>
        <!--Card-->
        <RouterLink
          class="onboarding-card"
          to="/app/config"
        >
          <div class="image-container">
            <img
              class="dark-image"
              src="/@src/assets/placeholders/system.svg"
              alt=""
            >
          </div>
          <div class="onboarding-content">
            <h3>
              系统配置
            </h3>
          </div>
        </RouterLink>
      </div>
    </div>
  </div>
</template>

<style lang="scss">
@import '/@src/scss/abstracts/all';

.is-navbar {
  .onboarding-wrapper {
    margin-top: 40px;
  }
}

.onboarding-wrapper {
  .title-wrap {
    text-align: center;

    h2 {
      font-family: var(--font-alt);
      font-weight: 600;
      font-size: 1.4rem;
      color: var(--dark-text);
      font-style: italic;
      margin-top: 20px;
    }
  }

  &.onboarding-v4 {
    .onboarding-wrap {
      padding: 30px 0;

      .onboarding-wrap-inner {
        display: flex;
        max-width: 1200px;
        margin: 0 auto;

        .onboarding-card {
          @include vuero-r-card;

          margin: 8px;
          width: calc(25% - 16px);
          text-align: center;
          padding: 0;
          overflow: hidden;

          .image-container {
            display: flex;
            align-items: center;
            justify-content: center;
            height: 260px;

            > img {
              display: block;
              width: 80%;
              transition: transform .3s ease;
            }
          }

          .onboarding-content {
            padding: 30px 20px;
            background: var(--dark-sidebar-light-4);
            border-color: var(--dark-sidebar-light-16);

            h3 {
              font-family: var(--font-alt);
              font-size: 1.6rem;
              font-weight: 600;
              color: var(--dark-dark-text);
              letter-spacing: .3em;
              transition: color .3s ease;
            }
          }

          &:hover {
            .image-container > img {
              transform: scale(1.1);
            }

            .onboarding-content h3 {
              color: var(--primary);
            }
          }
        }
      }
    }
  }
}

.is-dark {
  .onboarding-wrapper {
    &.onboarding-v4 {
      .onboarding-wrap {
        .onboarding-wrap-inner {
          .onboarding-card {
            @include vuero-card--dark;
          }
        }
      }
    }
  }
}

@media only screen and (width <= 767px) {
  .onboarding-wrapper {
    .standard-onboarding {
      padding: 20px 0;
    }

    .title-wrap {
      h2 {
        max-width: 280px;
        margin: 0 auto;
      }
    }

    .onboarding-wrap {
      .onboarding-wrap-inner {
        flex-wrap: wrap;

        .onboarding-card {
          min-width: calc(100% - 16px) !important;
        }
      }
    }
  }
}
</style>
