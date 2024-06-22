<script setup lang="ts">
import { useUserSession } from "@src/stores/userSession";

const locked = useLocalStorage('endswitch-dashboard-locked', false)
const boards = ref(
  Array.from({ length: 4 }, () => ({
    actived: false,
    status: 0
  }))
)

const usStore = useUserSession();
const titles = computed(() => usStore.settings.endswt_titles || []);

function configure() {
  bus.trigger(Signal.OpenApiServerPort)
}

function updateEndSwtTitle(idx: number) {
  bus.trigger(Signal.OpenEndSwtTitle, idx)
}
</script>

<template>
  <div class="action-page-wrapper action-page-endswt">
    <div class="action-toolbar">
      <h1 class="title is-4">
        {{ titles[0] || '4K 频道1末级切换' }}
        <span
          class="icon ml-2"
          style="cursor: pointer;"
          role="button"
          tabindex="-1"
          @keyup.space.prevent="updateEndSwtTitle(0)"
          @click.prevent="updateEndSwtTitle(0)"
        >
          <i aria-hidden="true" class="fas fa-edit" />
        </span>
      </h1>
      <VButtons>
        <VField class="lock-switch">
          <VControl>
            <VSwitchBlock
              v-model="locked"
              :label="locked ? '已锁定' : '可操作'"
              thin
              color="primary"
            />
          </VControl>
        </VField>
        <VButton
          color="primary"
          raised
          :disabled="locked"
          @click="configure"
        >
          <span class="icon">
            <i
              class="fas fa-cog"
              aria-hidden="true"
            />
          </span>
          <span>配置对外API地址</span>
        </VButton>
      </VButtons>
    </div>
    <div class="wrapper-inner" :class="[locked && 'is-locked']">
      <div v-for="(board, idx) in boards" :key="idx" class="action-board">
        <div class="board-status" :class="[board.actived && board.status === 0 && 'is-success', board.actived && board.status === 1 && 'is-danger']" />
        <VButton
          :color="board.actived ? 'success' : 'dark'"
          :disabled="locked"
        >
          {{ idx + 1 }}
        </VButton>
        <h3>
          {{ titles[idx + 1] || `信号${idx+1}` }}
        </h3>
        <span
          class="icon"
          style="cursor: pointer; margin-top: -14px;"
          role="button"
          tabindex="-1"
          @keyup.space.prevent="updateEndSwtTitle(idx+1)"
          @click.prevent="updateEndSwtTitle(idx+1)"
        >
          <i aria-hidden="true" class="fas fa-edit" />
        </span>
      </div>
    </div>
  </div>
  <APIServerPortModal />
  <EndSwitchTitleModal />
</template>

<style lang="scss">
@import '@src/scss/abstracts/all';

.action-page-wrapper {
  &.action-page-endswt {
    margin-inline-start: 2.625rem;

    .action-toolbar {
      display: flex;
      justify-content: space-between;
      align-items: center;
      margin-bottom: 20px;

      > h1 {
        color: var(--primary-grey-light-10) !important;
      }

      .lock-switch {
        margin-inline-end: 40px;
        margin-block-end: 4px;

        .thin-switch-block .text {
          margin-inline-start: 14px;

          span {
            top: 0;
          }
        }
      }
    }

    .wrapper-inner {
      @include vuero-s-card;

      width: 100%;
      max-width: 800px;
      min-height: 300px;
      margin: 80px auto 20px;
      display: flex;
      justify-content: center;
      align-items: center;
      gap: 20px;
      overflow: hidden;

      &.is-locked {
        position: relative;
        &:before {
          position: absolute;
          content: "";
          cursor: not-allowed;
          top: 0;
          left: 0;
          right: 0;
          bottom: 0;
          z-index: 10;
          background-color: rgba(255, 255, 255, 0.1);
        }
      }
    }

    .action-board {
      display: flex;
      flex-direction: column;
      align-items: center;
      gap: 16px;
      flex-basis: 120px;

      .board-status {
        width: 12px;
        height: 12px;
        border-radius: 100%;
        background: var(--dark-sidebar-light-20);
        transition: all .3s;

        &.is-success {
          background: var(--success);
        }

        &.is-danger {
          background: var(--danger);
        }
      }

      .button.v-button {
        padding: 0;
        width: 80px;
        height: 80px;
        border-radius: 12px;
        font-size: 1.5rem;
        font-weight: bold;
      }
    }
  }
}

.is-dark {
  .action-page-wrapper {
    &.action-page-endswt {

      .wrapper-inner {
        @include vuero-card--dark;
      }
    }
  }
}
</style>
