<script lang="js">

import lib_licence_js from '@/js/licences.js';


const the_licence_types = lib_licence_js.GetLicenceTypes();


export default {

  name: "licence-type-selector",

  components: {},


  computed: {
    options() {
      return the_licence_types;
    },
  },

  data() {
    const item = {}
    return { item }
  },

  methods: {

    // init() { },
    // save() { },

  },

  mounted() { },

  props: {}
}

</script>

<style></style>

<template>

  <el-select placeholder="请选择">
    <el-option v-for="item in options" :key="item.name" :label="item.label" :value="item.name" />
  </el-select>

</template>
