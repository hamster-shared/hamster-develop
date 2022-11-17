<template>
  <Header />
  <div class="w-[80%] mx-auto">
    <div class="flex justify-between mt-6 mb-2">
      <span>hamster-pipeline</span>
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
          <span v-if="record.status == 1">running</span>
          <span v-if="record.status == 2">passed</span>
          <span v-if="record.status == 3">failed</span>
          <span v-if="record.status == 4">stop</span>
        </template>
        <template v-else-if="column.key === 'execution_process'">
          <span> {{ record.execution_process[0].status }} </span>
        </template>
        <template v-else-if="column.key === 'action'">
          <div v-if="record.status == 1">
            <a>终止</a>
          </div>
          <div v-else>
            <a>删除</a>
          </div>
        </template>
      </template>
    </a-table>
  </div>
</template>

<script setup lang="ts">
import { reactive, onMounted } from "vue";
import { useRouter } from "vue-router";
import Header from "../../Header.vue";
import { apiGetPipelineInfo } from "@/apis/pipeline";

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

const getPipelineInfo = async () => {
  console.log(router.currentRoute.value.params.id);
  const data = await apiGetPipelineInfo(":name", {});
  console.log("apiGetPipelineInfo", data);
  Object.assign(pipelineInfo, data.pipeline.jobs);
  pagination.pageSize = data.pagination.size;
  pagination.total = data.pagination.total;
};

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

onMounted(() => {
  getPipelineInfo();
});
</script>
