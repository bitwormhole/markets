<script lang="js">
import MyPage from '@/components/roles/admin-role-query-page.vue'

export default {
  name: 'views-admin-roles-index',

  components: { MyPage },

  methods: {},

  mounted() {},

  props: {},
}
</script>

<style></style>

<template>
  <frame-for-admin>
    <h1>角色 (Role)</h1>

    <MyPage></MyPage>
  </frame-for-admin>
</template>
