<script lang="js">

import { ElButton, ElForm, ElFormItem, ElInput, ElTable, ElTableColumn } from 'element-plus';

import { useLicenceStore } from '@/stores/licences';


const theProductStore = useLicenceStore()


export default {

  name: "licence-table-view",

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
