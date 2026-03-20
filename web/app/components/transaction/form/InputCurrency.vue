<script setup lang="ts">
const model = defineModel<number>()
defineProps<{
    label: string,
    name: string
}>()

</script>

<template>
    <UFormField :label="label" :name="name">
        <UInputNumber v-model="model" class="w-full" :format-options="{
            style: 'currency',
            currency: 'IDR',
            currencyDisplay: 'narrowSymbol',
            maximumFractionDigits: 0
        }" />
    </UFormField>
</template>