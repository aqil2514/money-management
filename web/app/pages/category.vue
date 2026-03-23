<script setup lang="ts">
import Controller from '~/components/category/controller/index.vue';
import CategoryData from '~/components/category/data/index.vue'

useSeoMeta({
  title: "Kategori"
})

</script>

<template>
  <div class="space-y-4">
    <h1 class="text-2xl font-semibold">Kategori</h1>
    <USeparator />
    <Controller />
    <CategoryData />
  </div>
</template>