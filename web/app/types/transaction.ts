export type TransactionType =
  | "income"
  | "expense"
  | "transfer"
  | "payable"
  | "receivable";

export interface TransactionDb {
  assetFrom: string;
  assetFromCategory: string;
  assetTo: string;
  assetToCategory: string;
  category: string;
  createdAt: string;
  creditor: string;
  date: string;
  debtor: string;
  description: string;
  feeFromAsset: string;
  feeFromAssetCategory: string;
  id: string;
  isHaveTransferFee: boolean;
  nominal: number;
  note: string;
  subCategory: string;
  transferFee: number;
  type: TransactionType;
  updatedAt: string;
}

export interface TransactionSummary {
  label: string;
  nominal: number;
  color: string;
}

export interface TransactionItemHeader {
  date: string; // Ini untuk semua transaksi di tanggal yang sama
  income: number;
  expense: number;
}

export interface TransactionItemBody {
  category: string;
  date: string;
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
