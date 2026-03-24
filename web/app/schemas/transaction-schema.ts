import { CalendarDate, Time } from "@internationalized/date";
import z, { core } from "zod";

const baseSchema = z.object({
  date: z.instanceof(CalendarDate),
  time: z.instanceof(Time),
  type: z.enum(["expense", "income", "transfer", "payable", "receivable"]),
  nominal: z.number().min(1, "Nominal harus lebih dari 0"),
  categoryId: z.string().min(1, "Kategori wajib diisi"),
  subCategoryId: z.string().nullish(),
  assetFromId: z.number().min(1, "Aset asal wajib diisi"),
  note: z.string().min(1, "Catatan wajib diisi"),
  description: z.string().optional(),
});

const transferSchema = z.object({
  assetToId: z.number().optional(),
  isHaveTransferFee: z.boolean().default(false),
  transferFee: z.number().optional(),
  feeFromAssetId: z.number().optional(),
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

const transferRefine = (
  value: TransactionSchemaType,
  ctx: core.$RefinementCtx<TransactionSchemaType>,
) => {
  const { assetToId, isHaveTransferFee, feeFromAssetId, transferFee } = value;
  if (!assetToId)
    ctx.addIssue({
      code: "custom",
      message: "Aset tujuan wajib disertakan",
      path: ["assetTo"],
    });

  if (isHaveTransferFee && !feeFromAssetId)
    ctx.addIssue({
      code: "custom",
      message: "Uang keluar aset wajib diisi",
      path: ["feeFromAsset"],
    });

  if (isHaveTransferFee && (!transferFee || transferFee <= 0)) {
    ctx.addIssue({
      code: "custom",
      message: "Biaya transfer harus lebih dari 0",
      path: ["transferFee"],
    });
  }
};

export const transactionSchema = schema.superRefine((value, ctx) => {
  if (value.type === "transfer") transferRefine(value, ctx);
});

export type TransactionSchemaType = z.output<typeof schema>;
export type TransactionRequestPayload = Omit<TransactionSchemaType, "date"> & {
  date: string;
  id?: string; //Buat edit
};
export type TransactionSchemaTypeKey = keyof TransactionSchemaType;

const now = new Date();

export const defaultTransaction: TransactionSchemaType = {
  date: markRaw(
    new CalendarDate(now.getFullYear(), now.getMonth() + 1, now.getDate()),
  ),
  time: markRaw(new Time(now.getHours(), now.getMinutes())),
  type: "expense",
  nominal: 0,
  categoryId: "no-category",
  subCategoryId: "no-subcategory",
  note: "",
  description: "",

  assetFromId: 1,
  assetToId: undefined,
  debtor: undefined,
  creditor: undefined,
  isHaveTransferFee: false,
  transferFee: undefined,
  feeFromAssetId: undefined,
};
