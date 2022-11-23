<template>
  <div class="bg-[#FFFFFF] rounded-[12px] leading-[24px] mx-20">
    <div class="bg-[#121211] p-4 rounded-tl-[12px] rounded-tr-[12px]">
      <div class="flex justify-between">
        <div class="text-[24px] font-semibold text-[#FFFFFF]">{{ $t("template.title") }}</div>
        <div class="help-text">{{ $t("template.helpDoc") }}</div>
      </div>
      <div class="text-[#979797] text-[14px] mt-2">{{ $t("template.titleDesc") }}</div>
    </div>
    <div class="p-4 rounded-bl-[12px] rounded-br-[12px] border border-solid border-[#EFEFEF] box-border">
      <Tabs :defaultActiveKey="activeKey">
        <TabPane key="0" :tab="$t('template.allText')">
          <div class="card-div">
            <div class="card-item" @click="setCurrId(item.id)" :class="{'check-border':checkCurrId === item.id }" v-for="(item, index) in allTemplatesList" :key="index">
              <div class="card-img-div">
                <img
                  :src="getImageURL(`${item.imageName}.png`)"
                />
              </div>
              <div class="col-span-5">
                <div class="card-title">{{ item.name }}</div>
                <div class="card-desc show-line">{{ item.description }}</div>
              </div>
            </div>
          </div>
        </TabPane>
        <TabPane v-for="(data, index) in templatesList" :key="index+1" :tab="data.tag">
          <div class="card-div">
            <div class="card-item" @click="setCurrId(item.id)" :class="{'check-border':checkCurrId === item.id }" v-for="(item, index2) in data.items" :key="index2">
              <div>
                <img
                  :src="getImageURL(`${item.imageName}.png`)"
                />
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
        <Button type="primary" ghost @click="backStep">{{ $t("template.cancelBtn") }}</Button>
        <Button type="primary" class="ml-4" @click="nextStep">{{ $t("template.nextBtn") }}</Button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { reactive, ref, onMounted } from "vue";
import { useRouter } from 'vue-router';
import { apiGetTemplates } from "@/apis/template";
import { Tabs, TabPane, Button } from 'ant-design-vue';
import useAssets from "@/stores/useAssets";
const { getImageURL } = useAssets()
 
const router = useRouter();
const activeKey = ref('0');
const checkCurrId = ref(0);

