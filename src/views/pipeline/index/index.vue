<template>
  <Header />
  <div class="w-[80%] mx-auto">
    <div class="flex justify-between mt-6 mb-2">
      <a-input-search
        v-model:value="value"
        placeholder="input search text"
        style="width: 200px"
      />
      <a-button type="primary">创建</a-button>
    </div>

    <a-card v-for="(data, index) in dataList" :key="index">
      <div class="flex justify-between">
        <div>
          <div>{{ data.title }}</div>
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
      show-less-items
      class="block float-right"
    />
  </div>
</template>

<script setup lang="ts">
import { reactive } from "vue";
import Header from "../../Header.vue";

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

// const handlHQSJ = async (page) => {
//   const {total, pipelines} = await HQSU(page)
//   ...
// }

const pagination = reactive({
  // 分页配置器
  pageSize: 10, // 一页的数据限制
  current: 1, // 当前页
  total: 10, // 总数
  hideOnSinglePage: false, // 只有一页时是否隐藏分页器
  showQuickJumper: true, // 是否可以快速跳转至某页
  showSizeChanger: true, // 是否可以改变 pageSize
  pageSizeOptions: ["10", "20", "30"], // 指定每页可以显示多少条
  onShowSizeChange: (current, pagesize) => {
    // 改变 pageSize时的回调
    pagination.current = current;
    pagination.pageSize = pagesize;
  },
  onChange: (current) => {
    // 切换分页时的回调，
    pagination.current = current;
    // handlHQSJ(current);
  },
  // showTotal: total => `总数：${total}人`, // 可以展示总数
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
