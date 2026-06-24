<script lang="js">

import MyPage from '@/components/companies/company-detail-page.vue'

export default {

    name: "company-detail-view",

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
    <frame-for-user>
        <MyPage />
    </frame-for-user>
</template>
