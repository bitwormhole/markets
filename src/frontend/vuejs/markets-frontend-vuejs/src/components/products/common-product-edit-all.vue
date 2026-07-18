<script lang="js">
import { useAxiosLib } from '@/stores/axios'
import LibGetterJS from '@/js/getter'

import MySubViewCore from './common-product-edit-core.vue'
import MySubViewImageList from './common-product-edit-images.vue'
import MySubViewFactories from './common-product-edit-factories.vue'
import { ElButton } from 'element-plus'

const theAxiosLib = useAxiosLib()

export default {
  name: 'products-common-edit-all',

  components: { MySubViewCore, MySubViewImageList, MySubViewFactories },

  computed: {
    theProductID() {
      return this.$route.params['id']
    },
  },

  data() {
    const item = {}
    return { item }
  },

  methods: {
    init() {},

    fetch() {
      let pid = this.theProductID
      let url = '/api/v1/my/products/' + pid
      let method = 'GET'
      let p = theAxiosLib.execute({ method, url })
      p.then((res) => {
        let vo = res.data
        let def = {}
        let it = LibGetterJS.GetFirstItem(vo.products, def)
        this.item = it
        this.fireOnLoadItem(it)
      })
      return p
    },

    save() {
      let pid = this.theProductID
      let url = '/api/v1/my/products/' + pid
      let method = 'PUT'
      let item = this.item
      let data = {
        products: [item],
      }
      let p = theAxiosLib.execute({ method, url, data })
      p.then((res) => {
        console.log('save item OK, http.status=', res.status)
      })
      return p
    },

    fireOnLoadItem(item) {
      this.$emit('on-load-item', item)
    },
  },

  mounted() {
    this.fetch()
  },

  props: {
    // item: Object, // a product-dto
  },
}
</script>

<style scoped></style>

<template>
  <div>
    <el-tabs tab-position="left" class="demo-tabs">
      <el-tab-pane label="基本信息">
        <template #default> <MySubViewCore :item="item"> </MySubViewCore> </template>
      </el-tab-pane>

      <el-tab-pane label="生产厂商"> <MySubViewFactories> </MySubViewFactories> </el-tab-pane>

      <el-tab-pane label="产品图片"> <MySubViewImageList> </MySubViewImageList> </el-tab-pane>

      <!--  以下是食品特有的属性  -->

      <el-tab-pane label="配料表">Task</el-tab-pane>
      <el-tab-pane label="营养成分表">Task</el-tab-pane>
    </el-tabs>

    <div>
      <ElButton type="success" @click="save"> 保存 </ElButton>
    </div>
  </div>
</template>
