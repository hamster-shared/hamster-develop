import { createRouter, createWebHistory } from "vue-router";
import HomeView from "../views/home/HomeView.vue";
import PipelineIndex from "../views/pipeline/index/index.vue";
import Process from "../views/pipeline/process/index.vue";
import Stage from "../views/pipeline/stage/index.vue";
import Template from "../views/pipeline/template/index.vue";
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
      path: "/process",
      name: "process",
      component: Process,
    },
    {
      path: "/stage",
      name: "stage",
      component: Stage,
    },
    {
      path: "/template",
      name: "template",
      component: Template,
    },
    {
      path: "/create/config",
      name: "CreateConfig",
      component: CreatePipeline,
    },
  ],
});

export default router;
