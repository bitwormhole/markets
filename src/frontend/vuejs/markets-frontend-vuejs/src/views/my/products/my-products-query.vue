<script lang="js">
import InnerPage from '@/components/products/my-product-query-page.vue'

export default {
  name: 'views-my-products-query',

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
    <h1>产品 : 查询</h1>
    <InnerPage> </InnerPage>
  </frame-for-user>
</template>
