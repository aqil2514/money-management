import z from "zod";

const baseSchema = z.object({
  name: z.string().min(1, "Nama aset wajib diisi"),
  total: z.number(),
  category: z.string().min(1, "Kategori aset wajib diisi"),
  status: z.string().min(1, "Status aset wajib diisi"),
  description: z.string().min(1, "Deskripsi aset wajib diisi"),
  ownerType: z.string().min(1, "Kepemilikan aset wajib diisi"),
  assetType: z.string().min(1, "Tipe aset wajib diisi"),
  liquidityScore: z.number().min(1, "Tingkat liquiditas aset wajib diisi"),
  currency: z.string().min(1, "Mata uang aset wajib diisi"),
});

export const assetSchema = z.object({
  ...baseSchema.shape,
});

export type AssetSchemaType = z.output<typeof assetSchema>;

export const defaultAssetFormField: AssetSchemaType = {
  assetType: "Cash",
  category: "Tunai",
  currency: "idr",
  description: "",
  liquidityScore: 5,
  name: "",
  ownerType: "Pribadi",
  status: "active",
  total: 0,
};
