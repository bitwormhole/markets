<script lang="js">

import MyHeader from './header.vue'
import MyFooter from './footer.vue'
import MyNavTab from './admin-nav-table.vue'

export default {

    name: "frames-admin-index",

    components: { MyHeader, MyFooter, MyNavTab },

}

</script>

<style lang="css"></style>

<template>
    <div>
        <MyHeader>
            <!-- <MyNavTab> </MyNavTab> -->
        </MyHeader>

        <slot></slot>

        <MyFooter></MyFooter>
    </div>
</template>
