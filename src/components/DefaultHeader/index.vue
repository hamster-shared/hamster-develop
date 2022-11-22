<template>
  <div class="default-header h-[64px] bg-[#ccc] flex justify-between">
    <div>
      <img src="@/assets/images/logo.jpg" class="h-[28px] mt-[18px] mb-[18px] cursor-pointer" @click="goHome" />
    </div>
    <div>
      <img src="@/assets/icons/fill.svg" class="mt-[16px] mb-[16px] mr-[32px]" />
      <img src="@/assets/icons/git.svg" class="mt-[16px] mb-[16px]" />
    </div>
  </div>
</template>
<script lang='ts' setup>
import { useRouter } from 'vue-router';
const router = useRouter()
const goHome = () => {
  router.push('/pipeline')
}
</script>
<style lang='less' scoped>
.default-header {
  padding: 0 32px;
  width: 100%;
  height: 64px;
  background-color: #ffffff;
  border: 1px solid #F8F8F8;
}
</style>