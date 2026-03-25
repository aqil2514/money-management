<script lang="ts" setup>
const noteRef = ref("")

const clearSearch = () => {
  noteRef.value = ""
}
</script>

<template>
  <UPopover :popper="{ placement: 'bottom-start' }">
    <UButton size="md" icon="i-lucide-search" :color="noteRef ? 'primary' : 'neutral'" variant="outline"
      class="min-w-40 justify-between rounded-lg">
      <span class="truncate max-w-37.5">
        {{ noteRef || 'Cari Transaksi' }}
      </span>
      <UIcon name="i-lucide-chevron-down" class="w-4 h-4 opacity-50 ml-2" />
    </UButton>

    <template #content>
      <div class="p-3 w-72 flex flex-col gap-2">
        <div class="flex items-center justify-between px-1">
          <span class="text-xs font-medium text-gray-500 uppercase tracking-wider">Pencarian</span>
          <UButton v-if="noteRef" color="error" variant="ghost" icon="i-lucide-x" size="xs" label="Hapus"
            @click="clearSearch" />
        </div>

        <UInput v-model="noteRef" placeholder="Ketik detail transaksi..." icon="i-lucide-search" autofocus
          autocomplete="off">
          <template v-if="noteRef" #trailing>
            <UButton color="neutral" variant="link" icon="i-lucide-circle-x" :padded="false" @click="clearSearch" />
          </template>
        </UInput>

        <p class="text-[10px] text-gray-400 px-1">
          Tekan Enter untuk menerapkan filter
        </p>
      </div>
    </template>
  </UPopover>
</template>