<script lang="js">
import InnerLoader from '@/components/widgets/table-data-loader/main-table-loader.vue'

export default {
  name: 'sessions-admin-table-loader',

  components: { InnerLoader },

  data() {
    return {}
  },

  methods: {
    init() {},
  },

  mounted() {
    this.init()
  },

  props: {},
}
</script>

<style></style>

<template>
  <InnerLoader path="/api/v1/admin/roles" auto />
</template>
