<script lang="js">

import { useAxiosLib } from '@/stores/axios';
import { ElButton, ElForm, ElFormItem, ElInput } from 'element-plus';

import MyToolbar from '@/components/widgets/toolbar/index.vue';

const theAxiosLib = useAxiosLib()


export default {

  name: "example-index-page",

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
      let url = '/api/v1/examples'
      let params = {
        limit: 23,
        offset: 45,
      }
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
    </MyToolbar>

    <ElForm label-width="100">

      <ElFormItem label="商品名称">
        <ElInput v-model="item.name"></ElInput>
      </ElFormItem>

    </ElForm>

  </div>
</template>
