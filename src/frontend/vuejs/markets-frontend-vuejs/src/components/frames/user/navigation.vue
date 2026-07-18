<script lang="js">
import { ElMenu, ElMenuItem } from 'element-plus'

export default {
  name: 'frames-user-navigation',

  components: { ElMenu, ElMenuItem },

  computed: {
    current() {
      return this.$route.path
    },
  },
}
</script>

<style lang="css"></style>

<template>
  <div>
    <ElMenu mode="horizontal" :default-active="current" router>
      <ElMenuItem index="/my/standards"> 标准 </ElMenuItem>
      <ElMenuItem index="/my/licences"> 许可 </ElMenuItem>

      <ElMenuItem index="/my/trademarks"> 品牌 </ElMenuItem>
      <ElMenuItem index="/my/companies"> 企业 </ElMenuItem>
      <ElMenuItem index="/my/manufacturers"> 制造商 </ElMenuItem>
      <ElMenuItem index="/my/shops"> 店铺 </ElMenuItem>

      <ElMenuItem index="/my/products"> 产品 </ElMenuItem>
      <ElMenuItem index="/my/skus"> SKU </ElMenuItem>

      <ElMenuItem index="/my/favorites"> 收藏夹 </ElMenuItem>
      <ElMenuItem index="/my/settings"> 设置 </ElMenuItem>
    </ElMenu>
  </div>
</template>
