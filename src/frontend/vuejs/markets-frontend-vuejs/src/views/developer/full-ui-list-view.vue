<script lang="js">
import { ElButton } from 'element-plus';
import { RouterLink } from 'vue-router';


export default {

    name: "full-ui-list-view",

    components: {},

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
    <div>

        <h1> Full UI List </h1>

        <hr />

        <div>
            <h2> For User </h2>
        </div>

        <hr />

        <div>
            <h2> For Admin </h2>


            <div>
                <h3> 企业 (/admin/companies/*) </h3>
                <RouterLink to="/admin/companies/"> 列表 </RouterLink>
                <RouterLink to="/admin/companies/detail"> 详情 </RouterLink>
                <RouterLink to="/admin/companies/add"> 添加 </RouterLink>
                <RouterLink to="/admin/companies/edit"> 编辑 </RouterLink>
                <RouterLink to="/admin/companies/query"> 查询 </RouterLink>
            </div>

            <div>
                <h3> 许可 (/admin/licences/*) </h3>
                <RouterLink to="/admin/licences/"> 列表 </RouterLink>
                <RouterLink to="/admin/licences/detail"> 详情 </RouterLink>
                <RouterLink to="/admin/licences/add"> 添加 </RouterLink>
                <RouterLink to="/admin/licences/edit"> 编辑 </RouterLink>
                <RouterLink to="/admin/licences/query"> 查询 </RouterLink>
            </div>

            <div>
                <h3> 厂家 (/admin/manufacturers/*) </h3>
                <RouterLink to="/admin/manufacturers/"> 列表 </RouterLink>
                <RouterLink to="/admin/manufacturers/detail"> 详情 </RouterLink>
                <RouterLink to="/admin/manufacturers/add"> 添加 </RouterLink>
                <RouterLink to="/admin/manufacturers/edit"> 编辑 </RouterLink>
                <RouterLink to="/admin/manufacturers/query"> 查询 </RouterLink>
            </div>

            <div>
                <h3> 商品 (/admin/products/*) </h3>
                <RouterLink to="/admin/products/"> 列表 </RouterLink>
                <RouterLink to="/admin/products/detail"> 详情 </RouterLink>
                <RouterLink to="/admin/products/add"> 添加 </RouterLink>
                <RouterLink to="/admin/products/edit"> 编辑 </RouterLink>
                <RouterLink to="/admin/products/query"> 查询 </RouterLink>
            </div>

            <div>
                <h3> 店铺 (/admin/shops/*) </h3>
                <RouterLink to="/admin/shops/"> 列表 </RouterLink>
                <RouterLink to="/admin/shops/detail"> 详情 </RouterLink>
                <RouterLink to="/admin/shops/add"> 添加 </RouterLink>
                <RouterLink to="/admin/shops/edit"> 编辑 </RouterLink>
                <RouterLink to="/admin/shops/query"> 查询 </RouterLink>
            </div>

            <div>
                <h3> SKU (/admin/skus/*) </h3>
                <RouterLink to="/admin/skus/"> 列表 </RouterLink>
                <RouterLink to="/admin/skus/detail"> 详情 </RouterLink>
                <RouterLink to="/admin/skus/add"> 添加 </RouterLink>
                <RouterLink to="/admin/skus/edit"> 编辑 </RouterLink>
                <RouterLink to="/admin/skus/query"> 查询 </RouterLink>
            </div>

            <div>
                <h3> 标准 (/admin/standards/*) </h3>
                <RouterLink to="/admin/standards/"> 列表 </RouterLink>
                <RouterLink to="/admin/standards/detail"> 详情 </RouterLink>
                <RouterLink to="/admin/standards/add"> 添加 </RouterLink>
                <RouterLink to="/admin/standards/edit"> 编辑 </RouterLink>
                <RouterLink to="/admin/standards/query"> 查询 </RouterLink>
            </div>

            <div>
                <h3> 品牌 (/admin/trademarks/*) </h3>
                <RouterLink to="/admin/trademarks/"> 列表 </RouterLink>
                <RouterLink to="/admin/trademarks/detail"> 详情 </RouterLink>
                <RouterLink to="/admin/trademarks/add"> 添加 </RouterLink>
                <RouterLink to="/admin/trademarks/edit"> 编辑 </RouterLink>
                <RouterLink to="/admin/trademarks/query"> 查询 </RouterLink>
            </div>

        </div>

        <hr />

        <div>
            <h2> For Developer </h2>
        </div>


    </div>
</template>
