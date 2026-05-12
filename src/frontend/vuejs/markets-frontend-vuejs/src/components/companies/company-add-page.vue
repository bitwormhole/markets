<script lang="js">

import { useAxiosLib } from '@/stores/axios';
import { ElButton, ElForm, ElFormItem, ElInput } from 'element-plus';

import MyEditForm from './company-edit-form.vue'


import MyDateSelector from '@/components/widgets/date-time-box/date-selector.vue'


const theAxiosLib = useAxiosLib()


export default {

  name: "company-add-page",

  components: { MyDateSelector, MyEditForm },

  data() {
    const item = {
      name: '',
      ptype: '',
      category: '',
      code: '',
    }
    return { item }
  },

  methods: {
    init() { },


    save() {

      let item = this.item;

      let method = 'POST'
      let url = '/api/v1/companies'
      let data = {
        companies: [item]
      }
      theAxiosLib.execute({ method, url, data })
    },
  },

  mounted() {
    this.init()
  },

  props: {}
}

</script>

<style></style>

<template>
  <div>

    <!-- <ElForm label-width="150">

      <ElFormItem label="统一社会信用代码">
        <ElInput v-model="item.code"></ElInput>
      </ElFormItem>

      <ElFormItem label="企业名称">
        <ElInput v-model="item.name"></ElInput>
      </ElFormItem>

      <ElFormItem label="法定代表人">
        <ElInput v-model="item.representative"></ElInput>
      </ElFormItem>

      <ElFormItem label="类型">
        <ElInput v-model="item.company_type"></ElInput>
      </ElFormItem>

      <ElFormItem label="住所">
        <ElInput v-model="item.address"></ElInput>
      </ElFormItem>

      <ElFormItem label="成立日期">
        <MyDateSelector v-model="item.founded_at"></MyDateSelector>
      </ElFormItem>

      <ElFormItem label="核准日期">
        <MyDateSelector v-model="item.approved_at"></MyDateSelector>
      </ElFormItem>

      <ElFormItem label="注册资本">
        <ElInput v-model="item.capital"></ElInput>
      </ElFormItem>

      <ElFormItem label="登记机关">
        <ElInput v-model="item.registry"></ElInput>
      </ElFormItem>

      <ElFormItem label="登记状态">
        <ElInput v-model="item.state"></ElInput>
      </ElFormItem>

      <ElFormItem label="经营类别">
        <ElInput v-model="item.operation_category"></ElInput>
      </ElFormItem>

      <ElFormItem label="经营范围">
        <ElInput v-model="item.operation_range" type="textarea"></ElInput>
      </ElFormItem>

      <ElFormItem label="官网首页网址">
        <ElInput v-model="item.web"></ElInput>
      </ElFormItem>

      <ElFormItem label="参考网址">
        <ElInput v-model="item.reference"></ElInput>
      </ElFormItem>

      <ElFormItem label="备注">
        <ElInput v-model="item.remarks" type="textarea"></ElInput>
      </ElFormItem>

    </ElForm> -->

    <MyEditForm v-model="item" />

    <ElButton type="success" @click="save"> Save </ElButton>

  </div>
</template>
