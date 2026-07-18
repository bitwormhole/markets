<script lang="js">
import MySessionInfoBox from '@/components/sessions/session-state-info-box.vue'

export default {
  name: 'frames-user-header',

  components: { MySessionInfoBox },
}
</script>

<style lang="css"></style>

<template>
  <div>
    <h2>Markets-DB</h2>

    <hr />

    <MySessionInfoBox />
  </div>
</template>
