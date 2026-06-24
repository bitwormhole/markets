<script lang="js">

export default {

    name: "example-detail-view",

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
        <ElForm label-width="100">

            <ElFormItem label="商品名称">
                <ElInput v-model="item.name"></ElInput>
            </ElFormItem>

            <ElFormItem label="商品类别">
                <ElInput v-model="item.category"></ElInput>
            </ElFormItem>

            <ElFormItem label="商品类型">
                <ElInput v-model="item.ptype"></ElInput>
            </ElFormItem>

            <ElFormItem label="商品条码">
                <ElInput v-model="item.code"></ElInput>
            </ElFormItem>

        </ElForm>
    </div>
</template>
