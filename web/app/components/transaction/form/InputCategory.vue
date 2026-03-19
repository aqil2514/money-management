<script setup lang="ts">
const model = defineModel<string>()
const items = ref(['Backlog', 'Todo', 'In Progress', 'Done'])

function onCreate(item: string) {
  items.value.push(item)

  model.value = item
}
</script>

<template>
  <USelectMenu v-model="model" create-item :items="items" class="w-full" @create="onCreate" />
</template>