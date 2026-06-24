<script lang="js">

import MyPage from '@/components/skus/sku-detail-page.vue'

export default {

    name: "sku-detail-view",

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
