<script lang="js">

import MyPage from '@/components/licences/licence-index-page.vue';

export default {

    name: "licences-index-view",

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

            <h1> 许可证 : index </h1>

        </div>

        <MyPage></MyPage>

    </frame-for-admin>
</template>
