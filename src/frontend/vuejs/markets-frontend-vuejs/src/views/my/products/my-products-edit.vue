<script lang="js">
import InnerPage from '@/components/products/my-product-edit-page.vue'

export default {
  name: 'views-my-products-edit',

  components: { InnerPage },

  data() {
    const p_name = ''
    const p_code = ''
    return { p_name, p_code }
  },

  methods: {
    on_load_item(item) {
      this.p_name = item.name
      this.p_code = item.code
    },
  },

  mounted() {},
}
</script>

<style scoped></style>

<template>
  <frame-for-user>
    <h1>编辑产品(条码:{{ p_code }}) : {{ p_name }}</h1>
    <InnerPage @on-load-item="on_load_item"> </InnerPage>
  </frame-for-user>
</template>
