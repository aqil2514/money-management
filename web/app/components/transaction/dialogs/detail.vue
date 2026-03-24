<script lang="ts" setup>
import { transactionProviderKey } from '~/types/injectKey';
import type { TransactionDb } from '~/types/transaction';

const injectedData = inject(transactionProviderKey)
const transactionId = computed(() => injectedData?.modal.transactionId.value)
const selectedData = computed(() => injectedData?.fetcher.data.value?.data.find((d) => d.id === transactionId.value))
const internalData = ref<TransactionDb | null>(null)

const open = computed({
  get: () => injectedData?.modal.modalOpen.value === 'detail',
  set: (value) => {
    injectedData?.modal.updateModalOpen(value ? 'detail' : null)
  }
})

watch(selectedData, (newData) => {
  if (newData) {
    internalData.value = newData
  }
}, { immediate: true })

const getTypeColor = (type: string) => {
  switch (type) {
    case 'income':
      return 'success'
    case 'expense':
      return 'error'
    case 'transfer':
      return 'info'
    case 'payable':
      return 'warning'
    case 'receivable':
      return 'secondary'
    default:
      return 'neutral'
  }
}

</script>

<template>
  <UModal v-model:open="open" :title="internalData ? `Detail: ${internalData.note}` : 'Transaksi tidak ditemukan'"
    description="Detail transaksi">
    <template v-if="!internalData" #body>
      <div class="flex flex-col items-center justify-center py-10 text-gray-500">
        <UIcon name="i-heroicons-exclamation-circle" class="w-12 h-12 mb-2" />
        <p>Data transaksi tidak ditemukan</p>
      </div>
    </template>

    <template v-else #body>
      <div class="space-y-6">
        <div class="flex items-center justify-between p-4 rounded-xl bg-gray-50 dark:bg-gray-800/50">
          <div>
            <p class="text-sm text-gray-500 font-medium">Nominal</p>
            <h3 :class="[
              'text-2xl font-bold',
              internalData.type === 'income' ? 'text-green-600' : 'text-red-600'
            ]">
              {{ internalData.type === 'income' ? '+' : '-' }}
              {{ new Intl.NumberFormat('id-ID', { style: 'currency', currency: 'IDR' }).format(internalData.nominal) }}
            </h3>
          </div>
          <UBadge :color="getTypeColor(internalData.type)" variant="subtle" class="capitalize">
            {{ internalData.type }}
          </UBadge>
        </div>

        <div class="grid grid-cols-2 gap-4">
          <div class="space-y-1">
            <p class="text-xs text-gray-500 uppercase tracking-wider font-semibold">Kategori</p>
            <div class="flex items-center gap-2">
              <UIcon name="i-heroicons-tag" class="text-primary" />
              <span class="text-sm">{{ internalData.category }}</span>
              <span v-if="internalData.subCategory" class="text-xs text-gray-400">/ {{ internalData.subCategory
              }}</span>
            </div>
          </div>

          <div class="space-y-1">
            <p class="text-xs text-gray-500 uppercase tracking-wider font-semibold">Tanggal</p>
            <div class="flex items-center gap-2 text-sm">
              <UIcon name="i-heroicons-calendar" />
              {{ new Date(internalData.date).toLocaleDateString('id-ID', { dateStyle: 'long' }) }}
            </div>
          </div>
        </div>

        <UDivider />

        <div class="space-y-3">
          <p class="text-xs text-gray-500 uppercase tracking-wider font-semibold">Aliran Dana</p>
          <div class="flex items-center gap-4 bg-gray-50 dark:bg-gray-800/50 p-3 rounded-lg">
            <div class="flex-1 text-center">
              <p class="text-[10px] text-gray-400 uppercase">Dari</p>
              <p class="font-medium text-sm">{{ internalData.assetFrom }}</p>
            </div>

            <UIcon
              :name="internalData.type === 'transfer' ? 'i-heroicons-arrow-right-circle' : 'i-heroicons-minus-circle'"
              class="w-6 h-6 text-gray-300" />

            <div class="flex-1 text-center" v-if="internalData.type === 'transfer'">
              <p class="text-[10px] text-gray-400 uppercase">Ke</p>
              <p class="font-medium text-sm text-primary">{{ internalData.assetTo }}</p>
            </div>
            <div class="flex-1 text-center" v-else>
              <p class="text-[10px] text-gray-400 uppercase">Status</p>
              <p class="font-medium text-sm italic capitalize">{{ internalData.type }}</p>
            </div>
          </div>
        </div>

        <div v-if="internalData.description || internalData.isHaveTransferFee" class="space-y-3 pt-2">
          <div v-if="internalData.description" class="bg-blue-50 dark:bg-blue-900/20 p-3 rounded-lg">
            <p class="text-xs text-blue-600 dark:text-blue-400 font-semibold mb-1">Deskripsi</p>
            <p class="text-sm text-gray-600 dark:text-gray-300 italic">"{{ internalData.description }}"</p>
          </div>

          <div v-if="internalData.isHaveTransferFee"
            class="flex justify-between items-center text-sm border-t pt-3 border-dashed">
            <span class="text-gray-500">Biaya Transfer (via {{ internalData.feeFromAsset }})</span>
            <span class="font-medium text-red-500">-{{ new Intl.NumberFormat('id-ID', {
              style: 'currency', currency:
                'IDR'
            }).format(internalData.transferFee) }}</span>
          </div>
        </div>
      </div>
    </template>

    <template #footer>
      <div class="flex justify-end gap-2">
        <UButton color="neutral" variant="ghost" label="Tutup" @click="open = false" />
        <UButton color="primary" icon="i-heroicons-pencil-square"
          @click="injectedData?.modal.updateModalOpen('edit', internalData?.id)" label="Edit Transaksi" />
      </div>
    </template>
  </UModal>
</template>