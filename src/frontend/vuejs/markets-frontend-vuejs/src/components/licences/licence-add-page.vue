<script lang="js">

import { useAxiosLib } from '@/stores/axios';
import { ElButton, ElForm, ElFormItem, ElInput } from 'element-plus';


const theAxiosLib = useAxiosLib()


export default {

  name: "licence-add-page",

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
      let url = '/api/v1/licences'
      let data = {
        licences: [item]
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

      <ElFormItem label="许可证名称">
        <ElInput v-model="item.name"></ElInput>
      </ElFormItem>

      <ElFormItem label="许可证类别">
        <ElInput v-model="item.category"></ElInput>
      </ElFormItem>

      <ElFormItem label="许可证类型">
        <ElInput v-model="item.ptype"></ElInput>
      </ElFormItem>

      <ElFormItem label="许可证条码">
        <ElInput v-model="item.code"></ElInput>
      </ElFormItem>

    </ElForm>


    <ElButton type="success" @click="save"> Save </ElButton>

  </div>
</template>
