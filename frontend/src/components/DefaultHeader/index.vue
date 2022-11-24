<template>
  <div class="default-header h-[64px] bg-[#ccc] flex justify-between">
    <div>
      <img
        src="@/assets/images/logo.jpg"
        class="h-[28px] mt-[18px] mb-[18px] cursor-pointer"
        @click="goHome"
      />
    </div>
    <div>
      <a-button @click="changeLang" class="align-middle mr-[32px] normal-button"
        >中/En</a-button
      >
      <img
        src="@/assets/icons/fill.svg"
        class="mt-[16px] mb-[16px] mr-[32px]"
      />
      <img src="@/assets/icons/git.svg" class="mt-[16px] mb-[16px]" />
    </div>
  </div>
</template>

<script lang="ts" setup>
import { useRouter } from "vue-router";

const router = useRouter();
const goHome = () => {
  router.push("/pipeline");
};

const changeLang = () => {
  const language = window.localStorage.getItem("language");

  if (language == undefined || language == "zh") {
    window.localStorage.setItem("language", "en");
    location.reload();
  } else {
    window.localStorage.setItem("language", "zh");
    location.reload();
  }
};
</script>

<style lang="less" scoped>
.default-header {
  position: fixed;
  top: 0;
  padding: 0 32px;
  width: 100%;
  height: 64px;
  background-color: #ffffff;
  border: 1px solid #f8f8f8;
  z-index: 1;
}

.ant-btn {
  width: 60px;
  height: 32px;
  border-radius: 6px;
  font-size: 12px;

  &:hover,
  &:focus {
    color: #28c57c;
    border-color: #28c57c;
  }
}
.normal-button {
  color: #28c57c;
  border-color: #28c57c;
}
</style>
