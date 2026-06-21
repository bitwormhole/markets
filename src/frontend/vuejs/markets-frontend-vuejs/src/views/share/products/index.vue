<script lang="js">

import MyPage from '@/components/products/index.vue';

export default {

    name: "product-index-view",

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

            <h1> 产品 : index </h1>

        </div>

        <MyPage></MyPage>

    </frame-for-user>
</template>
