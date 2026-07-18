<script lang="js">
// common-user-table

import { ElButton, ElTable, ElTableColumn } from 'element-plus'

export default {
  name: 'users-common-table',

  components: { ElButton, ElTable, ElTableColumn },

  computed: {
    items() {
      return this.modelValue['items']
    },
  },

  data() {
    return {}
  },

  methods: {
    init() {},
  },

  mounted() {},

  props: {
    modelValue: Object, // a 'VMO'
  },
}
</script>

<style></style>

<template>
  <div>
    <ElTable :data="items">
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
