<script lang="js">
import { ElButton, ElTableColumn } from 'element-plus'

import LibGetterJS from '@/js/getter.js'

export default {
  name: 'sessions-common-table-view',

  components: { ElButton, ElTableColumn },

  computed: {
    tableRows() {
      let gett = LibGetterJS.NewGetter()
      return gett.From(this.modelValue).Get('items').AsAny().Result([])
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
    modelValue: Object, // a 'VMO' object
  },
}
</script>

<style></style>

<template>
  <div>
    <ElTable :data="tableRows">
      <ElTableColumn prop="id" label="ID"></ElTableColumn>

      <ElTableColumn prop="name" label="Name"></ElTableColumn>

      <ElTableColumn prop="description" label="Description"></ElTableColumn>

      <ElTableColumn label="操作">
        <template #default>
          <ElButton link type="primary">Info</ElButton>
          <ElButton link type="primary">Kill</ElButton>
        </template>
      </ElTableColumn>
    </ElTable>
  </div>
</template>
