<script setup lang="ts">

const model = defineModel<string | undefined>()
const items = ref(['No Subcategory', 'Backlog', 'Todo', 'In Progress', 'Done'])

watch(model, (newValue) => {
  if (newValue === "No Subcategory") {
    model.value = undefined
  }
})

function onCreate(item: string) {
  items.value.push(item)
  model.value = item
}
</script>

<template>
  <UFormField label="Sub Kategori" name="subCategory">
    <USelectMenu v-model="model" :items="items" create-item class="w-full" @create="onCreate" />
  </UFormField>
</template>