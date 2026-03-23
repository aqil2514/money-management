<script lang="ts" setup>
import type { FormSubmitEvent, RadioGroupItem } from '@nuxt/ui';
import { assetSchema, defaultAssetFormField, type AssetSchemaType } from '~/schemas/asset-schema';
import { assetProvideKey, type AssetProvideInjectKey } from '~/types/injectKey';

const props = defineProps<{
  submitHandler: (values: AssetSchemaType) => Promise<void> | void,
  defaultValues?: AssetSchemaType,
}>()

const state = reactive<AssetSchemaType>({ ...defaultAssetFormField })
const toast = useToast()
const isLoading = ref(false)

const injectData = inject<AssetProvideInjectKey>(assetProvideKey)

watch(() => props.defaultValues, (newValues) => {
  if (newValues)
    Object.assign(state, { ...newValues })
}, { immediate: true })

const onSubmit = async (values: FormSubmitEvent<AssetSchemaType>) => {
  try {
    isLoading.value = true
    await props.submitHandler(values.data)

    if (!props.defaultValues)
      Object.assign(state, { ...defaultAssetFormField })
  } catch (error) {
    console.error(error)
    toast.add({ title: "Gagal", description: "Terjadi kesalahan saat tambah data" })
    throw error
  } finally {
    isLoading.value = false
  }
}

const statusItem = injectData?.constant.statusItem.value ?? []
const currencyItem = injectData?.constant.currencyItem.value ?? []
const liquidityScoreMapping = injectData?.constant.liquidityScoreMapping ?? {}

</script>

<template>
  <UForm :state="state" :schema="assetSchema" @submit="onSubmit" @error="(e) => console.error(e)" class="space-y-4">
    <UFormField label="Nama Aset" name="name">
      <UInput v-model="state.name" class="w-full" />
    </UFormField>

    <UFormField label="Total Aset" name="total">
      <UInputNumber v-model="state.total" class="w-full" :format-options="{
        style: 'currency',
        currency: 'IDR',
        currencyDisplay: 'narrowSymbol',
        maximumFractionDigits: 0
      }" />
    </UFormField>

    <UFormField label="Kategori" name="category" help="Misal : Tunai, Digital, Bank">
      <UInput v-model="state.category" class="w-full" />
    </UFormField>

    <UFormField label="Status" name="status">
      <URadioGroup orientation="horizontal" v-model="state.status" :items="statusItem" class="w-full" />
    </UFormField>

    <UFormField label="Deskripsi" name="description">
      <UTextarea v-model="state.description" class="w-full" />
    </UFormField>

    <UFormField label="Jenis Kepemilikan" name="ownerType">
      <UInput v-model="state.ownerType" class="w-full" />
    </UFormField>

    <UFormField label="Tipe Aset" name="assetType">
      <UInput v-model="state.assetType" class="w-full" />
    </UFormField>

    <UFormField label="Skor Likuiditas"
      :help="`${state.liquidityScore} - ${liquidityScoreMapping[state.liquidityScore]}`">
      <USlider v-model="state.liquidityScore" :max="5" :min="1" :tooltip="{
        text: `${state.liquidityScore} - ${liquidityScoreMapping[state.liquidityScore]}`
      }" class="w-full" />
    </UFormField>

    <UFormField label="Mata Uang" name="currency">
      <URadioGroup orientation="horizontal" v-model="state.currency" :items="currencyItem" class="w-full" />
    </UFormField>

    <UButton :loading="isLoading" :disabled="isLoading" type="submit"> {{ isLoading ? 'Menyimpan...' : 'Simpan' }}
    </UButton>
  </UForm>
</template>