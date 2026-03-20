<script setup lang="ts">
const model = defineModel<string>()
</script>

<template>
  <UFormField label="Catatan">
    <UInput v-model="model" class="w-full" />
  </UFormField>
</template>