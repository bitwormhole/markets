<script lang="js">

import { ElButton } from 'element-plus';

import MyLoader from './product-table-page-loader.vue'
import MyTable from './product-table-view.vue'
import MyToolbar from '@/components/widgets/toolbar/index.vue'
import MyRefreshLoader from '@/components/widgets/table-data-loader/refresh-loader.vue'


export default {

  name: "product-query-page",

  components: { MyLoader, MyTable, MyToolbar, MyRefreshLoader },

  data() {
    const rev2 = 0;
    return { rev2 }
  },

  methods: {
    init() { },

    handleClickRefresh() {
      this.rev2++;
    },

    handleClickAdd() {
      let path = '/products/add'
      let lo = this.$router.resolve(path);
      let url = lo.fullPath;
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

    <MyToolbar>
      <ElButton @click="handleClickAdd"> Add </ElButton>
      <ElButton @click="handleClickRefresh"> Refresh </ElButton>
    </MyToolbar>

    <!-- <MyDebug> </MyDebug> -->

    <MyTable></MyTable>

    <MyLoader></MyLoader>

    <MyRefreshLoader :revision="rev2"> </MyRefreshLoader>

  </div>
</template>
