<script setup lang="ts">
const model = defineModel<string>()
defineProps<{
  label: string,
  name: string
}>()
const items = ref(['Tunai', 'Seabank', 'Kolega', 'Done'])

function onCreate(item: string) {
  items.value.push(item)

  model.value = item
}
</script>

<template>
  <UFormField :label="label" :name="name">
    <USelectMenu v-model="model" create-item :items="items" class="w-full" @create="onCreate" />
  </UFormField>
</template>