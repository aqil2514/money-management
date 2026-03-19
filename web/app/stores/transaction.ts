import { defineStore } from "pinia";

export const useTransactionStore = defineStore("transaction", () => {
  const transactions = ref([{ id: 1, name: "Saldo Awal", amount: 1000000 }]);

  const totalBalance = computed(() => {
    return transactions.value.reduce((acc, item) => acc + item.amount, 0);
  });

  function addTransaction(name: string, amount: number) {
    transactions.value.push({
      id: Date.now(),
      name,
      amount,
    });
  }

  return { addTransaction, totalBalance, transactions };
});
