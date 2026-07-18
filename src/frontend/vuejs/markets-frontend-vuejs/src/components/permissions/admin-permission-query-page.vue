<script lang="js">
// admin-permissions-query-page.vue

import { useAxiosLib } from '@/stores/axios.js'

import MyLoader from '@/components/widgets/table-data-loader/main-table-loader.vue'
import MyToolbar from '@/components/widgets/toolbar/index.vue'
import MyPager from '@/components/widgets/pager/index.vue'
import MyDebugBox from '@/components/widgets/debug-box/index.vue'

import MyTable from './common-permission-table.vue'

const theAxiosLib = useAxiosLib()

export default {
  name: 'permissions-admin-query-page',

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
      let path = '/admin/permissions/add'
      let lo = this.$router.resolve(path)
      let url = lo.fullPath
      window.open(url, '_blank')
    },

    handleClickReload() {
      let method = 'POST'
      let url = '/api/v1/admin/permissions/do/reload'
      let data = {}
      theAxiosLib.execute({ method, url, data })
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
      <ElButton @click="handleClickReload"> Reload RBAC.permissions.cache </ElButton>
    </MyToolbar>

    <MyLoader v-model="vmo" auto path="/api/v1/admin/permissions" items="permissions"></MyLoader>

    <MyPager v-model="vmo" @reload="reload"> </MyPager>
    <MyTable v-model="vmo"></MyTable>
    <MyPager v-model="vmo" @reload="reload"> </MyPager>

    <MyDebugBox title="Debug: VMO">
      <pre> {{ theDebugText }}  </pre>
    </MyDebugBox>
  </div>
</template>
