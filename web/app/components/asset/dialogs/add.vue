<script lang="ts" setup>
import { assetProvideKey, type AssetProvideInjectKey } from '~/types/injectKey';

import AssetForm from "~/components/asset/forms/index.vue"
import type { AssetSchemaType } from '~/schemas/asset-schema';
import { serverUrl } from '~/constants/server-url';

const state = inject<AssetProvideInjectKey>(assetProvideKey)
const isOpen = computed(({
  get: () => state?.modalOpen.value === "add",
  set: (value) => {
    if (!value) state?.updateModalOpen(null)
  }
}))
const toast = useToast()

const submitHandler = async (values: AssetSchemaType) => {
  await $fetch(`${serverUrl}/asset`, {
    body: values,
    method: "POST"
  })

  toast.add({
    title: "Berhasil",
    description: "Data aset berhasil ditambah"
  })
  state?.updateModalOpen(null)
  await refreshNuxtData("fetch-asset")
}

</script>

<template>
  <UModal v-model:open="isOpen" title="Tambah Aset" description="Isi data dibawah ini untuk meambahkan aset baru">
    <template #body>
      <AssetForm :submit-handler="submitHandler" />
    </template>
  </UModal>
</template>