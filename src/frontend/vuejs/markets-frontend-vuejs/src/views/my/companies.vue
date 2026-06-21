<script lang="js">

import MyInner from '@/components/home/my/companies/index.vue'

export default {

  name: "views-my-companies",

  components: { MyInner },

  data() {
    return {}
  },

  methods: {
  },

  mounted() {
  }
}

</script>


<style></style>


<template>
  <frame-for-user>

    <h1>厂商</h1>

    <MyInner />

  </frame-for-user>
</template>
