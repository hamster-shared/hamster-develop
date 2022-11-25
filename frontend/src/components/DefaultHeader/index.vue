<template>
  <div class="default-header h-[64px] bg-[#ccc] flex justify-between">
    <div>
      <img
        src="@/assets/images/logo.png"
        class="h-[28px] mt-[18px] mb-[18px] cursor-pointer"
        @click="goHome"
      />
    </div>
    <div>
      <div class=" cursor-pointer flex h-[36px]">
        <div @click="changeTheme('white')" class="w-[44px] border border-solid border-[#E2B578] rounded-tl-[6px] rounded-bl-[6px]">white</div>
        <div @click="changeTheme('dark')" class="w-[44px] border border-solid border-[#E2B578] rounded-tr-[6px] rounded-br-[6px]">dark</div>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { useRouter } from "vue-router";
import { onMounted } from "vue";

const router = useRouter();
const goHome = () => {
  router.push("/pipeline");
};

const changeTheme = (val) => {
  if (val === 'white') {
    document.documentElement.classList.remove('dark')
  } else {
    document.documentElement.classList.add('dark')
  }
}

onMounted(() => {
  changeTheme('dark');
});
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
