<script lang="ts" setup>
import type { SelectItem } from '@nuxt/ui';
import { serverUrl } from '~/constants/server-url';
import type { CategoryDb } from '~/types/category';

const model = defineModel<string | null | undefined>()
const { data, status } = useFetch<{ message: string, data: CategoryDb[] }>(`${serverUrl}/category/parents`)

const parentCategory = computed(() => data.value?.data ?? []);
const selectItems = computed<SelectItem[]>(() => parentCategory.value.map((val) => ({
  label: val.name,
  value: val.id
})))

</script>

<template>
  <div v-if="status === 'pending'">
    <p class="block font-medium text-default">Dari Kategori</p>
    <USkeleton class="w-full" />
  </div>
  <UFormField v-else label="Dari Kategori" name="parent">
    <USelect v-model="model" :items="[selectItems]" class="w-full" />
  </UFormField>
</template>