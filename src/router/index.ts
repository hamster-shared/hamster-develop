import { createRouter, createWebHistory } from "vue-router";
import HomeView from "../views/home/HomeView.vue";
import PipelineIndex from "../views/pipeline/index/index.vue";
import Process from "../views/pipeline/process/index.vue";
import Stage from "../views/pipeline/stage/index.vue";
import Create from "../views/pipeline/create/index.vue";
import Edit from "../views/pipeline/edit/index.vue";
import AllLogs from "../views/pipeline/allLogs/index.vue";
import CreatePipeline from "../views/pipeline/create/config/index.vue";

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
      path: "/pipeline/:id",
      name: "stage",
      component: Stage,
    },
    {
      path: "/pipeline/:id/:name",
      name: "process",
      component: Process,
    },
    {
      path: "/stage",
      name: "stage",
      component: Stage,
    },
    {
      path: "/allLogs/:id/:name",
      name: "allLogs",
      component: AllLogs,
      meta: {
        title: "全部日志",
        layout: "null",
      },
    },
    {
      path: "/create",
      name: "create",
      component: Create,
    },
    {
      path: "/create/config/:id",
      name: "CreateConfig",
      component: CreatePipeline,
    },
  ],
});

export default router;
