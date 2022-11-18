<template>
  <a-modal v-model:visible="visible" :closable="false" :footer="null" ref="modal"
    style="top: 0px; margin-right: 0px; padding: 0px;" width="800px">
    <div class="px-[24px]" ref="fullRef">
      <div class="flex justify-between">
        <span class="text-[24px] text-[#000000] font-semibold mb-[28px]">{{ props.text }}</span>
        <span class="text-[#28C57C]  cursor-pointer" @click="toggle">
          <img src="@/assets/icons/full.svg" class="w-[18px] mr-[10px]" />
          <span class="align-middle">全屏</span>
        </span>
      </div>
      <div class="content"></div>
    </div>
  </a-modal>
</template>
<script lang="ts">
import { ref, defineComponent, toRefs, reactive } from 'vue'
export default defineComponent({
  props: {
    text: { type: String, default: "" },
    content: { type: String, default: "" }
  },
  setup(props, context) {
    const root = ref()
    const fullRef = ref()
    const state = reactive({
      visible: false,
      bodyHeight: document.body.clientHeight - 162 + "px",
    })

    const toggle = () => {
      console.log(fullRef.value)
      const ro = new ResizeObserver((entries, observer) => {
        for (const entry of entries) {
          const { left, top, width, height } = entry.contentRect;
          console.log(left, top, width, height)
          // fullRef.value.style.height = height + 'px'
          // console.log('Element:', entry.target);
          // console.log(`Element's size: ${width}px x ${height}px`);
          // console.log(`Element's paddings: ${top}px ; ${left}px`);
        }
      });

      // ro.observe(fullRef.value);
      ro.observe(document.body);
    }
    // async function toggle() {
    //   await fullscreen.toggle(root.value.querySelector('.fullscreen-wrapper'), {
    //     teleport: state.teleport,
    //     pageOnly: true,
    //     callback: (isFullscreen) => {
    //       state.visible = true
    //       state.fullscreen = isFullscreen
    //     },
    //   })
    //   state.fullscreen = fullscreen.isFullscreen
    // }

    const showVisible = () => {
      state.visible = true
      // initResize()


    }


    return {
      root,
      props,
      fullRef,
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

.ant-modal-body {
  padding: 0px !important;
}

.fullStyle {
  border-radius: 0;
}
</style>