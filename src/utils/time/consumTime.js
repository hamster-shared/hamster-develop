// import dayJs from "dayjs";
// import duration from "dayjs/plugin/duration";
// import "dayjs/locale/zh-cn";
// dayJs.extend(duration);

// const language = window.localStorage.getItem("language");
// if (language == undefined || language == "zh") {
//   dayJs.locale("zh-cn");
// } else {
//   dayJs.locale("en");
// }

// export default function formatDurationTime(mss) {
//   const time = dayJs.duration(mss).$d;
//   console.log("dayJs", time);
//   dayJs.duration({
//     seconds: time.seconds,
//     minutes: time.minutes,
//     hours: time.hours,
//     days: time.days,
//     months: time.months,
//     years: time.years,
//   });
// }

export default function formatDurationTime(mss) {
  const days = parseInt(mss / (1000 * 60 * 60 * 24));
  const hours = parseInt((mss % (1000 * 60 * 60 * 24)) / (1000 * 60 * 60));
  const minutes = parseInt((mss % (1000 * 60 * 60)) / (1000 * 60));
  const seconds = parseInt((mss % (1000 * 60)) / 1000);

  const language = window.localStorage.getItem("language");

  if (days > 0) {
    if (language == undefined || language == "zh") {
      return (
        "耗时" +
        days +
        " 天 " +
        hours +
        " 小时 " +
        minutes +
        " 分钟 " +
        seconds +
        " 秒 "
      );
    } else {
      return (
        "Elapsed time " +
        days +
        " d " +
        hours +
        " h " +
        minutes +
        " m " +
        seconds +
        " s "
      );
    }
  } else {
    if (hours > 0) {
      if (language == undefined || language == "zh") {
        return (
          "耗时" + hours + " 小时 " + minutes + " 分钟 " + seconds + " 秒 "
        );
      } else {
        return (
          "Elapsed time " + hours + " h " + minutes + " m " + seconds + " s "
        );
      }
    } else {
      if (minutes > 0) {
        if (language == undefined || language == "zh") {
          return "耗时" + minutes + " 分钟 " + seconds + " 秒 ";
        } else {
          return "Elapsed time " + minutes + " m " + seconds + " s ";
        }
      } else {
        if (language == undefined || language == "zh") {
          return "耗时" + seconds + " 秒 ";
        } else {
          return "Elapsed time " + seconds + " s ";
        }
      }
    }
  }
}
