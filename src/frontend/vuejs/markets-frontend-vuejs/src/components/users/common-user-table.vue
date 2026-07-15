<script lang="js">
import { ElButton, ElTableColumn } from 'element-plus';


// import { useAxiosLib } from '@/stores/axios';
// import { ElButton, ElForm, ElFormItem, ElInput } from 'element-plus';
// import { useSessionStore } from '@/stores/sessions.js';
// import { ElButton, ElButtonGroup, ElTable } from 'element-plus';
// import MiniLoginDialog from './mini-login-dialog.vue'
// const theSessionStore = useSessionStore()


export default {

  name: "sessions-common-table-view",

  components: {},

  computed: {
  },

  data() {
    return {}
  },

  methods: {
    init() { },
  },

  mounted() {
  },

  props: {
    modelValue: Array
  }
}

</script>

<style></style>

<template>
  <div>
    <ElTable :data="modelValue">

      <ElTableColumn prop="id" label="ID"></ElTableColumn>


      <ElTableColumn prop="username" label="User.Name"></ElTableColumn>
      <ElTableColumn prop="nickname" label="Nickname"></ElTableColumn>



      <ElTableColumn label="操作">
        <template #default>
          <ElButton link type="primary">Info</ElButton>
          <ElButton link type="primary">Kill</ElButton>
        </template>
      </ElTableColumn>

    </ElTable>
  </div>
</template>
