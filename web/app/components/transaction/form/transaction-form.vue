<script setup lang="ts">
import InputCategory from "./InputCategory.vue";
import InputDate from "./InputDate.vue"
const { schema, state, transactionDate, onSubmit } = useTransactionForm()
</script>

<template>
  <UForm :schema="schema" :state="state" @submit="onSubmit" class="space-y-4">
    <UFormField label="Tanggal Transaksi" name="date">
      <InputDate v-model="transactionDate" />
    </UFormField>

    <UFormField label="Nominal" name="nominal">
      <UInputNumber v-model="state.nominal" class="w-full" :format-options="{
        style: 'currency',
        currency: 'IDR',
        currencyDisplay: 'narrowSymbol'
      }" />
    </UFormField>

    <UFormField label="Kategori" name="category">
      <InputCategory v-model="state.category" />
    </UFormField>

    <UFormField label="Notes" name="notes">
      <UTextarea v-model="state.note" :maxrows="4" autoresize class="w-full" />
    </UFormField>

    <UButton type="submit" label="Simpan" block />
  </UForm>
</template>