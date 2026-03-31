<script lang="js">

import MyPage from '@/components/companies/company-add-page.vue';

export default {

    name: "example-add-view",

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

            <h1> 企业 : 添加 </h1>

            <MyPage />

        </div>
    </frame-for-admin>
</template>
