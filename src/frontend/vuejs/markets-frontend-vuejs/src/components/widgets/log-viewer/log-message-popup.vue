<script lang="js">
import { ElMessage } from 'element-plus'

export default {
  name: 'widgets-log-msg-popup',

  data() {
    return {}
  },

  methods: {
    isItemSuccess(item) {
      let code = item.status
      let met = item.method + ''
      if (code == 200) {
        met = met.toUpperCase()
        if (met == 'POST' || met == 'PUT' || met == 'DELETE') {
          return true
        }
      }
      return false
    },

    isItemError(item) {
      let code = item.status
      return code < 200 || 299 < code
    },

    displayMessage(item) {
      // type can be [success|info|warning|error|primary]
      const level_type_map = {
        fatal: 'error',
        error: 'error',
        warn: 'warning',
        info: 'info',
        debug: 'primary',
        trace: 'primary',
      }

      let message = item.message
      let type = level_type_map[item.level]

      if (this.isItemSuccess(item)) {
        type = 'success'
      }
      if (type == null) {
        type = 'primary'
      }

      if (this.isItemError(item)) {
        let err = item.error
        message = message + ' : ' + err
        type = 'error'
      }

      ElMessage({ message, type })
    },
  },

  watch: {
    item(it) {
      this.displayMessage(it)
    },
  },

  props: {
    modelValue: Object,
    item: Object, // a log-record-item
  },
}
</script>

<style scoped></style>

<template>
  <div></div>
</template>
