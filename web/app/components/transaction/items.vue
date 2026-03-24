<script setup lang="ts">
import { transactionProviderKey } from '~/types/injectKey';
import formatDateParts from '~/utils/formatter/format-date-parts';
import formatToRupiah from '~/utils/formatter/format-to-rupiah';
import mapTransactionDbToTransactionItems from '~/utils/mapper/map-transaction-db-to-transaction-items';


const injectedState = inject(transactionProviderKey);
const data = computed(() => injectedState?.fetcher.data.value?.data ?? [])
const status = computed(() => injectedState?.fetcher.status.value ?? "pending")

const mappedData = computed(() => {
  return data.value.map(mapTransactionDbToTransactionItems) ?? []
})

const grouppedByDate = computed(() => {
  if (!mappedData.value) return {}

  return mappedData.value.reduce((acc, item) => {
    const dateKey = item.date.split("T")[0]!

    if (!acc[dateKey]) {
      acc[dateKey] = []
    }

    acc[dateKey].push(item)

    return acc
  }, {} as Record<string, typeof mappedData.value>)
})

const transactionEntries = computed(() => Object.entries(grouppedByDate.value).sort().reverse())

const items = computed(() => {
  return transactionEntries.value.map(([date, body]) => {
    const safeBody = body || []

    const header = safeBody.reduce((acc, curr) => {
      if (curr.type === "income") acc.income += curr.nominal
      else acc.expense += curr.nominal
      return acc
    }, { date, income: 0, expense: 0 })

    const dateParts = formatDateParts(header)

    return {
      header: {
        ...header,
        display: dateParts
      },
      body: safeBody
    }
  })
})

</script>

<template>
  <div v-if="status === 'pending'">
    <p>Loading...</p>
  </div>
  <UAccordion v-else :items="items" :ui="{
    root: 'space-y-4 mt-4',
    header: 'bg-white px-4 cursor-pointer flex ',
    content: 'bg-white p-4',
    label: 'w-full'
  }">
    <template #default="{ item: { header } }">
      <div class="flex items-center justify-between w-full">
        <div class="flex gap-2 items-center">
          <p class="font-bold text-xl text-gray-800 tracking-tighter">{{ header.display.date }}</p>

          <p class="font-bold text-[10px] text-white px-2 py-0.5 rounded-md shadow-sm transition-colors" :class="{
            'bg-blue-600': header.display.day === 'Sabtu',
            'bg-red-600': header.display.day === 'Minggu',
            'bg-gray-500': !['Sabtu', 'Minggu'].includes(header.display.day)
          }">
            {{ header.display.day }}
          </p>

          <p class="text-[11px] font-bold text-gray-400 tabular-nums">
            {{ header.display.month }}.{{ header.display.year }}
          </p>
        </div>

        <div class="grid grid-cols-2 gap-8 min-w-70">
          <div class="flex flex-col items-end">
            <span
              class="text-[9px] uppercase font-black text-gray-400 tracking-widest leading-none mb-1">Pemasukan</span>
            <p class="text-sm font-black text-blue-500 tabular-nums">
              {{ formatToRupiah(header.income) }}
            </p>
          </div>

          <div class="flex flex-col items-end mr-4">
            <span
              class="text-[9px] uppercase font-black text-gray-400 tracking-widest leading-none mb-1">Pengeluaran</span>
            <p class="text-sm font-black text-red-500 tabular-nums">
              {{ formatToRupiah(header.expense) }}
            </p>
          </div>
        </div>
      </div>
    </template>

    <template #content="{ item: { body } }">
      <div class="bg-gray-50/30 divide-y divide-gray-100">
        <div v-for="(trx, index) in body" :key="index"
          class="grid grid-cols-3 gap-4 p-4 items-center hover:bg-white transition-colors cursor-pointer group">
          <div class="flex flex-col">
            <span class="text-sm font-bold text-gray-700 group-hover:text-blue-600 transition-colors">
              {{ trx.category }}
            </span>
            <span class="text-[10px] uppercase tracking-wider text-gray-400 font-medium">
              {{ trx.subCategory }}
            </span>
          </div>

          <div class="flex flex-col items-center text-center">
            <p class="text-sm text-gray-600 line-clamp-1 italic font-medium">
              "{{ trx.note || '-' }}"
            </p>
            <div class="flex items-center gap-1 mt-1 px-2 py-0.5 bg-gray-100 rounded-md">
              <UIcon name="i-heroicons-wallet" class="w-3 h-3 text-gray-500" />
              <span class="text-[10px] text-gray-500 font-bold uppercase">{{ trx.asetName }}</span>
            </div>
          </div>

          <div class="flex flex-col items-end">
            <p class="text-sm font-black tracking-tight"
              :class="trx.type === 'income' ? 'text-blue-600' : 'text-red-600'">
              {{ trx.type === 'income' ? '+' : '-' }} {{ formatToRupiah(trx.nominal) }}
            </p>
            <p class="text-[10px] text-gray-400 mt-1 font-mono">
              {{ trx.date?.split('T')[1]?.substring(0, 5) || '00:00' }}
            </p>
          </div>
        </div>
      </div>
    </template>
  </UAccordion>

</template>
