<script lang="js">

// import { useAxiosLib } from '@/stores/axios';
import { ElButton, ElForm, ElFormItem, ElInput, ElTable, ElTableColumn } from 'element-plus';

import { useTrademarkStore } from '@/stores/trademarks';


const theProductStore = useTrademarkStore()


export default {

  name: "trademark-table-view",

  components: {},

  computed: {
    items() {
      return theProductStore.items;
    },
  },

  data() {
    const item = {}
    return { item }
  },

  methods: {
    init() { },

  },

  mounted() {
    this.init()
  },

  props: {}
}

</script>

<style></style>

<template>
  <div>

    <ElTable :data="items">
      <ElTableColumn label="ID" prop="id" :width="80"></ElTableColumn>
      <ElTableColumn label="Name" prop="name"></ElTableColumn>
      <ElTableColumn label="Code" prop="code"></ElTableColumn>
    </ElTable>

  </div>
</template>
