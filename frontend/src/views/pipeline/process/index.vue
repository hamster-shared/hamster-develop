<template>
  <div class="process">
    <div class="bg-[#121211] rounded-t-[12px] h-[92px] p-[24px] text-center">
      <a-row>
        <a-col :span="6">
          <div class="process-detail-item">
            <div class="process-detail-title">{{ jobData.name }}</div>
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
            <div class="process-detail-info">{{ fromNowexecutionTime(jobData.startTime, 'noThing') }}</div>
          </div>
        </a-col>
        <a-col :span="6">
          <div>
            <div class="process-detail-title">
              {{ $t("log.totalDuration") }}
            </div>
            <div class="process-detail-info">
              {{ formatDurationTime(jobData.duration, 'noThing') }}
            </div>
          </div>
        </a-col>
      </a-row>
    </div>
    <div class="p-[24px] border border-solid border-[#EFEFEF] rounded-b-[12px]">
      <div class="process-content">
        <div class="flex justify-between">
          <span class="process-content-title">{{
              $t("log.executionProcess")
          }}</span>
          <span class="text-[14px] text-[#28C57C] cursor-pointer" @click="checkAllLogs">{{ $t("log.viewAllLogs")
          }}</span>
        </div>
        <div class="process-scroll-box wrapper" ref="wrapper">
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
              <div class="inline-block border border-solid border-[#EFEFEF] p-[12px] rounded-[5px]"
                :class="item.status === 0 ? '' : 'cursorP'" @click="checkProcess(item, $event)">
                <img :src="getImageUrl(item.status)" class="w-[28px] mr-[24px] align-middle" v-if="item.status !== 1" />
                <img src="@/assets/images/run.gif" class="w-[28px]  mr-[24px] align-middle" v-else />
                <span class="align-middle">
                  <span class="text-[16px] text-[#121211] font-semibold mr-[24px]">{{ item.name }}</span>
                  <span class="text-[16px] text-[#7B7D7B]" v-if="item.status !== 0">{{
                      formatDurationTime(item.duration, 'noThing')
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
          <div v-for="it in jobData.actionResult.artifactorys" :key="it.id">
            {{ it.url }}
          </div>
          <a-empty v-if="jobData.actionResult.artifactorys.length <= 0" />
        </div>
      </div>
      <div class="process-content">
        <div class="process-content-title">{{ $t("log.report") }}</div>
        <div class="text-[#7B7D7B]">
          <div v-for="it in jobData.actionResult.reports" :key="it.id">{{ it.url }}</div>
          <a-empty v-if="jobData.actionResult.reports.length <= 0" />
        </div>
      </div>
    </div>
  </div>

  <ProcessModal ref="processModalRef" :text="title" :content="content" />
</template>
<script lang="ts" setup>
import { ref, onMounted, reactive, onUnmounted } from "vue";
import { useRouter } from "vue-router";
import { apiGetJobStageLogs } from "@/apis/jobs";
import { apiGetPipelineDetail } from "@/apis/pipeline";
import { fromNowexecutionTime, formatDurationTime } from '@/utils/time/dateUtils.js';
import BScroll from "@better-scroll/core";
import ProcessModal from "./components/ProcessModal.vue";
const router = useRouter();
const processModalRef = ref();
const title = ref("");
const content = ref([]);
const state = reactive({
  detailTimer: null,
  stagesTimer: null,
  running: true,
})
const wrapper = ref();
const jobData = reactive({
  id: undefined,
  stages: [],
  actionResult: {
    artifactorys: [],
    reports: [],
  }
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
  const { data } = await apiGetPipelineDetail(queryJson);
  Object.assign(jobData, data);
  state.running = data.stages.some((val: any) => val.status === 1)
  if (state.running) {
    state.detailTimer = setTimeout(() => {
      getPipelineDetail()
    }, 3000)
  } else {
    clearTimeout(state.detailTimer)
  }
}
const getImageUrl = (status: any) => {
  return new URL(`../../../assets/icons/Status${status}.svg`, import.meta.url)
    .href;
};

const checkAllLogs = () => {
  window.open(`#/allLogs/${queryJson.name}/${queryJson.id}`);
};

const checkProcess = async (item: any, e: Event) => {
  if (item.status === 0) {
    e.stopPropagation()
  } else {
    title.value = item.name;
    await getStageLogsData(item);
    processModalRef.value.showVisible();
  }

};

const getStageLogsData = async (item: any, start = 0, lastLine = 0) => {
  const query = Object.assign(queryJson, { stagename: item.name, start: start, lastLine: lastLine });
  const { data } = await apiGetJobStageLogs(query);
  if (data.end) {
    content.value = data?.content?.split('\r');
  } else {
    let t = data?.content?.split('\r')
    content.value.push(t)
  }

  if (!data.end) {
    state.stagesTimer = setTimeout(() => {
      getStageLogsData(item, 0, data.lastLine)
    }, 3000)
  } else {
    clearTimeout(state.stagesTimer)
  }
};


onMounted(async () => {
  await getPipelineDetail()
  let scroll = new BScroll(wrapper.value, {
    startX: 0,
    scrollX: true,
  });
});

onUnmounted(() => {
  clearTimeout(state.detailTimer)
  clearTimeout(state.stagesTimer)
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

      .cursorP {
        cursor: pointer;
      }
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
