<script lang="js">

import MyPage from '@/components/standards/standard-add-page.vue';

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
    <frame-for-user>
        <div>

            <h1> 标准 : 添加 </h1>

            <MyPage />

        </div>
    </frame-for-user>
</template>
