<template>
  <div class="inline-block mr-1" @click="handleClick">
    <div v-if="!checked">
      <img src="@/assets/icons/checkbox-uncheck.svg" />
    </div>
    <div v-else>
      <img src="@/assets/icons/checkbox-checked.svg" />
    </div>
  </div>
  <span class="text-[#7B7D7B] text-sm">{{ $t("template.buildAfterCreation") }}</span>
</template>

<script setup lang="ts">
import { ref, defineEmits } from 'vue';

const props = defineProps<{
  checked: Boolean;
}>();

const checked = ref(props.checked)

const emits = defineEmits(['update:checked'])

const handleClick = ()=>{
  checked.value = !checked.value
  emits('update:checked', checked.value)
}

</script>