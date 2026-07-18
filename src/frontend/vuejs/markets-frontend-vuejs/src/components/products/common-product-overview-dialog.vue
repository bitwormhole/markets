<script lang="js">
import { ElButton, ElButtonGroup, ElDialog, ElTable, ElTableColumn } from 'element-plus'
import MyInnerForm from './common-product-overview-form.vue'

export default {
  name: 'products-common-overview-dialog',

  components: { MyInnerForm },

  computed: {},

  data() {
    return {}
  },

  methods: {
    init() {},

    handleClickClose() {
      this.$emit('update:modelValue', false)
    },
  },

  mounted() {},

  props: {
    item: Object, // a product-dto
  },
}
</script>

<style></style>

<template>
  <ElDialog title="产品概览">
    <MyInnerForm :item="item"> </MyInnerForm>
    <div>
      <ElButton @click="handleClickClose"> Close </ElButton>
    </div>
  </ElDialog>
</template>
