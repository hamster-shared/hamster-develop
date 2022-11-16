// 导入axios实例
import httpRequest from "@/request/index";

// 定义接口的传参(例子)
interface PipelineInfoParam {
  userID: string;
  userName: string;
}

// 获取模板列表
export function apiGetTemplates() {
  return httpRequest({
    url: "/pipeline/templates",
    // url: "https://console-mock.apipost.cn/mock/ae73cd30-20d8-4975-b034-48b34891e956/pipeline/templates?apipost_id=1b2afa",
    method: "get",
  });
}
