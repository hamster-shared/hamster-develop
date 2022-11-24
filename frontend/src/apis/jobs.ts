import httpRequest from "@/request/index";

interface GegJobLogsParams {
  id: number | string;
  name: string;
}

interface GegJobStagelogsParams {
  id: number | string;
  name: string;
  stagename: string,
  start: number,
  lastLine: number,
}


// 查看所有日志  /pipeline/:name/logs/:id
export function apiGetAllJobLogs(params: GegJobLogsParams) {
  return httpRequest({
    url: `/pipeline/${params.name}/logs/${params.id}`,
    // url:'https://console-mock.apipost.cn/mock/ae73cd30-20d8-4975-b034-48b34891e956/pipeline/:name/logs/:id?apipost_id=a005bc',
    method: "get",
    params: params,
  });
}

//  获取指定stage日志 /pipeline/:name/logs/:id/:stagename
export function apiGetJobStageLogs(params: GegJobStagelogsParams) {
  return httpRequest({
    url: `/pipeline/${params.name}/logs/${params.id}/${params.stagename}`,
    // url:'https://console-mock.apipost.cn/mock/ae73cd30-20d8-4975-b034-48b34891e956/pipeline/:name/logs/:id/:stagename?apipost_id=510db1',
    method: "get",
    params: params,
  });
}