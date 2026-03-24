<script setup lang="ts">
import { type SelectItem } from '@nuxt/ui';
import { transactionProviderKey } from '~/types/injectKey';

const model = defineModel<number>()
defineProps<{
  label: string,
  name: string
}>()

const injectedData = inject(transactionProviderKey)

const items = computed<SelectItem[]>(() => {
  const data = injectedData?.resources.assets.value ?? []
  const formatted: SelectItem[] = data.map((val) => ({
    label: val.name,
    value: val.ID
  }))

  return formatted
}

)


function onCreate(item: number) {
  items.value.push(item)

  model.value = item
}
</script>

<template>
  <UFormField :label="label" :name="name">
    <USelectMenu v-model="model" value-key="value" create-item :items="items" class="w-full"
      @create="(item) => onCreate(Number(item))" />
  </UFormField>
</template>