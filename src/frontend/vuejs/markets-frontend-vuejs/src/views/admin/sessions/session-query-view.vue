<script lang="js">


import InnerPage from '@/components/sessions/admin-session-query-page.vue'


export default {

    name: "example-query-view",

    components: { InnerPage },

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
        <h1> 查询会话 (Session) </h1>
        <InnerPage> </InnerPage>
    </frame-for-admin>
</template>
