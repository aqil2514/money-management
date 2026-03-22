  <script lang="ts" setup>
  import { serverUrl } from '~/constants/server-url';
  import type { CategoryDb } from '~/types/category';

  const props = defineProps<{ data: CategoryDb | null }>()
  const emit = defineEmits(['close'])
  const toast = useToast()

  const isLoading = ref(false)
  const isOpen = ref(true)

  watch(() => props.data, (newVal) => {
    if (newVal) {
      isOpen.value = true
    } else isOpen.value = false
  })

  const handleUpdateOpen = (open: boolean) => {
    isOpen.value = open
    if (!open) {
      emit('close')
    }
  }

  const deleteHandler = async () => {
    const id = props.data?.id;
    if (!id) return
    try {
      isLoading.value = true
      await $fetch(`${serverUrl}/category/${id}`, {
        method: "DELETE"
      })

      isOpen.value = false
      await refreshNuxtData('fetch-category')
      toast.add({
        title: "Berhasil",
        description: `Kategori berhasil dihapus`
      })
      emit("close")

    } catch (error) {
      console.error(error);
      toast.add({
        title: "Gagal",
        description: "Terjadi kesalahan saat menghapus data",
        color: "error"
      })
      throw error
    } finally {
      isLoading.value = false

    }
  }

</script>

<template>
  <UModal v-model:open="isOpen" @update:open="handleUpdateOpen" title="Hapus Kategori"
    description="Tindakan ini tidak dapat dibatalkan.">
    <template #body>
      <div v-if="data" class="space-y-4">
        <UAlert icon="i-lucide-triangle-alert" color="error" variant="soft" title="Konfirmasi Penghapusan"
          :description="`Apakah Anda yakin ingin menghapus kategori '${data.name}'?`" />

        <div
          class="p-4 rounded-lg bg-gray-50 dark:bg-gray-800/50 border border-gray-200 dark:border-gray-700 space-y-2 text-sm">
          <div class="flex justify-between">
            <span class="text-gray-500">Nama:</span>
            <span class="font-medium">{{ data.name }}</span>
          </div>
          <div class="flex justify-between">
            <span class="text-gray-500">Tipe:</span>
            <UBadge size="sm" variant="subtle">{{ data.type }}</UBadge>
          </div>
          <div v-if="data.description" class="flex flex-col gap-1">
            <span class="text-gray-500">Deskripsi:</span>
            <p class="text-gray-700 dark:text-gray-300 italic">"{{ data.description }}"</p>
          </div>
        </div>
      </div>
    </template>

    <template #footer>
      <div class="flex justify-end gap-3">
        <UButton label="Batal" color="neutral" variant="ghost" @click="handleUpdateOpen(false)" />
        <UButton label="Ya, Hapus Kategori" color="error" icon="i-lucide-trash-2" @click="deleteHandler" />
      </div>
    </template>
  </UModal>
</template>