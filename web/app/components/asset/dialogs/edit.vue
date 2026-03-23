<script lang="ts" setup>
import { serverUrl } from '~/constants/server-url';
import type { AssetSchemaType } from '~/schemas/asset-schema';
import { assetProvideKey, type AssetProvideInjectKey } from '~/types/injectKey';

const state = inject<AssetProvideInjectKey>(assetProvideKey)
const isOpen = computed(({
  get: () => state?.modalOpen.value === "edit",
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

const editHandler = async (values: AssetSchemaType) => {
  await $fetch(`${serverUrl}/asset/${state?.assetId.value}`, {
    method: "PUT",
    body: values
  })

  toast.add({
    title: "Berhasil",
    description: "Data berhasil diubah"
  })

  await refreshNuxtData("fetch-asset")
  state?.updateModalOpen(null)
}

</script>

<template>
  <UModal v-model:open="isOpen" title="Tambah Aset" description="Isi data dibawah ini untuk meambahkan aset baru">
    <template #body>
      <AssetForms :submit-handler="editHandler" :default-values="selectedData" />
    </template>
  </UModal>
</template>