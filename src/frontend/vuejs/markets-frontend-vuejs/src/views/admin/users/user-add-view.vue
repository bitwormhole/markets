<script lang="js">

import MyPage from '@/components/users/user-add-page.vue';

export default {

    name: "user-add-view",

    components: { MyPage },

    methods: {
        // init() { },
    },

    mounted() {
        // this.init()
    },

    props: {}
}

</script>

<style></style>

<template>
    <frame-for-admin>
        <h1> 添加用户 (User) </h1>
        <MyPage />
    </frame-for-admin>
</template>
