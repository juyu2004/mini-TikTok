<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { useRoute } from 'vue-router'
import { api } from '../api'

const route = useRoute()
const user = ref<any>({})
const videos = ref<any[]>([])

onMounted(async () => {
  const uid = (route.params.id as string) || localStorage.getItem('user_id') || ''
  const { data } = await api.get('/user/profile', { params: { user_id: uid }})
  user.value = data
  const res = await api.get('/video/list', { params: { user_id: uid }})
  videos.value = res.data.items || []
})
</script>

<template>
  <div style="max-width:720px;margin:0 auto;padding:12px">
    <h2>{{ user.nickname }} <small style="color:#999">{{ user.email }}</small></h2>
    <p>粉丝：{{ user.followers }} · 关注：{{ user.following }}</p>
    <div class="grid">
      <div v-for="v in videos" :key="v.video_id" class="item">
        <video :src="v.playback_url" controls muted style="width:100%"></video>
        <div>{{ v.title }}</div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.grid { display:grid; grid-template-columns: repeat(3, 1fr); gap: 10px; }
.item { border: 1px solid #eee; border-radius: 8px; padding: 6px; }
</style>
