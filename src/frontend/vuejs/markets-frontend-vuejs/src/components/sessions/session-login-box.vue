<script lang="js">

// import { useAxiosLib } from '@/stores/axios';
// import { ElButton, ElForm, ElFormItem, ElInput } from 'element-plus';

import { useSessionStore } from '@/stores/sessions.js';
import { ElButton, ElInput } from 'element-plus';


const theSessionStore = useSessionStore()


export default {

  name: "session-login-box",

  components: {},

  data() {
    const username = 'foo'
    const password = 'bar'
    return { username, password }
  },

  methods: {
    handleClickLogin() {

      let username = this.username;
      let password = this.password;
      // let data = { username, password }

      let p = theSessionStore.login_basic({ password, username })

      p.then(() => {
        this.fireOk()
      }).catch(() => {
        this.fireFail()
      })

    },

    fireOk() {
      this.$emit('ok')
    },

    fireFail() {
      this.$emit('fail')
    },

  },

  mounted() {
    // this.fetch()
  },

  props: {}
}

</script>

<style></style>

<template>
  <div>

    <ElInput v-model="username" placeholder="请输入用户名"></ElInput>

    <ElInput v-model="password" type="password" placeholder="请输入密码"></ElInput>

    <ElButton type="primary" @click="handleClickLogin">登录</ElButton>

  </div>
</template>
