<template>
  <div class="w-[80%] mx-auto">
    <div class="flex justify-between mb-4">
      <a-input
        v-model:value="searchValue"
        placeholder="search here..."
        style="width: 340px"
        class="w-[340px] h-[48px]"
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
            class="mb-3 text-3xl font-semibold cursor-pointer"
          >
            {{ data.name }}
          </div>
          <div class="mb-3 text-xl font-normal">{{ data.description }}</div>
          <div>
            <div v-if="data.status == 0" class="text-xl font-normal leading-6">
              暂无执行数据
            </div>
            <div v-if="data.status == 1" class="text-xl font-normal leading-6">
              <img :src="runnngSVG" />
              执行中
            </div>
            <div v-if="data.status == 2" class="text-xl font-normal leading-6">
              <img :src="successSVG" />
              执行成功
            </div>
            <div v-if="data.status == 3" class="text-xl font-normal leading-6">
              <img :src="failedSVG" />
              推送到制品库失败
            </div>
            <div v-if="data.status == 4" class="text-xl font-normal leading-6">
              <img :src="stopSVG" />
              用户终止
            </div>
          </div>
        </div>
        <div class="self-center text-xl font-normal">
          {{ data.last_execution_time }}
        </div>
        <div class="self-center">
          <a-button type="primary">立即执行</a-button>
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
  total: 10, // 总数
  hideOnSinglePage: false, // 只有一页时是否隐藏分页器
  showQuickJumper: true, // 是否可以快速跳转至某页
  showSizeChanger: true,
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
  border: 2px solid #6481dc;
  border-radius: 6px;
  &:not(.ant-input-affix-wrapper-disabled):hover {
    border-color: #6481dc;
    border-right-width: 2px !important;
  }
}
.ant-card-bordered {
  margin-bottom: 20px;
}
.ant-btn {
  display: block;
  width: 200px;
  height: 48px;
}
.ant-btn-primary {
  margin-bottom: 10px;
  width: 200px;
  height: 48px;
  background: #6481dc;
  border-radius: 3px;
}
.ant-btn-primary:hover,
.ant-btn-primary:focus {
  border-color: #6481dc;
  background: #6481dc;
}
.ant-card-bordered {
  border: 1px solid #dedddc;
}
:deep(.ant-pagination-options-quick-jumper input) {
  margin: 0px 0px 0px 8px !important;
}
:deep(.ant-btn) {
  font-size: 16px;
}
.ant-pagination {
  padding-bottom: 24px !important;
}
</style>
