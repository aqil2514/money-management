<script setup lang="ts">
import Controller from '~/components/transaction/controller.vue';
import Items from '~/components/transaction/items.vue';
import SummaryHeader from '~/components/transaction/summary-header.vue';

useSeoMeta({
  title: "Transaksi"
})
</script>

<template>
  <div class="space-y-4">
    <SummaryHeader />
    <Controller />
    <Items />
  </div>
</template>