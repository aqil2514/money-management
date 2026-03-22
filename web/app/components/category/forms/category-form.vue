<script setup lang="ts">
import type { FormSubmitEvent, SelectItem } from '@nuxt/ui';
import { transactionType } from '~/constants/transaction-type';
import { categorySchema, type CategorySchemaType } from '~/schemas/category-schema';
import ParentSelect from './parent-select.vue';

const { state } = useCategoryForm()

const props = defineProps<{
  defaultValues?: CategorySchemaType,
  submitHandler: (values: FormSubmitEvent<CategorySchemaType>) => void | Promise<void>
  loading: boolean
}>()

watch(() => props.defaultValues, (newVal) => {
  if (newVal) {
    Object.assign(state, newVal)
  }
}, { immediate: true })

const onSubmit = async (values: FormSubmitEvent<CategorySchemaType>) => {
  await props.submitHandler(values)

  if (!props.defaultValues) {
    state.description = "";
    state.name = "";
    state.parentId = undefined;
    state.type = "expense"
  }
}

</script>

<template>
  <UForm :state="state" @submit="onSubmit" @error="(e) => console.log(e)" :schema="categorySchema" class="space-y-4">
    <UFormField label="Nama Kategori" name="name">
      <UInput v-model="state.name" class="w-full" />
    </UFormField>
    <UFormField label="Deskripsi Kategori" name="description">
      <UTextarea v-model="state.description" class="w-full" />
    </UFormField>
    <UFormField label="Tipe Kategori" name="type">
      <USelect v-model="state.type" :items="[...transactionType]" class="w-full" />
    </UFormField>
    <ParentSelect v-model="state.parentId" />
    <UButton type="submit" :disabled="loading" :loading="loading">
      {{ loading ? "Menyimpan..." : "Simpan" }}
    </UButton>
  </UForm>
</template>