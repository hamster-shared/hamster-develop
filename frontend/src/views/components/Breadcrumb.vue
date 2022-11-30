<template>
  <div class="breadcrumb-list">
    <span @click="backRoute" class="path-info">
      <img src="@/assets/icons/arrow-left.svg" class="fill-current svg-img text-[#28c57c]" />
      <span>back</span>
    </span>
    <span> | </span>
    <span>{{ currentName }}</span>
  </div>
</template>
<script lang='ts' setup>
import { toRefs } from "vue";
import { useRouter } from "vue-router";

const props = defineProps({
  currentName: {
    type: String,
    default: '',
  },
});
const router = useRouter();
const { currentName } = toRefs(props);

const backRoute = () => {
  router.go(-1)
}
</script>
<style lang='less' scoped>
.breadcrumb-list {
  font-size: 24px;
  color: #121211;
  // margin: 0px auto 24px;
  margin-bottom: 24px;
  background-color: #F9F9F9;

  span {
    font-weight: 600;
  }


  .path-info {
    cursor: pointer;

    &:hover {
      color: #28c57c;

      .svg-img {
        color: #28c57c;
        // fill: currentColor;
      }

      // img {
      //   transform: translateX(-2px);
      //   filter: drop-shadow(#28c57c 2px 0);
      // }
    }
  }
}
</style>