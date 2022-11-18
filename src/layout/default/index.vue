<template>
  <div>
    <!-- Default Layout -->
    <a-layout class="layout-default" id="layout-default" :class="layoutClass">
      <!-- Layout Header's Conditionally Fixed Wrapper -->
      <DefaultHeader />
      <!-- / Layout Header's Conditionally Fixed Wrapper -->

      <!-- Layout Content -->
      <a-layout>
        <!-- Page Content -->
        <a-layout-content class="layout-default-content">
          <keep-alive>
            <router-view v-if="$route.meta.keepAlive" />
          </keep-alive>
          <router-view v-if="!$route.meta.keepAlive" />
        </a-layout-content>
        <!-- / Page Content -->
      </a-layout>
      <!-- / Layout Content -->
    </a-layout>
    <!-- / Dashboard Layout -->
  </div>
</template>
<script lang="ts" setup>
import DefaultHeader from "../../components/DefaultHeader/index.vue";
import { RouterView, useRoute } from "vue-router";
import { computed } from "vue";
const routes = useRoute();
const layoutClass = computed(() => {
  return routes.meta.layoutClass;
});
</script>
<style lang="less" scoped>
.ant-layout {
  background: #f9f9f9;
}

.layout-default-content {
  width: 100%;
  padding: 24px;
  // max-width: 1280px;
  // margin: 24px auto;
  background-color: #f5f5f5;
  border-radius: 12px;
  border: 1px solid transparent;
}
</style>
