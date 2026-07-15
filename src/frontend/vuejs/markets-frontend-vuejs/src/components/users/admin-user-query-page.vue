<script lang="js">

import { useAxiosLib } from '@/stores/axios';
import { ElButton, ElForm, ElFormItem, ElInput } from 'element-plus';

import MyLoader from './admin-user-table-loader.vue'
import MyTable from './common-user-table.vue'
import MyToolbar from '@/components/widgets/toolbar/index.vue'
import MyRefreshLoader from '@/components/widgets/table-data-loader/refresh-loader.vue'


const theAxiosLib = useAxiosLib()


export default {

  name: "user-query-page",

  components: { MyLoader, MyTable, MyToolbar, MyRefreshLoader },

  data() {
    const rev2 = 0;
    const items = []
    return { rev2, items }
  },

  methods: {
    init() { },

    handleClickRefresh() {
      this.rev2++;
    },

    handleClickAdd() {
      let path = '/admin/users/add'
      let lo = this.$router.resolve(path);
      let url = lo.fullPath;
      window.open(url, '_blank')
    },

    onLoadItems(list) {
      this.items = list;
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

    <MyToolbar>
      <ElButton @click="handleClickAdd"> Add </ElButton>
      <ElButton @click="handleClickRefresh"> Refresh </ElButton>
    </MyToolbar>

    <!-- <MyDebug> </MyDebug> -->

    <MyTable v-model="items"></MyTable>

    <MyLoader @on-items="onLoadItems"></MyLoader>

    <MyRefreshLoader :revision="rev2"> </MyRefreshLoader>

  </div>
</template>
