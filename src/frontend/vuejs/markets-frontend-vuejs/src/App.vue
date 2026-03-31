<script setup>
import { RouterLink, RouterView } from 'vue-router'
import HelloWorld from './components/HelloWorld.vue'
</script>

<style scoped lang="css">
.router-link {
  padding: 3px;
  margin: 3px;
}
</style>

<template>
  <div>
    <div>
      <nav>
        <RouterLink class="router-link" to="/">Home</RouterLink>
        <RouterLink class="router-link" to="/about">About</RouterLink>

        <RouterLink class="router-link" to="/dev">Dev</RouterLink>
        <RouterLink class="router-link" to="/home">User</RouterLink>
        <RouterLink class="router-link" to="/admin">Admin</RouterLink>

      </nav>
    </div>
    <div>
      <RouterView />
    </div>
  </div>
</template>
