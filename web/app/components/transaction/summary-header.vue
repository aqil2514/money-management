<script setup lang="ts">
const summaries = [
  { label: "Pendapatan", nominal: 3000000, color: '#3b82f6' },
  { label: "Pengeluaran", nominal: 3000000, color: '#ef4444' },
  { label: "Total", nominal: 3000000, color: '#000000' },
]
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