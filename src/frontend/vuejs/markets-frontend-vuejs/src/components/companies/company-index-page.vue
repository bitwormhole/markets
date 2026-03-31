<script lang="js">

import { useAxiosLib } from '@/stores/axios';
import { ElButton, ElButtonGroup, ElForm, ElFormItem, ElInput } from 'element-plus';

import MyToolbar from '@/components/widgets/toolbar/index.vue';
import usePaginationJS from '@/js/pagination';

const theAxiosLib = useAxiosLib()
const theNewPaginationParamGetter = usePaginationJS.NewPaginationParamGetter()

export default {

  name: "company-index-page",

  components: { MyToolbar },

  data() {
    const item = {
      name: '',
      ptype: '',
      category: '',
      code: '',
    }
    return { item }
  },

  methods: {

    init() {
      this.fetch();
    },

    fetch() {
      let method = 'GET'
      let url = '/api/v1/companies'
      let params = {}
      let ctx = this;

      params = theNewPaginationParamGetter.GetParams(ctx, params)

      theAxiosLib.execute({ method, url, params })
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
      <ElButton @click="fetch"> Refresh </ElButton>
      <ElButton>btn-2 </ElButton>
      <ElButton>btn-3 </ElButton>
    </MyToolbar>

    <ElForm label-width="100">

      <ElFormItem label="商品名称">
        <ElInput v-model="item.name"></ElInput>
      </ElFormItem>

      <ElFormItem label="商品类别">
        <ElInput v-model="item.category"></ElInput>
      </ElFormItem>

      <ElFormItem label="商品类型">
        <ElInput v-model="item.ptype"></ElInput>
      </ElFormItem>

      <ElFormItem label="商品条码">
        <ElInput v-model="item.code"></ElInput>
      </ElFormItem>

    </ElForm>


    <ElButton type="success" @click="save"> Save </ElButton>

  </div>
</template>
