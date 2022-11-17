<template>
  <div class="title"><label>Pipelinefile Preview</label></div>
  <div ref="editContainer" class="code-editor"></div>
</template>
<script>
import { getCurrentInstance, onMounted, watch } from "vue";
import * as monaco from "monaco-editor/esm/vs/editor/editor.main.js";
import JsonWorker from "monaco-editor/esm/vs/language/json/json.worker?worker";
self.MonacoEnvironment = {
  getWorker() {
    return new JsonWorker();
  },
};
export default {
  name: "CodeEditor",
  props: {
    value: String,
  },
  setup(props, { emit }) {
    let monacoEditor = null;
    const { proxy } = getCurrentInstance();

    watch(
      () => props.value,
      (value) => {
        // 防止改变编辑器内容时光标重定向
        if (value !== monacoEditor?.getValue()) {
          monacoEditor.setValue(value);
        }
      }
    );

    onMounted(() => {
      monaco.editor.defineTheme("custom", {
        base: "vs",
        inherit: true,
        rules: [{ background: "#F7F8F9" }],
        colors: {
          // 相关颜色属性配置
          // 'editor.foreground': '#000000',
          "editor.background": "#F7F8F9", //背景色
          // 'editorCursor.foreground': '#8B0000',
          // 'editor.lineHighlightBackground': '#0000FF20',
          // 'editorLineNumber.foreground': '#008800',
          // 'editor.selectionBackground': '#88000030',
          // 'editor.inactiveSelectionBackground': '#88000015'
        },
      });
      //设置自定义主题
      monaco.editor.setTheme("custom");
      monacoEditor = monaco.editor.create(proxy.$refs.editContainer, {
        value: props.value,
        readOnly: true,
        language: "yaml",
        theme: "custom",
        automaticLayout: true,
        selectOnLineNumbers: false,
        renderSideBySide: false,
        minimap: {
          enabled: false,
        },
        fontSize: 16,
        fontWeight: "400",
        scrollBeyondLastLine: false,
        overviewRulerBorder: false,
      });
      // 监听值变化
      monacoEditor.onDidChangeModelContent(() => {
        const currenValue = monacoEditor?.getValue();
        emit("update:value", currenValue);
      });
    });
    return {};
  },
};
</script>
<style scoped lang="less">
.title {
  display: flex;
  flex-direction: row;
  align-items: flex-start;
  padding: 16px;
  gap: 10px;
  height: 56px;
  background: #121211;

  label {
    height: 24px;
    font-style: normal;
    font-weight: 600;
    font-size: 16px;
    line-height: 24px;
    /* identical to box height, or 150% */
    display: flex;
    align-items: flex-end;
    color: #FFFFFF;
  }
}
.code-editor {
  width: 100%;
  height: 100vh;
}

</style>
