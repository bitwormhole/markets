<script lang="js">

import MyPage from '@/components/licences/licence-detail-page.vue'

export default {

    name: "licence-detail-view",

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
