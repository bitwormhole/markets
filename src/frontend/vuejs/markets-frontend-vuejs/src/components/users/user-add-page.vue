<script lang="js">
import { useAxiosLib } from '@/stores/axios'
import { ElButton, ElForm, ElFormItem, ElInput, ElSwitch } from 'element-plus'

const theAxiosLib = useAxiosLib()

export default {
  name: 'example-add-page',

  components: {},

  data() {
    const item = {
      username: '',
      email: '',
      mobile: '',
      avatar: '',
      nickname: '',
      roles_as_str: '',
    }
    return { item }
  },

  methods: {
    init() {},

    parseRoleNameList(str) {
      let dst = []
      if (str == null) {
        return dst
      }
      let str2 = str + ''
      let src = str2.split(',')
      for (let i in src) {
        let item = src[i]
        item = item.trim()
        if (item == '') {
          continue
        }
        dst.push(item)
      }
      return dst
    },

    save() {
      let item = this.item
      let method = 'POST'
      let url = '/api/v1/admin/users'
      let data = {
        users: [item],
      }
      item.roles = this.parseRoleNameList(item.roles_as_str)

      theAxiosLib.execute({ method, url, data })
    },
  },

  mounted() {
    this.init()
  },

  props: {},
}
</script>

<style></style>

<template>
  <div>
    <ElForm label-width="100">
      <ElFormItem label="用户名">
        <ElInput v-model="item.username"></ElInput>
      </ElFormItem>

      <ElFormItem label="昵称">
        <ElInput v-model="item.nickname"></ElInput>
      </ElFormItem>

      <ElFormItem label="角色">
        <ElInput v-model="item.roles_as_str"></ElInput>
      </ElFormItem>

      <!-- <ElFormItem label="头像">
        <ElInput v-model="item.avatar"></ElInput>
      </ElFormItem> -->

      <ElFormItem label="手机号码">
        <ElInput v-model="item.mobile"></ElInput>
      </ElFormItem>

      <ElFormItem label="Email 地址">
        <ElInput v-model="item.email"></ElInput>
      </ElFormItem>

      <ElFormItem label="Enabled">
        <ElSwitch v-model="item.enabled"></ElSwitch>
      </ElFormItem>
    </ElForm>

    <ElButton type="success" @click="save"> 确定 </ElButton>
  </div>
</template>
