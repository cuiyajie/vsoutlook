<script lang="ts" setup>
// import { DeviceStatus } from '/@src/utils/enums'
import { confirm } from "/@src/utils/dialog";
import { useNotyf } from "/@src/composable/useNotyf";
import { useDevices } from '/@src/stores/device'

const props = defineProps<{
  device: DeviceDetail
}>()

const emit = defineEmits<{
  refresh: []
}>()

const notyf = useNotyf();
const deviceStore = useDevices();

function remove() {
  confirm({
    title: "删除设备",
    content: "确定要删除该设备吗？",
    onConfirm: async (hide) => {
      const res = await deviceStore.$remove(props.device.id);
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

function reboot() {
  confirm({
    title: "重启设备",
    content: "确定要重启该设备吗？",
    onConfirm: async (hide) => {
      const res = await deviceStore.$reboot(props.device);
      hide();
      if (res && res.result === "success") {
        notyf.success("重启成功");
        emit('refresh')
      } else if (res?.result === "error") {
        notyf.error(res.message)
      }
    },
  });
}

function config() {
  const { tmplID, tmplName, tmplTypeName, id, nodeName } = props.device
  bus.trigger(Signal.OpenResourceConfig, {
    tmpl: {
      id: tmplID,
      name: tmplName,
      typeName: tmplTypeName
    },
    device: {
      id,
      node: nodeName,
    },
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

      <template v-if="device.phase === 'Running'">
        <hr class="dropdown-divider">
        <a
          href="#"
          role="menuitem"
          class="dropdown-item is-media"
          @click="close(); reboot()"
        >
          <div class="icon">
            <i
              class="iconify"
              data-icon="feather:arrow-right-circle"
              aria-hidden="true"
            />
          </div>
          <div class="meta">
            <span>重启</span>
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
    min-width: 110px;
  }
}
</style>
