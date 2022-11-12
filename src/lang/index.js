import { createI18n } from "vue-i18n";
import en from "./en/index";
import zh from "./zh/index";

const messages = {
  en: {
    ...en,
  },
  zh: {
    ...zh,
  },
};

const i18n = new createI18n({
  globalInjection: true,
  locale: window.localStorage.getItem("language") || "en",
  messages,
  legacy: false,
});

export default i18n;
