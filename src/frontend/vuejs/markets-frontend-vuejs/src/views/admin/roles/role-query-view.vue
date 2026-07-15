<script lang="js">


import InnerView from '@/components/roles/admin-role-query-page.vue'


export default {

    name: "example-query-view",

    components: { InnerView },

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
        <h1> 查询角色 (Role) </h1>

        <InnerView> </InnerView>
    </frame-for-admin>
</template>
