<template>
  <div>
    测试
    <div class="scoll " ref="scrollOne">
      <div class="h-[1000px]">我要滚</div>
    </div>
    <div class="h-[500px]">我时后面要滚的</div>
  </div>
</template>
<script lang='ts' setup>
import { ref, onMounted, handleError } from "vue";
const scrollOne = ref();

const handleScroll = () => {
  console.log('scrollHeight===' + scrollOne.value.scrollHeight);
  console.log('scrollTop===' + scrollOne.value.scrollTop)
}

onMounted(() => {
  // window.addEventListener('scroll', function (e) {
  //   console.log(e.target, 'hhh')
  //   console.log('scrollHeight===' + scrollOne.value.scrollHeight);
  //   console.log('scrollTop===' + scrollOne.value.scrollTop)
  //   // console.log('scrollBottom===' + scrollOne.value.scrollBottom)
  //   console.log()

  //   // if (scrollOne.value.scrollTop >= 450) {
  //   //   document.body.scrollTop = 500 + 'px'
  //   // }
  // }, true)
})


</script>
<style lang='less' scoped>
.scoll {
  height: 500px;
  overflow-y: scroll;
  background-color: antiquewhite;
  // overscroll-behavior: auto;
}
</style>