<script lang="js">

import InnerHeaderImpl from './../common/header.vue'

export default {

    name: "frames-admin-header",

    components: { InnerHeaderImpl },

}

</script>

<style lang="css"></style>

<template>
    <div>
        <h2> (frame:admin) </h2>

        <slot></slot>

        <InnerHeaderImpl> </InnerHeaderImpl>
    </div>
</template>
