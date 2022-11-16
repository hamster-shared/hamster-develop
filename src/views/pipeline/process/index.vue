<template>
  <div class="process">
    <div class="bg-[#121211] rounded-t-[12px] h-[92px] p-[24px] text-center">
      <a-row>
        <a-col :span="6">
          <div class="process-detail-item">
            <div class="process-detail-title">Linda手动触发</div>
            <div class="process-detail-info">master | b4a0d99</div>
          </div>
        </a-col>
        <a-col :span="6">
          <div class="process-detail-item">
            <div class="process-detail-title">status</div>
            <div class="process-detail-info">Passed</div>
          </div>
        </a-col>
        <a-col :span="6">
          <div class="process-detail-item">
            <div class="process-detail-title">Execution Time</div>
            <div class="process-detail-info">7m 57s</div>
          </div>
        </a-col>
        <a-col :span="6">
          <div>
            <div class="process-detail-title">Total duration</div>
            <div class="process-detail-info">7m 57s</div>
          </div>
        </a-col>
      </a-row>
    </div>
    <div class="p-[24px] border border-solid border-[#EFEFEF] rounded-b-[12px]">
      <div class="process-content">
        <div class="flex justify-between">
          <span class="process-content-title">Execution
            Process</span>
          <span class="text-[14px] text-[#28C57C] cursor-pointer" @click="checkAllLogs">查看完整日志</span>
        </div>
        <div class="process-scroll-box wrapper" ref="wrapper">
          <div class="process-scroll content">
            <div v-for="item in data.execution_process" :key="item.id" class="inline-block execution_process_item ">
              <div class="inline-block border border-solid border-[#28C57C] p-[12px] cursor-pointer rounded-[5px]"
                @click="checkProcess(item)">
                <img src="@/assets/images/pro-1.jpg" class="w-[28px] mr-[24px] align-middle" />
                <span class="align-middle">
                  <span class="text-[16px] text-[#28C57C] font-semibold mr-[24px]">{{ item.name }}</span>
                  <span class="text-[16px] text-[#28C57C]">{{ item.time_consuming }}</span>
                </span>
              </div>
              <img src="@/assets/images/arrow-green.jpg" class="w-[28px] space-mark ml-[20px] mr-[20px]" />
            </div>
          </div>

        </div>

      </div>
      <div class="process-content">
        <div class="process-content-title">Artifats</div>
        <div class="text-[#7B7D7B]">{{ data.artifacts[0] }}</div>
      </div>
      <div class="process-content">
        <div class="process-content-title">Report</div>
        <div class="text-[#7B7D7B]">{{ data.report[0] }}</div>
      </div>
    </div>

  </div>
  <ProcessModal ref="processModalRef" :title="title" />
</template>
<script lang="ts" setup>
import { ref, onMounted } from "vue";
import BScroll from '@better-scroll/core'
import ProcessModal from "./components/ProcessModal.vue";
const processModalRef = ref()
const title = ref('')
const wrapper = ref()
const data = ref({
  id: 1,
  status: '啥啥啥',
  trigger_info: '@trigger_info',
  execution_process: [
    { status: 8888888, name: '开始', time_consuming: '5min' },
    { status: 8888888, name: '检出', time_consuming: '5min' },
    { status: 8888888, name: '合约检查', time_consuming: '5min' },
    { status: 8888888, name: '部署合约', time_consuming: '5min' },
    { status: 8888888, name: '部署合约', time_consuming: '5min' },
    { status: 8888888, name: '部署合约', time_consuming: '5min' },
    { status: 8888888, name: '部署合约', time_consuming: '5min' },
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

onMounted(() => {
  let scroll = new BScroll(wrapper.value, {
    startX: 0,
    scrollX: true,
  })
})

</script>
<style lang="less" scoped>
.process {
  width: 100%;
  font-size: 14px;

  .process-detail-item {
    position: relative;

    &::before {
      content: "";
      position: absolute;
      top: 0px;
      right: -2px;
      width: 1px;
      height: 44px;
      border: 1px dashed #3F4641;
    }
  }

  .process-detail-title {
    color: #ffffff;
    font-weight: 600;
  }

  .process-detail-info {
    color: #BCBEBC;
  }

  .process-scroll-box {

    // width: 100%;
    // overflow-x: scroll;
    white-space: nowrap;
    overflow: hidden;

    .process-scroll {
      // max-width: 100%;
      display: inline-block;
    }
  }

  .process-content {
    padding: 24px;
    border: 1px solid #EFEFEF;
    border-radius: 12px;
    margin-bottom: 24px;

    .process-content-title {
      font-size: 20px;
      font-weight: 600;
      color: #121211;
      margin-bottom: 12px;
    }

    .process-content-title:first {
      margin-bottom: 24px;
    }
  }

  .process-content:last-child {
    margin-bottom: 0px;
  }

  .execution_process_item:last-child {
    .space-mark {
      display: none;
    }
  }


}
</style>

