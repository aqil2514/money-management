<script setup lang="ts">
import AssetController from "~/components/asset/controller/index.vue"
import AssetData from "~/components/asset/data/index.vue"

import { assetProvideKey } from "~/types/injectKey";

// DIALOGS
import AddAssetDialog from "~/components/asset/dialogs/add.vue"
import EditAssetDialog from "~/components/asset/dialogs/edit.vue"
import DeleteAssetDialog from "~/components/asset/dialogs/delete.vue"

useSeoMeta({
  title: "Aset"
})

const { currencyItem, liquidityScoreMapping, modalOpen, statusItem, updateModalOpen, fetcher, assetId } = useUseAssetPage()


provide(assetProvideKey, {
  modalOpen,
  updateModalOpen,
  assetId,

  constant: {
    liquidityScoreMapping,
    currencyItem,
    statusItem,
  },

  fetcher
})
</script>

<template>
  <div class="space-y-4">

    <h1 class="text-2xl font-semibold">Aset Finansial</h1>
    <AssetController />
    <AssetData />
  </div>

  <!-- DIALOGS -->
  <AddAssetDialog />
  <EditAssetDialog />
  <DeleteAssetDialog />
</template>