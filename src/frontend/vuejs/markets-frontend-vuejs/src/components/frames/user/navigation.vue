<script lang="js">

// import MyHeader from './header.vue'
// import MyFooter from './footer.vue'

import { ElMenu, ElMenuItem } from 'element-plus';

export default {

    name: "frames-user-navigation",

    // components: { MyHeader, MyFooter },

    computed: {
        current() {
            let x = this.$route.path;
            return x; //  "/products";
        },
    },

}

</script>

<style lang="css"></style>

<template>
    <div>
        <ElMenu mode="horizontal" :default-active="current" router>

            <ElMenuItem index="/standards"> 标准 </ElMenuItem>
            <ElMenuItem index="/licences"> 许可 </ElMenuItem>

            <ElMenuItem index="/trademarks"> 品牌 </ElMenuItem>
            <ElMenuItem index="/companies"> 企业 </ElMenuItem>
            <ElMenuItem index="/manufacturers"> 制造商 </ElMenuItem>
            <ElMenuItem index="/shops"> 店铺 </ElMenuItem>

            <ElMenuItem index="/products"> 产品 </ElMenuItem>
            <ElMenuItem index="/skus"> SKU </ElMenuItem>

            <ElMenuItem index="/my/favorites"> 收藏夹 </ElMenuItem>

        </ElMenu>

    </div>
</template>
