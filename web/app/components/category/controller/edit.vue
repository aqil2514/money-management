<script lang="ts" setup>
import type { CategoryDb } from '~/types/category';
import CategoryForm from '../forms/category-form.vue';
import type { FormSubmitEvent } from '@nuxt/ui';
import type { CategorySchemaType } from '~/schemas/category-schema';
import { serverUrl } from '~/constants/server-url';

const props = defineProps<{ data: CategoryDb | null }>()
const emit = defineEmits(['close'])

const isOpen = ref(false)
const isLoading = ref(false)
const toast = useToast()

watch(() => props.data, (newVal) => {
  if (newVal) {
    isOpen.value = true
  }
})

const handleUpdateOpen = (open: boolean) => {
  isOpen.value = open
  if (!open) {
    emit('close')
  }
}

const handleEdit = async (values: FormSubmitEvent<CategorySchemaType>) => {
  const id = props.data?.id;
  if (!id) return;
  try {
    isLoading.value = true
    await $fetch(`${serverUrl}/category/${id}`, {
      method: "PUT",
      body: values.data
    })

    toast.add({
      title: "Berhasil",
      description: "Data berhasil diupdate"
    })

    isOpen.value = false
    await refreshNuxtData("fetch-category")
    emit("close")
  } catch (error) {
    console.error(error)
    throw error
  } finally {
    isLoading.value = false
  }
}

</script>

<template>
  <UModal v-model:open="isOpen" @update:open="handleUpdateOpen" title="Edit Kategori"
    description="Lengkapi data di bawah ini untuk edit data">
    <template #body>
      <CategoryForm v-if="data" :default-values="data as any" :loading="isLoading" :submit-handler="handleEdit" />
    </template>
  </UModal>
</template>