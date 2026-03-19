<script setup lang="ts">
useSeoMeta({
    title: "Dashboard",
    description: "Dashboard keuangan"
})
// Nuxt otomatis melakukan auto-import untuk useTransactionStore
const store = useTransactionStore()

const newName = ref('')
const newAmount = ref(0)

function handleSave() {
    store.addTransaction(newName.value, newAmount.value)
    newName.value = ''
    newAmount.value = 0
}
</script>

<template>
    <div>
        <h1>Saldo Utama: Rp {{ store.totalBalance.toLocaleString() }}</h1>

        <input v-model="newName" placeholder="Nama" />
        <input v-model.number="newAmount" type="number" />
        <button @click="handleSave">Simpan ke Store</button>

        <ul>
            <li v-for="t in store.transactions" :key="t.id">
                {{ t.name }} - Rp {{ t.amount.toLocaleString() }}
            </li>
        </ul>
    </div>
</template>