import { createApp } from "vue";
import { createPinia } from "pinia";
import Antd from "ant-design-vue";
import i18n from "./lang/index";

import App from "./App.vue";
import router from "./router";
import "./design/main.css";
import "ant-design-vue/dist/antd.css";

const app = createApp(App);

app.use(createPinia());
app.use(router);
app.use(Antd);
app.use(i18n);

app.mount("#app");
