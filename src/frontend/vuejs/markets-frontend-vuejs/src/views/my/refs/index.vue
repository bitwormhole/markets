<script lang="js">
import InnerPage from '@/components/examples/example-index-page.vue'

export default {
  name: 'views-my-xxxx-index',

  components: { InnerPage },

  data() {
    return {}
  },

  methods: {},

  mounted() {},
}
</script>

<style scoped></style>

<template>
  <frame-for-user>
    <h1>todo : Title</h1>
    <InnerPage> </InnerPage>
  </frame-for-user>
</template>
