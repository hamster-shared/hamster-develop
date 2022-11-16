// 导入axios实例
import httpRequest from "@/request/index";

// 定义接口的传参(例子)
interface PipelineInfoParam {
  userID: string;
  userName: string;
}

// 获取用户信息（例子）
export function apiGetUserInfo(param: PipelineInfoParam) {
  return httpRequest({
    url: "your api url",
    method: "post",
    data: param,
  });
}

interface GetPipelineInfoParams {
  query?: string;
  page?: number;
  size?: number;
}
export function apiGetPipelines(params: GetPipelineInfoParams) {
  return httpRequest({
    url: "/pipeline",
    method: "get",
    params: params,
  });
}

export function apiGetPipelineInfo(name: string, params) {
  return httpRequest({
    url: `/pipeline/${name}`,
    method: "get",
    params: params,
  });
}
