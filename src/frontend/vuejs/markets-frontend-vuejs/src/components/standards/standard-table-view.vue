<script lang="js">

// import { useAxiosLib } from '@/stores/axios';
import { ElButton, ElButtonGroup, ElForm, ElFormItem, ElInput, ElTable, ElTableColumn } from 'element-plus';

import { useStandardStore } from '@/stores/standards';


const theProductStore = useStandardStore()


export default {

  name: "standard-table-view",

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

      <ElTableColumn label="标准编号" prop="code">
        <template #default="scope">
          <ElButton link type="primary">{{ scope.row.code }}</ElButton>
        </template>
      </ElTableColumn>

      <!-- <ElTableColumn label="标准代号" prop="code"></ElTableColumn> -->

      <ElTableColumn label="标准名称" prop="title"></ElTableColumn>
      <ElTableColumn label="上传者" prop="owner"></ElTableColumn>

      <ElTableColumn label="操作">
        <template #default="scope">
          <ElButton link> 详情 </ElButton>
        </template>
      </ElTableColumn>

    </ElTable>

  </div>
</template>
