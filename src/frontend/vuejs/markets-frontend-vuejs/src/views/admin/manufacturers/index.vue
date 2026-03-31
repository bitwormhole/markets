<script lang="js">

import MyPage from '@/components/manufacturers/manufacturer-index-page.vue';

export default {

    name: "manufacturers-index-view",

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

            <h1> 生产商 : index </h1>

        </div>

        <MyPage></MyPage>

    </frame-for-admin>
</template>
