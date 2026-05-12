import './assets/main.css'

import { createApp } from 'vue'


import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'

import App from './App.vue'
import router from './router'
import lib_pinia_js from '@/js/pinia.js'

import Frame4admin from '@/components/frames/admin/index.vue'
import Frame4dev from '@/components/frames/developer/index.vue'
import Frame4user from '@/components/frames/user/index.vue'

////////////////////////////////////////////////////////////////////////////////

const app = createApp(App)

const pinia = lib_pinia_js.GetPinia()

////////////////////////////////////////////////////////////////////////////////

app.use(pinia)
app.use(router)
app.use(ElementPlus)

app.component("frame-for-admin", Frame4admin)
app.component("frame-for-developer", Frame4dev)
app.component("frame-for-user", Frame4user)

app.mount('#app')
