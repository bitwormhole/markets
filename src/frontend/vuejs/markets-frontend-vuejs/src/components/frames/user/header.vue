<script lang="js">
import MyHeaderBase from './../common/header.vue'
import MyNavigation from './navigation.vue'

export default {
  name: 'frames-user-header',

  components: { MyHeaderBase, MyNavigation },
}
</script>

<style lang="css"></style>

<template>
  <div>
    <h2>(frame:user)</h2>

    <MyHeaderBase> </MyHeaderBase>

    <MyNavigation> </MyNavigation>
  </div>
</template>
