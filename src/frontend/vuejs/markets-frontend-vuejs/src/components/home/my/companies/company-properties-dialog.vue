<script lang="js">

import { ElButton, ElDialog, ElDrawer, ElForm, ElFormItem, ElInput } from 'element-plus';

import MyForm from './company-edit-form.vue'


export default {

  name: "company-properties-dialog",

  components: { MyForm },

  data() {
    return {}
  },

  methods: {

    close() {
      this.$emit('update:modelValue', false)
    },

  },

  mounted() {
  },

  props: {
    item: Object,
  }
}

</script>

<style></style>

<template>
  <ElDrawer title="企业属性" size="80%" resizable>
    <div>

      <ElForm label-width="100">
        <ElFormItem label="ID"> {{ item.id }} </ElFormItem>
        <ElFormItem label="UUID"> {{ item.uuid }} </ElFormItem>
      </ElForm>

      <hr />

      <MyForm :model-value="item" readonly> </MyForm>

      <hr />

      <div>
        <ElButton type="default" @click="close"> 关闭 </ElButton>
      </div>
    </div>
  </ElDrawer>
</template>
