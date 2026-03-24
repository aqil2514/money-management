<script setup lang="ts">
import Controller from '~/components/transaction/controller.vue';
import Items from '~/components/transaction/items.vue';
import SummaryHeader from '~/components/transaction/summary-header.vue';
import { transactionProviderKey } from '~/types/injectKey';

import DetailDialog from '~/components/transaction/dialogs/detail.vue'

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
</template>