<script lang="js">
// import { ElButton } from 'element-plus'
// import { RouterLink } from 'vue-router'

import { UploadFilled } from '@element-plus/icons-vue'

export default {
  name: 'upload-media-file',

  components: { UploadFilled },

  methods: {
    init() {},
  },

  mounted() {
    this.init()
  },

  props: {},
}
</script>

<style scoped></style>

<template>
  <frame-for-developer>
    <h1>Upload Media File</h1>

    <el-upload class="upload-demo" drag action="/objects/upload" multiple>
      <el-icon class="el-icon--upload"><upload-filled /></el-icon>
      <div class="el-upload__text">Drop file here or <em>click to upload</em></div>
      <template #tip>
        <div class="el-upload__tip">jpg/png files with a size less than 500kb</div>
      </template>
    </el-upload>
  </frame-for-developer>
</template>
