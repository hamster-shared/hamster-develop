<template>
  <div class="bg-[#f9f9f9] mb-[20px] flex justify-between">
    <div class="flex flex-1">
      <img src="@/assets/icons/back-arrow.svg" class="w-[28px] h-[28px] self-center mr-2 cursor-pointer" @click="backStep"/>
      <span class="text-2xl font-semibold text-[#121211] self-center cursor-pointer" @click="backStep">back |</span>
      <span class="text-2xl font-semibold text-[#121211] self-center ml-2"> {{templateName}}</span>
      <img src="@/assets/icons/edit-pen.svg" class="w-[20px] h-[20px] self-center ml-2 cursor-pointer" @click="visible = true"/>
    </div>
    <a-button type="primary" danger class="delete-btn" @click="deletePipelineList">{{$t("template.deletePipeline")}}</a-button>
  </div>
  <div class="bg-[#FFFFFF] rounded-[12px] leading-[24px]">
    <div class="p-[24px] rounded-tl-[12px] rounded-tr-[12px]">
      <div class="flex justify-between">
        <div>
          <div class="text-[24px] font-semibold">
            {{ templateName }}
          </div>
          <div class="text-[#979797] text-[16px] mt-2">Pipelinefile</div>
        </div>
        <div>
          <a-button @click="backStep" class="normal-button">{{
            $t("template.discardChange")
          }}</a-button>
          <a-button type="primary" class="ml-4" @click="showModal">{{
            $t("template.saveBtn")
          }}</a-button>
        </div>
      </div>
    </div>
    <!-- :style="{ height: mainHeight }" -->
    <div class="mx-[24px] rounded-[12px] mb-[24px]" :style="{ height: mainHeight }">
      <CodeEditor @getYamlValue="getYamlValue" :readOnly="false" :value="codeValue"></CodeEditor>
    </div>
  </div>

  <a-modal :getContainer="false" v-model:visible="visible" :title="$t('template.modalTitle')" :footer="null">
    <div class="mb-8">
      <div class="flex justify-between mb-2">Pipeline Name</div>
      <a-input
        v-model:value="pipelineName"
        placeholder="Pipeline Name"
        @change="changeNameInput"
        @keyup="pipelineName=pipelineName.replace(/\s+/g,'')"
        allow-clear
      />
      <span class="text-[red]" v-if="showVerify">{{$t("template.cannotEmpty")}}</span>
    </div>
    <div class="text-center">
      <a-button @click="visible = false" class="normal-button">{{
        $t("template.cancelBtn")
      }}</a-button>
      <a-button type="primary" :loading="confirmLoading" class="ml-4" @click="handleOk">{{ $t("template.saveBtn") }}
      </a-button>
    </div>
  </a-modal>
</template>
<script lang="ts" setup>
import { ref, onMounted } from "vue";
import { useRoute, useRouter } from "vue-router";
import YAML from "yaml";
import { apiGetPipelineByName, apiEditPipeline } from "@/apis/template";
import { apiDeletePipelineList } from '@/apis/pipeline'
import CodeEditor from "../create/config/components/CodeEditor.vue";
import { message } from "ant-design-vue";

const codeValue = ref<String>();

const router = useRouter();
const { params } = useRoute();
const templateName = ref(params.id);

const mainHeight = ref("0px");

const confirmLoading = ref<boolean>(false);
const visible = ref<boolean>(false);
const showVerify = ref<boolean>(false)
const pipelineName = ref("");

const getYamlValue = async (value: String) => {
  codeValue.value = value;
};

const handleSave = async () => {
  const code = YAML.parse(codeValue.value);
  code.name = templateName.value
  const newCode = YAML.stringify(code)
  try {
    await apiEditPipeline(
      templateName.value.toString(),
      '',
      newCode
    );
    location.reload();
  } catch (error: any) {
    console.log("erro:", error);
    message.error("Failed");
  }
};

const changeNameInput = ()=>{
  if(pipelineName.value){
    showVerify.value = false
  } else{
    showVerify.value = true
  }
}

const handleOk = async () => {
  if(pipelineName.value){
    confirmLoading.value = true;
    const code = YAML.parse(codeValue.value);
    code.name = pipelineName.value
    const newCode = YAML.stringify(code)
    try {
      const result = await apiEditPipeline(
        templateName.value.toString(),
        pipelineName.value,
        newCode
      );
      if (result.code === 200) {
        visible.value = false;
        message.success(result.message);
        router.push({ path: "/pipeline" });
      }
      confirmLoading.value = false;
    } catch (error: any) {
      confirmLoading.value = false;
      console.log("erro:", error);
      message.error("Failed");
    }
  }else{
    showVerify.value = true
  }
};

onMounted(async () => {
  getTemplatesById(templateName.value.toString());
  mainHeight.value = document.body.clientHeight - 242 + "px";
});

const getTemplatesById = async (templateName: String) => {
  try {
    const { data } = await apiGetPipelineByName(templateName);
    codeValue.value = data;
  } catch (error: any) {
    console.log("erro:", error);
  }
};
const backStep = async () => {
  router.push({ path: "/pipeline" });
};

const deletePipelineList = async()=>{
  try {
    await apiDeletePipelineList(templateName.value)
    backStep()
  } catch(error: any) {
    console.log("erro:", error);
  }
}

</script>
<style scoped lang="less">
@baseColor: #28c57c;

:deep(.ant-btn) {
  border-radius: 6px;
}

:deep(.ant-btn-primary) {
  width: 120px;
  height: 40px;
}

:deep(.ant-btn-primary),
:deep(.ant-btn-primary:hover),
:deep(.ant-btn-primary:focus) {
  border-color: @baseColor;
  background: @baseColor;
}

:deep(.ant-btn-background-ghost.ant-btn-primary),
:deep(.ant-btn-background-ghost.ant-btn-primary:hover),
:deep(.ant-btn-background-ghost.ant-btn-primary:focus) {
  border-color: @baseColor;
  color: @baseColor;
}

:deep(.ant-input),
:deep(.ant-input-affix-wrapper) {
  border-color: #efefef;
  border-radius: 6px;
}

@placeholderColor: #bcbebc;

:deep(input::-webkit-input-placeholder) {
  /* WebKit browsers */
  color: @placeholderColor;
}

:deep(input:-moz-placeholder) {
  /* Mozilla Firefox 4 to 18 */
  color: @placeholderColor;
}

:deep(input::-moz-placeholder) {
  /* Mozilla Firefox 19+ */
  color: @placeholderColor;
}

:deep(input:-ms-input-placeholder) {
  /* Internet Explorer 10+ */
  color: @placeholderColor;
}
.normal-button {
  width: 120px;
  height: 40px;
  border-radius: 6px;
}
.delete-btn{
  background: #F52222;
  color: white;
  border-color: #F52222 !important;
}
.delete-btn:hover,.delete-btn:focus{
  background: #F52222 !important;
  color: white !important;
  border-color: #F52222 !important;
}
</style>
