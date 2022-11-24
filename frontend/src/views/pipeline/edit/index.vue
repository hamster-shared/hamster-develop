<template>
  <div class="bg-[#FFFFFF] rounded-[12px] leading-[24px]">
    <div class="p-4 rounded-tl-[12px] rounded-tr-[12px]">
      <div class="flex justify-between">
        <div>
          <div class="text-[24px] font-semibold">
            {{ templateName }}
          </div>
          <div class="text-[#979797] text-[14px] mt-2">
            Pipelinefile
          </div>
        </div>
        <div>
          <a-button type="primary" @click="backStep" ghost>{{ $t("template.discardChange") }}</a-button>
          <a-button type="primary" class="ml-4" @click="showModal">{{ $t("template.saveBtn") }}</a-button>
        </div>
      </div>
    </div>
    <div
      class="h-screen mx-4 rounded-[12px]"
    >
      <CodeEditor @getYamlValue="getYamlValue" :readOnly="false" :value="codeValue"></CodeEditor>
    </div>
  </div>

  <a-modal :getContainer = "false"
    v-model:visible="visible"
    :title="$t('template.modalTitle')"
    :footer="null"
  >
    <div class="mb-8">
      <div class="mb-2 flex justify-between">
        Pipeline Name
      </div>
      <a-input v-model:value="pipelineName" placeholder="Pipeline Name" allow-clear />
    </div>
    <div class="text-center">
      <a-button type="primary" @click="visible = false" ghost>{{ $t("template.cancelBtn") }}</a-button>
      <a-button type="primary" :loading="confirmLoading" class="ml-4" @click="handleOk">{{ $t("template.saveBtn") }}</a-button>
    </div>
  </a-modal>
</template>
<script lang="ts" setup>
import { ref, onMounted } from "vue";
import { useRoute, useRouter } from 'vue-router';
import { apiGetPipelineByName, apiEditPipeline } from "@/apis/template";
import CodeEditor from "../create/config/components/CodeEditor.vue";
import { message } from "ant-design-vue";

  const codeValue = ref<String>();

  const router = useRouter();
  const { params } = useRoute();
  const templateName = ref(params.id);

  const confirmLoading = ref<boolean>(false);
  const visible = ref<boolean>(false);
  const pipelineName = ref('');

  const showModal = async () => {
    visible.value = true;
  };

  const getYamlValue = async (value: String) => {
    codeValue.value = value;
  }

  const handleOk = async () => {
    confirmLoading.value = true;
    try {
      const result = await apiEditPipeline(templateName.value.toString(), pipelineName.value ,codeValue.value);
      if (result.code === 200) {
        visible.value = false;
        message.success(result.message);
        router.push({ path: '/pipeline' });
      }
      confirmLoading.value = false;
    } catch (error: any) {
      confirmLoading.value = false;
      console.log("erro:", error);
      message.error("Failed");
    }
};
  
  onMounted(async () => {
    getTemplatesById(templateName.value.toString());
  });
  
  const getTemplatesById = async (templateName: String) => {

    try {
      const { data } = await apiGetPipelineByName(templateName);
      codeValue.value = data;
      
    } catch (error: any) {
      console.log("erro:",error)
    }
  };
  const backStep = async () => {
    router.push({ path: '/pipeline' });
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
:deep(.ant-btn-primary), :deep(.ant-btn-primary:hover), :deep(.ant-btn-primary:focus){
  border-color: @baseColor;
  background: @baseColor;
}
:deep(.ant-btn-background-ghost.ant-btn-primary), :deep(.ant-btn-background-ghost.ant-btn-primary:hover), :deep(.ant-btn-background-ghost.ant-btn-primary:focus){
  border-color: @baseColor;
  color: @baseColor;
}
:deep(.ant-input),:deep(.ant-input-affix-wrapper){
  border-color: #EFEFEF;
  border-radius: 6px;
}
@placeholderColor: #BCBEBC;
:deep(input::-webkit-input-placeholder) { /* WebKit browsers */
  color: @placeholderColor;
}
:deep(input:-moz-placeholder) { /* Mozilla Firefox 4 to 18 */
  color: @placeholderColor;
}
:deep(input::-moz-placeholder) { /* Mozilla Firefox 19+ */
  color: @placeholderColor;
}
:deep(input:-ms-input-placeholder) { /* Internet Explorer 10+ */
  color: @placeholderColor;
}
</style>
