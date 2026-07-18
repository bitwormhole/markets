<script lang="js">
// import { ElButton, ElTable, ElTableColumn } from 'element-plus'

export default {
  name: 'products-common-overview-form',

  components: {},

  computed: {},

  data() {
    return {}
  },

  methods: {
    init() {},
  },

  mounted() {},

  props: {
    item: Object, // a product-dto
  },
}
</script>

<style></style>

<template>
  <ElForm label-width="150">
    <ElFormItem label="ID"> {{ item.id }} </ElFormItem>
    <ElFormItem label="品名"> {{ item.name }} </ElFormItem>
    <ElFormItem label="条码"> {{ item.code }} </ElFormItem>
    <ElFormItem label="执行标准"> {{ item.standard_code }} </ElFormItem>
  </ElForm>
</template>
