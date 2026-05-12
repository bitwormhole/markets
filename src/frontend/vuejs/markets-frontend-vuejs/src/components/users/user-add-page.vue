<script lang="js">

import { useAxiosLib } from '@/stores/axios';
import { ElButton, ElForm, ElFormItem, ElInput } from 'element-plus';


const theAxiosLib = useAxiosLib()


export default {

  name: "example-add-page",

  components: {},

  data() {
    const item = {
      username: '',
      email: '',
      mobile: '',
      avatar: '',
      nickname: '',
    }
    return { item }
  },

  methods: {
    init() { },


    save() {

      let item = this.item;

      let method = 'POST'
      let url = '/api/v1.admin/users'
      let data = {
        users: [item]
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

      <ElFormItem label="用户名">
        <ElInput v-model="item.username"></ElInput>
      </ElFormItem>

      <ElFormItem label="手机号码">
        <ElInput v-model="item.mobile"></ElInput>
      </ElFormItem>

      <ElFormItem label="Email 地址">
        <ElInput v-model="item.email"></ElInput>
      </ElFormItem>

      <ElFormItem label="头像">
        <ElInput v-model="item.avatar"></ElInput>
      </ElFormItem>

      <ElFormItem label="昵称">
        <ElInput v-model="item.nickname"></ElInput>
      </ElFormItem>

    </ElForm>


    <ElButton type="success" @click="save"> Save </ElButton>

  </div>
</template>
