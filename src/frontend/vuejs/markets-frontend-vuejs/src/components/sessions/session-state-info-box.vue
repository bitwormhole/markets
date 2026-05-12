<script lang="js">

// import { useAxiosLib } from '@/stores/axios';
// import { ElButton, ElForm, ElFormItem, ElInput } from 'element-plus';

import { useSessionStore } from '@/stores/sessions.js';


const theSessionStore = useSessionStore()


export default {

  name: "session-state-info-box",

  components: {},

  data() {
    return {}
  },

  methods: {
    fetch() {
      theSessionStore.fetch()
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
    session.current_state : {}
  </div>
</template>
