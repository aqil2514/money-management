<script setup lang="ts">
import Controller from '~/components/category/controller/index.vue';
import CategoryData from '~/components/category/data/index.vue'
import { serverUrl } from '~/constants/server-url';
import type { CategoryDb } from '~/types/category';

const { data, status } = useFetch<{ message: string, data: CategoryDb[] }>(`${serverUrl}/category`, {
  key: "fetch-category"
})

useSeoMeta({
  title: "Kategori"
})

const categoryItems = data.value?.data ?? [];

</script>

<template>
  <div class="space-y-4">
    <Controller />
    <CategoryData />
  </div>
</template>