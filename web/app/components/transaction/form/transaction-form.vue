<script setup lang="ts">
import { defaultTransaction, transactionSchema, type TransactionRequestPayload, type TransactionSchemaType } from "~/schemas/transaction-schema";
import InputCategory from "./InputCategory.vue";
import InputDate from "./InputDate.vue"
import InputTime from "./InputTime.vue";
import InputTransactionType from "./InputTransactionType.vue";
import InputCreditor from "./InputCreditor.vue";
import InputDebtor from "./InputDebtor.vue";
import InputNote from "./InputNote.vue";
import InputIsHaveTransferFee from "./InputIsHaveTransferFee.vue";
import InputCurrency from "./InputCurrency.vue";
import InputAsset from "./InputAsset.vue";
import type { FormSubmitEvent } from "@nuxt/ui";
import { CalendarDate, Time } from "@internationalized/date";

const props = defineProps<{
  onSubmit: (values: TransactionRequestPayload) => Promise<void> | void
  defaultValues?: TransactionSchemaType
}>()

const now = new Date();

const state = reactive({
  ...defaultTransaction,
  time: new Time(now.getHours(), now.getMinutes()),
}) as TransactionSchemaType;
const isLoading = ref(false)
const toast = useToast()

watch(() => state.categoryId, (newVal, oldVal) => {
  if (oldVal)
    state.subCategoryId = 'no-subcategory'
})

watch(() => state.type, (newVal, oldVal) => {
  if (oldVal)
    state.categoryId = 'no-category'
})

const submitHandler = async (event: FormSubmitEvent<TransactionSchemaType>) => {
  const rawDate: CalendarDate = event.data.date;
  const rawTime: Time = event.data.time;

  const createdDate = new Date(
    rawDate.year,
    rawDate.month - 1,
    rawDate.day,
    rawTime.hour,
    rawTime.minute,
  );

  const payload: TransactionRequestPayload = {
    ...event.data,
    date: createdDate.toISOString(),
    subCategoryId: event.data.subCategoryId === "no-subcategory" ? undefined : event.data.subCategoryId,
    categoryId: event.data.categoryId === "no-category" ? "" : event.data.categoryId,
  };

  try {
    isLoading.value = true
    await props.onSubmit(payload);
  } catch (error) {
    console.error(error);
    throw error;
  } finally {
    isLoading.value = false
  }
}

</script>

<template>
  <UForm :schema="transactionSchema" :state="state" @submit="submitHandler" @error="(e) => {
    console.log(e);
    toast.add({
      title: 'Gagal', description: 'Masih ada data yang tidak sesuai', color: 'error'
    })
  }" class="space-y-4">
    <InputTransactionType v-model="state.type" />
    <div class="grid grid-cols-2 gap-4">
      <InputDate v-model="state.date" />
      <InputTime v-model="state.time" />
    </div>

    <InputCurrency :name="'nominal'" :label="'Nominal'" v-model="state.nominal" />

    <div class="grid grid-cols-2 gap-4">
      <InputCategory v-model="state.categoryId" :transaction-type="state.type" field-name="categoryId" />
      <InputCategory v-model="state.subCategoryId" :transaction-type="state.type" field-name="subCategoryId"
        :parent-id="state.categoryId" />
    </div>

    <div class="grid gap-4"
      :class="{ 'grid-cols-2': state.type === 'transfer', 'grid-cols-1': state.type !== 'transfer' }">
      <InputAsset :name="'assetFrom'" :label="'Dari Aset'" v-model="state.assetFromId" />
      <InputAsset v-if="state.type === 'transfer'" :name="'assetTo'" :label="'Ke Aset'" v-model="state.assetToId" />
      <InputIsHaveTransferFee v-if="state.type === 'transfer'" class="col-span-2" v-model="state.isHaveTransferFee" />
      <InputCurrency :name="'transfer-fee'" :label="'Biaya Transfer'" v-if="state.isHaveTransferFee"
        v-model="state.transferFee" />
      <InputAsset :name="'feeFromAsset'" :label="'Biaya Dari Aset'" v-if="state.isHaveTransferFee"
        v-model="state.feeFromAssetId" />
    </div>

    <InputCreditor v-if="state.type === 'payable'" v-model="state.creditor" />
    <InputDebtor v-if="state.type === 'receivable'" v-model="state.debtor" />

    <InputNote v-model="state.note" />

    <UFormField label="Deskripsi" name="description">
      <UTextarea v-model="state.description" :maxrows="4" autoresize class="w-full" />
    </UFormField>

    <UButton type="submit" label="Simpan" :loading="isLoading" :disabled="isLoading" block />
  </UForm>
</template>