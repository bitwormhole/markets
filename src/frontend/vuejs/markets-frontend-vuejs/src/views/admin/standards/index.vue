<script lang="js">

import MyPage from '@/components/standards/standard-index-page.vue';

export default {

    name: "standards-index-view",

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

            <h1> 标准 : index </h1>

        </div>

        <MyPage></MyPage>
    </frame-for-admin>
</template>
