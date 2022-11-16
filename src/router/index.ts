import { createRouter, createWebHistory } from "vue-router";
import HomeView from "../views/home/HomeView.vue";
import PipelineIndex from "../views/pipeline/index/index.vue";
import Process from "../views/pipeline/process/index.vue";
import Stage from "../views/pipeline/stage/index.vue";
import AllLogs from "../views/pipeline/allLogs/index.vue";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/home",
      name: "home",
      component: HomeView,
    },
    {
      path: "/",
      redirect: "/home",
    },
    {
      path: "/pipeline",
      name: "Pipeline",
      component: PipelineIndex,
    },
    {
      path: "/process",
      name: "process",
      component: Process,
      meta: {
        title: '详情',
        layout: "default",
      }
    },
    {
      path: "/stage",
      name: "stage",
      component: Stage,
    },
    {
      path: "/allLogs",
      name: "allLogs",
      component: AllLogs,
      meta: {
        title: '全部日志',
        layout: "null",
      }
    },
  ],
});

export default router;
