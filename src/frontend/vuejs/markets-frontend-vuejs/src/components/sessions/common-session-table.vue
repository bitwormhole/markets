<script lang="js">
import { ElButton, ElTable, ElTableColumn } from 'element-plus'

export default {
  name: 'sessions-common-table',

  components: { ElButton, ElTable, ElTableColumn },

  computed: {
    items() {
      return this.modelValue['items']
    },
  },

  data() {
    return {}
  },

  methods: {},

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
      <ElTableColumn prop="id" label="ID" width="100"></ElTableColumn>

      <ElTableColumn prop="uuid" label="UUID" width="320"></ElTableColumn>

      <ElTableColumn prop="owner" label="User.ID"></ElTableColumn>
      <ElTableColumn prop="username" label="User.Name"></ElTableColumn>
      <ElTableColumn prop="nickname" label="Nickname"></ElTableColumn>

      <ElTableColumn prop="created_at" label="T1"></ElTableColumn>

      <ElTableColumn prop="expired_at" label="T2"></ElTableColumn>

      <ElTableColumn label="操作">
        <template #default>
          <ElButton link type="primary">Info</ElButton>
          <ElButton link type="primary">Kill</ElButton>
        </template>
      </ElTableColumn>
    </ElTable>
  </div>
</template>
