<script lang="js">

import { useAxiosLib } from '@/stores/axios';
import { ElButton, ElForm, ElFormItem, ElInput } from 'element-plus';
import MyLicenceTypeSelector from './licence-type-selector.vue'


const theAxiosLib = useAxiosLib()


export default {

  name: "licence-add-page",

  components: { MyLicenceTypeSelector },

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
      let url = '/api/v1/licences'
      let data = {
        licences: [item]
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

    <ElForm label-width="100">

      <ElFormItem label="许可证类型">
        <!-- <ElInput v-model="item.type"></ElInput> -->
        <MyLicenceTypeSelector v-model="item.type"> </MyLicenceTypeSelector>
      </ElFormItem>

      <ElFormItem label="许可证代码">
        <ElInput v-model="item.code"></ElInput>
      </ElFormItem>

      <ElFormItem label="生效日期">
        <!-- <ElInput v-model="item.not_before"></ElInput> -->
        <el-date-picker v-model="item.not_before" type="date" value-format="x" placeholder="请选择" />
      </ElFormItem>

      <ElFormItem label="失效日期">
        <el-date-picker v-model="item.not_after" type="date" value-format="x" placeholder="请选择" />
      </ElFormItem>

      <hr />

      <ElFormItem label="持有者-名称">
        <ElInput v-model="item.subject_name"></ElInput>
      </ElFormItem>

      <ElFormItem label="持有者-地址">
        <ElInput v-model="item.subject_address"></ElInput>
      </ElFormItem>

      <ElFormItem label="持有者-代码">
        <ElInput v-model="item.subject_code"></ElInput>
      </ElFormItem>

      <hr />

      <ElFormItem label="签发者-名称">
        <ElInput v-model="item.issuer_name"></ElInput>
      </ElFormItem>

      <!-- <ElFormItem label="签发者-地址">
        <ElInput v-model="item.issuer_address"></ElInput>
      </ElFormItem> -->


      <hr />


      <ElFormItem label="参考网址">
        <ElInput v-model="item.ref" />
      </ElFormItem>

      <ElFormItem label="备注">
        <ElInput v-model="item.remarks" type="textarea" />
      </ElFormItem>

    </ElForm>



    <ElButton type="success" @click="save"> Save </ElButton>

  </div>
</template>
