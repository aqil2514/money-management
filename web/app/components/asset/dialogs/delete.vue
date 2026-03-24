<script lang="ts" setup>
import { serverUrl } from '~/constants/server-url';
import type { AssetSchemaType } from '~/schemas/asset-schema';
import { assetProvideKey, type AssetProvideInjectKey } from '~/types/injectKey';

const state = inject<AssetProvideInjectKey>(assetProvideKey)
const isOpen = computed(({
  get: () => state?.modalOpen.value === "delete",
  set: (value) => {
    if (!value) state?.updateModalOpen(null)
  }
}))
const toast = useToast()

const selectedData = computed<AssetSchemaType | undefined>(() => {
  if (isOpen.value && state?.fetcher.data.value && state.assetId.value) {
    const raw = state.fetcher.data.value.data.find((d) => d.ID === state.assetId.value)

    if (raw) {
      return {
        name: raw.name,
        total: Number(raw.total),
        category: raw.category,
        status: raw.status,
        description: raw.description,
        ownerType: raw.ownerType,
        assetType: raw.assetType,
        liquidityScore: raw.liquidityScore,
        currency: raw.currency
      }
    }
  }
  return undefined
})

const deleteHandler = async () => {
  if (!state?.assetId.value) return

  await $fetch(`${serverUrl}/asset/${state.assetId.value}`, {
    method: "DELETE"
  })

  toast.add({
    title: "Berhasil",
    description: "Data berhasil dihapus"
  })

  await refreshNuxtData("fetch-asset")
  state?.updateModalOpen(null)
}

</script>

<template>
  <UModal v-model:open="isOpen" title="Hapus Aset"
    description="Apakah Anda yakin ingin menghapus aset ini? Tindakan ini tidak dapat dibatalkan.">
    <template #body>
      <div v-if="selectedData" class="space-y-3 text-sm">
        <div class="flex justify-between border-b pb-2">
          <span class="text-gray-500">Nama Aset</span>
          <span class="font-medium">{{ selectedData.name }}</span>
        </div>
        <div class="flex justify-between border-b pb-2">
          <span class="text-gray-500">Kategori</span>
          <span class="font-medium">{{ selectedData.category }}</span>
        </div>
        <div class="flex justify-between border-b pb-2">
          <span class="text-gray-500">Total</span>
          <span class="font-medium">{{ selectedData.currency }} {{ selectedData.total.toLocaleString() }}</span>
        </div>
        <div v-if="selectedData.description" class="pt-1">
          <span class="text-gray-500 block">Deskripsi</span>
          <p class="mt-1 text-gray-700 italic">"{{ selectedData.description }}"</p>
        </div>
      </div>
      <div v-else class="py-4 text-center">
        <p>Memuat data...</p>
      </div>
    </template>

    <template #footer>
      <div class="flex justify-end gap-3">
        <UButton color="neutral" variant="ghost" label="Batal" @click="isOpen = false" />
        <UButton color="error" label="Ya, Hapus Aset" @click="deleteHandler" />
      </div>
    </template>
  </UModal>
</template>