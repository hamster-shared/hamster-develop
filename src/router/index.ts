import { createRouter, createWebHistory } from "vue-router";
import HomeView from "../views/home/HomeView.vue";
import PipelineIndex from "../views/pipeline/index/index.vue";
import Process from "../views/pipeline/process/index.vue";
import Stage from "../views/pipeline/stage/index.vue";
import Create from "../views/pipeline/create/index.vue";
import Edit from "../views/pipeline/edit/index.vue";

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
      path: "/pipeline/:id/:id",
      name: "process",
      component: Process,
    },
    {
      path: "/create",
      name: "create",
      component: Create,
    },
    {
      path: "/edit",
      name: "edit",
      component: Edit,
    },
  ],
});

export default router;
