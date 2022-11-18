<template>
  <div class="bg-[#FFFFFF] rounded-[12px] leading-[24px]">
    <div class="bg-[#121211] p-4 rounded-tl-[12px] rounded-tr-[12px]">
      <div class="flex justify-between">
        <div class="text-[24px] font-semibold text-[#FFFFFF]">
          Substact + 智能合约
        </div>
      </div>
      <div class="text-[#979797] text-[14px] mt-2">
        该模版用于基于Substact框架编写的智能合约，实现全自动检出代码 ->
        检查合约质量 -> 编译合约 -> 测试合约->部署合约
      </div>
    </div>
    <div
      class="p-4 rounded-bl-[12px] rounded-br-[12px] border border-solid border-[#EFEFEF] box-border"
    >
      <a-row class="create" :gutter="24">
        <a-col :span="8">
          <div class="bg-[#EFEFEF] p-4 rounded-tl-[12px] rounded-tr-[12px]">
            <div class="flex justify-between">
              <div class="text-[16px] font-semibold text-[#121211]">
                Pipeline Process
              </div>
            </div>
          </div>
          <div
            class="p-4 rounded-bl-[12px] rounded-br-[12px] border border-solid border-[#EFEFEF] box-border"
          >
            <a-form
              :wrapper-col="{
                xs: { span: 24 },
                sm: { span: 12 },
                md: { span: 10 },
              }"
              :label-col="{ xs: { span: 24 }, sm: { span: 7 } }"
            >
              <a-form-item label="标题：" v-bind="validateInfos.title">
                <a-input v-model:value="modelRef.title" placeholder="请输入" />
              </a-form-item>
              <a-form-item label="起止日期" v-bind="validateInfos.date">
                <a-range-picker
                  v-model:value="modelRef.date"
                  style="width: 100%"
                />
              </a-form-item>

              <a-form-item label="下拉选择" v-bind="validateInfos.select">
                <a-select
                  v-model:value="modelRef.select"
                  placeholder="请选择"
                  allowClear
                >
                  <a-select-option value="1">select1</a-select-option>
                  <a-select-option value="2">select2</a-select-option>
                  <a-select-option value="3">select3</a-select-option>
                </a-select>
              </a-form-item>

              <a-form-item label="单选按钮1">
                <a-radio-group v-model:value="modelRef.radio1">
                  <a-radio value="1">item 1</a-radio>
                  <a-radio value="2">item 2</a-radio>
                  <a-radio value="3">item 3</a-radio>
                </a-radio-group>
              </a-form-item>

              <a-form-item label="单选按钮2" v-bind="validateInfos.radio2">
                <a-radio-group v-model:value="modelRef.radio2">
                  <a-radio-button value="1">item 1</a-radio-button>
                  <a-radio-button value="2">item 2</a-radio-button>
                  <a-radio-button value="2">item 3</a-radio-button>
                </a-radio-group>
              </a-form-item>

              <a-form-item label="复选框" v-bind="validateInfos.checkbox">
                <a-checkbox-group v-model:value="modelRef.checkbox">
                  <a-checkbox value="1" name="type"> Online </a-checkbox>
                  <a-checkbox value="2" name="type"> Promotion </a-checkbox>
                  <a-checkbox value="3" name="type"> Offline </a-checkbox>
                </a-checkbox-group>
              </a-form-item>

              <a-form-item label="备注" v-bind="validateInfos.remark">
                <a-textarea v-model:value="modelRef.remark" />
              </a-form-item>

              <a-form-item
                :wrapper-col="{
                  xs: { span: 24, offset: 0 },
                  sm: { span: 10, offset: 7 },
                }"
              >
                <a-button
                  type="primary"
                  @click="handleSubmit"
                  :loading="submitLoading"
                >
                  提交
                </a-button>
                <a-button
                  type="primary"
                  ghost
                  @click="resetFields"
                  style="margin-left: 10px"
                >
                  重置
                </a-button>
              </a-form-item>
            </a-form>
          </div>
        </a-col>
        <a-col :span="16">
          <CodeEditor :value="codeValue"></CodeEditor>
        </a-col>
      </a-row>
      <div class="text-center mt-8">
        <a-button type="primary" ghost>{{ $t("template.cancelBtn") }}</a-button>
        <a-button type="primary" class="ml-4">{{
          $t("template.nextBtn")
        }}</a-button>
      </div>
    </div>
  </div>
</template>
<script lang="ts">
import { defineComponent, reactive, ref } from "vue";
import type { Ref } from "vue";
import { message, Form } from "ant-design-vue";
const useForm = Form.useForm;
import CodeEditor from "./components/CodeEditor.vue";
import type { Props, validateInfos } from "ant-design-vue/lib/form/useForm";
import type { FormDataType } from "./data";

interface FormBasicPageSetupData {
  resetFields: (newValues?: Props) => void;
  validateInfos: validateInfos;
  modelRef: FormDataType;
  submitLoading: Ref<boolean>;
  handleSubmit: (e: MouseEvent) => void;
}

export default defineComponent({
  name: "CreatePipeline",
  components: {
    CodeEditor,
  },
  setup(): FormBasicPageSetupData {
    // 表单值
    const modelRef = reactive<FormDataType>({
      title: "",
      date: [],
      select: "",
      radio1: "",
      radio2: "",
      checkbox: [],
      remark: "",
    });
    // 表单验证
    const rulesRef = reactive({
      title: [
        {
          required: true,
          message: "必填",
        },
      ],
      date: [
        {
          required: true,
          message: "必填",
          trigger: "change",
          type: "array",
        },
      ],
      select: [
        {
          required: true,
          message: "请选择",
        },
      ],
      radio1: [],
      radio2: [
        {
          required: true,
          message: "请选择",
        },
      ],
      checkbox: [],
      remark: [],
    });
    // 获取表单内容
    const { resetFields, validate, validateInfos } = useForm(
      modelRef,
      rulesRef
    );
    // 重置 validateInfos 如果用到国际化需要此步骤
    //const validateInfosNew = useI18nAntdFormVaildateInfos(validateInfos);

    // 登录loading
    const submitLoading = ref<boolean>(false);
    // 登录
    const handleSubmit = async (e: MouseEvent) => {
      e.preventDefault();
      submitLoading.value = true;
      try {
        const fieldsValue = await validate<FormDataType>();
        message.success("提交成功");
        resetFields();
      } catch (error) {
        // console.log('error', error);
      }
      submitLoading.value = false;
    };

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

    return {
      resetFields,
      validateInfos,
      modelRef,
      submitLoading,
      handleSubmit,
      codeValue,
    };
  },
});
</script>
<style scoped lang="less">
.create {
  height: 100%;
}
@baseColor: #28c57c;
:deep(.ant-btn) {
  border-radius: 6px;
}
:deep(.ant-btn-primary) {
  border-color: @baseColor;
  background: @baseColor;
  width: 120px;
  height: 40px;
}
:deep(.ant-btn-background-ghost.ant-btn-primary) {
  border-color: @baseColor;
  color: @baseColor;
}
</style>
