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

    <a-table :columns="columns" :data-source="data" :pagination="pagination">
      <template #headerCell="{ column }">
        <template v-if="column.key === 'status'">
          <span> Status </span>
        </template>
      </template>

      <template #bodyCell="{ column, record }">
        <template v-if="column.key === 'status'">
          {{ record.status }}
        </template>
        <template v-else-if="column.key === 'stage'">
          <span>
            <a-tag
              v-for="tag in record.tags"
              :key="tag"
              :color="
                tag === 'loser'
                  ? 'volcano'
                  : tag.length > 5
                  ? 'geekblue'
                  : 'green'
              "
            >
              {{ tag.toUpperCase() }}
            </a-tag>
          </span>
        </template>
        <template v-else-if="column.key === 'action'">
          <span>
            <a>Delete</a>
          </span>
        </template>
      </template>
    </a-table>
  </div>
</template>

<script setup lang="ts">
import { reactive } from "vue";
import Header from "../../Header.vue";

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
    dataIndex: "info",
    key: "info",
  },
  {
    title: "Stage",
    key: "stage",
    dataIndex: "stage",
  },
  {
    title: "Time",
    key: "time",
    dataIndex: "time",
  },
  {
    title: "Action",
    key: "action",
  },
]);

const data = reactive([
  {
    key: "1",
    status: "running",
    id: "#1",
    info: "Linda手动触发master|b4a0d99",
    tags: ["nice", "developer"],
    time: "1小时前执行",
  },
  {
    key: "2",
    status: "passed",
    id: "#2",
    info: "Linda手动触发master|b4a0d99",
    tags: ["loser"],
    time: "1小时前执行,耗时6分钟",
  },
  {
    key: "3",
    status: "failed",
    id: "#3",
    info: "Linda手动触发master|b4a0d99",
    tags: ["cool", "teacher"],
    time: "1小时前执行",
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
    // handlHQSJ(current);
  },
  // showTotal: total => `总数：${total}人`, // 可以展示总数
});
</script>
