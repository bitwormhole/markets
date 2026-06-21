<script lang="js">

import { useProductStore } from '@/stores/products';

const theProductStore = useProductStore()

export default {

  name: "product-debug-view",

  components: {},

  computed: {
    output() {
      let params = theProductStore.params;
      return JSON.stringify(params);
    },
  },

  data() {
    return {}
  },

  methods: {
  },

  mounted() {
  },

  props: {}
}

</script>

<style>
.debug-box {
  color: gray;

  background-color: yellow;

  border-color: black;
  border-width: 3px;
  border-style: solid;

  margin: 3px;
  padding: 3px;
}
</style>

<template>
  <div class="debug-box">

    {{ output }}

  </div>
</template>
