<script lang="js">

import MyQueryPage from './user-query-page.vue';

export default {

  name: "users-index-page",

  components: { MyQueryPage },

  data() {
    return {}
  },

  methods: {
    fetch() {
    },
  },

  mounted() {
    // this.init()
  },

  props: {}
}

</script>

<style></style>

<template>
  <div>
    <MyQueryPage />
  </div>
</template>
