<script lang="js">

import MyPage from '@/components/companies/company-query-page.vue';

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
        <MyPage />
    </frame-for-admin>
</template>
