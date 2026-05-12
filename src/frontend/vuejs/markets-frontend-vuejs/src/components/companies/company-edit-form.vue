<script lang="js">

import { ElForm, ElFormItem, ElInput } from 'element-plus';

import MyDateSelector from '@/components/widgets/date-time-box/date-selector.vue'


export default {

  name: "company-edit-form",

  components: { MyDateSelector },

  computed: {
    item() {
      return this.modelValue;
    }
  },

  data() {
    return {}
  },

  methods: {
    init() { },

    updateValue(item) {
      if (item == null) {
        item = this.modelValue;
      }
      this.$emit('update:modelValue', item)
    },
  },

  mounted() {
    this.init()
  },

  props: {
    modelValue: Object,

    readonly: Boolean,
  }
}

</script>

<style></style>

<template>
  <div>

    <ElForm label-width="150">

      <ElFormItem label="统一社会信用代码">
        <ElInput v-model="item.code" :readonly="readonly"></ElInput>
      </ElFormItem>

      <ElFormItem label="企业名称">
        <ElInput v-model="item.name" :readonly="readonly"></ElInput>
      </ElFormItem>

      <ElFormItem label="法定代表人">
        <ElInput v-model="item.representative" :readonly="readonly"></ElInput>
      </ElFormItem>

      <ElFormItem label="类型">
        <ElInput v-model="item.company_type" :readonly="readonly"></ElInput>
      </ElFormItem>

      <ElFormItem label="住所">
        <ElInput v-model="item.address" :readonly="readonly"></ElInput>
      </ElFormItem>

      <ElFormItem label="成立日期">
        <MyDateSelector v-model="item.founded_at" :readonly="readonly"></MyDateSelector>
      </ElFormItem>

      <ElFormItem label="核准日期">
        <MyDateSelector v-model="item.approved_at" :readonly="readonly"></MyDateSelector>
      </ElFormItem>

      <ElFormItem label="注册资本">
        <ElInput v-model="item.capital" :readonly="readonly"></ElInput>
      </ElFormItem>

      <ElFormItem label="登记机关">
        <ElInput v-model="item.registry" :readonly="readonly"></ElInput>
      </ElFormItem>

      <ElFormItem label="登记状态">
        <ElInput v-model="item.state" :readonly="readonly"></ElInput>
      </ElFormItem>

      <ElFormItem label="经营类别">
        <ElInput v-model="item.operation_category" :readonly="readonly"></ElInput>
      </ElFormItem>

      <ElFormItem label="经营范围">
        <ElInput v-model="item.operation_range" type="textarea" :readonly="readonly"></ElInput>
      </ElFormItem>

      <ElFormItem label="官网首页网址">
        <ElInput v-model="item.web" :readonly="readonly"></ElInput>
      </ElFormItem>

      <ElFormItem label="参考网址">
        <ElInput v-model="item.reference" :readonly="readonly"></ElInput>
      </ElFormItem>

      <ElFormItem label="备注">
        <ElInput v-model="item.remarks" type="textarea" :readonly="readonly"></ElInput>
      </ElFormItem>

    </ElForm>

  </div>
</template>
