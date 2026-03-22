import z from "zod";

const baseSchema = z.object({
  name: z.string().min(1, "Nama kategori wajib diisi"),
  description: z.string().min(1, "Deskripsi kategori wajib diisi"),
  type: z.enum(
    ["income", "expense", "payable", "receivable", "transfer"],
    "Tipe kategori tidak valid",
  ),
  parentId: z.string().optional().nullable(),
});

export const categorySchema = z.object({
  ...baseSchema.shape,
});

export type CategorySchemaType = z.output<typeof categorySchema>;

export const defaultCategory: CategorySchemaType = {
  name: "",
  description: "",
  type: "expense",
  parentId: undefined,
};
