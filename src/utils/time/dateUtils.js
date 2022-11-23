import dayJs from "dayjs";
import relativeTime from "dayjs/plugin/relativeTime";
import duration from "dayjs/plugin/duration";
import "dayjs/locale/zh-cn";
import { useI18n } from "vue-i18n";
dayJs.extend(relativeTime);
dayJs.extend(duration);

export function changeDayJsLocale(language) {
  if (language == undefined || language == "zh") {
    dayJs.locale("zh-cn");
  } else {
    dayJs.locale("en");
  }
}

export function fromNowexecutionTime(time) {
  const { t } = useI18n();
  return dayJs(time).fromNow() + " " + t("datetime.operation");
}

export function formatDurationTime(ms) {
  const { t } = useI18n();
  const duration = dayJs.duration(ms).$d;
  const collection = [
    ["year", duration.year],
    ["months", duration.months],
    ["days", duration.days],
    ["hours", duration.hours],
    ["minutes", duration.minutes],
    ["seconds", duration.year],
  ];
  const validCollection = collection.filter(([_unit, value]) => value > 0);
  const durationText = validCollection
    .map(([unit, value]) => value + t(`datetime.${unit}`))
    .join(" ");

  return t("datetime.elapsedTime") + " " + durationText;
}

export default dayJs;
