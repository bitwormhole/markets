<script lang="js">
// file : 'scan-db-tables.vue'

import { useAxiosLib } from '@/stores/axios'
import { ElButton } from 'element-plus'

const theAxiosLib = useAxiosLib()

export default {
  name: 'views-dev-scan-db-tables',

  components: { ElButton },

  methods: {
    init() {},

    scan() {
      let method = 'POST'
      let url = '/api/v1/admin/tables/do/scan'
      let data = {}

      theAxiosLib.execute({ method, url, data })
    },
  },

  mounted() {
    // this.init()
  },

  props: {},
}
</script>

<style scoped></style>

<template>
  <frame-for-developer>
    <h1>Scan DB Tables</h1>

    <el-button type="primary" @click="scan"> Scan </el-button>
  </frame-for-developer>
</template>
