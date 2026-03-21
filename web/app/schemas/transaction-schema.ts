import { CalendarDate, Time } from "@internationalized/date";
import z from "zod";

const baseSchema = z.object({
  date: z.instanceof(CalendarDate),
  time: z.instanceof(Time),
  type: z.enum(["expense", "income", "transfer", "payable", "receivable"]),
  nominal: z.number().min(1, "Nominal tidak boleh 0"),
  category: z.string(),
  subCategory: z.string().optional(),
  assetFrom: z.string(),
  note: z.string(),
  description: z.string().optional(),
});

const transferSchema = z.object({
  assetTo: z.string().optional(),
  isHaveTransferFee: z.boolean().optional(),
  transferFee: z.number().optional(),
  feeFromAsset: z.string().optional(),
});

const payableSchema = z.object({
  debtor: z.string().optional(),
});

const receivableSchema = z.object({
  creditor: z.string().optional(),
});

const schema = z.object({
  ...baseSchema.shape,
  ...transferSchema.shape,
  ...payableSchema.shape,
  ...receivableSchema.shape,
});

export const transactionSchema = schema;

export type TransactionSchemaType = z.output<typeof schema>;

const now = new Date();

export const defaultTransaction: TransactionSchemaType = {
  date: markRaw(
    new CalendarDate(now.getFullYear(), now.getMonth() + 1, now.getDate()),
  ),
  time: markRaw(new Time(now.getHours(), now.getMinutes())),
  type: "expense",
  nominal: 0,
  category: "",
  subCategory: undefined,
  note: "",
  description: "",

  assetFrom: "Tunai",
  assetTo: undefined,
  debtor: undefined,
  creditor: undefined,
  isHaveTransferFee: undefined,
  transferFee: undefined,
  feeFromAsset: undefined,
};
