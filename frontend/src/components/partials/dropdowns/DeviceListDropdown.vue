<script lang="ts" setup>
// import { DeviceStatus } from '/@src/utils/enums'
import { confirm } from "/@src/utils/dialog";
import { useNotyf } from "/@src/composable/useNotyf";
import { useDevices } from '/@src/stores/device'
import { useClustNode } from '/@src/stores/node'
import { useTemplate } from '/@src/stores/template'
import { DeviceStatus } from "/@src/utils/enums-dic";

const props = defineProps<{
  device: DeviceDetail
}>()

const emit = defineEmits<{
  refresh: []
}>()

const notyf = useNotyf();
const deviceStore = useDevices();
const nodeStore = useClustNode();
const tmplStore = useTemplate();

function remove() {
  confirm({
    title: "删除设备",
    content: "确定要删除该设备吗？",
    onConfirm: async (hide) => {
      const res = await deviceStore.$remove(props.device.id, props.device.targetNode);
      hide();
      if (res && res.result === "ok") {
        notyf.success("删除成功");
        emit('refresh')
      } else if (res.message) {
        notyf.error(res.message)
      }
    },
  });
}

const stoppable = computed(() => {
  return props.device.status === DeviceStatus.Running
})

function stop() {
  if (!stoppable.value) return
  confirm({
    title: "停止设备",
    content: "确定要停止该设备吗？",
    onConfirm: async (hide) => {
      const res = await deviceStore.$stop(props.device.id);
      hide();
      if (res && res.result === "ok") {
        notyf.success("停止成功");
        emit('refresh')
      } else if (res?.error) {
        notyf.error(res.message)
      }
    },
  });
}

const startable = computed(() => {
  return [DeviceStatus.Failed, DeviceStatus.Unavailable, DeviceStatus.Terminated].includes(props.device.status as DeviceStatus)
})

function start() {
  if (!startable.value) return
  confirm({
    title: "启动设备",
    content: "确定要启动该设备吗？",
    onConfirm: async (hide) => {
      const res = await deviceStore.$start(props.device.id);
      hide();
      if (res && res.device) {
        notyf.success("启动成功");
        emit('refresh')
      } else if (res?.error) {
        notyf.error(res.message)
      }
    },
  });
}

async function config() {
  if (!startable.value) {
    notyf.error("设备正在运行中，无法更新配置")
    return
  }
  const { id } = props.device
  await nodeStore.$fetchList()
  const res = await Promise.all([
    deviceStore.$getById(id),
    nodeStore.$fetchList(),
    tmplStore.$fetchList()
  ])
  bus.trigger(Signal.OpenResourceConfig, {
    device: res[0],
    callbacks: {
      success: () => {
        emit('refresh')
      }
    }
  })
}
</script>
<template>
  <VDropdown
    icon="feather:more-vertical"
    spaced
    right
    class="device-list-dropdown"
  >
    <template #content="{ close }">
      <a
        href="#"
        role="menuitem"
        class="dropdown-item is-media"
        @click="close(); config()"
      >
        <div class="icon">
          <i
            class="iconify"
            data-icon="feather:cpu"
            aria-hidden="true"
          />
        </div>
        <div class="meta">
          <span>配置</span>
        </div>
      </a>

      <hr v-if="startable || stoppable" class="dropdown-divider">
      <!-- <a
        href="#"
        role="menuitem"
        class="dropdown-item is-media"
        @click="close(); viewContainer()"
      >
        <div class="icon"><i
          class="iconify"
          data-icon="feather:package"
          aria-hidden="true"
        />
        </div>
        <div class="meta">
          <span>查看容器</span>
        </div>
      </a> -->
      <template v-if="stoppable">
        <a
          href="#"
          role="menuitem"
          class="dropdown-item is-media"
          @click="close(); stop()"
        >
          <div class="icon is-primary">
            <i
              class="iconify"
              data-icon="feather:toggle-right"
              aria-hidden="true"
            />
          </div>
          <div class="meta">
            <span>停止</span>
          </div>
        </a>
      </template>
      <template v-if="startable">
        <a
          href="#"
          role="menuitem"
          class="dropdown-item is-media"
          @click="close(); start()"
        >
          <div class="icon">
            <i
              class="iconify"
              data-icon="feather:toggle-left"
              aria-hidden="true"
            />
          </div>
          <div class="meta">
            <span>启动</span>
          </div>
        </a>
      </template>

      <hr class="dropdown-divider">

      <a
        href="#"
        role="menuitem"
        class="dropdown-item is-media"
        @click="close(); emit('refresh')"
      >
        <div class="icon">
          <i
            class="iconify"
            data-icon="feather:refresh-cw"
            aria-hidden="true"
          />
        </div>
        <div class="meta">
          <span>刷新</span>
        </div>
      </a>

      <hr class="dropdown-divider">

      <a
        href="#"
        role="menuitem"
        class="dropdown-item is-media"
        @click="close(); remove()"
      >
        <div class="icon">
          <i
            class="iconify"
            data-icon="feather:trash-2"
            aria-hidden="true"
          />
        </div>
        <div class="meta">
          <span>删除</span>
        </div>
      </a>
    </template>
  </VDropdown>
</template>
<style lang="scss">
.dropdown.is-spaced.device-list-dropdown {
  .dropdown-menu {
    min-width: 140px;
  }

  .dropdown-content .is-media .icon.is-primary svg {
    color: var(--primary);
  }

}
</style>
