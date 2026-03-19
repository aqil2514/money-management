<script setup lang="ts">
const userName = ref("Muhamad Aqil Maulana")
</script>

<template>
  <UPopover>

    <UButton :avatar="{
      icon: 'i-lucide-user',
      loading: 'lazy'
    }" variant="outline" class="cursor-pointer">
      {{ userName }}
    </UButton>

    <template #content>
      <div class="p-4">
        <p>Ini adalah isi popover Anda</p>
      </div>
    </template>
  </UPopover>
</template>