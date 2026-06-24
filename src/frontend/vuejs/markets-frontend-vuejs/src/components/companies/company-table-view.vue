<script lang="js">

import { ElButton, ElForm, ElFormItem, ElInput, ElTable, ElTableColumn } from 'element-plus';

import { useCompanyStore } from '@/stores/companies';
import MyDeleteItemDialog from './company-delete-dialog.vue'
import MyItemPropsDialog from './company-properties-dialog.vue'

const theProductStore = useCompanyStore()


export default {

  name: "company-table-view",

  components: { MyDeleteItemDialog, MyItemPropsDialog },

  computed: {
    items() {
      return theProductStore.items;
    },
  },

  data() {
    const current_selection_item = {}
    const display_delete_item_dialog = false;
    const display_item_props_dialog = false;
    return { current_selection_item, display_delete_item_dialog, display_item_props_dialog }
  },

  methods: {
    init() { },

    handleClickItemDelete(item) {
      this.current_selection_item = item;
      this.display_delete_item_dialog = true;
    },

    handleClickItemProperties(item) {
      this.current_selection_item = item;
      this.display_item_props_dialog = true;
    },

    handleClickItemEdit(item) {
      // this.current_selection_item = item;
      // this.display_item_props_dialog = true;

      let id = item.id;
      let url = '/admin/companies/' + id + '/edit';

      window.open(url, '_blank')
    },

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

      <ElTableColumn label="企业名称" prop="name"></ElTableColumn>
      <ElTableColumn label="统一社会信用代码" prop="code"></ElTableColumn>

      <ElTableColumn label="操作">
        <template #default="scope">
          <ElButton link type="primary" @click="handleClickItemProperties(scope.row)"> 属性 </ElButton>
          <ElButton link type="primary" @click="handleClickItemEdit(scope.row)"> 编辑 </ElButton>
          <ElButton link type="danger" @click="handleClickItemDelete(scope.row)"> 删除 </ElButton>
        </template>
      </ElTableColumn>
    </ElTable>

    <MyDeleteItemDialog v-model="display_delete_item_dialog" :item="current_selection_item" />

    <MyItemPropsDialog v-model="display_item_props_dialog" :item="current_selection_item"> </MyItemPropsDialog>

  </div>
</template>
