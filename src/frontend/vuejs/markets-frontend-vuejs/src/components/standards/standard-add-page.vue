<script lang="js">

import { useAxiosLib } from '@/stores/axios';
import { ElButton, ElForm, ElFormItem, ElInput } from 'element-plus';


const theAxiosLib = useAxiosLib()


export default {

  name: "standard-add-page",

  components: {},

  data() {
    const item = {
      title: '标准-demo-1',
      uri: '',
      ref: 'https://example.com/std/1',
      code: 'std-demo-1',
    }
    return { item }
  },

  methods: {
    init() { },


    save() {

      let item = this.item;

      let method = 'POST'
      let url = '/api/v1/standards'
      let data = {
        standards: [item]
      }
      theAxiosLib.execute({ method, url, data })
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

    <ElForm label-width="100">

      <ElFormItem label="标准代码">
        <ElInput v-model="item.code"></ElInput>
      </ElFormItem>

      <ElFormItem label="标题">
        <ElInput v-model="item.title"></ElInput>
      </ElFormItem>

      <ElFormItem label="描述">
        <ElInput v-model="item.description" type="textarea"></ElInput>
      </ElFormItem>


      <!-- <ElFormItem label="URI">
        <ElInput v-model="item.uri"></ElInput>
      </ElFormItem> -->

      <ElFormItem label="参考网址">
        <ElInput v-model="item.ref" type="textarea"></ElInput>
      </ElFormItem>

    </ElForm>


    <ElButton type="success" @click="save"> Save </ElButton>

  </div>
</template>
