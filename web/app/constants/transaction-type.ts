/**
 *
 * Komponen yang makek
 *
 * 1. @link ('app\components\transaction\form\InputTransactionType.vue')
 *
 */

export const transactionType = [
  {
    label: "Pengeluaran",
    value: "expense",
  },
  {
    label: "Pemasukan",
    value: "income",
  },
  {
    label: "Transfer",
    value: "transfer",
  },
  {
    label: "Utang",
    value: "payable",
  },
  {
    label: "Piutang",
    value: "receivable",
  },
] as const;
