<script lang="js">

import { useAxiosLib } from '@/stores/axios';
import { ElButton, ElForm, ElFormItem, ElInput } from 'element-plus';


const theAxiosLib = useAxiosLib()


export default {

  name: "standard-add-page",

  components: {},

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

      <ElFormItem label="标准名称">
        <ElInput v-model="item.name"></ElInput>
      </ElFormItem>

      <ElFormItem label="标准类别">
        <ElInput v-model="item.category"></ElInput>
      </ElFormItem>

      <ElFormItem label="标准类型">
        <ElInput v-model="item.ptype"></ElInput>
      </ElFormItem>

      <ElFormItem label="标准条码">
        <ElInput v-model="item.code"></ElInput>
      </ElFormItem>

    </ElForm>


    <ElButton type="success" @click="save"> Save </ElButton>

  </div>
</template>
