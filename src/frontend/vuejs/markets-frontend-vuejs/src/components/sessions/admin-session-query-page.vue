<script lang="js">
// sessions-admin-query-page.vue

import MyLoader from '@/components/widgets/table-data-loader/main-table-loader.vue'
import MyToolbar from '@/components/widgets/toolbar/index.vue'
import MyPager from '@/components/widgets/pager/index.vue'
import MyDebugBox from '@/components/widgets/debug-box/index.vue'

import MyTable from './common-session-table.vue'

export default {
  name: 'sessions-admin-query-page',

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
      pagination: { page: 1, size: 5, total: 0 },
    }
    return { vmo }
  },

  methods: {
    init() {},

    reload() {
      this.vmo.revision++
    },

    handleClickDebugBox() {},

    handleClickRefresh() {
      this.reload()
    },

    handleClickAdd() {
      let path = '/admin/sessions/add'
      let lo = this.$router.resolve(path)
      let url = lo.fullPath
      window.open(url, '_blank')
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

    <MyLoader v-model="vmo" auto path="/api/v1/admin/sessions" items="sessions"></MyLoader>

    <MyPager v-model="vmo" @reload="reload"> </MyPager>
    <MyTable v-model="vmo"></MyTable>
    <MyPager v-model="vmo" @reload="reload"> </MyPager>

    <MyDebugBox @click="handleClickDebugBox">
      <pre> {{ theDebugText }}  </pre>
    </MyDebugBox>
  </div>
</template>
