<script lang="js">
import { ElButton, ElTable, ElTableColumn } from 'element-plus';

const the_table_data = [

    {
        "type": "examples",
        "GET": { label: "x" },
        "POST": { label: "x" },
        "PUT": { label: "x" },
        "DEL": { label: "x" },
    },
    {
        "type": "companies",
        "GET": { label: "x" },
        "POST": { label: "x" },
        "PUT": { label: "x" },
        "DEL": { label: "x" },
    },
    {
        "type": "shops",
        "GET": {},
        "POST": {},
        "PUT": {},
        "DEL": {},
    },

];




export default {

    name: "frames-admin-nav-table",

    components: {},

    data() {
        return { the_table_data }
    },


}

</script>

<style lang="css"></style>

<template>
    <div>


        <ElTable :data="the_table_data">

            <ElTableColumn label="Type" prop="type"></ElTableColumn>
            <ElTableColumn label="Query"></ElTableColumn>
            <ElTableColumn label="Insert"></ElTableColumn>
            <ElTableColumn label="Update"></ElTableColumn>
            <ElTableColumn label="Remove"></ElTableColumn>

        </ElTable>

    </div>
</template>
