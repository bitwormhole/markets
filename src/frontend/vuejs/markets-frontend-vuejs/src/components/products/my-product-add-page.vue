<script lang="js">
import { useAxiosLib } from '@/stores/axios'
import { ElButton, ElForm, ElFormItem, ElInput } from 'element-plus'

const theAxiosLib = useAxiosLib()

export default {
  name: 'product-add-page',

  components: { ElButton, ElForm, ElFormItem, ElInput },

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
    init() {},

    save() {
      let item = this.item
      let method = 'POST'
      let url = '/api/v1/my/products'
      let data = {
        products: [item],
      }
      theAxiosLib.execute({ method, url, data })
    },
  },

  mounted() {
    this.init()
  },

  props: {},
}
</script>

<style></style>

<template>
  <div>
    <ElForm label-width="100">
      <ElFormItem label="品名">
        <ElInput v-model="item.name"></ElInput>
      </ElFormItem>

      <ElFormItem label="品牌">
        <ElInput v-model="item.tm_name"></ElInput>
      </ElFormItem>

      <ElFormItem label="条码">
        <ElInput v-model="item.code"></ElInput>
      </ElFormItem>

      <ElFormItem label="备注">
        <ElInput v-model="item.remark"></ElInput>
      </ElFormItem>
    </ElForm>

    <ElButton type="success" @click="save"> Save </ElButton>
  </div>
</template>
