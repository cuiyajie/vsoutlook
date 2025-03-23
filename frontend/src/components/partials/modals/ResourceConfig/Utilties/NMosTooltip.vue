<script lang="ts" setup>
defineProps<{
  nmos: NMosNode
}>()

const tippyAppendTo = () => document.body
</script>
<template>
  <tippy
    animation="scale"
    theme="light"
    placement="right"
    :arrow="true"
    arrow-type="round"
    max-width="none"
    :interactive="true"
    :append-to="tippyAppendTo"
  >
    <template #content>
      <div class="kv-info tippy-detail">
        <div class="info-block-inner is-dark-bordered-12">
          <div class="info-block-line">
            <h4 class="dark-inverted">
              ID
            </h4>
            <span>
              {{ nmos.id }}
            </span>
          </div>
          <div class="info-block-line">
            <h4 class="dark-inverted">
              Version
            </h4>
            <span>
              {{ nmos.version }}
            </span>
          </div>
          <div class="info-block-line">
            <h4 class="dark-inverted">
              Label
            </h4>
            <span>
              {{ nmos.label }}
            </span>
          </div>
          <div class="info-block-line">
            <h4 class="dark-inverted">
              Tags
            </h4>
          </div>
          <div class="info-block-line">
            <h4 class="dark-inverted">
              Description
            </h4>
            <span>
              {{ nmos.description }}
            </span>
          </div>
          <div class="info-block-line">
            <h4 class="dark-inverted">
              Address
            </h4>
            <span>
              {{ nmos.href }}
            </span>
          </div>
          <div class="info-block-line">
            <h4 class="dark-inverted">
              Hostname
            </h4>
            <span>
              {{ nmos.hostname }}
            </span>
          </div>
          <div class="info-block-line">
            <h4 class="dark-inverted">
              API Versions
            </h4>
            <span>
              <span v-for="v in nmos.api.versions" :key="v" class="span-label">{{ v }}</span>
            </span>
          </div>
          <div class="info-block-line has-table">
            <h4 class="dark-inverted">
              API Endpoints
            </h4>
            <table v-if="nmos.api.endpoints.length > 0">
              <thead>
                <th>Host</th>
                <th>Port</th>
                <th>Protocol</th>
              </thead>
              <tr v-for="(ep, eidx) in nmos.api.endpoints" :key="eidx">
                <td>{{ ep.host }}</td>
                <td>{{ ep.port }}</td>
                <td>{{ ep.protocol }}</td>
              </tr>
            </table>
          </div>
          <div class="info-block-line has-table">
            <h4 class="dark-inverted">
              Clocks
            </h4>
            <table v-if="nmos.clocks.length > 0">
              <thead>
                <th>Name</th>
                <th>Ref Type</th>
              </thead>
              <tr v-for="(clock, cidx) in nmos.clocks" :key="cidx">
                <td>{{ clock.name }}</td>
                <td>{{ clock.ref_type }}</td>
              </tr>
            </table>
          </div>
          <div class="info-block-line">
            <h4 class="dark-inverted">
              Services
            </h4>
          </div>
          <div class="info-block-line has-table">
            <h4 class="dark-inverted">
              Interfaces
            </h4>
            <table v-if="nmos.interfaces.length > 0">
              <thead>
                <th>Name</th>
                <th>Local Chassis ID</th>
                <th>Local Port ID</th>
              </thead>
              <tr v-for="(inf, fidx) in nmos.interfaces" :key="fidx">
                <td>{{ inf.name }}</td>
                <td>{{ inf.chassis_id }}</td>
                <td>{{ inf.port_id }}</td>
              </tr>
            </table>
          </div>
        </div>
      </div>
    </template>
    <slot></slot>
  </tippy>
</template>
