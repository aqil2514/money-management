<script setup lang="ts">
import type { Time } from '@internationalized/date';

const model = defineModel<Time>()

</script>

<template>
  <UFormField label="Waktu" name="time">
    <UInputTime class="w-full" :hour-cycle="24" v-model="model" />
  </UFormField>
</template>