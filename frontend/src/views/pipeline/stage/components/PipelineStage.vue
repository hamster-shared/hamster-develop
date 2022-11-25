<template>
  <!-- <template v-if="isAnimation">
    <div v-for="(item, index) in animationStages">
      <div v-if="index == animationIndex">DONGHUA</div>
      <div v-else>Icon</div>
      <div v-if="index == animationIndex">
        <img :src="greyArrowSVG" />
      </div>
    </div>
  </template> -->
  <template v-if="stages">
    <div v-for="(item, index) in stages" :key="index" class="flex">
      <div>
        <div v-if="item?.status == 0" class="inline-block">
          <img :src="nodataSVG" class="w-[20px] h-[20px]" />
        </div>
        <div v-if="item?.status == 1" class="inline-block">
          <img src="@/assets/images/run.gif" class="w-[20px] h-[20px]" />
        </div>
        <div v-if="item?.status == 3" class="inline-block">
          <img :src="successSVG" class="w-[20px] h-[20px]" />
        </div>
        <div v-if="item?.status == 2" class="inline-block">
          <img :src="failedSVG" class="w-[20px] h-[20px]" />
        </div>
        <div v-if="item?.status == 4" class="inline-block">
          <img :src="stopSVG" class="w-[20px] h-[20px]" />
        </div>
      </div>
      <div v-if="index !== stages.length - 1">
        <img :src="greyArrowSVG" />
      </div>
    </div>
  </template>
</template>

<script setup lang="ts">
import { ref, computed, watch, watchEffect, onMounted, toRefs } from "vue";
import successSVG from "@/assets/icons/pipeline-success.svg";
import failedSVG from "@/assets/icons/pipeline-failed.svg";
import stopSVG from "@/assets/icons/pipeline-stop.svg";
import nodataSVG from "@/assets/icons/pipeline-no-data.svg";
import greyArrowSVG from "@/assets/icons/grey-arrow.svg";

const props = defineProps<{
  stages: [];
}>();

const { stages } = toRefs(props);

// const animationIndex = ref(0);
// const animationStages = computed(() => {
//   return props.stages.filter((stage, index) => index <= animationIndex.value);
// });

// const isAnimation = computed(() => {
//   return props.stages.some((x) => x.status == 1);
// });

// watch(
//   () => props.stages,
//   () => {
//     animationIndex.value = 0;
//   }
// );

// watchEffect((onInvalidate) => {
//   const timer = setInterval(() => {
//     isAnimation.value &&
//       animationIndex.value < props.stages.length - 1 &&
//       animationIndex.value++;
//   }, 1000);
//   onInvalidate(() => clearInterval(timer));
// });

onMounted(() => {
  if (stages.value.length > 8) {
    stages.value = [...stages.value.slice(-8, -1), ...stages.value.slice(-1)];
  }
  console.log("stages.value:", stages.value);
});
</script>
