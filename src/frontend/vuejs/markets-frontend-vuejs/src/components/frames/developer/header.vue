<script lang="js">

import InnerHeaderImpl from './../common/header.vue'
import MyNavTable from './dev_nav_table.vue'

export default {

    name: "frames-developer-header",

    components: { InnerHeaderImpl, MyNavTable },

}

</script>

<style lang="css"></style>

<template>
    <div>
        <h2> (frame:dev) </h2>

        <InnerHeaderImpl> </InnerHeaderImpl>

        <MyNavTable> </MyNavTable>


        <div>
            <RouterLink to="/admin/products"> Products </RouterLink>
            <RouterLink to="/admin/shops"> Shops </RouterLink>
            <RouterLink to="/admin/skus"> SKUs </RouterLink>

        </div>


    </div>
</template>
