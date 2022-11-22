<template>
  <div class="mb-4">
    <div class="mb-2">
      run<span class="text-[#FD0505] ml-1">*</span>
    </div>
    <a-textarea @change="setYamlValue('run', formData.run)" v-model:value="formData.run" placeholder="请输入" :auto-size="{ minRows: 3, maxRows: 6 }" allow-clear />
  </div>
  <div class="mb-4">
    <div class="mb-2">
      runs-on<span class="text-[#FD0505] ml-1">*</span>
    </div>
    <a-input @change="setYamlValue('runs-on', formData.runsOn)" v-model:value="formData.runsOn" placeholder="请输入" allow-clear />
  </div>
</template>

<script setup lang="ts">
  import { toRefs, reactive } from 'vue';
  
  const props = defineProps({
    stage: String,
    index: Number,
    run: String,
    runsOn: String,
  });
  const { stage, index, run, runsOn } = toRefs(props);
  const emit = defineEmits(["setYamlCode"])

  const formData = reactive({
    run: '',
    runsOn: '', 
  });
  Object.assign(formData, {run: run?.value, runsOn: runsOn?.value});

  const setYamlValue =async (item: string, val: string) => {
    emit("setYamlCode", false, stage?.value, index?.value, item, val);
  }
</script>