<script lang="js">

import MyPage from '@/components/trademarks/index.vue';

export default {

    name: "example-query-view",

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

            <h1> 商标 : index </h1>

        </div>

        <MyPage></MyPage>

    </frame-for-admin>
</template>
