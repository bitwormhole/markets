<script lang="js">

import { useAxiosLib } from '@/stores/axios';
import { ElButton, ElForm, ElFormItem, ElInput } from 'element-plus';

import MyEditForm from './company-edit-form.vue'
import lib_getter_js from '@/js/getter'

const theAxiosLib = useAxiosLib()
const the_getter = lib_getter_js.NewGetter()

export default {

  name: "company-edit-page",

  components: { MyEditForm },

  computed: {
    object_id() {
      return this.$route.params.id;
    },
  },

  data() {
    const item = {}
    return { item }
  },

  methods: {

    save() {

      let item = this.item;
      let id = this.object_id;

      let method = 'PUT'
      let url = '/api/v1/companies/' + id
      let data = {
        companies: [item]
      }

      theAxiosLib.execute({ method, url, data })
    },

    fetch() {

      let id = this.object_id;

      let method = 'GET'
      let url = '/api/v1/companies/' + id;

      theAxiosLib.execute({ method, url }).then((res) => {
        let vo = res.data;
        let item = the_getter.Get('companies/0').From(vo).Result({})
        this.item = item;
      })

    },
  },

  mounted() {
    this.fetch()
  },

  props: {}
}

</script>

<style></style>

<template>
  <div>

    <MyEditForm v-model="item" />

    <ElButton type="success" @click="save"> Save </ElButton>

  </div>
</template>
