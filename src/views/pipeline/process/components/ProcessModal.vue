<template>
  <a-modal v-model:visible="visible" :closable="false" :footer="null" ref="modal">
    <div>
      <div class="flex justify-between cursor-pointer">
        <span>{{ title }}</span>
        <!-- <span @click="toggle">全屏</span> -->
      </div>
      <div ref="root">
        <div class="fullscreen-wrapper">
          <a-button @click="toggle">{{ fullscreen ? '收起' : '全屏' }}</a-button>
          <div class="text-white bg-black p-[12px]">
            hhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhh
          </div>
        </div>
      </div>
    </div>
  </a-modal>
</template>
<script lang="ts">
import { ref, defineComponent, toRefs, reactive, } from 'vue'
import { api as fullscreen } from 'vue-fullscreen'
export default defineComponent({
  setup(props, context) {
    const root = ref()
    const state = reactive({
      fullscreen: false,
      teleport: true,
      visible: false,
      title: props.title || "",
    })

    async function toggle() {
      await fullscreen.toggle(root.value.querySelector('.fullscreen-wrapper'), {
        teleport: state.teleport,
        pageOnly: true,
        callback: (isFullscreen) => {
          state.visible = true
          // state.fullscreen = isFullscreen
        },
      })
      state.fullscreen = fullscreen.isFullscreen
    }

    const showVisible = () => {
      state.visible = true
    }

    return {
      root,
      ...toRefs(state),
      toggle,
      showVisible,
    }
  },
})
</script>
<style lang="less" scoped>
.fullscreen-wrapper {
  z-index: 999999;
  background-color: #000;
  overflow-y: auto;
}
</style>