<script lang="js">

import MyPage from '@/components/examples/index.vue';

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
    <frame-for-user>
        <div>

            <h1> Example : index </h1>

        </div>

        <MyPage></MyPage>

    </frame-for-user>
</template>
