<script setup lang="ts">
import { useNotyf } from "@src/composable/useNotyf";
import { useTemplate } from "@src/stores/template";
import { useTemplateType } from "@src/stores/templateType";
import { confirm } from "@src/utils/dialog";
import sleep from "@src/utils/sleep";

export interface DeviceTemplateData {
  id: string;
  name: string;
  description: string;
}

const router = useRouter();
const tmplStore = useTemplate();
const tmpls = computed(() => tmplStore.tmpls);
const tmplTypeStore = useTemplateType();
const tmplTypeMap = computed(() =>
  tmplTypeStore.tmplTypes.reduce((acc, type) => {
    acc[type.id] = type;
    return acc;
  }, {} as Record<string, any>)
);
const notyf = useNotyf();
const loading = ref(false);

const filters = ref("");
const tmplType = ref("");

const filteredData = computed(() => {
  let filtered;
  if (!filters.value) {
    filtered = tmpls.value;
  } else {
    filtered = tmpls.value.filter((item) => {
      return item.name.match(new RegExp(filters.value, "i"));
    });
  }
  if (!tmplType.value) {
    return filtered;
  } else {
    return filtered.filter((item) => {
      return item.type === tmplType.value;
    });
  }
});

function createTemplate() {
  router.push("/app/template/new");
}

function localImport() {
  bus.trigger(Signal.OpenTemplateImport);
}

function cloudImport() {
  notyf.dismissAll();
  notyf.info("即将支持，敬请期待");
}

function auth() {
  notyf.dismissAll();
  notyf.info("即将支持，敬请期待");
}

function deleteTmpl(tmpl: TemplateData) {
  confirm({
    title: "删除应用",
    content: `确定要删除该应用 ${tmpl.name} 吗？`,
    onConfirm: async (hide) => {
      const res = await tmplStore.$deleteTmpl(tmpl.id);
      if (res === "success") {
        hide();
        notyf.success("删除成功");
      }
    },
  });
}

function listTmpl(tmpl: TemplateData) {
  const action = tmpl.listed ? "下架" : "上架";
  confirm({
    title: `${action}应用`,
    content: `确定要${action}该应用吗？`,
    onConfirm: async (hide) => {
      const res = await tmplStore.$listed(tmpl);
      if (res === "success") {
        hide();
        notyf.success(`${action}成功`);
      }
    },
  });
}

(async () => {
  loading.value = true;
  await Promise.all([tmplTypeStore.$fetchList(), tmplStore.$fetchList()]);
  await sleep(1000);
  loading.value = false;
})();
</script>

<template>
  <div>
    <div class="card-grid-toolbar tmpl-list">
      <VControl icon="feather:search">
        <input v-model="filters" class="input custom-text-filter" placeholder="搜索...">
      </VControl>
      <TmplTypeSelect v-model="tmplType" />

      <div class="buttons">
        <VButtons>
          <VButton color="primary" raised @click="createTemplate">
            <span class="icon">
              <i aria-hidden="true" class="fas fa-plus" />
            </span>
            <span>新建应用</span>
          </VButton>
          <!-- <VButton color="primary" raised @click="localImport">
            <span class="icon">
              <i aria-hidden="true" class="fas fa-file-import" />
            </span>
            <span>本地导入</span>
          </VButton>
          <VButton color="primary" raised @click="cloudImport">
            <span class="icon">
              <i class="fas fa-cloud-download-alt" aria-hidden="true" />
            </span>
            <span>云市场导入</span>
          </VButton> -->
          <VButton color="primary" raised @click="auth">
            <span class="icon">
              <i aria-hidden="true" class="fas fa-plus" />
            </span>
            <span>应用授权</span>
          </VButton>
        </VButtons>
      </div>
    </div>

    <div class="card-grid card-grid-v2">
      <!--List Empty Search Placeholder -->
      <VPlaceholderPage
        :class="[(filteredData.length !== 0 || loading) && 'is-hidden']"
        title="没有找到符合搜索条件的应用"
        subtitle="很遗憾，看起来我们无法找到与您输入的搜索词或条件匹配的应用。请尝试不同的搜索词或条件。"
        larger
      />

      <!--Card Grid v2-->
      <TransitionGroup v-if="!loading" name="list" tag="div" class="columns is-multiline">
        <!--Grid Item-->
        <div v-for="(item, key) in filteredData" :key="key" class="column is-4">
          <div class="card-grid-item">
            <div class="card">
              <header class="card-header">
                <div
                  tabindex="0"
                  class="card-header-title"
                  role="button"
                  @keydown.space.prevent="tmplStore.navigate(item.id)"
                  @click.prevent="tmplStore.navigate(item.id)"
                >
                  <div class="meta">
                    <span class="dark-inverted">{{ item.name }}</span>
                  </div>
                </div>
                <div class="card-header-suffix">
                  <VIconWrap
                    size="small"
                    :picture="tmplTypeMap[item.type]?.icon"
                    class="mr-2 is-tt-icon"
                  />
                  {{ tmplTypeMap[item.type]?.name }}
                </div>
              </header>
              <div class="card-content">
                <div class="card-content-flex">
                  <div class="card-info">
                    <p v-if="item.description">{{ item.description }}</p>
                    <p v-else><em>未添加描述</em></p>
                  </div>
                </div>
              </div>
              <footer class="card-footer">
                <a
                  tabindex="-1"
                  role="button"
                  @keydown.space.prevent="deleteTmpl(item)"
                  @click.prevent="deleteTmpl(item)"
                >
                  删除
                </a>
                <a
                  tabindex="-1"
                  role="button"
                  @keydown.space.prevent="tmplStore.navigate(item.id)"
                  @click.prevent="tmplStore.navigate(item.id)"
                >
                  查看
                </a>
                <span style="text-align: center;" :class="item.listed ? 'is-success' : 'is-danger'">{{
                  item.listed ? "已授权" : "未授权"
                }}</span>
                <!-- <a
                  tabindex="0"
                  role="button"
                  class="with-icon"
                  :class="item.listed ? 'is-danger' : 'is-warning'"
                  @keydown.space.prevent="listTmpl(item)"
                  @click.prevent="listTmpl(item)"
                >
                  <i
                    class="fas"
                    :class="[
                      item.listed ? 'fa-long-arrow-alt-down' : 'fa-long-arrow-alt-up',
                    ]"
                    aria-hidden="true"
                  />
                  {{ item.listed ? "下架" : "上架" }}
                </a> -->
                <!-- <a
                  tabindex="0"
                  role="button"
                  @keydown.space.prevent="tmplStore.navigate(item.id)"
                  @click.prevent="tmplStore.navigate(item.id)"
                  >编辑</a
                > -->
              </footer>
            </div>
          </div>
        </div>
      </TransitionGroup>

      <VPlaceloadWrap v-if="loading" class="columns is-multiline">
        <div class="column is-4">
          <VPlaceload height="134px" width="100%" />
        </div>
        <div class="column is-4">
          <VPlaceload height="134px" width="100%" />
        </div>
        <div class="column is-4">
          <VPlaceload height="134px" width="100%" />
        </div>
      </VPlaceloadWrap>
    </div>
  </div>
  <NewTemplateModal />
  <TemplateImportModal />
