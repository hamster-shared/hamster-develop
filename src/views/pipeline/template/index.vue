<template>
  <Header />
  <div class="bg-[#FFFFFF] rounded-[12px] m-5 leading-[24px]">
    <div class="bg-[#121211] p-4 rounded-tl-[12px] rounded-tr-[12px]">
      <div class="flex justify-between">
        <div class="text-[24px] font-semibold text-[#FFFFFF]">{{ $t("template.title") }}</div>
        <div class="help-text">{{ $t("template.helpDoc") }}</div>
      </div>
      <div class="text-[#979797] text-[14px] mt-2">{{ $t("template.titleDesc") }}</div>
    </div>
    <div class=" p-4">
      <Tabs :defaultActiveKey="activeKey">
        <TabPane key="0" tab="全部">
          <div class="card-div">
            <div class="card-item" v-for="(item, index) in allTemplatesList" :key="index">
              <div class="card-img-div">
                <Image src="item.image" />
              </div>
              <div class="col-span-5">
                <div class="card-title">{{ item.name }}</div>
                <div class="card-desc show-line">{{ item.description }}</div>
              </div>
            </div>
          </div>
        </TabPane>
        <TabPane v-for="(data, index1) in templatesList.groups" :key="index1+1" :tab="data.name">
          <div class="card-div">
            <div class="card-item" v-for="(item, index2) in data.items" :key="index2">
              <div>
                <Image :src="item.image" />
              </div>
              <div class="col-span-5">
                <div class="card-title">{{ item.name }}</div>
                <div class="card-desc show-line">{{ item.description }}</div>
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
  </div>
</template>

<script setup lang="ts">
import { reactive, ref, onMounted } from "vue";
import Header from "../../Header.vue";
import { apiGetTemplates } from "@/apis/template";
import { Tabs, TabPane, Image, Button } from 'ant-design-vue';

const activeKey = ref('0');

const templatesList = reactive([]);
const allTemplatesList = ref([]);

onMounted(async () => {
  getTemplates();
});

const getTemplates = async () => {

  try {
    const data = await apiGetTemplates();
    Object.assign(templatesList, data.templates); //赋值
    //全部 tabs 下的数据
    templatesList.groups.forEach(data => {
      data.items.forEach((item) => {
        allTemplatesList.value.push(item);
      });
    });
  } catch (error: any) {
    console.log("erro:",error)
  }
};
</script>

<style scoped lang="less">
@baseColor: #28C57C;
.help-text{
  color: @baseColor;
  font-size: 14px;
}
:deep(.ant-tabs){
  color: #7B7D7B;
}
:deep(.ant-tabs-tab-btn:hover){
  color: @baseColor;
}
:deep(.ant-tabs-ink-bar) {
  background: @baseColor;
}
:deep(.ant-tabs-tab.ant-tabs-tab-active .ant-tabs-tab-btn){
  @apply font-semibold;
  color: #121211;
}
.card-div{
  @apply grid grid-cols-2 gap-4;
}
.card-item{
  @apply p-4 grid grid-cols-6 gap-4;
  border: 1px solid #EFEFEF;
  /* 投影 */
  box-shadow: 3px 3px 12px rgba(203, 217, 207, 0.1);
  border-radius: 12px;
}
.card-img-div{
  @apply flex justify-center;
}
.card-item :deep(.ant-image-img){
  width: 64px;
  height: 64px;
  border-radius: 12px;
}
.card-title{
  @apply font-semibold;
  font-size: 16px;
  color: #121211;
}
.card-desc{
  font-size: 12px;
  color: #7B7D7B;
  margin-top: 4px;
}
.show-line{
  display: -webkit-box;
  -webkit-box-orient: vertical;
  -webkit-line-clamp: 2;
  overflow: hidden;
}
:deep(.ant-btn){
  border-radius: 6px;
}
:deep(.ant-btn-primary){
  border-color: @baseColor;
  background: @baseColor;
}
:deep(.ant-btn-background-ghost.ant-btn-primary){
  border-color: @baseColor;
  color: @baseColor;
}
</style>