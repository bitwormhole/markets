<script lang="js">
import { ElDialog } from 'element-plus'

export default {
  name: 'widgets-log-msg-list',

  components: { ElDialog },

  data() {
    return {}
  },

  props: {
    item: Object, // a log-record-item
  },
}
</script>

<style scoped></style>

<template>
  <ElDialog title="Logs(list)"></ElDialog>
</template>
