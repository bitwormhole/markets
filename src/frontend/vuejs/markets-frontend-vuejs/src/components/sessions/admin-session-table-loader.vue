<script lang="js">

import { useAxiosLib } from '@/stores/axios';

const theAxiosLib = useAxiosLib();

export default {

  name: "sessions-admin-table-loader",

  components: {},

  data() {
    return {}
  },

  methods: {

    init() {
      this.fetch()
    },

    fetch() {
      let method = 'GET'
      let url = '/api/v1/admin/sessions'
      let params = {
        limit: 10,
        offset: (10 * 3),
      }
      let p = theAxiosLib.execute({ method, url, params })
      p.then((res) => {
        let vo = res.data;
        this.fireOnData(vo)
        this.fireOnItems(vo.sessions)
        this.fireOnPage(vo.pagination)
      })
      return p
    },

    tryUseDefaultValue(value, value_default) {
      if (value == null) {
        value = value_default;
      }
      return value;
    },

    fireOnData(vo) {
      vo = this.tryUseDefaultValue(vo, {})
      this.$emit('on-data', vo)
    },
    fireOnItems(list) {
      list = this.tryUseDefaultValue(list, [])
      this.$emit('on-items', list)
    },
    fireOnPage(pagination) {
      pagination = this.tryUseDefaultValue(pagination, {})
      this.$emit('on-page', pagination)
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


  </div>
</template>
