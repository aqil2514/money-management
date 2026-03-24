<script setup lang="ts">
import { transactionProviderKey } from '~/types/injectKey';
import formatToRupiah from '~/utils/formatter/format-to-rupiah';

const injectedState = inject(transactionProviderKey)
const transaction = computed(() => injectedState?.fetcher.data.value?.data ?? [])

const income = computed(() => transaction.value.reduce((acc, curr) => {
  if (curr.type === 'income') {
    acc.nominal += curr.nominal
  }

  return acc
}, { label: "Pendapatan", nominal: 0, color: '#3b82f6' }))

const expense = computed(() => transaction.value.reduce((acc, curr) => {
  if (curr.type === 'expense') {
    acc.nominal += curr.nominal
  }

  return acc
}, { label: "Pengeluaran", nominal: 0, color: '#ef4444' }))

const total = computed(() => {
  const nominal = income.value.nominal - expense.value.nominal
  return {
    label: "Total",
    nominal,
    color: "#000000"
  }
})

const summaries = computed(() => [
  income.value,
  expense.value,
  total.value
])
</script>

<template>
  <UCard>
    <div class="flex flex-row justify-between gap-4">
      <div v-for="summary in summaries" :key="summary.label" class="text-center">
        <p class="font-semibold">{{ summary.label }}</p>
        <p :style="{ color: summary.color }" class="text-lg font-semibold">
          {{ formatToRupiah(summary.nominal) }}
        </p>
      </div>
    </div>
  </UCard>
</template>