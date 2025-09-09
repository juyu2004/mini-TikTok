import { createRouter, createWebHistory } from 'vue-router'
import Login from './pages/Login.vue'
import Home from './pages/Home.vue'
import Profile from './pages/Profile.vue'
import Upload from './pages/Upload.vue'

export default createRouter({
  history: createWebHistory(),
  routes: [
    { path: '/', component: Home },
    { path: '/login', component: Login },
    { path: '/profile/:id?', component: Profile },
    { path: '/upload', component: Upload },
  ]
})
