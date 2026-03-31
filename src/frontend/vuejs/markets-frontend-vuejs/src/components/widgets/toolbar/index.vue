<script lang="js">

// import { ElButton, ElForm, ElFormItem, ElInput } from 'element-plus';

export default {

    name: "widgets-toolbar-index",

    components: {},

    data() {
        return {}
    },

    props: {}
}

</script>

<style>
.toolbar {
    margin: 5px;
    padding: 5px;

    background-color: rgb(161, 161, 161);

    border-radius: 3px;
    border-color: gray;
    border-width: 1px;
    border-style: solid;
}
</style>

<template>
    <div class="toolbar">
        <slot></slot>
    </div>
</template>
