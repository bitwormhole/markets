<script lang="js">
// this file : 'my-product-edit-page.vue'

import MyEditAll from './common-product-edit-all.vue'

export default {
  name: 'products-my-edit-page',

  components: { MyEditAll },

  data() {
    const item = {}
    return { item }
  },

  methods: {
    init() {},

    on_load_item(item) {
      this.$emit('on-load-item', item)
    },
  },

  mounted() {
    this.init()
  },

  props: {},
}
</script>

<style></style>

<template>
  <div><MyEditAll @on-load-item="on_load_item"> </MyEditAll></div>
</template>
