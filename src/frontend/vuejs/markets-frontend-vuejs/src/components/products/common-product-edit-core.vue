<script lang="js">
import { ElForm, ElFormItem, ElInput } from 'element-plus'

export default {
  name: 'products-common-edit-all',

  components: { ElForm, ElFormItem, ElInput },

  computed: {
    dto() {
      return this.item
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
    item: Object, // a product-dto
  },
}
</script>

<style scoped></style>

<template>
  <div>
    <ElForm label-width="150">
      <ElFormItem label="品名"> <ElInput v-model="dto.name"></ElInput> </ElFormItem>
      <ElFormItem label="条码"> <ElInput v-model="dto.code"></ElInput> </ElFormItem>
      <ElFormItem label="品牌"> <ElInput v-model="dto.tm_name"></ElInput> </ElFormItem>
      <ElFormItem label="标准代号"> <ElInput v-model="dto.standard_code"></ElInput> </ElFormItem>
      <ElFormItem label="标签"> <ElInput v-model="dto.label"></ElInput> </ElFormItem>
      <ElFormItem label="备注"> <ElInput v-model="dto.remark"></ElInput> </ElFormItem>
      <ElFormItem label="描述">
        <ElInput v-model="dto.description" type="textarea"></ElInput>
      </ElFormItem>
    </ElForm>
  </div>
</template>
