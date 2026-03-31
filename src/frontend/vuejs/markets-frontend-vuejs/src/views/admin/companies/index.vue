<script lang="js">

import MyPage from '@/components/companies/company-index-page.vue';

export default {

    name: "admin-companies-index-view",

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
        <div>

            <h1> 企业 : index </h1>

        </div>

        <MyPage></MyPage>

    </frame-for-admin>
</template>
