<script lang="js">

import InnerHeaderImpl from './../common/header.vue'
import MyNavTable from './admin-nav-table.vue'

export default {

    name: "frames-admin-header",

    components: { InnerHeaderImpl, MyNavTable },

}

</script>

<style lang="css"></style>

<template>
    <div>
        <h2> (frame:admin) </h2>

        <MyNavTable> </MyNavTable>

        <slot></slot>

        <InnerHeaderImpl> </InnerHeaderImpl>
    </div>
</template>
