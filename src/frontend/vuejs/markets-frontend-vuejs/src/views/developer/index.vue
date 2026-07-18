<script lang="js">
import { RouterLink } from 'vue-router'

export default {
  name: 'example-query-view',

  components: {},

  methods: {
    init() {},
  },

  mounted() {
    this.init()
  },

  props: {},
}
</script>

<style scoped>
.link {
  padding: 5px;
  margin: 5px;
  border-color: gray;
  border-style: solid;
  border-width: 1px;
}
</style>

<template>
  <frame-for-developer>
    <h1>developer: home</h1>

    <div>
      <router-link class="link" to="/dev/ui"> All UI </router-link>

      <router-link class="link" to="/dev/show-current-session-info">
        Show Current Session Info
      </router-link>

      <router-link class="link" to="/dev/upload-media-file"> Upload Media File </router-link>

      <router-link class="link" to="/dev/scan-db-tables"> Scan DB Tables </router-link>
    </div>
  </frame-for-developer>
</template>
