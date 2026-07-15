<script lang="js">


import MyPage from '@/components/users/admin-user-query-page.vue';


export default {

    name: "user-query-view",

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
        <h2> 查询用户 (User) </h2>
        <MyPage></MyPage>
    </frame-for-admin>
</template>
