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
    method: "get",
  });
}
