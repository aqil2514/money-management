<script setup lang="ts">
const model = defineModel < boolean > ()
</script>

<template>
  <USwitch label="Ada Biaya Transfer?" :description="model ? 'Ya, ada biaya transfer' : 'Tidak ada biaya transfer'"
    v-model="model" />
</template>