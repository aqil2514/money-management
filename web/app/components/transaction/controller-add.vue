<script setup lang="ts">
import TransactionForm from './form/transaction-form.vue';

const open = ref(false)
</script>

<template>
  <UButton class="cursor-pointer mt-4" @click="open = true">
    Tambah
  </UButton>

  <UModal v-model:open="open" :title="'Tambah Transaksi'"
    :description="'Isi form di bawah ini untuk menambah transaksi'">

    <template #body>
      <TransactionForm />
    </template>

  </UModal>
</template>