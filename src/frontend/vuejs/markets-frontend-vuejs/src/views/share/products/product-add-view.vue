<script lang="js">

import MyPage from '@/components/products/product-add-page.vue';

export default {

    name: "product-add-view",

    components: { MyPage },

    methods: {
        init() { },
    },

    mounted() {
        this.init()
    },

    props: {}
}

</script>

<style></style>

<template>
    <frame-for-user>
        <div>

            <h1> 商品 : 添加 </h1>

            <MyPage />

        </div>
    </frame-for-user>
</template>
