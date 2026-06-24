<script lang="js">

import { useAxiosLib } from '@/stores/axios';
import { ElButton, ElForm, ElFormItem, ElInput } from 'element-plus';

import MyForm from './trademark-info-form.vue'

const theAxiosLib = useAxiosLib()


export default {

  name: "trademark-detail-page",

  components: { MyForm },

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

      let method = 'GET'
      let url = '/api/v1/products'
      let data = {
        products: [item]
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
    <MyForm />
  </div>
</template>
