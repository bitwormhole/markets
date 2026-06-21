<script lang="js">

import InnerHeaderImpl from './../common/header.vue'
import InnerNavigation from './navigation.vue'

export default {

    name: "frames-user-header",

    components: { InnerHeaderImpl, InnerNavigation },

}

</script>

<style lang="css"></style>

<template>
    <div>
        <h2> (frame:user) </h2>

        <InnerHeaderImpl> </InnerHeaderImpl>

        <InnerNavigation> </InnerNavigation>

    </div>
</template>
