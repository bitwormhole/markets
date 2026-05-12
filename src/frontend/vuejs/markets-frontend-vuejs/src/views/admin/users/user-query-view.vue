<script lang="js">


import MyPage from '@/components/users/user-query-page.vue';


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
        <MyPage></MyPage>
    </frame-for-admin>
</template>
