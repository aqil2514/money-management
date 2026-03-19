<script setup lang="ts">
const model = defineModel<number>()

defineProps<{
    placeholder?: string
}>()
</script>

<template>
    <UInputNumber v-model="model" :placeholder="placeholder" class="w-full" :format-options="{
        style: 'currency',
        currency: 'IDR'
    }" />
</template>