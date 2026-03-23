<script lang="ts" setup>
import { assetProvideKey } from '~/types/injectKey';


const state = inject(assetProvideKey)

</script>

<template>
  <div class="flex justify-end gap-4">
    <UButton @click="() => state?.updateModalOpen('add')">
      Tambah Aset
    </UButton>
  </div>
</template>