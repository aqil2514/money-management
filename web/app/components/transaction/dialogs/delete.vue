<script lang="ts" setup>
import { serverUrl } from '~/constants/server-url';
import { transactionProviderKey } from '~/types/injectKey';
import type { TransactionDb } from '~/types/transaction';

const injectedData = inject(transactionProviderKey)
const toast = useToast()

const transactionId = computed(() => injectedData?.modal.transactionId.value)
const selectedData = computed(() =>
  injectedData?.fetcher.data.value?.data?.find((d) => d.id === transactionId.value)
)
const internalData = ref<TransactionDb | null>(null)

const loading = ref(false)

const open = computed({
  get: () => injectedData?.modal.modalOpen.value === 'delete',
  set: (value) => {
    injectedData?.modal.updateModalOpen(value ? 'delete' : null)
  }
})

watch(selectedData, (newData) => {
  if (newData) {
    internalData.value = newData
  }
}, { immediate: true })

const handleDelete = async () => {
  if (!transactionId.value) return

  loading.value = true
  try {
    await $fetch(`${serverUrl}/transactions/${transactionId.value}`, {
      method: 'DELETE'
    })

    toast.add({
      title: 'Berhasil',
      description: 'Transaksi telah dihapus dan saldo telah diperbarui',
      color: 'success'
    })

    await refreshNuxtData('transaction')
    injectedData?.modal.updateModalOpen(null)
  } catch (error: any) {
    toast.add({
      title: 'Gagal Menghapus',
      description: error.data?.message || 'Terjadi kesalahan pada server',
      color: 'error'
    })
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <UModal v-model:open="open" title="Konfirmasi Hapus" description="Tindakan ini tidak dapat dibatalkan">
    <template #body>
      <div v-if="internalData" class="space-y-4">
        <div class="p-4 rounded-lg bg-error-50 dark:bg-error-950/20 border border-error-100 dark:border-error-900">
          <div class="flex gap-3">
            <UIcon name="i-heroicons-exclamation-triangle" class="w-5 h-5 text-error-600 shrink-0" />
            <div>
              <p class="text-sm text-error-800 dark:text-error-200 font-semibold">
                Apakah Anda yakin ingin menghapus transaksi ini?
              </p>
              <p class="text-xs text-error-700 dark:text-error-300 mt-1">
                Menghapus transaksi ini akan secara otomatis mengembalikan (undo) saldo pada aset
                <strong>{{ internalData.assetFrom }}</strong>.
              </p>
            </div>
          </div>
        </div>

        <div class="border rounded-lg p-3 text-sm space-y-2 bg-gray-50 dark:bg-gray-800/50">
          <div class="flex justify-between">
            <span class="text-gray-500">Catatan:</span>
            <span class="font-medium">{{ internalData.note || '-' }}</span>
          </div>
          <div class="flex justify-between">
            <span class="text-gray-500">Nominal:</span>
            <span class="font-bold text-red-600">
              {{ new Intl.NumberFormat('id-ID', {
                style: 'currency', currency: 'IDR', maximumFractionDigits: 0
              }).format(internalData.nominal) }}
            </span>
          </div>
          <div class="flex justify-between">
            <span class="text-gray-500">Tanggal:</span>
            <span>{{ new Date(internalData.date).toLocaleDateString('id-ID', { dateStyle: 'medium' }) }}</span>
          </div>
        </div>
      </div>
    </template>

    <template #footer>
      <div class="flex justify-end gap-3">
        <UButton color="neutral" variant="ghost" label="Batal" :disabled="loading" @click="open = false" />
        <UButton color="error" icon="i-heroicons-trash" label="Hapus Permanen" :loading="loading"
          @click="handleDelete" />
      </div>
    </template>
  </UModal>
</template>