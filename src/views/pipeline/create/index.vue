<template>
  <Header />
  <div class="bg-[#FFFFFF] border border-solid border-[#DEDDDC] mx-10 my-5 p-4 leading-[24px]">
    <div class="text-[20px] font-semibold">{{ $t("template.title") }}</div>
    <div class="flex justify-between">
      <div class="text-[#979797]">{{ $t("template.titleDesc") }}</div>
      <div class="text-[#6481DC]">{{ $t("template.helpDoc") }}</div>
    </div>
    <Tabs :active-key="activeKey">
      <TabPane v-for="(data, index) in templatesList" :key="index" :tab="index">
        <div class="grid grid-cols-2 gap-4">
          <div class="p-4 border border-solid border-[#DEDDDC] grid grid-cols-4" v-for="(data2, index) in templatesList" :key="index">
            <div>
              <Image class="!w-[100px]" :src="data.img" />
            </div>
            <div class="col-span-3">
              <div>123</div>
            </div>
          </div>
        </div>
      </TabPane>
    </Tabs>
    <div class="text-center mt-8">
      <Button type="primary" ghost>{{ $t("template.cancelBtn") }}</Button>
      <Button type="primary" class="ml-4">{{ $t("template.nextBtn") }}</Button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { reactive, ref, onMounted } from "vue";
import Header from "../../Header.vue";
import { apiGetTemplates } from "@/apis/template";
import { Tabs, TabPane, Image, Button } from 'ant-design-vue';

const activeKey = ref(2);

const templatesList = reactive([])


onMounted(async () => {
  // getTemplates();
});

const getTemplates = async () => {
    try {
      const data = await apiGetTemplates();
      Object.assign(templatesList, data); //赋值
    } catch (error: any) {
      console.log("erro:",error)
    }
  };

// const templatesList = reactive([
//   {
//     title: "Hamster-pipeline-1",
//     description: "11111",
//     status: "success",
//     img: "https://zos.alipayobjects.com/rmsportal/jkjgkEfvpUPVyRjUImniVslZfWPnJuuZ.png",
//   },
//   {
//     title: "Hamster-pipeline-2",
//     description: "222222",
//     status: "fail",
//     img: "https://zos.alipayobjects.com/rmsportal/jkjgkEfvpUPVyRjUImniVslZfWPnJuuZ.png",
//   },
//   {
//     title: "Hamster-pipeline-3",
//     description: "333333",
//     status: "success",
//     img: "https://zos.alipayobjects.com/rmsportal/jkjgkEfvpUPVyRjUImniVslZfWPnJuuZ.png",
//   },
//   {
//     title: "Hamster-pipeline-4",
//     description: "44444",
//     status: "fail",
//     img: "https://zos.alipayobjects.com/rmsportal/jkjgkEfvpUPVyRjUImniVslZfWPnJuuZ.png",
//   },
// ]);


</script>

<style scoped lang="less">
</style>