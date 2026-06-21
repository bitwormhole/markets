<script lang="js">

// import { useAxiosLib } from '@/stores/axios';
import { ElButton, ElButtonGroup, ElForm, ElFormItem, ElInput, ElTable, ElTableColumn } from 'element-plus';

import { useStandardStore } from '@/stores/standards';


const theProductStore = useStandardStore()


export default {

  name: "standard-table-view",

  components: {},

  computed: {
    items() {
      return theProductStore.items;
    },
  },

  data() {
    const item = {}
    return { item }
  },

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

    <ElTable :data="items">
      <ElTableColumn label="ID" prop="id"></ElTableColumn>
      <ElTableColumn label="标准代号" prop="code"></ElTableColumn>
      <ElTableColumn label="标题" prop="title"></ElTableColumn>
      <ElTableColumn label="上传者" prop="owner"></ElTableColumn>

      <ElTableColumn>
        <template #default>
          <ElButton link> 详情 </ElButton>
        </template>
      </ElTableColumn>

    </ElTable>

  </div>
</template>
