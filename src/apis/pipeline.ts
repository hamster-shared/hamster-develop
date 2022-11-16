// 导入axios实例
import httpRequest from "@/request/index";

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
