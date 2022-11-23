<template>
  <div class="mx-auto bg-white py-[32px] mx-[24px] rounded-xl">
    <div class="flex justify-between mb-6">
      <span class="text-2xl font-semibold text-[#121211]"
        >Hamster-pipeline
      </span>
      <div>
        <a-button class="mr-2">{{ $t("pipeline.stage.set") }}</a-button>
        <a-button type="primary">
          {{ $t("pipeline.stage.immediateImplementation") }}</a-button
        >
      </div>
    </div>

    <a-table
      :columns="columns"
      :data-source="pipelineInfo"
      v-bind:pagination="pagination"
    >
      <template #headerCell="{ column }">
        <template v-if="column.key === 'status'">
          <span> Status </span>
        </template>
      </template>

      <template #bodyCell="{ column, record }">
        <template v-if="column.key === 'status'">
          <span
            v-if="record.status == 0"
            @click="handleToNextPage(record.id)"
            class="cursor-pointer"
            >no data</span
          >
          <span
            v-if="record.status == 1"
            @click="handleToNextPage(record.id)"
            class="cursor-pointer"
            >running</span
          >
          <span
            v-if="record.status == 2"
            @click="handleToNextPage(record.id)"
            class="cursor-pointer"
            >passed</span
          >
          <span
            v-if="record.status == 3"
            @click="handleToNextPage(record.id)"
            class="cursor-pointer"
            >failed</span
          >
          <span
            v-if="record.status == 4"
            @click="handleToNextPage(record.id)"
            class="cursor-pointer"
            >stop</span
          >
        </template>
        <template v-else-if="column.key === 'stages'">
          <div v-for="(item, index) in record.stages" :key="index" class="flex">
            <div>
              <div v-if="item?.status == 0" class="inline-block">
                <img :src="nodataSVG" class="w-[20px] h-[20px]" />
              </div>
              <div v-if="item?.status == 1" class="inline-block">
                <img :src="runnngSVG" class="w-[20px] h-[20px]" />
              </div>
              <div v-if="item?.status == 2" class="inline-block">
                <img :src="successSVG" class="w-[20px] h-[20px]" />
              </div>
              <div v-if="item?.status == 3" class="inline-block">
                <img :src="failedSVG" class="w-[20px] h-[20px]" />
              </div>
              <div v-if="item?.status == 4" class="inline-block">
                <img :src="stopSVG" class="w-[20px] h-[20px]" />
              </div>
            </div>
            <div v-if="index !== record.stages.length - 1">
              <img :src="greyArrowSVG" />
            </div>
          </div>
        </template>
        <template v-else-if="column.key === 'duration'">
          <span class="block">
            {{ fromNowexecutionTime(record.startTime) }}
          </span>
          <span>{{ formatDurationTime(record.duration) }}</span>
        </template>
        <template v-else-if="column.key === 'action'">
          <div v-if="record.status == 1">
            <img :src="stopButtonSVG" class="mr-2 align-text-bottom" />
            <a
              @click="handleStop(record.id)"
              class="text-[#FF842C] hover:text-[#FF842C]"
              >{{ $t("pipeline.stage.stop") }}</a
            >
          </div>
          <div v-else>
            <img :src="deleteButtonSVG" class="mr-1 align-text-bottom" />
            <a
              @click="handleDelete(record.id)"
              class="text-[#F52222] hover:text-[#F52222]"
              >{{ $t("pipeline.stage.delete") }}
            </a>
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
import stopButtonSVG from "@/assets/icons/pipeline-stop-button.svg";
import deleteButtonSVG from "@/assets/icons/pipeline-delete-button.svg";
import greyArrowSVG from "@/assets/icons/grey-arrow.svg";
import { message } from "ant-design-vue";
import { formatDurationTime } from "@/utils/time/dateUtils.js";
import { fromNowexecutionTime } from "@/utils/time/dateUtils.js";

const router = useRouter();

