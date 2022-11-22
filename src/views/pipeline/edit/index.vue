<template>
  <div class="bg-[#FFFFFF] rounded-[12px] leading-[24px] mx-20">
    <div class="p-4 rounded-tl-[12px] rounded-tr-[12px]">
      <div class="flex justify-between">
        <div>
          <div class="text-[24px] font-semibold">
            Hamster - pipeline
          </div>
          <div class="text-[#979797] text-[14px] mt-2">
            Pipelinefile
          </div>
        </div>
        <div>
          <a-button type="primary" @click="backStep" ghost>放弃修改</a-button>
          <a-button type="primary" class="ml-4" @click="showModal">保存</a-button>
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
    title="修改 PipeLineName"
    :footer="null"
  >
    <div class="mb-8">
      <div class="mb-2 flex justify-between">
        Pipeline Name
      </div>
      <a-input v-model:value="pipelineName" placeholder="Pipeline Name" allow-clear />
    </div>
    <div class="text-center">
      <a-button type="primary" @click="visible = false" ghost>取消</a-button>
      <a-button type="primary" :loading="confirmLoading" class="ml-4" @click="handleOk">保存</a-button>
    </div>
  </a-modal>
</template>
<script lang="ts" setup>
import { reactive, ref, onMounted } from "vue";
import { useRoute, useRouter } from 'vue-router';
import { apiGetTemplatesById, apiEditPipeline } from "@/apis/template";
import CodeEditor from "../create/config/components/CodeEditor.vue";
import { message } from "ant-design-vue";

  const codeValue = ref<String>(
    "version: 1.0\n" +
      "name: my-test\n" +
      "stages:\n" +
      "  git-clone:\n" +
      "    steps:\n" +
      "      - name: git-clone\n" +
      "        code: 1\n" +
      "        uses: git-checkout\n" +
      "        with:\n" +
      "          url: https://gitee.com/mohaijiang/spring-boot-example.git\n" +
      "          branch: master\n" +
      "  code-compile:\n" +
      "    needs:\n" +
      "      - git-clone\n" +
      "    steps:\n" +
      "      - name: code-compile\n" +
      "        code: 2\n" +
      "        runs-on: maven:3.5-jdk-8\n" +
      "        run: |\n" +
      "          mvn clean package -Dmaven.test.skip=true\n" +
      "      - name: save artifactory\n" +
      "        code: 2\n" +
      "        uses: hamster/artifactory\n" +
      "        with:\n" +
      "          name: some.zip\n" +
      "          path: contracts/*.sol\n" +
      "\n" +
      "  build-image:\n" +
      "    needs:\n" +
      "      - code-compile\n" +
      "    steps:\n" +
      "      - name: shell\n" +
      "        code: 3\n" +
      "        run: |\n" +
      "          docker build -t mohaijiang/spring-boot-example:20221109 ."
    );

  const router = useRouter();
  const { params } = useRoute();
  const templateId = ref(params.id);

  const confirmLoading = ref<boolean>(false);
  const visible = ref<boolean>(false);
  const pipelineName = ref('');
  
  const templateInfo = reactive({
    name: '',
    description: '',
    yaml: '',
  });

  const showModal = async () => {
    visible.value = true;
  };

  const getYamlValue = async (value: String) => {
    codeValue.value = value;
  }

  const handleOk = async () => {
    confirmLoading.value = true;
    try {
      const result = await apiEditPipeline(templateInfo.name, pipelineName.value ,codeValue.value);
      console.log("result:", result)
      if (result.code === 200) {
        confirmLoading.value = false;
        visible.value = false;
        message.success(result.message);
        router.push({ path: '/pipeline' });
      }
    } catch (error: any) {
      console.log("erro:", error);
      message.error("Failed");
    }
};
  
  onMounted(async () => {
    getTemplatesById(templateId.value.toString());
  });
  
  const getTemplatesById = async (templateId: String) => {

    try {
      const data = await apiGetTemplatesById(templateId);
      Object.assign(templateInfo, data.template); //赋值
      pipelineName.value = templateInfo.name;
      // codeValue.value = templateInfo.yaml; //@todo mock数据显示有问题，暂时用临时数据
      
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
