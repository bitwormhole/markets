<script lang="js">

import { ElButton, ElForm, ElFormItem, ElInput, ElTable, ElTableColumn } from 'element-plus';

import { useUserStore } from '@/stores/users';


const theUserStore = useUserStore()


export default {

  name: "user-table-view",

  components: {},

  computed: {
    items() {
      return theUserStore.items;
    },
  },

  data() {
    const item = {}
    return { item }
  },

  methods: {

    handleClickItemProps(item) { },

    handleClickItemEdit(item) { },

    handleClickItemDelete(item) { },

  },

  mounted() {
    // this.init()
  },

  props: {}
}

</script>

<style></style>

<template>
  <div>

    <ElTable :data="items">
      <ElTableColumn label="ID" prop="id"></ElTableColumn>
      <ElTableColumn label="用户名" prop="username"></ElTableColumn>
      <ElTableColumn label="昵称" prop="nickname"></ElTableColumn>
      <ElTableColumn label="手机号码" prop="mobile"></ElTableColumn>
      <ElTableColumn label="邮箱地址" prop="email"></ElTableColumn>
      <ElTableColumn label="角色" prop="roles"></ElTableColumn>

      <ElTableColumn label="操作">
        <template #default="scope">
          <ElButton type="primary" link @click="handleClickItemProps(scope.row)"> 属性 </ElButton>
          <ElButton type="primary" link @click="handleClickItemEdit(scope.row)"> 修改 </ElButton>
          <ElButton type="danger" link @click="handleClickItemDelete(scope.row)"> 删除 </ElButton>
        </template>
      </ElTableColumn>
    </ElTable>

  </div>
</template>
