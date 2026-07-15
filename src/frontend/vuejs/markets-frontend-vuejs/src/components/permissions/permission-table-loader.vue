<script lang="js">

import { useAxiosLib } from '@/stores/axios';

const theAxiosLib = useAxiosLib()

export default {

  name: "permission-table-loader",

  components: {},

  data() {
    const item = {
      name: '',
      ptype: '',
      category: '',
      code: '',
    }
    return { item }
  },

  methods: {
    init() {
      this.fetch();
    },

    fetch() {
      let method = 'GET'
      let url = '/api/v1/admin/permissions'
      let params = {
        limit: 23,
        offset: 45,
      }
      theAxiosLib.execute({ method, url, params })
    },
  },

  mounted() {
    this.init()
  },

  props: {
    modelValue: Array,
  }
}

</script>

<style></style>

<template>
  <div>



  </div>
</template>
