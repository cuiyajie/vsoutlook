<script setup lang="ts">
</script>

<template>
  <div class="page-content-inner card-hub">
    <div class="columns">
      <div class="column is-12">
        <div class="hub-wrapper">
          <!-- <div class="hub-header has-text-centered">
            <AnimatedLogo width="64px" />

            <h3 class="title is-3 mt-4 mb-2">
              欢迎进入维视远景系统
            </h3>
          </div> -->

          <div class="hub-body">
            <div class="body-inner">
              <!--Component/content-->
              <HubCard
                title="应用商店"
                link="/app/template"
                color="primary"
              >
                <template #icon>
                  <img
                    class="light-image"
                    src="/@src/assets/placeholders/modal.svg"
                    alt=""
                  >
                  <img
                    class="dark-image"
                    src="/@src/assets/placeholders/modal-dark.svg"
                    alt=""
                  >
                </template>
              </HubCard>
              <!--Component/content-->
              <HubCard
                title="设备管理"
                link="/app/devices"
                color="info"
              >
                <template #icon>
                  <img
                    class="light-image"
                    src="/@src/assets/placeholders/switches.svg"
                    alt=""
                  >
                  <img
                    class="dark-image"
                    src="/@src/assets/placeholders/switches-dark.svg"
                    alt=""
                  >
                </template>
              </HubCard>

              <!--Component/tags-->
              <HubCard
                title="资源管理"
                link="/app/resource"
                color="warning"
              >
                <template #icon>
                  <img
                    class="light-image"
                    src="/@src/assets/placeholders/table.svg"
                    alt=""
                  >
                  <img
                    class="dark-image"
                    src="/@src/assets/placeholders/table-dark.svg"
                    alt=""
                  >
                </template>
              </HubCard>

              <!--Component/tabs-->
              <HubCard
                title="系统配置"
                link="/app/config"
                color="danger"
              >
                <template #icon>
                  <img
                    class="light-image"
                    src="/@src/assets/placeholders/message.svg"
                    alt=""
                  >
                  <img
                    class="dark-image"
                    src="/@src/assets/placeholders/message-dark.svg"
                    alt=""
                  >
                </template>
              </HubCard>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style lang="scss">
@import "/@src/scss/abstracts/all";
.card-hub {
  max-width: 1440px;
  margin: 0 auto;

  .hub-header {
    text-align: center;

    > img {
      display: block;
      margin: 0 auto;
      max-width: 300px;
    }

    .v-avatar {
      margin: 0 auto 12px;
    }

    .anim-icon {
      margin-bottom: 12px;
    }

    p {
      font-size: 1.15rem;
    }
  }

  .hub-body {
    padding: 10px 0 20px;

    .body-inner {
      display: flex;
      flex-wrap: wrap;

      .hub-card {
        @include vuero-card;
        background-color: var(--white);
        border-radius: var(--radius-large);
        border: 1px solid var(--fade-grey-dark-3);
        width: calc(50% - 20px);
        max-width: calc(50% - 20px);
        margin: 10px;
        transition: all 0.3s; // transition-all test

        &:hover {
          box-shadow: var(--light-box-shadow);

          .hub-card-body {
            img {
              filter: grayscale(0);
              opacity: 1;
            }
          }
        }

        .hub-card-body {
          display: flex;
          flex-direction: row;
          justify-content: space-between;
          padding: 24px;

          .inner-content {
            flex-grow: 1;
            padding-left: 12px;
            padding-right: 12px;
            display: flex;
            align-items: center;
            justify-content: center;

            .title {
              font-size: 2rem;
              margin-bottom: 0;
            }

            p {
              max-width: 220px;
            }
          }

          img {
            display: block;
            width: 100%;
            max-width: 180px;
            height: 120px;
            filter: grayscale(1);
            opacity: 0.6;
            transition: all 0.3s; // transition-all test
          }
        }

        .hub-card-foot {
          border-top: 1px dashed var(--fade-grey-dark-4);
          border-top-width: 1.6px;
          padding: 18px 24px;

          a {
            color: var(--primary);
            opacity: 0.6;
            font-weight: 500;
            transition: opacity 0.3s;

            &:hover,
            &:focus {
              opacity: 1;
            }
          }
        }
      }
    }
  }
}

.is-dark {
  .card-hub {
    .hub-header {
      .v-avatar {
        .badge {
          border-color: var(--dark-sidebar-light-6) !important;
        }
      }
    }

    .hub-body {
      .body-inner {
        .hub-card {
          @include vuero-card;
          background: var(--dark-sidebar-light-6);
          border-color: var(--dark-sidebar-light-12);

          .hub-card-foot {
            border-color: var(--muted-grey-dark-25);

            a {
              color: var(--primary);
            }
          }
        }
      }
    }
  }
}

@media only screen and (width <= 767px) {
  .card-hub {
    .hub-body {
      .body-inner {
        .hub-card {
          width: calc(100% - 20px);
          max-width: calc(100% - 20px);
        }
      }
    }
  }
}
</style>
