<script lang="js">

import MyPage from '@/components/sessions/admin-session-query-page.vue';

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
        <h1> 会话 (Session) </h1>
        <MyPage></MyPage>
    </frame-for-admin>
</template>
