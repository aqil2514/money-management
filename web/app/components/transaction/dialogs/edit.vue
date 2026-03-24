<script setup lang="ts">
import { transactionProviderKey } from '~/types/injectKey';
import type { TransactionRequestPayload, TransactionSchemaType } from '~/schemas/transaction-schema';
import { serverUrl } from '~/constants/server-url';
import TransactionForm from '../form/transaction-form.vue';
import type { TransactionDb } from '~/types/transaction';
import { CalendarDate, Time } from '@internationalized/date';

const injectedData = inject(transactionProviderKey)
const toast = useToast()

const transactionId = computed(() => injectedData?.modal.transactionId.value)
const selectedData = computed(() => injectedData?.fetcher.data.value?.data.find((d) => d.id === transactionId.value))
const internalData = ref<TransactionDb | null>(null)

watch(selectedData, (newData) => {
  if (newData) {
    internalData.value = newData
  }
}, { immediate: true })

const open = computed({
  get: () => injectedData?.modal.modalOpen.value === 'edit',
  set: (value) => {
    injectedData?.modal.updateModalOpen(value ? 'edit' : null)
  }
})

const editHandler = async (values: TransactionRequestPayload) => {
  // await $fetch(`${serverUrl}/transactions`, {
  //   method: "POST",
  //   body: values
  // })

  console.log(values)

  toast.add({
    title: "Berhasil",
    description: "Data transaksi berhasil diedit"
  })
  // injectedData?.modal.updateModalOpen(null);
  // await refreshNuxtData("transaction")
}

const defaultValues = computed<TransactionSchemaType | undefined>(() => {
  if (!internalData.value) return undefined;

  const data = internalData.value;
  const d = new Date(data.date);

  return {
    type: data.type,
    nominal: data.nominal,
    note: data.note,
    description: data.description || "",

    date: new CalendarDate(d.getFullYear(), d.getMonth() + 1, d.getDate()),
    time: new Time(d.getHours(), d.getMinutes()),

    categoryId: data.categoryId,
    subCategoryId: data.subCategoryId || "no-subcategory",
    assetFromId: data.assetFromId,

    assetToId: data.assetToId || undefined,
    isHaveTransferFee: data.isHaveTransferFee,
    transferFee: data.transferFee || 0,
    feeFromAssetId: data.feeFromAssetId || undefined,
    debtor: data.debtor || "",
    creditor: data.creditor || "",
  } as unknown as TransactionSchemaType;
});

</script>

<template>
  <UModal v-model:open="open" :title="'Edit Transaksi'"
    :description="'Ubah form di bawah ini untuk mengedit transaksi'">

    <template #body>
      <TransactionForm v-on:submit="editHandler" :default-values="defaultValues" />
    </template>

  </UModal>
</template>