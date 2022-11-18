<template>
  <div class="w-[80%] mx-auto">
    <div class="flex justify-between mb-6">
      <span class="text-3xl font-semibold">hamster-pipeline</span>
      <div>
        <a-button class="mr-2">设置</a-button>
        <a-button type="primary">立即执行</a-button>
      </div>
    </div>

    <a-table
      :columns="columns"
      :data-source="pipelineInfo"
      :pagination="pagination"
    >
      <template #headerCell="{ column }">
        <template v-if="column.key === 'status'">
          <span> Status </span>
        </template>
      </template>

      <template #bodyCell="{ column, record }">
        <template v-if="column.key === 'status'">
          <span v-if="record.status == 0">no data</span>
          <span v-if="record.status == 1">running</span>
          <span v-if="record.status == 2">passed</span>
          <span v-if="record.status == 3">failed</span>
          <span v-if="record.status == 4">stop</span>
        </template>
        <template v-else-if="column.key === 'execution_process'">
          <div v-if="record.execution_process[0].status == 0">
            <img :src="nodataSVG" />
          </div>
          <div v-if="record.execution_process[0].status == 1">
            <img :src="runnngSVG" />
          </div>
          <div v-if="record.execution_process[0].status == 2">
            <img :src="successSVG" />
          </div>
          <div v-if="record.execution_process[0].status == 3">
            <img :src="failedSVG" />
          </div>
          <div v-if="record.execution_process[0].status == 4">
            <img :src="stopSVG" />
          </div>
        </template>
        <template v-else-if="column.key === 'action'">
          <div v-if="record.status == 1">
            <a @click="handleStop(record.id)">终止</a>
          </div>
          <div v-else>
            <a @click="handleDelete(record.id)">删除</a>
          </div>
        </template>
      </template>
    </a-table>
  </div>
</template>

<script setup lang="ts">
import { reactive, onMounted } from "vue";
import { useRouter } from "vue-router";
import {
  apiGetPipelineInfo,
  apiDeletePipelineInfo,
  apiOperationStopPipeline,
} from "@/apis/pipeline";
import runnngSVG from "@/assets/icons/pipeline-running.svg";
import successSVG from "@/assets/icons/pipeline-success.svg";
import failedSVG from "@/assets/icons/pipeline-failed.svg";
import stopSVG from "@/assets/icons/pipeline-stop.svg";
import nodataSVG from "@/assets/icons/pipeline-no-data.svg";

const router = useRouter();

const columns = reactive([
  {
    title: "Status",
    dataIndex: "status",
    key: "status",
  },
  {
    title: "ID",
    dataIndex: "id",
    key: "id",
  },
  {
    title: "Trigger Info",
    dataIndex: "trigger_info",
    key: "trigger_info",
  },
  {
    title: "Stage",
    key: "execution_process",
    dataIndex: "execution_process",
  },
  {
    title: "Time",
    key: "time_consuming",
    dataIndex: "time_consuming",
  },
  {
    title: "Action",
    key: "action",
  },
]);

const pipelineInfo = reactive([
  {
    key: "1",
    status: "1",
    id: "#1",
    last_execution_time: "Linda手动触发master|b4a0d99",
    execution_process: ["nice", "developer"],
    time_consuming: "1小时前执行",
    trigger_info: "111",
  },
  {
    key: "2",
    status: "2",
    id: "#2",
    last_execution_time: "Linda手动触发master|b4a0d99",
    execution_process: ["loser"],
    time_consuming: "1小时前执行,耗时6分钟",
    trigger_info: "222",
  },
  {
    key: "3",
    status: "3",
    id: "#3",
    info: "Linda手动触发master|b4a0d99",
    execution_process: ["cool", "teacher"],
    time_consuming: "1小时前执行",
    trigger_info: "333",
  },
]);

const pagination = reactive({
  // 分页配置器
  pageSize: 10, // 一页的数据限制
  current: 1, // 当前页
  total: 10, // 总数
  hideOnSinglePage: false, // 只有一页时是否隐藏分页器
  showQuickJumper: true, // 是否可以快速跳转至某页
  // showSizeChanger: true, // 是否可以改变 pageSize
  // pageSizeOptions: ["10", "20", "30"], // 指定每页可以显示多少条
  // onShowSizeChange: (current, pagesize) => {
  //   // 改变 pageSize时的回调
  //   pagination.current = current;
  //   pagination.pageSize = pagesize;
  // },
  onChange: (current) => {
    // 切换分页时的回调，
    pagination.current = current;
    getPipelineInfo(current);
  },
  // showTotal: total => `总数：${total}人`, // 可以展示总数
});

const getPipelineInfo = async () => {
  console.log(router.currentRoute.value.params.id);
  const data = await apiGetPipelineInfo(":name", {});
  console.log("apiGetPipelineInfo", data);
  Object.assign(pipelineInfo, data.pipeline.jobs);
  pagination.pageSize = data.pagination.size;
  pagination.total = data.pagination.total;
};

const handleDelete = async (id) => {
  await apiDeletePipelineInfo(id);
  console.log("id", id);
};

const handleStop = async (id) => {
  await apiOperationStopPipeline(id);
  console.log("id", id);
};

onMounted(() => {
  getPipelineInfo();
});
</script>

<style lang="less" scoped>
.ant-btn {
  font-size: 16px;
  border-radius: 3px;
  width: 200px;
  height: 48px;
}
.ant-btn-primary {
  background: #6481dc;
  &:hover,
  &:focus {
    border-color: #6481dc;
    background: #6481dc;
  }
}
:deep(.ant-table-thead > tr > th) {
  background: #808692;
  height: 48px;
  text-align: center;
  color: #ffffff;
}
:deep(.ant-table table) {
  text-align: center;
}
:deep(.ant-pagination-options-quick-jumper input) {
  margin: 0px 0px 0px 8px !important;
}
</style>
