<script setup lang="ts">
const model = defineModel<string>()
</script>

<template>
  <UFormField label="Pihak yang berhutang">
    <UInput v-model="model" class="w-full" />
  </UFormField>
</template>