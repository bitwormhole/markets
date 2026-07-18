<script lang="js">
import MyLoader from '@/components/widgets/table-data-loader/main-table-loader.vue'
import MyToolbar from '@/components/widgets/toolbar/index.vue'
import MyPager from '@/components/widgets/pager/index.vue'
import MyDebugBox from '@/components/widgets/debug-box/index.vue'

import MyTable from './common-product-table.vue'

export default {
  name: 'products-my-query-page',

  components: { MyLoader, MyToolbar, MyPager, MyDebugBox, MyTable },

  data() {
    const vmo = {
      revision: 0,
      items: [],
      pagination: { page: 1, size: 5, total: 0 },
    }
    return { vmo }
  },

  methods: {
    handleClickAdd() {
      let url = '/my/products/do/add'
      window.open(url, '_blank')
    },

    handleClickRefresh() {
      this.reload()
    },

    reload() {
      this.vmo.revision++
    },
  },

  mounted() {},

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

    <MyLoader v-model="vmo" auto path="/api/v1/my/products" items="products"></MyLoader>

    <MyPager v-model="vmo" @reload="reload"> </MyPager>
    <MyTable v-model="vmo"></MyTable>
    <MyPager v-model="vmo" @reload="reload"> </MyPager>

    <MyDebugBox @click="handleClickDebugBox">
      <pre> {{ theDebugText }}  </pre>
    </MyDebugBox>
  </div>
</template>
