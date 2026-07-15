<script lang="js">


import MyLoader from './admin-permission-table-loader.vue'
import MyTable from './common-permission-table.vue'
import MyToolbar from '@/components/widgets/toolbar/index.vue'
import MyRefreshLoader from '@/components/widgets/table-data-loader/refresh-loader.vue'


export default {

  name: "permission-query-page",

  components: { MyLoader, MyTable, MyToolbar, MyRefreshLoader },

  data() {
    const items = [];
    return { items }
  },

  methods: {
    init() { },

    handleClickRefresh() {
      this.rev2++;
    },

    handleClickAdd() {
      let path = '/admin/permissions/add'
      let lo = this.$router.resolve(path);
      let url = lo.fullPath;
      window.open(url, '_blank')
    },


    on_load_items(list) {
      this.items = list
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

    <MyLoader v-model="items" @on-items="on_load_items"></MyLoader>

    <MyRefreshLoader :revision="rev2"> </MyRefreshLoader>

  </div>
</template>
