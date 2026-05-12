<script lang="js">

export default {

    name: "frames-admin-nav-table",

    components: {},

    data() {
        return {}
    },

}

</script>

<style lang="css">
.nav-link-item {
    margin: 3px;
    padding: 3px;

    border-width: 1px;
    border-style: solid;
    border-color: rgb(200, 200, 200);
}
</style>

<template>
    <div>
        <div>

            <RouterLink class="nav-link-item" to="/admin/companies"> Companies </RouterLink>
            <RouterLink class="nav-link-item" to="/admin/licences"> Licences </RouterLink>
            <RouterLink class="nav-link-item" to="/admin/manufacturers"> Manufacturers </RouterLink>

            <RouterLink class="nav-link-item" to="/admin/products"> Products </RouterLink>
            <RouterLink class="nav-link-item" to="/admin/shops"> Shops </RouterLink>
            <RouterLink class="nav-link-item" to="/admin/skus"> SKUs </RouterLink>

            <RouterLink class="nav-link-item" to="/admin/standards"> Standards </RouterLink>
            <RouterLink class="nav-link-item" to="/admin/trademarks"> Trademarks </RouterLink>

            <RouterLink class="nav-link-item" to="/admin/users"> Users </RouterLink>

        </div>
    </div>
</template>
