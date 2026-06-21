<script lang="js">

// import { useAxiosLib } from '@/stores/axios';
// import { ElButton, ElForm, ElFormItem, ElInput } from 'element-plus';
// import { ElButton, ElInput } from 'element-plus';

import { useSessionStore } from '@/stores/sessions.js';
import { ElDialog } from 'element-plus';
import LoginBox from './session-login-box.vue'

const theSessionStore = useSessionStore()


export default {

    name: "sessions-mini-login-dialog",

    components: { LoginBox },

    data() {
        const username = 'foo'
        const password = 'bar'
        return { username, password }
    },

    methods: {


    },

    mounted() {
    },

    props: {}
}

</script>

<style></style>

<template>
    <ElDialog title="登录">

        <LoginBox />

    </ElDialog>
</template>
