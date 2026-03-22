<script setup lang="ts">
import type { DropdownMenuItem } from '@nuxt/ui';
import { serverUrl } from '~/constants/server-url';
import type { CategoryDb } from '~/types/category';
import Card from './card.vue';

const { data, status } = useFetch<{ message: string, data: CategoryDb[] }>(`${serverUrl}/category`, {
  key: "fetch-category"
})

useSeoMeta({
  title: "Kategori"
})

const categoryItems = computed(() => data.value?.data ?? []);

const emptyItems = Array.from({ length: 12 })
</script>

<template>
  <div v-if="status === 'pending'">
    <div class="grid grid-rows-3 gap-4">
      <USkeleton v-for="item in emptyItems" class="w-full" />
    </div>
  </div>
  <Card :category-items="categoryItems" />
</template>