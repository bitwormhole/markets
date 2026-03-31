<script lang="js">

import MyPage from '@/components/manufacturers/manufacturer-add-page.vue';

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
    <div>

        <h1> 制造商 : 添加 </h1>

        <MyPage />

    </div>
</template>
