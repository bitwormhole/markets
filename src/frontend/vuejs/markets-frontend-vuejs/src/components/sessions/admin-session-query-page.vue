<script lang="js">


import MyTable from './common-session-table.vue';
import MyTablePager from './common-session-table-pager.vue';
import MyTableLoader from './admin-session-table-loader.vue';
import MyToolbar from '@/components/widgets/toolbar/index.vue'


export default {

  name: "sessions-admin-query-page",

  components: { MyTableLoader, MyTablePager, MyTable },

  data() {
    const items = []
    return { items }
  },

  methods: {
    fetch() {
      // theSessionStore.fetch()
    },

    onLoadItems(list) {
      this.items = list
    },

    handleClickAdd() {
      // let path = '/admin/permissions/add'
      // let lo = this.$router.resolve(path);
      // let url = lo.fullPath;
      window.open(url, '_blank')
    },

    handleClickRefresh() {
      this.rev2++;
    },

  },

  mounted() {
    this.fetch()
  },

  props: {}
}

</script>

<style></style>

<template>
  <div>
    <MyTableLoader v-model="items" @on-items='onLoadItems' />

    <MyToolbar>
      <!-- <ElButton @click="handleClickAdd"> Add </ElButton> -->
      <ElButton @click="handleClickRefresh"> Refresh </ElButton>
    </MyToolbar>


    <MyTablePager />
    <MyTable v-model="items" />
    <MyTablePager />
  </div>
</template>
