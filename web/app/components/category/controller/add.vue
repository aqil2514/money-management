<script setup lang="ts">
import type { FormSubmitEvent } from '@nuxt/ui';
import type { CategorySchemaType } from '~/schemas/category-schema';
import { serverUrl } from '~/constants/server-url';
import CategoryForm from '../forms/category-form.vue';

const open = ref(false)
const loading = ref(false)
const toast = useToast()

const addHandler = async (values: FormSubmitEvent<CategorySchemaType>) => {
  try {
    loading.value = true

    await $fetch(`${serverUrl}/category`, {
      method: "POST",
      body: values.data
    })

    toast.add({
      title: "Berhasil",
      description: "Tambah kategori baru berhasil"
    })

    await refreshNuxtData("fetch-category")
    open.value = false
  } catch (error) {
    console.error(error)
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <UButton class="cursor-pointer mt-4" @click="open = true">
    Tambah
  </UButton>


  <UModal v-model:open="open" title="Tambah Kategori" description="Isi data di bawah ini untuk menambah kategori baru">
    <template #body>
      <CategoryForm :loading="loading" :submit-handler="addHandler" />
    </template>
  </UModal>
</template>