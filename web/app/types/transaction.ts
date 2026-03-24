export type TransactionType =
  | "income"
  | "expense"
  | "transfer"
  | "payable"
  | "receivable";

export interface TransactionDb {
  // Identitas Utama
  id: string;
  type: TransactionType;
  nominal: number;
  note: string;
  description: string;
  date: string; // ISO String dari Backend

  // Foreign Key IDs (Penting untuk Form Edit)
  categoryId: string;
  subCategoryId: string | null;
  assetFromId: number;
  assetToId: number | null;
  feeFromAssetId: number | null;

  // Nama/Label (Untuk Tampilan Detail)
  category: string;
  subCategory: string;
  assetFrom: string;
  assetFromCategory: string;
  assetTo: string;
  assetToCategory: string;
  feeFromAsset: string;
  feeFromAssetCategory: string;

  // Data Tambahan
  isHaveTransferFee: boolean;
  transferFee: number;
  debtor: string;
  creditor: string;

  // Timestamps
  createdAt: string;
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
  id: string;
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
