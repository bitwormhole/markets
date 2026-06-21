<script lang="js">


import MyPage from '@/components/companies/company-edit-page.vue';


export default {

    name: "example-edit-view",

    components: { MyPage },

    methods: {
    },

    mounted() {
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
