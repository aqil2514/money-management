<script setup lang="ts">
import { transactionProviderKey } from '~/types/injectKey';
import type { TransactionRequestPayload } from '~/schemas/transaction-schema';
import { serverUrl } from '~/constants/server-url';
import TransactionForm from '../form/transaction-form.vue';

const injectedData = inject(transactionProviderKey)
const toast = useToast()

const open = computed({
  get: () => injectedData?.modal.modalOpen.value === 'add',
  set: (value) => {
    injectedData?.modal.updateModalOpen(value ? 'add' : null)
  }
})

const addHandler = async (values: TransactionRequestPayload) => {
  await $fetch(`${serverUrl}/transactions`, {
    method: "POST",
    body: values
  })

  toast.add({
    title: "Berhasil",
    description: "Data transaksi berhasil ditambah"
  })
  injectedData?.modal.updateModalOpen(null);
  await refreshNuxtData("transaction")
}

</script>

<template>
  <UButton class="cursor-pointer mt-4" @click="open = true">
    Tambah
  </UButton>

  <UModal v-model:open="open" :title="'Tambah Transaksi'"
    :description="'Isi form di bawah ini untuk menambah transaksi'">

    <template #body>
      <TransactionForm v-on:submit="addHandler" />
    </template>

  </UModal>
</template>