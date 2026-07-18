<script lang="js">
import InnerFooterImpl from './../common/footer.vue'
import MyLogViewer from '@/components/widgets/log-viewer/index.vue'

export default {
  name: 'frames-admin-footer',

  components: { InnerFooterImpl, MyLogViewer },
}
</script>

<style lang="css"></style>

<template>
  <div>
    <InnerFooterImpl> </InnerFooterImpl>
    <MyLogViewer> </MyLogViewer>
  </div>
</template>
