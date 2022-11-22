<template>
  <div class="w-[80%] mx-auto bg-white py-[32px] px-[24px] rounded-xl">
    <div class="flex justify-between mb-4">
      <a-input
        v-model:value="searchValue"
        placeholder="search here..."
        style="width: 370px"
        class="w-[340px] h-[40px]"
      >
        <template #prefix>
          <img :src="searchSVG" />
        </template>
      </a-input>
      <a-button type="primary">
        <router-link to="/create">创建pipeline</router-link>
      </a-button>
    </div>

    <a-card v-for="(data, index) in PipelineList" :key="index">
      <div class="flex justify-between">
        <div class="self-center">
          <div
            @click="$router.push(`/pipeline/${data.id}`)"
            class="mb-3 text-xl font-semibold cursor-pointer text-[#121211]"
          >
            {{ data.name }}
          </div>
          <div class="mb-3 text-sm font-normal text-[#7B7D7B]">
            {{ data.description }}
          </div>
          <div>
            <div
              v-if="data.status == 0"
              class="text-sm font-normal text-[#7B7D7B]"
            >
              暂无执行数据
            </div>
            <div
              v-if="data.status == 1"
              class="text-sm font-normal text-[#2C5AFF]"
            >
              <img :src="runnngSVG" />
              执行中
            </div>
            <div
              v-if="data.status == 2"
              class="text-sm font-normal text-[#2DCE83]"
            >
              <img :src="successSVG" />
              执行成功
            </div>
            <div
              v-if="data.status == 3"
              class="text-sm font-normal text-[#F52222]"
            >
              <img :src="failedSVG" />
              推送到制品库失败
            </div>
            <div
              v-if="data.status == 4"
              class="text-sm font-normal text-[#FF842C]"
            >
              <img :src="stopSVG" />
              用户终止
            </div>
          </div>
        </div>
        <div class="self-center text-center">
          <span
            class="text-sm font-normal bg-[#F8F8F8] py-2 px-3 rounded text-[#3F4641] block mb-3"
            >{{ data.last_execution_time }}</span
          >
          <span class="text-xs">
            <img :src="wasteTimeSVG" />
            {{ data.time_consuming }}
          </span>
        </div>
        <div class="self-center">
          <a-button type="primary" v-if="data.status !== 1">立即执行</a-button>
          <a-button type="primary" danger v-if="data.status === 1"
            >终止执行</a-button
          >
          <a-button>设置</a-button>
        </div>
      </div>
    </a-card>

    <a-pagination class="block float-right" v-bind="pagination" />
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from "vue";
import { apiGetPipelines } from "@/apis/pipeline";
import searchSVG from "@/assets/icons/search.svg";
import runnngSVG from "@/assets/icons/pipeline-running.svg";
import successSVG from "@/assets/icons/pipeline-success.svg";
import failedSVG from "@/assets/icons/pipeline-failed.svg";
import stopSVG from "@/assets/icons/pipeline-stop.svg";
import wasteTimeSVG from "@/assets/icons/pipeline-waste-time.svg";

const searchValue = ref("");

const PipelineList = reactive([
  {
    title: "Hamster-pipeline-1",
    description: "11111",
    status: "success",
  },
  {
    title: "Hamster-pipeline-2",
    description: "222222",
    status: "fail",
  },
  {
    title: "Hamster-pipeline-3",
    description: "333333",
    status: "success",
  },
  {
    title: "Hamster-pipeline-4",
    description: "44444",
    status: "fail",
  },
]);

const getPipelineInfo = async () => {
  const data = await apiGetPipelines({ page: 1, size: 2 });
  console.log("apiGetPipelines", data);
  Object.assign(PipelineList, data.pipelines);
  pagination.pageSize = data.pagination.size;
  pagination.total = data.pagination.total;
};

const pagination = reactive({
  // 分页配置器
  pageSize: 10, // 一页的数据限制
  current: 1, // 当前页
  total: 10, // 总数,
  size: "small",
  hideOnSinglePage: false, // 只有一页时是否隐藏分页器
  showQuickJumper: false, // 是否可以快速跳转至某页
  showSizeChanger: false,
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

<style scoped lang="less">
.ant-input-affix-wrapper {
  border: 1px solid #efefef;
  border-radius: 6px;
  &:not(.ant-input-affix-wrapper-disabled):hover {
    border-color: #6481dc;
  }
}
.ant-card-bordered {
  margin-bottom: 20px;
  box-shadow: 3px 3px 12px rgba(203, 217, 207, 0.1);
  border-radius: 12px;
  border: 1px solid #efefef;
}
.ant-btn {
  display: block;
  width: 120px;
  height: 40px;
  border-radius: 6px;
  font-size: 12px;
  &:hover,
  &:focus {
    color: #28c57c;
    border-color: #28c57c;
  }
}
.ant-btn-primary {
  margin-bottom: 10px;
  border-radius: 6px;
  width: 120px;
  height: 40px;
  background: #28c57c;
  &:hover,
  &:focus {
    border-color: #28c57c;
    background: #28c57c;
    color: white;
  }
}
.ant-btn-dangerous.ant-btn-primary {
  border-color: #ff842c;
  background: #ff842c;
  &:hover,
  &:focus {
    border-color: #ff842c;
    background: #ff842c;
  }
}

.ant-card-bordered {
  border: 1px solid #dedddc;
}
ol,
ul,
dl {
  margin-bottom: 0px;
}
.float-right {
  float: unset;
}
.ant-pagination {
  text-align: center;
}
:deep(.ant-pagination-item-active) {
  background: #28c57c;
  border-color: #28c57c;
  & a {
    color: white;
  }
}
:deep(.ant-pagination-item:hover a) {
  color: #28c57c;
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
</style>
