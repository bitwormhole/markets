<script lang="js">
import InnerFooterImpl from './../common/footer.vue'
import InnerLogViewer from '@/components/widgets/log-viewer/index.vue'

export default {
  name: 'frames-developer-footer',

  components: { InnerFooterImpl, InnerLogViewer },
}
</script>

<style lang="css"></style>

<template>
  <div>
    <InnerFooterImpl> </InnerFooterImpl>
    <InnerLogViewer> </InnerLogViewer>
  </div>
</template>
