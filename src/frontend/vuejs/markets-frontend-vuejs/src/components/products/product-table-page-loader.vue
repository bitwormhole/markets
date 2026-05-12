<script lang="js">

import MyLoader from '@/components/widgets/table-data-loader/index.vue'
import { useProductStore } from '@/stores/products';


const theProductStore = useProductStore()


export default {

  name: "product-table-loader",

  components: { MyLoader },

  computed: {
  },

  data() {
    const items = []
    const pagination = {}
    return { items, pagination }
  },

  methods: {

    fetch(config) {
      theProductStore.fetch(config).then((res) => {
        let vo = res.data;
        this.items = vo.products;
        this.pagination = vo.pagination;
      });
    },

  },

  mounted() {


  },

  props: {}
}

</script>

<style></style>

<template>
  <div>
    <MyLoader bank="default" :items="items" :pagination='pagination' @onload="fetch" />
  </div>
</template>
