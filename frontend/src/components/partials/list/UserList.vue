<script setup lang="ts">
import type { VTagColor } from "@src/components/base/tags/VTag.vue";
import type { VAvatarProps } from "@src/components/base/avatar/VAvatar.vue";
import { useUserData } from "@src/stores/user";
import { UserRole } from "@src/utils/enums";

export interface UserData extends VAvatarProps {
  name: string;
  role: string;
}

const userStore = useUserData();

const filters = ref("");

const filteredData = computed(() => {
  if (!filters.value) {
    return userStore.users;
  } else {
    return userStore.users.filter((item) => {
      return item.name.match(new RegExp(filters.value, "i"));
    });
  }
});

const mapRoleData: Record<number, { color: VTagColor; label: string }> = {
  [UserRole.Admin]: {
    color: "primary" as VTagColor,
    label: "管理员",
  },
  [UserRole.Normal]: {
    color: "secondary" as VTagColor,
    label: "普通用户",
  },
};

function createUser() {
  bus.trigger(Signal.OpenNewUser);
}

userStore.$fetchList();
</script>

<template>
  <div>
    <div class="list-view-toolbar">
      <VField>
        <VControl icon="feather:search">
          <input
            v-model="filters"
            class="input custom-text-filter"
            placeholder="搜索..."
          />
        </VControl>
      </VField>

      <div class="buttons">
        <VButton
          color="primary"
          icon="fas fa-plus"
          elevated
          @keyup.space.prevent="createUser"
          @click.prevent="createUser"
        >
          添加用户
        </VButton>
      </div>
    </div>

    <div class="page-content-inner">
      <!--List-->
      <div class="list-view list-view-v1">
        <!--List Empty Search Placeholder -->
        <VPlaceholderPage
          :class="[filteredData.length !== 0 && 'is-hidden']"
          title="没有找到符合搜索条件的用户"
          subtitle="很遗憾，看起来我们无法找到与您输入的搜索词或条件匹配的用户。请尝试不同的搜索词或条件。"
          larger
        />

        <div class="list-view-inner">
          <!--Item-->
          <TransitionGroup name="list-complete" tag="div">
            <div v-for="(item, key) in filteredData" :key="key" class="list-view-item">
              <div class="list-view-item-inner">
                <div class="meta-left">
                  <h3>{{ item.name }}</h3>
                </div>
                <div class="meta-right">
                  <div class="tags">
                    <VTag
                      :label="mapRoleData[item.role].label"
                      :color="mapRoleData[item.role].color"
                      rounded
                      elevated
                    />
                  </div>
                </div>
              </div>
            </div>
          </TransitionGroup>
        </div>
      </div>

      <VFlexPagination
        v-if="filteredData.length > 5"
        :item-per-page="10"
        :total-items="873"
        :current-page="42"
        :max-links-displayed="5"
      />
    </div>
  </div>
  <NewUserModal />
</template>

<style lang="scss">
@import "@src/scss/abstracts/all";

.list-view-v1 {
  .list-view-item {
    @include vuero-r-card;

    margin-bottom: 16px;
    padding: 16px;

    .list-view-item-inner {
      display: flex;
      align-items: center;

      .meta-left {
        margin-inline-start: 16px;

        h3 {
          font-family: var(--font-alt);
          color: var(--dark-text);
          font-weight: 600;
          font-size: 1rem;
          line-height: 1;
        }

        > span:not(.tag) {
          font-size: 0.9rem;
          color: var(--light-text);

          svg {
            height: 12px;
            width: 12px;
          }
        }
      }

      .meta-right {
        margin-inline-start: auto;
        display: flex;
        justify-content: flex-end;
        align-items: center;

        .tags {
          margin-inline-end: 30px;
          margin-bottom: 0;

          .tag {
            margin-bottom: 0;
          }
        }

        .stats {
          display: flex;
          align-items: center;
          margin-inline-end: 30px;

          .stat {
            display: flex;
            align-items: center;
            flex-direction: column;
            text-align: center;
            color: var(--light-text);

            > span {
              font-family: var(--font);

              &:first-child {
                font-size: 1.2rem;
                font-weight: 600;
                color: var(--dark-text);
                line-height: 1.4;
              }

              &:nth-child(2) {
                text-transform: uppercase;
                font-family: var(--font-alt);
                font-size: 0.75rem;
              }
            }

            svg {
              height: 16px;
              width: 16px;
            }

            i {
              font-size: 1.4rem;
            }
          }

          .separator {
            height: 25px;
            width: 2px;
            border-inline-end: 1px solid var(--fade-grey-dark-3);
            margin: 0 16px;
          }
        }

        .network {
          display: flex;
          justify-content: flex-end;
          align-items: center;
          min-width: 245px;

          > span {
            font-family: var(--font);
            font-size: 0.9rem;
            color: var(--light-text);
            margin-inline-start: 6px;
          }
        }

        .dropdown {
          margin-inline-start: 30px;
        }
      }
    }
  }
}

.is-dark {
  .list-view-v1 {
    .list-view-item {
      @include vuero-card--dark;

      .list-view-item-inner {
        .meta-left {
          h3 {
            color: var(--dark-dark-text) !important;
          }
        }

        .meta-right {
          .stats {
            .stat {
              span {
                &:first-child {
                  color: var(--dark-dark-text);
                }
              }
            }

            .separator {
              border-color: var(--dark-sidebar-light-16) !important;
            }
          }
        }
      }
    }
  }
}

@media only screen and (width <= 767px) {
  .list-view-v1 {
    .list-view-item {
      .list-view-item-inner {
        position: relative;
        flex-direction: column;

        .v-avatar {
          margin-bottom: 10px;
        }

        .meta-left {
          margin-inline-start: 0;
        }

        .meta-right {
          flex-direction: column;
          margin-inline-start: 0;

          .tags {
            margin: 10px 0;
          }

          .stats {
            margin: 10px 0;
          }

          .network {
            margin: 10px 0 0;
            justify-content: center;

            > span {
              display: none;
            }
          }

          .dropdown {
            position: absolute;
            top: 0;
            inset-inline-end: 0;
            margin-inline-start: 0;
          }
        }
      }
    }
  }
}

@media only screen and (width >= 768px) and (width <= 1024px) and (orientation: portrait) {
  .list-view-v1 {
    display: flex;
    flex-wrap: wrap;

    .list-view-item {
      margin: 10px;
      width: calc(50% - 20px);

      .list-view-item-inner {
        position: relative;
        flex-direction: column;

        .v-avatar {
          margin-bottom: 10px;
        }

        .meta-left {
          margin-inline-start: 0;
        }

        .meta-right {
          flex-direction: column;
          margin-inline-start: 0;

          .tags {
            margin: 10px 0;
          }

          .stats {
            margin: 10px 0;
          }

          .network {
            margin: 10px 0 0;
            justify-content: center;

            > span {
              display: none;
            }
          }

          .dropdown {
            position: absolute;
            top: 0;
            inset-inline-end: 0;
            margin-inline-start: 0;
          }
        }
      }
    }
  }
}
</style>
