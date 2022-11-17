<template>
  <Header />
  <div class="w-[80%] mx-auto">
    <div class="flex justify-between mt-6 mb-2">
      <a-input
        v-model:value="searchValue"
        placeholder="search here..."
        style="width: 200px"
      >
        <template #prefix>
          <img :src="searchSVG" />
        </template>
      </a-input>
      <a-button type="primary">
        <router-link to="/create">创建</router-link>
      </a-button>
    </div>

    <a-card v-for="(data, index) in PipelineList" :key="index">
      <div class="flex justify-between">
        <div>
          <div @click="$router.push(`/pipeline/${data.id}`)">
            {{ data.name }}
          </div>
          <div>{{ data.description }}</div>
          <div>
            <div v-if="data.status == 0">暂无执行数据</div>
            <div v-if="data.status == 1">执行中</div>
            <div v-if="data.status == 2">成功</div>
            <div v-if="data.status == 3">失败</div>
            <div v-if="data.status == 4">用户终止</div>
          </div>
        </div>
        <div>
          {{ data.last_execution_time }}
        </div>
        <div>
          <a-button type="primary">立即执行</a-button>
          <a-button>设置</a-button>
        </div>
      </div>
    </a-card>

    <a-pagination
      :pageSize="pagination.pageSize"
      v-model:current="pagination.current"
      :total="pagination.total"
      @change="pagination.onChange"
      class="block float-right"
      show-quick-jumper
    />
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from "vue";
import { apiGetPipelines } from "@/apis/pipeline";
import Header from "../../Header.vue";
import searchSVG from "@/assets/search.svg";

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
// .ant-card {
//   width: 100%;
//   margin: 20px auto 0;
//   :deep(.ant-card-body) {
//     width: 100%;
//     display: flex;
//     justify-content: space-between;
//     align-items: center;
//     padding: 10px 0;
//   }
// }
.ant-card-bordered {
  margin-bottom: 20px;
}
.ant-btn {
  display: block;
  width: 90px;
}
.ant-btn-primary {
  margin-bottom: 10px;
}
</style>
