<script lang="js">

import MyPage from '@/components/permissions/permission-add-page.vue';

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
            <h1> 添加权限 (Permission) </h1>
            <MyPage />
        </div>
    </frame-for-admin>
</template>
