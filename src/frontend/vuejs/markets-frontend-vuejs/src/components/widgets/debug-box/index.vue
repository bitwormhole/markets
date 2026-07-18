<script>
export default {
  name: 'widgets-debug-box',

  computed: {
    theTitleStr() {
      let str = this.title
      if (str == null || str == '') {
        str = 'DebugBox'
      }
      return str
    },
  },

  data() {
    const hide = true
    return { hide }
  },

  methods: {
    handleClickTitleBar() {
      let h = this.hide
      this.hide = !h
    },
  },

  props: {
    title: String,
  },
}
</script>

<style scoped>
.debug-box {
  background-color: yellow;
  border-color: black;
  border-style: solid;
  border-width: 10px;
}

.debug-box-title {
  padding: 5px;
  border-color: gray;
  border-style: solid;
  border-width: 1px;
}
</style>

<template>
  <div class="debug-box">
    <div class="debug-box-title" @click="handleClickTitleBar">{{ theTitleStr }}</div>
    <div class="debug-box-client" v-show="!hide">
      <slot></slot>
    </div>
  </div>
</template>
