<script lang="js">

import InnerPage from './permission-query-page.vue';

export default {

  name: "permission-index-page",

  components: { InnerPage },

  data() {
    return {}
  },

  methods: {
    init() {
      this.fetch();
    },

    fetch() {
      // let method = 'GET'
      // let url = '/api/v1/examples'
      // let params = {
      //   limit: 23,
      //   offset: 45,
      // }
      // theAxiosLib.execute({ method, url, params })
    },
  },

  mounted() {
    this.init()
  },

  props: {}
}

</script>

<style></style>

<template>
  <div>
    <InnerPage> </InnerPage>
  </div>
</template>
