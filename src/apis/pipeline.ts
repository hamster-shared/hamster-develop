// 导入axios实例
import httpRequest from "@/request/index";

interface GetPipelinesParams {
  query?: string;
  page?: number;
  size?: number;
}

interface GetPipelineInfoParams {
  page?: number;
  size?: number;
}

export function apiGetPipelines(params: GetPipelinesParams) {
  return httpRequest({
    url: "/pipeline",
    method: "get",
    params: params,
  });
}

export function apiGetPipelineInfo(
  name: string,
  params: GetPipelineInfoParams
) {
  return httpRequest({
    url: `/pipeline/${name}`,
    method: "get",
    params: params,
  });
}

export function apiDeletePipelineInfo(name: string) {
  return httpRequest({
    url: `/pipeline/${name}`,
    method: "delete",
  });
}

export function apiOperationStopPipeline(name: string) {
  return httpRequest({
    url: `/pipeline/exec/${name}`,
    method: "post",
    data: name,
  });
}
