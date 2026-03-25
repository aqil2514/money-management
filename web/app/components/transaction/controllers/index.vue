<script setup lang="ts">
import ControllerAdd from '../dialogs/add.vue';
import SearchNote from './search-note.vue';

</script>

<template>
  <div class="flex justify-end gap-4 items-center">
    <SearchNote />
    <ControllerAdd />
  </div>
</template>