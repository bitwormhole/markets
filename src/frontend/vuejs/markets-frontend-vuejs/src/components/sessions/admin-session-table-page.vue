<script lang="js">

// import { useAxiosLib } from '@/stores/axios';
// import { ElButton, ElForm, ElFormItem, ElInput } from 'element-plus';

import MyLoginBox from './session-login-box.vue';


// const theSessionStore = useSessionStore()


export default {

  name: "session-login-page",

  components: { MyLoginBox },

  data() {
    return {}
  },

  methods: {
    fetch() {
      // theSessionStore.fetch()
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
    <MyLoginBox />
  </div>
</template>
