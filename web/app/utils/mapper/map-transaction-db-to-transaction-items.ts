import type { TransactionDb, TransactionItemBody } from "~/types/transaction";

export default (data: TransactionDb | undefined): TransactionItemBody => {
  if (!data) throw new Error("Data tidak ditemukan");

  switch (data.type) {
    case "transfer":
      return transferCase(data);
    default:
      return normalCase(data);
  }
};

const normalCase = (data: TransactionDb): TransactionItemBody => ({
  asetName: data.assetFrom,
  category: data.category,
  date: data.date,
  nominal: data.nominal,
  note: data.note,
  type: data.type,
  description: data.description,
  subCategory: data.subCategory,
});

const transferCase = (data: TransactionDb): TransactionItemBody => ({
  ...normalCase(data),
  asetName: `${data.assetFrom} => ${data.assetTo}`,
});
