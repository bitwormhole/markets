<script lang="js">
import { ElPagination, ElTableColumn } from 'element-plus';


export default {

  name: "sessions-common-table-pager",

  components: {},

  computed: {
  },

  data() {

    return {

      limit: 10,
      offset: 0,

    }
  },

  methods: {
    init() { },
  },

  mounted() {
  },

  props: {
    modelValue: Array
  }
}

</script>

<style></style>

<template>
  <div>
    <ElPagination></ElPagination>
  </div>
</template>
