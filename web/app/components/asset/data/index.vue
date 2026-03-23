<script lang="ts" setup>
import { assetProvideKey } from '~/types/injectKey';

const injectedData = inject(assetProvideKey)
const assetItems = computed(() => injectedData?.fetcher.data.value?.data ?? [])
const isLoading = computed(() => injectedData?.fetcher.status.value === 'pending')
const liquidityMapping = injectedData?.constant.liquidityScoreMapping

// Helper untuk format angka ke Rupiah/Dollar
const formatAmount = (amount: string, currency: string) => {
  return new Intl.NumberFormat('id-ID', {
    style: 'currency',
    currency: currency.toUpperCase(),
    maximumFractionDigits: 0
  }).format(parseFloat(amount))
}
</script>

<template>
  <div v-if="isLoading" class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
    <USkeleton v-for="n in 6" :key="n" class="w-full h-48" />
  </div>

  <div v-else-if="assetItems.length === 0" class="text-center py-20 border-2 border-dashed rounded-xl opacity-50">
    <UIcon name="i-heroicons-circle-stack" class="w-12 h-12 mx-auto mb-2" />
    <p>Belum ada aset finansial yang tercatat.</p>
  </div>

  <div v-else class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
    <UCard v-for="item in assetItems" :key="item.ID" class="flex flex-col">
      <template #header>
        <div class="flex justify-between items-start mb-2">
          <div>
            <p class="text-lg font-bold leading-tight">{{ item.name }}</p>
            <p class="text-xs text-gray-500 uppercase tracking-wide">{{ item.category }}</p>
          </div>
          <UBadge :color="item.status === 'active' ? 'success' : 'neutral'" variant="subtle" class="capitalize">
            {{ item.status }}
          </UBadge>
        </div>

        <p v-if="item.description" class="text-sm text-gray-600 dark:text-gray-400 line-clamp-1 italic">
          "{{ item.description }}"
        </p>

        <div class="flex flex-wrap gap-1 mt-3">
          <UBadge color="secondary" variant="soft">{{ item.ownerType }}</UBadge>
          <UBadge color="primary" variant="soft">{{ item.assetType }}</UBadge>
          <UBadge color="neutral" variant="outline" class="uppercase">{{ item.currency }}</UBadge>
        </div>
      </template>

      <div class="space-y-4">
        <div class="text-center py-2">
          <p class="text-sm text-gray-500 mb-1">Total Nilai Aset</p>
          <p class="text-2xl font-black text-primary">
            {{ formatAmount(item.total, item.currency) }}
          </p>
        </div>

        <UAlert variant="subtle"
          :color="item.liquidityScore >= 4 ? 'success' : (item.liquidityScore === 3 ? 'primary' : 'warning')"
          icon="i-heroicons-information-circle" title="Likuiditas"
          :description="liquidityMapping?.[item.liquidityScore] || 'Skor tidak diketahui'" />
      </div>

      <template #footer>
        <div class="flex justify-end gap-2">
          <UButton icon="i-heroicons-pencil-square" color="neutral" variant="ghost" size="sm"
            @click="injectedData?.updateModalOpen('edit', item.ID)" />
          <UButton icon="i-heroicons-trash" color="error" variant="ghost" size="sm"
            @click="injectedData?.updateModalOpen('delete')" />
        </div>
      </template>
    </UCard>
  </div>
</template>