<script lang="js">

// import { useAxiosLib } from '@/stores/axios';
import { ElButton, ElForm, ElFormItem, ElInput, ElTable, ElTableColumn } from 'element-plus';

import { useShopStore } from '@/stores/shops';


const theProductStore = useShopStore()


export default {

  name: "shop-table-view",

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

    handleClickItemProps(item) { },

    handleClickItemEdit(item) { },

    handleClickItemDelete(item) { },

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

      <ElTableColumn label="Label" prop="id">
        <template #default="scope">
          <ElButton link type="primary">{{ scope.row.code }}</ElButton>
        </template>
      </ElTableColumn>

      <ElTableColumn label="Name" prop="name"></ElTableColumn>
      <ElTableColumn label="Code" prop="code"></ElTableColumn>

      <ElTableColumn label="操作">
        <template #default="scope">
          <ElButton link type="primary" @click="handleClickItemProperties(scope.row)"> 属性 </ElButton>
          <ElButton link type="primary" @click="handleClickItemEdit(scope.row)"> 编辑 </ElButton>
          <ElButton link type="danger" @click="handleClickItemDelete(scope.row)"> 删除 </ElButton>
        </template>
      </ElTableColumn>

    </ElTable>

  </div>
</template>
