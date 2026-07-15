<script lang="js">

import MyPage from '@/components/permissions/admin-permission-query-page.vue';

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
        <h1> 权限 (Permission) </h1>
        <MyPage></MyPage>
    </frame-for-admin>
</template>
