<template>
  <div class="process">
    <div>
      <a-row>
        <a-col :span="6">
          <div>Linda手动触发</div>
          <div>master | b4a0d99</div>
        </a-col>
        <a-col :span="6">
          <div>status</div>
          <div>Passed</div>
        </a-col>
        <a-col :span="6">
          <div>Execution Time</div>
          <div>7m 57s</div>
        </a-col>
        <a-col :span="6">
          <div>Total duration</div>
          <div>7m 57s</div>
        </a-col>
      </a-row>
    </div>
    <div>
      <div class="flex justify-between">
        <span>Execution Process</span>
        <span class="text-red-300 cursor-pointer" @click="checkAllLogs">查看完整日志</span>
      </div>
      <div>
        <div v-for="item in data.execution_process" :key="item.id" class="inline-block execution_process_item ">
          <div class="inline-block border border-solid border-[#ccc] p-[24px] cursor-pointer"
            @click="checkProcess(item)">
            <img src="@/assets/logo.svg" class="w-[30px]" />
            <span>
              <span>{{ item.name }}</span>
              <span>{{ item.time_consuming }}</span>
            </span>
          </div>
          <div class="inline-block space-mark">-></div>
        </div>
      </div>

    </div>
    <div>
      <div>Artifats</div>
      <div>{{ data.artifacts[0] }}</div>
    </div>
    <div>
      <div>Report</div>
      <div>{{ data.report[0] }}</div>
    </div>
  </div>
  <ProcessModal ref="processModalRef" :title="title" />
</template>
<script lang="ts" setup>
import { ref } from "vue";
import ProcessModal from "./components/ProcessModal.vue";
const processModalRef = ref()
const title = ref('')
const data = ref({
  id: 1,
  status: '啥啥啥',
  trigger_info: '@trigger_info',
  execution_process: [
    { status: 8888888, name: '开始', time_consuming: '5min' },
    { status: 8888888, name: '检出', time_consuming: '5min' },
    { status: 8888888, name: '合约检查', time_consuming: '5min' },
    { status: 8888888, name: '部署合约', time_consuming: '5min' }
  ],
  last_execution_time: '1小时前',
  time_consuming: '耗时3分钟',
  artifacts: ['龙年员阶接自率些又斯听低便从统置题属思值出因前开为几意采'],
  report: ['装工边克思办须确太']
});
// console.log(data);

const checkAllLogs = () => {
  window.open('/allLogs')
}

const checkProcess = (item) => {
  title.value = item.name
  processModalRef.value.showVisible()
}
</script>
<style lang="less" scoped>
.process {
  width: 100%;
  padding: 24px;

  .execution_process_item:last-child {
    .space-mark {
      display: none;
    }
  }


}
</style>

