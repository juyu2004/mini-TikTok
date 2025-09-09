<script setup lang="ts">
import { ref } from 'vue'
import { api } from '../api'

const email = ref('test@example.com')
const password = ref('Passw0rd!')
const loading = ref(false)
async function login() {
  loading.value = true
  try {
    const { data } = await api.post('/auth/login', { email: email.value, password: password.value })
    localStorage.setItem('token', data.token)
    localStorage.setItem('user_id', data.user_id)
    location.href = '/'
  } catch (e:any) {
    alert('登录失败：' + (e.response?.data?.error || e.message))
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <div style="max-width:360px;margin:60px auto;">
    <h2>登录</h2>
    <el-input v-model="email" placeholder="邮箱" style="margin:8px 0" />
    <el-input v-model="password" type="password" placeholder="密码" style="margin:8px 0" />
    <el-button type="primary" :loading="loading" @click="login" style="width:100%">登录</el-button>
  </div>
</template>
