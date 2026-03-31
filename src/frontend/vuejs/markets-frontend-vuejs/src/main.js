import './assets/main.css'

import { createApp } from 'vue'
import { createPinia } from 'pinia'

import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'

import App from './App.vue'
import router from './router'

import Frame4admin from '@/components/frames/admin/index.vue'
import Frame4dev from '@/components/frames/developer/index.vue'
import Frame4user from '@/components/frames/user/index.vue'



const app = createApp(App)

app.use(createPinia())
app.use(router)
app.use(ElementPlus)

app.component("frame-for-admin", Frame4admin)
app.component("frame-for-developer", Frame4dev)
app.component("frame-for-user", Frame4user)

app.mount('#app')
