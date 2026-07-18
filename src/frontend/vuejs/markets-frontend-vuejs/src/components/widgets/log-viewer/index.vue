<script lang="js">
import { useLogStore } from '@/stores/logs'

import MyDetailBox from './log-message-detail.vue'
import MyListBox from './log-message-list.vue'
import MyPopup from './log-message-popup.vue'

import MyDebugBox from '@/components/widgets/debug-box/index.vue'

const theLogStore = useLogStore()

export default {
  name: 'widgets-log-viewer',

  components: { MyDetailBox, MyListBox, MyPopup, MyDebugBox },

  computed: {
    latest() {
      return theLogStore.latest
    },

    all() {
      return theLogStore.all
    },

    theLatestText() {
      let it = this.latest
      return JSON.stringify(it, null, '\t')
    },
  },

  data() {
    const displayListBox = false
    const displayDetailBox = false
    const current_item = {}
    return { displayDetailBox, displayListBox, current_item }
  },
}
</script>

<style scoped></style>

<template>
  <span>
    <MyDetailBox v-model="displayDetailBox" :item="current_item" />
    <MyListBox v-model="displayListBox" :item="current_item" />
    <MyPopup :item="latest" />

    <MyDebugBox title="logs-index">
      <pre>  {{ theLatestText }} </pre>
    </MyDebugBox>
  </span>
</template>
