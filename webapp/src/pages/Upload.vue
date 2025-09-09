<script setup lang="ts">
import { ref } from 'vue'
import { api } from '../api'

const title = ref('')
const description = ref('')
const file = ref<File|null>(null)
const loading = ref(false)

async function submit() {
  if (!file.value) return alert('请选择文件')
  const fd = new FormData()
  fd.append('title', title.value)
  fd.append('description', description.value)
  fd.append('file', file.value)
  loading.value = true
  try {
    const { data } = await api.post('/video/upload', fd, {
      headers: { 'Authorization': 'Bearer ' + localStorage.getItem('token') }
    })
    alert('上传成功：' + data.video_id)
    location.href = '/profile'
  } catch (e:any) {
    alert('上传失败：' + (e.response?.data?.error || e.message))
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <div style="max-width:520px;margin:20px auto;">
    <h2>上传视频</h2>
    <el-input v-model="title" placeholder="标题" style="margin:6px 0" />
    <el-input v-model="description" placeholder="描述" style="margin:6px 0" />
    <input type="file" accept="video/*" @change="e=>file = (e.target as HTMLInputElement).files?.[0] || null" />
    <div style="margin-top:10px">
      <el-button type="primary" :loading="loading" @click="submit">上传</el-button>
    </div>
  </div>
</template>
