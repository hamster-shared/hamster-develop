<template>
  <div class="process">
    <div class="bg-[#121211] rounded-t-[12px] h-[92px] p-[24px] text-center">
      <a-row>
        <a-col :span="6">
          <div class="process-detail-item">
            <div class="process-detail-title">{{ jobData.name }}</div>
            <!-- <div class="process-detail-info">master | b4a0d99</div> -->
            <div class="process-detail-info">{{ jobData.triggerMode }}</div>
          </div>
        </a-col>
        <a-col :span="6">
          <div class="process-detail-item">
            <div class="process-detail-title">{{ $t("log.status") }}</div>
            <div class="process-detail-info">
              {{ StatusEnum[jobData.status] }}
            </div>
          </div>
        </a-col>
        <a-col :span="6">
          <div class="process-detail-item">
            <div class="process-detail-title">{{ $t("log.executionTime") }}</div>
            <div class="process-detail-info">{{ jobData.startTimeText + $t("log.ago") }}</div>
          </div>
        </a-col>
        <a-col :span="6">
          <div>
            <div class="process-detail-title">{{ $t("log.totalDuration") }}</div>
            <div class="process-detail-info">
              {{ formatDuring(jobData.duration) }}
            </div>
          </div>
        </a-col>
      </a-row>
    </div>
    <div class="p-[24px] border border-solid border-[#EFEFEF] rounded-b-[12px]">
      <div class="process-content">
        <div class="flex justify-between">
          <span class="process-content-title">{{ $t("log.executionProcess") }}</span>
          <span class="text-[14px] text-[#28C57C] cursor-pointer" @click="checkAllLogs">{{ $t("log.viewAllLogs")
          }}</span>
        </div>
        <div class="process-scroll-box wrapper" ref="wrapper">
          <!-- <a-button @click="checkProcess({ name: 'hh' })">modal</a-button> -->
          <div class="process-scroll content">
            <div class="inline-block execution_process_item">
              <div class="inline-block border border-solid border-[#EFEFEF] p-[12px] rounded-[5px]">
                <img src="@/assets/icons/Frame.svg" class="w-[28px] mr-[24px] align-middle" />
                <span class="align-middle">
                  <span class="text-[16px] text-[#121211] font-semibold mr-[24px]">{{ $t("log.start") }}</span>
                </span>
              </div>
              <img src="@/assets/images/arrow-green.jpg" class="w-[28px] space-mark ml-[20px] mr-[20px]" />
            </div>
            <div v-for="item in jobData.stages" :key="item.name" class="inline-block execution_process_item">
              <div class="inline-block border border-solid border-[#EFEFEF] p-[12px] cursor-pointer rounded-[5px]"
                @click="checkProcess(item)">
                <!-- <img src="@/assets/icons/Status0.svg" class="w-[28px] mr-[24px] align-middle" /> -->
                <img :src="getImageUrl(item.status)" class="w-[28px] mr-[24px] align-middle" />
                <span class="align-middle">
                  <span class="text-[16px] text-[#121211] font-semibold mr-[24px]">{{ item.name }}</span>
                  <span class="text-[16px] text-[#7B7D7B]">{{
                      formatDuring(item.duration)
                  }}</span>
                </span>
              </div>
              <img src="@/assets/images/arrow-green.jpg" class="w-[28px] space-mark ml-[20px] mr-[20px]" />
            </div>
          </div>
        </div>
      </div>
      <div class="process-content">
        <div class="process-content-title">{{ $t("log.artifats") }}</div>
        <div class="text-[#7B7D7B]">
          <div v-for="it in jobData.artifactorys" :key="it.id">
            {{ it.url }}
          </div>
          <a-empty v-if="jobData.artifactorys.length <= 0" />
        </div>
      </div>
      <div class="process-content">
        <div class="process-content-title">{{ $t("log.report") }}</div>
        <div class="text-[#7B7D7B]">
          <div v-for="it in jobData.reports" :key="it.id">{{ it.url }}</div>
          <a-empty v-if="jobData.reports.length <= 0" />
        </div>
      </div>
    </div>
  </div>

  <ProcessModal ref="processModalRef" :text="title" :content="content" />
</template>
<script lang="ts" setup>
import { ref, onMounted, reactive } from "vue";
import { useRouter } from "vue-router";
import { apiGetJobStageLogs } from "@/apis/jobs";
import { apiGetPipelineDetail } from "@/apis/pipeline";
import { formatDuring } from "@/common/common";
import dayjs from "dayjs";
import BScroll from "@better-scroll/core";
import ProcessModal from "./components/ProcessModal.vue";
const router = useRouter();
const processModalRef = ref();
const title = ref("");
const content = ref("");
const wrapper = ref();
const jobData = reactive({
  id: undefined,
  stages: [],
  artifactorys: [],
  reports: [],
});

const enum StatusEnum {
  "Non-execution",
  "Running",
  "Failed",
  "Passed",
  "Stop",
}

const queryJson = reactive({
  name: router.currentRoute.value.params.name,
  id: router.currentRoute.value.params.id,
});

const getPipelineDetail = async () => {
  const data = await apiGetPipelineDetail(queryJson);
  Object.assign(jobData, data.data);
  // 计算开始时间
  getStarTime(jobData);
};

const getImageUrl = (status) => {
  return new URL(`../../../assets/icons/Status${status}.svg`, import.meta.url).href;
};

const getStarTime = (data: any) => {
  // 当前时间
  let date = dayjs(new Date()).valueOf();
  // 开始时间
  let startTime = dayjs(data.startTime).valueOf();
  let t = date - startTime;
  data.startTimeText = formatDuring(t);
};

const checkAllLogs = () => {
  window.open(`/allLogs/${queryJson.id}/${queryJson.name}`);
};

const checkProcess = (item) => {
  title.value = item.name;
  processModalRef.value.showVisible();
  getStageLogsData(item);
};

const getStageLogsData = async (item) => {
  const query = Object.assign(queryJson, { stagename: item.name });
  const data = await apiGetJobStageLogs(query);
  content.value = data.data?.content || '我暂时还没有值';
  // console.log(data.data, '9999')
};

onMounted(() => {
  getPipelineDetail();

  let scroll = new BScroll(wrapper.value, {
    startX: 0,
    scrollX: true,
  });
});
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
      border: 1px dashed #3f4641;
    }
  }

  .process-detail-title {
    color: #ffffff;
    font-weight: 600;
  }

  .process-detail-info {
    color: #bcbebc;
  }

  .process-scroll-box {
    white-space: nowrap;
    overflow: hidden;

    .process-scroll {
      display: inline-block;
    }
  }

  .process-content {
    padding: 24px;
    border: 1px solid #efefef;
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
