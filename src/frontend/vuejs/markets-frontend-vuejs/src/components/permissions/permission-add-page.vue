<script lang="js">

import { useAxiosLib } from '@/stores/axios';
import { ElButton, ElCheckbox, ElForm, ElFormItem, ElInput, ElSwitch } from 'element-plus';


const theAxiosLib = useAxiosLib()


export default {

  name: "example-add-page",

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
      let url = '/api/v1/admin/permissions'

      let data = {
        permissions: [item]
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

      <ElFormItem label="Method">
        <ElInput v-model="item.method"></ElInput>
      </ElFormItem>

      <ElFormItem label="Path">
        <ElInput v-model="item.path"></ElInput>
      </ElFormItem>

      <ElFormItem label="Roles">
        <ElInput v-model="item.roles"></ElInput>
      </ElFormItem>

      <ElFormItem label="Enabled">
        <ElSwitch v-model="item.enabled"></ElSwitch>
      </ElFormItem>

    </ElForm>

    <ElButton type="success" @click="save"> 确定 </ElButton>

  </div>
</template>
