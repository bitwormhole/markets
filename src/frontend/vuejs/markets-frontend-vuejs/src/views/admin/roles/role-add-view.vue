<script lang="js">

import MyPage from '@/components/roles/role-add-page.vue';

export default {

    name: "admin-roles-add-view",

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

        <h1> 添加角色 (Role) </h1>
        <MyPage />

    </frame-for-admin>
</template>
