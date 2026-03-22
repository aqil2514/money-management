<script lang="ts" setup>
import type { DropdownMenuItem } from '@nuxt/ui';
import type { CategoryDb } from '~/types/category';
import EditController from '~/components/category/controller/edit.vue'

const props = defineProps<{ categoryItems: CategoryDb[] }>()
const selectedData = ref<CategoryDb | null>(null)

const getDropdownItems = (item: CategoryDb): DropdownMenuItem[][] => ([
  [
    {
      label: item.name,
      type: "label"
    }
  ],
  [
    {
      label: "Edit",
      onSelect: () => selectedData.value = item
    }
  ]
])
</script>

<template>
  <div v-if="categoryItems.length === 0">
    <p>Kategori belum ada yang dibuat</p>
  </div>
  <div v-else class="grid grid-cols-3 gap-4">
    <UCard v-for="item in categoryItems">
      <template #header>
        <div class="space-y-2">
          <div class="flex justify-end">
            <UDropdownMenu :items="getDropdownItems(item)" :ui="{
              content: 'w-48'
            }">
              <UButton icon="i-lucide-menu" color="neutral" variant="outline" />
            </UDropdownMenu>
          </div>
          <p class="text-xl font-semibold">
            {{ item.name }}
          </p>
          <p class="text-sm text-muted">{{ item.description }}</p>
        </div>
      </template>
    </UCard>
  </div>

  <EditController :data="selectedData" @close="() => selectedData = null" />
</template>