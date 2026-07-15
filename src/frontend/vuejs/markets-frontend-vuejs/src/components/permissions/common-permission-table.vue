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

      <ElTableColumn prop="id" label="ID" width="100"></ElTableColumn>

      <ElTableColumn prop="method" label="Method"></ElTableColumn>

      <ElTableColumn prop="path" label="Path"></ElTableColumn>

      <ElTableColumn prop="enabled" label="Enabled" width="100"></ElTableColumn>

      <ElTableColumn prop="roles" label="Roles"></ElTableColumn>

      <ElTableColumn label="操作" width="200">
        <template #default>
          <ElButton link type="primary">Info</ElButton>
          <ElButton link type="primary">Kill</ElButton>
        </template>
      </ElTableColumn>

    </ElTable>
  </div>
</template>
