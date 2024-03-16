<template>
  <div class="mobile-subsidebar">
    <div class="inner">
      <ul>
        <li>
          <RouterLink to="/app/template">
            <i
              class="mr-1 fas fa-shopping-bag"
              aria-hidden="true"
            />
            应用商店
          </RouterLink>
        </li>
        <li>
          <RouterLink to="/app/devices">
            <i
              class="mr-1 fas fa-layer-group"
              aria-hidden="true"
            />
            设备管理
          </RouterLink>
        </li>
        <li>
          <RouterLink to="/app/resource">
            <i
              class="mr-1 fas fa-server"
              aria-hidden="true"
            />
            资源管理
          </RouterLink>
        </li>
        <li>
          <RouterLink to="/app/config">
            <i
              class="mr-1 fas fa-users-cog"
              aria-hidden="true"
            />
            系统配置
          </RouterLink>
        </li>
      </ul>
    </div>
  </div>
</template>
