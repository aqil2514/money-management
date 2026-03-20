<script setup lang="ts">
import { transactionSchema } from "~/schemas/transaction-schema";
import InputCategory from "./InputCategory.vue";
import InputDate from "./InputDate.vue"
import InputSubCategory from "./InputSubCategory.vue";
import InputTime from "./InputTime.vue";
import InputTransactionType from "./InputTransactionType.vue";
import InputCreditor from "./InputCreditor.vue";
import InputDebtor from "./InputDebtor.vue";
import InputNote from "./InputNote.vue";
import InputIsHaveTransferFee from "./InputIsHaveTransferFee.vue";
import InputCurrency from "./InputCurrency.vue";
import InputAsset from "./InputAsset.vue";
const { state, onSubmit } = useTransactionForm()

</script>

<template>
  <UForm :schema="transactionSchema" :state="state" @submit="onSubmit" class="space-y-4">
    <InputTransactionType v-model="state.type" />
    <div class="grid grid-cols-2 gap-4">
      <InputDate v-model="state.date" />
      <InputTime v-model="state.time" />
    </div>

    <InputCurrency :name="'nominal'" :label="'Nominal'" v-model="state.nominal" />

    <div class="grid grid-cols-2 gap-4">
      <InputCategory v-model="state.category" />
      <InputSubCategory v-model="state.subCategory" />
    </div>

    <div class="grid gap-4"
      :class="{ 'grid-cols-2': state.type === 'transfer', 'grid-cols-1': state.type !== 'transfer' }">
      <InputAsset :name="'asset-from'" :label="'Dari Aset'" v-model="state.assetFrom" />
      <InputAsset :name="'asset-to'" :label="'Ke Aset'" v-model="state.assetTo" />
      <InputIsHaveTransferFee v-if="state.type === 'transfer'" class="col-span-2" v-model="state.isHaveTransferFee" />
      <InputCurrency :name="'transfer-fee'" :label="'Biaya Transfer'" v-if="state.isHaveTransferFee"
        v-model="state.transferFee" />
      <InputAsset :name="'fee-asset-from'" :label="'Biaya Dari Aset'" v-if="state.isHaveTransferFee"
        v-model="state.feeFromAsset" />
    </div>

    <InputCreditor v-if="state.type === 'payable'" v-model="state.creditor" />
    <InputDebtor v-if="state.type === 'receivable'" v-model="state.debtor" />

    <InputNote v-model="state.note" />

    <UFormField label="Deskripsi" name="description">
      <UTextarea v-model="state.description" :maxrows="4" autoresize class="w-full" />
    </UFormField>

    <UButton type="submit" label="Simpan" block />
  </UForm>
</template>