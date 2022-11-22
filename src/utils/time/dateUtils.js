import dayJs from "dayjs";
import relativeTime from "dayjs/plugin/relativeTime";
import "dayjs/locale/zh-cn";
dayJs.extend(relativeTime);

const language = window.localStorage.getItem("language");
if (language == undefined || language == "zh") {
  dayJs.locale("zh-cn");
} else {
  dayJs.locale("en");
}

export default function executionTime(time) {
  return dayJs(time).fromNow();
}
