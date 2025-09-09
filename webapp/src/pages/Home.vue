<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { api } from '../api'

interface Video { video_id: string; title: string; description: string; playback_url: string; likes: number; comments: number }
const list = ref<Video[]>([])
onMounted(async () => {
  const { data } = await api.get('/video/feed')
  list.value = data.items || []
})
</script>

<template>
  <div class="feed">
    <div v-for="v in list" :key="v.video_id" class="card">
      <video :src="v.playback_url" controls playsinline style="width:100%;max-height:60vh"></video>
      <div class="meta">
        <h3>{{ v.title }}</h3>
        <p>{{ v.description }}</p>
      </div>
    </div>
  </div>
</template>

<style scoped>
.feed { max-width: 540px; margin: 0 auto; padding: 12px; }
.card { margin: 16px 0; border: 1px solid #eee; border-radius: 10px; padding: 10px; }
.meta { padding: 6px; }
</style>
