// 导入axios实例
import httpRequest from "@/request/index";

// 获取模板列表
export function apiGetTemplates() {
  return httpRequest({
    url: "/pipeline/templates",
    method: "get",
  });
}

export function apiGetTemplatesById(id: String) {
  return httpRequest({
    // url: `/pipeline/templates/${id}`,
    url: "https://console-mock.apipost.cn/mock/ae73cd30-20d8-4975-b034-48b34891e956/pipeline/template/:id?apipost_id=c60d54",
    method: "get",
  });
}

// 添加
export function apiAddPipeline(name: String, yaml: String) {
  return httpRequest({
    url: "/pipeline",
    method: "post",
    data: {
      name: name,
      yaml: yaml,
    },
  });
}

//修改 
export function apiEditPipeline(oldName: String, newName: String, yaml: String) {
  return httpRequest({
    url: `/pipeline/${oldName}`,
    method: "put",
    data: {
      newName: newName,
      yaml: yaml,
    },
  });
}