</template>

<style lang="scss">
.card-grid-toolbar.tmpl-list {
  .field {
    min-width: 180px;
  }
}

.card-grid {
  .columns {
    margin-inline-start: -0.5rem !important;
    margin-inline-end: -0.5rem !important;
    margin-top: -0.5rem !important;
  }

  .column {
    padding: 0.5rem !important;
  }
}

.card-grid-v2 {
  .card-grid-item {
    .card {
      border: 1px solid var(--fade-grey-dark-4);
      box-shadow: none;
      border-radius: var(--radius-large);

      .card-header {
        box-shadow: none;
        border-bottom: 1px solid var(--fade-grey-dark-4);

        .card-header-title {
          display: flex;
          align-items: center;
          cursor: pointer;

          .meta {
            line-height: 1.2;

            span {
              display: block;
              font-weight: 400;

              &:hover {
                text-decoration: underline;
              }

              &:first-child {
                font-family: var(--font-alt);
                font-size: 0.95rem;
                color: var(--dark-text);
                font-weight: 600;
              }

              &:nth-child(2) {
                font-size: 0.9rem;
                color: var(--light-text);
              }
            }
          }
        }

        .card-header-suffix {
          font-size: 0.75rem;
          color: var(--light-text);
          font-weight: 500;
          display: flex;
          align-items: center;
          justify-content: flex-end;
          margin-inline-end: 10px;
        }
      }

      .card-image {
        img {
          object-fit: cover;
        }
      }

      .card-content {
        padding: 1rem;

        .card-content-flex {
          display: flex;
          align-items: center;
          justify-content: space-between;

          .card-info {
            h3 {
              font-family: var(--font-alt);
              font-size: 1rem;
              color: var(--dark-text);
              font-weight: 600;
            }

            p {
              font-size: 0.9rem;

              svg {
                position: relative;
                top: 0;
                height: 14px;
                width: 14px;
                margin-inline-end: 4px;
              }
            }
          }
        }
      }

      .card-footer {
        span {
          padding: 1rem 0.75rem;
          flex: 1 1 0%;
          text-align: right;

          &.is-danger {
            color: var(--danger);
          }

          &.is-success {
            color: var(--success);
          }
        }

        a {
          font-family: var(--font);
          color: var(--light-text);
          padding: 1rem 0.75rem;
          transition: all 0.3s; // transition-all test
          flex: 1 1 0%;
          text-align: center;

          &:hover {
            background: var(--fade-grey-light-4);
            color: var(--primary);
          }

          &.with-icon {
            display: flex;
            align-items: center;
            justify-content: center;

            i {
              margin-inline-end: 8px;
            }

            &.is-danger {
              color: var(--danger);
            }

            &.is-warning {
              color: var(--warning);
            }
          }
        }
      }
    }
  }
}

.is-dark {
  .card-grid-v2 {
    .card-grid-item {
      border-color: var(--dark-sidebar-light-12);

      .card {
        background: var(--dark-sidebar-light-6);
        border-color: var(--dark-sidebar-light-12);

        .card-header {
          border-color: var(--dark-sidebar-light-12);
        }

        .card-content {
          border-color: var(--dark-sidebar-light-12);

          .avatar-stack {
            .avatar {
              border-color: var(--dark-sidebar-light-6);
            }
          }
        }

        .card-footer {
          border-color: var(--dark-sidebar-light-12);

          a {
            border-color: var(--dark-sidebar-light-12);

            &:hover {
              background: var(--dark-sidebar-light-2);
              color: var(--primary);
            }
          }
        }
      }
    }
  }
}
</style>
