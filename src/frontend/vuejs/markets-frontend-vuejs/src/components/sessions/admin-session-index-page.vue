<script lang="js">


import InnerPage from './admin-session-query-page.vue';


export default {

  name: "sessions-admin-index-page",

  components: { InnerPage },

  data() {
    return {}
  },

  methods: {
    fetch() {
      // theSessionStore.fetch()
    },
  },

  mounted() {
    this.fetch()
  },

  props: {}
}

</script>

<style></style>

<template>
  <div>
    <InnerPage />
  </div>
</template>
