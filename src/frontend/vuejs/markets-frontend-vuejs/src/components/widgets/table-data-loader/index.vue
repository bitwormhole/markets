<script lang="js">

import lib_getter_js from '@/js/getter'
import lib_params_js from '@/js/params'

const the_getter = lib_getter_js.NewGetter()

function nop() {

}

export default {

    name: "widgets-table-data-loader",

    components: {},

    computed: {

        current() {
            let str = the_getter.From(this.query).Get("current").Result(1);
            return Number(str);
        },

        limit() {
            let str = the_getter.From(this.query).Get("limit").Result(5)
            return Number(str);
        },

        offset() {
            let limit = this.limit;
            let current = this.current - 1;
            return (limit * current);
        },

        total() {
            return the_getter.From(this.pagination).Get("total").Result(666)
        },

        now() {
            return this.getNow();
        },

        query() {
            return this.$route.query;
        },

        revision() {
            return the_getter.From(this.query).Get("revision").Result(0)
        },

    },

    data() {
        const params = {}
        const revision2 = this.getNow();
        return { revision2, params }
    },

    methods: {

        handleSizeChange(value) {
            let limit = value;
            let params = { limit }
            this.nav2({ params });
        },

        handleCurrentChange(value) {
            let limit = this.limit;
            let current = value;
            let params = { limit, current }
            this.nav2({ params });
        },

        fireOnLoadEvent() {

            let pbuf = lib_params_js.NewParams()
            let offset = this.offset;

            pbuf.Import(this.query);
            pbuf.Import({ offset });

            let params = pbuf.Export(null);
            let config = { params }
            this.$emit('onload', config)
        },

        getNow() {
            let t = new Date();
            return t.getTime();
        },

        nav2({ params }) {

            let revision = this.getNow();
            let pbuf = lib_params_js.NewParams()

            pbuf.Import(this.query)
            pbuf.Import(params)
            pbuf.Import({ revision })
            let query = pbuf.Export(null);

            query.offset = query.limit * (query.current - 1);

            this.$router.push({ query })
        },
    },

    mounted() {

        let limit = this.limit;
        let current = this.current;
        let offset = this.offset;

        let params = { limit, current, offset };
        this.nav2({ params })
    },

    watch: {
        revision(r1) {
            let r2 = this.revision2;
            if (r1 != r2) {
                this.fireOnLoadEvent();
            }
            this.revision2 = r1;
        },
    },

    props: {
        items: Array,
        pagination: Object,
        bank: String,
    }
}

</script>

<style></style>

<template>
    <div class="td-loader">
        <ElPagination :total="total" :page-size="limit" :current-page="current" :page-sizes="[5, 10, 20, 50, 100]"
            layout="total, sizes , prev, pager, next" @size-change="handleSizeChange"
            @current-change="handleCurrentChange">
        </ElPagination>
    </div>
</template>
