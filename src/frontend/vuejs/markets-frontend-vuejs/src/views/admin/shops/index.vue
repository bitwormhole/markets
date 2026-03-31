<script lang="js">

import MyPage from '@/components/shops/shop-index-page.vue';

export default {

    name: "shops-index-view",

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
    <frame-for-admin>
        <div>

            <h1> 商店 : index </h1>

        </div>

        <MyPage></MyPage>

    </frame-for-admin>
</template>
