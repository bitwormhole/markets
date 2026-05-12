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
    <div>

        <h1> 添加用户 </h1>

        <MyPage />
    </div>
</template>
