<script lang="js">

import MyPage from '@/components/examples/admin-examples-add-page.vue';

export default {

    name: "example-add-view",

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
        <h1> 添加会话 (Session) </h1>
        <MyPage />
    </frame-for-admin>
</template>
