<script setup lang="ts">
import { type SelectItem } from '@nuxt/ui';
import { transactionProviderKey } from '~/types/injectKey'
import type { TransactionType } from '~/types/transaction';

const model = defineModel<string | undefined | null>()
const injectedData = inject(transactionProviderKey);
const props = defineProps<{
  fieldName: "categoryId" | "subCategoryId",
  transactionType: TransactionType,
  parentId?: string;
}>()
const isCategory = props.fieldName === "categoryId"

const items = computed<SelectItem[]>(() => {
  const data = injectedData?.resources.categories.value ?? [];
  const formatted: SelectItem[] = data.filter((d) => d.parentId === null && d.type === props.transactionType).map((d) => ({
    label: d.name,
    value: d.id
  }))

  const defaultValue: SelectItem[] = [
    {
      value: "no-category",
      label: "Pilih Kategori"
    }
  ]

  return [...defaultValue, ...formatted]
})

const subItems = computed<SelectItem[]>(() => {
  const data = injectedData?.resources.categories.value ?? [];
  const formatted: SelectItem[] = data.filter((d) => d.parentId === null && d.type === props.transactionType).map((d) => ({
    label: d.name,
    value: d.id
  }))

  const subCategory: SelectItem[] = [
    {
      label: "Tidak ada Subkategori",
      value: "no-subcategory"
    }
  ]

  return [...subCategory, ...formatted]
})

function onCreate(item: string) {
  items.value.push(item)

  model.value = item
}

</script>

<template>
  <UFormField :label="isCategory ? 'Kategori' : 'Sub Kategori'" :name="props.fieldName">
    <USelectMenu v-model="model" value-key="value" create-item :items="isCategory ? items : subItems" class="w-full"
      @create="onCreate" />
  </UFormField>
</template>