const templatesList = reactive([]);
const allTemplatesList = reactive([
{
"id": 10, //模版ID
"name": "Smart Contract Quality Check", //模版名称
"description": "Perform quality checks on a smart contract and output relevant reports.", //模版描述
"tag": "GENERAL_TEMPLATE", //模版类型，用于分组
"imageName": "contract_check" //图片名
},
{
"id": 12, //模版ID
"name": "Smart Contract Deployment", //模版名称
"description": "Deployment operation for a smart contract.", //模版描述
"tag": "GENERAL_TEMPLATE", //模版类型，用于分组
"imageName": "contract_deploy" //图片名
},
{
"id": 11, //模版ID
"name": "Smart Contract Test", //模版名称
"description": "Perform quality checks on a smart contract and output relevant reports.", //模版描述
"tag": "GENERAL_TEMPLATE", //模版类型，用于分组
"imageName": "contract_test" //图片名
},
{
"id": 1, //模版ID
"name": "Substrate + Smart Contracts", //模版名称
"description": "This template is used for smart contracts written based on the Substact framework, enabling fully automated code checkout -> check contract quality -> compile contract -> test contract -> deploy contract.", //模版描述
"tag": "SMART_CONTRACT_TEMPLATE", //模版类型，用于分组
"imageName": "substrate" //图片名
},
{
"id": 2, //模版ID
"name": "Hardhat + Smart Contracts", //模版名称
"description": "This template is used for smart contracts written based on Hardhat framework to fully automate code checkout -> check contract quality -> compile contract -> test contract -> deploy contract.", //模版描述
"tag": "SMART_CONTRACT_TEMPLATE", //模版类型，用于分组
"imageName": "hardhat" //图片名
},
{
"id": 3, //模版ID
"name": "Truffle + Smart Contracts", //模版名称
"description": "This template is used for smart contracts written based on Truffle framework, enabling fully automated code checkout -> check contract quality -> compile contract -> test contract -> deploy contract.", //模版描述
"tag": "SMART_CONTRACT_TEMPLATE", //模版类型，用于分组
"imageName": "truffle" //图片名
},
{
"id": 4, //模版ID
"name": "Substrate + DApp", //模版名称
"description": "This template is used for DApps written on Substact framework to fully automate code checkout -> check contract quality -> compile contract -> test contract -> deploy contract -> check front-end code quality -> edit front-end...", //模版描述
"tag": "DAPP_TEMPLATE", //模版类型，用于分组
"imageName": "substrate" //图片名
},
{
"id": 5, //模版ID
"name": "Hardhat + Smart Contracts", //模版名称
"description": "This template is used for DApps written on Hardhat framework to fully automate code checkout -> check contract quality -> compile contract -> test contract -> deploy contract -> check front-end code quality -> edit front-end...", //模版描述
"tag": "DAPP_TEMPLATE", //模版类型，用于分组
"imageName": "hardhat" //图片名
},
{
"id": 6, //模版ID
"name": "Truffle + DApp", //模版名称
"description": "This template is used for DApps written based on Truffle framework to fully automate code checkout -> check contract quality -> compile contract -> test contract -> deploy contract -> check front-end code quality -> edit front-end...", //模版描述
"tag": "DAPP_TEMPLATE", //模版类型，用于分组
"imageName": "truffle" //图片名
},
{
"id": 7, //模版ID
"name": "Substrate + DApp (front end only)", //模版名称
"description": "This template is used for DApps written based on Substact framework to fully automate code checkout -> check front-end code quality -> edit front-end code -> deploy front-end code.", //模版描述
"tag": "DAPP_TEMPLATE(Frontend)", //模版类型，用于分组
"imageName": "substrate" //图片名
},
{
"id": 8, //模版ID
"name": "Hardhat + DApp (front end only)", //模版名称
"description": "This template is used for DApps written on Hardhat framework to fully automate code checkout -> check front-end code quality -> edit front-end code -> deploy front-end code.", //模版描述
"tag": "DAPP_TEMPLATE(Frontend)", //模版类型，用于分组
"imageName": "hardhat" //图片名
},
{
"id": 9, //模版ID
"name": "Truffle + DApp (front end only)", //模版名称
"description": "This template is used for DApps written on Truffle framework to fully automate code checkout -> check front-end code quality -> edit front-end code -> deploy front-end code.", //模版描述
"tag": "DAPP_TEMPLATE(Frontend)", //模版类型，用于分组
"imageName": "truffle" //图片名
}
]);

onMounted(async () => {
  getTemplates();
});

const getTemplates = async () => {

  try {
    const { data } = await apiGetTemplates();
    Object.assign(allTemplatesList, data); //赋值
    //拆分相同 tabs 下的数据
    const templates: string | any[] = [];
    const templateTabs: any[] = [];
    data.forEach((item: any) => {
      
      if (templateTabs.includes(item.tag)) {
        templates.forEach((subItem, index) => {
          if (subItem.tag === item.tag) {
            templates[index]['items'].push(item);
          }
        })
      } else {
        templateTabs.push(item.tag);
        templates.push({ tag: item.tag, items: [item] });
      }
    });
    Object.assign(templatesList, templates); //赋值
    console.log("templateTabs:",templateTabs)
    console.log("templates:",templates)
  } catch (error: any) {
    console.log("erro:",error)
  }
};
const setCurrId = async (id: number) => {
  checkCurrId.value = id;
}
const backStep = async () => {
  router.push({ path: '/pipeline' });
}
const nextStep = async () => {
  router.push({ path: '/create/config/'+ checkCurrId.value });
}
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
  @apply p-4 grid grid-cols-6 gap-4 cursor-pointer;
  border: 1px solid #EFEFEF;
  /* 投影 */
  box-shadow: 3px 3px 12px rgba(203, 217, 207, 0.1);
  border-radius: 12px;
}
.card-img-div{
  @apply flex justify-center;
}
.card-item img{
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
  width: 120px;
  height: 40px;
}
:deep(.ant-btn-primary), :deep(.ant-btn-primary:hover), :deep(.ant-btn-primary:focus){
  border-color: @baseColor;
  background: @baseColor;
}

:deep(.ant-btn-background-ghost.ant-btn-primary), :deep(.ant-btn-background-ghost.ant-btn-primary:hover), :deep(.ant-btn-background-ghost.ant-btn-primary:focus){
  border-color: @baseColor;
  color: @baseColor;
}
.check-border{
   border-color: @baseColor;
}
</style>