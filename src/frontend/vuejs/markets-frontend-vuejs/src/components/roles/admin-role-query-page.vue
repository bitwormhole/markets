<script lang="js">
// admin-role-query-page.vue

import MyLoader from '@/components/widgets/table-data-loader/index.vue'
import MyToolbar from '@/components/widgets/toolbar/index.vue'
import MyPager from '@/components/widgets/pager/index.vue'
import MyDebugBox from '@/components/widgets/debug-box/index.vue'

import MyTable from './common-role-table.vue'

export default {
  name: 'roles-admin-query-page',

  components: { MyLoader, MyTable, MyToolbar, MyPager, MyDebugBox },

  computed: {
    theDebugText() {
      let vmo = this.vmo
      return JSON.stringify(vmo, null, '\t')
    },
  },

  data() {
    const vmo = {
      revision: 0,
      items: [],
      pagination: {
        page: 1,
        size: 5,
        total: 0,
      },
    }
    return { vmo }
  },

  methods: {
    init() {},

    handleClickRefresh() {
      this.reload()
    },

    handleClickAdd() {
      let path = '/admin/roles/add'
      let lo = this.$router.resolve(path)
      let url = lo.fullPath
      window.open(url, '_blank')
    },

    reload() {
      this.vmo.revision++
    },

    handleClickDebugBox() {},
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

    <MyLoader v-model="vmo" auto path="/api/v1/admin/roles" items="roles"></MyLoader>

    <MyPager v-model="vmo" @reload="reload"> </MyPager>
    <MyTable v-model="vmo"></MyTable>
    <MyPager v-model="vmo" @reload="reload"> </MyPager>

    <MyDebugBox @click="handleClickDebugBox">
      <pre> {{ theDebugText }}  </pre>
    </MyDebugBox>
  </div>
</template>
