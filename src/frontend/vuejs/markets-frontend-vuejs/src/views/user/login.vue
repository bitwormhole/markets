<script lang="js">


import MyPage from '@/components/sessions/session-login-page.vue'


export default {

    name: "views-user-login",

    components: { MyPage },

    methods: {

    },

    mounted() {
        // this.init()
    },

    props: {}
}

</script>

<style></style>

<template>
    <frame-for-user>
        <MyPage />
    </frame-for-user>
</template>
