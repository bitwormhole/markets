<script lang="js">
import { ElButton, ElTableColumn, ElTable } from 'element-plus'

export default {
  name: 'sessions-common-table',

  components: { ElTable, ElTableColumn, ElButton },

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
