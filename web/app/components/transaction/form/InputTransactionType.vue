<script setup lang="ts">
import type { RadioGroupItem } from '@nuxt/ui/runtime/components/RadioGroup.vue.js';
import { transactionType } from '~/constants/transaction-type';

const model = defineModel<string>()
</script>

<template>
    <UFormField label="Tipe Transaksi" name="transaction-type">
        <URadioGroup orientation="horizontal" :items="transactionType as any" v-model="model" />
    </UFormField>
</template>