<script lang="js">
import { ElButton, ElTable, ElTableColumn } from 'element-plus'

import MyOverviewDialog from './common-product-overview-dialog.vue'

export default {
  name: 'product-add-page',

  components: { ElTable, ElTableColumn, ElButton, MyOverviewDialog },

  computed: {
    items() {
      return this.modelValue['items']
    },
  },

  data() {
    const displayItemOverviewDialog = false
    const theCurrentItem = {}
    return { displayItemOverviewDialog, theCurrentItem }
  },

  methods: {
    init() {},

    handleClickItemInfo(item) {
      // aka. 'overview'

      console.log('todo: show item (product), id=', item.id)

      this.theCurrentItem = item
      this.displayItemOverviewDialog = true
    },

    handleClickItemEdit(item) {
      let id = '' + item.id
      let mark = '[[id]]'
      let url = '/my/products/[[id]]/edit'
      let location = url.replace(mark, id)
      console.log('todo: edit item (product), location=', location)
      window.open(location, '_blank')
    },

    handleClickItemRemove(item) {
      console.log('todo: remove item (product), id=', item.id)
    },
  },

  mounted() {},

  props: {
    modelValue: Object, // a VMO
  },
}
</script>

<style></style>

<template>
  <div>
    <ElTable :data="items">
      <ElTableColumn prop="id" label="ID"></ElTableColumn>
      <ElTableColumn prop="name" label="Name"></ElTableColumn>
      <ElTableColumn prop="code" label="Code"></ElTableColumn>

      <ElTableColumn label="操作">
        <template #default="scope">
          <ElButton @click="handleClickItemInfo(scope.row)"> Info </ElButton>
          <ElButton @click="handleClickItemEdit(scope.row)"> Edit </ElButton>
          <ElButton type="danger" @click="handleClickItemRemove(scope.row)"> Delete </ElButton>
        </template>
      </ElTableColumn>
    </ElTable>

    <MyOverviewDialog v-model="displayItemOverviewDialog" :item="theCurrentItem">
    </MyOverviewDialog>
  </div>
</template>
