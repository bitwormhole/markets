<script lang="js">

import { useAxiosLib } from '@/stores/axios';
import { ElButton, ElDialog, ElForm, ElFormItem, ElInput } from 'element-plus';


const theAxiosLib = useAxiosLib()


export default {

  name: "company-delete-dialog",

  components: {},

  data() {
    return {}
  },

  methods: {
    cancel() {
      this.$emit('update:modelValue', false)
    },

    remove() {

      let item = this.item;
      let id = item.id;

      let method = 'DELETE'
      let url = '/api/v1/companies/' + id

      theAxiosLib.execute({ method, url })
    },

  },

  mounted() {
    // this.init()
  },

  props: {
    item: Object,
  }
}

</script>

<style></style>

<template>
  <ElDialog title="删除条目 (企业)">
    <div>

      <h3> 确定要删除以下条目吗? </h3>

      <hr />

      <ElForm label-width="200">

        <ElFormItem label="ID"> {{ item.id }} </ElFormItem>
        <ElFormItem label="UUID"> {{ item.uuid }} </ElFormItem>

        <ElFormItem label="企业名称"> {{ item.name }} </ElFormItem>

        <ElFormItem label="统一社会信用代码"> {{ item.code }} </ElFormItem>
      </ElForm>

      <hr />

      <div>
        <ElButton type="default" @click="cancel"> 取消 </ElButton>
        <ElButton type="danger" @click="remove"> 删除 </ElButton>
      </div>
    </div>
  </ElDialog>
</template>
