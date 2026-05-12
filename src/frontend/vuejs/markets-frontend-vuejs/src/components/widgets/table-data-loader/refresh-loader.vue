<script lang="js">

import lib_getter_js from '@/js/getter'
import lib_params_js from '@/js/params'

const the_getter = lib_getter_js.NewGetter()

function nop() {

}

export default {

    name: "widgets-table-refresh-loader",

    components: {},

    computed: {


    },

    data() {
        const foo = {}
        return { foo }
    },

    methods: {

        getNow() {
            let t = new Date();
            return t.getTime();
        },

        refresh() {

            let pbuf = lib_params_js.NewParams();
            let revision = this.getNow()

            pbuf.Import(this.$route.query);
            pbuf.Import({ revision })

            let query = pbuf.Export();
            this.$router.push({ query })
        },
    },

    mounted() {

    },


    watch: {
        revision() {
            this.refresh();
        },
    },

    props: {
        revision: Number,
    }
}

</script>

<style></style>

<template>
    <span class="td-refresh-loader"> </span>
</template>
