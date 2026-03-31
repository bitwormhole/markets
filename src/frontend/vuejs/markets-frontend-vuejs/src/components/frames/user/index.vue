<script lang="js">

import MyHeader from './header.vue'
import MyFooter from './footer.vue'

export default {

    name: "frames-user-index",

    components: { MyHeader, MyFooter },

}

</script>

<style lang="css"></style>

<template>
    <div>
        <MyHeader></MyHeader>

        <slot></slot>

        <MyFooter></MyFooter>
    </div>
</template>
