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

interface GetPipelineDetailParams {
  name?: string;
  id?: number|string;
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


// /pipeline/:name/detail/:id
export function apiGetPipelineDetail(params:GetPipelineDetailParams) {
  return httpRequest({
    // url: `/pipeline/${params.name}/detail/${params.id}`,
    url:"https://console-mock.apipost.cn/mock/ae73cd30-20d8-4975-b034-48b34891e956/pipeline/:name/detail/:id?apipost_id=6bbbe6",
    method: "get",
    params: params,
  });
}
