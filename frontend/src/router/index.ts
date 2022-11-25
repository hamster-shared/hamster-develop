import { createRouter, createWebHashHistory } from "vue-router";
import HomeView from "../views/home/HomeView.vue";
import RpcsIndex from "../views/nodeService/rpcs/index.vue";
import AppsIndex from "../views/nodeService/apps/index.vue";

const router = createRouter({
  history: createWebHashHistory(),
  routes: [
    {
      path: "/home",
      name: "home",
      component: HomeView,
    },
    {
      path: "/",
      redirect: "/RPCs",
    },
    {
      path: "/RPCs",
      name: "RPCs",
      component: RpcsIndex,
    },
    {
      path: "/apps",
      name: "AppsIndex",
      component: AppsIndex,
    },
  ],
});

export default router;
