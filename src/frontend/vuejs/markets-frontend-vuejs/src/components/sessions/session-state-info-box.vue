<script lang="js">

// import { useAxiosLib } from '@/stores/axios';
// import { ElButton, ElForm, ElFormItem, ElInput } from 'element-plus';

import { useSessionStore } from '@/stores/sessions.js';
import { ElButton, ElButtonGroup } from 'element-plus';
import MiniLoginDialog from './mini-login-dialog.vue'

const theSessionStore = useSessionStore()


export default {

  name: "session-state-info-box",

  components: { MiniLoginDialog },

  computed: {
    theAvatar() {
      return this.getFieldValue('avatar');
    },

    theNickName() {
      return this.getFieldValue('nickname');
    },

    theIsAuth() {
      let ok = this.getFieldValue('authenticated');
      return (ok ? true : false);
    },

    theEmail() {
      return this.getFieldValue('email');
    },

    theUserName() { },

    theUserID() { },

    theSessionInfo() {



    },

  },

  data() {
    const displayLoginDialog = false
    return {
      displayLoginDialog
    }
  },

  methods: {

    getFieldValue(name) {
      let info = theSessionStore.current;
      if (info == null) {
        return '[null]'
      }
      return info[name]
    },

    fetch() {
      theSessionStore.fetch()
    },

    handleClickLogin() {
      this.displayLoginDialog = true
    },

    handleClickSignUp() { },

    handleClickExit() {
      let p = theSessionStore.exit()
      p.then(() => {
        this.fetch()
      })
    },

    on_login_fail() { },

    on_login_ok() {
      this.displayLoginDialog = false
      this.fetch();
    },

  },

  mounted() {
    this.fetch()
  },

  props: {}
}

</script>

<style></style>

<template>
  <div>

    <div v-show="!theIsAuth">
      <ElButton link type="primary" @click="handleClickLogin">登录</ElButton>
      <ElButton link type="primary" @click="handleClickSignUp">注册</ElButton>
    </div>

    <div v-show="theIsAuth">
      <el-popover :width="300"
        popper-style="box-shadow: rgb(14 18 22 / 35%) 0px 10px 38px -10px, rgb(14 18 22 / 20%) 0px 10px 20px -15px; padding: 20px;">
        <template #reference>
          <el-avatar :src="theAvatar" />
        </template>
        <template #default>
          <div class="demo-rich-conent" style="display: flex; gap: 16px; flex-direction: column">


            <el-avatar :size="60" :src="theAvatar" style="margin-bottom: 8px" />

            <div>
              <span> {{ theNickName }} </span>
            </div>
            <div>
              <span> {{ theEmail }} </span>
            </div>


            <ElButton> 设置 </ElButton>
            <ElButton @click="handleClickExit"> 退出 </ElButton>
          </div>
        </template>
      </el-popover>
    </div>


    <MiniLoginDialog v-model="displayLoginDialog" @ok="on_login_ok" @fail="on_login_fail" />

  </div>
</template>
