<script setup lang="ts">
import UserMenu from '~/components/UserMenu.vue';
import { NAVIGATION_ITEMS } from '~/constants/navigation-item';

const items = ref(NAVIGATION_ITEMS)
</script>

<template>
  <div class="min-h-screen bg-gray-50">
    <nav class="px-4 bg-white shadow-sm flex items-center justify-between">
      <h1>Money Management App</h1>
      <UNavigationMenu :items="items" />

      <user-menu />
    </nav>

    <main class="p-4 min-h-screen">
      <slot />
    </main>

    <footer class="p-4 text-center">
      © 2026 My Nuxt Project
    </footer>
  </div>
</template>