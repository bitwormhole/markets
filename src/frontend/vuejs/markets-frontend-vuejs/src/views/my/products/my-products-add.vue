<script lang="js">
import InnerPage from '@/components/products/my-product-add-page.vue'

export default {
  name: 'views-my-products-add',

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
    <h1>产品 : 添加</h1>
    <InnerPage> </InnerPage>
  </frame-for-user>
</template>
