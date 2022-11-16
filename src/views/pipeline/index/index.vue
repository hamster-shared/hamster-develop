<template>
  <Header />
  <div class="w-[80%] mx-auto">
    <div class="flex justify-between mt-6 mb-2">
      <a-input-search
        v-model:value="searchValue"
        placeholder="input search text"
        style="width: 200px"
      />
      <a-button type="primary">
        <router-link to="/create">创建</router-link>
      </a-button>
    </div>

    <a-card v-for="(data, index) in dataList" :key="index">
      <div class="flex justify-between">
        <div>
          <router-link to="/pipeline/1">
            <div>{{ data.title }}</div>
          </router-link>
          <div>{{ data.description }}</div>
          <div>{{ data.status }}</div>
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

const searchValue = ref("");

const dataList = reactive([
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
    // handlHQSJ(current);
  },
  // showTotal: total => `总数：${total}人`, // 可以展示总数
});

onMounted(async () => {
  const data = await apiGetPipelines({ page: 1, size: 2 });
  console.log("apiGetPipelines", data);
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
