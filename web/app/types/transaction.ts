export type TransactionType =
  | "income"
  | "expense"
  | "transfer"
  | "payable"
  | "receivable";

export interface TransactionSummary {
  label: string;
  nominal: number;
  textClass: string;
}

export interface TransactionItemHeader {
  date: string; // Ini untuk semua transaksi di tanggal yang sama
  income: number;
  expense: number;
}

export interface TransactionItemBody {
  category: string;
  date: string; // Ini untuk per item
  nominal: number;
  subCategory?: string;
  note: string;
  asetName: string;
  type: TransactionType;
  description?: string;
}

export interface TransactionItem {
  header: TransactionItemHeader;
  body: TransactionItemBody[];
}
