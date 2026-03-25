<script setup lang="ts">
import Controller from '~/components/transaction/controllers/index.vue';
import Items from '~/components/transaction/items.vue';
import SummaryHeader from '~/components/transaction/summary-header.vue';
import { transactionProviderKey } from '~/types/injectKey';

import DetailDialog from '~/components/transaction/dialogs/detail.vue'
import EditDialog from '~/components/transaction/dialogs/edit.vue'
import DeleteDialog from '~/components/transaction/dialogs/delete.vue'

useSeoMeta({
  title: "Transaksi"
})

const { fetcher, modal, resources } = await useTransactionPage()


provide(transactionProviderKey, {
  fetcher,
  modal,
  resources
})
</script>

<template>
  <div class="space-y-4">
    <SummaryHeader />
    <Controller />
    <Items />
  </div>

  <DetailDialog />
  <EditDialog />
  <DeleteDialog />
</template>