const pathName = router.currentRoute.value.params.name;

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
    dataIndex: "triggerMode",
    key: "triggerMode",
  },
  {
    title: "Stage",
    key: "stages",
    dataIndex: "stages",
  },
  {
    title: "Time",
    key: "duration",
    dataIndex: "duration",
  },
  {
    title: "Action",
    key: "action",
  },
]);

const pipelineInfo = reactive([
  {
    key: "0",
    id: 0,
    status: 3,
    triggerMode: "Manual trigger",
    startTime: "2022-11-21T17:46:26.041337+08:00",
    duration: 1109625375,
    stages: [],
  },
]);

const pagination = reactive({
  // 分页配置器
  pageSize: 10, // 一页的数据限制
  current: 1, // 当前页
  total: 10, // 总数
  size: "small",
  hideOnSinglePage: false, // 只有一页时是否隐藏分页器
  showQuickJumper: false, // 是否可以快速跳转至某页
  showSizeChanger: false, // 是否可以改变 pageSize
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

const handleToNextPage = (id) => {
  router.push(`/pipeline/${pathName}/${id}`);
};

const getPipelineInfo = async () => {
  console.log(
    "router.currentRoute.value.params",
    router.currentRoute.value.params,
    pathName
  );
  const { data } = await apiGetPipelineInfo(pathName, {});
  console.log("apiGetPipelineInfo", data);
  Object.assign(pipelineInfo, data.data);
  pagination.pageSize = data.pageSize;
  pagination.total = data.total;
};

const handleDelete = async (id) => {
  try {
    await apiDeletePipelineInfo(pathName, id);
    console.log("id", id);
    const newJobs = pipelineInfo.filter((x) => x.id !== id);
    Object.assign(pipelineInfo, newJobs);
    message.success("This is a success message");
  } catch {
    message.error("This is an error message");
  }
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
  font-size: 12px;
  border-radius: 6px;
  width: 120px;
  height: 40px;
  &:hover,
  &:focus {
    color: #28c57c;
    border-color: #28c57c;
  }
}
.ant-btn-primary {
  background: #28c57c;
  &:hover,
  &:focus {
    border-color: #28c57c;
    background: #28c57c;
    color: white;
  }
}
:deep(.ant-table-thead > tr > th) {
  background: #121211;
  height: 48px;
  text-align: center;
  color: #ffffff;
  border-top: 12px;
}
:deep(.ant-table table) {
  text-align: center;
  border: 1px solid #f8f8f8;
  box-shadow: 3px 3px 12px rgba(203, 217, 207, 0.1);
  border-radius: 12px;
}
:deep(.ant-table-container table > thead > tr:first-child th:first-child) {
  border-top-left-radius: 12px;
}
:deep(.ant-table-container table > thead > tr:first-child th:last-child) {
  border-top-right-radius: 12px;
}

:deep(.ant-table-tbody > tr > td) {
  font-size: 12px;
  color: #7b7d7b;
}

ol,
ul,
dl {
  margin-bottom: 0px;
}
:deep(.ant-table-pagination-right) {
  justify-content: unset !important;
}
:deep(.ant-table-pagination) {
  display: block;
  text-align: center;
}
:deep(.ant-table-pagination.ant-pagination) {
  margin-top: 20px;
  margin-bottom: 0px;
}
:deep(.ant-pagination-item-active) {
  background: #28c57c;
  border-color: #28c57c;
  & a {
    color: white;
  }
}
:deep(.ant-pagination-item:hover a) {
  color: #28c57c !important;
}
:deep(.ant-pagination-prev:hover),
:deep(.ant-pagination-next:hover) {
  .ant-pagination-item-link {
    color: #28c57c;
  }
}
:deep(.ant-pagination-jump-prev),
:deep(.ant-pagination-jump-next) {
  .ant-pagination-item-container .ant-pagination-item-link-icon {
    color: #28c57c;
  }
}
:deep(.ant-pagination-item-active:hover a) {
  color: white !important;
}

:deep(.ant-table-tbody > tr > td:nth-child(4)) {
  display: flex;
  border-bottom: unset;
  justify-content: center;
  align-items: center;
  height: 70.7px;
}
</style>
