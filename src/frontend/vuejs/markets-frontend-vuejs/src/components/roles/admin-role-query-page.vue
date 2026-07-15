<script lang="js">
import MyLoader from './admin-role-table-loader.vue'
import MyTable from './common-role-table.vue'
import MyToolbar from '@/components/widgets/toolbar/index.vue'
import MyPager from '@/components/widgets/pager/index.vue'

export default {
  name: 'example-query-page',

  components: { MyLoader, MyTable, MyToolbar, MyPager },

  data() {
    const rev2 = 0
    const items = []
    const pagination = {}
    return { rev2, items, pagination }
  },

  methods: {
    init() {},

    handleClickRefresh() {
      this.rev2++
    },

    handleClickAdd() {
      let path = '/admin/roles/add'
      let lo = this.$router.resolve(path)
      let url = lo.fullPath
      window.open(url, '_blank')
    },

    on_load_items(list) {
      this.items = list
    },

    on_load_pagination(pagination) {
      this.pagination = pagination
    },

    reload() {
      this.rev2++
    },
  },

  mounted() {
    this.init()
  },

  props: {},
}
</script>

<style></style>

<template>
  <div>
    <MyToolbar>
      <ElButton @click="handleClickAdd"> Add </ElButton>
      <ElButton @click="handleClickRefresh"> Refresh </ElButton>
    </MyToolbar>

    <MyPager :revision="rev2" :pagination="pagination" @reload="reload"> </MyPager>
    <MyTable :revision="rev2" :pagination="pagination" v-model="items"></MyTable>
    <MyPager :revision="rev2" :pagination="pagination" @reload="reload"> </MyPager>

    <MyLoader
      v-model="items"
      :revision="rev2"
      @on-items="on_load_items"
      @on-page="on_load_pagination"
    ></MyLoader>
  </div>
</template>
